// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// Optional. A number that identifies a particular job execution on a particular
	// device. If not specified, the latest job execution is returned.
	ExecutionNumber *int64 `location:"querystring" locationName:"executionNumber" type:"long"`

	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool `location:"querystring" locationName:"includeJobDocument" type:"boolean"`

	// The unique identifier assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`

	// The thing name associated with the device the job execution is running on.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeJobExecutionInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobExecutionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.IncludeJobDocument != nil {
		v := *s.IncludeJobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeJobDocument", protocol.BoolValue(v), metadata)
	}
	return nil
}

type DescribeJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	// Contains data about a job execution.
	Execution *JobExecution `locationName:"execution" type:"structure"`
}

// String returns the string representation
func (s DescribeJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeJobExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Execution != nil {
		v := s.Execution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "execution", v, metadata)
	}
	return nil
}

const opDescribeJobExecution = "DescribeJobExecution"

// DescribeJobExecutionRequest returns a request value for making API operation for
// AWS IoT Jobs Data Plane.
//
// Gets details of a job execution.
//
//    // Example sending a request using DescribeJobExecutionRequest.
//    req := client.DescribeJobExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/DescribeJobExecution
func (c *Client) DescribeJobExecutionRequest(input *DescribeJobExecutionInput) DescribeJobExecutionRequest {
	op := &aws.Operation{
		Name:       opDescribeJobExecution,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}/jobs/{jobId}",
	}

	if input == nil {
		input = &DescribeJobExecutionInput{}
	}

	req := c.newRequest(op, input, &DescribeJobExecutionOutput{})

	return DescribeJobExecutionRequest{Request: req, Input: input, Copy: c.DescribeJobExecutionRequest}
}

// DescribeJobExecutionRequest is the request type for the
// DescribeJobExecution API operation.
type DescribeJobExecutionRequest struct {
	*aws.Request
	Input *DescribeJobExecutionInput
	Copy  func(*DescribeJobExecutionInput) DescribeJobExecutionRequest
}

// Send marshals and sends the DescribeJobExecution API request.
func (r DescribeJobExecutionRequest) Send(ctx context.Context) (*DescribeJobExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobExecutionResponse{
		DescribeJobExecutionOutput: r.Request.Data.(*DescribeJobExecutionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeJobExecutionResponse is the response type for the
// DescribeJobExecution API operation.
type DescribeJobExecutionResponse struct {
	*DescribeJobExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJobExecution request.
func (r *DescribeJobExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
