// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteVoiceConnectorOriginationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVoiceConnectorOriginationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVoiceConnectorOriginationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVoiceConnectorOriginationInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorOriginationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVoiceConnectorOriginationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVoiceConnectorOriginationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorOriginationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVoiceConnectorOrigination = "DeleteVoiceConnectorOrigination"

// DeleteVoiceConnectorOriginationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the origination settings for the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using DeleteVoiceConnectorOriginationRequest.
//    req := client.DeleteVoiceConnectorOriginationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteVoiceConnectorOrigination
func (c *Client) DeleteVoiceConnectorOriginationRequest(input *DeleteVoiceConnectorOriginationInput) DeleteVoiceConnectorOriginationRequest {
	op := &aws.Operation{
		Name:       opDeleteVoiceConnectorOrigination,
		HTTPMethod: "DELETE",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/origination",
	}

	if input == nil {
		input = &DeleteVoiceConnectorOriginationInput{}
	}

	req := c.newRequest(op, input, &DeleteVoiceConnectorOriginationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteVoiceConnectorOriginationRequest{Request: req, Input: input, Copy: c.DeleteVoiceConnectorOriginationRequest}
}

// DeleteVoiceConnectorOriginationRequest is the request type for the
// DeleteVoiceConnectorOrigination API operation.
type DeleteVoiceConnectorOriginationRequest struct {
	*aws.Request
	Input *DeleteVoiceConnectorOriginationInput
	Copy  func(*DeleteVoiceConnectorOriginationInput) DeleteVoiceConnectorOriginationRequest
}

// Send marshals and sends the DeleteVoiceConnectorOrigination API request.
func (r DeleteVoiceConnectorOriginationRequest) Send(ctx context.Context) (*DeleteVoiceConnectorOriginationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVoiceConnectorOriginationResponse{
		DeleteVoiceConnectorOriginationOutput: r.Request.Data.(*DeleteVoiceConnectorOriginationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVoiceConnectorOriginationResponse is the response type for the
// DeleteVoiceConnectorOrigination API operation.
type DeleteVoiceConnectorOriginationResponse struct {
	*DeleteVoiceConnectorOriginationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVoiceConnectorOrigination request.
func (r *DeleteVoiceConnectorOriginationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
