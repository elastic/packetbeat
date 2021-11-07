// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type ResyncMFADeviceInput struct {
	_ struct{} `type:"structure"`

	// An authentication code emitted by the device.
	//
	// The format for this parameter is a sequence of six digits.
	//
	// AuthenticationCode1 is a required field
	AuthenticationCode1 *string `min:"6" type:"string" required:"true"`

	// A subsequent authentication code emitted by the device.
	//
	// The format for this parameter is a sequence of six digits.
	//
	// AuthenticationCode2 is a required field
	AuthenticationCode2 *string `min:"6" type:"string" required:"true"`

	// Serial number that uniquely identifies the MFA device.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// SerialNumber is a required field
	SerialNumber *string `min:"9" type:"string" required:"true"`

	// The name of the user whose MFA device you want to resynchronize.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResyncMFADeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResyncMFADeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResyncMFADeviceInput"}

	if s.AuthenticationCode1 == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationCode1"))
	}
	if s.AuthenticationCode1 != nil && len(*s.AuthenticationCode1) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationCode1", 6))
	}

	if s.AuthenticationCode2 == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationCode2"))
	}
	if s.AuthenticationCode2 != nil && len(*s.AuthenticationCode2) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationCode2", 6))
	}

	if s.SerialNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("SerialNumber"))
	}
	if s.SerialNumber != nil && len(*s.SerialNumber) < 9 {
		invalidParams.Add(aws.NewErrParamMinLen("SerialNumber", 9))
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

type ResyncMFADeviceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResyncMFADeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opResyncMFADevice = "ResyncMFADevice"

// ResyncMFADeviceRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Synchronizes the specified MFA device with its IAM resource object on the
// AWS servers.
//
// For more information about creating and working with virtual MFA devices,
// go to Using a Virtual MFA Device (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html)
// in the IAM User Guide.
//
//    // Example sending a request using ResyncMFADeviceRequest.
//    req := client.ResyncMFADeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ResyncMFADevice
func (c *Client) ResyncMFADeviceRequest(input *ResyncMFADeviceInput) ResyncMFADeviceRequest {
	op := &aws.Operation{
		Name:       opResyncMFADevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResyncMFADeviceInput{}
	}

	req := c.newRequest(op, input, &ResyncMFADeviceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return ResyncMFADeviceRequest{Request: req, Input: input, Copy: c.ResyncMFADeviceRequest}
}

// ResyncMFADeviceRequest is the request type for the
// ResyncMFADevice API operation.
type ResyncMFADeviceRequest struct {
	*aws.Request
	Input *ResyncMFADeviceInput
	Copy  func(*ResyncMFADeviceInput) ResyncMFADeviceRequest
}

// Send marshals and sends the ResyncMFADevice API request.
func (r ResyncMFADeviceRequest) Send(ctx context.Context) (*ResyncMFADeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResyncMFADeviceResponse{
		ResyncMFADeviceOutput: r.Request.Data.(*ResyncMFADeviceOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResyncMFADeviceResponse is the response type for the
// ResyncMFADevice API operation.
type ResyncMFADeviceResponse struct {
	*ResyncMFADeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResyncMFADevice request.
func (r *ResyncMFADeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
