// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListInstanceProfilesInput struct {
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

	// The path prefix for filtering the results. For example, the prefix /application_abc/component_xyz/
	// gets all instance profiles whose path starts with /application_abc/component_xyz/.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/), listing all instance profiles. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting
	// of either a forward slash (/) by itself or a string that must begin and end
	// with forward slashes. In addition, it can contain any ASCII character from
	// the ! (\u0021) through the DEL character (\u007F), including most punctuation
	// characters, digits, and upper and lowercased letters.
	PathPrefix *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListInstanceProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInstanceProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListInstanceProfilesInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}
	if s.PathPrefix != nil && len(*s.PathPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathPrefix", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListInstanceProfiles request.
type ListInstanceProfilesOutput struct {
	_ struct{} `type:"structure"`

	// A list of instance profiles.
	//
	// InstanceProfiles is a required field
	InstanceProfiles []InstanceProfile `type:"list" required:"true"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListInstanceProfilesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListInstanceProfiles = "ListInstanceProfiles"

// ListInstanceProfilesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the instance profiles that have the specified path prefix. If there
// are none, the operation returns an empty list. For more information about
// instance profiles, go to About Instance Profiles (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
//
// You can paginate the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using ListInstanceProfilesRequest.
//    req := client.ListInstanceProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListInstanceProfiles
func (c *Client) ListInstanceProfilesRequest(input *ListInstanceProfilesInput) ListInstanceProfilesRequest {
	op := &aws.Operation{
		Name:       opListInstanceProfiles,
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
		input = &ListInstanceProfilesInput{}
	}

	req := c.newRequest(op, input, &ListInstanceProfilesOutput{})

	return ListInstanceProfilesRequest{Request: req, Input: input, Copy: c.ListInstanceProfilesRequest}
}

// ListInstanceProfilesRequest is the request type for the
// ListInstanceProfiles API operation.
type ListInstanceProfilesRequest struct {
	*aws.Request
	Input *ListInstanceProfilesInput
	Copy  func(*ListInstanceProfilesInput) ListInstanceProfilesRequest
}

// Send marshals and sends the ListInstanceProfiles API request.
func (r ListInstanceProfilesRequest) Send(ctx context.Context) (*ListInstanceProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInstanceProfilesResponse{
		ListInstanceProfilesOutput: r.Request.Data.(*ListInstanceProfilesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListInstanceProfilesRequestPaginator returns a paginator for ListInstanceProfiles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListInstanceProfilesRequest(input)
//   p := iam.NewListInstanceProfilesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListInstanceProfilesPaginator(req ListInstanceProfilesRequest) ListInstanceProfilesPaginator {
	return ListInstanceProfilesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListInstanceProfilesInput
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

// ListInstanceProfilesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListInstanceProfilesPaginator struct {
	aws.Pager
}

func (p *ListInstanceProfilesPaginator) CurrentPage() *ListInstanceProfilesOutput {
	return p.Pager.CurrentPage().(*ListInstanceProfilesOutput)
}

// ListInstanceProfilesResponse is the response type for the
// ListInstanceProfiles API operation.
type ListInstanceProfilesResponse struct {
	*ListInstanceProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInstanceProfiles request.
func (r *ListInstanceProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
