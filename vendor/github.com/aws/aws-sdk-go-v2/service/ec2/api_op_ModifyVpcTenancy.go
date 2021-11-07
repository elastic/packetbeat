// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpcTenancyInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The instance tenancy attribute for the VPC.
	//
	// InstanceTenancy is a required field
	InstanceTenancy VpcTenancy `type:"string" required:"true" enum:"true"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcTenancyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcTenancyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpcTenancyInput"}
	if len(s.InstanceTenancy) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("InstanceTenancy"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpcTenancyOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, returns an error.
	ReturnValue *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyVpcTenancyOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVpcTenancy = "ModifyVpcTenancy"

// ModifyVpcTenancyRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the instance tenancy attribute of the specified VPC. You can change
// the instance tenancy attribute of a VPC to default only. You cannot change
// the instance tenancy attribute to dedicated.
//
// After you modify the tenancy of the VPC, any new instances that you launch
// into the VPC have a tenancy of default, unless you specify otherwise during
// launch. The tenancy of any existing instances in the VPC is not affected.
//
// For more information, see Dedicated Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using ModifyVpcTenancyRequest.
//    req := client.ModifyVpcTenancyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpcTenancy
func (c *Client) ModifyVpcTenancyRequest(input *ModifyVpcTenancyInput) ModifyVpcTenancyRequest {
	op := &aws.Operation{
		Name:       opModifyVpcTenancy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcTenancyInput{}
	}

	req := c.newRequest(op, input, &ModifyVpcTenancyOutput{})

	return ModifyVpcTenancyRequest{Request: req, Input: input, Copy: c.ModifyVpcTenancyRequest}
}

// ModifyVpcTenancyRequest is the request type for the
// ModifyVpcTenancy API operation.
type ModifyVpcTenancyRequest struct {
	*aws.Request
	Input *ModifyVpcTenancyInput
	Copy  func(*ModifyVpcTenancyInput) ModifyVpcTenancyRequest
}

// Send marshals and sends the ModifyVpcTenancy API request.
func (r ModifyVpcTenancyRequest) Send(ctx context.Context) (*ModifyVpcTenancyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpcTenancyResponse{
		ModifyVpcTenancyOutput: r.Request.Data.(*ModifyVpcTenancyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpcTenancyResponse is the response type for the
// ModifyVpcTenancy API operation.
type ModifyVpcTenancyResponse struct {
	*ModifyVpcTenancyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpcTenancy request.
func (r *ModifyVpcTenancyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
