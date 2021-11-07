// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBClustersInput struct {
	_ struct{} `type:"structure"`

	// The user-supplied DB cluster identifier. If this parameter is specified,
	// information from only the specific DB cluster is returned. This parameter
	// isn't case-sensitive.
	//
	// Constraints:
	//
	//    * If supplied, must match an existing DBClusterIdentifier.
	DBClusterIdentifier *string `type:"string"`

	// A filter that specifies one or more DB clusters to describe.
	//
	// Supported filters:
	//
	//    * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	//    Resource Names (ARNs). The results list will only include information
	//    about the DB clusters identified by these ARNs.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// Optional Boolean parameter that specifies whether the output includes information
	// about clusters shared from other AWS accounts.
	IncludeShared *bool `type:"boolean"`

	// An optional pagination token provided by a previous DescribeDBClusters request.
	// If this parameter is specified, the response includes only records beyond
	// the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBClustersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBClustersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBClustersInput"}
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

// Contains the result of a successful invocation of the DescribeDBClusters
// action.
type DescribeDBClustersOutput struct {
	_ struct{} `type:"structure"`

	// Contains a list of DB clusters for the user.
	DBClusters []DBCluster `locationNameList:"DBCluster" type:"list"`

	// A pagination token that can be used in a later DescribeDBClusters request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBClustersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBClusters = "DescribeDBClusters"

// DescribeDBClustersRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about provisioned Aurora DB clusters. This API supports
// pagination.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This operation can also return information for Amazon Neptune DB instances
// and Amazon DocumentDB instances.
//
//    // Example sending a request using DescribeDBClustersRequest.
//    req := client.DescribeDBClustersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBClusters
func (c *Client) DescribeDBClustersRequest(input *DescribeDBClustersInput) DescribeDBClustersRequest {
	op := &aws.Operation{
		Name:       opDescribeDBClusters,
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
		input = &DescribeDBClustersInput{}
	}

	req := c.newRequest(op, input, &DescribeDBClustersOutput{})

	return DescribeDBClustersRequest{Request: req, Input: input, Copy: c.DescribeDBClustersRequest}
}

// DescribeDBClustersRequest is the request type for the
// DescribeDBClusters API operation.
type DescribeDBClustersRequest struct {
	*aws.Request
	Input *DescribeDBClustersInput
	Copy  func(*DescribeDBClustersInput) DescribeDBClustersRequest
}

// Send marshals and sends the DescribeDBClusters API request.
func (r DescribeDBClustersRequest) Send(ctx context.Context) (*DescribeDBClustersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBClustersResponse{
		DescribeDBClustersOutput: r.Request.Data.(*DescribeDBClustersOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBClustersRequestPaginator returns a paginator for DescribeDBClusters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBClustersRequest(input)
//   p := rds.NewDescribeDBClustersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBClustersPaginator(req DescribeDBClustersRequest) DescribeDBClustersPaginator {
	return DescribeDBClustersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBClustersInput
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

// DescribeDBClustersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBClustersPaginator struct {
	aws.Pager
}

func (p *DescribeDBClustersPaginator) CurrentPage() *DescribeDBClustersOutput {
	return p.Pager.CurrentPage().(*DescribeDBClustersOutput)
}

// DescribeDBClustersResponse is the response type for the
// DescribeDBClusters API operation.
type DescribeDBClustersResponse struct {
	*DescribeDBClustersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBClusters request.
func (r *DescribeDBClustersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
