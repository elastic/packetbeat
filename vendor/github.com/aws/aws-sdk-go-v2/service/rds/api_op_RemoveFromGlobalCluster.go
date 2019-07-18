// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RemoveFromGlobalClusterMessage
type RemoveFromGlobalClusterInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) identifying the cluster that was detached
	// from the Aurora global database cluster.
	DbClusterIdentifier *string `type:"string"`

	// The cluster identifier to detach from the Aurora global database cluster.
	GlobalClusterIdentifier *string `type:"string"`
}

// String returns the string representation
func (s RemoveFromGlobalClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RemoveFromGlobalClusterResult
type RemoveFromGlobalClusterOutput struct {
	_ struct{} `type:"structure"`

	// A data type representing an Aurora global database.
	GlobalCluster *GlobalCluster `type:"structure"`
}

// String returns the string representation
func (s RemoveFromGlobalClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveFromGlobalCluster = "RemoveFromGlobalCluster"

// RemoveFromGlobalClusterRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Detaches an Aurora secondary cluster from an Aurora global database cluster.
// The cluster becomes a standalone cluster with read-write capability instead
// of being read-only and receiving data from a primary cluster in a different
// region.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using RemoveFromGlobalClusterRequest.
//    req := client.RemoveFromGlobalClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RemoveFromGlobalCluster
func (c *Client) RemoveFromGlobalClusterRequest(input *RemoveFromGlobalClusterInput) RemoveFromGlobalClusterRequest {
	op := &aws.Operation{
		Name:       opRemoveFromGlobalCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveFromGlobalClusterInput{}
	}

	req := c.newRequest(op, input, &RemoveFromGlobalClusterOutput{})
	return RemoveFromGlobalClusterRequest{Request: req, Input: input, Copy: c.RemoveFromGlobalClusterRequest}
}

// RemoveFromGlobalClusterRequest is the request type for the
// RemoveFromGlobalCluster API operation.
type RemoveFromGlobalClusterRequest struct {
	*aws.Request
	Input *RemoveFromGlobalClusterInput
	Copy  func(*RemoveFromGlobalClusterInput) RemoveFromGlobalClusterRequest
}

// Send marshals and sends the RemoveFromGlobalCluster API request.
func (r RemoveFromGlobalClusterRequest) Send(ctx context.Context) (*RemoveFromGlobalClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveFromGlobalClusterResponse{
		RemoveFromGlobalClusterOutput: r.Request.Data.(*RemoveFromGlobalClusterOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveFromGlobalClusterResponse is the response type for the
// RemoveFromGlobalCluster API operation.
type RemoveFromGlobalClusterResponse struct {
	*RemoveFromGlobalClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveFromGlobalCluster request.
func (r *RemoveFromGlobalClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
