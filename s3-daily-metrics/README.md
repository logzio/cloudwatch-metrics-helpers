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
- `P8S_LOGZIO_NAME`: The environment identifier

### Minimum AWS IAM Permissions Required
For the AWS Lambda function to be able to collect CloudWatch metrics and list S3 buckets, it will need the following permissions:

- `cloudwatch:GetMetricData`: Allows the function to retrieve CloudWatch metrics.
- `s3:ListBuckets`: Allows the function to list the S3 buckets in the account.
You can grant these permissions to the function by attaching the following IAM policy to the IAM role that is associated with the function:

```shell
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "CloudWatchMetrics",
            "Effect": "Allow",
            "Action": [
                "cloudwatch:GetMetricData"
            ],
            "Resource": "*"
        },
        {
            "Sid": "CloudWatchMetrics",
            "Effect": "Allow",
            "Action": [
                "cloudwatch:GetMetricData"
            ],
            "Resource": "*"
        },
        {
            "Sid": "S3ListBuckets",
            "Effect": "Allow",
            "Action": [
                "s3:ListBuckets",
                "s3:ListAllMyBuckets"
            ],
            "Resource": "*"
        }
    ]
}
```
