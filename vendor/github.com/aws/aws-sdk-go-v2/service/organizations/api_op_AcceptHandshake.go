// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AcceptHandshakeInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the handshake that you want to accept.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string
	// requires "h-" followed by from 8 to 32 lowercase letters or digits.
	//
	// HandshakeId is a required field
	HandshakeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptHandshakeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptHandshakeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptHandshakeInput"}

	if s.HandshakeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HandshakeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AcceptHandshakeOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the accepted handshake.
	Handshake *Handshake `type:"structure"`
}

// String returns the string representation
func (s AcceptHandshakeOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptHandshake = "AcceptHandshake"

// AcceptHandshakeRequest returns a request value for making API operation for
// AWS Organizations.
//
// Sends a response to the originator of a handshake agreeing to the action
// proposed by the handshake request.
//
// This operation can be called only by the following principals when they also
// have the relevant IAM permissions:
//
//    * Invitation to join or Approve all features request handshakes: only
//    a principal from the member account. The user who calls the API for an
//    invitation to join must have the organizations:AcceptHandshake permission.
//    If you enabled all features in the organization, the user must also have
//    the iam:CreateServiceLinkedRole permission so that AWS Organizations can
//    create the required service-linked role named AWSServiceRoleForOrganizations.
//    For more information, see AWS Organizations and Service-Linked Roles (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integration_services.html#orgs_integration_service-linked-roles)
//    in the AWS Organizations User Guide.
//
//    * Enable all features final confirmation handshake: only a principal from
//    the master account. For more information about invitations, see Inviting
//    an AWS Account to Join Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_invites.html)
//    in the AWS Organizations User Guide. For more information about requests
//    to enable all features in the organization, see Enabling All Features
//    in Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
//    in the AWS Organizations User Guide.
//
// After you accept a handshake, it continues to appear in the results of relevant
// APIs for only 30 days. After that, it's deleted.
//
//    // Example sending a request using AcceptHandshakeRequest.
//    req := client.AcceptHandshakeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/AcceptHandshake
func (c *Client) AcceptHandshakeRequest(input *AcceptHandshakeInput) AcceptHandshakeRequest {
	op := &aws.Operation{
		Name:       opAcceptHandshake,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptHandshakeInput{}
	}

	req := c.newRequest(op, input, &AcceptHandshakeOutput{})

	return AcceptHandshakeRequest{Request: req, Input: input, Copy: c.AcceptHandshakeRequest}
}

// AcceptHandshakeRequest is the request type for the
// AcceptHandshake API operation.
type AcceptHandshakeRequest struct {
	*aws.Request
	Input *AcceptHandshakeInput
	Copy  func(*AcceptHandshakeInput) AcceptHandshakeRequest
}

// Send marshals and sends the AcceptHandshake API request.
func (r AcceptHandshakeRequest) Send(ctx context.Context) (*AcceptHandshakeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptHandshakeResponse{
		AcceptHandshakeOutput: r.Request.Data.(*AcceptHandshakeOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptHandshakeResponse is the response type for the
// AcceptHandshake API operation.
type AcceptHandshakeResponse struct {
	*AcceptHandshakeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptHandshake request.
func (r *AcceptHandshakeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
