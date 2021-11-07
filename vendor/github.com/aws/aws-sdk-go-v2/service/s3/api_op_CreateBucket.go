// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateBucketInput struct {
	_ struct{} `type:"structure" payload:"CreateBucketConfiguration"`

	// The canned ACL to apply to the bucket.
	ACL BucketCannedACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// The name of the bucket to create.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The configuration information for the bucket.
	CreateBucketConfiguration *CreateBucketConfiguration `locationName:"CreateBucketConfiguration" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to list the objects in the bucket.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to create, overwrite, and delete any object in the bucket.
	GrantWrite *string `location:"header" locationName:"x-amz-grant-write" type:"string"`

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	ObjectLockEnabledForBucket *bool `location:"header" locationName:"x-amz-bucket-object-lock-enabled" type:"boolean"`
}

// String returns the string representation
func (s CreateBucketInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBucketInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBucketInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *CreateBucketInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBucketInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ACL) > 0 {
		v := s.ACL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-acl", v, metadata)
	}
	if s.GrantFullControl != nil {
		v := *s.GrantFullControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-full-control", protocol.StringValue(v), metadata)
	}
	if s.GrantRead != nil {
		v := *s.GrantRead

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read", protocol.StringValue(v), metadata)
	}
	if s.GrantReadACP != nil {
		v := *s.GrantReadACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read-acp", protocol.StringValue(v), metadata)
	}
	if s.GrantWrite != nil {
		v := *s.GrantWrite

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write", protocol.StringValue(v), metadata)
	}
	if s.GrantWriteACP != nil {
		v := *s.GrantWriteACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write-acp", protocol.StringValue(v), metadata)
	}
	if s.ObjectLockEnabledForBucket != nil {
		v := *s.ObjectLockEnabledForBucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bucket-object-lock-enabled", protocol.BoolValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.CreateBucketConfiguration != nil {
		v := s.CreateBucketConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "CreateBucketConfiguration", v, metadata)
	}
	return nil
}

type CreateBucketOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the Region where the bucket will be created. If you are creating
	// a bucket on the US East (N. Virginia) Region (us-east-1), you do not need
	// to specify the location.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s CreateBucketOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBucketOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCreateBucket = "CreateBucket"

// CreateBucketRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a new bucket. To create a bucket, you must register with Amazon S3
// and have a valid AWS Access Key ID to authenticate requests. Anonymous requests
// are never allowed to create buckets. By creating the bucket, you become the
// bucket owner.
//
// Not every string is an acceptable bucket name. For information on bucket
// naming restrictions, see Working with Amazon S3 Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html).
//
// By default, the bucket is created in the US East (N. Virginia) Region. You
// can optionally specify a Region in the request body. You might choose a Region
// to optimize latency, minimize costs, or address regulatory requirements.
// For example, if you reside in Europe, you will probably find it advantageous
// to create buckets in the Europe (Ireland) Region. For more information, see
// How to Select a Region for Your Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro).
//
// If you send your create bucket request to the s3.amazonaws.com endpoint,
// the request goes to the us-east-1 Region. Accordingly, the signature calculations
// in Signature Version 4 must use us-east-1 as the Region, even if the location
// constraint in the request specifies another Region where the bucket is to
// be created. If you create a bucket in a Region other than US East (N. Virginia),
// your application must be able to handle 307 redirect. For more information,
// see Virtual Hosting of Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html).
//
// When creating a bucket using this operation, you can optionally specify the
// accounts or groups that should be granted specific permissions on the bucket.
// There are two ways to grant the appropriate permissions using the request
// headers.
//
//    * Specify a canned ACL using the x-amz-acl request header. Amazon S3 supports
//    a set of predefined ACLs, known as canned ACLs. Each canned ACL has a
//    predefined set of grantees and permissions. For more information, see
//    Canned ACL (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
//    * Specify access permissions explicitly using the x-amz-grant-read, x-amz-grant-write,
//    x-amz-grant-read-acp, x-amz-grant-write-acp, and x-amz-grant-full-control
//    headers. These headers map to the set of permissions Amazon S3 supports
//    in an ACL. For more information, see Access Control List (ACL) Overview
//    (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html). You
//    specify each grantee as a type=value pair, where the type is one of the
//    following: id – if the value specified is the canonical user ID of an
//    AWS account uri – if you are granting permissions to a predefined group
//    emailAddress – if the value specified is the email address of an AWS
//    account Using email addresses to specify a grantee is only supported in
//    the following AWS Regions: US East (N. Virginia) US West (N. California)
//    US West (Oregon) Asia Pacific (Singapore) Asia Pacific (Sydney) Asia Pacific
//    (Tokyo) Europe (Ireland) South America (São Paulo) For a list of all
//    the Amazon S3 supported Regions and endpoints, see Regions and Endpoints
//    (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) in
//    the AWS General Reference. For example, the following x-amz-grant-read
//    header grants the AWS accounts identified by account IDs permissions to
//    read object data and its metadata: x-amz-grant-read: id="11112222333",
//    id="444455556666"
//
// You can use either a canned ACL or specify access permissions explicitly.
// You cannot do both.
//
// The following operations are related to CreateBucket:
//
//    * PutObject
//
//    * DeleteBucket
//
//    // Example sending a request using CreateBucketRequest.
//    req := client.CreateBucketRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CreateBucket
func (c *Client) CreateBucketRequest(input *CreateBucketInput) CreateBucketRequest {
	op := &aws.Operation{
		Name:       opCreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &CreateBucketInput{}
	}

	req := c.newRequest(op, input, &CreateBucketOutput{})

	return CreateBucketRequest{Request: req, Input: input, Copy: c.CreateBucketRequest}
}

// CreateBucketRequest is the request type for the
// CreateBucket API operation.
type CreateBucketRequest struct {
	*aws.Request
	Input *CreateBucketInput
	Copy  func(*CreateBucketInput) CreateBucketRequest
}

// Send marshals and sends the CreateBucket API request.
func (r CreateBucketRequest) Send(ctx context.Context) (*CreateBucketResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBucketResponse{
		CreateBucketOutput: r.Request.Data.(*CreateBucketOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBucketResponse is the response type for the
// CreateBucket API operation.
type CreateBucketResponse struct {
	*CreateBucketOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBucket request.
func (r *CreateBucketResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
