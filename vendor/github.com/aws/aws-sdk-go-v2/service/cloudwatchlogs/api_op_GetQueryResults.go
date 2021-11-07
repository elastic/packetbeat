// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetQueryResultsInput struct {
	_ struct{} `type:"structure"`

	// The ID number of the query.
	//
	// QueryId is a required field
	QueryId *string `locationName:"queryId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetQueryResultsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueryResultsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetQueryResultsInput"}

	if s.QueryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetQueryResultsOutput struct {
	_ struct{} `type:"structure"`

	// The log events that matched the query criteria during the most recent time
	// it ran.
	//
	// The results value is an array of arrays. Each log event is one object in
	// the top-level array. Each of these log event objects is an array of field/value
	// pairs.
	Results [][]ResultField `locationName:"results" type:"list"`

	// Includes the number of log events scanned by the query, the number of log
	// events that matched the query criteria, and the total number of bytes in
	// the log events that were scanned.
	Statistics *QueryStatistics `locationName:"statistics" type:"structure"`

	// The status of the most recent running of the query. Possible values are Cancelled,
	// Complete, Failed, Running, Scheduled, Timeout, and Unknown.
	//
	// Queries time out after 15 minutes of execution. To avoid having your queries
	// time out, reduce the time range being searched, or partition your query into
	// a number of queries.
	Status QueryStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetQueryResultsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetQueryResults = "GetQueryResults"

// GetQueryResultsRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Returns the results from the specified query.
//
// Only the fields requested in the query are returned, along with a @ptr field
// which is the identifier for the log record. You can use the value of @ptr
// in a GetLogRecord (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogRecord.html)
// operation to get the full log record.
//
// GetQueryResults does not start a query execution. To run a query, use StartQuery
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html).
//
// If the value of the Status field in the output is Running, this operation
// returns only partial results. If you see a value of Scheduled or Running
// for the status, you can retry the operation later to see the final results.
//
//    // Example sending a request using GetQueryResultsRequest.
//    req := client.GetQueryResultsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/GetQueryResults
func (c *Client) GetQueryResultsRequest(input *GetQueryResultsInput) GetQueryResultsRequest {
	op := &aws.Operation{
		Name:       opGetQueryResults,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetQueryResultsInput{}
	}

	req := c.newRequest(op, input, &GetQueryResultsOutput{})

	return GetQueryResultsRequest{Request: req, Input: input, Copy: c.GetQueryResultsRequest}
}

// GetQueryResultsRequest is the request type for the
// GetQueryResults API operation.
type GetQueryResultsRequest struct {
	*aws.Request
	Input *GetQueryResultsInput
	Copy  func(*GetQueryResultsInput) GetQueryResultsRequest
}

// Send marshals and sends the GetQueryResults API request.
func (r GetQueryResultsRequest) Send(ctx context.Context) (*GetQueryResultsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetQueryResultsResponse{
		GetQueryResultsOutput: r.Request.Data.(*GetQueryResultsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetQueryResultsResponse is the response type for the
// GetQueryResults API operation.
type GetQueryResultsResponse struct {
	*GetQueryResultsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetQueryResults request.
func (r *GetQueryResultsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
