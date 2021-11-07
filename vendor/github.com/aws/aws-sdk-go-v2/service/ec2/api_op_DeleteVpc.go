// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DeleteVpcInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpcInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVpcInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteVpcOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVpcOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteVpc = "DeleteVpc"

// DeleteVpcRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified VPC. You must detach or delete all gateways and resources
// that are associated with the VPC before you can delete it. For example, you
// must terminate all instances running in the VPC, delete all security groups
// associated with the VPC (except the default one), delete all route tables
// associated with the VPC (except the default one), and so on.
//
//    // Example sending a request using DeleteVpcRequest.
//    req := client.DeleteVpcRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpc
func (c *Client) DeleteVpcRequest(input *DeleteVpcInput) DeleteVpcRequest {
	op := &aws.Operation{
		Name:       opDeleteVpc,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcInput{}
	}

	req := c.newRequest(op, input, &DeleteVpcOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteVpcRequest{Request: req, Input: input, Copy: c.DeleteVpcRequest}
}

// DeleteVpcRequest is the request type for the
// DeleteVpc API operation.
type DeleteVpcRequest struct {
	*aws.Request
	Input *DeleteVpcInput
	Copy  func(*DeleteVpcInput) DeleteVpcRequest
}

// Send marshals and sends the DeleteVpc API request.
func (r DeleteVpcRequest) Send(ctx context.Context) (*DeleteVpcResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpcResponse{
		DeleteVpcOutput: r.Request.Data.(*DeleteVpcOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpcResponse is the response type for the
// DeleteVpc API operation.
type DeleteVpcResponse struct {
	*DeleteVpcOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpc request.
func (r *DeleteVpcResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
