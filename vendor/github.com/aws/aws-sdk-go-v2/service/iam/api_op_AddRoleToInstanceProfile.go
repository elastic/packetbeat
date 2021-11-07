// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type AddRoleToInstanceProfileInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance profile to update.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// InstanceProfileName is a required field
	InstanceProfileName *string `min:"1" type:"string" required:"true"`

	// The name of the role to add.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AddRoleToInstanceProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddRoleToInstanceProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddRoleToInstanceProfileInput"}

	if s.InstanceProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceProfileName"))
	}
	if s.InstanceProfileName != nil && len(*s.InstanceProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceProfileName", 1))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddRoleToInstanceProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddRoleToInstanceProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddRoleToInstanceProfile = "AddRoleToInstanceProfile"

// AddRoleToInstanceProfileRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Adds the specified IAM role to the specified instance profile. An instance
// profile can contain only one role. (The number and size of IAM resources
// in an AWS account are limited. For more information, see IAM and STS Quotas
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html)
// in the IAM User Guide.) You can remove the existing role and then add a different
// role to an instance profile. You must then wait for the change to appear
// across all of AWS because of eventual consistency (https://en.wikipedia.org/wiki/Eventual_consistency).
// To force the change, you must disassociate the instance profile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DisassociateIamInstanceProfile.html)
// and then associate the instance profile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociateIamInstanceProfile.html),
// or you can stop your instance and then restart it.
//
// The caller of this API must be granted the PassRole permission on the IAM
// role by a permissions policy.
//
// For more information about roles, go to Working with Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
// For more information about instance profiles, go to About Instance Profiles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
//
//    // Example sending a request using AddRoleToInstanceProfileRequest.
//    req := client.AddRoleToInstanceProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/AddRoleToInstanceProfile
func (c *Client) AddRoleToInstanceProfileRequest(input *AddRoleToInstanceProfileInput) AddRoleToInstanceProfileRequest {
	op := &aws.Operation{
		Name:       opAddRoleToInstanceProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddRoleToInstanceProfileInput{}
	}

	req := c.newRequest(op, input, &AddRoleToInstanceProfileOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AddRoleToInstanceProfileRequest{Request: req, Input: input, Copy: c.AddRoleToInstanceProfileRequest}
}

// AddRoleToInstanceProfileRequest is the request type for the
// AddRoleToInstanceProfile API operation.
type AddRoleToInstanceProfileRequest struct {
	*aws.Request
	Input *AddRoleToInstanceProfileInput
	Copy  func(*AddRoleToInstanceProfileInput) AddRoleToInstanceProfileRequest
}

// Send marshals and sends the AddRoleToInstanceProfile API request.
func (r AddRoleToInstanceProfileRequest) Send(ctx context.Context) (*AddRoleToInstanceProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddRoleToInstanceProfileResponse{
		AddRoleToInstanceProfileOutput: r.Request.Data.(*AddRoleToInstanceProfileOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddRoleToInstanceProfileResponse is the response type for the
// AddRoleToInstanceProfile API operation.
type AddRoleToInstanceProfileResponse struct {
	*AddRoleToInstanceProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddRoleToInstanceProfile request.
func (r *AddRoleToInstanceProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
