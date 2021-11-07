// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateCostCategoryDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for your Cost Category.
	//
	// CostCategoryArn is a required field
	CostCategoryArn *string `min:"20" type:"string" required:"true"`

	// The rule schema version in this particular Cost Category.
	//
	// RuleVersion is a required field
	RuleVersion CostCategoryRuleVersion `type:"string" required:"true" enum:"true"`

	// The Expression object used to categorize costs. For more information, see
	// CostCategoryRule (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html).
	//
	// Rules is a required field
	Rules []CostCategoryRule `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateCostCategoryDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCostCategoryDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCostCategoryDefinitionInput"}

	if s.CostCategoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CostCategoryArn"))
	}
	if s.CostCategoryArn != nil && len(*s.CostCategoryArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CostCategoryArn", 20))
	}
	if len(s.RuleVersion) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RuleVersion"))
	}

	if s.Rules == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rules"))
	}
	if s.Rules != nil && len(s.Rules) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Rules", 1))
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCostCategoryDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for your Cost Category.
	CostCategoryArn *string `min:"20" type:"string"`

	// The Cost Category's effective start date.
	EffectiveStart *string `min:"20" type:"string"`
}

// String returns the string representation
func (s UpdateCostCategoryDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateCostCategoryDefinition = "UpdateCostCategoryDefinition"

// UpdateCostCategoryDefinitionRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Updates an existing Cost Category. Changes made to the Cost Category rules
// will be used to categorize the current month’s expenses and future expenses.
// This won’t change categorization for the previous months.
//
//    // Example sending a request using UpdateCostCategoryDefinitionRequest.
//    req := client.UpdateCostCategoryDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/UpdateCostCategoryDefinition
func (c *Client) UpdateCostCategoryDefinitionRequest(input *UpdateCostCategoryDefinitionInput) UpdateCostCategoryDefinitionRequest {
	op := &aws.Operation{
		Name:       opUpdateCostCategoryDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateCostCategoryDefinitionInput{}
	}

	req := c.newRequest(op, input, &UpdateCostCategoryDefinitionOutput{})

	return UpdateCostCategoryDefinitionRequest{Request: req, Input: input, Copy: c.UpdateCostCategoryDefinitionRequest}
}

// UpdateCostCategoryDefinitionRequest is the request type for the
// UpdateCostCategoryDefinition API operation.
type UpdateCostCategoryDefinitionRequest struct {
	*aws.Request
	Input *UpdateCostCategoryDefinitionInput
	Copy  func(*UpdateCostCategoryDefinitionInput) UpdateCostCategoryDefinitionRequest
}

// Send marshals and sends the UpdateCostCategoryDefinition API request.
func (r UpdateCostCategoryDefinitionRequest) Send(ctx context.Context) (*UpdateCostCategoryDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCostCategoryDefinitionResponse{
		UpdateCostCategoryDefinitionOutput: r.Request.Data.(*UpdateCostCategoryDefinitionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCostCategoryDefinitionResponse is the response type for the
// UpdateCostCategoryDefinition API operation.
type UpdateCostCategoryDefinitionResponse struct {
	*UpdateCostCategoryDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCostCategoryDefinition request.
func (r *UpdateCostCategoryDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
