// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetHITInput struct {
	_ struct{} `type:"structure"`

	// The ID of the HIT to be retrieved.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetHITInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHITInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHITInput"}

	if s.HITId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITId"))
	}
	if s.HITId != nil && len(*s.HITId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetHITOutput struct {
	_ struct{} `type:"structure"`

	// Contains the requested HIT data.
	HIT *HIT `type:"structure"`
}

// String returns the string representation
func (s GetHITOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetHIT = "GetHIT"

// GetHITRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The GetHIT operation retrieves the details of the specified HIT.
//
//    // Example sending a request using GetHITRequest.
//    req := client.GetHITRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/GetHIT
func (c *Client) GetHITRequest(input *GetHITInput) GetHITRequest {
	op := &aws.Operation{
		Name:       opGetHIT,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetHITInput{}
	}

	req := c.newRequest(op, input, &GetHITOutput{})

	return GetHITRequest{Request: req, Input: input, Copy: c.GetHITRequest}
}

// GetHITRequest is the request type for the
// GetHIT API operation.
type GetHITRequest struct {
	*aws.Request
	Input *GetHITInput
	Copy  func(*GetHITInput) GetHITRequest
}

// Send marshals and sends the GetHIT API request.
func (r GetHITRequest) Send(ctx context.Context) (*GetHITResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHITResponse{
		GetHITOutput: r.Request.Data.(*GetHITOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHITResponse is the response type for the
// GetHIT API operation.
type GetHITResponse struct {
	*GetHITOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHIT request.
func (r *GetHITResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
