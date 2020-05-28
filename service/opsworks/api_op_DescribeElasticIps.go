// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeElasticIpsInput struct {
	_ struct{} `type:"structure"`

	// The instance ID. If you include this parameter, DescribeElasticIps returns
	// a description of the Elastic IP addresses associated with the specified instance.
	InstanceId *string `type:"string"`

	// An array of Elastic IP addresses to be described. If you include this parameter,
	// DescribeElasticIps returns a description of the specified Elastic IP addresses.
	// Otherwise, it returns a description of every Elastic IP address.
	Ips []string `type:"list"`

	// A stack ID. If you include this parameter, DescribeElasticIps returns a description
	// of the Elastic IP addresses that are registered with the specified stack.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s DescribeElasticIpsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeElasticIps request.
type DescribeElasticIpsOutput struct {
	_ struct{} `type:"structure"`

	// An ElasticIps object that describes the specified Elastic IP addresses.
	ElasticIps []ElasticIp `type:"list"`
}

// String returns the string representation
func (s DescribeElasticIpsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeElasticIps = "DescribeElasticIps"

// DescribeElasticIpsRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes Elastic IP addresses (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html).
//
// This call accepts only one resource-identifying parameter.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeElasticIpsRequest.
//    req := client.DescribeElasticIpsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeElasticIps
func (c *Client) DescribeElasticIpsRequest(input *DescribeElasticIpsInput) DescribeElasticIpsRequest {
	op := &aws.Operation{
		Name:       opDescribeElasticIps,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeElasticIpsInput{}
	}

	req := c.newRequest(op, input, &DescribeElasticIpsOutput{})

	return DescribeElasticIpsRequest{Request: req, Input: input, Copy: c.DescribeElasticIpsRequest}
}

// DescribeElasticIpsRequest is the request type for the
// DescribeElasticIps API operation.
type DescribeElasticIpsRequest struct {
	*aws.Request
	Input *DescribeElasticIpsInput
	Copy  func(*DescribeElasticIpsInput) DescribeElasticIpsRequest
}

// Send marshals and sends the DescribeElasticIps API request.
func (r DescribeElasticIpsRequest) Send(ctx context.Context) (*DescribeElasticIpsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeElasticIpsResponse{
		DescribeElasticIpsOutput: r.Request.Data.(*DescribeElasticIpsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeElasticIpsResponse is the response type for the
// DescribeElasticIps API operation.
type DescribeElasticIpsResponse struct {
	*DescribeElasticIpsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeElasticIps request.
func (r *DescribeElasticIpsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
