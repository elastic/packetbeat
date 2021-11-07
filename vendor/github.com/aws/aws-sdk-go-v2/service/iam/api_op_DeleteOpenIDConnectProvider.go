// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteOpenIDConnectProviderInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider resource
	// object to delete. You can get a list of OpenID Connect provider resource
	// ARNs by using the ListOpenIDConnectProviders operation.
	//
	// OpenIDConnectProviderArn is a required field
	OpenIDConnectProviderArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOpenIDConnectProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOpenIDConnectProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOpenIDConnectProviderInput"}

	if s.OpenIDConnectProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("OpenIDConnectProviderArn"))
	}
	if s.OpenIDConnectProviderArn != nil && len(*s.OpenIDConnectProviderArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("OpenIDConnectProviderArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOpenIDConnectProviderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOpenIDConnectProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteOpenIDConnectProvider = "DeleteOpenIDConnectProvider"

// DeleteOpenIDConnectProviderRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes an OpenID Connect identity provider (IdP) resource object in IAM.
//
// Deleting an IAM OIDC provider resource does not update any roles that reference
// the provider as a principal in their trust policies. Any attempt to assume
// a role that references a deleted provider fails.
//
// This operation is idempotent; it does not fail or return an error if you
// call the operation for a provider that does not exist.
//
//    // Example sending a request using DeleteOpenIDConnectProviderRequest.
//    req := client.DeleteOpenIDConnectProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteOpenIDConnectProvider
func (c *Client) DeleteOpenIDConnectProviderRequest(input *DeleteOpenIDConnectProviderInput) DeleteOpenIDConnectProviderRequest {
	op := &aws.Operation{
		Name:       opDeleteOpenIDConnectProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteOpenIDConnectProviderInput{}
	}

	req := c.newRequest(op, input, &DeleteOpenIDConnectProviderOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteOpenIDConnectProviderRequest{Request: req, Input: input, Copy: c.DeleteOpenIDConnectProviderRequest}
}

// DeleteOpenIDConnectProviderRequest is the request type for the
// DeleteOpenIDConnectProvider API operation.
type DeleteOpenIDConnectProviderRequest struct {
	*aws.Request
	Input *DeleteOpenIDConnectProviderInput
	Copy  func(*DeleteOpenIDConnectProviderInput) DeleteOpenIDConnectProviderRequest
}

// Send marshals and sends the DeleteOpenIDConnectProvider API request.
func (r DeleteOpenIDConnectProviderRequest) Send(ctx context.Context) (*DeleteOpenIDConnectProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOpenIDConnectProviderResponse{
		DeleteOpenIDConnectProviderOutput: r.Request.Data.(*DeleteOpenIDConnectProviderOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOpenIDConnectProviderResponse is the response type for the
// DeleteOpenIDConnectProvider API operation.
type DeleteOpenIDConnectProviderResponse struct {
	*DeleteOpenIDConnectProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOpenIDConnectProvider request.
func (r *DeleteOpenIDConnectProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
