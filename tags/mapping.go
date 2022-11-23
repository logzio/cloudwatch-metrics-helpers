package main

const (
	arnRegApiGateway           = "^arn\\:.+\\:apigateway\\:.+$"
	arnRegAppStream            = "^arn\\:.+\\:appstream\\:.+$"
	arnRegAppSync              = "^arn\\:.+\\:appsync\\:.+$"
	arnRegAthena               = "^arn\\:.+\\:athena\\:.+$"
	arnRegRds                  = "^arn\\:.+\\:rds\\:.+$"
	arnRegBackup               = "^arn\\:.+\\:backup\\:.+$:"
	arnRegAcm                  = "^arn\\:.+\\:acm\\:.+$"
	arnRegAcmPca               = "^arn\\:.+\\:acm-pca\\:.+$"
	arnRegCloudfront           = "^arn\\:.+\\:cloudfront\\:.+$"
	arnRegCloudHsm             = "^arn\\:.+\\:cloudhsm\\:.+$"
	arnRegCloudtrail           = "^arn\\:.+\\:cloudtrail\\:.+$"
	arnRegCloudwatchSynthetics = "^arn\\:.+\\:synthetics\\:.+$"
	arnRegCloudwatchLogs       = "^arn\\:.+\\:logs\\:.+$"
	arnRegCodeBuild            = "^arn\\:.+\\:codebuild\\:.+$"
	arnRegCodeGuru             = "^arn\\:.+\\:codeguru-profiler\\:.+$"
	arnRegCognitoIdentity      = "^arn\\:.+\\:cognito-identity\\:.+$"
	arnRegCognitoIdp           = "^arn\\:.+\\:cognito-idp\\:.+$"
	arnRegConnect              = "^arn\\:.+\\:connect\\:.+$"
	arnRegDataSync             = "^arn\\:.+\\:datasync\\:.+$"
	arnRegDms                  = "^arn\\:.+\\:dms\\:.+$"
	arnRegDirectConnect        = "^arn\\:.+\\:directconnect\\:.+$"
	arnRegDynamoDb             = "^arn\\:.+\\:dynamodb\\:.+$"
	arnRegEc2                  = "^arn\\:.+\\:ec2\\:.+\\:.+$"
	arnRegAutoScaling          = "^arn\\:.+\\:autoscaling\\:.+$"
	arnRegElasticbeanstalk     = "^arn\\:.+\\:elasticbeanstalk\\:.+$"
	arnRegEbs                  = "^arn\\:.+\\:ec2\\:.+\\:.+\\:volume\\/.+$"
	arnRegEcs                  = "^arn\\:.+\\:ecs\\:.+$"
	arnRegEfs                  = "^arn\\:.+\\:elasticfilesystem\\:.+$"
	arnRegElasticInference     = "^arn\\:.+\\:elastic-inference\\:.+$"
	arnRegAppElb               = "^arn\\:.+\\:elasticloadbalancing\\:.+\\:.+\\:loadbalancer\\/app\\/.+$"
	arnRegNetElb               = "^arn\\:.+\\:elasticloadbalancing\\:.+\\:.+\\:loadbalancer\\/net\\/.+$"
	arnRegElb                  = "^arn\\:.+\\:elasticloadbalancing\\:.+\\:.+\\:loadbalancer\\/[\\d\\w\\-]$"
	arnRegElasticache          = "^arn\\:.+\\:elasticache\\:.+$"
	arnRegEs                   = "^arn\\:.+\\:es\\:.+$"
	arnRegElasticMapReduce     = "^arn\\:.+\\:elasticmapreduce\\:.+$"
	arnRegMediaLive            = "^arn\\:.+\\:medialive\\:.+$"
	arnRegMediaPackage         = "^arn\\:.+\\:mediapackage\\:.+$"
	arnRegMediaTailor          = "^arn\\:.+\\:mediatailor\\:.+$"
	arnRegEventbridge          = "^arn\\:.+\\:events\\:.+$"
	arnRegFsx                  = "^arn\\:.+\\:fsx\\:.+$"
	arnRegGameLift             = "^arn\\:.+\\:gamelift\\:.+$"
	arnRegGlobalAccelerator    = "^arn\\:.+\\:globalaccelerator\\:.+$"
	arnRegGlue                 = "^arn\\:.+\\:glue\\:.+$"
	arnRegInspector            = "^arn\\:.+\\:inspector2\\:.+$"
	arnRegIvs                  = "^arn\\:.+\\:ivs\\:.+$"
	arnRegIotAnalytics         = "^arn\\:.+\\:iotanalytics\\:.+$"
	arnRegIotSiteWise          = "^arn\\:.+\\:iotsitewise\\:.+$"
	arnRegIotTwinMaker         = "^arn\\:.+\\:iottwinmaker\\:.+$"
	arnRegKms                  = "^arn\\:.+\\:kms\\:.+$"
	arnRegKinesisAnalytics     = "^arn\\:.+\\:kinesisanalytics\\:.+$"
	arnRegFirehose             = "^arn\\:.+\\:firehose\\:.+$"
	arnRegKinesis              = "^arn\\:.+\\:kinesis\\:.+$"
	arnRegLambda               = "^arn\\:.+\\:lambda\\:.+$"
	arnRegLex                  = "^arn\\:.+\\:lex\\:.+$"
	arnRegMl                   = "^arn\\:.+\\:machinelearning\\:.+$"
	arnRegKafka                = "^arn\\:.+\\:kafka\\:.+$"
	arnRegMq                   = "^arn\\:.+\\:mq\\:.+$"
	arnRegNeptune              = "^arn\\:.+\\:neptune-db\\:.+$"
	arnRegNetworkManager       = "^arn\\:.+\\:networkmanager\\:.+$"
	arnRegOpsWorks             = "^arn\\:.+\\:opsworks\\:.+$"
	arnRegQldb                 = "^arn\\:.+\\:qldb\\:.+$"
	arnRegRedshift             = "^arn\\:.+\\:redshift\\:.+$"
	arnRegRoboMaker            = "^arn\\:.+\\:robomaker\\:.+$"
	arnRegRoute53              = "^arn\\:.+\\:route53\\:.+$"
	arnRegSageMaker            = "^arn\\:.+\\:sagemaker\\:.+$"
	arnRegSecretManager        = "^arn\\:.+\\:secretsmanager\\:.+$"
	arnRegServiceCatalog       = "^arn\\:.+\\:servicecatalog\\:.+$"
	arnRegSes                  = "^arn\\:.+\\:ses\\:.+$"
	arnRegSns                  = "^arn\\:.+\\:sns\\:.+$"
	arnRegSqs                  = "^arn\\:.+\\:sqs\\:.+$"
	arnRegS3                   = "^arn\\:.+\\:s3\\:.+$"
	arnRegSwf                  = "^arn\\:.+\\:swf\\:.+$"
	arnRegStepFunction         = "^arn\\:.+\\:states\\:.+$"
	arnRegStorageGateway       = "^arn\\:.+\\:storagegateway\\:.+$"
	arnRegTransfer             = "^arn\\:.+:transfer\\:.+$"
	arnRegWaf                  = "^arn\\:.+\\:waf\\:.+$"
	arnRegWaf2                 = "^arn\\:.+\\:wafv2\\:.+$"
	arnRegWorkSpaces           = "^arn\\:.+\\:workspaces\\:.+$"
)

