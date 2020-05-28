// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateApnsVoipChannelInput struct {
	_ struct{} `type:"structure" payload:"APNSVoipChannelRequest"`

	// Specifies the status and settings of the APNs (Apple Push Notification service)
	// VoIP channel for an application.
	//
	// APNSVoipChannelRequest is a required field
	APNSVoipChannelRequest *APNSVoipChannelRequest `type:"structure" required:"true"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApnsVoipChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApnsVoipChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApnsVoipChannelInput"}

	if s.APNSVoipChannelRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("APNSVoipChannelRequest"))
	}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApnsVoipChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.APNSVoipChannelRequest != nil {
		v := s.APNSVoipChannelRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSVoipChannelRequest", v, metadata)
	}
	return nil
}

type UpdateApnsVoipChannelOutput struct {
	_ struct{} `type:"structure" payload:"APNSVoipChannelResponse"`

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP channel for an application.
	//
	// APNSVoipChannelResponse is a required field
	APNSVoipChannelResponse *APNSVoipChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApnsVoipChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApnsVoipChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.APNSVoipChannelResponse != nil {
		v := s.APNSVoipChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "APNSVoipChannelResponse", v, metadata)
	}
	return nil
}

const opUpdateApnsVoipChannel = "UpdateApnsVoipChannel"

// UpdateApnsVoipChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Enables the APNs VoIP channel for an application or updates the status and
// settings of the APNs VoIP channel for an application.
//
//    // Example sending a request using UpdateApnsVoipChannelRequest.
//    req := client.UpdateApnsVoipChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApnsVoipChannel
func (c *Client) UpdateApnsVoipChannelRequest(input *UpdateApnsVoipChannelInput) UpdateApnsVoipChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateApnsVoipChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/apns_voip",
	}

	if input == nil {
		input = &UpdateApnsVoipChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateApnsVoipChannelOutput{})

	return UpdateApnsVoipChannelRequest{Request: req, Input: input, Copy: c.UpdateApnsVoipChannelRequest}
}

// UpdateApnsVoipChannelRequest is the request type for the
// UpdateApnsVoipChannel API operation.
type UpdateApnsVoipChannelRequest struct {
	*aws.Request
	Input *UpdateApnsVoipChannelInput
	Copy  func(*UpdateApnsVoipChannelInput) UpdateApnsVoipChannelRequest
}

// Send marshals and sends the UpdateApnsVoipChannel API request.
func (r UpdateApnsVoipChannelRequest) Send(ctx context.Context) (*UpdateApnsVoipChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApnsVoipChannelResponse{
		UpdateApnsVoipChannelOutput: r.Request.Data.(*UpdateApnsVoipChannelOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApnsVoipChannelResponse is the response type for the
// UpdateApnsVoipChannel API operation.
type UpdateApnsVoipChannelResponse struct {
	*UpdateApnsVoipChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApnsVoipChannel request.
func (r *UpdateApnsVoipChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
