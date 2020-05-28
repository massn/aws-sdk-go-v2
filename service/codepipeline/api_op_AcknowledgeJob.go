// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an AcknowledgeJob action.
type AcknowledgeJobInput struct {
	_ struct{} `type:"structure"`

	// The unique system-generated ID of the job for which you want to confirm receipt.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`

	// A system-generated random number that AWS CodePipeline uses to ensure that
	// the job is being worked on by only one job worker. Get this number from the
	// response of the PollForJobs request that returned this job.
	//
	// Nonce is a required field
	Nonce *string `locationName:"nonce" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AcknowledgeJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcknowledgeJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcknowledgeJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.Nonce == nil {
		invalidParams.Add(aws.NewErrParamRequired("Nonce"))
	}
	if s.Nonce != nil && len(*s.Nonce) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Nonce", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an AcknowledgeJob action.
type AcknowledgeJobOutput struct {
	_ struct{} `type:"structure"`

	// Whether the job worker has received the specified job.
	Status JobStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s AcknowledgeJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcknowledgeJob = "AcknowledgeJob"

// AcknowledgeJobRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Returns information about a specified job and whether that job has been received
// by the job worker. Used for custom actions only.
//
//    // Example sending a request using AcknowledgeJobRequest.
//    req := client.AcknowledgeJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/AcknowledgeJob
func (c *Client) AcknowledgeJobRequest(input *AcknowledgeJobInput) AcknowledgeJobRequest {
	op := &aws.Operation{
		Name:       opAcknowledgeJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcknowledgeJobInput{}
	}

	req := c.newRequest(op, input, &AcknowledgeJobOutput{})

	return AcknowledgeJobRequest{Request: req, Input: input, Copy: c.AcknowledgeJobRequest}
}

// AcknowledgeJobRequest is the request type for the
// AcknowledgeJob API operation.
type AcknowledgeJobRequest struct {
	*aws.Request
	Input *AcknowledgeJobInput
	Copy  func(*AcknowledgeJobInput) AcknowledgeJobRequest
}

// Send marshals and sends the AcknowledgeJob API request.
func (r AcknowledgeJobRequest) Send(ctx context.Context) (*AcknowledgeJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcknowledgeJobResponse{
		AcknowledgeJobOutput: r.Request.Data.(*AcknowledgeJobOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcknowledgeJobResponse is the response type for the
// AcknowledgeJob API operation.
type AcknowledgeJobResponse struct {
	*AcknowledgeJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcknowledgeJob request.
func (r *AcknowledgeJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
