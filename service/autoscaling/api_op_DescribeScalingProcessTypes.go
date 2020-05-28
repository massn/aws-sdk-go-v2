// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeScalingProcessTypesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeScalingProcessTypesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeScalingProcessTypesOutput struct {
	_ struct{} `type:"structure"`

	// The names of the process types.
	Processes []ProcessType `type:"list"`
}

// String returns the string representation
func (s DescribeScalingProcessTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeScalingProcessTypes = "DescribeScalingProcessTypes"

// DescribeScalingProcessTypesRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes the scaling process types for use with the ResumeProcesses and
// SuspendProcesses APIs.
//
//    // Example sending a request using DescribeScalingProcessTypesRequest.
//    req := client.DescribeScalingProcessTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeScalingProcessTypes
func (c *Client) DescribeScalingProcessTypesRequest(input *DescribeScalingProcessTypesInput) DescribeScalingProcessTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeScalingProcessTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingProcessTypesInput{}
	}

	req := c.newRequest(op, input, &DescribeScalingProcessTypesOutput{})

	return DescribeScalingProcessTypesRequest{Request: req, Input: input, Copy: c.DescribeScalingProcessTypesRequest}
}

// DescribeScalingProcessTypesRequest is the request type for the
// DescribeScalingProcessTypes API operation.
type DescribeScalingProcessTypesRequest struct {
	*aws.Request
	Input *DescribeScalingProcessTypesInput
	Copy  func(*DescribeScalingProcessTypesInput) DescribeScalingProcessTypesRequest
}

// Send marshals and sends the DescribeScalingProcessTypes API request.
func (r DescribeScalingProcessTypesRequest) Send(ctx context.Context) (*DescribeScalingProcessTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScalingProcessTypesResponse{
		DescribeScalingProcessTypesOutput: r.Request.Data.(*DescribeScalingProcessTypesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeScalingProcessTypesResponse is the response type for the
// DescribeScalingProcessTypes API operation.
type DescribeScalingProcessTypesResponse struct {
	*DescribeScalingProcessTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScalingProcessTypes request.
func (r *DescribeScalingProcessTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
