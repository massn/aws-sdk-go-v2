// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xray

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutTelemetryRecordsInput struct {
	_ struct{} `type:"structure"`

	EC2InstanceId *string `type:"string"`

	Hostname *string `type:"string"`

	ResourceARN *string `type:"string"`

	// TelemetryRecords is a required field
	TelemetryRecords []TelemetryRecord `type:"list" required:"true"`
}

// String returns the string representation
func (s PutTelemetryRecordsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutTelemetryRecordsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutTelemetryRecordsInput"}

	if s.TelemetryRecords == nil {
		invalidParams.Add(aws.NewErrParamRequired("TelemetryRecords"))
	}
	if s.TelemetryRecords != nil {
		for i, v := range s.TelemetryRecords {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TelemetryRecords", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutTelemetryRecordsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EC2InstanceId != nil {
		v := *s.EC2InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EC2InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Hostname != nil {
		v := *s.Hostname

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Hostname", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceARN != nil {
		v := *s.ResourceARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TelemetryRecords != nil {
		v := s.TelemetryRecords

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TelemetryRecords", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type PutTelemetryRecordsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutTelemetryRecordsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutTelemetryRecordsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutTelemetryRecords = "PutTelemetryRecords"

// PutTelemetryRecordsRequest returns a request value for making API operation for
// AWS X-Ray.
//
// Used by the AWS X-Ray daemon to upload telemetry.
//
//    // Example sending a request using PutTelemetryRecordsRequest.
//    req := client.PutTelemetryRecordsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/xray-2016-04-12/PutTelemetryRecords
func (c *Client) PutTelemetryRecordsRequest(input *PutTelemetryRecordsInput) PutTelemetryRecordsRequest {
	op := &aws.Operation{
		Name:       opPutTelemetryRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/TelemetryRecords",
	}

	if input == nil {
		input = &PutTelemetryRecordsInput{}
	}

	req := c.newRequest(op, input, &PutTelemetryRecordsOutput{})

	return PutTelemetryRecordsRequest{Request: req, Input: input, Copy: c.PutTelemetryRecordsRequest}
}

// PutTelemetryRecordsRequest is the request type for the
// PutTelemetryRecords API operation.
type PutTelemetryRecordsRequest struct {
	*aws.Request
	Input *PutTelemetryRecordsInput
	Copy  func(*PutTelemetryRecordsInput) PutTelemetryRecordsRequest
}

// Send marshals and sends the PutTelemetryRecords API request.
func (r PutTelemetryRecordsRequest) Send(ctx context.Context) (*PutTelemetryRecordsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutTelemetryRecordsResponse{
		PutTelemetryRecordsOutput: r.Request.Data.(*PutTelemetryRecordsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutTelemetryRecordsResponse is the response type for the
// PutTelemetryRecords API operation.
type PutTelemetryRecordsResponse struct {
	*PutTelemetryRecordsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutTelemetryRecords request.
func (r *PutTelemetryRecordsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
