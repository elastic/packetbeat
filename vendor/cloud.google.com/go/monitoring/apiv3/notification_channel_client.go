// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newNotificationChannelClientHook clientHook

// NotificationChannelCallOptions contains the retry settings for each method of NotificationChannelClient.
type NotificationChannelCallOptions struct {
	ListNotificationChannelDescriptors      []gax.CallOption
	GetNotificationChannelDescriptor        []gax.CallOption
	ListNotificationChannels                []gax.CallOption
	GetNotificationChannel                  []gax.CallOption
	CreateNotificationChannel               []gax.CallOption
	UpdateNotificationChannel               []gax.CallOption
	DeleteNotificationChannel               []gax.CallOption
	SendNotificationChannelVerificationCode []gax.CallOption
	GetNotificationChannelVerificationCode  []gax.CallOption
	VerifyNotificationChannel               []gax.CallOption
}

func defaultNotificationChannelClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("monitoring.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultNotificationChannelCallOptions() *NotificationChannelCallOptions {
	return &NotificationChannelCallOptions{
		ListNotificationChannelDescriptors: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetNotificationChannelDescriptor: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListNotificationChannels: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetNotificationChannel: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateNotificationChannel: []gax.CallOption{},
		UpdateNotificationChannel: []gax.CallOption{},
		DeleteNotificationChannel: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SendNotificationChannelVerificationCode: []gax.CallOption{},
		GetNotificationChannelVerificationCode: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		VerifyNotificationChannel: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// NotificationChannelClient is a client for interacting with Cloud Monitoring API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type NotificationChannelClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	notificationChannelClient monitoringpb.NotificationChannelServiceClient

	// The call options for this service.
	CallOptions *NotificationChannelCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewNotificationChannelClient creates a new notification channel service client.
//
// The Notification Channel API provides access to configuration that
// controls how messages related to incidents are sent.
func NewNotificationChannelClient(ctx context.Context, opts ...option.ClientOption) (*NotificationChannelClient, error) {
	clientOpts := defaultNotificationChannelClientOptions()

	if newNotificationChannelClientHook != nil {
		hookOpts, err := newNotificationChannelClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &NotificationChannelClient{
		connPool:    connPool,
		CallOptions: defaultNotificationChannelCallOptions(),

		notificationChannelClient: monitoringpb.NewNotificationChannelServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *NotificationChannelClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *NotificationChannelClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *NotificationChannelClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListNotificationChannelDescriptors lists the descriptors for supported channel types. The use of descriptors
// makes it possible for new channel types to be dynamically added.
func (c *NotificationChannelClient) ListNotificationChannelDescriptors(ctx context.Context, req *monitoringpb.ListNotificationChannelDescriptorsRequest, opts ...gax.CallOption) *NotificationChannelDescriptorIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListNotificationChannelDescriptors[0:len(c.CallOptions.ListNotificationChannelDescriptors):len(c.CallOptions.ListNotificationChannelDescriptors)], opts...)
	it := &NotificationChannelDescriptorIterator{}
	req = proto.Clone(req).(*monitoringpb.ListNotificationChannelDescriptorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.NotificationChannelDescriptor, string, error) {
		var resp *monitoringpb.ListNotificationChannelDescriptorsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.notificationChannelClient.ListNotificationChannelDescriptors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.ChannelDescriptors, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// GetNotificationChannelDescriptor gets a single channel descriptor. The descriptor indicates which fields
// are expected / permitted for a notification channel of the given type.
func (c *NotificationChannelClient) GetNotificationChannelDescriptor(ctx context.Context, req *monitoringpb.GetNotificationChannelDescriptorRequest, opts ...gax.CallOption) (*monitoringpb.NotificationChannelDescriptor, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetNotificationChannelDescriptor[0:len(c.CallOptions.GetNotificationChannelDescriptor):len(c.CallOptions.GetNotificationChannelDescriptor)], opts...)
	var resp *monitoringpb.NotificationChannelDescriptor
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.notificationChannelClient.GetNotificationChannelDescriptor(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListNotificationChannels lists the notification channels that have been created for the project.
func (c *NotificationChannelClient) ListNotificationChannels(ctx context.Context, req *monitoringpb.ListNotificationChannelsRequest, opts ...gax.CallOption) *NotificationChannelIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListNotificationChannels[0:len(c.CallOptions.ListNotificationChannels):len(c.CallOptions.ListNotificationChannels)], opts...)
	it := &NotificationChannelIterator{}
	req = proto.Clone(req).(*monitoringpb.ListNotificationChannelsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.NotificationChannel, string, error) {
		var resp *monitoringpb.ListNotificationChannelsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.notificationChannelClient.ListNotificationChannels(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.NotificationChannels, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// GetNotificationChannel gets a single notification channel. The channel includes the relevant
// configuration details with which the channel was created. However, the
// response may truncate or omit passwords, API keys, or other private key
// matter and thus the response may not be 100% identical to the information
// that was supplied in the call to the create method.
func (c *NotificationChannelClient) GetNotificationChannel(ctx context.Context, req *monitoringpb.GetNotificationChannelRequest, opts ...gax.CallOption) (*monitoringpb.NotificationChannel, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetNotificationChannel[0:len(c.CallOptions.GetNotificationChannel):len(c.CallOptions.GetNotificationChannel)], opts...)
	var resp *monitoringpb.NotificationChannel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.notificationChannelClient.GetNotificationChannel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateNotificationChannel creates a new notification channel, representing a single notification
// endpoint such as an email address, SMS number, or PagerDuty service.
func (c *NotificationChannelClient) CreateNotificationChannel(ctx context.Context, req *monitoringpb.CreateNotificationChannelRequest, opts ...gax.CallOption) (*monitoringpb.NotificationChannel, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateNotificationChannel[0:len(c.CallOptions.CreateNotificationChannel):len(c.CallOptions.CreateNotificationChannel)], opts...)
	var resp *monitoringpb.NotificationChannel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.notificationChannelClient.CreateNotificationChannel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateNotificationChannel updates a notification channel. Fields not specified in the field mask
// remain unchanged.
func (c *NotificationChannelClient) UpdateNotificationChannel(ctx context.Context, req *monitoringpb.UpdateNotificationChannelRequest, opts ...gax.CallOption) (*monitoringpb.NotificationChannel, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "notification_channel.name", url.QueryEscape(req.GetNotificationChannel().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateNotificationChannel[0:len(c.CallOptions.UpdateNotificationChannel):len(c.CallOptions.UpdateNotificationChannel)], opts...)
	var resp *monitoringpb.NotificationChannel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.notificationChannelClient.UpdateNotificationChannel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteNotificationChannel deletes a notification channel.
func (c *NotificationChannelClient) DeleteNotificationChannel(ctx context.Context, req *monitoringpb.DeleteNotificationChannelRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteNotificationChannel[0:len(c.CallOptions.DeleteNotificationChannel):len(c.CallOptions.DeleteNotificationChannel)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.notificationChannelClient.DeleteNotificationChannel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SendNotificationChannelVerificationCode causes a verification code to be delivered to the channel. The code
// can then be supplied in VerifyNotificationChannel to verify the channel.
func (c *NotificationChannelClient) SendNotificationChannelVerificationCode(ctx context.Context, req *monitoringpb.SendNotificationChannelVerificationCodeRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SendNotificationChannelVerificationCode[0:len(c.CallOptions.SendNotificationChannelVerificationCode):len(c.CallOptions.SendNotificationChannelVerificationCode)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.notificationChannelClient.SendNotificationChannelVerificationCode(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// GetNotificationChannelVerificationCode requests a verification code for an already verified channel that can then
// be used in a call to VerifyNotificationChannel() on a different channel
// with an equivalent identity in the same or in a different project. This
// makes it possible to copy a channel between projects without requiring
// manual reverification of the channel. If the channel is not in the
// verified state, this method will fail (in other words, this may only be
// used if the SendNotificationChannelVerificationCode and
// VerifyNotificationChannel paths have already been used to put the given
// channel into the verified state).
//
// There is no guarantee that the verification codes returned by this method
// will be of a similar structure or form as the ones that are delivered
// to the channel via SendNotificationChannelVerificationCode; while
// VerifyNotificationChannel() will recognize both the codes delivered via
// SendNotificationChannelVerificationCode() and returned from
// GetNotificationChannelVerificationCode(), it is typically the case that
// the verification codes delivered via
// SendNotificationChannelVerificationCode() will be shorter and also
// have a shorter expiration (e.g. codes such as “G-123456”) whereas
// GetVerificationCode() will typically return a much longer, websafe base
// 64 encoded string that has a longer expiration time.
func (c *NotificationChannelClient) GetNotificationChannelVerificationCode(ctx context.Context, req *monitoringpb.GetNotificationChannelVerificationCodeRequest, opts ...gax.CallOption) (*monitoringpb.GetNotificationChannelVerificationCodeResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetNotificationChannelVerificationCode[0:len(c.CallOptions.GetNotificationChannelVerificationCode):len(c.CallOptions.GetNotificationChannelVerificationCode)], opts...)
	var resp *monitoringpb.GetNotificationChannelVerificationCodeResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.notificationChannelClient.GetNotificationChannelVerificationCode(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VerifyNotificationChannel verifies a NotificationChannel by proving receipt of the code
// delivered to the channel as a result of calling
// SendNotificationChannelVerificationCode.
func (c *NotificationChannelClient) VerifyNotificationChannel(ctx context.Context, req *monitoringpb.VerifyNotificationChannelRequest, opts ...gax.CallOption) (*monitoringpb.NotificationChannel, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.VerifyNotificationChannel[0:len(c.CallOptions.VerifyNotificationChannel):len(c.CallOptions.VerifyNotificationChannel)], opts...)
	var resp *monitoringpb.NotificationChannel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.notificationChannelClient.VerifyNotificationChannel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NotificationChannelDescriptorIterator manages a stream of *monitoringpb.NotificationChannelDescriptor.
type NotificationChannelDescriptorIterator struct {
	items    []*monitoringpb.NotificationChannelDescriptor
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.NotificationChannelDescriptor, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NotificationChannelDescriptorIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NotificationChannelDescriptorIterator) Next() (*monitoringpb.NotificationChannelDescriptor, error) {
	var item *monitoringpb.NotificationChannelDescriptor
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NotificationChannelDescriptorIterator) bufLen() int {
	return len(it.items)
}

func (it *NotificationChannelDescriptorIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// NotificationChannelIterator manages a stream of *monitoringpb.NotificationChannel.
type NotificationChannelIterator struct {
	items    []*monitoringpb.NotificationChannel
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.NotificationChannel, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NotificationChannelIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NotificationChannelIterator) Next() (*monitoringpb.NotificationChannel, error) {
	var item *monitoringpb.NotificationChannel
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NotificationChannelIterator) bufLen() int {
	return len(it.items)
}

func (it *NotificationChannelIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
