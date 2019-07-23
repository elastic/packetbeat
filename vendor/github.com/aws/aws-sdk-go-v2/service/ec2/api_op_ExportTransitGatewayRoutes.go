// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExportTransitGatewayRoutesRequest
type ExportTransitGatewayRoutesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * attachment.transit-gateway-attachment-id- The id of the transit gateway
	//    attachment.
	//
	//    * attachment.resource-id - The resource id of the transit gateway attachment.
	//
	//    * route-search.exact-match - The exact match of the specified filter.
	//
	//    * route-search.longest-prefix-match - The longest prefix that matches
	//    the route.
	//
	//    * route-search.subnet-of-match - The routes with a subnet that match the
	//    specified CIDR filter.
	//
	//    * route-search.supernet-of-match - The routes with a CIDR that encompass
	//    the CIDR filter. For example, if you have 10.0.1.0/29 and 10.0.1.0/31
	//    routes in your route table and you specify supernet-of-match as 10.0.1.0/30,
	//    then the result returns 10.0.1.0/29.
	//
	//    * state - The state of the attachment (available | deleted | deleting
	//    | failed | modifying | pendingAcceptance | pending | rollingBack | rejected
	//    | rejecting).
	//
	//    * transit-gateway-route-destination-cidr-block - The CIDR range.
	//
	//    * type - The type of roue (active | blackhole).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The name of the S3 bucket.
	//
	// S3Bucket is a required field
	S3Bucket *string `type:"string" required:"true"`

	// The ID of the route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ExportTransitGatewayRoutesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportTransitGatewayRoutesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportTransitGatewayRoutesInput"}

	if s.S3Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Bucket"))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExportTransitGatewayRoutesResult
type ExportTransitGatewayRoutesOutput struct {
	_ struct{} `type:"structure"`

	// The URL of the exported file in Amazon S3. For example, s3://bucket_name/VPCTransitGateway/TransitGatewayRouteTables/file_name.
	S3Location *string `locationName:"s3Location" type:"string"`
}

// String returns the string representation
func (s ExportTransitGatewayRoutesOutput) String() string {
	return awsutil.Prettify(s)
}

const opExportTransitGatewayRoutes = "ExportTransitGatewayRoutes"

// ExportTransitGatewayRoutesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Exports routes from the specified transit gateway route table to the specified
// S3 bucket. By default, all routes are exported. Alternatively, you can filter
// by CIDR range.
//
//    // Example sending a request using ExportTransitGatewayRoutesRequest.
//    req := client.ExportTransitGatewayRoutesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExportTransitGatewayRoutes
func (c *Client) ExportTransitGatewayRoutesRequest(input *ExportTransitGatewayRoutesInput) ExportTransitGatewayRoutesRequest {
	op := &aws.Operation{
		Name:       opExportTransitGatewayRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExportTransitGatewayRoutesInput{}
	}

	req := c.newRequest(op, input, &ExportTransitGatewayRoutesOutput{})
	return ExportTransitGatewayRoutesRequest{Request: req, Input: input, Copy: c.ExportTransitGatewayRoutesRequest}
}

// ExportTransitGatewayRoutesRequest is the request type for the
// ExportTransitGatewayRoutes API operation.
type ExportTransitGatewayRoutesRequest struct {
	*aws.Request
	Input *ExportTransitGatewayRoutesInput
	Copy  func(*ExportTransitGatewayRoutesInput) ExportTransitGatewayRoutesRequest
}

// Send marshals and sends the ExportTransitGatewayRoutes API request.
func (r ExportTransitGatewayRoutesRequest) Send(ctx context.Context) (*ExportTransitGatewayRoutesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportTransitGatewayRoutesResponse{
		ExportTransitGatewayRoutesOutput: r.Request.Data.(*ExportTransitGatewayRoutesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportTransitGatewayRoutesResponse is the response type for the
// ExportTransitGatewayRoutes API operation.
type ExportTransitGatewayRoutesResponse struct {
	*ExportTransitGatewayRoutesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportTransitGatewayRoutes request.
func (r *ExportTransitGatewayRoutesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
