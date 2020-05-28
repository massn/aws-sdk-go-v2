// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeRdsDbInstancesInput struct {
	_ struct{} `type:"structure"`

	// An array containing the ARNs of the instances to be described.
	RdsDbInstanceArns []string `type:"list"`

	// The ID of the stack with which the instances are registered. The operation
	// returns descriptions of all registered Amazon RDS instances.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRdsDbInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRdsDbInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRdsDbInstancesInput"}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a DescribeRdsDbInstances request.
type DescribeRdsDbInstancesOutput struct {
	_ struct{} `type:"structure"`

	// An a array of RdsDbInstance objects that describe the instances.
	RdsDbInstances []RdsDbInstance `type:"list"`
}

// String returns the string representation
func (s DescribeRdsDbInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRdsDbInstances = "DescribeRdsDbInstances"

// DescribeRdsDbInstancesRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes Amazon RDS instances.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
// This call accepts only one resource-identifying parameter.
//
//    // Example sending a request using DescribeRdsDbInstancesRequest.
//    req := client.DescribeRdsDbInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeRdsDbInstances
func (c *Client) DescribeRdsDbInstancesRequest(input *DescribeRdsDbInstancesInput) DescribeRdsDbInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeRdsDbInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRdsDbInstancesInput{}
	}

	req := c.newRequest(op, input, &DescribeRdsDbInstancesOutput{})

	return DescribeRdsDbInstancesRequest{Request: req, Input: input, Copy: c.DescribeRdsDbInstancesRequest}
}

// DescribeRdsDbInstancesRequest is the request type for the
// DescribeRdsDbInstances API operation.
type DescribeRdsDbInstancesRequest struct {
	*aws.Request
	Input *DescribeRdsDbInstancesInput
	Copy  func(*DescribeRdsDbInstancesInput) DescribeRdsDbInstancesRequest
}

// Send marshals and sends the DescribeRdsDbInstances API request.
func (r DescribeRdsDbInstancesRequest) Send(ctx context.Context) (*DescribeRdsDbInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRdsDbInstancesResponse{
		DescribeRdsDbInstancesOutput: r.Request.Data.(*DescribeRdsDbInstancesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRdsDbInstancesResponse is the response type for the
// DescribeRdsDbInstances API operation.
type DescribeRdsDbInstancesResponse struct {
	*DescribeRdsDbInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRdsDbInstances request.
func (r *DescribeRdsDbInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
