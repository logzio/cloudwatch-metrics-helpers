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

	nsAll                      = "all-namespaces"
	nsAmplify                  = "AWS/AmplifyHosting"
	nsApiGateway               = "AWS/ApiGateway"
	nsAppRunner                = "AWS/AppRunner"
	nsAppStream                = "AWS/AppStream"
	nsAppSync                  = "AWS/AppSync"
	nsAthena                   = "AWS/Athena"
	nsRds                      = "AWS/RDS"
	nsBackup                   = "AWS/Backup"
	nsBilling                  = "AWS/Billing"
	nsAcm                      = "AWS/CertificateManager"
	nsAcmPca                   = "AWS/ACMPrivateCA"
	nsChatBot                  = "AWS/Chatbot"
	nsChimeSdk                 = "AWS/ChimeSDK"
	nsClientVpn                = "AWS/ClientVPN"
	nsCloudFront               = "AWS/CloudFront"
	nsCloudHsm                 = "AWS/CloudHSM"
	nsCloudSearch              = "AWS/CloudSearch"
	nsCloudTrail               = "AWS/CloudTrail"
	nsCwAgent                  = "CWAgent"
	nsMetricStream             = "AWS/CloudWatch/MetricStreams"
	nsRum                      = "AWS/RUM"
	nsCwSynthetics             = "CloudWatchSynthetics"
	nsLogs                     = "AWS/Logs"
	nsCodeBuild                = "AWS/CodeBuild"
	nsCodeGuroProfiler         = "AWS/CodeGuruProfiler"
	nsCognito                  = "AWS/Cognito"
	nsConnect                  = "AWS/Connect"
	nsDataLifecycleManager     = "AWS/DataLifecycleManager"
	nsDataSync                 = "AWS/DataSync"
	nsDms                      = "AWS/DMS"
	nsDx                       = "AWS/DX"
	nsDocDb                    = "AWS/DocDB"
	nsDynamoDb                 = "AWS/DynamoDB"
	nsDax                      = "AWS/DAX"
	nsEc2                      = "AWS/EC2"
	nsElasticGpus              = "AWS/ElasticGPUs"
	nsEc2Spot                  = "AWS/EC2Spot"
	nsAutoScaling              = "AWS/AutoScaling"
	nsElasticBeanstalk         = "AWS/ElasticBeanstalk"
	nsEbs                      = "AWS/EBS"
	nsEcs                      = "AWS/ECS"
	nsEcsManagedScaling        = "AWS/ECS/ManagedScaling"
	nsEfs                      = "AWS/EFS"
	nsElasticInference         = "AWS/ElasticInference"
	nsAppElb                   = "AWS/ApplicationELB"
	nsNetElb                   = "AWS/NetworkELB"
	nsGatewayElb               = "AWS/GatewayELB"
	nsElb                      = "AWS/ELB"
	nsElasticTranscoder        = "AWS/ElasticTranscoder"
	nsElastiCache              = "AWS/ElastiCache"
	nsEs                       = "AWS/ES"
	nsEmr                      = "AWS/ElasticMapReduce"
	nsMediaConnect             = "AWS/MediaConnect"
	nsMediaConvert             = "AWS/MediaConvert"
	nsMediaLive                = "AWS/MediaLive"
	nsMediaPackage             = "AWS/MediaPackage"
	nsMediaStore               = "AWS/MediaStore"
	nsMediaTailor              = "AWS/MediaTailor"
	nsEvents                   = "AWS/Events"
	nsFsx                      = "AWS/FSx"
	nsGameLift                 = "AWS/GameLift"
	nsGlobalAccelerator        = "AWS/GlobalAccelerator"
	nsGlue                     = "Glue"
	nsGroundStation            = "AWS/GroundStation"
	nsHealthLake               = "AWS/HealthLake"
	nsInspector                = "AWS/Inspector"
	nsIvs                      = "AWS/IVS"
	nsIvsChat                  = "AWS/IVSChat"
	nsIot                      = "AWS/IoT"
	nsIotAnalytics             = "AWS/IoTAnalytics"
	nsIotSiteWise              = "AWS/IoTSiteWise"
	nsThingsGraph              = "AWS/ThingsGraph"
	nsIotTwinMaker             = "AWS/IoTTwinMaker"
	nsKms                      = "AWS/KMS"
	nsCassandra                = "AWS/Cassandra"
	nsKinesisAnalytics         = "AWS/KinesisAnalytics"
	nsFirehose                 = "AWS/Firehose"
	nsKinesis                  = "AWS/Kinesis"
	nsKinesisVideo             = "AWS/KinesisVideo"
	nsLambda                   = "AWS/Lambda"
	nsLex                      = "AWS/Lex"
	nsLocation                 = "AWS/Location"
	nsLookoutMetrics           = "AWS/LookoutMetrics"
	nsMl                       = "AWS/ML"
	nsKafka                    = "AWS/Kafka"
	nsMq                       = "AWS/AmazonMQ"
	nsNeptune                  = "AWS/Neptune"
	nsNetworkFirewall          = "AWS/NetworkFirewall"
	nsNetworkManager           = "AWS/NetworkManager"
	nsNimbleStudio             = "AWS/NimbleStudio"
	nsOpsWorks                 = "AWS/OpsWorks"
	nsPolly                    = "AWS/Polly"
	nsPrivateLinkEndpoints     = "AWS/PrivateLinkEndpoints"
	nsPrivateLinkServices      = "AWS/PrivateLinkServices"
	nsQldb                     = "AWS/QLDB"
	nsQuickSight               = "AWS/QuickSight"
	nsRedshift                 = "AWS/Redshift"
	nsRoboMaker                = "AWS/Robomaker"
	nsRoute53                  = "AWS/Route53"
	nsRoute53RecoveryReadiness = "AWS/Route53RecoveryReadiness"
	nsSageMaker                = "AWS/SageMaker"
	nsSageMakerMBP             = "AWS/SageMaker/ModelBuildingPipeline"
	nsSecretsManager           = "AWS/SecretsManager"
	nsServiceCatalog           = "AWS/ServiceCatalog"
	nsShieldAdvanced           = "AWS/DDoSProtection"
	nsSes                      = "AWS/SES"
	nsSns                      = "AWS/SNS"
	nsSqs                      = "AWS/SQS"
	nsS3                       = "AWS/S3"
	nsS3StorageLens            = "AWS/S3/Storage-Lens"
	nsSwf                      = "AWS/SWF"
	nsStates                   = "AWS/States"
	nsStorageGateway           = "AWS/StorageGateway"
	nsSsmRunCommand            = "AWS/SSM-RunCommand"
	nsTextract                 = "AWS/Textract"
	nsTimestream               = "AWS/Timestream"
	nsTransfer                 = "AWS/Transfer"
	nsTranslate                = "AWS/Translate"
	nsTrustedAdvisor           = "AWS/TrustedAdvisor"
	nsVpcNatgateway            = "AWS/NATGateway"
	nsTransitGateway           = "AWS/TransitGateway"
	nsVpn                      = "AWS/VPN"
	nsIpam                     = "AWS/IPAM"
	nsWaf                      = "WAF"
	nsWaf2                     = "AWS/WAFV2"
	nsWorkMail                 = "AWS/WorkMail"
	nsWorkspaces               = "AWS/WorkSpaces"
	nsWorkSpacesWeb            = "AWS/WorkSpacesWeb"
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
		nsApiGateway:        {"apigateway"},
		nsAppStream:         {"appstream"},
		nsAppSync:           {"appsync"},
		nsAthena:            {"athena"},
		nsRds:               {"rds"},
		nsBackup:            {"backup"},
		nsAcm:               {"acm"},
		nsAcmPca:            {"acm-pca"},
		nsCloudFront:        {"cloudfront"},
		nsCloudHsm:          {"cloudhsm"},
		nsCloudTrail:        {"cloudtrail"},
		nsCwSynthetics:      {"synthetics"},
		nsLogs:              {"logs"},
		nsCodeBuild:         {"codebuild"},
		nsCodeGuroProfiler:  {"codeguru-profiler"},
		nsCognito:           {"cognito-identity", "cognito-idp"},
		nsConnect:           {"connect"},
		nsDataSync:          {"datasync"},
		nsDms:               {"dms"},
		nsDx:                {"directconnect"},
		nsDynamoDb:          {"dynamodb"},
		nsEc2:               {"ec2"},
		nsAutoScaling:       {"autoscaling"},
		nsElasticBeanstalk:  {"elasticbeanstalk"},
		nsEbs:               {"ec2:volume"},
		nsEcs:               {"ecs"},
		nsEfs:               {"elasticfilesystem"},
		nsElasticInference:  {"elastic-inference"},
		nsAppElb:            {"elasticloadbalancing:loadbalancer/app"},
		nsNetElb:            {"elasticloadbalancing:loadbalancer/net"},
		nsElb:               {"elasticloadbalancing"},
		nsElastiCache:       {"elasticache"},
		nsEs:                {"es"},
		nsEmr:               {"emr-serverless"},
		nsMediaLive:         {"medialive"},
		nsMediaPackage:      {"mediapackage"},
		nsMediaTailor:       {"mediatailor"},
		nsEvents:            {"schemas"},
		nsFsx:               {"fsx"},
		nsGameLift:          {"gamelift"},
		nsGlobalAccelerator: {"globalaccelerator"},
		nsGlue:              {"glue"},
		nsGroundStation:     {"groundstation"},
		nsInspector:         {"inspector", "inspector2"},
		nsIvs:               {"ivs"},
		nsIotAnalytics:      {"iotanalytics"},
		nsIotSiteWise:       {"iotsitewise"},
		nsIotTwinMaker:      {"iottwinmaker"},
		nsKms:               {"kms"},
		nsKinesisAnalytics:  {"kinesisanalytics"},
		nsFirehose:          {"firehose"},
		nsKinesis:           {"kinesis"},
		nsLambda:            {"lambda"},
		nsLex:               {"lex"},
		nsMl:                {"machinelearning"},
		nsKafka:             {"kafka"},
		nsMq:                {"mq"},
		nsNetworkManager:    {"networkmanager"},
		nsOpsWorks:          {"opsworks"},
		nsQldb:              {"qldb"},
		nsRedshift:          {"redshift"},
		nsRoboMaker:         {"robomaker"},
		nsRoute53:           {"route53"},
		nsSageMaker:         {"sagemaker"},
		nsSecretsManager:    {"secretsmanager"},
		nsServiceCatalog:    {"servicecatalog"},
		nsSes:               {"ses"},
		nsSns:               {"sns"},
		nsSqs:               {"sqs"},
		nsS3:                {"s3"},
		nsSwf:               {"swf"},
		nsStates:            {"states"},
		nsStorageGateway:    {"storagegateway"},
		nsTransfer:          {"transfer"},
		nsWaf:               {"waf"},
		nsWaf2:              {"wafv2"},
		nsWorkspaces:        {"workspaces"},
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

