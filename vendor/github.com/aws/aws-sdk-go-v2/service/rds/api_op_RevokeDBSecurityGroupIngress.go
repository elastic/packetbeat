// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RevokeDBSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	// The IP range to revoke access from. Must be a valid CIDR range. If CIDRIP
	// is specified, EC2SecurityGroupName, EC2SecurityGroupId and EC2SecurityGroupOwnerId
	// can't be provided.
	CIDRIP *string `type:"string"`

	// The name of the DB security group to revoke ingress from.
	//
	// DBSecurityGroupName is a required field
	DBSecurityGroupName *string `type:"string" required:"true"`

	// The id of the EC2 security group to revoke access from. For VPC DB security
	// groups, EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId
	// and either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupId *string `type:"string"`

	// The name of the EC2 security group to revoke access from. For VPC DB security
	// groups, EC2SecurityGroupId must be provided. Otherwise, EC2SecurityGroupOwnerId
	// and either EC2SecurityGroupName or EC2SecurityGroupId must be provided.
	EC2SecurityGroupName *string `type:"string"`

	// The AWS account number of the owner of the EC2 security group specified in
	// the EC2SecurityGroupName parameter. The AWS access key ID isn't an acceptable
	// value. For VPC DB security groups, EC2SecurityGroupId must be provided. Otherwise,
	// EC2SecurityGroupOwnerId and either EC2SecurityGroupName or EC2SecurityGroupId
	// must be provided.
	EC2SecurityGroupOwnerId *string `type:"string"`
}

// String returns the string representation
func (s RevokeDBSecurityGroupIngressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeDBSecurityGroupIngressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevokeDBSecurityGroupIngressInput"}

	if s.DBSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSecurityGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RevokeDBSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details for an Amazon RDS DB security group.
	//
	// This data type is used as a response element in the DescribeDBSecurityGroups
	// action.
	DBSecurityGroup *DBSecurityGroup `type:"structure"`
}

// String returns the string representation
func (s RevokeDBSecurityGroupIngressOutput) String() string {
	return awsutil.Prettify(s)
}

const opRevokeDBSecurityGroupIngress = "RevokeDBSecurityGroupIngress"

// RevokeDBSecurityGroupIngressRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Revokes ingress from a DBSecurityGroup for previously authorized IP ranges
// or EC2 or VPC Security Groups. Required parameters for this API are one of
// CIDRIP, EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId).
//
//    // Example sending a request using RevokeDBSecurityGroupIngressRequest.
//    req := client.RevokeDBSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RevokeDBSecurityGroupIngress
func (c *Client) RevokeDBSecurityGroupIngressRequest(input *RevokeDBSecurityGroupIngressInput) RevokeDBSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opRevokeDBSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokeDBSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &RevokeDBSecurityGroupIngressOutput{})

	return RevokeDBSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.RevokeDBSecurityGroupIngressRequest}
}

// RevokeDBSecurityGroupIngressRequest is the request type for the
// RevokeDBSecurityGroupIngress API operation.
type RevokeDBSecurityGroupIngressRequest struct {
	*aws.Request
	Input *RevokeDBSecurityGroupIngressInput
	Copy  func(*RevokeDBSecurityGroupIngressInput) RevokeDBSecurityGroupIngressRequest
}

// Send marshals and sends the RevokeDBSecurityGroupIngress API request.
func (r RevokeDBSecurityGroupIngressRequest) Send(ctx context.Context) (*RevokeDBSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeDBSecurityGroupIngressResponse{
		RevokeDBSecurityGroupIngressOutput: r.Request.Data.(*RevokeDBSecurityGroupIngressOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeDBSecurityGroupIngressResponse is the response type for the
// RevokeDBSecurityGroupIngress API operation.
type RevokeDBSecurityGroupIngressResponse struct {
	*RevokeDBSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeDBSecurityGroupIngress request.
func (r *RevokeDBSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
