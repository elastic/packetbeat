// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type PutObjectInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The canned ACL to apply to the object. For more information, see Canned ACL
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
	ACL ObjectCannedACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// Object data.
	Body io.ReadSeeker `type:"blob"`

	// Bucket name to which the PUT operation was initiated.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Can be used to specify caching behavior along the request/reply chain. For
	// more information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9).
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1).
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field. For more information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11).
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// Size of the body in bytes. This parameter is useful when the size of the
	// body cannot be determined automatically. For more information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13).
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The base64-encoded 128-bit MD5 digest of the message (without the headers)
	// according to RFC 1864. This header can be used as a message integrity check
	// to verify that the data is the same data that was originally sent. Although
	// it is optional, we recommend using the Content-MD5 mechanism as an end-to-end
	// integrity check. For more information about REST request authentication,
	// see REST Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html).
	ContentMD5 *string `location:"header" locationName:"Content-MD5" type:"string"`

	// A standard MIME type describing the format of the contents. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The date and time at which the object is no longer cacheable. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21).
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to read the object data and its metadata.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the object ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Object key for which the PUT operation was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// Specifies whether a legal hold will be applied to this object. For more information
	// about S3 Object Lock, see Object Lock (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockLegalHoldStatus ObjectLockLegalHoldStatus `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"true"`

	// The Object Lock mode that you want to apply to this object.
	ObjectLockMode ObjectLockMode `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"true"`

	// The date and time when you want this object's Object Lock to expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// S3 does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// Specifies the AWS KMS Encryption Context to use for object encryption. The
	// value of this header is a base64-encoded UTF-8 string holding JSON with the
	// encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// If x-amz-server-side-encryption is present and has the value of aws:kms,
	// this header specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetrical customer managed customer master key (CMK) that was used for
	// the object.
	//
	// If the value of x-amz-server-side-encryption is aws:kms, this header specifies
	// the ID of the symmetric customer managed AWS KMS CMK that will be used for
	// the object. If you specify x-amz-server-side-encryption:aws:kms, but do not
	// providex-amz-server-side-encryption-aws-kms-key-id, Amazon S3 uses the AWS
	// managed CMK in AWS to protect the data.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The server-side encryption algorithm used when storing this object in Amazon
	// S3 (for example, AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// If you don't specify, S3 Standard is the default storage class. Amazon S3
	// supports other storage classes.
	StorageClass StorageClass `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"true"`

	// The tag-set for the object. The tag-set must be encoded as URL Query parameters.
	// (For example, "Key1=Value1")
	Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata. For information about object
	// metadata, see Object Key and Metadata (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html).
	//
	// In the following example, the request header sets the redirect to an object
	// (anotherPage.html) in the same bucket:
	//
	// x-amz-website-redirect-location: /anotherPage.html
	//
	// In the following example, the request header sets the object redirect to
	// another website:
	//
	// x-amz-website-redirect-location: http://www.example.com/
	//
	// For more information about website hosting in Amazon S3, see Hosting Websites
	// on Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html)
	// and How to Configure Website Page Redirects (https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s PutObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutObjectInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutObjectInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *PutObjectInput) getSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ACL) > 0 {
		v := s.ACL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-acl", v, metadata)
	}
	if s.CacheControl != nil {
		v := *s.CacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Cache-Control", protocol.StringValue(v), metadata)
	}
	if s.ContentDisposition != nil {
		v := *s.ContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Disposition", protocol.StringValue(v), metadata)
	}
	if s.ContentEncoding != nil {
		v := *s.ContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Encoding", protocol.StringValue(v), metadata)
	}
	if s.ContentLanguage != nil {
		v := *s.ContentLanguage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Language", protocol.StringValue(v), metadata)
	}
	if s.ContentLength != nil {
		v := *s.ContentLength

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Length", protocol.Int64Value(v), metadata)
	}
	if s.ContentMD5 != nil {
		v := *s.ContentMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-MD5", protocol.StringValue(v), metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue(v), metadata)
	}
	if s.Expires != nil {
		v := *s.Expires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Expires",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
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
	if s.GrantWriteACP != nil {
		v := *s.GrantWriteACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write-acp", protocol.StringValue(v), metadata)
	}
	if len(s.ObjectLockLegalHoldStatus) > 0 {
		v := s.ObjectLockLegalHoldStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-legal-hold", v, metadata)
	}
	if len(s.ObjectLockMode) > 0 {
		v := s.ObjectLockMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-mode", v, metadata)
	}
	if s.ObjectLockRetainUntilDate != nil {
		v := *s.ObjectLockRetainUntilDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-retain-until-date",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKey != nil {
		v := *s.SSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSEncryptionContext != nil {
		v := *s.SSEKMSEncryptionContext

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-context", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if len(s.StorageClass) > 0 {
		v := s.StorageClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-storage-class", v, metadata)
	}
	if s.Tagging != nil {
		v := *s.Tagging

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-tagging", protocol.StringValue(v), metadata)
	}
	if s.WebsiteRedirectLocation != nil {
		v := *s.WebsiteRedirectLocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-website-redirect-location", protocol.StringValue(v), metadata)
	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.HeadersTarget, "x-amz-meta-", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.StringValue(v1))
		}
		ms0.End()

	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Body", protocol.ReadSeekerStream{V: v}, metadata)
	}
	return nil
}

