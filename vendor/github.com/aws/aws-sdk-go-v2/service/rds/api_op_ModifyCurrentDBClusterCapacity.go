// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyCurrentDBClusterCapacityMessage
type ModifyCurrentDBClusterCapacityInput struct {
	_ struct{} `type:"structure"`

	// The DB cluster capacity.
	//
	// When you change the capacity of a paused Aurora Serverless DB cluster, it
	// automatically resumes.
	//
	// Constraints:
	//
	//    * Value must be 2, 4, 8, 16, 32, 64, 128, or 256.
	Capacity *int64 `type:"integer"`

	// The DB cluster identifier for the cluster being modified. This parameter
	// is not case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DB cluster.
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The amount of time, in seconds, that Aurora Serverless tries to find a scaling
	// point to perform seamless scaling before enforcing the timeout action. The
	// default is 300.
	//
	//    * Value must be from 10 through 600.
	SecondsBeforeTimeout *int64 `type:"integer"`

	// The action to take when the timeout is reached, either ForceApplyCapacityChange
	// or RollbackCapacityChange.
	//
	// ForceApplyCapacityChange, the default, sets the capacity to the specified
	// value as soon as possible.
	//
	// RollbackCapacityChange ignores the capacity change if a scaling point is
	// not found in the timeout period.
	TimeoutAction *string `type:"string"`
}

// String returns the string representation
func (s ModifyCurrentDBClusterCapacityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyCurrentDBClusterCapacityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyCurrentDBClusterCapacityInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DBClusterCapacityInfo
type ModifyCurrentDBClusterCapacityOutput struct {
	_ struct{} `type:"structure"`

	// The current capacity of the DB cluster.
	CurrentCapacity *int64 `type:"integer"`

	// A user-supplied DB cluster identifier. This identifier is the unique key
	// that identifies a DB cluster.
	DBClusterIdentifier *string `type:"string"`

	// A value that specifies the capacity that the DB cluster scales to next.
	PendingCapacity *int64 `type:"integer"`

	// The number of seconds before a call to ModifyCurrentDBClusterCapacity times
	// out.
	SecondsBeforeTimeout *int64 `type:"integer"`

	// The timeout action of a call to ModifyCurrentDBClusterCapacity, either ForceApplyCapacityChange
	// or RollbackCapacityChange.
	TimeoutAction *string `type:"string"`
}

// String returns the string representation
func (s ModifyCurrentDBClusterCapacityOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyCurrentDBClusterCapacity = "ModifyCurrentDBClusterCapacity"

// ModifyCurrentDBClusterCapacityRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Set the capacity of an Aurora Serverless DB cluster to a specific value.
//
// Aurora Serverless scales seamlessly based on the workload on the DB cluster.
// In some cases, the capacity might not scale fast enough to meet a sudden
// change in workload, such as a large number of new transactions. Call ModifyCurrentDBClusterCapacity
// to set the capacity explicitly.
//
// After this call sets the DB cluster capacity, Aurora Serverless can automatically
// scale the DB cluster based on the cooldown period for scaling up and the
// cooldown period for scaling down.
//
// For more information about Aurora Serverless, see Using Amazon Aurora Serverless
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html)
// in the Amazon Aurora User Guide.
//
// If you call ModifyCurrentDBClusterCapacity with the default TimeoutAction,
// connections that prevent Aurora Serverless from finding a scaling point might
// be dropped. For more information about scaling points, see Autoscaling for
// Aurora Serverless (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using ModifyCurrentDBClusterCapacityRequest.
//    req := client.ModifyCurrentDBClusterCapacityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyCurrentDBClusterCapacity
func (c *Client) ModifyCurrentDBClusterCapacityRequest(input *ModifyCurrentDBClusterCapacityInput) ModifyCurrentDBClusterCapacityRequest {
	op := &aws.Operation{
		Name:       opModifyCurrentDBClusterCapacity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyCurrentDBClusterCapacityInput{}
	}

	req := c.newRequest(op, input, &ModifyCurrentDBClusterCapacityOutput{})
	return ModifyCurrentDBClusterCapacityRequest{Request: req, Input: input, Copy: c.ModifyCurrentDBClusterCapacityRequest}
}

// ModifyCurrentDBClusterCapacityRequest is the request type for the
// ModifyCurrentDBClusterCapacity API operation.
type ModifyCurrentDBClusterCapacityRequest struct {
	*aws.Request
	Input *ModifyCurrentDBClusterCapacityInput
	Copy  func(*ModifyCurrentDBClusterCapacityInput) ModifyCurrentDBClusterCapacityRequest
}

// Send marshals and sends the ModifyCurrentDBClusterCapacity API request.
func (r ModifyCurrentDBClusterCapacityRequest) Send(ctx context.Context) (*ModifyCurrentDBClusterCapacityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyCurrentDBClusterCapacityResponse{
		ModifyCurrentDBClusterCapacityOutput: r.Request.Data.(*ModifyCurrentDBClusterCapacityOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyCurrentDBClusterCapacityResponse is the response type for the
// ModifyCurrentDBClusterCapacity API operation.
type ModifyCurrentDBClusterCapacityResponse struct {
	*ModifyCurrentDBClusterCapacityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyCurrentDBClusterCapacity request.
func (r *ModifyCurrentDBClusterCapacityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
