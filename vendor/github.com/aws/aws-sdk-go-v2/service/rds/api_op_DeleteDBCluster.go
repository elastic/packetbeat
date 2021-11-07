// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The DB cluster identifier for the DB cluster to be deleted. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must match an existing DBClusterIdentifier.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The DB cluster snapshot identifier of the new DB cluster snapshot created
	// when SkipFinalSnapshot is disabled.
	//
	// Specifying this parameter and also skipping the creation of a final DB cluster
	// snapshot with the SkipFinalShapshot parameter results in an error.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	FinalDBSnapshotIdentifier *string `type:"string"`

	// A value that indicates whether to skip the creation of a final DB cluster
	// snapshot before the DB cluster is deleted. If skip is specified, no DB cluster
	// snapshot is created. If skip isn't specified, a DB cluster snapshot is created
	// before the DB cluster is deleted. By default, skip isn't specified, and the
	// DB cluster snapshot is created. By default, this parameter is disabled.
	//
	// You must specify a FinalDBSnapshotIdentifier parameter if SkipFinalSnapshot
	// is disabled.
	SkipFinalSnapshot *bool `type:"boolean"`
}

// String returns the string representation
func (s DeleteDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Aurora DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters, StopDBCluster,
	// and StartDBCluster actions.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s DeleteDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBCluster = "DeleteDBCluster"

// DeleteDBClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// The DeleteDBCluster action deletes a previously provisioned DB cluster. When
// you delete a DB cluster, all automated backups for that DB cluster are deleted
// and can't be recovered. Manual DB cluster snapshots of the specified DB cluster
// are not deleted.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using DeleteDBClusterRequest.
//    req := client.DeleteDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBCluster
func (c *Client) DeleteDBClusterRequest(input *DeleteDBClusterInput) DeleteDBClusterRequest {
	op := &aws.Operation{
		Name:       opDeleteDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBClusterInput{}
	}

	req := c.newRequest(op, input, &DeleteDBClusterOutput{})

	return DeleteDBClusterRequest{Request: req, Input: input, Copy: c.DeleteDBClusterRequest}
}

// DeleteDBClusterRequest is the request type for the
// DeleteDBCluster API operation.
type DeleteDBClusterRequest struct {
	*aws.Request
	Input *DeleteDBClusterInput
	Copy  func(*DeleteDBClusterInput) DeleteDBClusterRequest
}

// Send marshals and sends the DeleteDBCluster API request.
func (r DeleteDBClusterRequest) Send(ctx context.Context) (*DeleteDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBClusterResponse{
		DeleteDBClusterOutput: r.Request.Data.(*DeleteDBClusterOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBClusterResponse is the response type for the
// DeleteDBCluster API operation.
type DeleteDBClusterResponse struct {
	*DeleteDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBCluster request.
func (r *DeleteDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
