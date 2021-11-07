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

type GetBucketMetricsConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket containing the metrics configuration to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The ID used to identify the metrics configuration.
	//
	// Id is a required field
	Id *string `location:"querystring" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketMetricsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketMetricsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketMetricsConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketMetricsConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketMetricsConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "id", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetBucketMetricsConfigurationInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetBucketMetricsConfigurationInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetBucketMetricsConfigurationOutput struct {
	_ struct{} `type:"structure" payload:"MetricsConfiguration"`

	// Specifies the metrics configuration.
	MetricsConfiguration *MetricsConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetBucketMetricsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketMetricsConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MetricsConfiguration != nil {
		v := s.MetricsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "MetricsConfiguration", v, metadata)
	}
	return nil
}

const opGetBucketMetricsConfiguration = "GetBucketMetricsConfiguration"

// GetBucketMetricsConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Gets a metrics configuration (specified by the metrics configuration ID)
// from the bucket. Note that this doesn't include the daily storage metrics.
//
// To use this operation, you must have permissions to perform the s3:GetMetricsConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// For information about CloudWatch request metrics for Amazon S3, see Monitoring
// Metrics with Amazon CloudWatch (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html).
//
// The following operations are related to GetBucketMetricsConfiguration:
//
//    * PutBucketMetricsConfiguration
//
//    * DeleteBucketMetricsConfiguration
//
//    * ListBucketMetricsConfigurations
//
//    * Monitoring Metrics with Amazon CloudWatch (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html)
//
//    // Example sending a request using GetBucketMetricsConfigurationRequest.
//    req := client.GetBucketMetricsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketMetricsConfiguration
func (c *Client) GetBucketMetricsConfigurationRequest(input *GetBucketMetricsConfigurationInput) GetBucketMetricsConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketMetricsConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?metrics",
	}

	if input == nil {
		input = &GetBucketMetricsConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetBucketMetricsConfigurationOutput{})

	return GetBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketMetricsConfigurationRequest}
}

// GetBucketMetricsConfigurationRequest is the request type for the
// GetBucketMetricsConfiguration API operation.
type GetBucketMetricsConfigurationRequest struct {
	*aws.Request
	Input *GetBucketMetricsConfigurationInput
	Copy  func(*GetBucketMetricsConfigurationInput) GetBucketMetricsConfigurationRequest
}

// Send marshals and sends the GetBucketMetricsConfiguration API request.
func (r GetBucketMetricsConfigurationRequest) Send(ctx context.Context) (*GetBucketMetricsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketMetricsConfigurationResponse{
		GetBucketMetricsConfigurationOutput: r.Request.Data.(*GetBucketMetricsConfigurationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketMetricsConfigurationResponse is the response type for the
// GetBucketMetricsConfiguration API operation.
type GetBucketMetricsConfigurationResponse struct {
	*GetBucketMetricsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketMetricsConfiguration request.
func (r *GetBucketMetricsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
