// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type PutMetricDataInput struct {
	_ struct{} `type:"structure"`

	// The data for the metric. The array can include no more than 20 metrics per
	// call.
	//
	// MetricData is a required field
	MetricData []MetricDatum `type:"list" required:"true"`

	// The namespace for the metric data.
	//
	// To avoid conflicts with AWS service namespaces, you should not specify a
	// namespace that begins with AWS/
	//
	// Namespace is a required field
	Namespace *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutMetricDataInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutMetricDataInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutMetricDataInput"}

	if s.MetricData == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricData"))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}
	if s.MetricData != nil {
		for i, v := range s.MetricData {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MetricData", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutMetricDataOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutMetricDataOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutMetricData = "PutMetricData"

// PutMetricDataRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Publishes metric data points to Amazon CloudWatch. CloudWatch associates
// the data points with the specified metric. If the specified metric does not
// exist, CloudWatch creates the metric. When CloudWatch creates a metric, it
// can take up to fifteen minutes for the metric to appear in calls to ListMetrics
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html).
//
// You can publish either individual data points in the Value field, or arrays
// of values and the number of times each value occurred during the period by
// using the Values and Counts fields in the MetricDatum structure. Using the
// Values and Counts method enables you to publish up to 150 values per metric
// with one PutMetricData request, and supports retrieving percentile statistics
// on this data.
//
// Each PutMetricData request is limited to 40 KB in size for HTTP POST requests.
// You can send a payload compressed by gzip. Each request is also limited to
// no more than 20 different metrics.
//
// Although the Value parameter accepts numbers of type Double, CloudWatch rejects
// values that are either too small or too large. Values must be in the range
// of -2^360 to 2^360. In addition, special values (for example, NaN, +Infinity,
// -Infinity) are not supported.
//
// You can use up to 10 dimensions per metric to further clarify what data the
// metric collects. Each dimension consists of a Name and Value pair. For more
// information about specifying dimensions, see Publishing Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
// in the Amazon CloudWatch User Guide.
//
// Data points with time stamps from 24 hours ago or longer can take at least
// 48 hours to become available for GetMetricData (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html)
// from the time they are submitted. Data points with time stamps between 3
// and 24 hours ago can take as much as 2 hours to become available for for
// GetMetricData (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html).
//
// CloudWatch needs raw data points to calculate percentile statistics. If you
// publish data using a statistic set instead, you can only retrieve percentile
// statistics for this data if one of the following conditions is true:
//
//    * The SampleCount value of the statistic set is 1 and Min, Max, and Sum
//    are all equal.
//
//    * The Min and Max are equal, and Sum is equal to Min multiplied by SampleCount.
//
//    // Example sending a request using PutMetricDataRequest.
//    req := client.PutMetricDataRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutMetricData
func (c *Client) PutMetricDataRequest(input *PutMetricDataInput) PutMetricDataRequest {
	op := &aws.Operation{
		Name:       opPutMetricData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutMetricDataInput{}
	}

	req := c.newRequest(op, input, &PutMetricDataOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return PutMetricDataRequest{Request: req, Input: input, Copy: c.PutMetricDataRequest}
}

// PutMetricDataRequest is the request type for the
// PutMetricData API operation.
type PutMetricDataRequest struct {
	*aws.Request
	Input *PutMetricDataInput
	Copy  func(*PutMetricDataInput) PutMetricDataRequest
}

// Send marshals and sends the PutMetricData API request.
func (r PutMetricDataRequest) Send(ctx context.Context) (*PutMetricDataResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutMetricDataResponse{
		PutMetricDataOutput: r.Request.Data.(*PutMetricDataOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutMetricDataResponse is the response type for the
// PutMetricData API operation.
type PutMetricDataResponse struct {
	*PutMetricDataOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutMetricData request.
func (r *PutMetricDataResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
