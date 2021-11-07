// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetUserInput struct {
	_ struct{} `type:"structure"`

	// The name of the user to get information about.
	//
	// This parameter is optional. If it is not included, it defaults to the user
	// making the request. This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserInput"}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetUser request.
type GetUserOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the IAM user.
	//
	// Due to a service issue, password last used data does not include password
	// use from May 3, 2018 22:50 PDT to May 23, 2018 14:08 PDT. This affects last
	// sign-in (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_finding-unused.html)
	// dates shown in the IAM console and password last used dates in the IAM credential
	// report (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html),
	// and returned by this GetUser API. If users signed in during the affected
	// time, the password last used date that is returned is the date the user last
	// signed in before May 3, 2018. For users that signed in after May 23, 2018
	// 14:08 PDT, the returned password last used date is accurate.
	//
	// You can use password last used information to identify unused credentials
	// for deletion. For example, you might delete users who did not sign in to
	// AWS in the last 90 days. In cases like this, we recommend that you adjust
	// your evaluation window to include dates after May 23, 2018. Alternatively,
	// if your users use access keys to access AWS programmatically you can refer
	// to access key last used information because it is accurate for all dates.
	//
	// User is a required field
	User *User `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUser = "GetUser"

// GetUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves information about the specified IAM user, including the user's
// creation date, path, unique ID, and ARN.
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the AWS access key ID used to sign the request to this API.
//
//    // Example sending a request using GetUserRequest.
//    req := client.GetUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetUser
func (c *Client) GetUserRequest(input *GetUserInput) GetUserRequest {
	op := &aws.Operation{
		Name:       opGetUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUserInput{}
	}

	req := c.newRequest(op, input, &GetUserOutput{})

	return GetUserRequest{Request: req, Input: input, Copy: c.GetUserRequest}
}

// GetUserRequest is the request type for the
// GetUser API operation.
type GetUserRequest struct {
	*aws.Request
	Input *GetUserInput
	Copy  func(*GetUserInput) GetUserRequest
}

// Send marshals and sends the GetUser API request.
func (r GetUserRequest) Send(ctx context.Context) (*GetUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserResponse{
		GetUserOutput: r.Request.Data.(*GetUserOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUserResponse is the response type for the
// GetUser API operation.
type GetUserResponse struct {
	*GetUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUser request.
func (r *GetUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
