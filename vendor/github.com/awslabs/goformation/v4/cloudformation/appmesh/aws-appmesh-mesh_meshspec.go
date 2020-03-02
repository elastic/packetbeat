package appmesh

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Mesh_MeshSpec AWS CloudFormation Resource (AWS::AppMesh::Mesh.MeshSpec)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-meshspec.html
type Mesh_MeshSpec struct {

	// EgressFilter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-mesh-meshspec.html#cfn-appmesh-mesh-meshspec-egressfilter
	EgressFilter *Mesh_EgressFilter `json:"EgressFilter,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Mesh_MeshSpec) AWSCloudFormationType() string {
	return "AWS::AppMesh::Mesh.MeshSpec"
}
