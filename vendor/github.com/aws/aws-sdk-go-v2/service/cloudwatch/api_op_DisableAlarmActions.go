// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DisableAlarmActionsInput struct {
	_ struct{} `type:"structure"`

	// The names of the alarms.
	//
	// AlarmNames is a required field
	AlarmNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DisableAlarmActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableAlarmActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableAlarmActionsInput"}

	if s.AlarmNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("AlarmNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisableAlarmActionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableAlarmActionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableAlarmActions = "DisableAlarmActions"

// DisableAlarmActionsRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Disables the actions for the specified alarms. When an alarm's actions are
// disabled, the alarm actions do not execute when the alarm state changes.
//
//    // Example sending a request using DisableAlarmActionsRequest.
//    req := client.DisableAlarmActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DisableAlarmActions
func (c *Client) DisableAlarmActionsRequest(input *DisableAlarmActionsInput) DisableAlarmActionsRequest {
	op := &aws.Operation{
		Name:       opDisableAlarmActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableAlarmActionsInput{}
	}

	req := c.newRequest(op, input, &DisableAlarmActionsOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DisableAlarmActionsRequest{Request: req, Input: input, Copy: c.DisableAlarmActionsRequest}
}

// DisableAlarmActionsRequest is the request type for the
// DisableAlarmActions API operation.
type DisableAlarmActionsRequest struct {
	*aws.Request
	Input *DisableAlarmActionsInput
	Copy  func(*DisableAlarmActionsInput) DisableAlarmActionsRequest
}

// Send marshals and sends the DisableAlarmActions API request.
func (r DisableAlarmActionsRequest) Send(ctx context.Context) (*DisableAlarmActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableAlarmActionsResponse{
		DisableAlarmActionsOutput: r.Request.Data.(*DisableAlarmActionsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableAlarmActionsResponse is the response type for the
// DisableAlarmActions API operation.
type DisableAlarmActionsResponse struct {
	*DisableAlarmActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableAlarmActions request.
func (r *DisableAlarmActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
