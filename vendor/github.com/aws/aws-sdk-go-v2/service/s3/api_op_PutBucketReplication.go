// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketReplicationRequest
type PutBucketReplicationInput struct {
	_ struct{} `type:"structure" payload:"ReplicationConfiguration"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	//
	// ReplicationConfiguration is a required field
	ReplicationConfiguration *ReplicationConfiguration `locationName:"ReplicationConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	Token *string `location:"header" locationName:"x-amz-bucket-object-lock-token" type:"string"`
}

// String returns the string representation
func (s PutBucketReplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketReplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketReplicationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.ReplicationConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationConfiguration"))
	}
	if s.ReplicationConfiguration != nil {
		if err := s.ReplicationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ReplicationConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketReplicationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketReplicationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bucket-object-lock-token", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ReplicationConfiguration != nil {
		v := s.ReplicationConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "ReplicationConfiguration", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketReplicationOutput
type PutBucketReplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketReplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketReplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketReplication = "PutBucketReplication"

// PutBucketReplicationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a replication configuration or replaces an existing one. For more
// information, see Cross-Region Replication (CRR) (https://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html)
// in the Amazon S3 Developer Guide.
//
//    // Example sending a request using PutBucketReplicationRequest.
//    req := client.PutBucketReplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketReplication
func (c *Client) PutBucketReplicationRequest(input *PutBucketReplicationInput) PutBucketReplicationRequest {
	op := &aws.Operation{
		Name:       opPutBucketReplication,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?replication",
	}

	if input == nil {
		input = &PutBucketReplicationInput{}
	}

	req := c.newRequest(op, input, &PutBucketReplicationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketReplicationRequest{Request: req, Input: input, Copy: c.PutBucketReplicationRequest}
}

// PutBucketReplicationRequest is the request type for the
// PutBucketReplication API operation.
type PutBucketReplicationRequest struct {
	*aws.Request
	Input *PutBucketReplicationInput
	Copy  func(*PutBucketReplicationInput) PutBucketReplicationRequest
}

// Send marshals and sends the PutBucketReplication API request.
func (r PutBucketReplicationRequest) Send(ctx context.Context) (*PutBucketReplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketReplicationResponse{
		PutBucketReplicationOutput: r.Request.Data.(*PutBucketReplicationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketReplicationResponse is the response type for the
// PutBucketReplication API operation.
type PutBucketReplicationResponse struct {
	*PutBucketReplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketReplication request.
func (r *PutBucketReplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
