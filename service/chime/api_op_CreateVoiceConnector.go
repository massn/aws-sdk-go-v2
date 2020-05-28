// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateVoiceConnectorInput struct {
	_ struct{} `type:"structure"`

	// The AWS Region in which the Amazon Chime Voice Connector is created. Default
	// value: us-east-1.
	AwsRegion VoiceConnectorAwsRegion `type:"string" enum:"true"`

	// The name of the Amazon Chime Voice Connector.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// When enabled, requires encryption for the Amazon Chime Voice Connector.
	//
	// RequireEncryption is a required field
	RequireEncryption *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s CreateVoiceConnectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVoiceConnectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVoiceConnectorInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.RequireEncryption == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequireEncryption"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVoiceConnectorInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.AwsRegion) > 0 {
		v := s.AwsRegion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AwsRegion", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequireEncryption != nil {
		v := *s.RequireEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequireEncryption", protocol.BoolValue(v), metadata)
	}
	return nil
}

type CreateVoiceConnectorOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector details.
	VoiceConnector *VoiceConnector `type:"structure"`
}

// String returns the string representation
func (s CreateVoiceConnectorOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVoiceConnectorOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VoiceConnector != nil {
		v := s.VoiceConnector

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VoiceConnector", v, metadata)
	}
	return nil
}

const opCreateVoiceConnector = "CreateVoiceConnector"

// CreateVoiceConnectorRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates an Amazon Chime Voice Connector under the administrator's AWS account.
// You can choose to create an Amazon Chime Voice Connector in a specific AWS
// Region.
//
// Enabling CreateVoiceConnectorRequest$RequireEncryption configures your Amazon
// Chime Voice Connector to use TLS transport for SIP signaling and Secure RTP
// (SRTP) for media. Inbound calls use TLS transport, and unencrypted outbound
// calls are blocked.
//
//    // Example sending a request using CreateVoiceConnectorRequest.
//    req := client.CreateVoiceConnectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreateVoiceConnector
func (c *Client) CreateVoiceConnectorRequest(input *CreateVoiceConnectorInput) CreateVoiceConnectorRequest {
	op := &aws.Operation{
		Name:       opCreateVoiceConnector,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors",
	}

	if input == nil {
		input = &CreateVoiceConnectorInput{}
	}

	req := c.newRequest(op, input, &CreateVoiceConnectorOutput{})

	return CreateVoiceConnectorRequest{Request: req, Input: input, Copy: c.CreateVoiceConnectorRequest}
}

// CreateVoiceConnectorRequest is the request type for the
// CreateVoiceConnector API operation.
type CreateVoiceConnectorRequest struct {
	*aws.Request
	Input *CreateVoiceConnectorInput
	Copy  func(*CreateVoiceConnectorInput) CreateVoiceConnectorRequest
}

// Send marshals and sends the CreateVoiceConnector API request.
func (r CreateVoiceConnectorRequest) Send(ctx context.Context) (*CreateVoiceConnectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVoiceConnectorResponse{
		CreateVoiceConnectorOutput: r.Request.Data.(*CreateVoiceConnectorOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVoiceConnectorResponse is the response type for the
// CreateVoiceConnector API operation.
type CreateVoiceConnectorResponse struct {
	*CreateVoiceConnectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVoiceConnector request.
func (r *CreateVoiceConnectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
