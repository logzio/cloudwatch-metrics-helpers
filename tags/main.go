package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"os"
	"strings"
)

const (
	envAwsRegion     = "AWS_REGION" // reserved env
	envAwsNamespaces = "AWS_NAMESPACES"

	emptyString   = ""
	listSeparator = ","
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest() error {
	services, sess, err := initialize()
	if err != nil {
		return err
	}

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

func initialize() ([]string, *session.Session, error) {
	namespaces, err := getAwsNamespaces()
	if err != nil {
		return nil, nil, fmt.Errorf("could not retrieve namespaces: %s. Aborting", err.Error())
	}

	services := namespacesToServices(namespaces)

	sess, err := getSession()
	if err != nil {
		return nil, nil, fmt.Errorf("could not create AWS session: %s. Aborting", err.Error())
	}

	return services, sess, nil
}
