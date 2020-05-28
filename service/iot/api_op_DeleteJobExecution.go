// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the job execution to be deleted. The executionNumber refers to
	// the execution of a particular job on a particular device.
	//
	// Note that once a job execution is deleted, the executionNumber may be reused
	// by IoT, so be sure you get and use the correct value here.
	//
	// ExecutionNumber is a required field
	ExecutionNumber *int64 `location:"uri" locationName:"executionNumber" type:"long" required:"true"`

	// (Optional) When true, you can delete a job execution which is "IN_PROGRESS".
	// Otherwise, you can only delete a job execution which is in a terminal state
	// ("SUCCEEDED", "FAILED", "REJECTED", "REMOVED" or "CANCELED") or an exception
	// will occur. The default is false.
	//
	// Deleting a job execution which is "IN_PROGRESS", will cause the device to
	// be unable to access job information or update the job execution status. Use
	// caution and ensure that the device is able to recover to a valid state.
	Force *bool `location:"querystring" locationName:"force" type:"boolean"`

	// The ID of the job whose execution on a particular device will be deleted.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// The name of the thing whose job execution will be deleted.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteJobExecutionInput"}

	if s.ExecutionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExecutionNumber"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
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
func (s DeleteJobExecutionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
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
	if s.Force != nil {
		v := *s.Force

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "force", protocol.BoolValue(v), metadata)
	}
	return nil
}

type DeleteJobExecutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteJobExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteJobExecution = "DeleteJobExecution"

// DeleteJobExecutionRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a job execution.
//
//    // Example sending a request using DeleteJobExecutionRequest.
//    req := client.DeleteJobExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteJobExecutionRequest(input *DeleteJobExecutionInput) DeleteJobExecutionRequest {
	op := &aws.Operation{
		Name:       opDeleteJobExecution,
		HTTPMethod: "DELETE",
		HTTPPath:   "/things/{thingName}/jobs/{jobId}/executionNumber/{executionNumber}",
	}

	if input == nil {
		input = &DeleteJobExecutionInput{}
	}

	req := c.newRequest(op, input, &DeleteJobExecutionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteJobExecutionRequest{Request: req, Input: input, Copy: c.DeleteJobExecutionRequest}
}

// DeleteJobExecutionRequest is the request type for the
// DeleteJobExecution API operation.
type DeleteJobExecutionRequest struct {
	*aws.Request
	Input *DeleteJobExecutionInput
	Copy  func(*DeleteJobExecutionInput) DeleteJobExecutionRequest
}

// Send marshals and sends the DeleteJobExecution API request.
func (r DeleteJobExecutionRequest) Send(ctx context.Context) (*DeleteJobExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteJobExecutionResponse{
		DeleteJobExecutionOutput: r.Request.Data.(*DeleteJobExecutionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteJobExecutionResponse is the response type for the
// DeleteJobExecution API operation.
type DeleteJobExecutionResponse struct {
	*DeleteJobExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteJobExecution request.
func (r *DeleteJobExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
