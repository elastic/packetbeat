// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeStaleSecurityGroupsRequest
type DescribeStaleSecurityGroupsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The maximum number of items to return for this request. The request returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a prior call.)
	NextToken *string `min:"1" type:"string"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStaleSecurityGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStaleSecurityGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStaleSecurityGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeStaleSecurityGroupsResult
type DescribeStaleSecurityGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the stale security groups.
	StaleSecurityGroupSet []StaleSecurityGroup `locationName:"staleSecurityGroupSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeStaleSecurityGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStaleSecurityGroups = "DescribeStaleSecurityGroups"

// DescribeStaleSecurityGroupsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// [VPC only] Describes the stale security group rules for security groups in
// a specified VPC. Rules are stale when they reference a deleted security group
// in a peer VPC, or a security group in a peer VPC for which the VPC peering
// connection has been deleted.
//
//    // Example sending a request using DescribeStaleSecurityGroupsRequest.
//    req := client.DescribeStaleSecurityGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeStaleSecurityGroups
func (c *Client) DescribeStaleSecurityGroupsRequest(input *DescribeStaleSecurityGroupsInput) DescribeStaleSecurityGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeStaleSecurityGroups,
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
		input = &DescribeStaleSecurityGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeStaleSecurityGroupsOutput{})
	return DescribeStaleSecurityGroupsRequest{Request: req, Input: input, Copy: c.DescribeStaleSecurityGroupsRequest}
}

// DescribeStaleSecurityGroupsRequest is the request type for the
// DescribeStaleSecurityGroups API operation.
type DescribeStaleSecurityGroupsRequest struct {
	*aws.Request
	Input *DescribeStaleSecurityGroupsInput
	Copy  func(*DescribeStaleSecurityGroupsInput) DescribeStaleSecurityGroupsRequest
}

// Send marshals and sends the DescribeStaleSecurityGroups API request.
func (r DescribeStaleSecurityGroupsRequest) Send(ctx context.Context) (*DescribeStaleSecurityGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStaleSecurityGroupsResponse{
		DescribeStaleSecurityGroupsOutput: r.Request.Data.(*DescribeStaleSecurityGroupsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeStaleSecurityGroupsRequestPaginator returns a paginator for DescribeStaleSecurityGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeStaleSecurityGroupsRequest(input)
//   p := ec2.NewDescribeStaleSecurityGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeStaleSecurityGroupsPaginator(req DescribeStaleSecurityGroupsRequest) DescribeStaleSecurityGroupsPaginator {
	return DescribeStaleSecurityGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeStaleSecurityGroupsInput
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

// DescribeStaleSecurityGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeStaleSecurityGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeStaleSecurityGroupsPaginator) CurrentPage() *DescribeStaleSecurityGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeStaleSecurityGroupsOutput)
}

// DescribeStaleSecurityGroupsResponse is the response type for the
// DescribeStaleSecurityGroups API operation.
type DescribeStaleSecurityGroupsResponse struct {
	*DescribeStaleSecurityGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStaleSecurityGroups request.
func (r *DescribeStaleSecurityGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
