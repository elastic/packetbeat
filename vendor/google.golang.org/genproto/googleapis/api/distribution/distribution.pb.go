// Copyright 2015 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/api/distribution.proto

package distribution

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `Distribution` contains summary statistics for a population of values. It
// optionally contains a histogram representing the distribution of those values
// across a set of buckets.
//
// The summary statistics are the count, mean, sum of the squared deviation from
// the mean, the minimum, and the maximum of the set of population of values.
// The histogram is based on a sequence of buckets and gives a count of values
// that fall into each bucket. The boundaries of the buckets are given either
// explicitly or by formulas for buckets of fixed or exponentially increasing
// widths.
//
// Although it is not forbidden, it is generally a bad idea to include
// non-finite values (infinities or NaNs) in the population of values, as this
// will render the `mean` and `sum_of_squared_deviation` fields meaningless.
type Distribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of values in the population. Must be non-negative. This value
	// must equal the sum of the values in `bucket_counts` if a histogram is
	// provided.
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// The arithmetic mean of the values in the population. If `count` is zero
	// then this field must be zero.
	Mean float64 `protobuf:"fixed64,2,opt,name=mean,proto3" json:"mean,omitempty"`
	// The sum of squared deviations from the mean of the values in the
	// population. For values x_i this is:
	//
	//     Sum[i=1..n]((x_i - mean)^2)
	//
	// Knuth, "The Art of Computer Programming", Vol. 2, page 232, 3rd edition
	// describes Welford's method for accumulating this sum in one pass.
	//
	// If `count` is zero then this field must be zero.
	SumOfSquaredDeviation float64 `protobuf:"fixed64,3,opt,name=sum_of_squared_deviation,json=sumOfSquaredDeviation,proto3" json:"sum_of_squared_deviation,omitempty"`
	// If specified, contains the range of the population values. The field
	// must not be present if the `count` is zero.
	Range *Distribution_Range `protobuf:"bytes,4,opt,name=range,proto3" json:"range,omitempty"`
	// Defines the histogram bucket boundaries. If the distribution does not
	// contain a histogram, then omit this field.
	BucketOptions *Distribution_BucketOptions `protobuf:"bytes,6,opt,name=bucket_options,json=bucketOptions,proto3" json:"bucket_options,omitempty"`
	// The number of values in each bucket of the histogram, as described in
	// `bucket_options`. If the distribution does not have a histogram, then omit
	// this field. If there is a histogram, then the sum of the values in
	// `bucket_counts` must equal the value in the `count` field of the
	// distribution.
	//
	// If present, `bucket_counts` should contain N values, where N is the number
	// of buckets specified in `bucket_options`. If you supply fewer than N
	// values, the remaining values are assumed to be 0.
	//
	// The order of the values in `bucket_counts` follows the bucket numbering
	// schemes described for the three bucket types. The first value must be the
	// count for the underflow bucket (number 0). The next N-2 values are the
	// counts for the finite buckets (number 1 through N-2). The N'th value in
	// `bucket_counts` is the count for the overflow bucket (number N-1).
	BucketCounts []int64 `protobuf:"varint,7,rep,packed,name=bucket_counts,json=bucketCounts,proto3" json:"bucket_counts,omitempty"`
	// Must be in increasing order of `value` field.
	Exemplars []*Distribution_Exemplar `protobuf:"bytes,10,rep,name=exemplars,proto3" json:"exemplars,omitempty"`
}

func (x *Distribution) Reset() {
	*x = Distribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_distribution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution) ProtoMessage() {}

func (x *Distribution) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_distribution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution.ProtoReflect.Descriptor instead.
func (*Distribution) Descriptor() ([]byte, []int) {
	return file_google_api_distribution_proto_rawDescGZIP(), []int{0}
}

func (x *Distribution) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Distribution) GetMean() float64 {
	if x != nil {
		return x.Mean
	}
	return 0
}

func (x *Distribution) GetSumOfSquaredDeviation() float64 {
	if x != nil {
		return x.SumOfSquaredDeviation
	}
	return 0
}

