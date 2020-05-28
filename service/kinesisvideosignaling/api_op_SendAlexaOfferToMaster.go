// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideosignaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type SendAlexaOfferToMasterInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the signaling channel by which Alexa and the master peer communicate.
	//
	// ChannelARN is a required field
	ChannelARN *string `min:"1" type:"string" required:"true"`

	// The base64-encoded SDP offer content.
	//
	// MessagePayload is a required field
	MessagePayload *string `min:"1" type:"string" required:"true"`

	// The unique identifier for the sender client.
	//
	// SenderClientId is a required field
	SenderClientId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SendAlexaOfferToMasterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendAlexaOfferToMasterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendAlexaOfferToMasterInput"}

	if s.ChannelARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelARN"))
	}
	if s.ChannelARN != nil && len(*s.ChannelARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChannelARN", 1))
	}

	if s.MessagePayload == nil {
		invalidParams.Add(aws.NewErrParamRequired("MessagePayload"))
	}
	if s.MessagePayload != nil && len(*s.MessagePayload) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MessagePayload", 1))
	}

	if s.SenderClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SenderClientId"))
	}
	if s.SenderClientId != nil && len(*s.SenderClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SenderClientId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendAlexaOfferToMasterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ChannelARN != nil {
		v := *s.ChannelARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ChannelARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MessagePayload != nil {
		v := *s.MessagePayload

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MessagePayload", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SenderClientId != nil {
		v := *s.SenderClientId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SenderClientId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type SendAlexaOfferToMasterOutput struct {
	_ struct{} `type:"structure"`

	// The base64-encoded SDP answer content.
	Answer *string `min:"1" type:"string"`
}

// String returns the string representation
func (s SendAlexaOfferToMasterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SendAlexaOfferToMasterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Answer != nil {
		v := *s.Answer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Answer", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opSendAlexaOfferToMaster = "SendAlexaOfferToMaster"

// SendAlexaOfferToMasterRequest returns a request value for making API operation for
// Amazon Kinesis Video Signaling Channels.
//
// This API allows you to connect WebRTC-enabled devices with Alexa display
// devices. When invoked, it sends the Alexa Session Description Protocol (SDP)
// offer to the master peer. The offer is delivered as soon as the master is
// connected to the specified signaling channel. This API returns the SDP answer
// from the connected master. If the master is not connected to the signaling
// channel, redelivery requests are made until the message expires.
//
//    // Example sending a request using SendAlexaOfferToMasterRequest.
//    req := client.SendAlexaOfferToMasterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-signaling-2019-12-04/SendAlexaOfferToMaster
func (c *Client) SendAlexaOfferToMasterRequest(input *SendAlexaOfferToMasterInput) SendAlexaOfferToMasterRequest {
	op := &aws.Operation{
		Name:       opSendAlexaOfferToMaster,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/send-alexa-offer-to-master",
	}

	if input == nil {
		input = &SendAlexaOfferToMasterInput{}
	}

	req := c.newRequest(op, input, &SendAlexaOfferToMasterOutput{})

	return SendAlexaOfferToMasterRequest{Request: req, Input: input, Copy: c.SendAlexaOfferToMasterRequest}
}

// SendAlexaOfferToMasterRequest is the request type for the
// SendAlexaOfferToMaster API operation.
type SendAlexaOfferToMasterRequest struct {
	*aws.Request
	Input *SendAlexaOfferToMasterInput
	Copy  func(*SendAlexaOfferToMasterInput) SendAlexaOfferToMasterRequest
}

// Send marshals and sends the SendAlexaOfferToMaster API request.
func (r SendAlexaOfferToMasterRequest) Send(ctx context.Context) (*SendAlexaOfferToMasterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendAlexaOfferToMasterResponse{
		SendAlexaOfferToMasterOutput: r.Request.Data.(*SendAlexaOfferToMasterOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendAlexaOfferToMasterResponse is the response type for the
// SendAlexaOfferToMaster API operation.
type SendAlexaOfferToMasterResponse struct {
	*SendAlexaOfferToMasterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendAlexaOfferToMaster request.
func (r *SendAlexaOfferToMasterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
