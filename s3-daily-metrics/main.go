package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	metricsExporter "github.com/logzio/go-metrics-sdk/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
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

	meterProvider, err := configureMetricsExporter(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		handleErr(meterProvider.Shutdown(shutdownCtx))
	}()

	otel.SetMeterProvider(meterProvider)
	meter := otel.Meter("aws_s3")

	fmt.Println("Creating new AWS session")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		fmt.Printf(err.Error())
		return err.Error(), err
	}

	fmt.Println("Creating new CloudWatch client")
	cw := cloudwatch.NewFromConfig(cfg)

	fmt.Println("Creating new S3 client")
	s3Client := s3.NewFromConfig(cfg)

	fmt.Println("Listing all the buckets in the S3 namespace")
	buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return err.Error(), err
	}

	fmt.Println("Iterating through the buckets and get the NumberOfObjects and BucketSizeBytes metrics for each bucket")
	for _, bucket := range buckets.Buckets {
		NumberOfObjectsErr := collectCloudwatchMetric(metricNameNumObjects, unitCount, storageTypeAll, bucket, ctx, meter, cw)
		if NumberOfObjectsErr != nil {
			fmt.Printf("Error while collecting NumberOfObjects metric for bucket %s: %s", *bucket.Name, NumberOfObjectsErr.Error())
		}
		BucketSizeBytesErr := collectCloudwatchMetric(metricNameSizeBytes, unitBytes, storageTypeStandard, bucket, ctx, meter, cw)
		if BucketSizeBytesErr != nil {
			fmt.Printf("Error while collecting BucketSizeBytes metric for bucket %s: %s", *bucket.Name, BucketSizeBytesErr.Error())
		}
	}
	return "Success", nil
}

// collectCloudwatchMetric Collects a Cloudwatch metric for a given bucket and metric name
func collectCloudwatchMetric(name string, unit string, storageType string, bucket s3types.Bucket, ctx context.Context, meter metric.Meter, cw *cloudwatch.Client) error {
	var cloudwatchMetric *cloudwatch.GetMetricDataOutput
	var err error
	var backoff = 1

	for i := 0; i < 3; i++ {
		cloudwatchMetric, err = cw.GetMetricData(ctx, &cloudwatch.GetMetricDataInput{
			MetricDataQueries: []types.MetricDataQuery{
				{
					Id: aws.String("m1"),
					MetricStat: &types.MetricStat{
						Metric: &types.Metric{
							Namespace:  aws.String("AWS/S3"),
							MetricName: aws.String(name),
							Dimensions: []types.Dimension{
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
						Period: aws.Int32(86400),
						Stat:   aws.String("Maximum"),
						Unit:   types.StandardUnit(unit),
					},
				},
			},
			StartTime: aws.Time(time.Now().Add(-48 * time.Hour)),
			EndTime:   aws.Time(time.Now()),
		})
		if err == nil {
			break
		}
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
		fmt.Printf("Error while performing GetMetricData api call for bucket %s. Trying again in %v seconds. Error: %s", *bucket.Name, backoff, err.Error())
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
			attributes := []attribute.KeyValue{
				// Add the bucket name as an attribute
				attribute.String(fieldBucketName, *bucket.Name),
				// Add logzio_agent_version as an attribute
				attribute.String(fieldLogzioAgentVersion, fieldLogzioAgentVersionValue),
				// Add namespace attribute
				attribute.String(fieldNameSpace, fieldNameSpaceValue),
				// Add p8s_logzio_name attribute
				attribute.String(fieldP8slogzioName, getEnvTag()),
				// Add aws region attribute
				attribute.String(fieldRegion, os.Getenv("AWS_REGION")),
				// Add aws account attribute
				attribute.String(fieldAccount, os.Getenv("AWS_ACCOUNT_ID")),
			}

			metricValue := int64(cloudwatchMetric.MetricDataResults[0].Values[0])

			counter, err := meter.Int64UpDownCounter("aws_s3_" + strings.ToLower(name) + "_max")
			if err != nil {
				return fmt.Errorf("failed to create counter: %w", err)
			}

			counter.Add(ctx, metricValue, metric.WithAttributes(attributes...))
		}
	}
	return nil
}

// configureMetricsExporter Configures the Logz.io metrics exporter
func configureMetricsExporter(ctx context.Context) (*sdkmetric.MeterProvider, error) {
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

	exporter, err := metricsExporter.New(config)
	if err != nil {
		return nil, fmt.Errorf("error while creating metrics exporter: %s", err.Error())
	}

	reader := sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(2*time.Second))
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))

	return meterProvider, nil
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
	token := os.Getenv(envLogzioMetricsToken)
	if token == "" {
		return "", fmt.Errorf("%s must be set", envLogzioMetricsToken)
	}

	return token, nil
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
