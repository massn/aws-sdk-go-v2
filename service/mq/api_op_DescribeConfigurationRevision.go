// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeConfigurationRevisionInput struct {
	_ struct{} `type:"structure"`

	// ConfigurationId is a required field
	ConfigurationId *string `location:"uri" locationName:"configuration-id" type:"string" required:"true"`

	// ConfigurationRevision is a required field
	ConfigurationRevision *string `location:"uri" locationName:"configuration-revision" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeConfigurationRevisionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConfigurationRevisionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConfigurationRevisionInput"}

	if s.ConfigurationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationId"))
	}

	if s.ConfigurationRevision == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationRevision"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeConfigurationRevisionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigurationId != nil {
		v := *s.ConfigurationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "configuration-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationRevision != nil {
		v := *s.ConfigurationRevision

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "configuration-revision", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeConfigurationRevisionOutput struct {
	_ struct{} `type:"structure"`

	ConfigurationId *string `locationName:"configurationId" type:"string"`

	Created *time.Time `locationName:"created" type:"timestamp" timestampFormat:"iso8601"`

	Data *string `locationName:"data" type:"string"`

	Description *string `locationName:"description" type:"string"`
}

// String returns the string representation
func (s DescribeConfigurationRevisionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeConfigurationRevisionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConfigurationId != nil {
		v := *s.ConfigurationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configurationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Created != nil {
		v := *s.Created

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "created",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Data != nil {
		v := *s.Data

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "data", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeConfigurationRevision = "DescribeConfigurationRevision"

// DescribeConfigurationRevisionRequest returns a request value for making API operation for
// AmazonMQ.
//
// Returns the specified configuration revision for the specified configuration.
//
//    // Example sending a request using DescribeConfigurationRevisionRequest.
//    req := client.DescribeConfigurationRevisionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DescribeConfigurationRevision
func (c *Client) DescribeConfigurationRevisionRequest(input *DescribeConfigurationRevisionInput) DescribeConfigurationRevisionRequest {
	op := &aws.Operation{
		Name:       opDescribeConfigurationRevision,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/configurations/{configuration-id}/revisions/{configuration-revision}",
	}

	if input == nil {
		input = &DescribeConfigurationRevisionInput{}
	}

	req := c.newRequest(op, input, &DescribeConfigurationRevisionOutput{})

	return DescribeConfigurationRevisionRequest{Request: req, Input: input, Copy: c.DescribeConfigurationRevisionRequest}
}

// DescribeConfigurationRevisionRequest is the request type for the
// DescribeConfigurationRevision API operation.
type DescribeConfigurationRevisionRequest struct {
	*aws.Request
	Input *DescribeConfigurationRevisionInput
	Copy  func(*DescribeConfigurationRevisionInput) DescribeConfigurationRevisionRequest
}

// Send marshals and sends the DescribeConfigurationRevision API request.
func (r DescribeConfigurationRevisionRequest) Send(ctx context.Context) (*DescribeConfigurationRevisionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationRevisionResponse{
		DescribeConfigurationRevisionOutput: r.Request.Data.(*DescribeConfigurationRevisionOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationRevisionResponse is the response type for the
// DescribeConfigurationRevision API operation.
type DescribeConfigurationRevisionResponse struct {
	*DescribeConfigurationRevisionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfigurationRevision request.
func (r *DescribeConfigurationRevisionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
