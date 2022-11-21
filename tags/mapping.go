package main

func namespacesToServices(namespaces []string) []*string {
	services := make([]*string, 0)
	ns2sMap := getServicesMap()
	for _, ns := range namespaces {
		if serviceArr, ok := ns2sMap[ns]; ok {
			for _, service := range serviceArr {
				servicePtr := new(string)
				*servicePtr = service
				services = append(services, servicePtr)
			}
		} else {
			logger.Println("namespace ", ns, " does not appear in service map. will not try to retrieve its tags")
		}
	}

	return services
}

// The string for each service name and resource type is the same as that embedded in a resource’s Amazon Resource Name (ARN).
// https://awscli.amazonaws.com/v2/documentation/api/latest/reference/resourcegroupstaggingapi/get-resources.html
// if namespace not in map - tags are not available for it
// https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/supported-services.html
func getServicesMap() map[string][]string {

	return map[string][]string{
		"AWS/ApiGateway":         {"apigateway"},
		"AWS/AppStream":          {"appstream"},
		"AWS/AppSync":            {"appsync"},
		"AWS/Athena":             {"athena"},
		"AWS/RDS":                {"rds"},
		"AWS/Backup":             {"backup"},
		"AWS/CertificateManager": {"acm"},
		"AWS/ACMPrivateCA":       {"acm-pca"},
		"AWS/CloudFront":         {"cloudfront"},
		"AWS/CloudHSM":           {"cloudhsm"},
		"AWS/CloudTrail":         {"cloudtrail"},
		"CloudWatchSynthetics":   {"synthetics"},
		"AWS/Logs":               {"logs"},
		"AWS/CodeBuild":          {"codebuild"},
		"AWS/CodeGuruProfiler":   {"codeguru-profiler"},
		"AWS/Cognito":            {"cognito-identity", "cognito-idp"},
		"AWS/Connect":            {"connect"},
		"AWS/DataSync":           {"datasync"},
		"AWS/DMS":                {"dms"},
		"AWS/DX":                 {"directconnect"},
		"AWS/DynamoDB":           {"dynamodb"},
		"AWS/EC2":                {"ec2"},
		"AWS/AutoScaling":        {"autoscaling"},
		"AWS/ElasticBeanstalk":   {"elasticbeanstalk"},
		"AWS/EBS":                {"ebs"},
		"AWS/ECS":                {"ecs"},
		"AWS/EFS":                {"elasticfilesystem"},
		"AWS/ElasticInference":   {"elastic-inference"},
		"AWS/ApplicationELB":     {"elasticloadbalancing:loadbalancer/app"},
		"AWS/NetworkELB":         {"elasticloadbalancing:loadbalancer/net"},
		"AWS/ELB":                {"elasticloadbalancing"},
		"AWS/ElastiCache":        {"elasticache"},
		"AWS/ES":                 {"es"},
		"AWS/ElasticMapReduce":   {"emr-serverless"},
		"AWS/MediaLive":          {"medialive"},
		"AWS/MediaPackage":       {"mediapackage"},
		"AWS/MediaTailor":        {"mediatailor"},
		"AWS/Events":             {"schemas"},
		"AWS/FSx":                {"fsx"},
		"AWS/GameLift":           {"gamelift"},
		"AWS/GlobalAccelerator":  {"globalaccelerator"},
		"Glue":                   {"glue"},
		"AWS/GroundStation":      {"groundstation"},
		"AWS/Inspector":          {"inspector"},
		"AWS/IVS":                {"ivs"},
		"AWS/IoTAnalytics":       {"iotanalytics"},
		"AWS/IoTSiteWise":        {"iotsitewise"},
		"AWS/IoTTwinMaker":       {"iottwinmaker"},
		"AWS/KMS":                {"kms"},
		"AWS/KinesisAnalytics":   {"kinesisanalytics"},
		"AWS/Firehose":           {"firehose"},
		"AWS/Kinesis":            {"kinesis"},
		"AWS/Lambda":             {"lambda"},
		"AWS/Lex":                {"lex"},
		"AWS/ML":                 {"machinelearning"},
		"AWS/Kafka":              {"kafka"},
		"AWS/AmazonMQ":           {"mq"},
		"AWS/Neptune":            {"neptune-db"},
		"AWS/NetworkManager":     {"networkmanager"},
		"AWS/OpsWorks":           {"opsworks"},
		"AWS/QLDB":               {"qldb"},
		"AWS/Redshift":           {"redshift"},
		"AWS/Robomaker":          {"robomaker"},
		"AWS/Route53":            {"route53"},
		"AWS/SageMaker":          {"sagemaker"},
		"AWS/SecretsManager":     {"secretsmanager"},
		"AWS/ServiceCatalog":     {"servicecatalog"},
		"AWS/SES":                {"ses"},
		"AWS/SNS":                {"sns"},
		"AWS/SQS":                {"sqs"},
		"AWS/S3":                 {"s3"},
		"AWS/SWF":                {"swf"},
		"AWS/States":             {"states"},
		"AWS/StorageGateway":     {"storagegateway"},
		"AWS/Transfer":           {"transfer"},
		"WAF":                    {"waf"},
		"AWS/WAFV2":              {"wafv2"},
		"AWS/WorkSpaces":         {"workspaces"},
	}
}
