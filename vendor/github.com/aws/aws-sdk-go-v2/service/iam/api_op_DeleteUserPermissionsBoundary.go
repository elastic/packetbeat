// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteUserPermissionsBoundaryInput struct {
	_ struct{} `type:"structure"`

	// The name (friendly name, not ARN) of the IAM user from which you want to
	// remove the permissions boundary.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserPermissionsBoundaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserPermissionsBoundaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserPermissionsBoundaryInput"}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserPermissionsBoundaryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserPermissionsBoundaryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUserPermissionsBoundary = "DeleteUserPermissionsBoundary"

// DeleteUserPermissionsBoundaryRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the permissions boundary for the specified IAM user.
//
// Deleting the permissions boundary for a user might increase its permissions
// by allowing the user to perform all the actions granted in its permissions
// policies.
//
//    // Example sending a request using DeleteUserPermissionsBoundaryRequest.
//    req := client.DeleteUserPermissionsBoundaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteUserPermissionsBoundary
func (c *Client) DeleteUserPermissionsBoundaryRequest(input *DeleteUserPermissionsBoundaryInput) DeleteUserPermissionsBoundaryRequest {
	op := &aws.Operation{
		Name:       opDeleteUserPermissionsBoundary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserPermissionsBoundaryInput{}
	}

	req := c.newRequest(op, input, &DeleteUserPermissionsBoundaryOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteUserPermissionsBoundaryRequest{Request: req, Input: input, Copy: c.DeleteUserPermissionsBoundaryRequest}
}

// DeleteUserPermissionsBoundaryRequest is the request type for the
// DeleteUserPermissionsBoundary API operation.
type DeleteUserPermissionsBoundaryRequest struct {
	*aws.Request
	Input *DeleteUserPermissionsBoundaryInput
	Copy  func(*DeleteUserPermissionsBoundaryInput) DeleteUserPermissionsBoundaryRequest
}

// Send marshals and sends the DeleteUserPermissionsBoundary API request.
func (r DeleteUserPermissionsBoundaryRequest) Send(ctx context.Context) (*DeleteUserPermissionsBoundaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserPermissionsBoundaryResponse{
		DeleteUserPermissionsBoundaryOutput: r.Request.Data.(*DeleteUserPermissionsBoundaryOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserPermissionsBoundaryResponse is the response type for the
// DeleteUserPermissionsBoundary API operation.
type DeleteUserPermissionsBoundaryResponse struct {
	*DeleteUserPermissionsBoundaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUserPermissionsBoundary request.
func (r *DeleteUserPermissionsBoundaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