func namespacesToServices(namespaces []string) []*string {
	services := make([]*string, 0)
	ns2sMap := getNamespacesToServicesMap()
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
func getNamespacesToServicesMap() map[string][]string {
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
		"AWS/Inspector":          {"inspector", "inspector2"},
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

func getArnRegexToIdsMap() map[string][]string {
	return map[string][]string{
		arnRegApiGateway:           {"ApiId"},
		arnRegAppStream:            {"Fleet"},
		arnRegAppSync:              {"GraphQLAPIId"},
		arnRegAthena:               {"WorkGroup"},
		arnRegRds:                  {"DBClusterIdentifier"},
		arnRegBackup:               {"BackupVaultName"},
		arnRegAcm:                  {"CertificateArn"},
		arnRegAcmPca:               {"PrivateCAArn"},
		arnRegCloudfront:           {"DistributionId"},
		arnRegCloudHsm:             {"HsmId"},
		arnRegCloudtrail:           {},
		arnRegCloudwatchSynthetics: {"CanaryName"},
		arnRegCloudwatchLogs:       {"LogGroupName"},
		arnRegCodeBuild:            {"BuildId"},
		arnRegCodeGuru:             {},
		arnRegCognitoIdentity:      {"UserPoolId"},
		arnRegCognitoIdp:           {"UserPoolId"},
		arnRegConnect:              {"InstanceId"},
		arnRegDataSync:             {"TaskId"},
		arnRegDms:                  {"ReplicationTaskIdentifier"},
		arnRegDirectConnect:        {"ConnectionId"},
		arnRegDynamoDb:             {"TableName"},
		arnRegEc2:                  {"InstanceId"},
		arnRegAutoScaling:          {"AutoScalingGroupName"},
		arnRegElasticbeanstalk:     {"InstanceId"},
		arnRegEbs:                  {"VolumeId"},
		arnRegEcs:                  {"ServiceName"},
		arnRegEfs:                  {"FileSystemId"},
		arnRegElasticInference:     {"ElasticInferenceAcceleratorId"},
		arnRegAppElb:               {"LoadBalancer"},
		arnRegNetElb:               {"LoadBalancer"},
		arnRegElb:                  {"LoadBalancerName"},
		arnRegElasticache:          {"CacheClusterId"},
		arnRegEs:                   {"ClientId"},
		arnRegElasticMapReduce:     {"JobFlowId"},
		arnRegMediaLive:            {"ChannelID"},
		arnRegMediaPackage:         {"Channel"},
		arnRegMediaTailor:          {"Configuration Name"},
		arnRegEventbridge:          {"RuleName"},
		arnRegFsx:                  {"FileSystemId"},
		arnRegGameLift:             {"FleetId"},
		arnRegGlobalAccelerator:    {},
		arnRegGlue:                 {"JobName"},
		arnRegInspector:            {},
		arnRegIvs:                  {"Channel"},
		arnRegIotAnalytics:         {"DatasetName"},
		arnRegIotSiteWise:          {},
		arnRegIotTwinMaker:         {"WorkspaceId"},
		arnRegKms:                  {"KeyId"},
		arnRegKinesisAnalytics:     {"Id"},
		arnRegFirehose:             {"DeliveryStreamName"},
		arnRegKinesis:              {"StreamName"},
		arnRegLambda:               {"FunctionName"},
		arnRegLex:                  {"BotName"},
		arnRegMl:                   {},
		arnRegKafka:                {"Cluster Name"},
		arnRegMq:                   {"Broker"},
		arnRegNeptune:              {"DBClusterIdentifier"},
		arnRegNetworkManager:       {"FirewallName"},
		arnRegOpsWorks:             {"StackId"},
		arnRegQldb:                 {"LedgerName"},
		arnRegRedshift:             {"ClusterIdentifier"},
		arnRegRoboMaker:            {"SimulationJobId"},
		arnRegRoute53:              {"HostedZoneId"},
		arnRegSageMaker:            {"EndpointName"},
		arnRegSecretManager:        {"Service"},
		arnRegServiceCatalog:       {"ProductId"},
		arnRegSes:                  {},
		arnRegSns:                  {"TopicName"},
		arnRegSqs:                  {"QueueName"},
		arnRegS3:                   {"BucketName"},
		arnRegSwf:                  {"ActivityTypeName"},
		arnRegStepFunction:         {"ActivityArn"},
		arnRegStorageGateway:       {"GatewayName", "GatewayId"},
		arnRegTransfer:             {"ServerId"},
		arnRegWaf:                  {"WebACL"},
		arnRegWaf2:                 {"WebACL"},
		arnRegWorkSpaces:           {"WorkspaceId", "DirectoryId"},
	}
}
