package s3_daily_metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {
	lambda.Start(getS3Metrics)
}

func getS3Metrics(ctx context.Context) error {
	// Create an AWS session
	sess, err := session.NewSession()
	if err != nil {
		return fmt.Errorf("error creating AWS session: %s", err)
	}

	// Create a CloudWatch client
	svc := cloudwatch.New(sess)

	// Set the start and end times for the metric data
	now := time.Now()
	endTime := now.Add(time.Hour * 24)
	startTime := endTime.Add(time.Hour * -48)

	// Set the CloudWatch namespace to "AWS/S3"
	namespace := "AWS/S3"

	// Set the metrics to retrieve
	metrics := []string{
		"NumberOfObjects",
		"BucketSizeBytes",
	}

	// Set the dimensions for the metric data
	dimensions := []*cloudwatch.Dimension{
		{
			Name:  aws.String("BucketName"),
			Value: aws.String("*"),
		},
		{
			Name:  aws.String("StorageType"),
			Value: aws.String("*"),
		},
	}

	// Create a CloudWatch GetMetricDataInput struct
	input := &cloudwatch.GetMetricDataInput{
		EndTime:   &endTime,
		StartTime: &startTime,
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			{
				Id: aws.String("m1"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{
						Namespace:  &namespace,
						MetricName: &metrics[0],
						Dimensions: dimensions,
					},
					Period: aws.Int64(3600),
					Stat:   aws.String("Average"),
				},
			},
			{
				Id: aws.String("m2"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{
						Namespace:  &namespace,
						MetricName: &metrics[1],
						Dimensions: dimensions,
					},
					Period: aws.Int64(3600),
					Stat:   aws.String("Average"),
				},
			},
		},
		MaxDatapoints: aws.Int64(100),
	}

	// Call the CloudWatch GetMetricData API
	result, err := svc.GetMetricData(input)
	if err != nil {
		return fmt.Errorf("error calling CloudWatch GetMetricData API: %s", err)
	}

	// Print the metric data
	for _, metricDataResult := range result.MetricDataResults {
		fmt.Printf("Metric: %s\n", *metricDataResult.Label)
		for _, datapoint := range metricDataResult.Values {
			fmt.Printf("Timestamp: %s, Value: %f\n", *datapoint.Timestamp, *datapoint.Value)
		}
	}

	return nil
}
