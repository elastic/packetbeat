// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBLogFilesMessage
type DescribeDBLogFilesInput struct {
	_ struct{} `type:"structure"`

	// The customer-assigned name of the DB instance that contains the log files
	// you want to list.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBInstance.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// Filters the available log files for files written since the specified date,
	// in POSIX timestamp format with milliseconds.
	FileLastWritten *int64 `type:"long"`

	// Filters the available log files for files larger than the specified size.
	FileSize *int64 `type:"long"`

	// Filters the available log files for log file names that contain the specified
	// string.
	FilenameContains *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// The pagination token provided in the previous request. If this parameter
	// is specified the response includes only records beyond the marker, up to
	// MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBLogFilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBLogFilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBLogFilesInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response from a call to DescribeDBLogFiles.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBLogFilesResponse
type DescribeDBLogFilesOutput struct {
	_ struct{} `type:"structure"`

	// The DB log files returned.
	DescribeDBLogFiles []DescribeDBLogFilesDetails `locationNameList:"DescribeDBLogFilesDetails" type:"list"`

	// A pagination token that can be used in a subsequent DescribeDBLogFiles request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBLogFilesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBLogFiles = "DescribeDBLogFiles"

// DescribeDBLogFilesRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns a list of DB log files for the DB instance.
//
//    // Example sending a request using DescribeDBLogFilesRequest.
//    req := client.DescribeDBLogFilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBLogFiles
func (c *Client) DescribeDBLogFilesRequest(input *DescribeDBLogFilesInput) DescribeDBLogFilesRequest {
	op := &aws.Operation{
		Name:       opDescribeDBLogFiles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBLogFilesInput{}
	}

	req := c.newRequest(op, input, &DescribeDBLogFilesOutput{})
	return DescribeDBLogFilesRequest{Request: req, Input: input, Copy: c.DescribeDBLogFilesRequest}
}

// DescribeDBLogFilesRequest is the request type for the
// DescribeDBLogFiles API operation.
type DescribeDBLogFilesRequest struct {
	*aws.Request
	Input *DescribeDBLogFilesInput
	Copy  func(*DescribeDBLogFilesInput) DescribeDBLogFilesRequest
}

// Send marshals and sends the DescribeDBLogFiles API request.
func (r DescribeDBLogFilesRequest) Send(ctx context.Context) (*DescribeDBLogFilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBLogFilesResponse{
		DescribeDBLogFilesOutput: r.Request.Data.(*DescribeDBLogFilesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBLogFilesRequestPaginator returns a paginator for DescribeDBLogFiles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBLogFilesRequest(input)
//   p := rds.NewDescribeDBLogFilesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBLogFilesPaginator(req DescribeDBLogFilesRequest) DescribeDBLogFilesPaginator {
	return DescribeDBLogFilesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBLogFilesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeDBLogFilesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBLogFilesPaginator struct {
	aws.Pager
}

func (p *DescribeDBLogFilesPaginator) CurrentPage() *DescribeDBLogFilesOutput {
	return p.Pager.CurrentPage().(*DescribeDBLogFilesOutput)
}

// DescribeDBLogFilesResponse is the response type for the
// DescribeDBLogFiles API operation.
type DescribeDBLogFilesResponse struct {
	*DescribeDBLogFilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBLogFiles request.
func (r *DescribeDBLogFilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