func getAllNamespaces() []string {
	return []string{
		nsAmplify,
		nsApiGateway,
		nsAppRunner,
		nsAppStream,
		nsAppSync,
		nsAthena,
		nsRds,
		nsBackup,
		nsBilling,
		nsAcm,
		nsAcmPca,
		nsChatBot,
		nsChimeSdk,
		nsClientVpn,
		nsCloudFront,
		nsCloudHsm,
		nsCloudSearch,
		nsCloudTrail,
		nsCwAgent,
		nsMetricStream,
		nsRum,
		nsCwSynthetics,
		nsLogs,
		nsCodeBuild,
		nsCodeGuroProfiler,
		nsCognito,
		nsConnect,
		nsDataLifecycleManager,
		nsDataSync,
		nsDms,
		nsDx,
		nsDocDb,
		nsDynamoDb,
		nsDax,
		nsEc2,
		nsElasticGpus,
		nsEc2Spot,
		nsAutoScaling,
		nsElasticBeanstalk,
		nsEbs,
		nsEcs,
		nsEcsManagedScaling,
		nsEfs,
		nsElasticInference,
		nsAppElb,
		nsNetElb,
		nsGatewayElb,
		nsElb,
		nsElasticTranscoder,
		nsElastiCache,
		nsEs,
		nsEmr,
		nsMediaConnect,
		nsMediaConvert,
		nsMediaLive,
		nsMediaPackage,
		nsMediaStore,
		nsMediaTailor,
		nsEvents,
		nsFsx,
		nsGameLift,
		nsGlobalAccelerator,
		nsGlue,
		nsGroundStation,
		nsHealthLake,
		nsInspector,
		nsIvs,
		nsIvsChat,
		nsIot,
		nsIotAnalytics,
		nsIotSiteWise,
		nsThingsGraph,
		nsIotTwinMaker,
		nsKms,
		nsCassandra,
		nsKinesisAnalytics,
		nsFirehose,
		nsKinesis,
		nsKinesisVideo,
		nsLambda,
		nsLex,
		nsLocation,
		nsLookoutMetrics,
		nsMl,
		nsKafka,
		nsMq,
		nsNeptune,
		nsNetworkFirewall,
		nsNetworkManager,
		nsNimbleStudio,
		nsOpsWorks,
		nsPolly,
		nsPrivateLinkEndpoints,
		nsPrivateLinkServices,
		nsQldb,
		nsQuickSight,
		nsRedshift,
		nsRoboMaker,
		nsRoute53,
		nsRoute53RecoveryReadiness,
		nsSageMaker,
		nsSageMakerMBP,
		nsSecretsManager,
		nsServiceCatalog,
		nsShieldAdvanced,
		nsSes,
		nsSns,
		nsSqs,
		nsS3,
		nsS3StorageLens,
		nsSwf,
		nsStates,
		nsStorageGateway,
		nsSsmRunCommand,
		nsTextract,
		nsTimestream,
		nsTransfer,
		nsTranslate,
		nsTrustedAdvisor,
		nsVpcNatgateway,
		nsTransitGateway,
		nsVpn,
		nsIpam,
		nsWaf,
		nsWaf2,
		nsWorkMail,
		nsWorkspaces,
		nsWorkSpacesWeb,
	}
}
