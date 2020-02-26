package s3

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Bucket_SseKmsEncryptedObjects AWS CloudFormation Resource (AWS::S3::Bucket.SseKmsEncryptedObjects)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html
type Bucket_SseKmsEncryptedObjects struct {

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html#cfn-s3-bucket-ssekmsencryptedobjects-status
	Status string `json:"Status,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Bucket_SseKmsEncryptedObjects) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.SseKmsEncryptedObjects"
}
