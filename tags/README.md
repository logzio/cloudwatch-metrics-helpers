### Logz.io Metric Stream Helper - Tags

This Lambda collects tags from resources under selected AWS namespaces, and sends them as metrics to Logz.io.

#### Environment variables:

**All environment variables mentioned here are required!**

| Name                      | Description                                                                                                                                                                                                                                                                                |
|---------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `AWS_NAMESPACES`          | Comma-separated list of the AWS namespaces to collect metrics from. Possible values can be found [here](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html). If you want to automatically add all namespaces, use value `all-namespaces`. |
| `LOGZIO_METRICS_LISTENER` | Full URL for the Logz.io metrics listener, for example `https://listener.logz.io:8053`                                                                                                                                                                                                     |
| `LOGZIO_METRICS_TOKEN`    | Logz.io metrics shipping token                                                                                                                                                                                                                                                             |


#### Changelog:

- **1.1.2**:
  - Added debug mode support to aid in operational logging and diagnostics.
- **1.1.1**:
  - Add label `logzio_agent_version` to metrics.
- **1.1.0**:
  - Support adding all namespaces by setting `AWS_NAMESPACES` with `all-namespaces`.
  - Fix EBS service details
- **1.0.0**: Initial release.
