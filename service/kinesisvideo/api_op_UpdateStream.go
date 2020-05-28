// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateStreamInput struct {
	_ struct{} `type:"structure"`

	// The version of the stream whose metadata you want to update.
	//
	// CurrentVersion is a required field
	CurrentVersion *string `min:"1" type:"string" required:"true"`

	// The name of the device that is writing to the stream.
	//
	// In the current implementation, Kinesis Video Streams does not use this name.
	DeviceName *string `min:"1" type:"string"`

	// The stream's media type. Use MediaType to specify the type of content that
	// the stream contains to the consumers of the stream. For more information
	// about media types, see Media Types (http://www.iana.org/assignments/media-types/media-types.xhtml).
	// If you choose to specify the MediaType, see Naming Requirements (https://tools.ietf.org/html/rfc6838#section-4.2).
	//
	// To play video on the console, you must specify the correct video type. For
	// example, if the video in the stream is H.264, specify video/h264 as the MediaType.
	MediaType *string `min:"1" type:"string"`

	// The ARN of the stream whose metadata you want to update.
	StreamARN *string `min:"1" type:"string"`

	// The name of the stream whose metadata you want to update.
	//
	// The stream name is an identifier for the stream, and must be unique for each
	// account and region.
	StreamName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateStreamInput"}

	if s.CurrentVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentVersion"))
	}
	if s.CurrentVersion != nil && len(*s.CurrentVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentVersion", 1))
	}
	if s.DeviceName != nil && len(*s.DeviceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceName", 1))
	}
	if s.MediaType != nil && len(*s.MediaType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MediaType", 1))
	}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateStreamInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CurrentVersion != nil {
		v := *s.CurrentVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CurrentVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeviceName != nil {
		v := *s.DeviceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeviceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MediaType != nil {
		v := *s.MediaType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MediaType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamARN != nil {
		v := *s.StreamARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateStream = "UpdateStream"

// UpdateStreamRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Updates stream metadata, such as the device name and media type.
//
// You must provide the stream name or the Amazon Resource Name (ARN) of the
// stream.
//
// To make sure that you have the latest version of the stream before updating
// it, you can specify the stream version. Kinesis Video Streams assigns a version
// to each stream. When you update a stream, Kinesis Video Streams assigns a
// new version number. To get the latest stream version, use the DescribeStream
// API.
//
// UpdateStream is an asynchronous operation, and takes time to complete.
//
//    // Example sending a request using UpdateStreamRequest.
//    req := client.UpdateStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/UpdateStream
func (c *Client) UpdateStreamRequest(input *UpdateStreamInput) UpdateStreamRequest {
	op := &aws.Operation{
		Name:       opUpdateStream,
		HTTPMethod: "POST",
		HTTPPath:   "/updateStream",
	}

	if input == nil {
		input = &UpdateStreamInput{}
	}

	req := c.newRequest(op, input, &UpdateStreamOutput{})

	return UpdateStreamRequest{Request: req, Input: input, Copy: c.UpdateStreamRequest}
}

// UpdateStreamRequest is the request type for the
// UpdateStream API operation.
type UpdateStreamRequest struct {
	*aws.Request
	Input *UpdateStreamInput
	Copy  func(*UpdateStreamInput) UpdateStreamRequest
}

// Send marshals and sends the UpdateStream API request.
func (r UpdateStreamRequest) Send(ctx context.Context) (*UpdateStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateStreamResponse{
		UpdateStreamOutput: r.Request.Data.(*UpdateStreamOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateStreamResponse is the response type for the
// UpdateStream API operation.
type UpdateStreamResponse struct {
	*UpdateStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateStream request.
func (r *UpdateStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
