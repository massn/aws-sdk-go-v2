// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateWebsiteAuthorizationProviderInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the authorization provider.
	//
	// AuthorizationProviderId is a required field
	AuthorizationProviderId *string `min:"1" type:"string" required:"true"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateWebsiteAuthorizationProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateWebsiteAuthorizationProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateWebsiteAuthorizationProviderInput"}

	if s.AuthorizationProviderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizationProviderId"))
	}
	if s.AuthorizationProviderId != nil && len(*s.AuthorizationProviderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthorizationProviderId", 1))
	}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateWebsiteAuthorizationProviderInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthorizationProviderId != nil {
		v := *s.AuthorizationProviderId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AuthorizationProviderId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisassociateWebsiteAuthorizationProviderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateWebsiteAuthorizationProviderOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateWebsiteAuthorizationProviderOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisassociateWebsiteAuthorizationProvider = "DisassociateWebsiteAuthorizationProvider"

// DisassociateWebsiteAuthorizationProviderRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Disassociates a website authorization provider from a specified fleet. After
// the disassociation, users can't load any associated websites that require
// this authorization provider.
//
//    // Example sending a request using DisassociateWebsiteAuthorizationProviderRequest.
//    req := client.DisassociateWebsiteAuthorizationProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DisassociateWebsiteAuthorizationProvider
func (c *Client) DisassociateWebsiteAuthorizationProviderRequest(input *DisassociateWebsiteAuthorizationProviderInput) DisassociateWebsiteAuthorizationProviderRequest {
	op := &aws.Operation{
		Name:       opDisassociateWebsiteAuthorizationProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/disassociateWebsiteAuthorizationProvider",
	}

	if input == nil {
		input = &DisassociateWebsiteAuthorizationProviderInput{}
	}

	req := c.newRequest(op, input, &DisassociateWebsiteAuthorizationProviderOutput{})

	return DisassociateWebsiteAuthorizationProviderRequest{Request: req, Input: input, Copy: c.DisassociateWebsiteAuthorizationProviderRequest}
}

// DisassociateWebsiteAuthorizationProviderRequest is the request type for the
// DisassociateWebsiteAuthorizationProvider API operation.
type DisassociateWebsiteAuthorizationProviderRequest struct {
	*aws.Request
	Input *DisassociateWebsiteAuthorizationProviderInput
	Copy  func(*DisassociateWebsiteAuthorizationProviderInput) DisassociateWebsiteAuthorizationProviderRequest
}

// Send marshals and sends the DisassociateWebsiteAuthorizationProvider API request.
func (r DisassociateWebsiteAuthorizationProviderRequest) Send(ctx context.Context) (*DisassociateWebsiteAuthorizationProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateWebsiteAuthorizationProviderResponse{
		DisassociateWebsiteAuthorizationProviderOutput: r.Request.Data.(*DisassociateWebsiteAuthorizationProviderOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateWebsiteAuthorizationProviderResponse is the response type for the
// DisassociateWebsiteAuthorizationProvider API operation.
type DisassociateWebsiteAuthorizationProviderResponse struct {
	*DisassociateWebsiteAuthorizationProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateWebsiteAuthorizationProvider request.
func (r *DisassociateWebsiteAuthorizationProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
