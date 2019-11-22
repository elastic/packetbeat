// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBClusterFromS3Message
type RestoreDBClusterFromS3Input struct {
	_ struct{} `type:"structure"`

	// A list of Availability Zones (AZs) where instances in the restored DB cluster
	// can be created.
	AvailabilityZones []string `locationNameList:"AvailabilityZone" type:"list"`

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0.
	//
	// Default: 0
	//
	// Constraints:
	//
	//    * If specified, this value must be set to a number from 0 to 259,200 (72
	//    hours).
	BacktrackWindow *int64 `type:"long"`

	// The number of days for which automated backups of the restored DB cluster
	// are retained. You must specify a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 1 to 35
	BackupRetentionPeriod *int64 `type:"integer"`

	// A value that indicates that the restored DB cluster should be associated
	// with the specified CharacterSet.
	CharacterSetName *string `type:"string"`

	// True to copy all tags from the restored DB cluster to snapshots of the restored
	// DB cluster, and otherwise false. The default is false.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The name of the DB cluster to create from the source data in the Amazon S3
	// bucket. This parameter is isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with the restored
	// DB cluster. If this argument is omitted, default.aurora5.6 is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string `type:"string"`

	// A DB subnet group to associate with the restored DB cluster.
	//
	// Constraints: If supplied, must match the name of an existing DBSubnetGroup.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// The database name for the restored DB cluster.
	DatabaseName *string `type:"string"`

	// Indicates if the DB cluster should have deletion protection enabled. The
	// database can't be deleted when this value is set to true. The default is
	// false.
	DeletionProtection *bool `type:"boolean"`

	// The list of logs that the restored DB cluster is to export to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false.
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The name of the database engine to be used for the restored DB cluster.
	//
	// Valid Values: aurora, aurora-postgresql
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The version number of the database engine to use.
	//
	// Aurora MySQL
	//
	// Example: 5.6.10a
	//
	// Aurora PostgreSQL
	//
	// Example: 9.6.3
	EngineVersion *string `type:"string"`

	// The AWS KMS key identifier for an encrypted DB cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a DB cluster with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB cluster, then you can use
	// the KMS key alias instead of the ARN for the KM encryption key.
	//
	// If the StorageEncrypted parameter is true, and you do not specify a value
	// for the KmsKeyId parameter, then Amazon RDS will use your default encryption
	// key. AWS KMS creates the default encryption key for your AWS account. Your
	// AWS account has a different default encryption key for each AWS Region.
	KmsKeyId *string `type:"string"`

	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@".
	//
	// Constraints: Must contain from 8 to 41 characters.
	//
	// MasterUserPassword is a required field
	MasterUserPassword *string `type:"string" required:"true"`

	// The name of the master user for the restored DB cluster.
	//
	// Constraints:
	//
	//    * Must be 1 to 16 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Can't be a reserved word for the chosen database engine.
	//
	// MasterUsername is a required field
	MasterUsername *string `type:"string" required:"true"`

	// A value that indicates that the restored DB cluster should be associated
	// with the specified option group.
	//
	// Permanent options can't be removed from an option group. An option group
	// can't be removed from a DB cluster once it is associated with a DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the instances in the restored DB cluster accept
	// connections.
	//
	// Default: 3306
	Port *int64 `type:"integer"`

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region. To see the time blocks available, see Adjusting
	// the Preferred Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `type:"string"`

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week. To see
	// the time blocks available, see Adjusting the Preferred Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// The name of the Amazon S3 bucket that contains the data used to create the
	// Amazon Aurora DB cluster.
	//
	// S3BucketName is a required field
	S3BucketName *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) role that authorizes Amazon RDS to access the Amazon S3 bucket on your
	// behalf.
	//
	// S3IngestionRoleArn is a required field
	S3IngestionRoleArn *string `type:"string" required:"true"`

	// The prefix for all of the file names that contain the data used to create
	// the Amazon Aurora DB cluster. If you do not specify a SourceS3Prefix value,
	// then the Amazon Aurora DB cluster is created by using all of the files in
	// the Amazon S3 bucket.
	S3Prefix *string `type:"string"`

	// The identifier for the database engine that was backed up to create the files
	// stored in the Amazon S3 bucket.
	//
	// Valid values: mysql
	//
	// SourceEngine is a required field
	SourceEngine *string `type:"string" required:"true"`

	// The version of the database that the backup files were created from.
	//
	// MySQL version 5.5 and 5.6 are supported.
	//
	// Example: 5.6.22
	//
	// SourceEngineVersion is a required field
	SourceEngineVersion *string `type:"string" required:"true"`

	// Specifies whether the restored DB cluster is encrypted.
	StorageEncrypted *bool `type:"boolean"`

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A list of EC2 VPC security groups to associate with the restored DB cluster.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBClusterFromS3Input) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBClusterFromS3Input) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBClusterFromS3Input"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if s.MasterUserPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("MasterUserPassword"))
	}

	if s.MasterUsername == nil {
		invalidParams.Add(aws.NewErrParamRequired("MasterUsername"))
	}

	if s.S3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3BucketName"))
	}

	if s.S3IngestionRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3IngestionRoleArn"))
	}

	if s.SourceEngine == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEngine"))
	}

	if s.SourceEngineVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEngineVersion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBClusterFromS3Result
type RestoreDBClusterFromS3Output struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Aurora DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters, StopDBCluster,
	// and StartDBCluster actions.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s RestoreDBClusterFromS3Output) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBClusterFromS3 = "RestoreDBClusterFromS3"

// RestoreDBClusterFromS3Request returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates an Amazon Aurora DB cluster from data stored in an Amazon S3 bucket.
// Amazon RDS must be authorized to access the Amazon S3 bucket and the data
// must be created using the Percona XtraBackup utility as described in Migrating
// Data to an Amazon Aurora MySQL DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Migrating.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using RestoreDBClusterFromS3Request.
//    req := client.RestoreDBClusterFromS3Request(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBClusterFromS3
func (c *Client) RestoreDBClusterFromS3Request(input *RestoreDBClusterFromS3Input) RestoreDBClusterFromS3Request {
	op := &aws.Operation{
		Name:       opRestoreDBClusterFromS3,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBClusterFromS3Input{}
	}

	req := c.newRequest(op, input, &RestoreDBClusterFromS3Output{})
	return RestoreDBClusterFromS3Request{Request: req, Input: input, Copy: c.RestoreDBClusterFromS3Request}
}

// RestoreDBClusterFromS3Request is the request type for the
// RestoreDBClusterFromS3 API operation.
type RestoreDBClusterFromS3Request struct {
	*aws.Request
	Input *RestoreDBClusterFromS3Input
	Copy  func(*RestoreDBClusterFromS3Input) RestoreDBClusterFromS3Request
}

// Send marshals and sends the RestoreDBClusterFromS3 API request.
func (r RestoreDBClusterFromS3Request) Send(ctx context.Context) (*RestoreDBClusterFromS3Response, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBClusterFromS3Response{
		RestoreDBClusterFromS3Output: r.Request.Data.(*RestoreDBClusterFromS3Output),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBClusterFromS3Response is the response type for the
// RestoreDBClusterFromS3 API operation.
type RestoreDBClusterFromS3Response struct {
	*RestoreDBClusterFromS3Output

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBClusterFromS3 request.
func (r *RestoreDBClusterFromS3Response) SDKResponseMetdata() *aws.Response {
	return r.response
}
