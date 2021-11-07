// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTransitGatewayPeeringAttachmentsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * transit-gateway-attachment-id - The ID of the transit gateway attachment.
	//
	//    * local-owner-id - The ID of your AWS account.
	//
	//    * remote-owner-id - The ID of the AWS account in the remote Region that
	//    owns the transit gateway.
	//
	//    * state - The state of the peering attachment (available | deleted | deleting
	//    | failed | modifying | pendingAcceptance | pending | rollingBack | rejected
	//    | rejecting).
	//
	//    * transit-gateway-id - The ID of the transit gateway.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// One or more IDs of the transit gateway peering attachments.
	TransitGatewayAttachmentIds []string `type:"list"`
}

// String returns the string representation
func (s DescribeTransitGatewayPeeringAttachmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTransitGatewayPeeringAttachmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTransitGatewayPeeringAttachmentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTransitGatewayPeeringAttachmentsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The transit gateway peering attachments.
	TransitGatewayPeeringAttachments []TransitGatewayPeeringAttachment `locationName:"transitGatewayPeeringAttachments" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeTransitGatewayPeeringAttachmentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTransitGatewayPeeringAttachments = "DescribeTransitGatewayPeeringAttachments"

// DescribeTransitGatewayPeeringAttachmentsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes your transit gateway peering attachments.
//
//    // Example sending a request using DescribeTransitGatewayPeeringAttachmentsRequest.
//    req := client.DescribeTransitGatewayPeeringAttachmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeTransitGatewayPeeringAttachments
func (c *Client) DescribeTransitGatewayPeeringAttachmentsRequest(input *DescribeTransitGatewayPeeringAttachmentsInput) DescribeTransitGatewayPeeringAttachmentsRequest {
	op := &aws.Operation{
		Name:       opDescribeTransitGatewayPeeringAttachments,
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
		input = &DescribeTransitGatewayPeeringAttachmentsInput{}
	}

	req := c.newRequest(op, input, &DescribeTransitGatewayPeeringAttachmentsOutput{})

	return DescribeTransitGatewayPeeringAttachmentsRequest{Request: req, Input: input, Copy: c.DescribeTransitGatewayPeeringAttachmentsRequest}
}

// DescribeTransitGatewayPeeringAttachmentsRequest is the request type for the
// DescribeTransitGatewayPeeringAttachments API operation.
type DescribeTransitGatewayPeeringAttachmentsRequest struct {
	*aws.Request
	Input *DescribeTransitGatewayPeeringAttachmentsInput
	Copy  func(*DescribeTransitGatewayPeeringAttachmentsInput) DescribeTransitGatewayPeeringAttachmentsRequest
}

// Send marshals and sends the DescribeTransitGatewayPeeringAttachments API request.
func (r DescribeTransitGatewayPeeringAttachmentsRequest) Send(ctx context.Context) (*DescribeTransitGatewayPeeringAttachmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTransitGatewayPeeringAttachmentsResponse{
		DescribeTransitGatewayPeeringAttachmentsOutput: r.Request.Data.(*DescribeTransitGatewayPeeringAttachmentsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTransitGatewayPeeringAttachmentsRequestPaginator returns a paginator for DescribeTransitGatewayPeeringAttachments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTransitGatewayPeeringAttachmentsRequest(input)
//   p := ec2.NewDescribeTransitGatewayPeeringAttachmentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTransitGatewayPeeringAttachmentsPaginator(req DescribeTransitGatewayPeeringAttachmentsRequest) DescribeTransitGatewayPeeringAttachmentsPaginator {
	return DescribeTransitGatewayPeeringAttachmentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTransitGatewayPeeringAttachmentsInput
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

// DescribeTransitGatewayPeeringAttachmentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTransitGatewayPeeringAttachmentsPaginator struct {
	aws.Pager
}

func (p *DescribeTransitGatewayPeeringAttachmentsPaginator) CurrentPage() *DescribeTransitGatewayPeeringAttachmentsOutput {
	return p.Pager.CurrentPage().(*DescribeTransitGatewayPeeringAttachmentsOutput)
}

// DescribeTransitGatewayPeeringAttachmentsResponse is the response type for the
// DescribeTransitGatewayPeeringAttachments API operation.
type DescribeTransitGatewayPeeringAttachmentsResponse struct {
	*DescribeTransitGatewayPeeringAttachmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTransitGatewayPeeringAttachments request.
func (r *DescribeTransitGatewayPeeringAttachmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
