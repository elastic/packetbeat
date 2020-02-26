package kinesisanalyticsv2

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Application_RecordFormat AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.RecordFormat)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-recordformat.html
type Application_RecordFormat struct {

	// MappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-recordformat.html#cfn-kinesisanalyticsv2-application-recordformat-mappingparameters
	MappingParameters *Application_MappingParameters `json:"MappingParameters,omitempty"`

	// RecordFormatType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-recordformat.html#cfn-kinesisanalyticsv2-application-recordformat-recordformattype
	RecordFormatType string `json:"RecordFormatType,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Application_RecordFormat) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.RecordFormat"
}
