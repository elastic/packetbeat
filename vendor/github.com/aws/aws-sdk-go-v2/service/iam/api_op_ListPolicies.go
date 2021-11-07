// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListPoliciesInput struct {
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

	// A flag to filter the results to only the attached policies.
	//
	// When OnlyAttached is true, the returned list contains only the policies that
	// are attached to an IAM user, group, or role. When OnlyAttached is false,
	// or when the parameter is not included, all policies are returned.
	OnlyAttached *bool `type:"boolean"`

	// The path prefix for filtering the results. This parameter is optional. If
	// it is not included, it defaults to a slash (/), listing all policies. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	PathPrefix *string `min:"1" type:"string"`

	// The policy usage method to use for filtering the results.
	//
	// To list only permissions policies, set PolicyUsageFilter to PermissionsPolicy.
	// To list only the policies used to set permissions boundaries, set the value
	// to PermissionsBoundary.
	//
	// This parameter is optional. If it is not included, all policies are returned.
	PolicyUsageFilter PolicyUsageType `type:"string" enum:"true"`

	// The scope to use for filtering the results.
	//
	// To list only AWS managed policies, set Scope to AWS. To list only the customer
	// managed policies in your AWS account, set Scope to Local.
	//
	// This parameter is optional. If it is not included, or if it is set to All,
	// all policies are returned.
	Scope PolicyScopeType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPoliciesInput"}
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

// Contains the response to a successful ListPolicies request.
type ListPoliciesOutput struct {
	_ struct{} `type:"structure"`

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

	// A list of policies.
	Policies []Policy `type:"list"`
}

// String returns the string representation
func (s ListPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPolicies = "ListPolicies"

// ListPoliciesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists all the managed policies that are available in your AWS account, including
// your own customer-defined managed policies and all AWS managed policies.
//
// You can filter the list of policies that is returned using the optional OnlyAttached,
// Scope, and PathPrefix parameters. For example, to list only the customer
// managed policies in your AWS account, set Scope to Local. To list only AWS
// managed policies, set Scope to AWS.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// For more information about managed policies, see Managed Policies and Inline
// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using ListPoliciesRequest.
//    req := client.ListPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListPolicies
func (c *Client) ListPoliciesRequest(input *ListPoliciesInput) ListPoliciesRequest {
	op := &aws.Operation{
		Name:       opListPolicies,
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
		input = &ListPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListPoliciesOutput{})

	return ListPoliciesRequest{Request: req, Input: input, Copy: c.ListPoliciesRequest}
}

// ListPoliciesRequest is the request type for the
// ListPolicies API operation.
type ListPoliciesRequest struct {
	*aws.Request
	Input *ListPoliciesInput
	Copy  func(*ListPoliciesInput) ListPoliciesRequest
}

// Send marshals and sends the ListPolicies API request.
func (r ListPoliciesRequest) Send(ctx context.Context) (*ListPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPoliciesResponse{
		ListPoliciesOutput: r.Request.Data.(*ListPoliciesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPoliciesRequestPaginator returns a paginator for ListPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPoliciesRequest(input)
//   p := iam.NewListPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPoliciesPaginator(req ListPoliciesRequest) ListPoliciesPaginator {
	return ListPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPoliciesInput
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

// ListPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPoliciesPaginator struct {
	aws.Pager
}

func (p *ListPoliciesPaginator) CurrentPage() *ListPoliciesOutput {
	return p.Pager.CurrentPage().(*ListPoliciesOutput)
}

// ListPoliciesResponse is the response type for the
// ListPolicies API operation.
type ListPoliciesResponse struct {
	*ListPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicies request.
func (r *ListPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
