package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cwtypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

const (
	envAwsRegion             = "AWS_REGION" // reserved env
	envAwsNamespaces         = "AWS_NAMESPACES"
	envCustomNamespace       = "CUSTOM_NAMESPACE"
	envIncludeMetricsFilters = "INCLUDE_METRICS_FILTERS"
	envStreamName            = "METRIC_STREAM_NAME"
	envFirehoseArn           = "FIREHOSE_ARN"
	envRoleArn               = "METRIC_STREAM_ROLE_ARN"
	envDebugMode             = "DEBUG_MODE"
	envP8slogzioName         = "P8S_LOGZIO_NAME"

	emptyString   = ""
	listSeparator = ","

	paramLogzioMetricsListener = "logzioListener"
	paramLogzioMetricsToken    = "logzioToken"
	paramP8slogzioName         = "p8sLogzioName"
	envLogzioMetricsListener   = "LOGZIO_METRICS_LISTENER"
	envLogzioMetricsToken      = "LOGZIO_METRICS_TOKEN"
	envStackName               = "STACK_NAME"
	version                    = "latest"
	defualtP8sLogzioName       = "cloudwatch-helpers"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func debugLog(format string, v ...interface{}) {
	if strings.ToLower(os.Getenv(envDebugMode)) == "true" {
		logger.Printf(format, v...)
	}
}

// parseIncludeMetricsFilters parses comma-separated namespace:metric pairs with robust error handling
// Example: "AWS/EC2:CPUUtilization,AWS/EC2:NetworkIn,AWS/S3:BucketSizeBytes"
// Returns: (map[namespace][]metrics, warnings)
func parseIncludeMetricsFilters(input string) (map[string][]string, []string) {
	result := make(map[string][]string)
	var warnings []string

	if strings.TrimSpace(input) == "" {
		return result, warnings
	}

	pairs := strings.Split(input, listSeparator)
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			warnings = append(warnings, fmt.Sprintf("invalid format '%s', expected 'namespace:metric' - skipping", pair))
			continue
		}

		namespace := strings.TrimSpace(parts[0])
		metric := strings.TrimSpace(parts[1])

		if namespace == "" || metric == "" {
			warnings = append(warnings, fmt.Sprintf("empty namespace or metric in '%s' - skipping", pair))
			continue
		}

		if !isValidASCII(namespace) || !isValidASCII(metric) {
			warnings = append(warnings, fmt.Sprintf("invalid characters in '%s' - skipping", pair))
			continue
		}

		found := false
		for _, existing := range result[namespace] {
			if existing == metric {
				found = true
				break
			}
		}
		if !found {
			result[namespace] = append(result[namespace], metric)
		}
	}

	return result, warnings
}

