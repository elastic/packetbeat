// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeInstanceAttributeRequest
type DescribeInstanceAttributeInput struct {
	_ struct{} `type:"structure"`

	// The instance attribute.
	//
	// Note: The enaSupport attribute is not supported at this time.
	//
	// Attribute is a required field
	Attribute InstanceAttributeName `locationName:"attribute" type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `locationName:"instanceId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeInstanceAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstanceAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes an instance attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/InstanceAttribute
type DescribeInstanceAttributeOutput struct {
	_ struct{} `type:"structure"`

	// The block device mapping of the instance.
	BlockDeviceMappings []InstanceBlockDeviceMapping `locationName:"blockDeviceMapping" locationNameList:"item" type:"list"`

	// If the value is true, you can't terminate the instance through the Amazon
	// EC2 console, CLI, or API; otherwise, you can.
	DisableApiTermination *AttributeBooleanValue `locationName:"disableApiTermination" type:"structure"`

	// Indicates whether the instance is optimized for Amazon EBS I/O.
	EbsOptimized *AttributeBooleanValue `locationName:"ebsOptimized" type:"structure"`

	// Indicates whether enhanced networking with ENA is enabled.
	EnaSupport *AttributeBooleanValue `locationName:"enaSupport" type:"structure"`

	// The security groups associated with the instance.
	Groups []GroupIdentifier `locationName:"groupSet" locationNameList:"item" type:"list"`

	// The ID of the instance.
	InstanceId *string `locationName:"instanceId" type:"string"`

	// Indicates whether an instance stops or terminates when you initiate shutdown
	// from the instance (using the operating system command for system shutdown).
	InstanceInitiatedShutdownBehavior *AttributeValue `locationName:"instanceInitiatedShutdownBehavior" type:"structure"`

	// The instance type.
	InstanceType *AttributeValue `locationName:"instanceType" type:"structure"`

	// The kernel ID.
	KernelId *AttributeValue `locationName:"kernel" type:"structure"`

	// A list of product codes.
	ProductCodes []ProductCode `locationName:"productCodes" locationNameList:"item" type:"list"`

	// The RAM disk ID.
	RamdiskId *AttributeValue `locationName:"ramdisk" type:"structure"`

	// The device name of the root device volume (for example, /dev/sda1).
	RootDeviceName *AttributeValue `locationName:"rootDeviceName" type:"structure"`

	// Indicates whether source/destination checking is enabled. A value of true
	// means that checking is enabled, and false means that checking is disabled.
	// This value must be false for a NAT instance to perform NAT.
	SourceDestCheck *AttributeBooleanValue `locationName:"sourceDestCheck" type:"structure"`

	// Indicates whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport *AttributeValue `locationName:"sriovNetSupport" type:"structure"`

	// The user data.
	UserData *AttributeValue `locationName:"userData" type:"structure"`
}

// String returns the string representation
func (s DescribeInstanceAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstanceAttribute = "DescribeInstanceAttribute"

// DescribeInstanceAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified instance. You can specify
// only one attribute at a time. Valid attribute values are: instanceType |
// kernel | ramdisk | userData | disableApiTermination | instanceInitiatedShutdownBehavior
// | rootDeviceName | blockDeviceMapping | productCodes | sourceDestCheck |
// groupSet | ebsOptimized | sriovNetSupport
//
//    // Example sending a request using DescribeInstanceAttributeRequest.
//    req := client.DescribeInstanceAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeInstanceAttribute
func (c *Client) DescribeInstanceAttributeRequest(input *DescribeInstanceAttributeInput) DescribeInstanceAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceAttributeOutput{})
	return DescribeInstanceAttributeRequest{Request: req, Input: input, Copy: c.DescribeInstanceAttributeRequest}
}

// DescribeInstanceAttributeRequest is the request type for the
// DescribeInstanceAttribute API operation.
type DescribeInstanceAttributeRequest struct {
	*aws.Request
	Input *DescribeInstanceAttributeInput
	Copy  func(*DescribeInstanceAttributeInput) DescribeInstanceAttributeRequest
}

// Send marshals and sends the DescribeInstanceAttribute API request.
func (r DescribeInstanceAttributeRequest) Send(ctx context.Context) (*DescribeInstanceAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceAttributeResponse{
		DescribeInstanceAttributeOutput: r.Request.Data.(*DescribeInstanceAttributeOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstanceAttributeResponse is the response type for the
// DescribeInstanceAttribute API operation.
type DescribeInstanceAttributeResponse struct {
	*DescribeInstanceAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceAttribute request.
func (r *DescribeInstanceAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
