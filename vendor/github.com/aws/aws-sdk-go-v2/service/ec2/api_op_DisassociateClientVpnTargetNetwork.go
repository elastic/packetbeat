// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateClientVpnTargetNetworkRequest
type DisassociateClientVpnTargetNetworkInput struct {
	_ struct{} `type:"structure"`

	// The ID of the target network association.
	//
	// AssociationId is a required field
	AssociationId *string `type:"string" required:"true"`

	// The ID of the Client VPN endpoint from which to disassociate the target network.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DisassociateClientVpnTargetNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateClientVpnTargetNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateClientVpnTargetNetworkInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateClientVpnTargetNetworkResult
type DisassociateClientVpnTargetNetworkOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the target network association.
	AssociationId *string `locationName:"associationId" type:"string"`

	// The current state of the target network association.
	Status *AssociationStatus `locationName:"status" type:"structure"`
}

// String returns the string representation
func (s DisassociateClientVpnTargetNetworkOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateClientVpnTargetNetwork = "DisassociateClientVpnTargetNetwork"

// DisassociateClientVpnTargetNetworkRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates a target network from the specified Client VPN endpoint. When
// you disassociate the last target network from a Client VPN, the following
// happens:
//
//    * The route that was automatically added for the VPC is deleted
//
//    * All active client connections are terminated
//
//    * New client connections are disallowed
//
//    * The Client VPN endpoint's status changes to pending-associate
//
//    // Example sending a request using DisassociateClientVpnTargetNetworkRequest.
//    req := client.DisassociateClientVpnTargetNetworkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateClientVpnTargetNetwork
func (c *Client) DisassociateClientVpnTargetNetworkRequest(input *DisassociateClientVpnTargetNetworkInput) DisassociateClientVpnTargetNetworkRequest {
	op := &aws.Operation{
		Name:       opDisassociateClientVpnTargetNetwork,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateClientVpnTargetNetworkInput{}
	}

	req := c.newRequest(op, input, &DisassociateClientVpnTargetNetworkOutput{})
	return DisassociateClientVpnTargetNetworkRequest{Request: req, Input: input, Copy: c.DisassociateClientVpnTargetNetworkRequest}
}

// DisassociateClientVpnTargetNetworkRequest is the request type for the
// DisassociateClientVpnTargetNetwork API operation.
type DisassociateClientVpnTargetNetworkRequest struct {
	*aws.Request
	Input *DisassociateClientVpnTargetNetworkInput
	Copy  func(*DisassociateClientVpnTargetNetworkInput) DisassociateClientVpnTargetNetworkRequest
}

// Send marshals and sends the DisassociateClientVpnTargetNetwork API request.
func (r DisassociateClientVpnTargetNetworkRequest) Send(ctx context.Context) (*DisassociateClientVpnTargetNetworkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateClientVpnTargetNetworkResponse{
		DisassociateClientVpnTargetNetworkOutput: r.Request.Data.(*DisassociateClientVpnTargetNetworkOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateClientVpnTargetNetworkResponse is the response type for the
// DisassociateClientVpnTargetNetwork API operation.
type DisassociateClientVpnTargetNetworkResponse struct {
	*DisassociateClientVpnTargetNetworkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateClientVpnTargetNetwork request.
func (r *DisassociateClientVpnTargetNetworkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
