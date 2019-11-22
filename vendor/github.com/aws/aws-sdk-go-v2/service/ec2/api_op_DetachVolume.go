// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DetachVolume.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachVolumeRequest
type DetachVolumeInput struct {
	_ struct{} `type:"structure"`

	// The device name.
	Device *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Forces detachment if the previous detachment attempt did not occur cleanly
	// (for example, logging into an instance, unmounting the volume, and detaching
	// normally). This option can lead to data loss or a corrupted file system.
	// Use this option only as a last resort to detach a volume from a failed instance.
	// The instance won't have an opportunity to flush file system caches or file
	// system metadata. If you use this option, you must perform file system check
	// and repair procedures.
	Force *bool `type:"boolean"`

	// The ID of the instance.
	InstanceId *string `type:"string"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes volume attachment details.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/VolumeAttachment
type DetachVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The time stamp when the attachment initiated.
	AttachTime *time.Time `locationName:"attachTime" type:"timestamp" timestampFormat:"iso8601"`

	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination *bool `locationName:"deleteOnTermination" type:"boolean"`

	// The device name.
	Device *string `locationName:"device" type:"string"`

	// The ID of the instance.
	InstanceId *string `locationName:"instanceId" type:"string"`

	// The attachment state of the volume.
	State VolumeAttachmentState `locationName:"status" type:"string" enum:"true"`

	// The ID of the volume.
	VolumeId *string `locationName:"volumeId" type:"string"`
}

// String returns the string representation
func (s DetachVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachVolume = "DetachVolume"

// DetachVolumeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Detaches an EBS volume from an instance. Make sure to unmount any file systems
// on the device within your operating system before detaching the volume. Failure
// to do so can result in the volume becoming stuck in the busy state while
// detaching. If this happens, detachment can be delayed indefinitely until
// you unmount the volume, force detachment, reboot the instance, or all three.
// If an EBS volume is the root device of an instance, it can't be detached
// while the instance is running. To detach the root volume, stop the instance
// first.
//
// When a volume with an AWS Marketplace product code is detached from an instance,
// the product code is no longer associated with the instance.
//
// For more information, see Detaching an Amazon EBS Volume (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DetachVolumeRequest.
//    req := client.DetachVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachVolume
func (c *Client) DetachVolumeRequest(input *DetachVolumeInput) DetachVolumeRequest {
	op := &aws.Operation{
		Name:       opDetachVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachVolumeInput{}
	}

	req := c.newRequest(op, input, &DetachVolumeOutput{})
	return DetachVolumeRequest{Request: req, Input: input, Copy: c.DetachVolumeRequest}
}

// DetachVolumeRequest is the request type for the
// DetachVolume API operation.
type DetachVolumeRequest struct {
	*aws.Request
	Input *DetachVolumeInput
	Copy  func(*DetachVolumeInput) DetachVolumeRequest
}

// Send marshals and sends the DetachVolume API request.
func (r DetachVolumeRequest) Send(ctx context.Context) (*DetachVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachVolumeResponse{
		DetachVolumeOutput: r.Request.Data.(*DetachVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachVolumeResponse is the response type for the
// DetachVolume API operation.
type DetachVolumeResponse struct {
	*DetachVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachVolume request.
func (r *DetachVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
