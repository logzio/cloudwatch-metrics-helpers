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
	"strings"
	"time"
)

const (
	envAwsRegion             = "AWS_REGION" // reserved env
	envAwsNamespaces         = "AWS_NAMESPACES"
	envLogzioMetricsListener = "LOGZIO_METRICS_LISTENER"
	envLogzioMetricsToken    = "LOGZIO_METRICS_TOKEN"

	emptyString   = ""
	listSeparator = ","
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) error {
	services, client, exporter, err := initialize()
	if err != nil {
		return err
	}

	defer func() {
		handleErr(exporter.Stop(ctx))
	}()
	meter := exporter.Meter("aws_resource_info")
	intUpDownCounter := metric.Must(meter).NewInt64UpDownCounter("aws_resource_info")

	var nextToken *string
	nextToken = nil
	resourcesPerPage := int64(100)
	callsCounter := 0
	for {
		input := resourcegroupstaggingapi.GetResourcesInput{
			PaginationToken:     nextToken,
			ResourceTypeFilters: services,
			ResourcesPerPage:    &resourcesPerPage,
		}

		getResourcesOutput, err := client.GetResources(&input)
		callsCounter += 1

		if err != nil {
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

	logger.Println("finished lambda run")
	return nil
}

func getAwsNamespaces() ([]string, error) {
	nsStr := os.Getenv(envAwsNamespaces)
	if nsStr == emptyString {
		return nil, fmt.Errorf("env %s must be set", envAwsNamespaces)
	}

	nsStr = strings.ReplaceAll(nsStr, " ", "")

	ns := strings.Split(nsStr, listSeparator)
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

		logger.Println("handling resource ", *resource.ResourceARN)
		attributes := make([]attribute.KeyValue, 0)
		for _, tag := range resource.Tags {
			attributes = append(attributes, attribute.KeyValue{
				Key:   attribute.Key(*tag.Key),
				Value: attribute.StringValue(*tag.Value),
			})
		}

		attributes = append(attributes, attribute.KeyValue{
			Key:   "arn",
			Value: attribute.StringValue(*resource.ResourceARN),
		})

		counter.Add(ctx, int64(1), attributes...)
	}

}

func handleErr(err error) {
	if err != nil {
		logger.Println("encountered error: ", err)
	}
}
