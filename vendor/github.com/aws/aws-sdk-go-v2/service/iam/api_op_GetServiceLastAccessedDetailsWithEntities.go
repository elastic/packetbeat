// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServiceLastAccessedDetailsWithEntitiesRequest
type GetServiceLastAccessedDetailsWithEntitiesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the request generated by the GenerateServiceLastAccessedDetails
	// operation.
	//
	// JobId is a required field
	JobId *string `min:"36" type:"string" required:"true"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The service namespace for an AWS service. Provide the service namespace to
	// learn when the IAM entity last attempted to access the specified service.
	//
	// To learn the service namespace for a service, go to Actions, Resources, and
	// Condition Keys for AWS Services (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_actions-resources-contextkeys.html)
	// in the IAM User Guide. Choose the name of the service to view details for
	// that service. In the first paragraph, find the service prefix. For example,
	// (service prefix: a4b). For more information about service namespaces, see
	// AWS Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces)
	// in the AWS General Reference.
	//
	// ServiceNamespace is a required field
	ServiceNamespace *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetServiceLastAccessedDetailsWithEntitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetServiceLastAccessedDetailsWithEntitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetServiceLastAccessedDetailsWithEntitiesInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 36))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.ServiceNamespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceNamespace"))
	}
	if s.ServiceNamespace != nil && len(*s.ServiceNamespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceNamespace", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServiceLastAccessedDetailsWithEntitiesResponse
type GetServiceLastAccessedDetailsWithEntitiesOutput struct {
	_ struct{} `type:"structure"`

	// An EntityDetailsList object that contains details about when an IAM entity
	// (user or role) used group or policy permissions in an attempt to access the
	// specified AWS service.
	//
	// EntityDetailsList is a required field
	EntityDetailsList []EntityDetails `type:"list" required:"true"`

	// An object that contains details about the reason the operation failed.
	Error *ErrorDetails `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601),
	// when the generated report job was completed or failed.
	//
	// JobCompletionDate is a required field
	JobCompletionDate *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601),
	// when the report job was created.
	//
	// JobCreationDate is a required field
	JobCreationDate *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The status of the job.
	//
	// JobStatus is a required field
	JobStatus JobStatusType `type:"string" required:"true" enum:"true"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s GetServiceLastAccessedDetailsWithEntitiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetServiceLastAccessedDetailsWithEntities = "GetServiceLastAccessedDetailsWithEntities"

// GetServiceLastAccessedDetailsWithEntitiesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// After you generate a group or policy report using the GenerateServiceLastAccessedDetails
// operation, you can use the JobId parameter in GetServiceLastAccessedDetailsWithEntities.
// This operation retrieves the status of your report job and a list of entities
// that could have used group or policy permissions to access the specified
// service.
//
//    * Group – For a group report, this operation returns a list of users
//    in the group that could have used the group’s policies in an attempt
//    to access the service.
//
//    * Policy – For a policy report, this operation returns a list of entities
//    (users or roles) that could have used the policy in an attempt to access
//    the service.
//
// You can also use this operation for user or role reports to retrieve details
// about those entities.
//
// If the operation fails, the GetServiceLastAccessedDetailsWithEntities operation
// returns the reason that it failed.
//
// By default, the list of associated entities is sorted by date, with the most
// recent access listed first.
//
//    // Example sending a request using GetServiceLastAccessedDetailsWithEntitiesRequest.
//    req := client.GetServiceLastAccessedDetailsWithEntitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetServiceLastAccessedDetailsWithEntities
func (c *Client) GetServiceLastAccessedDetailsWithEntitiesRequest(input *GetServiceLastAccessedDetailsWithEntitiesInput) GetServiceLastAccessedDetailsWithEntitiesRequest {
	op := &aws.Operation{
		Name:       opGetServiceLastAccessedDetailsWithEntities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetServiceLastAccessedDetailsWithEntitiesInput{}
	}

	req := c.newRequest(op, input, &GetServiceLastAccessedDetailsWithEntitiesOutput{})
	return GetServiceLastAccessedDetailsWithEntitiesRequest{Request: req, Input: input, Copy: c.GetServiceLastAccessedDetailsWithEntitiesRequest}
}

// GetServiceLastAccessedDetailsWithEntitiesRequest is the request type for the
// GetServiceLastAccessedDetailsWithEntities API operation.
type GetServiceLastAccessedDetailsWithEntitiesRequest struct {
	*aws.Request
	Input *GetServiceLastAccessedDetailsWithEntitiesInput
	Copy  func(*GetServiceLastAccessedDetailsWithEntitiesInput) GetServiceLastAccessedDetailsWithEntitiesRequest
}

// Send marshals and sends the GetServiceLastAccessedDetailsWithEntities API request.
func (r GetServiceLastAccessedDetailsWithEntitiesRequest) Send(ctx context.Context) (*GetServiceLastAccessedDetailsWithEntitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServiceLastAccessedDetailsWithEntitiesResponse{
		GetServiceLastAccessedDetailsWithEntitiesOutput: r.Request.Data.(*GetServiceLastAccessedDetailsWithEntitiesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServiceLastAccessedDetailsWithEntitiesResponse is the response type for the
// GetServiceLastAccessedDetailsWithEntities API operation.
type GetServiceLastAccessedDetailsWithEntitiesResponse struct {
	*GetServiceLastAccessedDetailsWithEntitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetServiceLastAccessedDetailsWithEntities request.
func (r *GetServiceLastAccessedDetailsWithEntitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
