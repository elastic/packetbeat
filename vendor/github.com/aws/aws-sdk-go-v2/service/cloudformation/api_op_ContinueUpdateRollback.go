// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ContinueUpdateRollback action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ContinueUpdateRollbackInput
type ContinueUpdateRollbackInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for this ContinueUpdateRollback request. Specify this
	// token if you plan to retry requests so that AWS CloudFormation knows that
	// you're not attempting to continue the rollback to a stack with the same name.
	// You might retry ContinueUpdateRollback requests to ensure that AWS CloudFormation
	// successfully received them.
	ClientRequestToken *string `min:"1" type:"string"`

	// A list of the logical IDs of the resources that AWS CloudFormation skips
	// during the continue update rollback operation. You can specify only resources
	// that are in the UPDATE_FAILED state because a rollback failed. You can't
	// specify resources that are in the UPDATE_FAILED state for other reasons,
	// for example, because an update was cancelled. To check why a resource update
	// failed, use the DescribeStackResources action, and view the resource status
	// reason.
	//
	// Specify this property to skip rolling back resources that AWS CloudFormation
	// can't successfully roll back. We recommend that you troubleshoot (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/troubleshooting.html#troubleshooting-errors-update-rollback-failed)
	// resources before skipping them. AWS CloudFormation sets the status of the
	// specified resources to UPDATE_COMPLETE and continues to roll back the stack.
	// After the rollback is complete, the state of the skipped resources will be
	// inconsistent with the state of the resources in the stack template. Before
	// performing another stack update, you must update the stack or resources to
	// be consistent with each other. If you don't, subsequent stack updates might
	// fail, and the stack will become unrecoverable.
	//
	// Specify the minimum number of resources required to successfully roll back
	// your stack. For example, a failed resource update might cause dependent resources
	// to fail. In this case, it might not be necessary to skip the dependent resources.
	//
	// To skip resources that are part of nested stacks, use the following format:
	// NestedStackName.ResourceLogicalID. If you want to specify the logical ID
	// of a stack resource (Type: AWS::CloudFormation::Stack) in the ResourcesToSkip
	// list, then its corresponding embedded stack must be in one of the following
	// states: DELETE_IN_PROGRESS, DELETE_COMPLETE, or DELETE_FAILED.
	//
	// Don't confuse a child stack's name with its corresponding logical ID defined
	// in the parent stack. For an example of a continue update rollback operation
	// with nested stacks, see Using ResourcesToSkip to recover a nested stacks
	// hierarchy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html#nested-stacks).
	ResourcesToSkip []string `type:"list"`

	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM)
	// role that AWS CloudFormation assumes to roll back the stack. AWS CloudFormation
	// uses the role's credentials to make calls on your behalf. AWS CloudFormation
	// always uses this role for all future operations on the stack. As long as
	// users have permission to operate on the stack, AWS CloudFormation uses this
	// role even if the users don't have permission to pass it. Ensure that the
	// role grants least privilege.
	//
	// If you don't specify a value, AWS CloudFormation uses the role that was previously
	// associated with the stack. If no role is available, AWS CloudFormation uses
	// a temporary session that is generated from your user credentials.
	RoleARN *string `min:"20" type:"string"`

	// The name or the unique ID of the stack that you want to continue rolling
	// back.
	//
	// Don't specify the name of a nested stack (a stack that was created by using
	// the AWS::CloudFormation::Stack resource). Instead, use this operation on
	// the parent stack (the stack that contains the AWS::CloudFormation::Stack
	// resource).
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ContinueUpdateRollbackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ContinueUpdateRollbackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ContinueUpdateRollbackInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.RoleARN != nil && len(*s.RoleARN) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleARN", 20))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for a ContinueUpdateRollback action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ContinueUpdateRollbackOutput
type ContinueUpdateRollbackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ContinueUpdateRollbackOutput) String() string {
	return awsutil.Prettify(s)
}

const opContinueUpdateRollback = "ContinueUpdateRollback"

// ContinueUpdateRollbackRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// For a specified stack that is in the UPDATE_ROLLBACK_FAILED state, continues
// rolling it back to the UPDATE_ROLLBACK_COMPLETE state. Depending on the cause
// of the failure, you can manually fix the error (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/troubleshooting.html#troubleshooting-errors-update-rollback-failed)
// and continue the rollback. By continuing the rollback, you can return your
// stack to a working state (the UPDATE_ROLLBACK_COMPLETE state), and then try
// to update the stack again.
//
// A stack goes into the UPDATE_ROLLBACK_FAILED state when AWS CloudFormation
// cannot roll back all changes after a failed stack update. For example, you
// might have a stack that is rolling back to an old database instance that
// was deleted outside of AWS CloudFormation. Because AWS CloudFormation doesn't
// know the database was deleted, it assumes that the database instance still
// exists and attempts to roll back to it, causing the update rollback to fail.
//
//    // Example sending a request using ContinueUpdateRollbackRequest.
//    req := client.ContinueUpdateRollbackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ContinueUpdateRollback
func (c *Client) ContinueUpdateRollbackRequest(input *ContinueUpdateRollbackInput) ContinueUpdateRollbackRequest {
	op := &aws.Operation{
		Name:       opContinueUpdateRollback,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ContinueUpdateRollbackInput{}
	}

	req := c.newRequest(op, input, &ContinueUpdateRollbackOutput{})
	return ContinueUpdateRollbackRequest{Request: req, Input: input, Copy: c.ContinueUpdateRollbackRequest}
}

// ContinueUpdateRollbackRequest is the request type for the
// ContinueUpdateRollback API operation.
type ContinueUpdateRollbackRequest struct {
	*aws.Request
	Input *ContinueUpdateRollbackInput
	Copy  func(*ContinueUpdateRollbackInput) ContinueUpdateRollbackRequest
}

// Send marshals and sends the ContinueUpdateRollback API request.
func (r ContinueUpdateRollbackRequest) Send(ctx context.Context) (*ContinueUpdateRollbackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ContinueUpdateRollbackResponse{
		ContinueUpdateRollbackOutput: r.Request.Data.(*ContinueUpdateRollbackOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ContinueUpdateRollbackResponse is the response type for the
// ContinueUpdateRollback API operation.
type ContinueUpdateRollbackResponse struct {
	*ContinueUpdateRollbackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ContinueUpdateRollback request.
func (r *ContinueUpdateRollbackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
