// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DeleteDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the DHCP options set.
	//
	// DhcpOptionsId is a required field
	DhcpOptionsId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s DeleteDhcpOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDhcpOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDhcpOptionsInput"}

	if s.DhcpOptionsId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DhcpOptionsId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDhcpOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDhcpOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDhcpOptions = "DeleteDhcpOptions"

// DeleteDhcpOptionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified set of DHCP options. You must disassociate the set
// of DHCP options before you can delete it. You can disassociate the set of
// DHCP options by associating either a new set of options or the default set
// of options with the VPC.
//
//    // Example sending a request using DeleteDhcpOptionsRequest.
//    req := client.DeleteDhcpOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteDhcpOptions
func (c *Client) DeleteDhcpOptionsRequest(input *DeleteDhcpOptionsInput) DeleteDhcpOptionsRequest {
	op := &aws.Operation{
		Name:       opDeleteDhcpOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDhcpOptionsInput{}
	}

	req := c.newRequest(op, input, &DeleteDhcpOptionsOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteDhcpOptionsRequest{Request: req, Input: input, Copy: c.DeleteDhcpOptionsRequest}
}

// DeleteDhcpOptionsRequest is the request type for the
// DeleteDhcpOptions API operation.
type DeleteDhcpOptionsRequest struct {
	*aws.Request
	Input *DeleteDhcpOptionsInput
	Copy  func(*DeleteDhcpOptionsInput) DeleteDhcpOptionsRequest
}

// Send marshals and sends the DeleteDhcpOptions API request.
func (r DeleteDhcpOptionsRequest) Send(ctx context.Context) (*DeleteDhcpOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDhcpOptionsResponse{
		DeleteDhcpOptionsOutput: r.Request.Data.(*DeleteDhcpOptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDhcpOptionsResponse is the response type for the
// DeleteDhcpOptions API operation.
type DeleteDhcpOptionsResponse struct {
	*DeleteDhcpOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDhcpOptions request.
func (r *DeleteDhcpOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
