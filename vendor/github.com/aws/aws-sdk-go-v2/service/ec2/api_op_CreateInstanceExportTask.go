// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateInstanceExportTask.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateInstanceExportTaskRequest
type CreateInstanceExportTaskInput struct {
	_ struct{} `type:"structure"`

	// A description for the conversion task or the resource being exported. The
	// maximum length is 255 bytes.
	Description *string `locationName:"description" type:"string"`

	// The format and location for an instance export task.
	ExportToS3Task *ExportToS3TaskSpecification `locationName:"exportToS3" type:"structure"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`

	// The target virtualization environment.
	TargetEnvironment ExportEnvironment `locationName:"targetEnvironment" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateInstanceExportTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInstanceExportTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateInstanceExportTaskInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output for CreateInstanceExportTask.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateInstanceExportTaskResult
type CreateInstanceExportTaskOutput struct {
	_ struct{} `type:"structure"`

	// Information about the instance export task.
	ExportTask *ExportTask `locationName:"exportTask" type:"structure"`
}

// String returns the string representation
func (s CreateInstanceExportTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateInstanceExportTask = "CreateInstanceExportTask"

// CreateInstanceExportTaskRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Exports a running or stopped instance to an S3 bucket.
//
// For information about the supported operating systems, image formats, and
// known limitations for the types of instances you can export, see Exporting
// an Instance as a VM Using VM Import/Export (https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport.html)
// in the VM Import/Export User Guide.
//
//    // Example sending a request using CreateInstanceExportTaskRequest.
//    req := client.CreateInstanceExportTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateInstanceExportTask
func (c *Client) CreateInstanceExportTaskRequest(input *CreateInstanceExportTaskInput) CreateInstanceExportTaskRequest {
	op := &aws.Operation{
		Name:       opCreateInstanceExportTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateInstanceExportTaskInput{}
	}

	req := c.newRequest(op, input, &CreateInstanceExportTaskOutput{})
	return CreateInstanceExportTaskRequest{Request: req, Input: input, Copy: c.CreateInstanceExportTaskRequest}
}

// CreateInstanceExportTaskRequest is the request type for the
// CreateInstanceExportTask API operation.
type CreateInstanceExportTaskRequest struct {
	*aws.Request
	Input *CreateInstanceExportTaskInput
	Copy  func(*CreateInstanceExportTaskInput) CreateInstanceExportTaskRequest
}

// Send marshals and sends the CreateInstanceExportTask API request.
func (r CreateInstanceExportTaskRequest) Send(ctx context.Context) (*CreateInstanceExportTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInstanceExportTaskResponse{
		CreateInstanceExportTaskOutput: r.Request.Data.(*CreateInstanceExportTaskOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInstanceExportTaskResponse is the response type for the
// CreateInstanceExportTask API operation.
type CreateInstanceExportTaskResponse struct {
	*CreateInstanceExportTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInstanceExportTask request.
func (r *CreateInstanceExportTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
