package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	metricsExporter "github.com/logzio/go-metrics-sdk"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	envAwsRegion                = "AWS_REGION" // reserved env
	envAwsNamespaces            = "AWS_NAMESPACES"
	envLogzioMetricsListener    = "LOGZIO_METRICS_LISTENER"
	envLogzioMetricsToken       = "LOGZIO_METRICS_TOKEN"
	envAwsLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION" // reserved env
	envDebugMode                = "DEBUG_MODE"                  // Added debug mode environment variable

	emptyString             = ""
	listSeparator           = ","
	fieldLogzioAgentVersion = "logzio_agent_version"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func debugLog(format string, v ...interface{}) {
	if strings.ToLower(os.Getenv(envDebugMode)) == "true" {
		logger.Printf(format, v...)
	}
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) error {
	debugLog("Debug mode enabled: Starting Lambda function execution.")

	services, client, exporter, err := initialize()
	if err != nil {
		logger.Printf("Error in initialize: %s", err)
		return err
	}

	defer func() {
		debugLog("Stopping metrics exporter and finalizing function execution.")
		handleErr(exporter.Stop(ctx))
	}()
	meter := exporter.Meter("aws_resource_info")
	intUpDownCounter := metric.Must(meter).NewInt64UpDownCounter("aws_resource_info")

	var nextToken *string
	nextToken = nil
	resourcesPerPage := int64(100)
	callsCounter := 0
	for {
		debugLog("Fetching resources from AWS Resource Groups Tagging API, call number: %d", callsCounter+1)

		input := resourcegroupstaggingapi.GetResourcesInput{
			PaginationToken:     nextToken,
			ResourceTypeFilters: services,
			ResourcesPerPage:    &resourcesPerPage,
		}

		getResourcesOutput, err := client.GetResources(&input)
		callsCounter += 1

		if err != nil {
			logger.Printf("Error fetching resources: %s", err)
			return fmt.Errorf("error occurred while trying to get resources: %s", err.Error())
		}

		if getResourcesOutput != nil {
			nextToken = getResourcesOutput.PaginationToken
			sendTags(ctx, getResourcesOutput.ResourceTagMappingList, &intUpDownCounter)
		}

		if *nextToken == emptyString {
			break
		}
	}

	logger.Printf("Finished lambda run after %d API calls", callsCounter)
	return nil
}

func getAwsNamespaces() ([]string, error) {
	nsStr := os.Getenv(envAwsNamespaces)
	if nsStr == emptyString {
		return nil, fmt.Errorf("env %s must be set", envAwsNamespaces)
	}

	nsStr = strings.ReplaceAll(nsStr, " ", "")

	ns := strings.Split(nsStr, listSeparator)
	for _, namespace := range ns {
		if namespace == nsAll {
			logger.Println("detected ALL namespaces")
			return getAllNamespaces(), nil
		}
	}

	logger.Printf("detected the following services: %v", ns)
	return ns, nil
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

func initialize() ([]*string, *resourcegroupstaggingapi.ResourceGroupsTaggingAPI, *basic.Controller, error) {
	debugLog("Initializing AWS session, metrics exporter, and resource group tagging API client.")
	namespaces, err := getAwsNamespaces()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not retrieve namespaces: %s. Aborting", err.Error())
	}

	services := namespacesToServices(namespaces)
	if len(services) == 0 {
		return nil, nil, nil, fmt.Errorf("could not translate any namespace to service. Aborting")
	}

	sess, err := getSession()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("could not create AWS session: %s. Aborting", err.Error())
	}

	tagsClient := resourcegroupstaggingapi.New(sess)
	if tagsClient == nil {
		return nil, nil, nil, fmt.Errorf("could not initialize resource groups tags client. Aborting")
	}

	exporter, err := configureMetricsExporter()
	if err != nil {
		return nil, nil, nil, err
	}

	return services, tagsClient, exporter, nil
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
	if listener == emptyString {
		return emptyString, fmt.Errorf("%s must be set", envLogzioMetricsListener)
	}

	return listener, nil
}

func getLogzioToken() (string, error) {
	listener := os.Getenv(envLogzioMetricsToken)
	if listener == emptyString {
		return emptyString, fmt.Errorf("%s must be set", envLogzioMetricsListener)
	}

	return listener, nil
}

func sendTags(ctx context.Context, resources []*resourcegroupstaggingapi.ResourceTagMapping, counter *metric.Int64UpDownCounter) {
	for _, resource := range resources {
		// Since GetResources returns all resources that ever had a tag, it can return a resource that is now without a tag.
		// Skip a tag-less resource
		if len(resource.Tags) == 0 {
			continue
		}

		attributes := make([]attribute.KeyValue, 0)
		for _, tag := range resource.Tags {
			attributes = append(attributes, attribute.KeyValue{
				Key:   attribute.Key(*tag.Key),
				Value: attribute.StringValue(*tag.Value),
			})
		}

		idLabels := addIdentifier(*resource.ResourceARN)
		if idLabels == nil {
			attributes = append(attributes, attribute.KeyValue{
				Key:   "Arn",
				Value: attribute.StringValue(*resource.ResourceARN),
			})
		} else {
			attributes = append(attributes, idLabels...)
		}

		attributes = append(attributes, attribute.KeyValue{
			Key:   fieldLogzioAgentVersion,
			Value: attribute.StringValue(os.Getenv(envAwsLambdaFunctionVersion)),
		})

		counter.Add(ctx, int64(1), attributes...)
	}

}

func addIdentifier(resourceArn string) []attribute.KeyValue {
	arnToId := getArnRegexToIdsMap()
	for arn, _ := range arnToId {
		matched, err := regexp.MatchString(arn, resourceArn)
		if err != nil {
			logger.Println("error while trying to match resource arn ", resourceArn, " with regex ", arn)
			continue
		}

		if matched {
			return getLabels(resourceArn, arn, arnToId)
		}
	}

	return nil
}

func handleErr(err error) {
	if err != nil {
		logger.Println("encountered error: ", err)
	}
}

func getLabels(resourceArn, matchedRegexArn string, arnToIds map[string][]string) []attribute.KeyValue {
	// Special cases
	switch matchedRegexArn {
	case arnRegEc2:
		// In this case we need to check if it's indeed EC2, or EBS
		matched, err := regexp.MatchString(arnRegEbs, resourceArn)
		if err != nil {
		}
		if err == nil && matched {
			return getKeyValuePairs(arnToIds[arnRegEbs], resourceArn, true)
		} else {
			return getKeyValuePairs(arnToIds[arnRegEc2], resourceArn, true)
		}
	case arnRegAcm:
	case arnRegAcmPca:
	case arnRegStepFunction:
		// Resources that their identifier is an entire ARN
		return getKeyValuePairs(arnToIds[matchedRegexArn], resourceArn, false)
	}

	return getKeyValuePairs(arnToIds[matchedRegexArn], resourceArn, true)
}

func getKeyValuePairs(keys []string, arn string, idOnly bool) []attribute.KeyValue {
	value := arn
	if idOnly {
		arnArr := strings.Split(arn, ":")
		value = arnArr[len(arnArr)-1]
		if strings.Contains(value, "/") {
			resourceArray := strings.Split(value, "/")
			value = resourceArray[len(resourceArray)-1]
		}
	}

	logger.Println("for resource ", arn, " the id is: ", value)

	attributes := make([]attribute.KeyValue, 0)
	for _, key := range keys {
		attributes = append(attributes, attribute.KeyValue{
			Key:   attribute.Key(strings.ToLower(key)),
			Value: attribute.StringValue(value),
		})
	}

	return attributes
}
