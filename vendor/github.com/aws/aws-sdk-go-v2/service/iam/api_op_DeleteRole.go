// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteRoleInput struct {
	_ struct{} `type:"structure"`

	// The name of the role to delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRoleInput"}

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

type DeleteRoleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRole = "DeleteRole"

// DeleteRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified role. The role must not have any policies attached.
// For more information about roles, go to Working with Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
//
// Make sure that you do not have any Amazon EC2 instances running with the
// role you are about to delete. Deleting a role or instance profile that is
// associated with a running instance will break any applications running on
// the instance.
//
//    // Example sending a request using DeleteRoleRequest.
//    req := client.DeleteRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteRole
func (c *Client) DeleteRoleRequest(input *DeleteRoleInput) DeleteRoleRequest {
	op := &aws.Operation{
		Name:       opDeleteRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRoleInput{}
	}

	req := c.newRequest(op, input, &DeleteRoleOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteRoleRequest{Request: req, Input: input, Copy: c.DeleteRoleRequest}
}

// DeleteRoleRequest is the request type for the
// DeleteRole API operation.
type DeleteRoleRequest struct {
	*aws.Request
	Input *DeleteRoleInput
	Copy  func(*DeleteRoleInput) DeleteRoleRequest
}

// Send marshals and sends the DeleteRole API request.
func (r DeleteRoleRequest) Send(ctx context.Context) (*DeleteRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRoleResponse{
		DeleteRoleOutput: r.Request.Data.(*DeleteRoleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRoleResponse is the response type for the
// DeleteRole API operation.
type DeleteRoleResponse struct {
	*DeleteRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRole request.
func (r *DeleteRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
