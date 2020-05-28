// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteApnsVoipChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApnsVoipChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApnsVoipChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApnsVoipChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApnsVoipChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteApnsVoipChannelOutput struct {
	_ struct{} `type:"structure" payload:"APNSVoipChannelResponse"`

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP channel for an application.
	//
	// APNSVoipChannelResponse is a required field
	APNSVoipChannelResponse *APNSVoipChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteApnsVoipChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApnsVoipChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.APNSVoipChannelResponse != nil {
		v := s.APNSVoipChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSVoipChannelResponse", v, metadata)
	}
	return nil
}

const opDeleteApnsVoipChannel = "DeleteApnsVoipChannel"

// DeleteApnsVoipChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Disables the APNs VoIP channel for an application and deletes any existing
// settings for the channel.
//
//    // Example sending a request using DeleteApnsVoipChannelRequest.
//    req := client.DeleteApnsVoipChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteApnsVoipChannel
func (c *Client) DeleteApnsVoipChannelRequest(input *DeleteApnsVoipChannelInput) DeleteApnsVoipChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteApnsVoipChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/channels/apns_voip",
	}

	if input == nil {
		input = &DeleteApnsVoipChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteApnsVoipChannelOutput{})

	return DeleteApnsVoipChannelRequest{Request: req, Input: input, Copy: c.DeleteApnsVoipChannelRequest}
}

// DeleteApnsVoipChannelRequest is the request type for the
// DeleteApnsVoipChannel API operation.
type DeleteApnsVoipChannelRequest struct {
	*aws.Request
	Input *DeleteApnsVoipChannelInput
	Copy  func(*DeleteApnsVoipChannelInput) DeleteApnsVoipChannelRequest
}

// Send marshals and sends the DeleteApnsVoipChannel API request.
func (r DeleteApnsVoipChannelRequest) Send(ctx context.Context) (*DeleteApnsVoipChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApnsVoipChannelResponse{
		DeleteApnsVoipChannelOutput: r.Request.Data.(*DeleteApnsVoipChannelOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApnsVoipChannelResponse is the response type for the
// DeleteApnsVoipChannel API operation.
type DeleteApnsVoipChannelResponse struct {
	*DeleteApnsVoipChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApnsVoipChannel request.
func (r *DeleteApnsVoipChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
