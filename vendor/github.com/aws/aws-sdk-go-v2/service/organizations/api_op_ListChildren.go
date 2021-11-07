// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListChildrenInput struct {
	_ struct{} `type:"structure"`

	// Filters the output to include only the specified child type.
	//
	// ChildType is a required field
	ChildType ChildType `type:"string" required:"true" enum:"true"`

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific
	// to the operation. If additional items exist beyond the maximum you specify,
	// the NextToken response element is present and has a value (is not null).
	// Include that value as the NextToken request parameter in the next call to
	// the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that
	// you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more
	// output is available. Set this parameter to the value of the previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string `type:"string"`

	// The unique identifier (ID) for the parent root or OU whose children you want
	// to list.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//    * Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//    letters or digits.
	//
	//    * Organizational unit (OU) - A string that begins with "ou-" followed
	//    by from 4 to 32 lowercase letters or digits (the ID of the root that the
	//    OU is in). This string is followed by a second "-" dash and from 8 to
	//    32 additional lowercase letters or digits.
	//
	// ParentId is a required field
	ParentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListChildrenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListChildrenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListChildrenInput"}
	if len(s.ChildType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ChildType"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ParentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListChildrenOutput struct {
	_ struct{} `type:"structure"`

	// The list of children of the specified parent container.
	Children []Child `type:"list"`

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You
	// should repeat this until the NextToken response element comes back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListChildrenOutput) String() string {
	return awsutil.Prettify(s)
}

const opListChildren = "ListChildren"

// ListChildrenRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists all of the organizational units (OUs) or accounts that are contained
// in the specified parent OU or root. This operation, along with ListParents
// enables you to traverse the tree structure that makes up this root.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's master account
// or by a member account that is a delegated administrator for an AWS service.
//
//    // Example sending a request using ListChildrenRequest.
//    req := client.ListChildrenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListChildren
func (c *Client) ListChildrenRequest(input *ListChildrenInput) ListChildrenRequest {
	op := &aws.Operation{
		Name:       opListChildren,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListChildrenInput{}
	}

	req := c.newRequest(op, input, &ListChildrenOutput{})

	return ListChildrenRequest{Request: req, Input: input, Copy: c.ListChildrenRequest}
}

// ListChildrenRequest is the request type for the
// ListChildren API operation.
type ListChildrenRequest struct {
	*aws.Request
	Input *ListChildrenInput
	Copy  func(*ListChildrenInput) ListChildrenRequest
}

// Send marshals and sends the ListChildren API request.
func (r ListChildrenRequest) Send(ctx context.Context) (*ListChildrenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListChildrenResponse{
		ListChildrenOutput: r.Request.Data.(*ListChildrenOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListChildrenRequestPaginator returns a paginator for ListChildren.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListChildrenRequest(input)
//   p := organizations.NewListChildrenRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListChildrenPaginator(req ListChildrenRequest) ListChildrenPaginator {
	return ListChildrenPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListChildrenInput
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

// ListChildrenPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListChildrenPaginator struct {
	aws.Pager
}

func (p *ListChildrenPaginator) CurrentPage() *ListChildrenOutput {
	return p.Pager.CurrentPage().(*ListChildrenOutput)
}

// ListChildrenResponse is the response type for the
// ListChildren API operation.
type ListChildrenResponse struct {
	*ListChildrenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListChildren request.
func (r *ListChildrenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
