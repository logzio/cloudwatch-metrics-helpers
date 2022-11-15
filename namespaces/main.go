package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"log"
	"os"
	"strings"
)

const (
	envAwsRegion     = "AWS_REGION" // reserved env
	envAwsNamespaces = "AWS_NAMESPACES"
	envStreamName    = "METRIC_STREAM_NAME"
	envFirehoseArn   = "FIREHOSE_ARN"
	envRoleArn       = "METRIC_STREAM_ROLE_ARN"

	emptyString   = ""
	listSeparator = ","
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event cfn.Event) (string, error) {
	log.Println("Starting triggers run")

	// If requestID is empty - the lambda call is not from a custom resource
	if event.RequestID != "" && event.RequestType == cfn.RequestCreate {
		// Custom resource CREATE invocation
		lambda.Start(cfn.LambdaWrap(run))
	}

	return "Lambda finished", nil
}

func run(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	awsNs, err := getAwsNamespaces()
	if err != nil {
		return
	}

	sess, err := getSession()
	if err != nil {
		return
	}

	client := cloudwatch.New(sess)
	streamName := os.Getenv(envStreamName)
	firehoseArn := os.Getenv(envFirehoseArn)
	roleArn := os.Getenv(envRoleArn)
	outputFormat := "opentelemetry0.7"
	filters := make([]*cloudwatch.MetricStreamFilter, 0)
	for _, namespace := range awsNs {
		filter := new(cloudwatch.MetricStreamFilter)
		filter.Namespace = &namespace
		filters = append(filters, filter)
	}

	putFilterOutput, err := client.PutMetricStream(&cloudwatch.PutMetricStreamInput{
		FirehoseArn:    &firehoseArn,
		IncludeFilters: filters,
		Name:           &streamName,
		OutputFormat:   &outputFormat,
		RoleArn:        &roleArn,
	})

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(putFilterOutput.String())

	return
}

func getSession() (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(os.Getenv(envAwsRegion)),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error occurred while trying to create a connection to aws: %s. Aborting", err.Error())
	}

	return sess, nil
}

func getAwsNamespaces() ([]string, error) {
	nsStr := os.Getenv(envAwsNamespaces)
	if nsStr == emptyString {
		return nil, fmt.Errorf("env %s must be set", envAwsNamespaces)
	}

	nsStr = strings.ReplaceAll(nsStr, " ", "")

	ns := strings.Split(nsStr, listSeparator)
	log.Printf("detected the following services: %v", ns)
	return ns, nil
}
