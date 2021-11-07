// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Forces the instances to stop. The instances do not have an opportunity to
	// flush file system caches or file system metadata. If you use this option,
	// you must perform file system check and repair procedures. This option is
	// not recommended for Windows instances.
	//
	// Default: false
	Force *bool `locationName:"force" type:"boolean"`

	// Hibernates the instance if the instance was enabled for hibernation at launch.
	// If the instance cannot hibernate successfully, a normal shutdown occurs.
	// For more information, see Hibernate your instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Default: false
	Hibernate *bool `type:"boolean"`

	// The IDs of the instances.
	//
	// InstanceIds is a required field
	InstanceIds []string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list" required:"true"`
}

// String returns the string representation
func (s StopInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopInstancesInput"}

	if s.InstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopInstancesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the stopped instances.
	StoppingInstances []InstanceStateChange `locationName:"instancesSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s StopInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopInstances = "StopInstances"

// StopInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Stops an Amazon EBS-backed instance.
//
// You can use the Stop action to hibernate an instance if the instance is enabled
// for hibernation (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#enabling-hibernation)
// and it meets the hibernation prerequisites (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#hibernating-prerequisites).
// For more information, see Hibernate your instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// We don't charge usage for a stopped instance, or data transfer fees; however,
// your root partition Amazon EBS volume remains and continues to persist your
// data, and you are charged for Amazon EBS volume usage. Every time you start
// your Windows instance, Amazon EC2 charges you for a full instance hour. If
// you stop and restart your Windows instance, a new instance hour begins and
// Amazon EC2 charges you for another full instance hour even if you are still
// within the same 60-minute period when it was stopped. Every time you start
// your Linux instance, Amazon EC2 charges a one-minute minimum for instance
// usage, and thereafter charges per second for instance usage.
//
// You can't stop or hibernate instance store-backed instances. You can't use
// the Stop action to hibernate Spot Instances, but you can specify that Amazon
// EC2 should hibernate Spot Instances when they are interrupted. For more information,
// see Hibernating interrupted Spot Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-interruptions.html#hibernate-spot-instances)
// in the Amazon Elastic Compute Cloud User Guide.
//
// When you stop or hibernate an instance, we shut it down. You can restart
// your instance at any time. Before stopping or hibernating an instance, make
// sure it is in a state from which it can be restarted. Stopping an instance
// does not preserve data stored in RAM, but hibernating an instance does preserve
// data stored in RAM. If an instance cannot hibernate successfully, a normal
// shutdown occurs.
//
// Stopping and hibernating an instance is different to rebooting or terminating
// it. For example, when you stop or hibernate an instance, the root device
// and any other devices attached to the instance persist. When you terminate
// an instance, the root device and any other devices attached during the instance
// launch are automatically deleted. For more information about the differences
// between rebooting, stopping, hibernating, and terminating instances, see
// Instance lifecycle (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// When you stop an instance, we attempt to shut it down forcibly after a short
// while. If your instance appears stuck in the stopping state after a period
// of time, there may be an issue with the underlying host computer. For more
// information, see Troubleshooting stopping your instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesStopping.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using StopInstancesRequest.
//    req := client.StopInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/StopInstances
func (c *Client) StopInstancesRequest(input *StopInstancesInput) StopInstancesRequest {
	op := &aws.Operation{
		Name:       opStopInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopInstancesInput{}
	}

	req := c.newRequest(op, input, &StopInstancesOutput{})

	return StopInstancesRequest{Request: req, Input: input, Copy: c.StopInstancesRequest}
}

// StopInstancesRequest is the request type for the
// StopInstances API operation.
type StopInstancesRequest struct {
	*aws.Request
	Input *StopInstancesInput
	Copy  func(*StopInstancesInput) StopInstancesRequest
}

// Send marshals and sends the StopInstances API request.
func (r StopInstancesRequest) Send(ctx context.Context) (*StopInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopInstancesResponse{
		StopInstancesOutput: r.Request.Data.(*StopInstancesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopInstancesResponse is the response type for the
// StopInstances API operation.
type StopInstancesResponse struct {
	*StopInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopInstances request.
func (r *StopInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
