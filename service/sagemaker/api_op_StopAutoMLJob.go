// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StopAutoMLJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the object you are requesting.
	//
	// AutoMLJobName is a required field
	AutoMLJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopAutoMLJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopAutoMLJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopAutoMLJobInput"}

	if s.AutoMLJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoMLJobName"))
	}
	if s.AutoMLJobName != nil && len(*s.AutoMLJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoMLJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopAutoMLJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopAutoMLJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopAutoMLJob = "StopAutoMLJob"

// StopAutoMLJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// A method for forcing the termination of a running job.
//
//    // Example sending a request using StopAutoMLJobRequest.
//    req := client.StopAutoMLJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopAutoMLJob
func (c *Client) StopAutoMLJobRequest(input *StopAutoMLJobInput) StopAutoMLJobRequest {
	op := &aws.Operation{
		Name:       opStopAutoMLJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopAutoMLJobInput{}
	}

	req := c.newRequest(op, input, &StopAutoMLJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return StopAutoMLJobRequest{Request: req, Input: input, Copy: c.StopAutoMLJobRequest}
}

// StopAutoMLJobRequest is the request type for the
// StopAutoMLJob API operation.
type StopAutoMLJobRequest struct {
	*aws.Request
	Input *StopAutoMLJobInput
	Copy  func(*StopAutoMLJobInput) StopAutoMLJobRequest
}

// Send marshals and sends the StopAutoMLJob API request.
func (r StopAutoMLJobRequest) Send(ctx context.Context) (*StopAutoMLJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopAutoMLJobResponse{
		StopAutoMLJobOutput: r.Request.Data.(*StopAutoMLJobOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopAutoMLJobResponse is the response type for the
// StopAutoMLJob API operation.
type StopAutoMLJobResponse struct {
	*StopAutoMLJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopAutoMLJob request.
func (r *StopAutoMLJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
