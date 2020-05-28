// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveAutoScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the ID of a cluster. The instance group to which the automatic
	// scaling policy is applied is within this cluster.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	// Specifies the ID of the instance group to which the scaling policy is applied.
	//
	// InstanceGroupId is a required field
	InstanceGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveAutoScalingPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveAutoScalingPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveAutoScalingPolicyInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if s.InstanceGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveAutoScalingPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveAutoScalingPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveAutoScalingPolicy = "RemoveAutoScalingPolicy"

// RemoveAutoScalingPolicyRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Removes an automatic scaling policy from a specified instance group within
// an EMR cluster.
//
//    // Example sending a request using RemoveAutoScalingPolicyRequest.
//    req := client.RemoveAutoScalingPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/RemoveAutoScalingPolicy
func (c *Client) RemoveAutoScalingPolicyRequest(input *RemoveAutoScalingPolicyInput) RemoveAutoScalingPolicyRequest {
	op := &aws.Operation{
		Name:       opRemoveAutoScalingPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveAutoScalingPolicyInput{}
	}

	req := c.newRequest(op, input, &RemoveAutoScalingPolicyOutput{})

	return RemoveAutoScalingPolicyRequest{Request: req, Input: input, Copy: c.RemoveAutoScalingPolicyRequest}
}

// RemoveAutoScalingPolicyRequest is the request type for the
// RemoveAutoScalingPolicy API operation.
type RemoveAutoScalingPolicyRequest struct {
	*aws.Request
	Input *RemoveAutoScalingPolicyInput
	Copy  func(*RemoveAutoScalingPolicyInput) RemoveAutoScalingPolicyRequest
}

// Send marshals and sends the RemoveAutoScalingPolicy API request.
func (r RemoveAutoScalingPolicyRequest) Send(ctx context.Context) (*RemoveAutoScalingPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveAutoScalingPolicyResponse{
		RemoveAutoScalingPolicyOutput: r.Request.Data.(*RemoveAutoScalingPolicyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveAutoScalingPolicyResponse is the response type for the
// RemoveAutoScalingPolicy API operation.
type RemoveAutoScalingPolicyResponse struct {
	*RemoveAutoScalingPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveAutoScalingPolicy request.
func (r *RemoveAutoScalingPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
