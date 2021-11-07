// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type UpdateSSHPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the SSH public key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// SSHPublicKeyId is a required field
	SSHPublicKeyId *string `min:"20" type:"string" required:"true"`

	// The status to assign to the SSH public key. Active means that the key can
	// be used for authentication with an AWS CodeCommit repository. Inactive means
	// that the key cannot be used.
	//
	// Status is a required field
	Status StatusType `type:"string" required:"true" enum:"true"`

	// The name of the IAM user associated with the SSH public key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateSSHPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSSHPublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSSHPublicKeyInput"}

	if s.SSHPublicKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SSHPublicKeyId"))
	}
	if s.SSHPublicKeyId != nil && len(*s.SSHPublicKeyId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("SSHPublicKeyId", 20))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}

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

type UpdateSSHPublicKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateSSHPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSSHPublicKey = "UpdateSSHPublicKey"

// UpdateSSHPublicKeyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Sets the status of an IAM user's SSH public key to active or inactive. SSH
// public keys that are inactive cannot be used for authentication. This operation
// can be used to disable a user's SSH public key as part of a key rotation
// work flow.
//
// The SSH public key affected by this operation is used only for authenticating
// the associated IAM user to an AWS CodeCommit repository. For more information
// about using SSH keys to authenticate to an AWS CodeCommit repository, see
// Set up AWS CodeCommit for SSH Connections (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html)
// in the AWS CodeCommit User Guide.
//
//    // Example sending a request using UpdateSSHPublicKeyRequest.
//    req := client.UpdateSSHPublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateSSHPublicKey
func (c *Client) UpdateSSHPublicKeyRequest(input *UpdateSSHPublicKeyInput) UpdateSSHPublicKeyRequest {
	op := &aws.Operation{
		Name:       opUpdateSSHPublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSSHPublicKeyInput{}
	}

	req := c.newRequest(op, input, &UpdateSSHPublicKeyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return UpdateSSHPublicKeyRequest{Request: req, Input: input, Copy: c.UpdateSSHPublicKeyRequest}
}

// UpdateSSHPublicKeyRequest is the request type for the
// UpdateSSHPublicKey API operation.
type UpdateSSHPublicKeyRequest struct {
	*aws.Request
	Input *UpdateSSHPublicKeyInput
	Copy  func(*UpdateSSHPublicKeyInput) UpdateSSHPublicKeyRequest
}

// Send marshals and sends the UpdateSSHPublicKey API request.
func (r UpdateSSHPublicKeyRequest) Send(ctx context.Context) (*UpdateSSHPublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSSHPublicKeyResponse{
		UpdateSSHPublicKeyOutput: r.Request.Data.(*UpdateSSHPublicKeyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSSHPublicKeyResponse is the response type for the
// UpdateSSHPublicKey API operation.
type UpdateSSHPublicKeyResponse struct {
	*UpdateSSHPublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSSHPublicKey request.
func (r *UpdateSSHPublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
