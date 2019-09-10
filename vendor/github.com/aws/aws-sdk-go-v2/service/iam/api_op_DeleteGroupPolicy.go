// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteGroupPolicyRequest
type DeleteGroupPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name (friendly name, not ARN) identifying the group that the policy is
	// embedded in.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// GroupName is a required field
	GroupName *string `min:"1" type:"string" required:"true"`

	// The name identifying the policy document to delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGroupPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGroupPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGroupPolicyInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteGroupPolicyOutput
type DeleteGroupPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGroupPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteGroupPolicy = "DeleteGroupPolicy"

// DeleteGroupPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified inline policy that is embedded in the specified IAM
// group.
//
// A group can also have managed policies attached to it. To detach a managed
// policy from a group, use DetachGroupPolicy. For more information about policies,
// refer to Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using DeleteGroupPolicyRequest.
//    req := client.DeleteGroupPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteGroupPolicy
func (c *Client) DeleteGroupPolicyRequest(input *DeleteGroupPolicyInput) DeleteGroupPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteGroupPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteGroupPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteGroupPolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteGroupPolicyRequest{Request: req, Input: input, Copy: c.DeleteGroupPolicyRequest}
}

// DeleteGroupPolicyRequest is the request type for the
// DeleteGroupPolicy API operation.
type DeleteGroupPolicyRequest struct {
	*aws.Request
	Input *DeleteGroupPolicyInput
	Copy  func(*DeleteGroupPolicyInput) DeleteGroupPolicyRequest
}

// Send marshals and sends the DeleteGroupPolicy API request.
func (r DeleteGroupPolicyRequest) Send(ctx context.Context) (*DeleteGroupPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGroupPolicyResponse{
		DeleteGroupPolicyOutput: r.Request.Data.(*DeleteGroupPolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGroupPolicyResponse is the response type for the
// DeleteGroupPolicy API operation.
type DeleteGroupPolicyResponse struct {
	*DeleteGroupPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGroupPolicy request.
func (r *DeleteGroupPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
