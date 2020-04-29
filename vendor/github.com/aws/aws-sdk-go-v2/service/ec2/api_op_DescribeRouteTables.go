// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeRouteTablesRequest
type DescribeRouteTablesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * association.route-table-association-id - The ID of an association ID
	//    for the route table.
	//
	//    * association.route-table-id - The ID of the route table involved in the
	//    association.
	//
	//    * association.subnet-id - The ID of the subnet involved in the association.
	//
	//    * association.main - Indicates whether the route table is the main route
	//    table for the VPC (true | false). Route tables that do not have an association
	//    ID are not returned in the response.
	//
	//    * owner-id - The ID of the AWS account that owns the route table.
	//
	//    * route-table-id - The ID of the route table.
	//
	//    * route.destination-cidr-block - The IPv4 CIDR range specified in a route
	//    in the table.
	//
	//    * route.destination-ipv6-cidr-block - The IPv6 CIDR range specified in
	//    a route in the route table.
	//
	//    * route.destination-prefix-list-id - The ID (prefix) of the AWS service
	//    specified in a route in the table.
	//
	//    * route.egress-only-internet-gateway-id - The ID of an egress-only Internet
	//    gateway specified in a route in the route table.
	//
	//    * route.gateway-id - The ID of a gateway specified in a route in the table.
	//
	//    * route.instance-id - The ID of an instance specified in a route in the
	//    table.
	//
	//    * route.nat-gateway-id - The ID of a NAT gateway.
	//
	//    * route.transit-gateway-id - The ID of a transit gateway.
	//
	//    * route.origin - Describes how the route was created. CreateRouteTable
	//    indicates that the route was automatically created when the route table
	//    was created; CreateRoute indicates that the route was manually added to
	//    the route table; EnableVgwRoutePropagation indicates that the route was
	//    propagated by route propagation.
	//
	//    * route.state - The state of a route in the route table (active | blackhole).
	//    The blackhole state indicates that the route's target isn't available
	//    (for example, the specified gateway isn't attached to the VPC, the specified
	//    NAT instance has been terminated, and so on).
	//
	//    * route.vpc-peering-connection-id - The ID of a VPC peering connection
	//    specified in a route in the table.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * transit-gateway-id - The ID of a transit gateway.
	//
	//    * vpc-id - The ID of the VPC for the route table.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// One or more route table IDs.
	//
	// Default: Describes all your route tables.
	RouteTableIds []string `locationName:"RouteTableId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeRouteTablesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRouteTablesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRouteTablesInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DescribeRouteTables.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeRouteTablesResult
type DescribeRouteTablesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about one or more route tables.
	RouteTables []RouteTable `locationName:"routeTableSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeRouteTablesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRouteTables = "DescribeRouteTables"

// DescribeRouteTablesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your route tables.
//
// Each subnet in your VPC must be associated with a route table. If a subnet
// is not explicitly associated with any route table, it is implicitly associated
// with the main route table. This command does not return the subnet ID for
// implicit associations.
//
// For more information, see Route Tables (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using DescribeRouteTablesRequest.
//    req := client.DescribeRouteTablesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeRouteTables
func (c *Client) DescribeRouteTablesRequest(input *DescribeRouteTablesInput) DescribeRouteTablesRequest {
	op := &aws.Operation{
		Name:       opDescribeRouteTables,
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
		input = &DescribeRouteTablesInput{}
	}

	req := c.newRequest(op, input, &DescribeRouteTablesOutput{})
	return DescribeRouteTablesRequest{Request: req, Input: input, Copy: c.DescribeRouteTablesRequest}
}

// DescribeRouteTablesRequest is the request type for the
// DescribeRouteTables API operation.
type DescribeRouteTablesRequest struct {
	*aws.Request
	Input *DescribeRouteTablesInput
	Copy  func(*DescribeRouteTablesInput) DescribeRouteTablesRequest
}

// Send marshals and sends the DescribeRouteTables API request.
func (r DescribeRouteTablesRequest) Send(ctx context.Context) (*DescribeRouteTablesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRouteTablesResponse{
		DescribeRouteTablesOutput: r.Request.Data.(*DescribeRouteTablesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeRouteTablesRequestPaginator returns a paginator for DescribeRouteTables.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeRouteTablesRequest(input)
//   p := ec2.NewDescribeRouteTablesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeRouteTablesPaginator(req DescribeRouteTablesRequest) DescribeRouteTablesPaginator {
	return DescribeRouteTablesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeRouteTablesInput
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

// DescribeRouteTablesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeRouteTablesPaginator struct {
	aws.Pager
}

func (p *DescribeRouteTablesPaginator) CurrentPage() *DescribeRouteTablesOutput {
	return p.Pager.CurrentPage().(*DescribeRouteTablesOutput)
}

// DescribeRouteTablesResponse is the response type for the
// DescribeRouteTables API operation.
type DescribeRouteTablesResponse struct {
	*DescribeRouteTablesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRouteTables request.
func (r *DescribeRouteTablesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
