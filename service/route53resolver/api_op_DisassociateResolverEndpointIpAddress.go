// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateResolverEndpointIpAddressInput struct {
	_ struct{} `type:"structure"`

	// The IPv4 address that you want to remove from a resolver endpoint.
	//
	// IpAddress is a required field
	IpAddress *IpAddressUpdate `type:"structure" required:"true"`

	// The ID of the resolver endpoint that you want to disassociate an IP address
	// from.
	//
	// ResolverEndpointId is a required field
	ResolverEndpointId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateResolverEndpointIpAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateResolverEndpointIpAddressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateResolverEndpointIpAddressInput"}

	if s.IpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpAddress"))
	}

	if s.ResolverEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverEndpointId"))
	}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}
	if s.IpAddress != nil {
		if err := s.IpAddress.Validate(); err != nil {
			invalidParams.AddNested("IpAddress", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateResolverEndpointIpAddressOutput struct {
	_ struct{} `type:"structure"`

	// The response to an DisassociateResolverEndpointIpAddress request.
	ResolverEndpoint *ResolverEndpoint `type:"structure"`
}

// String returns the string representation
func (s DisassociateResolverEndpointIpAddressOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateResolverEndpointIpAddress = "DisassociateResolverEndpointIpAddress"

// DisassociateResolverEndpointIpAddressRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Removes IP addresses from an inbound or an outbound resolver endpoint. If
// you want to remove more than one IP address, submit one DisassociateResolverEndpointIpAddress
// request for each IP address.
//
// To add an IP address to an endpoint, see AssociateResolverEndpointIpAddress.
//
//    // Example sending a request using DisassociateResolverEndpointIpAddressRequest.
//    req := client.DisassociateResolverEndpointIpAddressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/DisassociateResolverEndpointIpAddress
func (c *Client) DisassociateResolverEndpointIpAddressRequest(input *DisassociateResolverEndpointIpAddressInput) DisassociateResolverEndpointIpAddressRequest {
	op := &aws.Operation{
		Name:       opDisassociateResolverEndpointIpAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateResolverEndpointIpAddressInput{}
	}

	req := c.newRequest(op, input, &DisassociateResolverEndpointIpAddressOutput{})

	return DisassociateResolverEndpointIpAddressRequest{Request: req, Input: input, Copy: c.DisassociateResolverEndpointIpAddressRequest}
}

// DisassociateResolverEndpointIpAddressRequest is the request type for the
// DisassociateResolverEndpointIpAddress API operation.
type DisassociateResolverEndpointIpAddressRequest struct {
	*aws.Request
	Input *DisassociateResolverEndpointIpAddressInput
	Copy  func(*DisassociateResolverEndpointIpAddressInput) DisassociateResolverEndpointIpAddressRequest
}

// Send marshals and sends the DisassociateResolverEndpointIpAddress API request.
func (r DisassociateResolverEndpointIpAddressRequest) Send(ctx context.Context) (*DisassociateResolverEndpointIpAddressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateResolverEndpointIpAddressResponse{
		DisassociateResolverEndpointIpAddressOutput: r.Request.Data.(*DisassociateResolverEndpointIpAddressOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateResolverEndpointIpAddressResponse is the response type for the
// DisassociateResolverEndpointIpAddress API operation.
type DisassociateResolverEndpointIpAddressResponse struct {
	*DisassociateResolverEndpointIpAddressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateResolverEndpointIpAddress request.
func (r *DisassociateResolverEndpointIpAddressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}