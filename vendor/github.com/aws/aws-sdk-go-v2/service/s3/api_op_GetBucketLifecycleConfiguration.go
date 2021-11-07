// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type GetBucketLifecycleConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket for which to get the lifecycle information.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketLifecycleConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketLifecycleConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketLifecycleConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketLifecycleConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLifecycleConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetBucketLifecycleConfigurationInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetBucketLifecycleConfigurationInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetBucketLifecycleConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Container for a lifecycle rule.
	Rules []LifecycleRule `locationName:"Rule" type:"list" flattened:"true"`
}

// String returns the string representation
func (s GetBucketLifecycleConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketLifecycleConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Rules != nil {
		v := s.Rules

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Rule", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetBucketLifecycleConfiguration = "GetBucketLifecycleConfiguration"

// GetBucketLifecycleConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
//
// Bucket lifecycle configuration now supports specifying a lifecycle rule using
// an object key name prefix, one or more object tags, or a combination of both.
// Accordingly, this section describes the latest API. The response describes
// the new filter element that you can use to specify a filter to select a subset
// of objects to which the rule applies. If you are still using previous version
// of the lifecycle configuration, it works. For the earlier API description,
// see GetBucketLifecycle.
//
// Returns the lifecycle configuration information set on the bucket. For information
// about lifecycle configuration, see Object Lifecycle Management (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html).
//
// To use this operation, you must have permission to perform the s3:GetLifecycleConfiguration
// action. The bucket owner has this permission, by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// GetBucketLifecycleConfiguration has the following special error:
//
//    * Error code: NoSuchLifecycleConfiguration Description: The lifecycle
//    configuration does not exist. HTTP Status Code: 404 Not Found SOAP Fault
//    Code Prefix: Client
//
// The following operations are related to GetBucketLifecycleConfiguration:
//
//    * GetBucketLifecycle
//
//    * PutBucketLifecycle
//
//    * DeleteBucketLifecycle
//
//    // Example sending a request using GetBucketLifecycleConfigurationRequest.
//    req := client.GetBucketLifecycleConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketLifecycleConfiguration
func (c *Client) GetBucketLifecycleConfigurationRequest(input *GetBucketLifecycleConfigurationInput) GetBucketLifecycleConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketLifecycleConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &GetBucketLifecycleConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetBucketLifecycleConfigurationOutput{})

	return GetBucketLifecycleConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketLifecycleConfigurationRequest}
}

// GetBucketLifecycleConfigurationRequest is the request type for the
// GetBucketLifecycleConfiguration API operation.
type GetBucketLifecycleConfigurationRequest struct {
	*aws.Request
	Input *GetBucketLifecycleConfigurationInput
	Copy  func(*GetBucketLifecycleConfigurationInput) GetBucketLifecycleConfigurationRequest
}

// Send marshals and sends the GetBucketLifecycleConfiguration API request.
func (r GetBucketLifecycleConfigurationRequest) Send(ctx context.Context) (*GetBucketLifecycleConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketLifecycleConfigurationResponse{
		GetBucketLifecycleConfigurationOutput: r.Request.Data.(*GetBucketLifecycleConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketLifecycleConfigurationResponse is the response type for the
// GetBucketLifecycleConfiguration API operation.
type GetBucketLifecycleConfigurationResponse struct {
	*GetBucketLifecycleConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketLifecycleConfiguration request.
func (r *GetBucketLifecycleConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
