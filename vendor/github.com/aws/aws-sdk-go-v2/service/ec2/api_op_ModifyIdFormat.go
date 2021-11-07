// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type ModifyIdFormatInput struct {
	_ struct{} `type:"structure"`

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log
	// | image | import-task | internet-gateway | network-acl | network-acl-association
	// | network-interface | network-interface-attachment | prefix-list | route-table
	// | route-table-association | security-group | subnet | subnet-cidr-block-association
	// | vpc | vpc-cidr-block-association | vpc-endpoint | vpc-peering-connection
	// | vpn-connection | vpn-gateway.
	//
	// Alternatively, use the all-current option to include all resource types that
	// are currently within their opt-in period for longer IDs.
	//
	// Resource is a required field
	Resource *string `type:"string" required:"true"`

	// Indicate whether the resource should use longer IDs (17-character IDs).
	//
	// UseLongIds is a required field
	UseLongIds *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s ModifyIdFormatInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyIdFormatInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyIdFormatInput"}

	if s.Resource == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resource"))
	}

	if s.UseLongIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("UseLongIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyIdFormatOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyIdFormatOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyIdFormat = "ModifyIdFormat"

// ModifyIdFormatRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the ID format for the specified resource on a per-Region basis.
// You can specify that resources should receive longer IDs (17-character IDs)
// when they are created.
//
// This request can only be used to modify longer ID settings for resource types
// that are within the opt-in period. Resources currently in their opt-in period
// include: bundle | conversion-task | customer-gateway | dhcp-options | elastic-ip-allocation
// | elastic-ip-association | export-task | flow-log | image | import-task |
// internet-gateway | network-acl | network-acl-association | network-interface
// | network-interface-attachment | prefix-list | route-table | route-table-association
// | security-group | subnet | subnet-cidr-block-association | vpc | vpc-cidr-block-association
// | vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway.
//
// This setting applies to the IAM user who makes the request; it does not apply
// to the entire AWS account. By default, an IAM user defaults to the same settings
// as the root user. If you're using this action as the root user, then these
// settings apply to the entire account, unless an IAM user explicitly overrides
// these settings for themselves. For more information, see Resource IDs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/resource-ids.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// Resources created with longer IDs are visible to all IAM roles and users,
// regardless of these settings and provided that they have permission to use
// the relevant Describe command for the resource type.
//
//    // Example sending a request using ModifyIdFormatRequest.
//    req := client.ModifyIdFormatRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyIdFormat
func (c *Client) ModifyIdFormatRequest(input *ModifyIdFormatInput) ModifyIdFormatRequest {
	op := &aws.Operation{
		Name:       opModifyIdFormat,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyIdFormatInput{}
	}

	req := c.newRequest(op, input, &ModifyIdFormatOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return ModifyIdFormatRequest{Request: req, Input: input, Copy: c.ModifyIdFormatRequest}
}

// ModifyIdFormatRequest is the request type for the
// ModifyIdFormat API operation.
type ModifyIdFormatRequest struct {
	*aws.Request
	Input *ModifyIdFormatInput
	Copy  func(*ModifyIdFormatInput) ModifyIdFormatRequest
}

// Send marshals and sends the ModifyIdFormat API request.
func (r ModifyIdFormatRequest) Send(ctx context.Context) (*ModifyIdFormatResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyIdFormatResponse{
		ModifyIdFormatOutput: r.Request.Data.(*ModifyIdFormatOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyIdFormatResponse is the response type for the
// ModifyIdFormat API operation.
type ModifyIdFormatResponse struct {
	*ModifyIdFormatOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyIdFormat request.
func (r *ModifyIdFormatResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
