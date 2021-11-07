// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDBInstanceReadReplicaInput struct {
	_ struct{} `type:"structure"`

	// A value that indicates whether minor engine upgrades are applied automatically
	// to the read replica during the maintenance window.
	//
	// Default: Inherits from the source DB instance
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The Availability Zone (AZ) where the read replica will be created.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS
	// Region.
	//
	// Example: us-east-1d
	AvailabilityZone *string `type:"string"`

	// A value that indicates whether to copy all tags from the read replica to
	// snapshots of the read replica. By default, tags are not copied.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The compute and memory capacity of the read replica, for example, db.m4.large.
	// Not all DB instance classes are available in all AWS Regions, or for all
	// database engines. For the full list of DB instance classes, and availability
	// for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide.
	//
	// Default: Inherits from the source DB instance.
	DBInstanceClass *string `type:"string"`

	// The DB instance identifier of the read replica. This identifier is the unique
	// key that identifies a DB instance. This parameter is stored as a lowercase
	// string.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The name of the DB parameter group to associate with this DB instance.
	//
	// If you do not specify a value for DBParameterGroupName, then Amazon RDS uses
	// the DBParameterGroup of source DB instance for a same region read replica,
	// or the default DBParameterGroup for the specified DB engine for a cross region
	// read replica.
	//
	// Currently, specifying a parameter group for this operation is only supported
	// for Oracle DB instances.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	DBParameterGroupName *string `type:"string"`

	// Specifies a DB subnet group for the DB instance. The new DB instance is created
	// in the VPC associated with the DB subnet group. If no DB subnet group is
	// specified, then the new DB instance isn't created in a VPC.
	//
	// Constraints:
	//
	//    * Can only be specified if the source DB instance identifier specifies
	//    a DB instance in another AWS Region.
	//
	//    * If supplied, must match the name of an existing DBSubnetGroup.
	//
	//    * The specified DB subnet group must be in the same AWS Region in which
	//    the operation is running.
	//
	//    * All read replicas in one AWS Region that are created from the same source
	//    DB instance must either:> Specify DB subnet groups from the same VPC.
	//    All these read replicas are created in the same VPC. Not specify a DB
	//    subnet group. All these read replicas are created outside of any VPC.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see Deleting a DB
	// Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool `type:"boolean"`

	// DestinationRegion is used for presigning the request to a given region.
	DestinationRegion *string `type:"string"`

	// The Active Directory directory ID to create the DB instance in.
	//
	// For Oracle DB instances, Amazon RDS can use Kerberos authentication to authenticate
	// users that connect to the DB instance. For more information, see Using Kerberos
	// Authentication with Amazon RDS for Oracle (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos.html)
	// in the Amazon RDS User Guide.
	//
	// For Microsoft SQL Server DB instances, Amazon RDS can use Windows Authentication
	// to authenticate users that connect to the DB instance. For more information,
	// see Using Windows Authentication with an Amazon RDS DB Instance Running Microsoft
	// SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_SQLServerWinAuth.html)
	// in the Amazon RDS User Guide.
	Domain *string `type:"string"`

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string `type:"string"`

	// The list of logs that the new DB instance is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For information about the supported DB engines, see CreateDBInstance.
	//
	// For more information about IAM database authentication, see IAM Database
	// Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide.
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// A value that indicates whether to enable Performance Insights for the read
	// replica.
	//
	// For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon RDS User Guide.
	EnablePerformanceInsights *bool `type:"boolean"`

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance.
	Iops *int64 `type:"integer"`

	// The AWS KMS key ID for an encrypted read replica. The KMS key ID is the Amazon
	// Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS
	// encryption key.
	//
	// If you create an encrypted read replica in the same AWS Region as the source
	// DB instance, then you do not have to specify a value for this parameter.
	// The read replica is encrypted with the same KMS key as the source DB instance.
	//
	// If you create an encrypted read replica in a different AWS Region, then you
	// must specify a KMS key for the destination AWS Region. KMS encryption keys
	// are specific to the AWS Region that they are created in, and you can't use
	// encryption keys from one AWS Region in another AWS Region.
	//
	// You can't create an encrypted read replica from an unencrypted DB instance.
	KmsKeyId *string `type:"string"`

	// The interval, in seconds, between points when Enhanced Monitoring metrics
	// are collected for the read replica. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0.
	//
	// If MonitoringRoleArn is specified, then you must also set MonitoringInterval
	// to a value other than 0.
	//
	// Valid Values: 0, 1, 5, 10, 15, 30, 60
	MonitoringInterval *int64 `type:"integer"`

	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics
	// to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess.
	// For information on creating a monitoring role, go to To create an IAM role
	// for Amazon RDS Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html#USER_Monitoring.OS.IAMRole)
	// in the Amazon RDS User Guide.
	//
	// If MonitoringInterval is set to a value other than 0, then you must supply
	// a MonitoringRoleArn value.
	MonitoringRoleArn *string `type:"string"`

	// A value that indicates whether the read replica is in a Multi-AZ deployment.
	//
	// You can create a read replica as a Multi-AZ DB instance. RDS creates a standby
	// of your replica in another Availability Zone for failover support for the
	// replica. Creating your read replica as a Multi-AZ DB instance is independent
	// of whether the source database is a Multi-AZ DB instance.
	MultiAZ *bool `type:"boolean"`

	// The option group the DB instance is associated with. If omitted, the option
	// group associated with the source instance is used.
	//
	// For SQL Server, you must use the option group associated with the source
	// instance.
	OptionGroupName *string `type:"string"`

	// The AWS KMS key identifier for encryption of Performance Insights data. The
	// KMS key ID is the Amazon Resource Name (ARN), KMS key identifier, or the
	// KMS key alias for the KMS encryption key.
	//
	// If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon
	// RDS uses your default encryption key. AWS KMS creates the default encryption
	// key for your AWS account. Your AWS account has a different default encryption
	// key for each AWS Region.
	PerformanceInsightsKMSKeyId *string `type:"string"`

	// The amount of time, in days, to retain Performance Insights data. Valid values
	// are 7 or 731 (2 years).
	PerformanceInsightsRetentionPeriod *int64 `type:"integer"`

	// The port number that the DB instance uses for connections.
	//
	// Default: Inherits from the source DB instance
	//
	// Valid Values: 1150-65535
	Port *int64 `type:"integer"`

	// The URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplica
	// API action in the source AWS Region that contains the source DB instance.
	//
	// You must specify this parameter when you create an encrypted read replica
	// from another AWS Region by using the Amazon RDS API. Don't specify PreSignedUrl
	// when you are creating an encrypted read replica in the same AWS Region.
	//
	// The presigned URL must be a valid request for the CreateDBInstanceReadReplica
	// API action that can be executed in the source AWS Region that contains the
	// encrypted source DB instance. The presigned URL request must contain the
	// following parameter values:
	//
	//    * DestinationRegion - The AWS Region that the encrypted read replica is
	//    created in. This AWS Region is the same one where the CreateDBInstanceReadReplica
	//    action is called that contains this presigned URL. For example, if you
	//    create an encrypted DB instance in the us-west-1 AWS Region, from a source
	//    DB instance in the us-east-2 AWS Region, then you call the CreateDBInstanceReadReplica
	//    action in the us-east-1 AWS Region and provide a presigned URL that contains
	//    a call to the CreateDBInstanceReadReplica action in the us-west-2 AWS
	//    Region. For this example, the DestinationRegion in the presigned URL must
	//    be set to the us-east-1 AWS Region.
	//
	//    * KmsKeyId - The AWS KMS key identifier for the key to use to encrypt
	//    the read replica in the destination AWS Region. This is the same identifier
	//    for both the CreateDBInstanceReadReplica action that is called in the
	//    destination AWS Region, and the action contained in the presigned URL.
	//
	//    * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted
	//    DB instance to be replicated. This identifier must be in the Amazon Resource
	//    Name (ARN) format for the source AWS Region. For example, if you are creating
	//    an encrypted read replica from a DB instance in the us-west-2 AWS Region,
	//    then your SourceDBInstanceIdentifier looks like the following example:
	//    arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115.
	//
	// To learn how to generate a Signature Version 4 signed request, see Authenticating
	// Requests: Using Query Parameters (AWS Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
	//
	// If you are using an AWS SDK tool or the AWS CLI, you can specify SourceRegion
	// (or --source-region for the AWS CLI) instead of specifying PreSignedUrl manually.
	// Specifying SourceRegion autogenerates a presigned URL that is a valid request
	// for the operation that can be executed in the source AWS Region.
	//
	// SourceRegion isn't supported for SQL Server, because SQL Server on Amazon
	// RDS doesn't support cross-region read replicas.
	PreSignedUrl *string `type:"string"`

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []ProcessorFeature `locationNameList:"ProcessorFeature" type:"list"`

	// A value that indicates whether the DB instance is publicly accessible.
	//
	// When the DB instance is publicly accessible, its DNS endpoint resolves to
	// the private IP address from within the DB instance's VPC, and to the public
	// IP address from outside of the DB instance's VPC. Access to the DB instance
	// is ultimately controlled by the security group it uses, and that public access
	// is not permitted if the security group assigned to the DB instance doesn't
	// permit it.
	//
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address.
	//
	// For more information, see CreateDBInstance.
	PubliclyAccessible *bool `type:"boolean"`

	// The identifier of the DB instance that will act as the source for the read
	// replica. Each DB instance can have up to five read replicas.
	//
	// Constraints:
	//
	//    * Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL,
	//    or SQL Server DB instance.
	//
	//    * Can specify a DB instance that is a MySQL read replica only if the source
	//    is running MySQL 5.6 or later.
	//
	//    * For the limitations of Oracle read replicas, see Read Replica Limitations
	//    with Oracle (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html)
	//    in the Amazon RDS User Guide.
	//
	//    * For the limitations of SQL Server read replicas, see Read Replica Limitations
	//    with Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.Limitations.html)
	//    in the Amazon RDS User Guide.
	//
	//    * Can specify a PostgreSQL DB instance only if the source is running PostgreSQL
	//    9.3.5 or later (9.4.7 and higher for cross-region replication).
	//
	//    * The specified DB instance must have automatic backups enabled, that
	//    is, its backup retention period must be greater than 0.
	//
	//    * If the source DB instance is in the same AWS Region as the read replica,
	//    specify a valid DB instance identifier.
	//
	//    * If the source DB instance is in a different AWS Region from the read
	//    replica, specify a valid DB instance ARN. For more information, see Constructing
	//    an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	//    in the Amazon RDS User Guide. This doesn't apply to SQL Server, which
	//    doesn't support cross-region replicas.
	//
	// SourceDBInstanceIdentifier is a required field
	SourceDBInstanceIdentifier *string `type:"string" required:"true"`

	// SourceRegion is the source region where the resource exists. This is not
	// sent over the wire and is only used for presigning. This value should always
	// have the same region as the source ARN.
	SourceRegion *string `type:"string" ignore:"true"`

	// Specifies the storage type to be associated with the read replica.
	//
	// Valid values: standard | gp2 | io1
	//
	// If you specify io1, you must also include a value for the Iops parameter.
	//
	// Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string `type:"string"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A value that indicates whether the DB instance class of the DB instance uses
	// its default processor features.
	UseDefaultProcessorFeatures *bool `type:"boolean"`

	// A list of EC2 VPC security groups to associate with the read replica.
	//
	// Default: The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s CreateDBInstanceReadReplicaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBInstanceReadReplicaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBInstanceReadReplicaInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if s.SourceDBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDBInstanceReadReplicaOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s CreateDBInstanceReadReplicaOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBInstanceReadReplica = "CreateDBInstanceReadReplica"

// CreateDBInstanceReadReplicaRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new DB instance that acts as a read replica for an existing source
// DB instance. You can create a read replica for a DB instance running MySQL,
// MariaDB, Oracle, PostgreSQL, or SQL Server. For more information, see Working
// with Read Replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)
// in the Amazon RDS User Guide.
//
// Amazon Aurora doesn't support this action. Call the CreateDBInstance action
// to create a DB instance for an Aurora DB cluster.
//
// All read replica DB instances are created with backups disabled. All other
// DB instance attributes (including DB security groups and DB parameter groups)
// are inherited from the source DB instance, except as specified.
//
// Your source DB instance must have backup retention enabled.
//
//    // Example sending a request using CreateDBInstanceReadReplicaRequest.
//    req := client.CreateDBInstanceReadReplicaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateDBInstanceReadReplica
func (c *Client) CreateDBInstanceReadReplicaRequest(input *CreateDBInstanceReadReplicaInput) CreateDBInstanceReadReplicaRequest {
	op := &aws.Operation{
		Name:       opCreateDBInstanceReadReplica,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBInstanceReadReplicaInput{}
	}

	req := c.newRequest(op, input, &CreateDBInstanceReadReplicaOutput{})

	return CreateDBInstanceReadReplicaRequest{Request: req, Input: input, Copy: c.CreateDBInstanceReadReplicaRequest}
}

// CreateDBInstanceReadReplicaRequest is the request type for the
// CreateDBInstanceReadReplica API operation.
type CreateDBInstanceReadReplicaRequest struct {
	*aws.Request
	Input *CreateDBInstanceReadReplicaInput
	Copy  func(*CreateDBInstanceReadReplicaInput) CreateDBInstanceReadReplicaRequest
}

// Send marshals and sends the CreateDBInstanceReadReplica API request.
func (r CreateDBInstanceReadReplicaRequest) Send(ctx context.Context) (*CreateDBInstanceReadReplicaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBInstanceReadReplicaResponse{
		CreateDBInstanceReadReplicaOutput: r.Request.Data.(*CreateDBInstanceReadReplicaOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBInstanceReadReplicaResponse is the response type for the
// CreateDBInstanceReadReplica API operation.
type CreateDBInstanceReadReplicaResponse struct {
	*CreateDBInstanceReadReplicaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBInstanceReadReplica request.
func (r *CreateDBInstanceReadReplicaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
