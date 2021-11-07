// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ModifyReservedInstances.
type ModifyReservedInstancesInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive token you provide to ensure idempotency of your
	// modification request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The IDs of the Reserved Instances to modify.
	//
	// ReservedInstancesIds is a required field
	ReservedInstancesIds []string `locationName:"ReservedInstancesId" locationNameList:"ReservedInstancesId" type:"list" required:"true"`

	// The configuration settings for the Reserved Instances to modify.
	//
	// TargetConfigurations is a required field
	TargetConfigurations []ReservedInstancesConfiguration `locationName:"ReservedInstancesConfigurationSetItemType" locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyReservedInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyReservedInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyReservedInstancesInput"}

	if s.ReservedInstancesIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedInstancesIds"))
	}

	if s.TargetConfigurations == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetConfigurations"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of ModifyReservedInstances.
type ModifyReservedInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The ID for the modification.
	ReservedInstancesModificationId *string `locationName:"reservedInstancesModificationId" type:"string"`
}

// String returns the string representation
func (s ModifyReservedInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyReservedInstances = "ModifyReservedInstances"

// ModifyReservedInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the Availability Zone, instance count, instance type, or network
// platform (EC2-Classic or EC2-VPC) of your Reserved Instances. The Reserved
// Instances to be modified must be identical, except for Availability Zone,
// network platform, and instance type.
//
// For more information, see Modifying Reserved Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using ModifyReservedInstancesRequest.
//    req := client.ModifyReservedInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyReservedInstances
func (c *Client) ModifyReservedInstancesRequest(input *ModifyReservedInstancesInput) ModifyReservedInstancesRequest {
	op := &aws.Operation{
		Name:       opModifyReservedInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyReservedInstancesInput{}
	}

	req := c.newRequest(op, input, &ModifyReservedInstancesOutput{})

	return ModifyReservedInstancesRequest{Request: req, Input: input, Copy: c.ModifyReservedInstancesRequest}
}

// ModifyReservedInstancesRequest is the request type for the
// ModifyReservedInstances API operation.
type ModifyReservedInstancesRequest struct {
	*aws.Request
	Input *ModifyReservedInstancesInput
	Copy  func(*ModifyReservedInstancesInput) ModifyReservedInstancesRequest
}

// Send marshals and sends the ModifyReservedInstances API request.
func (r ModifyReservedInstancesRequest) Send(ctx context.Context) (*ModifyReservedInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyReservedInstancesResponse{
		ModifyReservedInstancesOutput: r.Request.Data.(*ModifyReservedInstancesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyReservedInstancesResponse is the response type for the
// ModifyReservedInstances API operation.
type ModifyReservedInstancesResponse struct {
	*ModifyReservedInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyReservedInstances request.
func (r *ModifyReservedInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
