// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/checksum"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type PutBucketRequestPaymentInput struct {
	_ struct{} `type:"structure" payload:"RequestPaymentConfiguration"`

	// The bucket name.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Container for Payer.
	//
	// RequestPaymentConfiguration is a required field
	RequestPaymentConfiguration *RequestPaymentConfiguration `locationName:"RequestPaymentConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketRequestPaymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketRequestPaymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketRequestPaymentInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.RequestPaymentConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestPaymentConfiguration"))
	}
	if s.RequestPaymentConfiguration != nil {
		if err := s.RequestPaymentConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RequestPaymentConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketRequestPaymentInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketRequestPaymentInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.RequestPaymentConfiguration != nil {
		v := s.RequestPaymentConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "RequestPaymentConfiguration", v, metadata)
	}
	return nil
}

func (s *PutBucketRequestPaymentInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *PutBucketRequestPaymentInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type PutBucketRequestPaymentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketRequestPaymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketRequestPaymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketRequestPayment = "PutBucketRequestPayment"

// PutBucketRequestPaymentRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets the request payment configuration for a bucket. By default, the bucket
// owner pays for downloads from the bucket. This configuration parameter enables
// the bucket owner (only) to specify that the person requesting the download
// will be charged for the download. For more information, see Requester Pays
// Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html).
//
// The following operations are related to PutBucketRequestPayment:
//
//    * CreateBucket
//
//    * GetBucketRequestPayment
//
//    // Example sending a request using PutBucketRequestPaymentRequest.
//    req := client.PutBucketRequestPaymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketRequestPayment
func (c *Client) PutBucketRequestPaymentRequest(input *PutBucketRequestPaymentInput) PutBucketRequestPaymentRequest {
	op := &aws.Operation{
		Name:       opPutBucketRequestPayment,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?requestPayment",
	}

	if input == nil {
		input = &PutBucketRequestPaymentInput{}
	}

	req := c.newRequest(op, input, &PutBucketRequestPaymentOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	req.Handlers.Build.PushBackNamed(aws.NamedHandler{
		Name: "contentMd5Handler",
		Fn:   checksum.AddBodyContentMD5Handler,
	})

	return PutBucketRequestPaymentRequest{Request: req, Input: input, Copy: c.PutBucketRequestPaymentRequest}
}

// PutBucketRequestPaymentRequest is the request type for the
// PutBucketRequestPayment API operation.
type PutBucketRequestPaymentRequest struct {
	*aws.Request
	Input *PutBucketRequestPaymentInput
	Copy  func(*PutBucketRequestPaymentInput) PutBucketRequestPaymentRequest
}

// Send marshals and sends the PutBucketRequestPayment API request.
func (r PutBucketRequestPaymentRequest) Send(ctx context.Context) (*PutBucketRequestPaymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketRequestPaymentResponse{
		PutBucketRequestPaymentOutput: r.Request.Data.(*PutBucketRequestPaymentOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketRequestPaymentResponse is the response type for the
// PutBucketRequestPayment API operation.
type PutBucketRequestPaymentResponse struct {
	*PutBucketRequestPaymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketRequestPayment request.
func (r *PutBucketRequestPaymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
