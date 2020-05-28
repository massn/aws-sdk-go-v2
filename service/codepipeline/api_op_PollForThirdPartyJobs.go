// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a PollForThirdPartyJobs action.
type PollForThirdPartyJobsInput struct {
	_ struct{} `type:"structure"`

	// Represents information about an action type.
	//
	// ActionTypeId is a required field
	ActionTypeId *ActionTypeId `locationName:"actionTypeId" type:"structure" required:"true"`

	// The maximum number of jobs to return in a poll for jobs call.
	MaxBatchSize *int64 `locationName:"maxBatchSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s PollForThirdPartyJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PollForThirdPartyJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PollForThirdPartyJobsInput"}

	if s.ActionTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionTypeId"))
	}
	if s.MaxBatchSize != nil && *s.MaxBatchSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxBatchSize", 1))
	}
	if s.ActionTypeId != nil {
		if err := s.ActionTypeId.Validate(); err != nil {
			invalidParams.AddNested("ActionTypeId", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a PollForThirdPartyJobs action.
type PollForThirdPartyJobsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the jobs to take action on.
	Jobs []ThirdPartyJob `locationName:"jobs" type:"list"`
}

// String returns the string representation
func (s PollForThirdPartyJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPollForThirdPartyJobs = "PollForThirdPartyJobs"

// PollForThirdPartyJobsRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Determines whether there are any third party jobs for a job worker to act
// on. Used for partner actions only.
//
// When this API is called, AWS CodePipeline returns temporary credentials for
// the S3 bucket used to store artifacts for the pipeline, if the action requires
// access to that S3 bucket for input or output artifacts.
//
//    // Example sending a request using PollForThirdPartyJobsRequest.
//    req := client.PollForThirdPartyJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/PollForThirdPartyJobs
func (c *Client) PollForThirdPartyJobsRequest(input *PollForThirdPartyJobsInput) PollForThirdPartyJobsRequest {
	op := &aws.Operation{
		Name:       opPollForThirdPartyJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PollForThirdPartyJobsInput{}
	}

	req := c.newRequest(op, input, &PollForThirdPartyJobsOutput{})

	return PollForThirdPartyJobsRequest{Request: req, Input: input, Copy: c.PollForThirdPartyJobsRequest}
}

// PollForThirdPartyJobsRequest is the request type for the
// PollForThirdPartyJobs API operation.
type PollForThirdPartyJobsRequest struct {
	*aws.Request
	Input *PollForThirdPartyJobsInput
	Copy  func(*PollForThirdPartyJobsInput) PollForThirdPartyJobsRequest
}

// Send marshals and sends the PollForThirdPartyJobs API request.
func (r PollForThirdPartyJobsRequest) Send(ctx context.Context) (*PollForThirdPartyJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PollForThirdPartyJobsResponse{
		PollForThirdPartyJobsOutput: r.Request.Data.(*PollForThirdPartyJobsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PollForThirdPartyJobsResponse is the response type for the
// PollForThirdPartyJobs API operation.
type PollForThirdPartyJobsResponse struct {
	*PollForThirdPartyJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PollForThirdPartyJobs request.
func (r *PollForThirdPartyJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
