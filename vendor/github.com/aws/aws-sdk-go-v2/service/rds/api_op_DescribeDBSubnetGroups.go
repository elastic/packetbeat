// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBSubnetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB subnet group to return details for.
	DBSubnetGroupName *string `type:"string"`

	// This parameter isn't currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBSubnetGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBSubnetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBSubnetGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBSubnetGroupsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the result of a successful invocation of the DescribeDBSubnetGroups
// action.
type DescribeDBSubnetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DBSubnetGroup instances.
	DBSubnetGroups []DBSubnetGroup `locationNameList:"DBSubnetGroup" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBSubnetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBSubnetGroups = "DescribeDBSubnetGroups"

// DescribeDBSubnetGroupsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns a list of DBSubnetGroup descriptions. If a DBSubnetGroupName is specified,
// the list will contain only the descriptions of the specified DBSubnetGroup.
//
// For an overview of CIDR ranges, go to the Wikipedia Tutorial (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
//
//    // Example sending a request using DescribeDBSubnetGroupsRequest.
//    req := client.DescribeDBSubnetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBSubnetGroups
func (c *Client) DescribeDBSubnetGroupsRequest(input *DescribeDBSubnetGroupsInput) DescribeDBSubnetGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBSubnetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBSubnetGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBSubnetGroupsOutput{})

	return DescribeDBSubnetGroupsRequest{Request: req, Input: input, Copy: c.DescribeDBSubnetGroupsRequest}
}

// DescribeDBSubnetGroupsRequest is the request type for the
// DescribeDBSubnetGroups API operation.
type DescribeDBSubnetGroupsRequest struct {
	*aws.Request
	Input *DescribeDBSubnetGroupsInput
	Copy  func(*DescribeDBSubnetGroupsInput) DescribeDBSubnetGroupsRequest
}

// Send marshals and sends the DescribeDBSubnetGroups API request.
func (r DescribeDBSubnetGroupsRequest) Send(ctx context.Context) (*DescribeDBSubnetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBSubnetGroupsResponse{
		DescribeDBSubnetGroupsOutput: r.Request.Data.(*DescribeDBSubnetGroupsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBSubnetGroupsRequestPaginator returns a paginator for DescribeDBSubnetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBSubnetGroupsRequest(input)
//   p := rds.NewDescribeDBSubnetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBSubnetGroupsPaginator(req DescribeDBSubnetGroupsRequest) DescribeDBSubnetGroupsPaginator {
	return DescribeDBSubnetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBSubnetGroupsInput
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

// DescribeDBSubnetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBSubnetGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeDBSubnetGroupsPaginator) CurrentPage() *DescribeDBSubnetGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeDBSubnetGroupsOutput)
}

// DescribeDBSubnetGroupsResponse is the response type for the
// DescribeDBSubnetGroups API operation.
type DescribeDBSubnetGroupsResponse struct {
	*DescribeDBSubnetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBSubnetGroups request.
func (r *DescribeDBSubnetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
