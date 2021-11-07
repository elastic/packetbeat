// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteOrganizationalUnitInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the organizational unit that you want to delete.
	// You can get the ID from the ListOrganizationalUnitsForParent operation.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an organizational
	// unit ID string requires "ou-" followed by from 4 to 32 lowercase letters
	// or digits (the ID of the root that contains the OU). This string is followed
	// by a second "-" dash and from 8 to 32 additional lowercase letters or digits.
	//
	// OrganizationalUnitId is a required field
	OrganizationalUnitId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOrganizationalUnitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOrganizationalUnitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOrganizationalUnitInput"}

	if s.OrganizationalUnitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationalUnitId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOrganizationalUnitOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOrganizationalUnitOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteOrganizationalUnit = "DeleteOrganizationalUnit"

// DeleteOrganizationalUnitRequest returns a request value for making API operation for
// AWS Organizations.
//
// Deletes an organizational unit (OU) from a root or another OU. You must first
// remove all accounts and child OUs from the OU that you want to delete.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using DeleteOrganizationalUnitRequest.
//    req := client.DeleteOrganizationalUnitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DeleteOrganizationalUnit
func (c *Client) DeleteOrganizationalUnitRequest(input *DeleteOrganizationalUnitInput) DeleteOrganizationalUnitRequest {
	op := &aws.Operation{
		Name:       opDeleteOrganizationalUnit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteOrganizationalUnitInput{}
	}

	req := c.newRequest(op, input, &DeleteOrganizationalUnitOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteOrganizationalUnitRequest{Request: req, Input: input, Copy: c.DeleteOrganizationalUnitRequest}
}

// DeleteOrganizationalUnitRequest is the request type for the
// DeleteOrganizationalUnit API operation.
type DeleteOrganizationalUnitRequest struct {
	*aws.Request
	Input *DeleteOrganizationalUnitInput
	Copy  func(*DeleteOrganizationalUnitInput) DeleteOrganizationalUnitRequest
}

// Send marshals and sends the DeleteOrganizationalUnit API request.
func (r DeleteOrganizationalUnitRequest) Send(ctx context.Context) (*DeleteOrganizationalUnitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOrganizationalUnitResponse{
		DeleteOrganizationalUnitOutput: r.Request.Data.(*DeleteOrganizationalUnitOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOrganizationalUnitResponse is the response type for the
// DeleteOrganizationalUnit API operation.
type DeleteOrganizationalUnitResponse struct {
	*DeleteOrganizationalUnitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOrganizationalUnit request.
func (r *DeleteOrganizationalUnitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
