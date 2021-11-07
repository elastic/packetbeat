// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRoleInput struct {
	_ struct{} `type:"structure"`

	// The name of the IAM role to get information about.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRoleInput"}

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

// Contains the response to a successful GetRole request.
type GetRoleOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the IAM role.
	//
	// Role is a required field
	Role *Role `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRole = "GetRole"

// GetRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves information about the specified role, including the role's path,
// GUID, ARN, and the role's trust policy that grants permission to assume the
// role. For more information about roles, see Working with Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
//
// Policies returned by this API are URL-encoded compliant with RFC 3986 (https://tools.ietf.org/html/rfc3986).
// You can use a URL decoding method to convert the policy back to plain JSON
// text. For example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality.
//
//    // Example sending a request using GetRoleRequest.
//    req := client.GetRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetRole
func (c *Client) GetRoleRequest(input *GetRoleInput) GetRoleRequest {
	op := &aws.Operation{
		Name:       opGetRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRoleInput{}
	}

	req := c.newRequest(op, input, &GetRoleOutput{})

	return GetRoleRequest{Request: req, Input: input, Copy: c.GetRoleRequest}
}

// GetRoleRequest is the request type for the
// GetRole API operation.
type GetRoleRequest struct {
	*aws.Request
	Input *GetRoleInput
	Copy  func(*GetRoleInput) GetRoleRequest
}

// Send marshals and sends the GetRole API request.
func (r GetRoleRequest) Send(ctx context.Context) (*GetRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRoleResponse{
		GetRoleOutput: r.Request.Data.(*GetRoleOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRoleResponse is the response type for the
// GetRole API operation.
type GetRoleResponse struct {
	*GetRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRole request.
func (r *GetRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