func (x *Distribution) GetRange() *Distribution_Range {
	if x != nil {
		return x.Range
	}
	return nil
}

func (x *Distribution) GetBucketOptions() *Distribution_BucketOptions {
	if x != nil {
		return x.BucketOptions
	}
	return nil
}

func (x *Distribution) GetBucketCounts() []int64 {
	if x != nil {
		return x.BucketCounts
	}
	return nil
}

func (x *Distribution) GetExemplars() []*Distribution_Exemplar {
	if x != nil {
		return x.Exemplars
	}
	return nil
}

// The range of the population values.
type Distribution_Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The minimum of the population values.
	Min float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	// The maximum of the population values.
	Max float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Distribution_Range) Reset() {
	*x = Distribution_Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_distribution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_Range) ProtoMessage() {}

func (x *Distribution_Range) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_distribution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_Range.ProtoReflect.Descriptor instead.
func (*Distribution_Range) Descriptor() ([]byte, []int) {
	return file_google_api_distribution_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Distribution_Range) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Distribution_Range) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

// `BucketOptions` describes the bucket boundaries used to create a histogram
// for the distribution. The buckets can be in a linear sequence, an
// exponential sequence, or each bucket can be specified explicitly.
// `BucketOptions` does not include the number of values in each bucket.
//
// A bucket has an inclusive lower bound and exclusive upper bound for the
// values that are counted for that bucket. The upper bound of a bucket must
// be strictly greater than the lower bound. The sequence of N buckets for a
// distribution consists of an underflow bucket (number 0), zero or more
// finite buckets (number 1 through N - 2) and an overflow bucket (number N -
// 1). The buckets are contiguous: the lower bound of bucket i (i > 0) is the
// same as the upper bound of bucket i - 1. The buckets span the whole range
// of finite values: lower bound of the underflow bucket is -infinity and the
// upper bound of the overflow bucket is +infinity. The finite buckets are
// so-called because both bounds are finite.
type Distribution_BucketOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Exactly one of these three fields must be set.
	//
	// Types that are assignable to Options:
	//	*Distribution_BucketOptions_LinearBuckets
	//	*Distribution_BucketOptions_ExponentialBuckets
	//	*Distribution_BucketOptions_ExplicitBuckets
	Options isDistribution_BucketOptions_Options `protobuf_oneof:"options"`
}

func (x *Distribution_BucketOptions) Reset() {
	*x = Distribution_BucketOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_distribution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_BucketOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_BucketOptions) ProtoMessage() {}

func (x *Distribution_BucketOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_distribution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_BucketOptions.ProtoReflect.Descriptor instead.
func (*Distribution_BucketOptions) Descriptor() ([]byte, []int) {
	return file_google_api_distribution_proto_rawDescGZIP(), []int{0, 1}
}

func (m *Distribution_BucketOptions) GetOptions() isDistribution_BucketOptions_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *Distribution_BucketOptions) GetLinearBuckets() *Distribution_BucketOptions_Linear {
	if x, ok := x.GetOptions().(*Distribution_BucketOptions_LinearBuckets); ok {
		return x.LinearBuckets
	}
	return nil
}

func (x *Distribution_BucketOptions) GetExponentialBuckets() *Distribution_BucketOptions_Exponential {
	if x, ok := x.GetOptions().(*Distribution_BucketOptions_ExponentialBuckets); ok {
		return x.ExponentialBuckets
	}
	return nil
}

func (x *Distribution_BucketOptions) GetExplicitBuckets() *Distribution_BucketOptions_Explicit {
	if x, ok := x.GetOptions().(*Distribution_BucketOptions_ExplicitBuckets); ok {
		return x.ExplicitBuckets
	}
	return nil
}

type isDistribution_BucketOptions_Options interface {
	isDistribution_BucketOptions_Options()
}

