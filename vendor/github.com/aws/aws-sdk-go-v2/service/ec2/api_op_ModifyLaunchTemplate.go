// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	//
	// Constraint: Maximum 128 ASCII characters.
	ClientToken *string `type:"string"`

	// The version number of the launch template to set as the default version.
	DefaultVersion *string `locationName:"SetDefaultVersion" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateId *string `type:"string"`

	// The name of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateName *string `min:"3" type:"string"`
}

// String returns the string representation
func (s ModifyLaunchTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyLaunchTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyLaunchTemplateInput"}
	if s.LaunchTemplateName != nil && len(*s.LaunchTemplateName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchTemplateName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyLaunchTemplateOutput struct {
	_ struct{} `type:"structure"`

	// Information about the launch template.
	LaunchTemplate *LaunchTemplate `locationName:"launchTemplate" type:"structure"`
}

// String returns the string representation
func (s ModifyLaunchTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyLaunchTemplate = "ModifyLaunchTemplate"

// ModifyLaunchTemplateRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies a launch template. You can specify which version of the launch template
// to set as the default version. When launching an instance, the default version
// applies when a launch template version is not specified.
//
//    // Example sending a request using ModifyLaunchTemplateRequest.
//    req := client.ModifyLaunchTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyLaunchTemplate
func (c *Client) ModifyLaunchTemplateRequest(input *ModifyLaunchTemplateInput) ModifyLaunchTemplateRequest {
	op := &aws.Operation{
		Name:       opModifyLaunchTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyLaunchTemplateInput{}
	}

	req := c.newRequest(op, input, &ModifyLaunchTemplateOutput{})

	return ModifyLaunchTemplateRequest{Request: req, Input: input, Copy: c.ModifyLaunchTemplateRequest}
}

// ModifyLaunchTemplateRequest is the request type for the
// ModifyLaunchTemplate API operation.
type ModifyLaunchTemplateRequest struct {
	*aws.Request
	Input *ModifyLaunchTemplateInput
	Copy  func(*ModifyLaunchTemplateInput) ModifyLaunchTemplateRequest
}

// Send marshals and sends the ModifyLaunchTemplate API request.
func (r ModifyLaunchTemplateRequest) Send(ctx context.Context) (*ModifyLaunchTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyLaunchTemplateResponse{
		ModifyLaunchTemplateOutput: r.Request.Data.(*ModifyLaunchTemplateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyLaunchTemplateResponse is the response type for the
// ModifyLaunchTemplate API operation.
type ModifyLaunchTemplateResponse struct {
	*ModifyLaunchTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyLaunchTemplate request.
func (r *ModifyLaunchTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
