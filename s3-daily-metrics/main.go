package main

import (
	"context"
	"fmt"
	metricsExporter "github.com/logzio/go-metrics-sdk"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	envLogzioMetricsListener = "LOGZIO_METRICS_LISTENER"
	envLogzioMetricsToken    = "LOGZIO_METRICS_TOKEN"
	fieldLogzioAgentVersion  = "logzio_agent_version"
	unitBytes                = "Bytes"
	unitCount                = "Count"
	metricNameNumObjects     = "NumberOfObjects"
	metricNameSizeBytes      = "BucketSizeBytes"
)

func main() {
	//lambda.Start(handleRequest)
	request, err := handleRequest(context.Background())
	if err != nil {
		print(request)
	}
}

// handleRequest is the entry point for the Lambda function
func handleRequest(ctx context.Context) (string, error) {
	exporter, err := configureMetricsExporter()
	if err != nil {
		return "", err
	}
	defer func() {
		handleErr(exporter.Stop(ctx))
	}()
	meter := exporter.Meter("aws_s3")
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		fmt.Printf(err.Error())
	}
	// Create a new CloudWatch client
	cw := cloudwatch.New(sess)
	// Create a new S3 client
	s3Client := s3.New(sess)
	// List all the buckets in the S3 namespace
	buckets, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err.Error(), err
	}
	// Iterate through the buckets and get the NumberOfObjects and BucketSizeBytes metrics for each bucket
	for _, bucket := range buckets.Buckets {
		fmt.Println(*bucket.Name)
		// Get the NumberOfObjects metric for the bucket
		NumberOfObjectsErr := collectCloudwatchMetric(metricNameNumObjects, unitCount, bucket, ctx, &meter, cw)
		if NumberOfObjectsErr != nil {
			return NumberOfObjectsErr.Error(), NumberOfObjectsErr
		}
		// Get the BucketSizeBytes metric for the bucket
		BucketSizeBytesErr := collectCloudwatchMetric(metricNameSizeBytes, unitBytes, bucket, ctx, &meter, cw)
		if BucketSizeBytesErr != nil {
			return BucketSizeBytesErr.Error(), BucketSizeBytesErr
		}
	}
	fmt.Println("Done")
	return "Success", nil
}

// collectCloudwatchMetric Collects a Cloudwatch metric for a given bucket and metric name
func collectCloudwatchMetric(name string, unit string, bucket *s3.Bucket, ctx context.Context, meter *metric.Meter, cw *cloudwatch.CloudWatch) error {
	// Get the NumberOfObjects metric for the bucket
	cloudwatchMetric, err := cw.GetMetricData(&cloudwatch.GetMetricDataInput{
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
								Value: bucket.Name,
							},
							{
								Name:  aws.String("StorageType"),
								Value: aws.String("AllStorageTypes"),
							},
						},
					},
					Period: aws.Int64(86400),
					Stat:   aws.String("Maximum"),
					Unit:   aws.String(unit),
				},
			},
		},
		// Set the start time to 1 day ago
		StartTime: aws.Time(time.Now().Add(-48 * time.Hour)),
		EndTime:   aws.Time(time.Now()),
	})
	if err != nil {
		return err
	}
	if len(cloudwatchMetric.MetricDataResults) > 0 {
		if len(cloudwatchMetric.MetricDataResults[0].Values) > 0 {
			metricValue := int64(*cloudwatchMetric.MetricDataResults[0].Values[0])
			prometheusCloudwacthMetric := metric.Must(*meter).NewInt64UpDownCounter("aws_s3_" + strings.ToLower(name) + "_max")
			bucketNameAtt := attribute.KeyValue{
				Key:   "bucketname",
				Value: attribute.StringValue(*bucket.Name),
			}
			prometheusCloudwacthMetric.Add(ctx, metricValue, bucketNameAtt)
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

// handleErr Handles defer errors
func handleErr(err error) {
	if err != nil {
		fmt.Println("encountered error: ", err)
	}
}
