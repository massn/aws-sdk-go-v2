// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeLoggingOptionsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeLoggingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeLoggingOptionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type DescribeLoggingOptionsOutput struct {
	_ struct{} `type:"structure"`

	// The current settings of the AWS IoT Analytics logging options.
	LoggingOptions *LoggingOptions `locationName:"loggingOptions" type:"structure"`
}

// String returns the string representation
func (s DescribeLoggingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeLoggingOptionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LoggingOptions != nil {
		v := s.LoggingOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "loggingOptions", v, metadata)
	}
	return nil
}

const opDescribeLoggingOptions = "DescribeLoggingOptions"

// DescribeLoggingOptionsRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves the current settings of the AWS IoT Analytics logging options.
//
//    // Example sending a request using DescribeLoggingOptionsRequest.
//    req := client.DescribeLoggingOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeLoggingOptions
func (c *Client) DescribeLoggingOptionsRequest(input *DescribeLoggingOptionsInput) DescribeLoggingOptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeLoggingOptions,
		HTTPMethod: "GET",
		HTTPPath:   "/logging",
	}

	if input == nil {
		input = &DescribeLoggingOptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeLoggingOptionsOutput{})

	return DescribeLoggingOptionsRequest{Request: req, Input: input, Copy: c.DescribeLoggingOptionsRequest}
}

// DescribeLoggingOptionsRequest is the request type for the
// DescribeLoggingOptions API operation.
type DescribeLoggingOptionsRequest struct {
	*aws.Request
	Input *DescribeLoggingOptionsInput
	Copy  func(*DescribeLoggingOptionsInput) DescribeLoggingOptionsRequest
}

// Send marshals and sends the DescribeLoggingOptions API request.
func (r DescribeLoggingOptionsRequest) Send(ctx context.Context) (*DescribeLoggingOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoggingOptionsResponse{
		DescribeLoggingOptionsOutput: r.Request.Data.(*DescribeLoggingOptionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLoggingOptionsResponse is the response type for the
// DescribeLoggingOptions API operation.
type DescribeLoggingOptionsResponse struct {
	*DescribeLoggingOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoggingOptions request.
func (r *DescribeLoggingOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
