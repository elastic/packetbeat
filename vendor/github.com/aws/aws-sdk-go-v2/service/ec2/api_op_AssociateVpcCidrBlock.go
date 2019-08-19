// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateVpcCidrBlockRequest
type AssociateVpcCidrBlockInput struct {
	_ struct{} `type:"structure"`

	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for
	// the VPC. You cannot specify the range of IPv6 addresses, or the size of the
	// CIDR block.
	AmazonProvidedIpv6CidrBlock *bool `locationName:"amazonProvidedIpv6CidrBlock" type:"boolean"`

	// An IPv4 CIDR block to associate with the VPC.
	CidrBlock *string `type:"string"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `locationName:"vpcId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateVpcCidrBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateVpcCidrBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateVpcCidrBlockInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateVpcCidrBlockResult
type AssociateVpcCidrBlockOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IPv4 CIDR block association.
	CidrBlockAssociation *VpcCidrBlockAssociation `locationName:"cidrBlockAssociation" type:"structure"`

	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *VpcIpv6CidrBlockAssociation `locationName:"ipv6CidrBlockAssociation" type:"structure"`

	// The ID of the VPC.
	VpcId *string `locationName:"vpcId" type:"string"`
}

// String returns the string representation
func (s AssociateVpcCidrBlockOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateVpcCidrBlock = "AssociateVpcCidrBlock"

// AssociateVpcCidrBlockRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates a CIDR block with your VPC. You can associate a secondary IPv4
// CIDR block, or you can associate an Amazon-provided IPv6 CIDR block. The
// IPv6 CIDR block size is fixed at /56.
//
// For more information about associating CIDR blocks with your VPC and applicable
// restrictions, see VPC and Subnet Sizing (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Subnets.html#VPC_Sizing)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using AssociateVpcCidrBlockRequest.
//    req := client.AssociateVpcCidrBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateVpcCidrBlock
func (c *Client) AssociateVpcCidrBlockRequest(input *AssociateVpcCidrBlockInput) AssociateVpcCidrBlockRequest {
	op := &aws.Operation{
		Name:       opAssociateVpcCidrBlock,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateVpcCidrBlockInput{}
	}

	req := c.newRequest(op, input, &AssociateVpcCidrBlockOutput{})
	return AssociateVpcCidrBlockRequest{Request: req, Input: input, Copy: c.AssociateVpcCidrBlockRequest}
}

// AssociateVpcCidrBlockRequest is the request type for the
// AssociateVpcCidrBlock API operation.
type AssociateVpcCidrBlockRequest struct {
	*aws.Request
	Input *AssociateVpcCidrBlockInput
	Copy  func(*AssociateVpcCidrBlockInput) AssociateVpcCidrBlockRequest
}

// Send marshals and sends the AssociateVpcCidrBlock API request.
func (r AssociateVpcCidrBlockRequest) Send(ctx context.Context) (*AssociateVpcCidrBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateVpcCidrBlockResponse{
		AssociateVpcCidrBlockOutput: r.Request.Data.(*AssociateVpcCidrBlockOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateVpcCidrBlockResponse is the response type for the
// AssociateVpcCidrBlock API operation.
type AssociateVpcCidrBlockResponse struct {
	*AssociateVpcCidrBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateVpcCidrBlock request.
func (r *AssociateVpcCidrBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