func (s *PutObjectInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *PutObjectInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type PutObjectOutput struct {
	_ struct{} `type:"structure"`

	// Entity tag for the uploaded object.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the expiration is configured for the object (see PutBucketLifecycleConfiguration),
	// the response includes this header. It includes the expiry-date and rule-id
	// key-value pairs that provide information about object expiration. The value
	// of the rule-id is URL encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the AWS KMS Encryption Context to use for object encryption.
	// The value of this header is a base64-encoded UTF-8 string holding JSON with
	// the encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// If x-amz-server-side-encryption is present and has the value of aws:kms,
	// this header specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// If you specified server-side encryption either with an AWS KMS customer master
	// key (CMK) or Amazon S3-managed encryption key in your PUT request, the response
	// includes this header. It confirms the encryption algorithm that Amazon S3
	// used to encrypt the object.
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s PutObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Expiration != nil {
		v := *s.Expiration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-expiration", protocol.StringValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSEncryptionContext != nil {
		v := *s.SSEKMSEncryptionContext

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-context", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	return nil
}

const opPutObject = "PutObject"

// PutObjectRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Adds an object to a bucket. You must have WRITE permissions on a bucket to
// add an object to it.
//
// Amazon S3 never adds partial objects; if you receive a success response,
// Amazon S3 added the entire object to the bucket.
//
// Amazon S3 is a distributed system. If it receives multiple write requests
// for the same object simultaneously, it overwrites all but the last object
// written. Amazon S3 does not provide object locking; if you need this, make
// sure to build it into your application layer or use versioning instead.
//
// To ensure that data is not corrupted traversing the network, use the Content-MD5
// header. When you use this header, Amazon S3 checks the object against the
// provided MD5 value and, if they do not match, returns an error. Additionally,
// you can calculate the MD5 while putting an object to Amazon S3 and compare
// the returned ETag to the calculated MD5 value.
//
// The Content-MD5 header is required for any request to upload an object with
// a retention period configured using Amazon S3 Object Lock. For more information
// about Amazon S3 Object Lock, see Amazon S3 Object Lock Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// Server-side Encryption
//
// You can optionally request server-side encryption. With server-side encryption,
// Amazon S3 encrypts your data as it writes it to disks in its data centers
// and decrypts the data when you access it. You have the option to provide
// your own encryption key or use AWS managed encryption keys. For more information,
// see Using Server-Side Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html).
//
// Access Control List (ACL)-Specific Request Headers
//
// You can use headers to grant ACL- based permissions. By default, all objects
// are private. Only the owner has full access control. When adding a new object,
// you can grant permissions to individual AWS accounts or to predefined groups
// defined by Amazon S3. These permissions are then added to the ACL on the
// object. For more information, see Access Control List (ACL) Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html)
// and Managing ACLs Using the REST API (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-using-rest-api.html).
//
// Storage Class Options
//
// By default, Amazon S3 uses the STANDARD storage class to store newly created
// objects. The STANDARD storage class provides high durability and high availability.
// Depending on performance needs, you can specify a different storage class.
// For more information, see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
// in the Amazon S3 Service Developer Guide.
//
// Versioning
//
// If you enable versioning for a bucket, Amazon S3 automatically generates
// a unique version ID for the object being stored. Amazon S3 returns this ID
// in the response. When you enable versioning for a bucket, if Amazon S3 receives
// multiple write requests for the same object simultaneously, it stores all
// of the objects.
//
// For more information about versioning, see Adding Objects to Versioning Enabled
// Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/AddingObjectstoVersioningEnabledBuckets.html).
// For information about returning the versioning state of a bucket, see GetBucketVersioning.
//
// Related Resources
//
//    * CopyObject
//
//    * DeleteObject
//
//    // Example sending a request using PutObjectRequest.
//    req := client.PutObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObject
func (c *Client) PutObjectRequest(input *PutObjectInput) PutObjectRequest {
	op := &aws.Operation{
		Name:       opPutObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &PutObjectInput{}
	}

	req := c.newRequest(op, input, &PutObjectOutput{})

	return PutObjectRequest{Request: req, Input: input, Copy: c.PutObjectRequest}
}

// PutObjectRequest is the request type for the
// PutObject API operation.
type PutObjectRequest struct {
	*aws.Request
	Input *PutObjectInput
	Copy  func(*PutObjectInput) PutObjectRequest
}

// Send marshals and sends the PutObject API request.
func (r PutObjectRequest) Send(ctx context.Context) (*PutObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectResponse{
		PutObjectOutput: r.Request.Data.(*PutObjectOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectResponse is the response type for the
// PutObject API operation.
type PutObjectResponse struct {
	*PutObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObject request.
func (r *PutObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
