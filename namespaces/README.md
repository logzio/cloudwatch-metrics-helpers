### Logz.io Metric Stream Helper - Namespaces

This lambda function is intended to run once at Cloudformation stack creation, and to add namespaces to the Logz.io metrics stream.

#### Environment variables:

**All environment variables mentioned here are required!**

| Name                     | Description                                                                                                                                                                                                                                                                                |
|--------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `AWS_NAMESPACES`         | Comma delimited list of the AWS namespaces to collect metrics from. Possible values can be found [here](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html). If you want to automatically add all namespaces, use value `all-namespaces`. |
| `CUSTOM_NAMESPACE`       | A Parameter to define a unique namespace for CloudWatch metrics, separate from standard AWS namespaces.                                                                                                                                                                                    |
| `METRIC_STREAM_NAME`     | Name of the metric stream you want to add namespace to                                                                                                                                                                                                                                     |
| `FIREHOSE_ARN`           | ARN of the Kinesis Firehose Delivery Stream that's attached to the metric stream.                                                                                                                                                                                                          |
| `METRIC_STREAM_ROLE_ARN` | ARN of your metric stream IAM role.                                                                                                                                                                                                                                                        | 


#### Changelog:

- **1.2.1**:
  - Added debug mode support for enhanced troubleshooting.
  - Added generation and return of physicalResourceId for improved CloudFormation resource management.
- **1.2.0**: Added the functionality of `CUSTOM_NAMESPACE` to specify a unique namespace.
- **1.1.0**: Support adding all namespaces by setting `AWS_NAMESPACES` with `all-namespaces`.
- **1.0.0**: Initial release.