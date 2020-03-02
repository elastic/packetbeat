package kinesisanalyticsv2

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Application_KinesisStreamsInput AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.KinesisStreamsInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-kinesisstreamsinput.html
type Application_KinesisStreamsInput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-kinesisstreamsinput.html#cfn-kinesisanalyticsv2-application-kinesisstreamsinput-resourcearn
	ResourceARN string `json:"ResourceARN,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Application_KinesisStreamsInput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.KinesisStreamsInput"
}
