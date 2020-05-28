// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StopTrainingJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the training job to stop.
	//
	// TrainingJobName is a required field
	TrainingJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopTrainingJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopTrainingJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopTrainingJobInput"}

	if s.TrainingJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrainingJobName"))
	}
	if s.TrainingJobName != nil && len(*s.TrainingJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrainingJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopTrainingJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopTrainingJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopTrainingJob = "StopTrainingJob"

// StopTrainingJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Stops a training job. To stop a job, Amazon SageMaker sends the algorithm
// the SIGTERM signal, which delays job termination for 120 seconds. Algorithms
// might use this 120-second window to save the model artifacts, so the results
// of the training is not lost.
//
// When it receives a StopTrainingJob request, Amazon SageMaker changes the
// status of the job to Stopping. After Amazon SageMaker stops the job, it sets
// the status to Stopped.
//
//    // Example sending a request using StopTrainingJobRequest.
//    req := client.StopTrainingJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/StopTrainingJob
func (c *Client) StopTrainingJobRequest(input *StopTrainingJobInput) StopTrainingJobRequest {
	op := &aws.Operation{
		Name:       opStopTrainingJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopTrainingJobInput{}
	}

	req := c.newRequest(op, input, &StopTrainingJobOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return StopTrainingJobRequest{Request: req, Input: input, Copy: c.StopTrainingJobRequest}
}

// StopTrainingJobRequest is the request type for the
// StopTrainingJob API operation.
type StopTrainingJobRequest struct {
	*aws.Request
	Input *StopTrainingJobInput
	Copy  func(*StopTrainingJobInput) StopTrainingJobRequest
}

// Send marshals and sends the StopTrainingJob API request.
func (r StopTrainingJobRequest) Send(ctx context.Context) (*StopTrainingJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopTrainingJobResponse{
		StopTrainingJobOutput: r.Request.Data.(*StopTrainingJobOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopTrainingJobResponse is the response type for the
// StopTrainingJob API operation.
type StopTrainingJobResponse struct {
	*StopTrainingJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopTrainingJob request.
func (r *StopTrainingJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
