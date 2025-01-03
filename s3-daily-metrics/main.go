package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/s3"
	metricsExporter "github.com/logzio/go-metrics-sdk"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric/controller/basic"
)

const (
	// envLogzioMetricsListener The environment variable name for the Logz.io listener
	envLogzioMetricsListener = "LOGZIO_METRICS_LISTENER"
	// envLogzioMetricsToken The environment variable name for the Logz.io token
	envLogzioMetricsToken = "LOGZIO_METRICS_TOKEN"
	envP8slogzioName      = "P8S_LOGZIO_NAME"
	// fieldLogzioAgentVersion The field name for the Logz.io agent version
	fieldLogzioAgentVersion      = "logzio_agent_version"
	fieldLogzioAgentVersionValue = "1.0.0"
	// fieldP8slogzioName
	fieldP8slogzioName      = "p8s_logzio_name"
	fieldP8slogzioNameValue = "cloudwatch-helpers"
	// fieldNameSpace The field name for the namespace
	fieldNameSpace      = "namespace"
	fieldNameSpaceValue = "aws/s3"
	// fieldBucketName The field name for the bucket name
	fieldBucketName = "bucketname"
	// fieldRegion The field name for the region
	fieldRegion = "region"
	// fieldAccount The field name for the aws account
	fieldAccount         = "account"
	unitBytes            = "Bytes"
	unitCount            = "Count"
	metricNameNumObjects = "NumberOfObjects"
	metricNameSizeBytes  = "BucketSizeBytes"
	storageTypeStandard  = "StandardStorage"
	storageTypeAll       = "AllStorageTypes"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest is the entry point for the Lambda function
func HandleRequest(ctx context.Context, event cfn.Event) (string, error) {
	//
	if event.RequestID == "" {
		fmt.Println("Scheduled invocation")
		// Scheduled invocation
		_, err := s3DailyMetricsHandler(ctx)
		if err != nil {
			fmt.Printf("Encountered an error: %s", err.Error())
			return "Encountered an error, lambda finished with error", err
		}
	} else {
		if event.RequestType == cfn.RequestCreate {
			fmt.Println("Custom resource invocation")
			// Custom resource invocation
			lambda.Start(cfn.LambdaWrap(customResourceHandler))
		} else {
			lambda.Start(cfn.LambdaWrap(customResourceDoNothing))
		}
	}
	return "Lambda finished", nil
}

// Wrapper for Read, Update, Delete requests
func customResourceDoNothing(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	return
}

// Wrapper for first invocation from cloud formation custom resource
func customResourceHandler(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	fmt.Println("Starting customResourceHandler")
	_, err = s3DailyMetricsHandler(ctx)
	if err != nil {
		fmt.Printf("Encountered an error: %s", err.Error())
		return
	}
	return
}

// s3DailyMetricsHandler Handles the scheduled invocation
func s3DailyMetricsHandler(ctx context.Context) (string, error) {
	fmt.Println("Starting s3DailyMetricsHandler")
	exporter, err := configureMetricsExporter()
	if err != nil {
		return "", err
	}
	defer func() {
		handleErr(exporter.Stop(ctx))
	}()
	meter := exporter.Meter("aws_s3")
	// Create a new AWS session
	fmt.Println("Creating new AWS session")
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		fmt.Printf(err.Error())
	}
	// Create a new CloudWatch client
	fmt.Println("Creating new CloudWatch client")
	cw := cloudwatch.New(sess)
	// Create a new S3 client
	fmt.Println("Creating new S3 client")
	s3Client := s3.New(sess)
	// List all the buckets in the S3 namespace
	fmt.Println("Listing all the buckets in the S3 namespace")
	buckets, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err.Error(), err
	}
	// Iterate through the buckets and get the NumberOfObjects and BucketSizeBytes metrics for each bucket
	fmt.Println("Iterating through the buckets and get the NumberOfObjects and BucketSizeBytes metrics for each bucket")
	for _, bucket := range buckets.Buckets {
		// Get the NumberOfObjects metric for the bucket
		NumberOfObjectsErr := collectCloudwatchMetric(metricNameNumObjects, unitCount, storageTypeAll, bucket, ctx, &meter, cw)
		if NumberOfObjectsErr != nil {
			fmt.Printf("Error while collecting NumberOfObjects metric for bucket %s: %s", *bucket.Name, NumberOfObjectsErr.Error())
		}
		// Get the BucketSizeBytes metric for the bucket
		BucketSizeBytesErr := collectCloudwatchMetric(metricNameSizeBytes, unitBytes, storageTypeStandard, bucket, ctx, &meter, cw)
		if BucketSizeBytesErr != nil {
			fmt.Printf("Error while collecting BucketSizeBytes metric for bucket %s: %s", *bucket.Name, BucketSizeBytesErr.Error())
		}
	}
	return "Success", nil
}

