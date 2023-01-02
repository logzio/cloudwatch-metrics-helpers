package main

import (
	"context"
	"fmt"
	metricsExporter "github.com/logzio/go-metrics-sdk"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"os"
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
)

func main() {
	//lambda.Start(handleRequest)
	request, err := handleRequest(context.Background())
	if err != nil {
		print(request)
	}
}

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
		return "", err
	}

	// Iterate through the buckets and get the NumberOfObjects and BucketSizeBytes metrics for each bucket
	for _, bucket := range buckets.Buckets {
		// Get the NumberOfObjects metric for the bucket
		numberOfObjectsMetric, err := cw.GetMetricData(&cloudwatch.GetMetricDataInput{
			MetricDataQueries: []*cloudwatch.MetricDataQuery{
				{
					Id: aws.String("m1"),
					MetricStat: &cloudwatch.MetricStat{
						Metric: &cloudwatch.Metric{
							Namespace:  aws.String("AWS/S3"),
							MetricName: aws.String("NumberOfObjects"),
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
						Stat:   aws.String("Average"),
						Unit:   aws.String("Count"),
					},
				},
			},
			StartTime: aws.Time(time.Now().Add(-48 * time.Hour)),
			EndTime:   aws.Time(time.Now()),
		})
		if err != nil {
			return "", err
		}
		if len(numberOfObjectsMetric.MetricDataResults) > 0 {
			numberOfObjects := int64(*numberOfObjectsMetric.MetricDataResults[0].Values[0])
			prometheusNumberOfObjectsMetric := metric.Must(meter).NewInt64UpDownCounter("NumberOfObjects")
			bucketNameAtt := attribute.KeyValue{
				Key:   "bucket_name",
				Value: attribute.StringValue(*bucket.Name),
			}
			prometheusNumberOfObjectsMetric.Add(ctx, numberOfObjects, bucketNameAtt)
		}

		// Get the BucketSizeBytes metric for the bucket
		bucketSizeBytesMetric, err := cw.GetMetricData(&cloudwatch.GetMetricDataInput{
			MetricDataQueries: []*cloudwatch.MetricDataQuery{
				{
					Id: aws.String("m2"),
					MetricStat: &cloudwatch.MetricStat{
						Metric: &cloudwatch.Metric{
							Namespace:  aws.String("AWS/S3"),
							MetricName: aws.String("BucketSizeBytes"),
							Dimensions: []*cloudwatch.Dimension{
								{
									Name:  aws.String("BucketName"),
									Value: bucket.Name,
								},
								{
									Name:  aws.String("StorageType"),
									Value: aws.String("StandardStorage"),
								},
							},
						},
						Period: aws.Int64(86400),
						Stat:   aws.String("Average"),
						Unit:   aws.String("Bytes"),
					},
				},
			},
			StartTime: aws.Time(time.Now().Add(-48 * time.Hour)),
			EndTime:   aws.Time(time.Now()),
		})
		if err != nil {
			return "", err
		}
		if len(bucketSizeBytesMetric.MetricDataResults) > 0 {
			bucketSizeBytes := int64(*bucketSizeBytesMetric.MetricDataResults[0].Values[0])
			prometheusBucketSizeBytesMetric := metric.Must(meter).NewInt64UpDownCounter("BucketSizeBytes")
			bucketNameAtt := attribute.KeyValue{
				Key:   "bucket_name",
				Value: attribute.StringValue(*bucket.Name),
			}
			prometheusBucketSizeBytesMetric.Add(ctx, bucketSizeBytes, bucketNameAtt)
		}
	}

	return "Success", nil
}

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

func getListener() (string, error) {
	listener := os.Getenv(envLogzioMetricsListener)
	if listener == "" {
		return "", fmt.Errorf("%s must be set", envLogzioMetricsListener)
	}

	return listener, nil
}

func getLogzioToken() (string, error) {
	listener := os.Getenv(envLogzioMetricsToken)
	if listener == "" {
		return "", fmt.Errorf("%s must be set", envLogzioMetricsListener)
	}

	return listener, nil
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("encountered error: ", err)
	}
}
