// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the GetTemplateSummary action.
type GetTemplateSummaryInput struct {
	_ struct{} `type:"structure"`

	// The name or the stack ID that is associated with the stack, which are not
	// always interchangeable. For running stacks, you can specify either the stack's
	// name or its unique stack ID. For deleted stack, you must specify the unique
	// stack ID.
	//
	// Conditional: You must specify only one of the following parameters: StackName,
	// StackSetName, TemplateBody, or TemplateURL.
	StackName *string `min:"1" type:"string"`

	// The name or unique ID of the stack set from which the stack was created.
	//
	// Conditional: You must specify only one of the following parameters: StackName,
	// StackSetName, TemplateBody, or TemplateURL.
	StackSetName *string `type:"string"`

	// Structure containing the template body with a minimum length of 1 byte and
	// a maximum length of 51,200 bytes. For more information about templates, see
	// Template Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must specify only one of the following parameters: StackName,
	// StackSetName, TemplateBody, or TemplateURL.
	TemplateBody *string `min:"1" type:"string"`

	// Location of file containing the template body. The URL must point to a template
	// (max size: 460,800 bytes) that is located in an Amazon S3 bucket. For more
	// information about templates, see Template Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must specify only one of the following parameters: StackName,
	// StackSetName, TemplateBody, or TemplateURL.
	TemplateURL *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetTemplateSummaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTemplateSummaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTemplateSummaryInput"}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}
	if s.TemplateBody != nil && len(*s.TemplateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateBody", 1))
	}
	if s.TemplateURL != nil && len(*s.TemplateURL) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateURL", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the GetTemplateSummary action.
type GetTemplateSummaryOutput struct {
	_ struct{} `type:"structure"`

	// The capabilities found within the template. If your template contains IAM
	// resources, you must specify the CAPABILITY_IAM or CAPABILITY_NAMED_IAM value
	// for this parameter when you use the CreateStack or UpdateStack actions with
	// your template; otherwise, those actions return an InsufficientCapabilities
	// error.
	//
	// For more information, see Acknowledging IAM Resources in AWS CloudFormation
	// Templates (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	Capabilities []Capability `type:"list"`

	// The list of resources that generated the values in the Capabilities response
	// element.
	CapabilitiesReason *string `type:"string"`

	// A list of the transforms that are declared in the template.
	DeclaredTransforms []string `type:"list"`

	// The value that is defined in the Description property of the template.
	Description *string `min:"1" type:"string"`

	// The value that is defined for the Metadata property of the template.
	Metadata *string `type:"string"`

	// A list of parameter declarations that describe various properties for each
	// parameter.
	Parameters []ParameterDeclaration `type:"list"`

	// A list of resource identifier summaries that describe the target resources
	// of an import operation and the properties you can provide during the import
	// to identify the target resources. For example, BucketName is a possible identifier
	// property for an AWS::S3::Bucket resource.
	ResourceIdentifierSummaries []ResourceIdentifierSummary `type:"list"`

	// A list of all the template resource types that are defined in the template,
	// such as AWS::EC2::Instance, AWS::Dynamo::Table, and Custom::MyCustomInstance.
	ResourceTypes []string `type:"list"`

	// The AWS template format version, which identifies the capabilities of the
	// template.
	Version *string `type:"string"`
}

// String returns the string representation
func (s GetTemplateSummaryOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTemplateSummary = "GetTemplateSummary"

// GetTemplateSummaryRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns information about a new or existing template. The GetTemplateSummary
// action is useful for viewing parameter information, such as default parameter
// values and parameter types, before you create or update a stack or stack
// set.
//
// You can use the GetTemplateSummary action when you submit a template, or
// you can get template information for a stack set, or a running or deleted
// stack.
//
// For deleted stacks, GetTemplateSummary returns the template information for
// up to 90 days after the stack has been deleted. If the template does not
// exist, a ValidationError is returned.
//
//    // Example sending a request using GetTemplateSummaryRequest.
//    req := client.GetTemplateSummaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/GetTemplateSummary
func (c *Client) GetTemplateSummaryRequest(input *GetTemplateSummaryInput) GetTemplateSummaryRequest {
	op := &aws.Operation{
		Name:       opGetTemplateSummary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTemplateSummaryInput{}
	}

	req := c.newRequest(op, input, &GetTemplateSummaryOutput{})

	return GetTemplateSummaryRequest{Request: req, Input: input, Copy: c.GetTemplateSummaryRequest}
}

// GetTemplateSummaryRequest is the request type for the
// GetTemplateSummary API operation.
type GetTemplateSummaryRequest struct {
	*aws.Request
	Input *GetTemplateSummaryInput
	Copy  func(*GetTemplateSummaryInput) GetTemplateSummaryRequest
}

// Send marshals and sends the GetTemplateSummary API request.
func (r GetTemplateSummaryRequest) Send(ctx context.Context) (*GetTemplateSummaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTemplateSummaryResponse{
		GetTemplateSummaryOutput: r.Request.Data.(*GetTemplateSummaryOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTemplateSummaryResponse is the response type for the
// GetTemplateSummary API operation.
type GetTemplateSummaryResponse struct {
	*GetTemplateSummaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTemplateSummary request.
func (r *GetTemplateSummaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
