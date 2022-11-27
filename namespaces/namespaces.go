package main

const (
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
