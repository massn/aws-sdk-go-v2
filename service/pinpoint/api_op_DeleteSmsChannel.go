// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteSmsChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSmsChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSmsChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSmsChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSmsChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSmsChannelOutput struct {
	_ struct{} `type:"structure" payload:"SMSChannelResponse"`

	// Provides information about the status and settings of the SMS channel for
	// an application.
	//
	// SMSChannelResponse is a required field
	SMSChannelResponse *SMSChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteSmsChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSmsChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SMSChannelResponse != nil {
		v := s.SMSChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SMSChannelResponse", v, metadata)
	}
	return nil
}

const opDeleteSmsChannel = "DeleteSmsChannel"

// DeleteSmsChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Disables the SMS channel for an application and deletes any existing settings
// for the channel.
//
//    // Example sending a request using DeleteSmsChannelRequest.
//    req := client.DeleteSmsChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteSmsChannel
func (c *Client) DeleteSmsChannelRequest(input *DeleteSmsChannelInput) DeleteSmsChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteSmsChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/channels/sms",
	}

	if input == nil {
		input = &DeleteSmsChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteSmsChannelOutput{})

	return DeleteSmsChannelRequest{Request: req, Input: input, Copy: c.DeleteSmsChannelRequest}
}

// DeleteSmsChannelRequest is the request type for the
// DeleteSmsChannel API operation.
type DeleteSmsChannelRequest struct {
	*aws.Request
	Input *DeleteSmsChannelInput
	Copy  func(*DeleteSmsChannelInput) DeleteSmsChannelRequest
}

// Send marshals and sends the DeleteSmsChannel API request.
func (r DeleteSmsChannelRequest) Send(ctx context.Context) (*DeleteSmsChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSmsChannelResponse{
		DeleteSmsChannelOutput: r.Request.Data.(*DeleteSmsChannelOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSmsChannelResponse is the response type for the
// DeleteSmsChannel API operation.
type DeleteSmsChannelResponse struct {
	*DeleteSmsChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSmsChannel request.
func (r *DeleteSmsChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
