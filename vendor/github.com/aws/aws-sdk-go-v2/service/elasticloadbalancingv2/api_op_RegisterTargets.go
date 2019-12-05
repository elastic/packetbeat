// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/RegisterTargetsInput
type RegisterTargetsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the target group.
	//
	// TargetGroupArn is a required field
	TargetGroupArn *string `type:"string" required:"true"`

	// The targets.
	//
	// To register a target by instance ID, specify the instance ID. To register
	// a target by IP address, specify the IP address. To register a Lambda function,
	// specify the ARN of the Lambda function.
	//
	// Targets is a required field
	Targets []TargetDescription `type:"list" required:"true"`
}

// String returns the string representation
func (s RegisterTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterTargetsInput"}

	if s.TargetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetGroupArn"))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/RegisterTargetsOutput
type RegisterTargetsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterTargets = "RegisterTargets"

// RegisterTargetsRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Registers the specified targets with the specified target group.
//
// If the target is an EC2 instance, it must be in the running state when you
// register it.
//
// By default, the load balancer routes requests to registered targets using
// the protocol and port for the target group. Alternatively, you can override
// the port for a target when you register it. You can register each EC2 instance
// or IP address with the same target group multiple times using different ports.
//
// With a Network Load Balancer, you cannot register instances by instance ID
// if they have the following instance types: C1, CC1, CC2, CG1, CG2, CR1, CS1,
// G1, G2, HI1, HS1, M1, M2, M3, and T1. You can register instances of these
// types by IP address.
//
// To remove a target from a target group, use DeregisterTargets.
//
//    // Example sending a request using RegisterTargetsRequest.
//    req := client.RegisterTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/RegisterTargets
func (c *Client) RegisterTargetsRequest(input *RegisterTargetsInput) RegisterTargetsRequest {
	op := &aws.Operation{
		Name:       opRegisterTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterTargetsInput{}
	}

	req := c.newRequest(op, input, &RegisterTargetsOutput{})
	return RegisterTargetsRequest{Request: req, Input: input, Copy: c.RegisterTargetsRequest}
}

// RegisterTargetsRequest is the request type for the
// RegisterTargets API operation.
type RegisterTargetsRequest struct {
	*aws.Request
	Input *RegisterTargetsInput
	Copy  func(*RegisterTargetsInput) RegisterTargetsRequest
}

// Send marshals and sends the RegisterTargets API request.
func (r RegisterTargetsRequest) Send(ctx context.Context) (*RegisterTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterTargetsResponse{
		RegisterTargetsOutput: r.Request.Data.(*RegisterTargetsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterTargetsResponse is the response type for the
// RegisterTargets API operation.
type RegisterTargetsResponse struct {
	*RegisterTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterTargets request.
func (r *RegisterTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
