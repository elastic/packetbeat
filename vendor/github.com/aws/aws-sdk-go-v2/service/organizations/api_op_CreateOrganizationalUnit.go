// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateOrganizationalUnitInput struct {
	_ struct{} `type:"structure"`

	// The friendly name to assign to the new OU.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The unique identifier (ID) of the parent root or OU that you want to create
	// the new OU in.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//    * Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//    letters or digits.
	//
	//    * Organizational unit (OU) - A string that begins with "ou-" followed
	//    by from 4 to 32 lowercase letters or digits (the ID of the root that the
	//    OU is in). This string is followed by a second "-" dash and from 8 to
	//    32 additional lowercase letters or digits.
	//
	// ParentId is a required field
	ParentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateOrganizationalUnitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOrganizationalUnitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOrganizationalUnitInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.ParentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateOrganizationalUnitOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the newly created OU.
	OrganizationalUnit *OrganizationalUnit `type:"structure"`
}

// String returns the string representation
func (s CreateOrganizationalUnitOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateOrganizationalUnit = "CreateOrganizationalUnit"

// CreateOrganizationalUnitRequest returns a request value for making API operation for
// AWS Organizations.
//
// Creates an organizational unit (OU) within a root or parent OU. An OU is
// a container for accounts that enables you to organize your accounts to apply
// policies according to your business requirements. The number of levels deep
// that you can nest OUs is dependent upon the policy types enabled for that
// root. For service control policies, the limit is five.
//
// For more information about OUs, see Managing Organizational Units (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_ous.html)
// in the AWS Organizations User Guide.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using CreateOrganizationalUnitRequest.
//    req := client.CreateOrganizationalUnitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/CreateOrganizationalUnit
func (c *Client) CreateOrganizationalUnitRequest(input *CreateOrganizationalUnitInput) CreateOrganizationalUnitRequest {
	op := &aws.Operation{
		Name:       opCreateOrganizationalUnit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateOrganizationalUnitInput{}
	}

	req := c.newRequest(op, input, &CreateOrganizationalUnitOutput{})

	return CreateOrganizationalUnitRequest{Request: req, Input: input, Copy: c.CreateOrganizationalUnitRequest}
}

// CreateOrganizationalUnitRequest is the request type for the
// CreateOrganizationalUnit API operation.
type CreateOrganizationalUnitRequest struct {
	*aws.Request
	Input *CreateOrganizationalUnitInput
	Copy  func(*CreateOrganizationalUnitInput) CreateOrganizationalUnitRequest
}

// Send marshals and sends the CreateOrganizationalUnit API request.
func (r CreateOrganizationalUnitRequest) Send(ctx context.Context) (*CreateOrganizationalUnitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOrganizationalUnitResponse{
		CreateOrganizationalUnitOutput: r.Request.Data.(*CreateOrganizationalUnitOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOrganizationalUnitResponse is the response type for the
// CreateOrganizationalUnit API operation.
type CreateOrganizationalUnitResponse struct {
	*CreateOrganizationalUnitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOrganizationalUnit request.
func (r *CreateOrganizationalUnitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
