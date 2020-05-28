// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteSessionInput struct {
	_ struct{} `type:"structure"`

	// The alias in use for the bot that contains the session data.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// The name of the bot that contains the session data.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// The identifier of the user associated with the session data.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSessionInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSessionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BotAlias != nil {
		v := *s.BotAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSessionOutput struct {
	_ struct{} `type:"structure"`

	// The alias in use for the bot associated with the session data.
	BotAlias *string `locationName:"botAlias" type:"string"`

	// The name of the bot associated with the session data.
	BotName *string `locationName:"botName" type:"string"`

	// The unique identifier for the session.
	SessionId *string `locationName:"sessionId" type:"string"`

	// The ID of the client application user.
	UserId *string `locationName:"userId" min:"2" type:"string"`
}

// String returns the string representation
func (s DeleteSessionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSessionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BotAlias != nil {
		v := *s.BotAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "botAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SessionId != nil {
		v := *s.SessionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sessionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteSession = "DeleteSession"

// DeleteSessionRequest returns a request value for making API operation for
// Amazon Lex Runtime Service.
//
// Removes session information for a specified bot, alias, and user ID.
//
//    // Example sending a request using DeleteSessionRequest.
//    req := client.DeleteSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/DeleteSession
func (c *Client) DeleteSessionRequest(input *DeleteSessionInput) DeleteSessionRequest {
	op := &aws.Operation{
		Name:       opDeleteSession,
		HTTPMethod: "DELETE",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/session",
	}

	if input == nil {
		input = &DeleteSessionInput{}
	}

	req := c.newRequest(op, input, &DeleteSessionOutput{})

	return DeleteSessionRequest{Request: req, Input: input, Copy: c.DeleteSessionRequest}
}

// DeleteSessionRequest is the request type for the
// DeleteSession API operation.
type DeleteSessionRequest struct {
	*aws.Request
	Input *DeleteSessionInput
	Copy  func(*DeleteSessionInput) DeleteSessionRequest
}

// Send marshals and sends the DeleteSession API request.
func (r DeleteSessionRequest) Send(ctx context.Context) (*DeleteSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSessionResponse{
		DeleteSessionOutput: r.Request.Data.(*DeleteSessionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSessionResponse is the response type for the
// DeleteSession API operation.
type DeleteSessionResponse struct {
	*DeleteSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSession request.
func (r *DeleteSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
