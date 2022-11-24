### Logz.io Metric Stream Helper - Tags

This Lambda collects tags from resources under selected AWS namespaces, and sends them as metrics to Logz.io.

#### Environment variables:

**All environment variables mentioned here are required!**

| Name                      | Description                                                                                                                                                                                                   |
|---------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `AWS_NAMESPACES`          | Comma-separated list of the AWS namespaces to collect metrics from. Possible values can be found [here](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html). |
| `LOGZIO_METRICS_LISTENER` | Full URL for the Logz.io metrics listener, for example `https://listener.logz.io:8053`                                                                                                                                |
| `LOGZIO_METRICS_TOKEN`    | Logz.io metrics shipping token                                                                                                                                                                                |


#### Changelog:

- **1.0.0**: Initial release.
