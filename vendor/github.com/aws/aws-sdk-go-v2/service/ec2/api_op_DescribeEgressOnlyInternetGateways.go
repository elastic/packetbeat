// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeEgressOnlyInternetGatewaysRequest
type DescribeEgressOnlyInternetGatewaysInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more egress-only internet gateway IDs.
	EgressOnlyInternetGatewayIds []string `locationName:"EgressOnlyInternetGatewayId" locationNameList:"item" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEgressOnlyInternetGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeEgressOnlyInternetGatewaysResult
type DescribeEgressOnlyInternetGatewaysOutput struct {
	_ struct{} `type:"structure"`

	// Information about the egress-only internet gateways.
	EgressOnlyInternetGateways []EgressOnlyInternetGateway `locationName:"egressOnlyInternetGatewaySet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEgressOnlyInternetGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEgressOnlyInternetGateways = "DescribeEgressOnlyInternetGateways"

// DescribeEgressOnlyInternetGatewaysRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your egress-only internet gateways.
//
//    // Example sending a request using DescribeEgressOnlyInternetGatewaysRequest.
//    req := client.DescribeEgressOnlyInternetGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeEgressOnlyInternetGateways
func (c *Client) DescribeEgressOnlyInternetGatewaysRequest(input *DescribeEgressOnlyInternetGatewaysInput) DescribeEgressOnlyInternetGatewaysRequest {
	op := &aws.Operation{
		Name:       opDescribeEgressOnlyInternetGateways,
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
		input = &DescribeEgressOnlyInternetGatewaysInput{}
	}

	req := c.newRequest(op, input, &DescribeEgressOnlyInternetGatewaysOutput{})
	return DescribeEgressOnlyInternetGatewaysRequest{Request: req, Input: input, Copy: c.DescribeEgressOnlyInternetGatewaysRequest}
}

// DescribeEgressOnlyInternetGatewaysRequest is the request type for the
// DescribeEgressOnlyInternetGateways API operation.
type DescribeEgressOnlyInternetGatewaysRequest struct {
	*aws.Request
	Input *DescribeEgressOnlyInternetGatewaysInput
	Copy  func(*DescribeEgressOnlyInternetGatewaysInput) DescribeEgressOnlyInternetGatewaysRequest
}

// Send marshals and sends the DescribeEgressOnlyInternetGateways API request.
func (r DescribeEgressOnlyInternetGatewaysRequest) Send(ctx context.Context) (*DescribeEgressOnlyInternetGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEgressOnlyInternetGatewaysResponse{
		DescribeEgressOnlyInternetGatewaysOutput: r.Request.Data.(*DescribeEgressOnlyInternetGatewaysOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEgressOnlyInternetGatewaysRequestPaginator returns a paginator for DescribeEgressOnlyInternetGateways.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEgressOnlyInternetGatewaysRequest(input)
//   p := ec2.NewDescribeEgressOnlyInternetGatewaysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEgressOnlyInternetGatewaysPaginator(req DescribeEgressOnlyInternetGatewaysRequest) DescribeEgressOnlyInternetGatewaysPaginator {
	return DescribeEgressOnlyInternetGatewaysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEgressOnlyInternetGatewaysInput
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

// DescribeEgressOnlyInternetGatewaysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEgressOnlyInternetGatewaysPaginator struct {
	aws.Pager
}

func (p *DescribeEgressOnlyInternetGatewaysPaginator) CurrentPage() *DescribeEgressOnlyInternetGatewaysOutput {
	return p.Pager.CurrentPage().(*DescribeEgressOnlyInternetGatewaysOutput)
}

// DescribeEgressOnlyInternetGatewaysResponse is the response type for the
// DescribeEgressOnlyInternetGateways API operation.
type DescribeEgressOnlyInternetGatewaysResponse struct {
	*DescribeEgressOnlyInternetGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEgressOnlyInternetGateways request.
func (r *DescribeEgressOnlyInternetGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
