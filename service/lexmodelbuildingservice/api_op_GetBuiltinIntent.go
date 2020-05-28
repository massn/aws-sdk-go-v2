// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBuiltinIntentInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for a built-in intent. To find the signature for an
	// intent, see Standard Built-in Intents (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	// in the Alexa Skills Kit.
	//
	// Signature is a required field
	Signature *string `location:"uri" locationName:"signature" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBuiltinIntentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBuiltinIntentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBuiltinIntentInput"}

	if s.Signature == nil {
		invalidParams.Add(aws.NewErrParamRequired("Signature"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBuiltinIntentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Signature != nil {
		v := *s.Signature

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "signature", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetBuiltinIntentOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for a built-in intent.
	Signature *string `locationName:"signature" type:"string"`

	// An array of BuiltinIntentSlot objects, one entry for each slot type in the
	// intent.
	Slots []BuiltinIntentSlot `locationName:"slots" type:"list"`

	// A list of locales that the intent supports.
	SupportedLocales []Locale `locationName:"supportedLocales" type:"list"`
}

// String returns the string representation
func (s GetBuiltinIntentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBuiltinIntentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Signature != nil {
		v := *s.Signature

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "signature", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Slots != nil {
		v := s.Slots

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "slots", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SupportedLocales != nil {
		v := s.SupportedLocales

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "supportedLocales", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opGetBuiltinIntent = "GetBuiltinIntent"

// GetBuiltinIntentRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns information about a built-in intent.
//
// This operation requires permission for the lex:GetBuiltinIntent action.
//
//    // Example sending a request using GetBuiltinIntentRequest.
//    req := client.GetBuiltinIntentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBuiltinIntent
func (c *Client) GetBuiltinIntentRequest(input *GetBuiltinIntentInput) GetBuiltinIntentRequest {
	op := &aws.Operation{
		Name:       opGetBuiltinIntent,
		HTTPMethod: "GET",
		HTTPPath:   "/builtins/intents/{signature}",
	}

	if input == nil {
		input = &GetBuiltinIntentInput{}
	}

	req := c.newRequest(op, input, &GetBuiltinIntentOutput{})

	return GetBuiltinIntentRequest{Request: req, Input: input, Copy: c.GetBuiltinIntentRequest}
}

// GetBuiltinIntentRequest is the request type for the
// GetBuiltinIntent API operation.
type GetBuiltinIntentRequest struct {
	*aws.Request
	Input *GetBuiltinIntentInput
	Copy  func(*GetBuiltinIntentInput) GetBuiltinIntentRequest
}

// Send marshals and sends the GetBuiltinIntent API request.
func (r GetBuiltinIntentRequest) Send(ctx context.Context) (*GetBuiltinIntentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBuiltinIntentResponse{
		GetBuiltinIntentOutput: r.Request.Data.(*GetBuiltinIntentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBuiltinIntentResponse is the response type for the
// GetBuiltinIntent API operation.
type GetBuiltinIntentResponse struct {
	*GetBuiltinIntentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBuiltinIntent request.
func (r *GetBuiltinIntentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
