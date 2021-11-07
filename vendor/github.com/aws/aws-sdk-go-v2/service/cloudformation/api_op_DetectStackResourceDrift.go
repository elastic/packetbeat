// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetectStackResourceDriftInput struct {
	_ struct{} `type:"structure"`

	// The logical name of the resource for which to return drift information.
	//
	// LogicalResourceId is a required field
	LogicalResourceId *string `type:"string" required:"true"`

	// The name of the stack to which the resource belongs.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DetectStackResourceDriftInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetectStackResourceDriftInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetectStackResourceDriftInput"}

	if s.LogicalResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogicalResourceId"))
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

type DetectStackResourceDriftOutput struct {
	_ struct{} `type:"structure"`

	// Information about whether the resource's actual configuration has drifted
	// from its expected template configuration, including actual and expected property
	// values and any differences detected.
	//
	// StackResourceDrift is a required field
	StackResourceDrift *StackResourceDrift `type:"structure" required:"true"`
}

// String returns the string representation
func (s DetectStackResourceDriftOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetectStackResourceDrift = "DetectStackResourceDrift"

// DetectStackResourceDriftRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns information about whether a resource's actual configuration differs,
// or has drifted, from it's expected configuration, as defined in the stack
// template and any values specified as template parameters. This information
// includes actual and expected property values for resources in which AWS CloudFormation
// detects drift. Only resource properties explicitly defined in the stack template
// are checked for drift. For more information about stack and resource drift,
// see Detecting Unregulated Configuration Changes to Stacks and Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html).
//
// Use DetectStackResourceDrift to detect drift on individual resources, or
// DetectStackDrift to detect drift on all resources in a given stack that support
// drift detection.
//
// Resources that do not currently support drift detection cannot be checked.
// For a list of resources that support drift detection, see Resources that
// Support Drift Detection (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html).
//
//    // Example sending a request using DetectStackResourceDriftRequest.
//    req := client.DetectStackResourceDriftRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DetectStackResourceDrift
func (c *Client) DetectStackResourceDriftRequest(input *DetectStackResourceDriftInput) DetectStackResourceDriftRequest {
	op := &aws.Operation{
		Name:       opDetectStackResourceDrift,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetectStackResourceDriftInput{}
	}

	req := c.newRequest(op, input, &DetectStackResourceDriftOutput{})

	return DetectStackResourceDriftRequest{Request: req, Input: input, Copy: c.DetectStackResourceDriftRequest}
}

// DetectStackResourceDriftRequest is the request type for the
// DetectStackResourceDrift API operation.
type DetectStackResourceDriftRequest struct {
	*aws.Request
	Input *DetectStackResourceDriftInput
	Copy  func(*DetectStackResourceDriftInput) DetectStackResourceDriftRequest
}

// Send marshals and sends the DetectStackResourceDrift API request.
func (r DetectStackResourceDriftRequest) Send(ctx context.Context) (*DetectStackResourceDriftResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetectStackResourceDriftResponse{
		DetectStackResourceDriftOutput: r.Request.Data.(*DetectStackResourceDriftOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetectStackResourceDriftResponse is the response type for the
// DetectStackResourceDrift API operation.
type DetectStackResourceDriftResponse struct {
	*DetectStackResourceDriftOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetectStackResourceDrift request.
func (r *DetectStackResourceDriftResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
