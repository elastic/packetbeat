// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListMFADevicesInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The name of the user whose MFA devices you want to list.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListMFADevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMFADevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMFADevicesInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListMFADevices request.
type ListMFADevicesOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// A list of MFA devices.
	//
	// MFADevices is a required field
	MFADevices []MFADevice `type:"list" required:"true"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListMFADevicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListMFADevices = "ListMFADevices"

// ListMFADevicesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the MFA devices for an IAM user. If the request includes a IAM user
// name, then this operation lists all the MFA devices associated with the specified
// user. If you do not specify a user name, IAM determines the user name implicitly
// based on the AWS access key ID signing the request for this API.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using ListMFADevicesRequest.
//    req := client.ListMFADevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListMFADevices
func (c *Client) ListMFADevicesRequest(input *ListMFADevicesInput) ListMFADevicesRequest {
	op := &aws.Operation{
		Name:       opListMFADevices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListMFADevicesInput{}
	}

	req := c.newRequest(op, input, &ListMFADevicesOutput{})

	return ListMFADevicesRequest{Request: req, Input: input, Copy: c.ListMFADevicesRequest}
}

// ListMFADevicesRequest is the request type for the
// ListMFADevices API operation.
type ListMFADevicesRequest struct {
	*aws.Request
	Input *ListMFADevicesInput
	Copy  func(*ListMFADevicesInput) ListMFADevicesRequest
}

// Send marshals and sends the ListMFADevices API request.
func (r ListMFADevicesRequest) Send(ctx context.Context) (*ListMFADevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMFADevicesResponse{
		ListMFADevicesOutput: r.Request.Data.(*ListMFADevicesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMFADevicesRequestPaginator returns a paginator for ListMFADevices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMFADevicesRequest(input)
//   p := iam.NewListMFADevicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMFADevicesPaginator(req ListMFADevicesRequest) ListMFADevicesPaginator {
	return ListMFADevicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMFADevicesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListMFADevicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMFADevicesPaginator struct {
	aws.Pager
}

func (p *ListMFADevicesPaginator) CurrentPage() *ListMFADevicesOutput {
	return p.Pager.CurrentPage().(*ListMFADevicesOutput)
}

// ListMFADevicesResponse is the response type for the
// ListMFADevices API operation.
type ListMFADevicesResponse struct {
	*ListMFADevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMFADevices request.
func (r *ListMFADevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
