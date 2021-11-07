// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AcceptTransitGatewayPeeringAttachmentInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptTransitGatewayPeeringAttachmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptTransitGatewayPeeringAttachmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptTransitGatewayPeeringAttachmentInput"}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AcceptTransitGatewayPeeringAttachmentOutput struct {
	_ struct{} `type:"structure"`

	// The transit gateway peering attachment.
	TransitGatewayPeeringAttachment *TransitGatewayPeeringAttachment `locationName:"transitGatewayPeeringAttachment" type:"structure"`
}

// String returns the string representation
func (s AcceptTransitGatewayPeeringAttachmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptTransitGatewayPeeringAttachment = "AcceptTransitGatewayPeeringAttachment"

// AcceptTransitGatewayPeeringAttachmentRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Accepts a transit gateway peering attachment request. The peering attachment
// must be in the pendingAcceptance state.
//
//    // Example sending a request using AcceptTransitGatewayPeeringAttachmentRequest.
//    req := client.AcceptTransitGatewayPeeringAttachmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AcceptTransitGatewayPeeringAttachment
func (c *Client) AcceptTransitGatewayPeeringAttachmentRequest(input *AcceptTransitGatewayPeeringAttachmentInput) AcceptTransitGatewayPeeringAttachmentRequest {
	op := &aws.Operation{
		Name:       opAcceptTransitGatewayPeeringAttachment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptTransitGatewayPeeringAttachmentInput{}
	}

	req := c.newRequest(op, input, &AcceptTransitGatewayPeeringAttachmentOutput{})

	return AcceptTransitGatewayPeeringAttachmentRequest{Request: req, Input: input, Copy: c.AcceptTransitGatewayPeeringAttachmentRequest}
}

// AcceptTransitGatewayPeeringAttachmentRequest is the request type for the
// AcceptTransitGatewayPeeringAttachment API operation.
type AcceptTransitGatewayPeeringAttachmentRequest struct {
	*aws.Request
	Input *AcceptTransitGatewayPeeringAttachmentInput
	Copy  func(*AcceptTransitGatewayPeeringAttachmentInput) AcceptTransitGatewayPeeringAttachmentRequest
}

// Send marshals and sends the AcceptTransitGatewayPeeringAttachment API request.
func (r AcceptTransitGatewayPeeringAttachmentRequest) Send(ctx context.Context) (*AcceptTransitGatewayPeeringAttachmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptTransitGatewayPeeringAttachmentResponse{
		AcceptTransitGatewayPeeringAttachmentOutput: r.Request.Data.(*AcceptTransitGatewayPeeringAttachmentOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptTransitGatewayPeeringAttachmentResponse is the response type for the
// AcceptTransitGatewayPeeringAttachment API operation.
type AcceptTransitGatewayPeeringAttachmentResponse struct {
	*AcceptTransitGatewayPeeringAttachmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptTransitGatewayPeeringAttachment request.
func (r *AcceptTransitGatewayPeeringAttachmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
