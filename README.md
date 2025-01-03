# Cloudwatch Metric Stream

Deploy this integration to send your CloudWatch metrics to Logz.io.


## Overview

This integration creates a Kinesis Data Firehose delivery stream that links to your CloudWatch metrics stream and then sends the metrics to your Logz.io account. It also creates a Lambda function that adds AWS namespaces to the metric stream, and a Lambda function that collects and ships the resources' tags.

## Instructions

To deploy this project, click the button that matches the region you wish to deploy your Stack to:

| Region           | Deployment                                                                                                                                                                                                                                                                                                                                                   |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `us-east-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-us-east-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `us-east-2`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-east-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-us-east-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `us-west-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-us-west-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `us-west-2`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=us-west-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-us-west-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `eu-central-1`   | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-central-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)               |
| `eu-central-2`   | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-central-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)               |
| `eu-north-1`     | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-north-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-north-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                   |
| `eu-west-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-west-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `eu-west-2`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-west-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `eu-west-3`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-west-3#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-west-3.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `sa-east-1`      | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=sa-east-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-sa-east-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                     |
| `ap-northeast-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-northeast-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ap-northeast-2` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-northeast-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ap-northeast-3` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-northeast-3#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-northeast-3.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ap-south-1`     | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-south-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-south-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                 |
| `ap-southeast-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-southeast-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ap-southeast-2` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-southeast-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ca-central-1`   | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ca-central-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-ca-central-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)               |
| `eu-central-2`   | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-central-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-central-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)               |
| `eu-south-1`     | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-south-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-south-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                   |
| `eu-south-2`     | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=eu-south-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-eu-south-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                   |
| `ap-south-2`     | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-south-2#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-south-2.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)                 |
| `ap-southeast-3` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-3#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-southeast-3.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ap-southeast-4` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-southeast-4#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-southeast-4.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ap-east-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ap-east-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-ap-east-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `ca-west-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=ca-west-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-ca-west-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `af-south-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=af-south-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-af-south-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `me-central-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=me-central-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-me-central-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |
| `il-central-1` | [![Deploy to AWS](https://dytvr9ot2sszz.cloudfront.net/logz-docs/lights/LightS-button.png)](https://console.aws.amazon.com/cloudformation/home?region=il-central-1#/stacks/create/review?templateURL=https://logzio-aws-integrations-il-central-1.s3.amazonaws.com/metric-stream-helpers/aws/1.4.0/sam-template.yaml&stackName=logzio-metric-stream)           |


### 1. Specify stack details

Specify the stack details as per the table below, check the checkboxes and select **Create stack**.

| Parameter                                  | Description                                                                                                                                                                                                                                                                       | Required/Default                                                 |
|--------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------|
| `logzioListener`                           | The Logz.io listener URL for your region. (For more details, see the [regions page](https://docs.logz.io/user-guide/accounts/account-region.html). For example - `https://listener.logz.io:8053`                                                                                  | **Required**                                                     |
| `logzioToken`                              | Your Logz.io metrics shipping token.                                                                                                                                                                                                                                              | **Required**                                                     |
| `awsNamespaces`                            | Comma-separated list of the AWS namespaces you want to monitor. See [this list]((https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) of namespaces. If you want to automatically add all namespaces, use value `all-namespaces`. | At least one of `awsNamespaces` or `customNamespace` is required |
| `customNamespace`                          | A custom namespace for CloudWatch metrics. This is used to specify a namespace unique to your setup, separate from the standard AWS namespaces.                                                                                                                                   | At least one of `awsNamespaces` or `customNamespace` is required |
| `logzioDestination`                        | Your Logz.io destination URL.                                                                                                                                                                                                                                                     | **Required**                                                     |
| `httpEndpointDestinationIntervalInSeconds` | The length of time, in seconds, that Kinesis Data Firehose buffers incoming data before delivering it to the destination.                                                                                                                                                         | `60`                                                             |
| `httpEndpointDestinationSizeInMBs`         | The size of the buffer, in MBs, that Kinesis Data Firehose uses for incoming data before delivering it to the destination.                                                                                                                                                        | `5`                                                              |
| `debugMode`                                | Enable debug mode for detailed logging (true/false).                                                                                                                                                                                                                              | `false`                                                          |

### 2. Send your metrics

Give your Cloudformation a few minutes to be created, and that's it!


## Changelog:
- **1.4.0**
  - Added Support for otlp 1.0 protocol
  - Support adding custom 'p8s_logzio_name` label
- **1.3.4**
  - Added debug mode to namespaces and tags to provide detailed operational logs, facilitating easier debugging.
  - Added physicalResourceId in response, enhancing CloudFormation stack resource tracking.
- **1.3.3**
  - Expanded the REGIONS list with new AWS regions.
  - Enhanced the script for better AWS multi-region support by implementing region-specific S3 client usage.
- **1.3.2**
  - Upgrade the runtime from Go 1.x to the provided Amazon Linux 2023.
- **1.3.1**
  - Support collecting `DBInstanceIdentifier` as label for the `aws_resource_info` metric for RDS resources.
- **1.3.0**
  - Added the functionality of `customNamespace` to specify a unique namespace. 
- **1.2.4**
  - Add more logzio destinations to Cloudwatch template.
- **1.2.3**
  - Remove 3 logzio destinations from Cloudformation template.
- **1.2.2**
  - Add `s3 metrics collector` function
  - Updated aws cloudformation templates
  - Added `s3 metrics collector` deployment logic to `namespaces function`
- **1.2.1**:
  - Template - rename field to avoid warning upon deployment.
- **1.2.0**:
  - Add label `logzio_agent_version` to `aws_resouece_info` metrics.
  - Create log group and log stream for Firehose Delivery Stream.
- **1.1.0**:
  - Support adding all namespaces by setting `awsNamespaces` with `all-namespaces`.
  - Tags function: fix EBS service details.
- **1.0.0**: Initial release.
