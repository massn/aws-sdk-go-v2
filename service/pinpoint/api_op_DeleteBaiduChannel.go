// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteBaiduChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBaiduChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBaiduChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBaiduChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBaiduChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteBaiduChannelOutput struct {
	_ struct{} `type:"structure" payload:"BaiduChannelResponse"`

	// Provides information about the status and settings of the Baidu (Baidu Cloud
	// Push) channel for an application.
	//
	// BaiduChannelResponse is a required field
	BaiduChannelResponse *BaiduChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteBaiduChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBaiduChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BaiduChannelResponse != nil {
		v := s.BaiduChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "BaiduChannelResponse", v, metadata)
	}
	return nil
}

const opDeleteBaiduChannel = "DeleteBaiduChannel"

// DeleteBaiduChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Disables the Baidu channel for an application and deletes any existing settings
// for the channel.
//
//    // Example sending a request using DeleteBaiduChannelRequest.
//    req := client.DeleteBaiduChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteBaiduChannel
func (c *Client) DeleteBaiduChannelRequest(input *DeleteBaiduChannelInput) DeleteBaiduChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteBaiduChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/channels/baidu",
	}

	if input == nil {
		input = &DeleteBaiduChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteBaiduChannelOutput{})

	return DeleteBaiduChannelRequest{Request: req, Input: input, Copy: c.DeleteBaiduChannelRequest}
}

// DeleteBaiduChannelRequest is the request type for the
// DeleteBaiduChannel API operation.
type DeleteBaiduChannelRequest struct {
	*aws.Request
	Input *DeleteBaiduChannelInput
	Copy  func(*DeleteBaiduChannelInput) DeleteBaiduChannelRequest
}

// Send marshals and sends the DeleteBaiduChannel API request.
func (r DeleteBaiduChannelRequest) Send(ctx context.Context) (*DeleteBaiduChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBaiduChannelResponse{
		DeleteBaiduChannelOutput: r.Request.Data.(*DeleteBaiduChannelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBaiduChannelResponse is the response type for the
// DeleteBaiduChannel API operation.
type DeleteBaiduChannelResponse struct {
	*DeleteBaiduChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBaiduChannel request.
func (r *DeleteBaiduChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
