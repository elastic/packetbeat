// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeSourceRegionsMessage
type DescribeSourceRegionsInput struct {
	_ struct{} `type:"structure"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeSourceRegions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The source AWS Region name. For example, us-east-1.
	//
	// Constraints:
	//
	//    * Must specify a valid AWS Region name.
	RegionName *string `type:"string"`
}

// String returns the string representation
func (s DescribeSourceRegionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSourceRegionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSourceRegionsInput"}
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

// Contains the result of a successful invocation of the DescribeSourceRegions
// action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/SourceRegionMessage
type DescribeSourceRegionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of SourceRegion instances that contains each source AWS Region that
	// the current AWS Region can get a Read Replica or a DB snapshot from.
	SourceRegions []SourceRegion `locationNameList:"SourceRegion" type:"list"`
}

// String returns the string representation
func (s DescribeSourceRegionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSourceRegions = "DescribeSourceRegions"

// DescribeSourceRegionsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns a list of the source AWS Regions where the current AWS Region can
// create a Read Replica or copy a DB snapshot from. This API action supports
// pagination.
//
//    // Example sending a request using DescribeSourceRegionsRequest.
//    req := client.DescribeSourceRegionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeSourceRegions
func (c *Client) DescribeSourceRegionsRequest(input *DescribeSourceRegionsInput) DescribeSourceRegionsRequest {
	op := &aws.Operation{
		Name:       opDescribeSourceRegions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSourceRegionsInput{}
	}

	req := c.newRequest(op, input, &DescribeSourceRegionsOutput{})
	return DescribeSourceRegionsRequest{Request: req, Input: input, Copy: c.DescribeSourceRegionsRequest}
}

// DescribeSourceRegionsRequest is the request type for the
// DescribeSourceRegions API operation.
type DescribeSourceRegionsRequest struct {
	*aws.Request
	Input *DescribeSourceRegionsInput
	Copy  func(*DescribeSourceRegionsInput) DescribeSourceRegionsRequest
}

// Send marshals and sends the DescribeSourceRegions API request.
func (r DescribeSourceRegionsRequest) Send(ctx context.Context) (*DescribeSourceRegionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSourceRegionsResponse{
		DescribeSourceRegionsOutput: r.Request.Data.(*DescribeSourceRegionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSourceRegionsResponse is the response type for the
// DescribeSourceRegions API operation.
type DescribeSourceRegionsResponse struct {
	*DescribeSourceRegionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSourceRegions request.
func (r *DescribeSourceRegionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
