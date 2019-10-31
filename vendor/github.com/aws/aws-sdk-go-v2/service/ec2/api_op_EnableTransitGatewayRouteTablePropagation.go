// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableTransitGatewayRouteTablePropagationRequest
type EnableTransitGatewayRouteTablePropagationInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`

	// The ID of the propagation route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableTransitGatewayRouteTablePropagationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableTransitGatewayRouteTablePropagationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableTransitGatewayRouteTablePropagationInput"}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableTransitGatewayRouteTablePropagationResult
type EnableTransitGatewayRouteTablePropagationOutput struct {
	_ struct{} `type:"structure"`

	// Information about route propagation.
	Propagation *TransitGatewayPropagation `locationName:"propagation" type:"structure"`
}

// String returns the string representation
func (s EnableTransitGatewayRouteTablePropagationOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableTransitGatewayRouteTablePropagation = "EnableTransitGatewayRouteTablePropagation"

// EnableTransitGatewayRouteTablePropagationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables the specified attachment to propagate routes to the specified propagation
// route table.
//
//    // Example sending a request using EnableTransitGatewayRouteTablePropagationRequest.
//    req := client.EnableTransitGatewayRouteTablePropagationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableTransitGatewayRouteTablePropagation
func (c *Client) EnableTransitGatewayRouteTablePropagationRequest(input *EnableTransitGatewayRouteTablePropagationInput) EnableTransitGatewayRouteTablePropagationRequest {
	op := &aws.Operation{
		Name:       opEnableTransitGatewayRouteTablePropagation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableTransitGatewayRouteTablePropagationInput{}
	}

	req := c.newRequest(op, input, &EnableTransitGatewayRouteTablePropagationOutput{})
	return EnableTransitGatewayRouteTablePropagationRequest{Request: req, Input: input, Copy: c.EnableTransitGatewayRouteTablePropagationRequest}
}

// EnableTransitGatewayRouteTablePropagationRequest is the request type for the
// EnableTransitGatewayRouteTablePropagation API operation.
type EnableTransitGatewayRouteTablePropagationRequest struct {
	*aws.Request
	Input *EnableTransitGatewayRouteTablePropagationInput
	Copy  func(*EnableTransitGatewayRouteTablePropagationInput) EnableTransitGatewayRouteTablePropagationRequest
}

// Send marshals and sends the EnableTransitGatewayRouteTablePropagation API request.
func (r EnableTransitGatewayRouteTablePropagationRequest) Send(ctx context.Context) (*EnableTransitGatewayRouteTablePropagationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableTransitGatewayRouteTablePropagationResponse{
		EnableTransitGatewayRouteTablePropagationOutput: r.Request.Data.(*EnableTransitGatewayRouteTablePropagationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableTransitGatewayRouteTablePropagationResponse is the response type for the
// EnableTransitGatewayRouteTablePropagation API operation.
type EnableTransitGatewayRouteTablePropagationResponse struct {
	*EnableTransitGatewayRouteTablePropagationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableTransitGatewayRouteTablePropagation request.
func (r *EnableTransitGatewayRouteTablePropagationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
