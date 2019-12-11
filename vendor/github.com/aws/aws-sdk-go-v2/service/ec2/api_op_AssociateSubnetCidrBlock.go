// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateSubnetCidrBlockRequest
type AssociateSubnetCidrBlockInput struct {
	_ struct{} `type:"structure"`

	// The IPv6 CIDR block for your subnet. The subnet must have a /64 prefix length.
	//
	// Ipv6CidrBlock is a required field
	Ipv6CidrBlock *string `locationName:"ipv6CidrBlock" type:"string" required:"true"`

	// The ID of your subnet.
	//
	// SubnetId is a required field
	SubnetId *string `locationName:"subnetId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateSubnetCidrBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateSubnetCidrBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateSubnetCidrBlockInput"}

	if s.Ipv6CidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ipv6CidrBlock"))
	}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateSubnetCidrBlockResult
type AssociateSubnetCidrBlockOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *SubnetIpv6CidrBlockAssociation `locationName:"ipv6CidrBlockAssociation" type:"structure"`

	// The ID of the subnet.
	SubnetId *string `locationName:"subnetId" type:"string"`
}

// String returns the string representation
func (s AssociateSubnetCidrBlockOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateSubnetCidrBlock = "AssociateSubnetCidrBlock"

// AssociateSubnetCidrBlockRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates a CIDR block with your subnet. You can only associate a single
// IPv6 CIDR block with your subnet. An IPv6 CIDR block must have a prefix length
// of /64.
//
//    // Example sending a request using AssociateSubnetCidrBlockRequest.
//    req := client.AssociateSubnetCidrBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateSubnetCidrBlock
func (c *Client) AssociateSubnetCidrBlockRequest(input *AssociateSubnetCidrBlockInput) AssociateSubnetCidrBlockRequest {
	op := &aws.Operation{
		Name:       opAssociateSubnetCidrBlock,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateSubnetCidrBlockInput{}
	}

	req := c.newRequest(op, input, &AssociateSubnetCidrBlockOutput{})
	return AssociateSubnetCidrBlockRequest{Request: req, Input: input, Copy: c.AssociateSubnetCidrBlockRequest}
}

// AssociateSubnetCidrBlockRequest is the request type for the
// AssociateSubnetCidrBlock API operation.
type AssociateSubnetCidrBlockRequest struct {
	*aws.Request
	Input *AssociateSubnetCidrBlockInput
	Copy  func(*AssociateSubnetCidrBlockInput) AssociateSubnetCidrBlockRequest
}

// Send marshals and sends the AssociateSubnetCidrBlock API request.
func (r AssociateSubnetCidrBlockRequest) Send(ctx context.Context) (*AssociateSubnetCidrBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateSubnetCidrBlockResponse{
		AssociateSubnetCidrBlockOutput: r.Request.Data.(*AssociateSubnetCidrBlockOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateSubnetCidrBlockResponse is the response type for the
// AssociateSubnetCidrBlock API operation.
type AssociateSubnetCidrBlockResponse struct {
	*AssociateSubnetCidrBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateSubnetCidrBlock request.
func (r *AssociateSubnetCidrBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
