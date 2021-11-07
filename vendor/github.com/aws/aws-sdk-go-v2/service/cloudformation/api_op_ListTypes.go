// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTypesInput struct {
	_ struct{} `type:"structure"`

	// The deprecation status of the types that you want to get summary information
	// about.
	//
	// Valid values include:
	//
	//    * LIVE: The type is registered for use in CloudFormation operations.
	//
	//    * DEPRECATED: The type has been deregistered and can no longer be used
	//    in CloudFormation operations.
	DeprecatedStatus DeprecatedStatus `type:"string" enum:"true"`

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to
	// the request object's NextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string `min:"1" type:"string"`

	// The provisioning behavior of the type. AWS CloudFormation determines the
	// provisioning type during registration, based on the types of handlers in
	// the schema handler package submitted.
	//
	// Valid values include:
	//
	//    * FULLY_MUTABLE: The type includes an update handler to process updates
	//    to the type during stack update operations.
	//
	//    * IMMUTABLE: The type does not include an update handler, so the type
	//    cannot be updated and must instead be replaced during stack update operations.
	//
	//    * NON_PROVISIONABLE: The type does not include create, read, and delete
	//    handlers, and therefore cannot actually be provisioned.
	ProvisioningType ProvisioningType `type:"string" enum:"true"`

	// The scope at which the type is visible and usable in CloudFormation operations.
	//
	// Valid values include:
	//
	//    * PRIVATE: The type is only visible and usable within the account in which
	//    it is registered. Currently, AWS CloudFormation marks any types you create
	//    as PRIVATE.
	//
	//    * PUBLIC: The type is publically visible and usable within any Amazon
	//    account.
	//
	// The default is PRIVATE.
	Visibility Visibility `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTypesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTypesOutput struct {
	_ struct{} `type:"structure"`

	// If the request doesn't return all of the remaining results, NextToken is
	// set to a token. To retrieve the next set of results, call this action again
	// and assign that token to the request object's NextToken parameter. If the
	// request returns all results, NextToken is set to null.
	NextToken *string `min:"1" type:"string"`

	// A list of TypeSummary structures that contain information about the specified
	// types.
	TypeSummaries []TypeSummary `type:"list"`
}

// String returns the string representation
func (s ListTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTypes = "ListTypes"

// ListTypesRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns summary information about types that have been registered with CloudFormation.
//
//    // Example sending a request using ListTypesRequest.
//    req := client.ListTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListTypes
func (c *Client) ListTypesRequest(input *ListTypesInput) ListTypesRequest {
	op := &aws.Operation{
		Name:       opListTypes,
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
		input = &ListTypesInput{}
	}

	req := c.newRequest(op, input, &ListTypesOutput{})

	return ListTypesRequest{Request: req, Input: input, Copy: c.ListTypesRequest}
}

// ListTypesRequest is the request type for the
// ListTypes API operation.
type ListTypesRequest struct {
	*aws.Request
	Input *ListTypesInput
	Copy  func(*ListTypesInput) ListTypesRequest
}

// Send marshals and sends the ListTypes API request.
func (r ListTypesRequest) Send(ctx context.Context) (*ListTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTypesResponse{
		ListTypesOutput: r.Request.Data.(*ListTypesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTypesRequestPaginator returns a paginator for ListTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTypesRequest(input)
//   p := cloudformation.NewListTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTypesPaginator(req ListTypesRequest) ListTypesPaginator {
	return ListTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTypesInput
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

// ListTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTypesPaginator struct {
	aws.Pager
}

func (p *ListTypesPaginator) CurrentPage() *ListTypesOutput {
	return p.Pager.CurrentPage().(*ListTypesOutput)
}

// ListTypesResponse is the response type for the
// ListTypes API operation.
type ListTypesResponse struct {
	*ListTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTypes request.
func (r *ListTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
