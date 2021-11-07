// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DeleteChangeSet action.
type DeleteChangeSetInput struct {
	_ struct{} `type:"structure"`

	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// delete.
	//
	// ChangeSetName is a required field
	ChangeSetName *string `min:"1" type:"string" required:"true"`

	// If you specified the name of a change set to delete, specify the stack name
	// or ID (ARN) that is associated with it.
	StackName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteChangeSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteChangeSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteChangeSetInput"}

	if s.ChangeSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeSetName"))
	}
	if s.ChangeSetName != nil && len(*s.ChangeSetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeSetName", 1))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the DeleteChangeSet action.
type DeleteChangeSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteChangeSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteChangeSet = "DeleteChangeSet"

// DeleteChangeSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Deletes the specified change set. Deleting change sets ensures that no one
// executes the wrong change set.
//
// If the call successfully completes, AWS CloudFormation successfully deleted
// the change set.
//
//    // Example sending a request using DeleteChangeSetRequest.
//    req := client.DeleteChangeSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeleteChangeSet
func (c *Client) DeleteChangeSetRequest(input *DeleteChangeSetInput) DeleteChangeSetRequest {
	op := &aws.Operation{
		Name:       opDeleteChangeSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteChangeSetInput{}
	}

	req := c.newRequest(op, input, &DeleteChangeSetOutput{})

	return DeleteChangeSetRequest{Request: req, Input: input, Copy: c.DeleteChangeSetRequest}
}

// DeleteChangeSetRequest is the request type for the
// DeleteChangeSet API operation.
type DeleteChangeSetRequest struct {
	*aws.Request
	Input *DeleteChangeSetInput
	Copy  func(*DeleteChangeSetInput) DeleteChangeSetRequest
}

// Send marshals and sends the DeleteChangeSet API request.
func (r DeleteChangeSetRequest) Send(ctx context.Context) (*DeleteChangeSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteChangeSetResponse{
		DeleteChangeSetOutput: r.Request.Data.(*DeleteChangeSetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteChangeSetResponse is the response type for the
// DeleteChangeSet API operation.
type DeleteChangeSetResponse struct {
	*DeleteChangeSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteChangeSet request.
func (r *DeleteChangeSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
