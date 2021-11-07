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

type GetBucketRequestPaymentInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket for which to get the payment request configuration
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketRequestPaymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketRequestPaymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketRequestPaymentInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketRequestPaymentInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketRequestPaymentInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetBucketRequestPaymentInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetBucketRequestPaymentInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetBucketRequestPaymentOutput struct {
	_ struct{} `type:"structure"`

	// Specifies who pays for the download and request fees.
	Payer Payer `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetBucketRequestPaymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketRequestPaymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Payer) > 0 {
		v := s.Payer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Payer", v, metadata)
	}
	return nil
}

const opGetBucketRequestPayment = "GetBucketRequestPayment"

// GetBucketRequestPaymentRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the request payment configuration of a bucket. To use this version
// of the operation, you must be the bucket owner. For more information, see
// Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html).
//
// The following operations are related to GetBucketRequestPayment:
//
//    * ListObjects
//
//    // Example sending a request using GetBucketRequestPaymentRequest.
//    req := client.GetBucketRequestPaymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketRequestPayment
func (c *Client) GetBucketRequestPaymentRequest(input *GetBucketRequestPaymentInput) GetBucketRequestPaymentRequest {
	op := &aws.Operation{
		Name:       opGetBucketRequestPayment,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?requestPayment",
	}

	if input == nil {
		input = &GetBucketRequestPaymentInput{}
	}

	req := c.newRequest(op, input, &GetBucketRequestPaymentOutput{})

	return GetBucketRequestPaymentRequest{Request: req, Input: input, Copy: c.GetBucketRequestPaymentRequest}
}

// GetBucketRequestPaymentRequest is the request type for the
// GetBucketRequestPayment API operation.
type GetBucketRequestPaymentRequest struct {
	*aws.Request
	Input *GetBucketRequestPaymentInput
	Copy  func(*GetBucketRequestPaymentInput) GetBucketRequestPaymentRequest
}

// Send marshals and sends the GetBucketRequestPayment API request.
func (r GetBucketRequestPaymentRequest) Send(ctx context.Context) (*GetBucketRequestPaymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketRequestPaymentResponse{
		GetBucketRequestPaymentOutput: r.Request.Data.(*GetBucketRequestPaymentOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketRequestPaymentResponse is the response type for the
// GetBucketRequestPayment API operation.
type GetBucketRequestPaymentResponse struct {
	*GetBucketRequestPaymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketRequestPayment request.
func (r *GetBucketRequestPaymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
