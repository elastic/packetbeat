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

type PutBucketLoggingInput struct {
	_ struct{} `type:"structure" payload:"BucketLoggingStatus"`

	// The name of the bucket for which to set the logging parameters.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Container for logging status information.
	//
	// BucketLoggingStatus is a required field
	BucketLoggingStatus *BucketLoggingStatus `locationName:"BucketLoggingStatus" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketLoggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.BucketLoggingStatus == nil {
		invalidParams.Add(aws.NewErrParamRequired("BucketLoggingStatus"))
	}
	if s.BucketLoggingStatus != nil {
		if err := s.BucketLoggingStatus.Validate(); err != nil {
			invalidParams.AddNested("BucketLoggingStatus", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketLoggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketLoggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.BucketLoggingStatus != nil {
		v := s.BucketLoggingStatus

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "BucketLoggingStatus", v, metadata)
	}
	return nil
}

func (s *PutBucketLoggingInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *PutBucketLoggingInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type PutBucketLoggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketLoggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketLoggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketLogging = "PutBucketLogging"

// PutBucketLoggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Set the logging parameters for a bucket and to specify permissions for who
// can view and modify the logging parameters. All logs are saved to buckets
// in the same AWS Region as the source bucket. To set the logging status of
// a bucket, you must be the bucket owner.
//
// The bucket owner is automatically granted FULL_CONTROL to all logs. You use
// the Grantee request element to grant access to other people. The Permissions
// request element specifies the kind of access the grantee has to the logs.
//
// Grantee Values
//
// You can specify the person (grantee) to whom you're assigning access rights
// (using request elements) in the following ways:
//
//    * By the person's ID: <Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//    xsi:type="CanonicalUser"><ID><>ID<></ID><DisplayName><>GranteesEmail<></DisplayName>
//    </Grantee> DisplayName is optional and ignored in the request.
//
//    * By Email address: <Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//    xsi:type="AmazonCustomerByEmail"><EmailAddress><>Grantees@email.com<></EmailAddress></Grantee>
//    The grantee is resolved to the CanonicalUser and, in a response to a GET
//    Object acl request, appears as the CanonicalUser.
//
//    * By URI: <Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//    xsi:type="Group"><URI><>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<></URI></Grantee>
//
// To enable logging, you use LoggingEnabled and its children request elements.
// To disable logging, you use an empty BucketLoggingStatus request element:
//
// <BucketLoggingStatus xmlns="http://doc.s3.amazonaws.com/2006-03-01" />
//
// For more information about server access logging, see Server Access Logging
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html).
//
// For more information about creating a bucket, see CreateBucket. For more
// information about returning the logging status of a bucket, see GetBucketLogging.
//
// The following operations are related to PutBucketLogging:
//
//    * PutObject
//
//    * DeleteBucket
//
//    * CreateBucket
//
//    * GetBucketLogging
//
//    // Example sending a request using PutBucketLoggingRequest.
//    req := client.PutBucketLoggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketLogging
func (c *Client) PutBucketLoggingRequest(input *PutBucketLoggingInput) PutBucketLoggingRequest {
	op := &aws.Operation{
		Name:       opPutBucketLogging,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?logging",
	}

	if input == nil {
		input = &PutBucketLoggingInput{}
	}

	req := c.newRequest(op, input, &PutBucketLoggingOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	req.Handlers.Build.PushBackNamed(aws.NamedHandler{
		Name: "contentMd5Handler",
		Fn:   checksum.AddBodyContentMD5Handler,
	})

	return PutBucketLoggingRequest{Request: req, Input: input, Copy: c.PutBucketLoggingRequest}
}

// PutBucketLoggingRequest is the request type for the
// PutBucketLogging API operation.
type PutBucketLoggingRequest struct {
	*aws.Request
	Input *PutBucketLoggingInput
	Copy  func(*PutBucketLoggingInput) PutBucketLoggingRequest
}

// Send marshals and sends the PutBucketLogging API request.
func (r PutBucketLoggingRequest) Send(ctx context.Context) (*PutBucketLoggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketLoggingResponse{
		PutBucketLoggingOutput: r.Request.Data.(*PutBucketLoggingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketLoggingResponse is the response type for the
// PutBucketLogging API operation.
type PutBucketLoggingResponse struct {
	*PutBucketLoggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketLogging request.
func (r *PutBucketLoggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
