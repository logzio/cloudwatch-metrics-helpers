package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"log"
	"os"
	"strings"
)

const (
	envAwsRegion       = "AWS_REGION" // reserved env
	envAwsNamespaces   = "AWS_NAMESPACES"
	envCustomNamespace = "CUSTOM_NAMESPACE"
	envStreamName      = "METRIC_STREAM_NAME"
	envFirehoseArn     = "FIREHOSE_ARN"
	envRoleArn         = "METRIC_STREAM_ROLE_ARN"
	envDebugMode       = "DEBUG_MODE"

	emptyString   = ""
	listSeparator = ","

	paramLogzioMetricsListener = "logzioListener"
	paramLogzioMetricsToken    = "logzioToken"
	envLogzioMetricsListener   = "LOGZIO_METRICS_LISTENER"
	envLogzioMetricsToken      = "LOGZIO_METRICS_TOKEN"
	envStackName               = "STACK_NAME"
	version                    = "latest"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func debugLog(format string, v ...interface{}) {
	if strings.ToLower(os.Getenv(envDebugMode)) == "true" {
		logger.Printf(format, v...)
	}
}

func generatePhysicalResourceId(event cfn.Event) string {
	// Concatenate StackId and LogicalResourceId to form a unique PhysicalResourceId
	physicalResourceId := fmt.Sprintf("%s-%s", event.StackID, event.LogicalResourceID)
	return physicalResourceId
}

func main() {
	lambda.Start(cfn.LambdaWrap(HandleRequest))
}

func HandleRequest(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	debugLog("Debug mode enabled: Starting Lambda function execution.")
	debugLog("Lambda HandleRequest invoked with event: %+v", event)

	logger.Println("Starting triggers run")

	// If requestID is empty - the lambda call is not from a custom resource
	if event.RequestID != "" && event.RequestType == cfn.RequestCreate {
		// Custom resource CREATE invocation
		return customResourceRun(ctx, event)
	} else {
		return customResourceDoNothing(ctx, event)
	}
}

// Wrapper for Read, Update, Delete requests
func customResourceDoNothing(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	debugLog("Executing 'Do Nothing' operation for request: %+v", event)
	return generatePhysicalResourceId(event), nil, nil
}

// Wrapper for first invocation from cloud formation custom resource
func customResourceRun(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	debugLog("Executing custom resource 'Create' operation for event: %+v", event)

	if event.RequestType == cfn.RequestCreate {
		err = run()
		if err != nil {
			logger.Printf("Error encountered during 'Create' operation: %s", err.Error())
			return
		}
	} else {
		logger.Println("Got ", event.RequestType, " request")
	}

	return generatePhysicalResourceId(event), nil, nil
}

func run() error {
	debugLog("Beginning main logic execution for the custom resource deployment.")
	DeployS3Function := false
	awsNs, err := getAwsNamespaces()
	if err != nil {
		logger.Printf("Failed to get AWS namespaces: %s", err.Error())
		return err
	}

	sess, sessErr := getSession()
	if sessErr != nil {
		logger.Printf("Failed to create AWS session: %s", sessErr.Error())
		return sessErr
	}

	debugLog("AWS session created successfully. AWS Namespaces to include: %v", awsNs)

	client := cloudwatch.New(sess)
	streamName := os.Getenv(envStreamName)
	firehoseArn := os.Getenv(envFirehoseArn)
	roleArn := os.Getenv(envRoleArn)

	debugLog("Preparing to create CloudWatch metric stream with name: %s", streamName)

	outputFormat := "opentelemetry0.7"
	filters := make([]*cloudwatch.MetricStreamFilter, 0)
	for _, namespace := range awsNs {
		filter := &cloudwatch.MetricStreamFilter{Namespace: aws.String(namespace)}
		filters = append(filters, filter)
		if namespace == nsS3 {
			DeployS3Function = true
		}
	}

	logger.Printf("Filters prepared: %+v", filters)

	putFilterOutput, err := client.PutMetricStream(&cloudwatch.PutMetricStreamInput{
		FirehoseArn:    &firehoseArn,
		IncludeFilters: filters,
		Name:           &streamName,
		OutputFormat:   &outputFormat,
		RoleArn:        &roleArn,
	})

	if err != nil {
		logger.Printf("Failed to create/update CloudWatch metric stream: %s", err.Error())
		return err
	}

	debugLog("CloudWatch metric stream created/updated successfully")

	logger.Println(putFilterOutput.String())

	// deploy s3 function if needed
	if DeployS3Function {
		log.Printf("Deploying S3 function")
		cloudformationClient := cloudformation.New(sess)
		listener, getListenerErr := getListener()
		if getListenerErr != nil {
			return fmt.Errorf("error while getting logzio listener address: %s", getListenerErr.Error())
		}
		token, tokenErr := getLogzioToken()
		if tokenErr != nil {
			return fmt.Errorf("error while getting logzio token: %s", tokenErr.Error())
		}
		currentStack, stackErr := getStackName()
		if stackErr != nil {
			return fmt.Errorf("error while getting stack name: %s", stackErr.Error())
		}
		params := []*cloudformation.Parameter{
			{
				ParameterKey:   aws.String(paramLogzioMetricsListener),
				ParameterValue: aws.String(listener),
			},
			{
				ParameterKey:   aws.String(paramLogzioMetricsToken),
				ParameterValue: aws.String(token),
			},
		}
		stackName := fmt.Sprintf("%v-s3", currentStack)
		templateUrl := fmt.Sprintf("https://logzio-aws-integrations-%v.s3.amazonaws.com/metric-stream-helpers/aws/%v/sam-s3-daily-metrics.yaml", os.Getenv(envAwsRegion), version)
		// Create a new CloudFormation stack
		_, cfErr := cloudformationClient.CreateStack(&cloudformation.CreateStackInput{
			StackName:   aws.String(stackName),
			TemplateURL: aws.String(templateUrl),
			Parameters:  params,
			Capabilities: []*string{
				aws.String(cloudformation.CapabilityCapabilityAutoExpand),
				aws.String(cloudformation.CapabilityCapabilityIam),
				aws.String(cloudformation.CapabilityCapabilityNamedIam),
			},
		})
		if cfErr != nil {
			logger.Printf("Error while creating stack: %s", cfErr.Error())
			return err
		}
		logger.Printf("Stack %v created successfully", stackName)
	}
	return nil
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
	awsNsStr := os.Getenv(envAwsNamespaces)
	customNs := os.Getenv(envCustomNamespace)

	if awsNsStr == emptyString && customNs == emptyString {
		return nil, fmt.Errorf("either %s or %s must be set", envAwsNamespaces, envCustomNamespace)
	}

	// Create a slice to hold the namespaces
	var namespaces []string

	// Add awsNsStr and customNs to the slice if they are not empty
	if awsNsStr != emptyString {
		namespaces = append(namespaces, awsNsStr)
	}
	if customNs != emptyString {
		namespaces = append(namespaces, customNs)
	}

	// Join namespace strings with the list separator
	fullNsStr := strings.Join(namespaces, listSeparator)
	// Remove all spaces from the final namespace string
	fullNsStr = strings.ReplaceAll(fullNsStr, " ", "")

	ns := strings.Split(fullNsStr, listSeparator)
	for _, namespace := range ns {
		if namespace == nsAll {
			logger.Println("detected ALL namespaces")
			return getAllNamespaces(), nil
		}
	}
	logger.Printf("detected the following services: %v", ns)
	return ns, nil
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
		return "", fmt.Errorf("%s must be set", envLogzioMetricsToken)
	}

	return listener, nil
}

// getStackName gets the name of the cfn stack from environment variables
func getStackName() (string, error) {
	stackName := os.Getenv(envStackName)
	if stackName == "" {
		return "", fmt.Errorf("%s must be set", envStackName)
	}

	return stackName, nil
}