// collectCloudwatchMetric Collects a Cloudwatch metric for a given bucket and metric name
func collectCloudwatchMetric(name string, unit string, storageType string, bucket *s3.Bucket, ctx context.Context, meter *metric.Meter, cw *cloudwatch.CloudWatch) error {
	var cloudwatchMetric *cloudwatch.GetMetricDataOutput
	var err error
	var backoff = 1
	// retry logic
	for i := 0; i < 3; i++ {
		cloudwatchMetric, err = cw.GetMetricData(&cloudwatch.GetMetricDataInput{
			MetricDataQueries: []*cloudwatch.MetricDataQuery{
				{
					Id: aws.String("m1"),
					MetricStat: &cloudwatch.MetricStat{
						Metric: &cloudwatch.Metric{
							Namespace:  aws.String("AWS/S3"),
							MetricName: aws.String(name),
							Dimensions: []*cloudwatch.Dimension{
								{
									Name:  aws.String("BucketName"),
									Value: aws.String(*bucket.Name),
								},
								{
									Name:  aws.String("StorageType"),
									Value: aws.String(storageType),
								},
							},
						},
						Period: aws.Int64(86400),
						Stat:   aws.String("Maximum"),
						Unit:   aws.String(unit),
					},
				},
			},
			// Set the start time to 2 day ago
			StartTime: aws.Time(time.Now().Add(-48 * time.Hour)),
			EndTime:   aws.Time(time.Now()),
		})
		// No error, break out of the loop
		if err == nil {
			break
		}
		// Check for 4XX errors and cancel the retry if we encounter one
		if strings.Contains(err.Error(), "403") {
			fmt.Printf("Received 403 error, Retry aborted. Error: %s", err.Error())
			break
		}
		if strings.Contains(err.Error(), "404") {
			fmt.Printf("Received 404 error, Retry aborted. Error: %s", err.Error())
			break
		}
		if strings.Contains(err.Error(), "400") {
			fmt.Printf("Received 400 error, Retry aborted. Error: %s", err.Error())
			break
		}
		// Handle non 4XX errors with retry logic
		fmt.Printf("Error while performing GetMetricData api call for bucket %s. Trying again in %v seconds. Error: %s", *bucket.Name, backoff, err.Error())
		// wait before retrying (exponential backoff)
		time.Sleep(time.Duration(backoff) * time.Second)
		backoff = backoff * 2
	}
	if err != nil {
		fmt.Printf("GetMetricData api call for bucket %s failed. Error: %s", *bucket.Name, err.Error())
		return err
	}
	if len(cloudwatchMetric.MetricDataResults) > 0 {
		if len(cloudwatchMetric.MetricDataResults[0].Values) > 0 {
			// Attributes for the metric
			attributes := make([]attribute.KeyValue, 0)
			// Add the bucket name as an attribute
			attributes = append(attributes, attribute.KeyValue{
				Key:   fieldBucketName,
				Value: attribute.StringValue(*bucket.Name),
			})
			// Add logzio_agent_version as an attribute
			attributes = append(attributes, attribute.KeyValue{
				Key:   fieldLogzioAgentVersion,
				Value: attribute.StringValue(fieldLogzioAgentVersionValue),
			})
			// Add namespace attribute
			attributes = append(attributes, attribute.KeyValue{
				Key:   fieldNameSpace,
				Value: attribute.StringValue(fieldNameSpaceValue),
			})
			// Add p8s_logzio_name attribute
			envTag := getEnvTag()
			attributes = append(attributes, attribute.KeyValue{
				Key:   fieldP8slogzioName,
				Value: attribute.StringValue(envTag),
			})
			// Add aws region attribute
			attributes = append(attributes, attribute.KeyValue{
				Key:   fieldRegion,
				Value: attribute.StringValue(os.Getenv("AWS_REGION")),
			})
			// Add aws account attribute
			attributes = append(attributes, attribute.KeyValue{
				Key:   fieldAccount,
				Value: attribute.StringValue(os.Getenv("AWS_ACCOUNT_ID")),
			})
			metricValue := int64(*cloudwatchMetric.MetricDataResults[0].Values[0])
			prometheusCloudwatchMetric := metric.Must(*meter).NewInt64UpDownCounter("aws_s3_" + strings.ToLower(name) + "_max")
			prometheusCloudwatchMetric.Add(ctx, metricValue, attributes...)
		}
	}
	return nil
}

// configureMetricsExporter Configures the Logz.io metrics exporter
func configureMetricsExporter() (*basic.Controller, error) {
	listener, err := getListener()
	if err != nil {
		return nil, fmt.Errorf("error while configuring metrics exporter: %s", err.Error())
	}

	token, err := getLogzioToken()
	if err != nil {
		return nil, fmt.Errorf("error while configuring metrics exporter: %s", err.Error())
	}

	config := metricsExporter.Config{
		LogzioMetricsListener: listener,
		LogzioMetricsToken:    token,
		RemoteTimeout:         30 * time.Second,
		PushInterval:          2 * time.Second,
	}

	exporter, err := metricsExporter.InstallNewPipeline(config, basic.WithCollectPeriod(2*time.Second))
	if err != nil {
		return nil, fmt.Errorf("error while configuring metrics exporter: %s", err.Error())
	}
	return exporter, nil
}

// getListener Gets the listener from the environment variables
func getListener() (string, error) {
	listener := os.Getenv(envLogzioMetricsListener)
	if listener == "" {
		return "", fmt.Errorf("%s must be set", envLogzioMetricsListener)
	}

	return listener, nil
}

// getLogzioToken Gets the Logz.io token from the environment variables
func getLogzioToken() (string, error) {
	listener := os.Getenv(envLogzioMetricsToken)
	if listener == "" {
		return "", fmt.Errorf("%s must be set", envLogzioMetricsListener)
	}

	return listener, nil
}

func getEnvTag() string {
	tag := os.Getenv(envP8slogzioName)
	if tag == "" {
		tag = fieldP8slogzioNameValue
	}
	return tag
}

// handleErr Handles defer errors
func handleErr(err error) {
	if err != nil {
		fmt.Println("encountered error: ", err)
	}
}