// isValidASCII checks if string contains only ASCII printable characters (32-126)
func isValidASCII(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

// buildIncludeFilters creates IncludeFilters from namespaces and optional metric filters
// Applies 1000-name limit gracefully
func buildIncludeFilters(namespaces []string, includeMap map[string][]string) ([]cwtypes.MetricStreamFilter, []string) {
	filters := make([]cwtypes.MetricStreamFilter, 0, len(namespaces))
	var warnings []string

	totalNames := len(namespaces)
	for _, metrics := range includeMap {
		totalNames += len(metrics)
	}

	remainingBudget := 1000 - len(namespaces)
	if remainingBudget < 0 {
		remainingBudget = 0
	}

	if totalNames > 1000 {
		warnings = append(warnings, fmt.Sprintf("total filter names (%d) exceeds CloudWatch limit (1000), some metric names will be dropped", totalNames))
	}

	for _, ns := range namespaces {
		if metrics, ok := includeMap[ns]; ok && len(metrics) > 0 {
			metricsToInclude := metrics
			if len(metrics) > remainingBudget {
				metricsToInclude = metrics[:remainingBudget]
				if remainingBudget > 0 {
					warnings = append(warnings, fmt.Sprintf("dropped %d metric names from namespace %s due to 1000-name limit", len(metrics)-remainingBudget, ns))
				}
				remainingBudget = 0
			} else {
				remainingBudget -= len(metrics)
			}

			if len(metricsToInclude) > 0 {
				filters = append(filters, cwtypes.MetricStreamFilter{
					Namespace:   aws.String(ns),
					MetricNames: metricsToInclude,
				})
			} else {
				filters = append(filters, cwtypes.MetricStreamFilter{
					Namespace: aws.String(ns),
				})
			}
		} else {
			filters = append(filters, cwtypes.MetricStreamFilter{
				Namespace: aws.String(ns),
			})
		}
	}

	return filters, warnings
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
	debugLog("Since the request type is not 'Create', no further action will be taken for this event: %+v", event)
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
	debugLog("Starting main logic execution for the custom resource deployment.")
	DeployS3Function := false
	awsNs, err := getAwsNamespaces()
	if err != nil {
		logger.Printf("Failed to get AWS namespaces: %s", err.Error())
		return err
	}

	cfg, sessErr := getSession()
	if sessErr != nil {
		logger.Printf("Failed to create AWS session: %s", sessErr.Error())
		return sessErr
	}

	debugLog("AWS session created successfully. AWS Namespaces to include: %v", awsNs)

	client := cloudwatch.NewFromConfig(*cfg)
	streamName := os.Getenv(envStreamName)
	firehoseArn := os.Getenv(envFirehoseArn)
	roleArn := os.Getenv(envRoleArn)

	debugLog("Preparing to create CloudWatch metric stream with name: %s", streamName)

	includeMetricsInput := os.Getenv(envIncludeMetricsFilters)
	includeMap, parseWarnings := parseIncludeMetricsFilters(includeMetricsInput)

	for _, warning := range parseWarnings {
		logger.Printf("Warning: %s", warning)
	}

	filteredIncludeMap := make(map[string][]string)
	namespaceSet := make(map[string]bool)
	for _, ns := range awsNs {
		namespaceSet[ns] = true
	}

	for ns, metrics := range includeMap {
		if namespaceSet[ns] {
			filteredIncludeMap[ns] = metrics
		} else {
			logger.Printf("Warning: ignoring metrics for namespace '%s' not in AWS_NAMESPACES", ns)
		}
	}

	outputFormat := cwtypes.MetricStreamOutputFormat("opentelemetry1.0")

	for _, namespace := range awsNs {
		if namespace == nsS3 {
			DeployS3Function = true
			break
		}
	}

	includeFilters, filterWarnings := buildIncludeFilters(awsNs, filteredIncludeMap)

	for _, warning := range filterWarnings {
		logger.Printf("Warning: %s", warning)
	}

	debugLog("Using include filters for namespaces: %v with metric filters: %v", awsNs, filteredIncludeMap)

	input := &cloudwatch.PutMetricStreamInput{
		FirehoseArn:    aws.String(firehoseArn),
		Name:           aws.String(streamName),
		OutputFormat:   outputFormat,
		RoleArn:        aws.String(roleArn),
		IncludeFilters: includeFilters,
	}

	logger.Printf("Filters prepared: Include=%d", len(input.IncludeFilters))

	putFilterOutput, err := client.PutMetricStream(context.TODO(), input)
	if err != nil {
		logger.Printf("Failed to create/update CloudWatch metric stream: %s", err.Error())
		return err
	}

	debugLog("CloudWatch metric stream created/updated successfully")
	logger.Printf("Metric stream ARN: %s", aws.ToString(putFilterOutput.Arn))

	// deploy s3 function if needed
	if DeployS3Function {
		log.Printf("Deploying S3 function")
		cloudformationClient := cloudformation.NewFromConfig(*cfg)
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
		params := []cftypes.Parameter{
			{
				ParameterKey:   aws.String(paramLogzioMetricsListener),
				ParameterValue: aws.String(listener),
			},
			{
				ParameterKey:   aws.String(paramLogzioMetricsToken),
				ParameterValue: aws.String(token),
			},
			{
				ParameterKey:   aws.String(paramP8slogzioName),
				ParameterValue: aws.String(getP8sLogzioName()),
			},
		}
		stackName := fmt.Sprintf("%v-s3", currentStack)
		templateUrl := fmt.Sprintf("https://logzio-aws-integrations-%v.s3.amazonaws.com/metric-stream-helpers/aws/%v/sam-s3-daily-metrics.yaml", os.Getenv(envAwsRegion), version)
		// Create a new CloudFormation stack
		_, cfErr := cloudformationClient.CreateStack(context.TODO(), &cloudformation.CreateStackInput{
			StackName:   aws.String(stackName),
			TemplateURL: aws.String(templateUrl),
			Parameters:  params,
			Capabilities: []cftypes.Capability{
				cftypes.CapabilityCapabilityAutoExpand,
				cftypes.CapabilityCapabilityIam,
				cftypes.CapabilityCapabilityNamedIam,
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

func getSession() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv(envAwsRegion)))
	if err != nil {
		return nil, fmt.Errorf("error occurred while trying to create a connection to aws: %s. Aborting", err.Error())
	}

	return &cfg, nil
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
		debugLog("Appending AWS namespace string to namespaces: %s", awsNsStr)
	}
	if customNs != emptyString {
		namespaces = append(namespaces, customNs)
		debugLog("Appending custom namespace string to namespaces: %s", customNs)
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

func getP8sLogzioName() string {
	envTag := os.Getenv(envP8slogzioName)
	if envTag == "" {
		return defualtP8sLogzioName
	}

	return envTag
}

// getStackName gets the name of the cfn stack from environment variables
func getStackName() (string, error) {
	stackName := os.Getenv(envStackName)
	if stackName == "" {
		return "", fmt.Errorf("%s must be set", envStackName)
	}

	return stackName, nil
}
