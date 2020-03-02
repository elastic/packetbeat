package azblob

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// appendBlobClient is the client for the AppendBlob methods of the Azblob service.
type appendBlobClient struct {
	managementClient
}

// newAppendBlobClient creates an instance of the appendBlobClient client.
func newAppendBlobClient(url url.URL, p pipeline.Pipeline) appendBlobClient {
	return appendBlobClient{newManagementClient(url, p)}
}

// AppendBlock the Append Block operation commits a new block of data to the end of an existing append blob. The Append
// Block operation is permitted only if the blob was created with x-ms-blob-type set to AppendBlob. Append Block is
// supported only on version 2015-02-21 version or later.
//
// body is initial data body will be closed upon successful return. Callers should ensure closure when receiving an
// error.contentLength is the length of the request. timeout is the timeout parameter is expressed in seconds. For more
// information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> transactionalContentMD5 is specify the transactional md5 for the body, to
// be validated by the service. leaseID is if specified, the operation only succeeds if the resource's lease is active
// and matches this ID. maxSize is optional conditional header. The max length in bytes permitted for the append blob.
// If the Append Block operation would cause the blob to exceed that limit or if the blob size is already greater than
// the value specified in this header, the request will fail with MaxBlobSizeConditionNotMet error (HTTP status code
// 412 - Precondition Failed). appendPosition is optional conditional header, used only for the Append Block operation.
// A number indicating the byte offset to compare. Append Block will succeed only if the append position is equal to
// this number. If it is not, the request will fail with the AppendPositionConditionNotMet error (HTTP status code 412
// - Precondition Failed). ifModifiedSince is specify this header value to operate only on a blob if it has been
// modified since the specified date/time. ifUnmodifiedSince is specify this header value to operate only on a blob if
// it has not been modified since the specified date/time. ifMatch is specify an ETag value to operate only on blobs
// with a matching value. ifNoneMatch is specify an ETag value to operate only on blobs without a matching value.
// requestID is provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics
// logs when storage analytics logging is enabled.
func (client appendBlobClient) AppendBlock(ctx context.Context, body io.ReadSeeker, contentLength int64, timeout *int32, transactionalContentMD5 []byte, leaseID *string, maxSize *int64, appendPosition *int64, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatch *ETag, ifNoneMatch *ETag, requestID *string) (*AppendBlobAppendBlockResponse, error) {
	if err := validate([]validation{
		{targetValue: body,
			constraints: []constraint{{target: "body", name: null, rule: true, chain: nil}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.appendBlockPreparer(body, contentLength, timeout, transactionalContentMD5, leaseID, maxSize, appendPosition, ifModifiedSince, ifUnmodifiedSince, ifMatch, ifNoneMatch, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.appendBlockResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*AppendBlobAppendBlockResponse), err
}

// appendBlockPreparer prepares the AppendBlock request.
func (client appendBlobClient) appendBlockPreparer(body io.ReadSeeker, contentLength int64, timeout *int32, transactionalContentMD5 []byte, leaseID *string, maxSize *int64, appendPosition *int64, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatch *ETag, ifNoneMatch *ETag, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "appendblock")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if transactionalContentMD5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(transactionalContentMD5))
	}
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	if maxSize != nil {
		req.Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*maxSize, 10))
	}
	if appendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPosition, 10))
	}
	if ifModifiedSince != nil {
		req.Header.Set("If-Modified-Since", (*ifModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", (*ifUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifMatch != nil {
		req.Header.Set("If-Match", string(*ifMatch))
	}
	if ifNoneMatch != nil {
		req.Header.Set("If-None-Match", string(*ifNoneMatch))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// appendBlockResponder handles the response to the AppendBlock request.
func (client appendBlobClient) appendBlockResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &AppendBlobAppendBlockResponse{rawResponse: resp.Response()}, err
}

// AppendBlockFromURL the Append Block operation commits a new block of data to the end of an existing append blob
// where the contents are read from a source url. The Append Block operation is permitted only if the blob was created
// with x-ms-blob-type set to AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
//
// sourceURL is specify a URL to the copy source. contentLength is the length of the request. sourceRange is bytes of
// source data in the specified range. sourceContentMD5 is specify the md5 calculated for the range of bytes that must
// be read from the copy source. timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> leaseID is if specified, the operation only succeeds if the resource's
// lease is active and matches this ID. maxSize is optional conditional header. The max length in bytes permitted for
// the append blob. If the Append Block operation would cause the blob to exceed that limit or if the blob size is
// already greater than the value specified in this header, the request will fail with MaxBlobSizeConditionNotMet error
// (HTTP status code 412 - Precondition Failed). appendPosition is optional conditional header, used only for the
// Append Block operation. A number indicating the byte offset to compare. Append Block will succeed only if the append
// position is equal to this number. If it is not, the request will fail with the AppendPositionConditionNotMet error
// (HTTP status code 412 - Precondition Failed). ifModifiedSince is specify this header value to operate only on a blob
// if it has been modified since the specified date/time. ifUnmodifiedSince is specify this header value to operate
// only on a blob if it has not been modified since the specified date/time. ifMatch is specify an ETag value to
// operate only on blobs with a matching value. ifNoneMatch is specify an ETag value to operate only on blobs without a
// matching value. sourceIfModifiedSince is specify this header value to operate only on a blob if it has been modified
// since the specified date/time. sourceIfUnmodifiedSince is specify this header value to operate only on a blob if it
// has not been modified since the specified date/time. sourceIfMatch is specify an ETag value to operate only on blobs
// with a matching value. sourceIfNoneMatch is specify an ETag value to operate only on blobs without a matching value.
// requestID is provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics
// logs when storage analytics logging is enabled.
func (client appendBlobClient) AppendBlockFromURL(ctx context.Context, sourceURL string, contentLength int64, sourceRange *string, sourceContentMD5 []byte, timeout *int32, leaseID *string, maxSize *int64, appendPosition *int64, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatch *ETag, ifNoneMatch *ETag, sourceIfModifiedSince *time.Time, sourceIfUnmodifiedSince *time.Time, sourceIfMatch *ETag, sourceIfNoneMatch *ETag, requestID *string) (*AppendBlobAppendBlockFromURLResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.appendBlockFromURLPreparer(sourceURL, contentLength, sourceRange, sourceContentMD5, timeout, leaseID, maxSize, appendPosition, ifModifiedSince, ifUnmodifiedSince, ifMatch, ifNoneMatch, sourceIfModifiedSince, sourceIfUnmodifiedSince, sourceIfMatch, sourceIfNoneMatch, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.appendBlockFromURLResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*AppendBlobAppendBlockFromURLResponse), err
}

// appendBlockFromURLPreparer prepares the AppendBlockFromURL request.
func (client appendBlobClient) appendBlockFromURLPreparer(sourceURL string, contentLength int64, sourceRange *string, sourceContentMD5 []byte, timeout *int32, leaseID *string, maxSize *int64, appendPosition *int64, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatch *ETag, ifNoneMatch *ETag, sourceIfModifiedSince *time.Time, sourceIfUnmodifiedSince *time.Time, sourceIfMatch *ETag, sourceIfNoneMatch *ETag, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "appendblock")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("x-ms-copy-source", sourceURL)
	if sourceRange != nil {
		req.Header.Set("x-ms-source-range", *sourceRange)
	}
	if sourceContentMD5 != nil {
		req.Header.Set("x-ms-source-content-md5", base64.StdEncoding.EncodeToString(sourceContentMD5))
	}
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	if maxSize != nil {
		req.Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*maxSize, 10))
	}
	if appendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPosition, 10))
	}
	if ifModifiedSince != nil {
		req.Header.Set("If-Modified-Since", (*ifModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", (*ifUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifMatch != nil {
		req.Header.Set("If-Match", string(*ifMatch))
	}
	if ifNoneMatch != nil {
		req.Header.Set("If-None-Match", string(*ifNoneMatch))
	}
	if sourceIfModifiedSince != nil {
		req.Header.Set("x-ms-source-if-modified-since", (*sourceIfModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if sourceIfUnmodifiedSince != nil {
		req.Header.Set("x-ms-source-if-unmodified-since", (*sourceIfUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if sourceIfMatch != nil {
		req.Header.Set("x-ms-source-if-match", string(*sourceIfMatch))
	}
	if sourceIfNoneMatch != nil {
		req.Header.Set("x-ms-source-if-none-match", string(*sourceIfNoneMatch))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// appendBlockFromURLResponder handles the response to the AppendBlockFromURL request.
func (client appendBlobClient) appendBlockFromURLResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &AppendBlobAppendBlockFromURLResponse{rawResponse: resp.Response()}, err
}

// Create the Create Append Blob operation creates a new append blob.
//
// contentLength is the length of the request. timeout is the timeout parameter is expressed in seconds. For more
// information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> blobContentType is optional. Sets the blob's content type. If specified,
// this property is stored with the blob and returned with a read request. blobContentEncoding is optional. Sets the
// blob's content encoding. If specified, this property is stored with the blob and returned with a read request.
// blobContentLanguage is optional. Set the blob's content language. If specified, this property is stored with the
// blob and returned with a read request. blobContentMD5 is optional. An MD5 hash of the blob content. Note that this
// hash is not validated, as the hashes for the individual blocks were validated when each was uploaded.
// blobCacheControl is optional. Sets the blob's cache control. If specified, this property is stored with the blob and
// returned with a read request. metadata is optional. Specifies a user-defined name-value pair associated with the
// blob. If no name-value pairs are specified, the operation will copy the metadata from the source blob or file to the
// destination blob. If one or more name-value pairs are specified, the destination blob is created with the specified
// metadata, and metadata is not copied from the source blob or file. Note that beginning with version 2009-09-19,
// metadata names must adhere to the naming rules for C# identifiers. See Naming and Referencing Containers, Blobs, and
// Metadata for more information. leaseID is if specified, the operation only succeeds if the resource's lease is
// active and matches this ID. blobContentDisposition is optional. Sets the blob's Content-Disposition header.
// ifModifiedSince is specify this header value to operate only on a blob if it has been modified since the specified
// date/time. ifUnmodifiedSince is specify this header value to operate only on a blob if it has not been modified
// since the specified date/time. ifMatch is specify an ETag value to operate only on blobs with a matching value.
// ifNoneMatch is specify an ETag value to operate only on blobs without a matching value. requestID is provides a
// client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when storage
// analytics logging is enabled.
func (client appendBlobClient) Create(ctx context.Context, contentLength int64, timeout *int32, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 []byte, blobCacheControl *string, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatch *ETag, ifNoneMatch *ETag, requestID *string) (*AppendBlobCreateResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPreparer(contentLength, timeout, blobContentType, blobContentEncoding, blobContentLanguage, blobContentMD5, blobCacheControl, metadata, leaseID, blobContentDisposition, ifModifiedSince, ifUnmodifiedSince, ifMatch, ifNoneMatch, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*AppendBlobCreateResponse), err
}

// createPreparer prepares the Create request.
func (client appendBlobClient) createPreparer(contentLength int64, timeout *int32, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 []byte, blobCacheControl *string, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatch *ETag, ifNoneMatch *ETag, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if blobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobContentType)
	}
	if blobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobContentEncoding)
	}
	if blobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobContentLanguage)
	}
	if blobContentMD5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(blobContentMD5))
	}
	if blobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobCacheControl)
	}
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	if blobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobContentDisposition)
	}
	if ifModifiedSince != nil {
		req.Header.Set("If-Modified-Since", (*ifModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", (*ifUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifMatch != nil {
		req.Header.Set("If-Match", string(*ifMatch))
	}
	if ifNoneMatch != nil {
		req.Header.Set("If-None-Match", string(*ifNoneMatch))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	req.Header.Set("x-ms-blob-type", "AppendBlob")
	return req, nil
}

// createResponder handles the response to the Create request.
func (client appendBlobClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &AppendBlobCreateResponse{rawResponse: resp.Response()}, err
}
