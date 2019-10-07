package cloudformation

// AllResources fetches an iterable map all CloudFormation and SAM resources
func AllResources() map[string]interface{} {
	return map[string]interface{}{
		"AWS::AmazonMQ::Broker":                                 AWSAmazonMQBroker{},
		"AWS::AmazonMQ::Configuration":                          AWSAmazonMQConfiguration{},
		"AWS::ApiGateway::Account":                              AWSApiGatewayAccount{},
		"AWS::ApiGateway::ApiKey":                               AWSApiGatewayApiKey{},
		"AWS::ApiGateway::Authorizer":                           AWSApiGatewayAuthorizer{},
		"AWS::ApiGateway::BasePathMapping":                      AWSApiGatewayBasePathMapping{},
		"AWS::ApiGateway::ClientCertificate":                    AWSApiGatewayClientCertificate{},
		"AWS::ApiGateway::Deployment":                           AWSApiGatewayDeployment{},
		"AWS::ApiGateway::DocumentationPart":                    AWSApiGatewayDocumentationPart{},
		"AWS::ApiGateway::DocumentationVersion":                 AWSApiGatewayDocumentationVersion{},
		"AWS::ApiGateway::DomainName":                           AWSApiGatewayDomainName{},
		"AWS::ApiGateway::GatewayResponse":                      AWSApiGatewayGatewayResponse{},
		"AWS::ApiGateway::Method":                               AWSApiGatewayMethod{},
		"AWS::ApiGateway::Model":                                AWSApiGatewayModel{},
		"AWS::ApiGateway::RequestValidator":                     AWSApiGatewayRequestValidator{},
		"AWS::ApiGateway::Resource":                             AWSApiGatewayResource{},
		"AWS::ApiGateway::RestApi":                              AWSApiGatewayRestApi{},
		"AWS::ApiGateway::Stage":                                AWSApiGatewayStage{},
		"AWS::ApiGateway::UsagePlan":                            AWSApiGatewayUsagePlan{},
		"AWS::ApiGateway::UsagePlanKey":                         AWSApiGatewayUsagePlanKey{},
		"AWS::ApiGateway::VpcLink":                              AWSApiGatewayVpcLink{},
		"AWS::AppStream::DirectoryConfig":                       AWSAppStreamDirectoryConfig{},
		"AWS::AppStream::Fleet":                                 AWSAppStreamFleet{},
		"AWS::AppStream::ImageBuilder":                          AWSAppStreamImageBuilder{},
		"AWS::AppStream::Stack":                                 AWSAppStreamStack{},
		"AWS::AppStream::StackFleetAssociation":                 AWSAppStreamStackFleetAssociation{},
		"AWS::AppStream::StackUserAssociation":                  AWSAppStreamStackUserAssociation{},
		"AWS::AppStream::User":                                  AWSAppStreamUser{},
		"AWS::AppSync::ApiKey":                                  AWSAppSyncApiKey{},
		"AWS::AppSync::DataSource":                              AWSAppSyncDataSource{},
		"AWS::AppSync::GraphQLApi":                              AWSAppSyncGraphQLApi{},
		"AWS::AppSync::GraphQLSchema":                           AWSAppSyncGraphQLSchema{},
		"AWS::AppSync::Resolver":                                AWSAppSyncResolver{},
		"AWS::ApplicationAutoScaling::ScalableTarget":           AWSApplicationAutoScalingScalableTarget{},
		"AWS::ApplicationAutoScaling::ScalingPolicy":            AWSApplicationAutoScalingScalingPolicy{},
		"AWS::Athena::NamedQuery":                               AWSAthenaNamedQuery{},
		"AWS::AutoScaling::AutoScalingGroup":                    AWSAutoScalingAutoScalingGroup{},
		"AWS::AutoScaling::LaunchConfiguration":                 AWSAutoScalingLaunchConfiguration{},
		"AWS::AutoScaling::LifecycleHook":                       AWSAutoScalingLifecycleHook{},
		"AWS::AutoScaling::ScalingPolicy":                       AWSAutoScalingScalingPolicy{},
		"AWS::AutoScaling::ScheduledAction":                     AWSAutoScalingScheduledAction{},
		"AWS::AutoScalingPlans::ScalingPlan":                    AWSAutoScalingPlansScalingPlan{},
		"AWS::Batch::ComputeEnvironment":                        AWSBatchComputeEnvironment{},
		"AWS::Batch::JobDefinition":                             AWSBatchJobDefinition{},
		"AWS::Batch::JobQueue":                                  AWSBatchJobQueue{},
		"AWS::Budgets::Budget":                                  AWSBudgetsBudget{},
		"AWS::CertificateManager::Certificate":                  AWSCertificateManagerCertificate{},
		"AWS::Cloud9::EnvironmentEC2":                           AWSCloud9EnvironmentEC2{},
		"AWS::CloudFormation::CustomResource":                   AWSCloudFormationCustomResource{},
		"AWS::CloudFormation::Stack":                            AWSCloudFormationStack{},
		"AWS::CloudFormation::WaitCondition":                    AWSCloudFormationWaitCondition{},
		"AWS::CloudFormation::WaitConditionHandle":              AWSCloudFormationWaitConditionHandle{},
		"AWS::CloudFront::CloudFrontOriginAccessIdentity":       AWSCloudFrontCloudFrontOriginAccessIdentity{},
		"AWS::CloudFront::Distribution":                         AWSCloudFrontDistribution{},
		"AWS::CloudFront::StreamingDistribution":                AWSCloudFrontStreamingDistribution{},
		"AWS::CloudTrail::Trail":                                AWSCloudTrailTrail{},
		"AWS::CloudWatch::Alarm":                                AWSCloudWatchAlarm{},
		"AWS::CloudWatch::Dashboard":                            AWSCloudWatchDashboard{},
		"AWS::CodeBuild::Project":                               AWSCodeBuildProject{},
		"AWS::CodeCommit::Repository":                           AWSCodeCommitRepository{},
		"AWS::CodeDeploy::Application":                          AWSCodeDeployApplication{},
		"AWS::CodeDeploy::DeploymentConfig":                     AWSCodeDeployDeploymentConfig{},
		"AWS::CodeDeploy::DeploymentGroup":                      AWSCodeDeployDeploymentGroup{},
		"AWS::CodePipeline::CustomActionType":                   AWSCodePipelineCustomActionType{},
		"AWS::CodePipeline::Pipeline":                           AWSCodePipelinePipeline{},
		"AWS::CodePipeline::Webhook":                            AWSCodePipelineWebhook{},
		"AWS::Cognito::IdentityPool":                            AWSCognitoIdentityPool{},
		"AWS::Cognito::IdentityPoolRoleAttachment":              AWSCognitoIdentityPoolRoleAttachment{},
		"AWS::Cognito::UserPool":                                AWSCognitoUserPool{},
		"AWS::Cognito::UserPoolClient":                          AWSCognitoUserPoolClient{},
		"AWS::Cognito::UserPoolGroup":                           AWSCognitoUserPoolGroup{},
		"AWS::Cognito::UserPoolUser":                            AWSCognitoUserPoolUser{},
		"AWS::Cognito::UserPoolUserToGroupAttachment":           AWSCognitoUserPoolUserToGroupAttachment{},
		"AWS::Config::AggregationAuthorization":                 AWSConfigAggregationAuthorization{},
		"AWS::Config::ConfigRule":                               AWSConfigConfigRule{},
		"AWS::Config::ConfigurationAggregator":                  AWSConfigConfigurationAggregator{},
		"AWS::Config::ConfigurationRecorder":                    AWSConfigConfigurationRecorder{},
		"AWS::Config::DeliveryChannel":                          AWSConfigDeliveryChannel{},
		"AWS::DAX::Cluster":                                     AWSDAXCluster{},
		"AWS::DAX::ParameterGroup":                              AWSDAXParameterGroup{},
		"AWS::DAX::SubnetGroup":                                 AWSDAXSubnetGroup{},
		"AWS::DMS::Certificate":                                 AWSDMSCertificate{},
		"AWS::DMS::Endpoint":                                    AWSDMSEndpoint{},
		"AWS::DMS::EventSubscription":                           AWSDMSEventSubscription{},
		"AWS::DMS::ReplicationInstance":                         AWSDMSReplicationInstance{},
		"AWS::DMS::ReplicationSubnetGroup":                      AWSDMSReplicationSubnetGroup{},
		"AWS::DMS::ReplicationTask":                             AWSDMSReplicationTask{},
		"AWS::DataPipeline::Pipeline":                           AWSDataPipelinePipeline{},
		"AWS::DirectoryService::MicrosoftAD":                    AWSDirectoryServiceMicrosoftAD{},
		"AWS::DirectoryService::SimpleAD":                       AWSDirectoryServiceSimpleAD{},
		"AWS::DynamoDB::Table":                                  AWSDynamoDBTable{},
		"AWS::EC2::CustomerGateway":                             AWSEC2CustomerGateway{},
		"AWS::EC2::DHCPOptions":                                 AWSEC2DHCPOptions{},
		"AWS::EC2::EIP":                                         AWSEC2EIP{},
		"AWS::EC2::EIPAssociation":                              AWSEC2EIPAssociation{},
		"AWS::EC2::EgressOnlyInternetGateway":                   AWSEC2EgressOnlyInternetGateway{},
		"AWS::EC2::FlowLog":                                     AWSEC2FlowLog{},
		"AWS::EC2::Host":                                        AWSEC2Host{},
		"AWS::EC2::Instance":                                    AWSEC2Instance{},
		"AWS::EC2::InternetGateway":                             AWSEC2InternetGateway{},
		"AWS::EC2::LaunchTemplate":                              AWSEC2LaunchTemplate{},
		"AWS::EC2::NatGateway":                                  AWSEC2NatGateway{},
		"AWS::EC2::NetworkAcl":                                  AWSEC2NetworkAcl{},
		"AWS::EC2::NetworkAclEntry":                             AWSEC2NetworkAclEntry{},
		"AWS::EC2::NetworkInterface":                            AWSEC2NetworkInterface{},
		"AWS::EC2::NetworkInterfaceAttachment":                  AWSEC2NetworkInterfaceAttachment{},
		"AWS::EC2::NetworkInterfacePermission":                  AWSEC2NetworkInterfacePermission{},
		"AWS::EC2::PlacementGroup":                              AWSEC2PlacementGroup{},
		"AWS::EC2::Route":                                       AWSEC2Route{},
		"AWS::EC2::RouteTable":                                  AWSEC2RouteTable{},
		"AWS::EC2::SecurityGroup":                               AWSEC2SecurityGroup{},
		"AWS::EC2::SecurityGroupEgress":                         AWSEC2SecurityGroupEgress{},
		"AWS::EC2::SecurityGroupIngress":                        AWSEC2SecurityGroupIngress{},
		"AWS::EC2::SpotFleet":                                   AWSEC2SpotFleet{},
		"AWS::EC2::Subnet":                                      AWSEC2Subnet{},
		"AWS::EC2::SubnetCidrBlock":                             AWSEC2SubnetCidrBlock{},
		"AWS::EC2::SubnetNetworkAclAssociation":                 AWSEC2SubnetNetworkAclAssociation{},
		"AWS::EC2::SubnetRouteTableAssociation":                 AWSEC2SubnetRouteTableAssociation{},
		"AWS::EC2::TrunkInterfaceAssociation":                   AWSEC2TrunkInterfaceAssociation{},
		"AWS::EC2::VPC":                                         AWSEC2VPC{},
		"AWS::EC2::VPCCidrBlock":                                AWSEC2VPCCidrBlock{},
		"AWS::EC2::VPCDHCPOptionsAssociation":                   AWSEC2VPCDHCPOptionsAssociation{},
		"AWS::EC2::VPCEndpoint":                                 AWSEC2VPCEndpoint{},
		"AWS::EC2::VPCEndpointConnectionNotification":           AWSEC2VPCEndpointConnectionNotification{},
		"AWS::EC2::VPCEndpointServicePermissions":               AWSEC2VPCEndpointServicePermissions{},
		"AWS::EC2::VPCGatewayAttachment":                        AWSEC2VPCGatewayAttachment{},
		"AWS::EC2::VPCPeeringConnection":                        AWSEC2VPCPeeringConnection{},
		"AWS::EC2::VPNConnection":                               AWSEC2VPNConnection{},
		"AWS::EC2::VPNConnectionRoute":                          AWSEC2VPNConnectionRoute{},
		"AWS::EC2::VPNGateway":                                  AWSEC2VPNGateway{},
		"AWS::EC2::VPNGatewayRoutePropagation":                  AWSEC2VPNGatewayRoutePropagation{},
		"AWS::EC2::Volume":                                      AWSEC2Volume{},
		"AWS::EC2::VolumeAttachment":                            AWSEC2VolumeAttachment{},
		"AWS::ECR::Repository":                                  AWSECRRepository{},
		"AWS::ECS::Cluster":                                     AWSECSCluster{},
		"AWS::ECS::Service":                                     AWSECSService{},
		"AWS::ECS::TaskDefinition":                              AWSECSTaskDefinition{},
		"AWS::EFS::FileSystem":                                  AWSEFSFileSystem{},
		"AWS::EFS::MountTarget":                                 AWSEFSMountTarget{},
		"AWS::EKS::Cluster":                                     AWSEKSCluster{},
		"AWS::EMR::Cluster":                                     AWSEMRCluster{},
		"AWS::EMR::InstanceFleetConfig":                         AWSEMRInstanceFleetConfig{},
		"AWS::EMR::InstanceGroupConfig":                         AWSEMRInstanceGroupConfig{},
		"AWS::EMR::SecurityConfiguration":                       AWSEMRSecurityConfiguration{},
		"AWS::EMR::Step":                                        AWSEMRStep{},
		"AWS::ElastiCache::CacheCluster":                        AWSElastiCacheCacheCluster{},
		"AWS::ElastiCache::ParameterGroup":                      AWSElastiCacheParameterGroup{},
		"AWS::ElastiCache::ReplicationGroup":                    AWSElastiCacheReplicationGroup{},
		"AWS::ElastiCache::SecurityGroup":                       AWSElastiCacheSecurityGroup{},
		"AWS::ElastiCache::SecurityGroupIngress":                AWSElastiCacheSecurityGroupIngress{},
		"AWS::ElastiCache::SubnetGroup":                         AWSElastiCacheSubnetGroup{},
		"AWS::ElasticBeanstalk::Application":                    AWSElasticBeanstalkApplication{},
		"AWS::ElasticBeanstalk::ApplicationVersion":             AWSElasticBeanstalkApplicationVersion{},
		"AWS::ElasticBeanstalk::ConfigurationTemplate":          AWSElasticBeanstalkConfigurationTemplate{},
		"AWS::ElasticBeanstalk::Environment":                    AWSElasticBeanstalkEnvironment{},
		"AWS::ElasticLoadBalancing::LoadBalancer":               AWSElasticLoadBalancingLoadBalancer{},
		"AWS::ElasticLoadBalancingV2::Listener":                 AWSElasticLoadBalancingV2Listener{},
		"AWS::ElasticLoadBalancingV2::ListenerCertificate":      AWSElasticLoadBalancingV2ListenerCertificate{},
		"AWS::ElasticLoadBalancingV2::ListenerRule":             AWSElasticLoadBalancingV2ListenerRule{},
		"AWS::ElasticLoadBalancingV2::LoadBalancer":             AWSElasticLoadBalancingV2LoadBalancer{},
		"AWS::ElasticLoadBalancingV2::TargetGroup":              AWSElasticLoadBalancingV2TargetGroup{},
		"AWS::Elasticsearch::Domain":                            AWSElasticsearchDomain{},
		"AWS::Events::EventBusPolicy":                           AWSEventsEventBusPolicy{},
		"AWS::Events::Rule":                                     AWSEventsRule{},
		"AWS::GameLift::Alias":                                  AWSGameLiftAlias{},
		"AWS::GameLift::Build":                                  AWSGameLiftBuild{},
		"AWS::GameLift::Fleet":                                  AWSGameLiftFleet{},
		"AWS::Glue::Classifier":                                 AWSGlueClassifier{},
		"AWS::Glue::Connection":                                 AWSGlueConnection{},
		"AWS::Glue::Crawler":                                    AWSGlueCrawler{},
		"AWS::Glue::Database":                                   AWSGlueDatabase{},
		"AWS::Glue::DevEndpoint":                                AWSGlueDevEndpoint{},
		"AWS::Glue::Job":                                        AWSGlueJob{},
		"AWS::Glue::Partition":                                  AWSGluePartition{},
		"AWS::Glue::Table":                                      AWSGlueTable{},
		"AWS::Glue::Trigger":                                    AWSGlueTrigger{},
		"AWS::GuardDuty::Detector":                              AWSGuardDutyDetector{},
		"AWS::GuardDuty::Filter":                                AWSGuardDutyFilter{},
		"AWS::GuardDuty::IPSet":                                 AWSGuardDutyIPSet{},
		"AWS::GuardDuty::Master":                                AWSGuardDutyMaster{},
		"AWS::GuardDuty::Member":                                AWSGuardDutyMember{},
		"AWS::GuardDuty::ThreatIntelSet":                        AWSGuardDutyThreatIntelSet{},
		"AWS::IAM::AccessKey":                                   AWSIAMAccessKey{},
		"AWS::IAM::Group":                                       AWSIAMGroup{},
		"AWS::IAM::InstanceProfile":                             AWSIAMInstanceProfile{},
		"AWS::IAM::ManagedPolicy":                               AWSIAMManagedPolicy{},
		"AWS::IAM::Policy":                                      AWSIAMPolicy{},
		"AWS::IAM::Role":                                        AWSIAMRole{},
		"AWS::IAM::ServiceLinkedRole":                           AWSIAMServiceLinkedRole{},
		"AWS::IAM::User":                                        AWSIAMUser{},
		"AWS::IAM::UserToGroupAddition":                         AWSIAMUserToGroupAddition{},
		"AWS::Inspector::AssessmentTarget":                      AWSInspectorAssessmentTarget{},
		"AWS::Inspector::AssessmentTemplate":                    AWSInspectorAssessmentTemplate{},
		"AWS::Inspector::ResourceGroup":                         AWSInspectorResourceGroup{},
		"AWS::IoT1Click::Device":                                AWSIoT1ClickDevice{},
		"AWS::IoT1Click::Placement":                             AWSIoT1ClickPlacement{},
		"AWS::IoT1Click::Project":                               AWSIoT1ClickProject{},
		"AWS::IoT::Certificate":                                 AWSIoTCertificate{},
		"AWS::IoT::Policy":                                      AWSIoTPolicy{},
		"AWS::IoT::PolicyPrincipalAttachment":                   AWSIoTPolicyPrincipalAttachment{},
		"AWS::IoT::Thing":                                       AWSIoTThing{},
		"AWS::IoT::ThingPrincipalAttachment":                    AWSIoTThingPrincipalAttachment{},
		"AWS::IoT::TopicRule":                                   AWSIoTTopicRule{},
		"AWS::KMS::Alias":                                       AWSKMSAlias{},
		"AWS::KMS::Key":                                         AWSKMSKey{},
		"AWS::Kinesis::Stream":                                  AWSKinesisStream{},
		"AWS::KinesisAnalytics::Application":                    AWSKinesisAnalyticsApplication{},
		"AWS::KinesisAnalytics::ApplicationOutput":              AWSKinesisAnalyticsApplicationOutput{},
		"AWS::KinesisAnalytics::ApplicationReferenceDataSource": AWSKinesisAnalyticsApplicationReferenceDataSource{},
		"AWS::KinesisFirehose::DeliveryStream":                  AWSKinesisFirehoseDeliveryStream{},
		"AWS::Lambda::Alias":                                    AWSLambdaAlias{},
		"AWS::Lambda::EventSourceMapping":                       AWSLambdaEventSourceMapping{},
		"AWS::Lambda::Function":                                 AWSLambdaFunction{},
		"AWS::Lambda::Permission":                               AWSLambdaPermission{},
		"AWS::Lambda::Version":                                  AWSLambdaVersion{},
		"AWS::Logs::Destination":                                AWSLogsDestination{},
		"AWS::Logs::LogGroup":                                   AWSLogsLogGroup{},
		"AWS::Logs::LogStream":                                  AWSLogsLogStream{},
		"AWS::Logs::MetricFilter":                               AWSLogsMetricFilter{},
		"AWS::Logs::SubscriptionFilter":                         AWSLogsSubscriptionFilter{},
		"AWS::Neptune::DBCluster":                               AWSNeptuneDBCluster{},
		"AWS::Neptune::DBClusterParameterGroup":                 AWSNeptuneDBClusterParameterGroup{},
		"AWS::Neptune::DBInstance":                              AWSNeptuneDBInstance{},
		"AWS::Neptune::DBParameterGroup":                        AWSNeptuneDBParameterGroup{},
		"AWS::Neptune::DBSubnetGroup":                           AWSNeptuneDBSubnetGroup{},
		"AWS::OpsWorks::App":                                    AWSOpsWorksApp{},
		"AWS::OpsWorks::ElasticLoadBalancerAttachment":          AWSOpsWorksElasticLoadBalancerAttachment{},
		"AWS::OpsWorks::Instance":                               AWSOpsWorksInstance{},
		"AWS::OpsWorks::Layer":                                  AWSOpsWorksLayer{},
		"AWS::OpsWorks::Stack":                                  AWSOpsWorksStack{},
		"AWS::OpsWorks::UserProfile":                            AWSOpsWorksUserProfile{},
		"AWS::OpsWorks::Volume":                                 AWSOpsWorksVolume{},
		"AWS::RDS::DBCluster":                                   AWSRDSDBCluster{},
		"AWS::RDS::DBClusterParameterGroup":                     AWSRDSDBClusterParameterGroup{},
		"AWS::RDS::DBInstance":                                  AWSRDSDBInstance{},
		"AWS::RDS::DBParameterGroup":                            AWSRDSDBParameterGroup{},
		"AWS::RDS::DBSecurityGroup":                             AWSRDSDBSecurityGroup{},
		"AWS::RDS::DBSecurityGroupIngress":                      AWSRDSDBSecurityGroupIngress{},
		"AWS::RDS::DBSubnetGroup":                               AWSRDSDBSubnetGroup{},
		"AWS::RDS::EventSubscription":                           AWSRDSEventSubscription{},
		"AWS::RDS::OptionGroup":                                 AWSRDSOptionGroup{},
		"AWS::Redshift::Cluster":                                AWSRedshiftCluster{},
		"AWS::Redshift::ClusterParameterGroup":                  AWSRedshiftClusterParameterGroup{},
		"AWS::Redshift::ClusterSecurityGroup":                   AWSRedshiftClusterSecurityGroup{},
		"AWS::Redshift::ClusterSecurityGroupIngress":            AWSRedshiftClusterSecurityGroupIngress{},
		"AWS::Redshift::ClusterSubnetGroup":                     AWSRedshiftClusterSubnetGroup{},
		"AWS::Route53::HealthCheck":                             AWSRoute53HealthCheck{},
		"AWS::Route53::HostedZone":                              AWSRoute53HostedZone{},
		"AWS::Route53::RecordSet":                               AWSRoute53RecordSet{},
		"AWS::Route53::RecordSetGroup":                          AWSRoute53RecordSetGroup{},
		"AWS::S3::Bucket":                                       AWSS3Bucket{},
		"AWS::S3::BucketPolicy":                                 AWSS3BucketPolicy{},
		"AWS::SDB::Domain":                                      AWSSDBDomain{},
		"AWS::SES::ConfigurationSet":                            AWSSESConfigurationSet{},
		"AWS::SES::ConfigurationSetEventDestination":            AWSSESConfigurationSetEventDestination{},
		"AWS::SES::ReceiptFilter":                               AWSSESReceiptFilter{},
		"AWS::SES::ReceiptRule":                                 AWSSESReceiptRule{},
		"AWS::SES::ReceiptRuleSet":                              AWSSESReceiptRuleSet{},
		"AWS::SES::Template":                                    AWSSESTemplate{},
		"AWS::SNS::Subscription":                                AWSSNSSubscription{},
		"AWS::SNS::Topic":                                       AWSSNSTopic{},
		"AWS::SNS::TopicPolicy":                                 AWSSNSTopicPolicy{},
		"AWS::SQS::Queue":                                       AWSSQSQueue{},
		"AWS::SQS::QueuePolicy":                                 AWSSQSQueuePolicy{},
		"AWS::SSM::Association":                                 AWSSSMAssociation{},
		"AWS::SSM::Document":                                    AWSSSMDocument{},
		"AWS::SSM::MaintenanceWindowTask":                       AWSSSMMaintenanceWindowTask{},
		"AWS::SSM::Parameter":                                   AWSSSMParameter{},
		"AWS::SSM::PatchBaseline":                               AWSSSMPatchBaseline{},
		"AWS::SSM::ResourceDataSync":                            AWSSSMResourceDataSync{},
		"AWS::SageMaker::Endpoint":                              AWSSageMakerEndpoint{},
		"AWS::SageMaker::EndpointConfig":                        AWSSageMakerEndpointConfig{},
		"AWS::SageMaker::Model":                                 AWSSageMakerModel{},
		"AWS::SageMaker::NotebookInstance":                      AWSSageMakerNotebookInstance{},
		"AWS::SageMaker::NotebookInstanceLifecycleConfig":       AWSSageMakerNotebookInstanceLifecycleConfig{},
		"AWS::Serverless::Api":                                  AWSServerlessApi{},
		"AWS::Serverless::Function":                             AWSServerlessFunction{},
		"AWS::Serverless::SimpleTable":                          AWSServerlessSimpleTable{},
		"AWS::ServiceCatalog::AcceptedPortfolioShare":           AWSServiceCatalogAcceptedPortfolioShare{},
		"AWS::ServiceCatalog::CloudFormationProduct":            AWSServiceCatalogCloudFormationProduct{},
		"AWS::ServiceCatalog::CloudFormationProvisionedProduct": AWSServiceCatalogCloudFormationProvisionedProduct{},
		"AWS::ServiceCatalog::LaunchNotificationConstraint":     AWSServiceCatalogLaunchNotificationConstraint{},
		"AWS::ServiceCatalog::LaunchRoleConstraint":             AWSServiceCatalogLaunchRoleConstraint{},
		"AWS::ServiceCatalog::LaunchTemplateConstraint":         AWSServiceCatalogLaunchTemplateConstraint{},
		"AWS::ServiceCatalog::Portfolio":                        AWSServiceCatalogPortfolio{},
		"AWS::ServiceCatalog::PortfolioPrincipalAssociation":    AWSServiceCatalogPortfolioPrincipalAssociation{},
		"AWS::ServiceCatalog::PortfolioProductAssociation":      AWSServiceCatalogPortfolioProductAssociation{},
		"AWS::ServiceCatalog::PortfolioShare":                   AWSServiceCatalogPortfolioShare{},
		"AWS::ServiceCatalog::TagOption":                        AWSServiceCatalogTagOption{},
		"AWS::ServiceCatalog::TagOptionAssociation":             AWSServiceCatalogTagOptionAssociation{},
		"AWS::ServiceDiscovery::Instance":                       AWSServiceDiscoveryInstance{},
		"AWS::ServiceDiscovery::PrivateDnsNamespace":            AWSServiceDiscoveryPrivateDnsNamespace{},
		"AWS::ServiceDiscovery::PublicDnsNamespace":             AWSServiceDiscoveryPublicDnsNamespace{},
		"AWS::ServiceDiscovery::Service":                        AWSServiceDiscoveryService{},
		"AWS::StepFunctions::Activity":                          AWSStepFunctionsActivity{},
		"AWS::StepFunctions::StateMachine":                      AWSStepFunctionsStateMachine{},
		"AWS::WAF::ByteMatchSet":                                AWSWAFByteMatchSet{},
		"AWS::WAF::IPSet":                                       AWSWAFIPSet{},
		"AWS::WAF::Rule":                                        AWSWAFRule{},
		"AWS::WAF::SizeConstraintSet":                           AWSWAFSizeConstraintSet{},
		"AWS::WAF::SqlInjectionMatchSet":                        AWSWAFSqlInjectionMatchSet{},
		"AWS::WAF::WebACL":                                      AWSWAFWebACL{},
		"AWS::WAF::XssMatchSet":                                 AWSWAFXssMatchSet{},
		"AWS::WAFRegional::ByteMatchSet":                        AWSWAFRegionalByteMatchSet{},
		"AWS::WAFRegional::IPSet":                               AWSWAFRegionalIPSet{},
		"AWS::WAFRegional::Rule":                                AWSWAFRegionalRule{},
		"AWS::WAFRegional::SizeConstraintSet":                   AWSWAFRegionalSizeConstraintSet{},
		"AWS::WAFRegional::SqlInjectionMatchSet":                AWSWAFRegionalSqlInjectionMatchSet{},
		"AWS::WAFRegional::WebACL":                              AWSWAFRegionalWebACL{},
		"AWS::WAFRegional::WebACLAssociation":                   AWSWAFRegionalWebACLAssociation{},
		"AWS::WAFRegional::XssMatchSet":                         AWSWAFRegionalXssMatchSet{},
		"AWS::WorkSpaces::Workspace":                            AWSWorkSpacesWorkspace{},
	}
}
