## AWS CloudWatch S3 Metrics Collector for AWS Lambda
This Go program is designed to be used as an AWS Lambda function that collects CloudWatch metrics for all S3 buckets in an AWS account, and exports the metrics to Logz.io. The program collects the `NumberOfObjects` and `BucketSizeBytes` metrics for each bucket, and exports them with the following dimensions:

- `BucketName`: The name of the S3 bucket

### Prerequisites
Before you can use this program, you will need to do the following:

- Set up a Logz.io account and obtain a Logz.io listener URL and token.
- Set up the AWS CLI and configure it with your AWS credentials.

### Configuration
- To configure the program, you will need to set the following environment variables in the AWS Lambda function's configuration:

- `LOGZIO_METRICS_LISTENER`: The Logz.io listener URL
- `LOGZIO_METRICS_TOKEN`: The Logz.io token

### Deploying to AWS Lambda using the AWS CLI
To deploy the source code to AWS Lambda using the AWS CLI, you will need to do the following:

- Install the AWS CLI and configure it with your AWS credentials.
- Create a new AWS Lambda function by running the following command:
```shell
aws lambda create-function \
--function-name <function-name> \
--runtime go1.x \
--handler main \
--environment Variables={LOGZIO_METRICS_LISTENER=<listener-url>,LOGZIO_METRICS_TOKEN=<token>} \
--zip-file fileb://function.zip
```
- Replace `<function-name>`, `<listener-url>`, and `<token>` with your desired function name, Logz.io listener URL, and Logz.io token, respectively. The function.zip file should contain the compiled Go binary and all its dependencies.

To create the function.zip file, you can run the `make function` command.

- Test the function by running the following command:
```shell
aws lambda invoke \
--function-name <function-name> \
--payload '{}' \
output.json
```
This will invoke the function and save the output to the output.json file.