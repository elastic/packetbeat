// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSavingsPlansUtilizationInput struct {
	_ struct{} `type:"structure"`

	// Filters Savings Plans utilization coverage data for active Savings Plans
	// dimensions. You can filter data with the following dimensions:
	//
	//    * LINKED_ACCOUNT
	//
	//    * SAVINGS_PLAN_ARN
	//
	//    * SAVINGS_PLANS_TYPE
	//
	//    * REGION
	//
	//    * PAYMENT_OPTION
	//
	//    * INSTANCE_TYPE_FAMILY
	//
	// GetSavingsPlansUtilization uses the same Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	Filter *Expression `type:"structure"`

	// The granularity of the Amazon Web Services utillization data for your Savings
	// Plans.
	//
	// The GetSavingsPlansUtilization operation supports only DAILY and MONTHLY
	// granularities.
	Granularity Granularity `type:"string" enum:"true"`

	// The time period that you want the usage and costs for. The Start date must
	// be within 13 months. The End date must be after the Start date, and before
	// the current date. Future dates can't be used as an End date.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSavingsPlansUtilizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSavingsPlansUtilizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSavingsPlansUtilizationInput"}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetSavingsPlansUtilizationOutput struct {
	_ struct{} `type:"structure"`

	// The amount of cost/commitment you used your Savings Plans. This allows you
	// to specify date ranges.
	SavingsPlansUtilizationsByTime []SavingsPlansUtilizationByTime `type:"list"`

	// The total amount of cost/commitment that you used your Savings Plans, regardless
	// of date ranges.
	//
	// Total is a required field
	Total *SavingsPlansUtilizationAggregates `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSavingsPlansUtilizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSavingsPlansUtilization = "GetSavingsPlansUtilization"

// GetSavingsPlansUtilizationRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves the Savings Plans utilization for your account across date ranges
// with daily or monthly granularity. Master accounts in an organization have
// access to member accounts. You can use GetDimensionValues in SAVINGS_PLANS
// to determine the possible dimension values.
//
// You cannot group by any dimension values for GetSavingsPlansUtilization.
//
//    // Example sending a request using GetSavingsPlansUtilizationRequest.
//    req := client.GetSavingsPlansUtilizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetSavingsPlansUtilization
func (c *Client) GetSavingsPlansUtilizationRequest(input *GetSavingsPlansUtilizationInput) GetSavingsPlansUtilizationRequest {
	op := &aws.Operation{
		Name:       opGetSavingsPlansUtilization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSavingsPlansUtilizationInput{}
	}

	req := c.newRequest(op, input, &GetSavingsPlansUtilizationOutput{})

	return GetSavingsPlansUtilizationRequest{Request: req, Input: input, Copy: c.GetSavingsPlansUtilizationRequest}
}

// GetSavingsPlansUtilizationRequest is the request type for the
// GetSavingsPlansUtilization API operation.
type GetSavingsPlansUtilizationRequest struct {
	*aws.Request
	Input *GetSavingsPlansUtilizationInput
	Copy  func(*GetSavingsPlansUtilizationInput) GetSavingsPlansUtilizationRequest
}

// Send marshals and sends the GetSavingsPlansUtilization API request.
func (r GetSavingsPlansUtilizationRequest) Send(ctx context.Context) (*GetSavingsPlansUtilizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSavingsPlansUtilizationResponse{
		GetSavingsPlansUtilizationOutput: r.Request.Data.(*GetSavingsPlansUtilizationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSavingsPlansUtilizationResponse is the response type for the
// GetSavingsPlansUtilization API operation.
type GetSavingsPlansUtilizationResponse struct {
	*GetSavingsPlansUtilizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSavingsPlansUtilization request.
func (r *GetSavingsPlansUtilizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
