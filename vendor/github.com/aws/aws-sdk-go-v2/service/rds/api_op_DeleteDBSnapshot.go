// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDBSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The DB snapshot identifier.
	//
	// Constraints: Must be the name of an existing DB snapshot in the available
	// state.
	//
	// DBSnapshotIdentifier is a required field
	DBSnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBSnapshotInput"}

	if s.DBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB snapshot.
	//
	// This data type is used as a response element in the DescribeDBSnapshots action.
	DBSnapshot *DBSnapshot `type:"structure"`
}

// String returns the string representation
func (s DeleteDBSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDBSnapshot = "DeleteDBSnapshot"

// DeleteDBSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes a DB snapshot. If the snapshot is being copied, the copy operation
// is terminated.
//
// The DB snapshot must be in the available state to be deleted.
//
//    // Example sending a request using DeleteDBSnapshotRequest.
//    req := client.DeleteDBSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteDBSnapshot
func (c *Client) DeleteDBSnapshotRequest(input *DeleteDBSnapshotInput) DeleteDBSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteDBSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteDBSnapshotOutput{})

	return DeleteDBSnapshotRequest{Request: req, Input: input, Copy: c.DeleteDBSnapshotRequest}
}

// DeleteDBSnapshotRequest is the request type for the
// DeleteDBSnapshot API operation.
type DeleteDBSnapshotRequest struct {
	*aws.Request
	Input *DeleteDBSnapshotInput
	Copy  func(*DeleteDBSnapshotInput) DeleteDBSnapshotRequest
}

// Send marshals and sends the DeleteDBSnapshot API request.
func (r DeleteDBSnapshotRequest) Send(ctx context.Context) (*DeleteDBSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBSnapshotResponse{
		DeleteDBSnapshotOutput: r.Request.Data.(*DeleteDBSnapshotOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBSnapshotResponse is the response type for the
// DeleteDBSnapshot API operation.
type DeleteDBSnapshotResponse struct {
	*DeleteDBSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBSnapshot request.
func (r *DeleteDBSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
