// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateDomainInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain.
	//
	// DomainName is a required field
	DomainName *string `min:"1" type:"string" required:"true"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 1))
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
func (s DisassociateDomainInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisassociateDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateDomainOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisassociateDomain = "DisassociateDomain"

// DisassociateDomainRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Disassociates a domain from Amazon WorkLink. End users lose the ability to
// access the domain with Amazon WorkLink.
//
//    // Example sending a request using DisassociateDomainRequest.
//    req := client.DisassociateDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DisassociateDomain
func (c *Client) DisassociateDomainRequest(input *DisassociateDomainInput) DisassociateDomainRequest {
	op := &aws.Operation{
		Name:       opDisassociateDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/disassociateDomain",
	}

	if input == nil {
		input = &DisassociateDomainInput{}
	}

	req := c.newRequest(op, input, &DisassociateDomainOutput{})

	return DisassociateDomainRequest{Request: req, Input: input, Copy: c.DisassociateDomainRequest}
}

// DisassociateDomainRequest is the request type for the
// DisassociateDomain API operation.
type DisassociateDomainRequest struct {
	*aws.Request
	Input *DisassociateDomainInput
	Copy  func(*DisassociateDomainInput) DisassociateDomainRequest
}

// Send marshals and sends the DisassociateDomain API request.
func (r DisassociateDomainRequest) Send(ctx context.Context) (*DisassociateDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateDomainResponse{
		DisassociateDomainOutput: r.Request.Data.(*DisassociateDomainOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateDomainResponse is the response type for the
// DisassociateDomain API operation.
type DisassociateDomainResponse struct {
	*DisassociateDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateDomain request.
func (r *DisassociateDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
