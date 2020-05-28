// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetMeetingInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMeetingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMeetingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMeetingInput"}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMeetingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetMeetingOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting information.
	Meeting *Meeting `type:"structure"`
}

// String returns the string representation
func (s GetMeetingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMeetingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Meeting != nil {
		v := s.Meeting

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Meeting", v, metadata)
	}
	return nil
}

const opGetMeeting = "GetMeeting"

// GetMeetingRequest returns a request value for making API operation for
// Amazon Chime.
//
// Gets the Amazon Chime SDK meeting details for the specified meeting ID. For
// more information about the Amazon Chime SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
//
//    // Example sending a request using GetMeetingRequest.
//    req := client.GetMeetingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetMeeting
func (c *Client) GetMeetingRequest(input *GetMeetingInput) GetMeetingRequest {
	op := &aws.Operation{
		Name:       opGetMeeting,
		HTTPMethod: "GET",
		HTTPPath:   "/meetings/{meetingId}",
	}

	if input == nil {
		input = &GetMeetingInput{}
	}

	req := c.newRequest(op, input, &GetMeetingOutput{})

	return GetMeetingRequest{Request: req, Input: input, Copy: c.GetMeetingRequest}
}

// GetMeetingRequest is the request type for the
// GetMeeting API operation.
type GetMeetingRequest struct {
	*aws.Request
	Input *GetMeetingInput
	Copy  func(*GetMeetingInput) GetMeetingRequest
}

// Send marshals and sends the GetMeeting API request.
func (r GetMeetingRequest) Send(ctx context.Context) (*GetMeetingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMeetingResponse{
		GetMeetingOutput: r.Request.Data.(*GetMeetingOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMeetingResponse is the response type for the
// GetMeeting API operation.
type GetMeetingResponse struct {
	*GetMeetingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMeeting request.
func (r *GetMeetingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
