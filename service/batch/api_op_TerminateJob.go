// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type TerminateJobInput struct {
	_ struct{} `type:"structure"`

	// The AWS Batch job ID of the job to terminate.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`

	// A message to attach to the job that explains the reason for canceling it.
	// This message is returned by future DescribeJobs operations on the job. This
	// message is also recorded in the AWS Batch activity logs.
	//
	// Reason is a required field
	Reason *string `locationName:"reason" type:"string" required:"true"`
}

// String returns the string representation
func (s TerminateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.Reason == nil {
		invalidParams.Add(aws.NewErrParamRequired("Reason"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TerminateJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Reason != nil {
		v := *s.Reason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "reason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type TerminateJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TerminateJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TerminateJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opTerminateJob = "TerminateJob"

// TerminateJobRequest returns a request value for making API operation for
// AWS Batch.
//
// Terminates a job in a job queue. Jobs that are in the STARTING or RUNNING
// state are terminated, which causes them to transition to FAILED. Jobs that
// have not progressed to the STARTING state are cancelled.
//
//    // Example sending a request using TerminateJobRequest.
//    req := client.TerminateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/TerminateJob
func (c *Client) TerminateJobRequest(input *TerminateJobInput) TerminateJobRequest {
	op := &aws.Operation{
		Name:       opTerminateJob,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/terminatejob",
	}

	if input == nil {
		input = &TerminateJobInput{}
	}

	req := c.newRequest(op, input, &TerminateJobOutput{})

	return TerminateJobRequest{Request: req, Input: input, Copy: c.TerminateJobRequest}
}

// TerminateJobRequest is the request type for the
// TerminateJob API operation.
type TerminateJobRequest struct {
	*aws.Request
	Input *TerminateJobInput
	Copy  func(*TerminateJobInput) TerminateJobRequest
}

// Send marshals and sends the TerminateJob API request.
func (r TerminateJobRequest) Send(ctx context.Context) (*TerminateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TerminateJobResponse{
		TerminateJobOutput: r.Request.Data.(*TerminateJobOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TerminateJobResponse is the response type for the
// TerminateJob API operation.
type TerminateJobResponse struct {
	*TerminateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TerminateJob request.
func (r *TerminateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