type Distribution_BucketOptions_LinearBuckets struct {
	// The linear bucket.
	LinearBuckets *Distribution_BucketOptions_Linear `protobuf:"bytes,1,opt,name=linear_buckets,json=linearBuckets,proto3,oneof"`
}

type Distribution_BucketOptions_ExponentialBuckets struct {
	// The exponential buckets.
	ExponentialBuckets *Distribution_BucketOptions_Exponential `protobuf:"bytes,2,opt,name=exponential_buckets,json=exponentialBuckets,proto3,oneof"`
}

type Distribution_BucketOptions_ExplicitBuckets struct {
	// The explicit buckets.
	ExplicitBuckets *Distribution_BucketOptions_Explicit `protobuf:"bytes,3,opt,name=explicit_buckets,json=explicitBuckets,proto3,oneof"`
}

func (*Distribution_BucketOptions_LinearBuckets) isDistribution_BucketOptions_Options() {}

func (*Distribution_BucketOptions_ExponentialBuckets) isDistribution_BucketOptions_Options() {}

func (*Distribution_BucketOptions_ExplicitBuckets) isDistribution_BucketOptions_Options() {}

// Exemplars are example points that may be used to annotate aggregated
// distribution values. They are metadata that gives information about a
// particular value added to a Distribution bucket, such as a trace ID that
// was active when a value was added. They may contain further information,
// such as a example values and timestamps, origin, etc.
type Distribution_Exemplar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value of the exemplar point. This value determines to which bucket the
	// exemplar belongs.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// The observation (sampling) time of the above value.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Contextual information about the example value. Examples are:
	//
	//   Trace: type.googleapis.com/google.monitoring.v3.SpanContext
	//
	//   Literal string: type.googleapis.com/google.protobuf.StringValue
	//
	//   Labels dropped during aggregation:
	//     type.googleapis.com/google.monitoring.v3.DroppedLabels
	//
	// There may be only a single attachment of any given message type in a
	// single exemplar, and this is enforced by the system.
	Attachments []*anypb.Any `protobuf:"bytes,3,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *Distribution_Exemplar) Reset() {
	*x = Distribution_Exemplar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_distribution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_Exemplar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_Exemplar) ProtoMessage() {}

func (x *Distribution_Exemplar) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_distribution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_Exemplar.ProtoReflect.Descriptor instead.
func (*Distribution_Exemplar) Descriptor() ([]byte, []int) {
	return file_google_api_distribution_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Distribution_Exemplar) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Distribution_Exemplar) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Distribution_Exemplar) GetAttachments() []*anypb.Any {
	if x != nil {
		return x.Attachments
	}
	return nil
}

// Specifies a linear sequence of buckets that all have the same width
// (except overflow and underflow). Each bucket represents a constant
// absolute uncertainty on the specific value in the bucket.
//
// There are `num_finite_buckets + 2` (= N) buckets. Bucket `i` has the
// following boundaries:
//
//    Upper bound (0 <= i < N-1):     offset + (width * i).
//    Lower bound (1 <= i < N):       offset + (width * (i - 1)).
type Distribution_BucketOptions_Linear struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must be greater than 0.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets,proto3" json:"num_finite_buckets,omitempty"`
	// Must be greater than 0.
	Width float64 `protobuf:"fixed64,2,opt,name=width,proto3" json:"width,omitempty"`
	// Lower bound of the first bucket.
	Offset float64 `protobuf:"fixed64,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Distribution_BucketOptions_Linear) Reset() {
	*x = Distribution_BucketOptions_Linear{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_distribution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_BucketOptions_Linear) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_BucketOptions_Linear) ProtoMessage() {}

func (x *Distribution_BucketOptions_Linear) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_distribution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_BucketOptions_Linear.ProtoReflect.Descriptor instead.
func (*Distribution_BucketOptions_Linear) Descriptor() ([]byte, []int) {
	return file_google_api_distribution_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Distribution_BucketOptions_Linear) GetNumFiniteBuckets() int32 {
	if x != nil {
		return x.NumFiniteBuckets
	}
	return 0
}

func (x *Distribution_BucketOptions_Linear) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Distribution_BucketOptions_Linear) GetOffset() float64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Specifies an exponential sequence of buckets that have a width that is
// proportional to the value of the lower bound. Each bucket represents a
// constant relative uncertainty on a specific value in the bucket.
//
// There are `num_finite_buckets + 2` (= N) buckets. Bucket `i` has the
// following boundaries:
//
//    Upper bound (0 <= i < N-1):     scale * (growth_factor ^ i).
//    Lower bound (1 <= i < N):       scale * (growth_factor ^ (i - 1)).
type Distribution_BucketOptions_Exponential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must be greater than 0.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets,proto3" json:"num_finite_buckets,omitempty"`
	// Must be greater than 1.
	GrowthFactor float64 `protobuf:"fixed64,2,opt,name=growth_factor,json=growthFactor,proto3" json:"growth_factor,omitempty"`
	// Must be greater than 0.
	Scale float64 `protobuf:"fixed64,3,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *Distribution_BucketOptions_Exponential) Reset() {
	*x = Distribution_BucketOptions_Exponential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_distribution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_BucketOptions_Exponential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_BucketOptions_Exponential) ProtoMessage() {}

func (x *Distribution_BucketOptions_Exponential) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_distribution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_BucketOptions_Exponential.ProtoReflect.Descriptor instead.
func (*Distribution_BucketOptions_Exponential) Descriptor() ([]byte, []int) {
	return file_google_api_distribution_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *Distribution_BucketOptions_Exponential) GetNumFiniteBuckets() int32 {
	if x != nil {
		return x.NumFiniteBuckets
	}
	return 0
}

func (x *Distribution_BucketOptions_Exponential) GetGrowthFactor() float64 {
	if x != nil {
		return x.GrowthFactor
	}
	return 0
}

func (x *Distribution_BucketOptions_Exponential) GetScale() float64 {
	if x != nil {
		return x.Scale
	}
	return 0
}

// Specifies a set of buckets with arbitrary widths.
//
// There are `size(bounds) + 1` (= N) buckets. Bucket `i` has the following
// boundaries:
//
//    Upper bound (0 <= i < N-1):     bounds[i]
//    Lower bound (1 <= i < N);       bounds[i - 1]
//
// The `bounds` field must contain at least one element. If `bounds` has
// only one element, then there are no finite buckets, and that single
// element is the common boundary of the overflow and underflow buckets.
type Distribution_BucketOptions_Explicit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The values must be monotonically increasing.
	Bounds []float64 `protobuf:"fixed64,1,rep,packed,name=bounds,proto3" json:"bounds,omitempty"`
}

func (x *Distribution_BucketOptions_Explicit) Reset() {
	*x = Distribution_BucketOptions_Explicit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_distribution_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution_BucketOptions_Explicit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution_BucketOptions_Explicit) ProtoMessage() {}

func (x *Distribution_BucketOptions_Explicit) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_distribution_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution_BucketOptions_Explicit.ProtoReflect.Descriptor instead.
func (*Distribution_BucketOptions_Explicit) Descriptor() ([]byte, []int) {
	return file_google_api_distribution_proto_rawDescGZIP(), []int{0, 1, 2}
}

func (x *Distribution_BucketOptions_Explicit) GetBounds() []float64 {
	if x != nil {
		return x.Bounds
	}
	return nil
}

var File_google_api_distribution_proto protoreflect.FileDescriptor

var file_google_api_distribution_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x08, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6d, 0x65,
	0x61, 0x6e, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x75, 0x6d, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x73, 0x75, 0x6d, 0x4f, 0x66, 0x53, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x05, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x52, 0x09, 0x65, 0x78, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x72, 0x73, 0x1a, 0x2b, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6d, 0x61, 0x78, 0x1a, 0xb9, 0x04, 0x0a, 0x0d, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x56, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x5f,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x48, 0x00, 0x52, 0x0d,
	0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x65, 0x0a,
	0x13, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x48, 0x00,
	0x52, 0x12, 0x65, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x12, 0x5c, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74,
	0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x48,
	0x00, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x1a, 0x64, 0x0a, 0x06, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x12, 0x2c, 0x0a, 0x12,
	0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x6e,
	0x69, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x1a, 0x76, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x5f, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x67, 0x72,
	0x6f, 0x77, 0x74, 0x68, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x1a, 0x22, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x06, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x92, 0x01, 0x0a, 0x08, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x36, 0x0a, 0x0b,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x71, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x11, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x3b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0xa2, 0x02, 0x04, 0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_distribution_proto_rawDescOnce sync.Once
	file_google_api_distribution_proto_rawDescData = file_google_api_distribution_proto_rawDesc
)

func file_google_api_distribution_proto_rawDescGZIP() []byte {
	file_google_api_distribution_proto_rawDescOnce.Do(func() {
		file_google_api_distribution_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_distribution_proto_rawDescData)
	})
	return file_google_api_distribution_proto_rawDescData
}

var file_google_api_distribution_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_api_distribution_proto_goTypes = []interface{}{
	(*Distribution)(nil),                           // 0: google.api.Distribution
	(*Distribution_Range)(nil),                     // 1: google.api.Distribution.Range
	(*Distribution_BucketOptions)(nil),             // 2: google.api.Distribution.BucketOptions
	(*Distribution_Exemplar)(nil),                  // 3: google.api.Distribution.Exemplar
	(*Distribution_BucketOptions_Linear)(nil),      // 4: google.api.Distribution.BucketOptions.Linear
	(*Distribution_BucketOptions_Exponential)(nil), // 5: google.api.Distribution.BucketOptions.Exponential
	(*Distribution_BucketOptions_Explicit)(nil),    // 6: google.api.Distribution.BucketOptions.Explicit
	(*timestamppb.Timestamp)(nil),                  // 7: google.protobuf.Timestamp
	(*anypb.Any)(nil),                              // 8: google.protobuf.Any
}
var file_google_api_distribution_proto_depIdxs = []int32{
	1, // 0: google.api.Distribution.range:type_name -> google.api.Distribution.Range
	2, // 1: google.api.Distribution.bucket_options:type_name -> google.api.Distribution.BucketOptions
	3, // 2: google.api.Distribution.exemplars:type_name -> google.api.Distribution.Exemplar
	4, // 3: google.api.Distribution.BucketOptions.linear_buckets:type_name -> google.api.Distribution.BucketOptions.Linear
	5, // 4: google.api.Distribution.BucketOptions.exponential_buckets:type_name -> google.api.Distribution.BucketOptions.Exponential
	6, // 5: google.api.Distribution.BucketOptions.explicit_buckets:type_name -> google.api.Distribution.BucketOptions.Explicit
	7, // 6: google.api.Distribution.Exemplar.timestamp:type_name -> google.protobuf.Timestamp
	8, // 7: google.api.Distribution.Exemplar.attachments:type_name -> google.protobuf.Any
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_api_distribution_proto_init() }
func file_google_api_distribution_proto_init() {
	if File_google_api_distribution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_distribution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_distribution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_Range); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_distribution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_BucketOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_distribution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_Exemplar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_distribution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_BucketOptions_Linear); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_distribution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_BucketOptions_Exponential); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_distribution_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution_BucketOptions_Explicit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_api_distribution_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Distribution_BucketOptions_LinearBuckets)(nil),
		(*Distribution_BucketOptions_ExponentialBuckets)(nil),
		(*Distribution_BucketOptions_ExplicitBuckets)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_distribution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_distribution_proto_goTypes,
		DependencyIndexes: file_google_api_distribution_proto_depIdxs,
		MessageInfos:      file_google_api_distribution_proto_msgTypes,
	}.Build()
	File_google_api_distribution_proto = out.File
	file_google_api_distribution_proto_rawDesc = nil
	file_google_api_distribution_proto_goTypes = nil
	file_google_api_distribution_proto_depIdxs = nil
}
