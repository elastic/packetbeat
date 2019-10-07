// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/GetMetricWidgetImageInput
type GetMetricWidgetImageInput struct {
	_ struct{} `type:"structure"`

	// A JSON string that defines the bitmap graph to be retrieved. The string includes
	// the metrics to include in the graph, statistics, annotations, title, axis
	// limits, and so on. You can include only one MetricWidget parameter in each
	// GetMetricWidgetImage call.
	//
	// For more information about the syntax of MetricWidget see CloudWatch-Metric-Widget-Structure.
	//
	// If any metric on the graph could not load all the requested data points,
	// an orange triangle with an exclamation point appears next to the graph legend.
	//
	// MetricWidget is a required field
	MetricWidget *string `type:"string" required:"true"`

	// The format of the resulting image. Only PNG images are supported.
	//
	// The default is png. If you specify png, the API returns an HTTP response
	// with the content-type set to text/xml. The image data is in a MetricWidgetImage
	// field. For example:
	//
	// <GetMetricWidgetImageResponse xmlns=<URLstring>>
	//
	// <GetMetricWidgetImageResult>
	//
	// <MetricWidgetImage>
	//
	// iVBORw0KGgoAAAANSUhEUgAAAlgAAAGQEAYAAAAip...
	//
	// </MetricWidgetImage>
	//
	// </GetMetricWidgetImageResult>
	//
	// <ResponseMetadata>
	//
	// <RequestId>6f0d4192-4d42-11e8-82c1-f539a07e0e3b</RequestId>
	//
	// </ResponseMetadata>
	//
	// </GetMetricWidgetImageResponse>
	//
	// The image/png setting is intended only for custom HTTP requests. For most
	// use cases, and all actions using an AWS SDK, you should use png. If you specify
	// image/png, the HTTP response has a content-type set to image/png, and the
	// body of the response is a PNG image.
	OutputFormat *string `type:"string"`
}

// String returns the string representation
func (s GetMetricWidgetImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMetricWidgetImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMetricWidgetImageInput"}

	if s.MetricWidget == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricWidget"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/GetMetricWidgetImageOutput
type GetMetricWidgetImageOutput struct {
	_ struct{} `type:"structure"`

	// The image of the graph, in the output format specified.
	//
	// MetricWidgetImage is automatically base64 encoded/decoded by the SDK.
	MetricWidgetImage []byte `type:"blob"`
}

// String returns the string representation
func (s GetMetricWidgetImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMetricWidgetImage = "GetMetricWidgetImage"

// GetMetricWidgetImageRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// You can use the GetMetricWidgetImage API to retrieve a snapshot graph of
// one or more Amazon CloudWatch metrics as a bitmap image. You can then embed
// this image into your services and products, such as wiki pages, reports,
// and documents. You could also retrieve images regularly, such as every minute,
// and create your own custom live dashboard.
//
// The graph you retrieve can include all CloudWatch metric graph features,
// including metric math and horizontal and vertical annotations.
//
// There is a limit of 20 transactions per second for this API. Each GetMetricWidgetImage
// action has the following limits:
//
//    * As many as 100 metrics in the graph.
//
//    * Up to 100 KB uncompressed payload.
//
//    // Example sending a request using GetMetricWidgetImageRequest.
//    req := client.GetMetricWidgetImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/GetMetricWidgetImage
func (c *Client) GetMetricWidgetImageRequest(input *GetMetricWidgetImageInput) GetMetricWidgetImageRequest {
	op := &aws.Operation{
		Name:       opGetMetricWidgetImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMetricWidgetImageInput{}
	}

	req := c.newRequest(op, input, &GetMetricWidgetImageOutput{})
	return GetMetricWidgetImageRequest{Request: req, Input: input, Copy: c.GetMetricWidgetImageRequest}
}

// GetMetricWidgetImageRequest is the request type for the
// GetMetricWidgetImage API operation.
type GetMetricWidgetImageRequest struct {
	*aws.Request
	Input *GetMetricWidgetImageInput
	Copy  func(*GetMetricWidgetImageInput) GetMetricWidgetImageRequest
}

// Send marshals and sends the GetMetricWidgetImage API request.
func (r GetMetricWidgetImageRequest) Send(ctx context.Context) (*GetMetricWidgetImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMetricWidgetImageResponse{
		GetMetricWidgetImageOutput: r.Request.Data.(*GetMetricWidgetImageOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMetricWidgetImageResponse is the response type for the
// GetMetricWidgetImage API operation.
type GetMetricWidgetImageResponse struct {
	*GetMetricWidgetImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMetricWidgetImage request.
func (r *GetMetricWidgetImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
