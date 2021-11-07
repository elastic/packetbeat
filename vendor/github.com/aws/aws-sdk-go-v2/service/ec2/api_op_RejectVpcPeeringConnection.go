// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RejectVpcPeeringConnectionInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC peering connection.
	//
	// VpcPeeringConnectionId is a required field
	VpcPeeringConnectionId *string `locationName:"vpcPeeringConnectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s RejectVpcPeeringConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectVpcPeeringConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectVpcPeeringConnectionInput"}

	if s.VpcPeeringConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcPeeringConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RejectVpcPeeringConnectionOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s RejectVpcPeeringConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opRejectVpcPeeringConnection = "RejectVpcPeeringConnection"

// RejectVpcPeeringConnectionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Rejects a VPC peering connection request. The VPC peering connection must
// be in the pending-acceptance state. Use the DescribeVpcPeeringConnections
// request to view your outstanding VPC peering connection requests. To delete
// an active VPC peering connection, or to delete a VPC peering connection request
// that you initiated, use DeleteVpcPeeringConnection.
//
//    // Example sending a request using RejectVpcPeeringConnectionRequest.
//    req := client.RejectVpcPeeringConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RejectVpcPeeringConnection
func (c *Client) RejectVpcPeeringConnectionRequest(input *RejectVpcPeeringConnectionInput) RejectVpcPeeringConnectionRequest {
	op := &aws.Operation{
		Name:       opRejectVpcPeeringConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RejectVpcPeeringConnectionInput{}
	}

	req := c.newRequest(op, input, &RejectVpcPeeringConnectionOutput{})

	return RejectVpcPeeringConnectionRequest{Request: req, Input: input, Copy: c.RejectVpcPeeringConnectionRequest}
}

// RejectVpcPeeringConnectionRequest is the request type for the
// RejectVpcPeeringConnection API operation.
type RejectVpcPeeringConnectionRequest struct {
	*aws.Request
	Input *RejectVpcPeeringConnectionInput
	Copy  func(*RejectVpcPeeringConnectionInput) RejectVpcPeeringConnectionRequest
}

// Send marshals and sends the RejectVpcPeeringConnection API request.
func (r RejectVpcPeeringConnectionRequest) Send(ctx context.Context) (*RejectVpcPeeringConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectVpcPeeringConnectionResponse{
		RejectVpcPeeringConnectionOutput: r.Request.Data.(*RejectVpcPeeringConnectionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectVpcPeeringConnectionResponse is the response type for the
// RejectVpcPeeringConnection API operation.
type RejectVpcPeeringConnectionResponse struct {
	*RejectVpcPeeringConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectVpcPeeringConnection request.
func (r *RejectVpcPeeringConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
