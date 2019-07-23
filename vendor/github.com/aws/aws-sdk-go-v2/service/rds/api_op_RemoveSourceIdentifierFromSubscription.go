// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RemoveSourceIdentifierFromSubscriptionMessage
type RemoveSourceIdentifierFromSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// The source identifier to be removed from the subscription, such as the DB
	// instance identifier for a DB instance or the name of a security group.
	//
	// SourceIdentifier is a required field
	SourceIdentifier *string `type:"string" required:"true"`

	// The name of the RDS event notification subscription you want to remove a
	// source identifier from.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveSourceIdentifierFromSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveSourceIdentifierFromSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveSourceIdentifierFromSubscriptionInput"}

	if s.SourceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceIdentifier"))
	}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RemoveSourceIdentifierFromSubscriptionResult
type RemoveSourceIdentifierFromSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful invocation of the DescribeEventSubscriptions
	// action.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s RemoveSourceIdentifierFromSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveSourceIdentifierFromSubscription = "RemoveSourceIdentifierFromSubscription"

// RemoveSourceIdentifierFromSubscriptionRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Removes a source identifier from an existing RDS event notification subscription.
//
//    // Example sending a request using RemoveSourceIdentifierFromSubscriptionRequest.
//    req := client.RemoveSourceIdentifierFromSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RemoveSourceIdentifierFromSubscription
func (c *Client) RemoveSourceIdentifierFromSubscriptionRequest(input *RemoveSourceIdentifierFromSubscriptionInput) RemoveSourceIdentifierFromSubscriptionRequest {
	op := &aws.Operation{
		Name:       opRemoveSourceIdentifierFromSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveSourceIdentifierFromSubscriptionInput{}
	}

	req := c.newRequest(op, input, &RemoveSourceIdentifierFromSubscriptionOutput{})
	return RemoveSourceIdentifierFromSubscriptionRequest{Request: req, Input: input, Copy: c.RemoveSourceIdentifierFromSubscriptionRequest}
}

// RemoveSourceIdentifierFromSubscriptionRequest is the request type for the
// RemoveSourceIdentifierFromSubscription API operation.
type RemoveSourceIdentifierFromSubscriptionRequest struct {
	*aws.Request
	Input *RemoveSourceIdentifierFromSubscriptionInput
	Copy  func(*RemoveSourceIdentifierFromSubscriptionInput) RemoveSourceIdentifierFromSubscriptionRequest
}

// Send marshals and sends the RemoveSourceIdentifierFromSubscription API request.
func (r RemoveSourceIdentifierFromSubscriptionRequest) Send(ctx context.Context) (*RemoveSourceIdentifierFromSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveSourceIdentifierFromSubscriptionResponse{
		RemoveSourceIdentifierFromSubscriptionOutput: r.Request.Data.(*RemoveSourceIdentifierFromSubscriptionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveSourceIdentifierFromSubscriptionResponse is the response type for the
// RemoveSourceIdentifierFromSubscription API operation.
type RemoveSourceIdentifierFromSubscriptionResponse struct {
	*RemoveSourceIdentifierFromSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveSourceIdentifierFromSubscription request.
func (r *RemoveSourceIdentifierFromSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
