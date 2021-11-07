// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetContextKeysForPrincipalPolicyInput struct {
	_ struct{} `type:"structure"`

	// An optional list of additional policies for which you want the list of context
	// keys that are referenced.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	PolicyInputList []string `type:"list"`

	// The ARN of a user, group, or role whose policies contain the context keys
	// that you want listed. If you specify a user, the list includes context keys
	// that are found in all policies that are attached to the user. The list also
	// includes all groups that the user is a member of. If you pick a group or
	// a role, then it includes only those context keys that are found in policies
	// attached to that entity. Note that all parameters are shown in unencoded
	// form here for clarity, but must be URL encoded to be included as a part of
	// a real HTML request.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicySourceArn is a required field
	PolicySourceArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s GetContextKeysForPrincipalPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetContextKeysForPrincipalPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetContextKeysForPrincipalPolicyInput"}

	if s.PolicySourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicySourceArn"))
	}
	if s.PolicySourceArn != nil && len(*s.PolicySourceArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicySourceArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetContextKeysForPrincipalPolicy or
// GetContextKeysForCustomPolicy request.
type GetContextKeysForPrincipalPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The list of context keys that are referenced in the input policies.
	ContextKeyNames []string `type:"list"`
}

// String returns the string representation
func (s GetContextKeysForPrincipalPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetContextKeysForPrincipalPolicy = "GetContextKeysForPrincipalPolicy"

// GetContextKeysForPrincipalPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Gets a list of all of the context keys referenced in all the IAM policies
// that are attached to the specified IAM entity. The entity can be an IAM user,
// group, or role. If you specify a user, then the request also includes all
// of the policies attached to groups that the user is a member of.
//
// You can optionally include a list of one or more additional policies, specified
// as strings. If you want to include only a list of policies by string, use
// GetContextKeysForCustomPolicy instead.
//
// Note: This API discloses information about the permissions granted to other
// users. If you do not want users to see other user's permissions, then consider
// allowing them to use GetContextKeysForCustomPolicy instead.
//
// Context keys are variables maintained by AWS and its services that provide
// details about the context of an API query request. Context keys can be evaluated
// by testing against a value in an IAM policy. Use GetContextKeysForPrincipalPolicy
// to understand what key names and values you must supply when you call SimulatePrincipalPolicy.
//
//    // Example sending a request using GetContextKeysForPrincipalPolicyRequest.
//    req := client.GetContextKeysForPrincipalPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetContextKeysForPrincipalPolicy
func (c *Client) GetContextKeysForPrincipalPolicyRequest(input *GetContextKeysForPrincipalPolicyInput) GetContextKeysForPrincipalPolicyRequest {
	op := &aws.Operation{
		Name:       opGetContextKeysForPrincipalPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetContextKeysForPrincipalPolicyInput{}
	}

	req := c.newRequest(op, input, &GetContextKeysForPrincipalPolicyOutput{})

	return GetContextKeysForPrincipalPolicyRequest{Request: req, Input: input, Copy: c.GetContextKeysForPrincipalPolicyRequest}
}

// GetContextKeysForPrincipalPolicyRequest is the request type for the
// GetContextKeysForPrincipalPolicy API operation.
type GetContextKeysForPrincipalPolicyRequest struct {
	*aws.Request
	Input *GetContextKeysForPrincipalPolicyInput
	Copy  func(*GetContextKeysForPrincipalPolicyInput) GetContextKeysForPrincipalPolicyRequest
}

// Send marshals and sends the GetContextKeysForPrincipalPolicy API request.
func (r GetContextKeysForPrincipalPolicyRequest) Send(ctx context.Context) (*GetContextKeysForPrincipalPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetContextKeysForPrincipalPolicyResponse{
		GetContextKeysForPrincipalPolicyOutput: r.Request.Data.(*GetContextKeysForPrincipalPolicyOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetContextKeysForPrincipalPolicyResponse is the response type for the
// GetContextKeysForPrincipalPolicy API operation.
type GetContextKeysForPrincipalPolicyResponse struct {
	*GetContextKeysForPrincipalPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetContextKeysForPrincipalPolicy request.
func (r *GetContextKeysForPrincipalPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
