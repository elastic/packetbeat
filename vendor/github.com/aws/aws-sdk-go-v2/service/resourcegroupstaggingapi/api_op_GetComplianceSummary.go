// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetComplianceSummaryInput struct {
	_ struct{} `type:"structure"`

	// A list of attributes to group the counts of noncompliant resources by. If
	// supplied, the counts are sorted by those attributes.
	GroupBy []GroupByAttribute `type:"list"`

	// A limit that restricts the number of results that are returned per page.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string that indicates that additional data is available. Leave this value
	// empty for your initial request. If the response includes a PaginationToken,
	// use that string for this value to request an additional page of data.
	PaginationToken *string `type:"string"`

	// A list of Regions to limit the output by. If you use this parameter, the
	// count of returned noncompliant resources includes only resources in the specified
	// Regions.
	RegionFilters []string `min:"1" type:"list"`

	// The constraints on the resources that you want returned. The format of each
	// resource type is service[:resourceType]. For example, specifying a resource
	// type of ec2 returns all Amazon EC2 resources (which includes EC2 instances).
	// Specifying a resource type of ec2:instance returns only EC2 instances.
	//
	// The string for each service name and resource type is the same as that embedded
	// in a resource's Amazon Resource Name (ARN). Consult the AWS General Reference
	// for the following:
	//
	//    * For a list of service name strings, see AWS Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces).
	//
	//    * For resource type strings, see Example ARNs (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arns-syntax).
	//
	//    * For more information about ARNs, see Amazon Resource Names (ARNs) and
	//    AWS Service Namespaces (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// You can specify multiple resource types by using an array. The array can
	// include up to 100 items. Note that the length constraint requirement applies
	// to each resource type filter.
	ResourceTypeFilters []string `type:"list"`

	// A list of tag keys to limit the output by. If you use this parameter, the
	// count of returned noncompliant resources includes only resources that have
	// the specified tag keys.
	TagKeyFilters []string `min:"1" type:"list"`

	// The target identifiers (usually, specific account IDs) to limit the output
	// by. If you use this parameter, the count of returned noncompliant resources
	// includes only resources with the specified target IDs.
	TargetIdFilters []string `min:"1" type:"list"`
}

// String returns the string representation
func (s GetComplianceSummaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetComplianceSummaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetComplianceSummaryInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.RegionFilters != nil && len(s.RegionFilters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegionFilters", 1))
	}
	if s.TagKeyFilters != nil && len(s.TagKeyFilters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagKeyFilters", 1))
	}
	if s.TargetIdFilters != nil && len(s.TargetIdFilters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetIdFilters", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetComplianceSummaryOutput struct {
	_ struct{} `type:"structure"`

	// A string that indicates that the response contains more data than can be
	// returned in a single response. To receive additional data, specify this string
	// for the PaginationToken value in a subsequent request.
	PaginationToken *string `type:"string"`

	// A table that shows counts of noncompliant resources.
	SummaryList []Summary `type:"list"`
}

// String returns the string representation
func (s GetComplianceSummaryOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetComplianceSummary = "GetComplianceSummary"

// GetComplianceSummaryRequest returns a request value for making API operation for
// AWS Resource Groups Tagging API.
//
// Returns a table that shows counts of resources that are noncompliant with
// their tag policies.
//
// For more information on tag policies, see Tag Policies (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
// in the AWS Organizations User Guide.
//
// You can call this operation only from the organization's master account and
// from the us-east-1 Region.
//
//    // Example sending a request using GetComplianceSummaryRequest.
//    req := client.GetComplianceSummaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26/GetComplianceSummary
func (c *Client) GetComplianceSummaryRequest(input *GetComplianceSummaryInput) GetComplianceSummaryRequest {
	op := &aws.Operation{
		Name:       opGetComplianceSummary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PaginationToken"},
			OutputTokens:    []string{"PaginationToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetComplianceSummaryInput{}
	}

	req := c.newRequest(op, input, &GetComplianceSummaryOutput{})

	return GetComplianceSummaryRequest{Request: req, Input: input, Copy: c.GetComplianceSummaryRequest}
}

// GetComplianceSummaryRequest is the request type for the
// GetComplianceSummary API operation.
type GetComplianceSummaryRequest struct {
	*aws.Request
	Input *GetComplianceSummaryInput
	Copy  func(*GetComplianceSummaryInput) GetComplianceSummaryRequest
}

// Send marshals and sends the GetComplianceSummary API request.
func (r GetComplianceSummaryRequest) Send(ctx context.Context) (*GetComplianceSummaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetComplianceSummaryResponse{
		GetComplianceSummaryOutput: r.Request.Data.(*GetComplianceSummaryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetComplianceSummaryRequestPaginator returns a paginator for GetComplianceSummary.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetComplianceSummaryRequest(input)
//   p := resourcegroupstaggingapi.NewGetComplianceSummaryRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetComplianceSummaryPaginator(req GetComplianceSummaryRequest) GetComplianceSummaryPaginator {
	return GetComplianceSummaryPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetComplianceSummaryInput
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

// GetComplianceSummaryPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetComplianceSummaryPaginator struct {
	aws.Pager
}

func (p *GetComplianceSummaryPaginator) CurrentPage() *GetComplianceSummaryOutput {
	return p.Pager.CurrentPage().(*GetComplianceSummaryOutput)
}

// GetComplianceSummaryResponse is the response type for the
// GetComplianceSummary API operation.
type GetComplianceSummaryResponse struct {
	*GetComplianceSummaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetComplianceSummary request.
func (r *GetComplianceSummaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
