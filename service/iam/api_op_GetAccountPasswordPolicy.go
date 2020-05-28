// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccountPasswordPolicyInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountPasswordPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a successful GetAccountPasswordPolicy request.
type GetAccountPasswordPolicyOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the account's password policy.
	//
	// PasswordPolicy is a required field
	PasswordPolicy *PasswordPolicy `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetAccountPasswordPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAccountPasswordPolicy = "GetAccountPasswordPolicy"

// GetAccountPasswordPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves the password policy for the AWS account. For more information about
// using a password policy, go to Managing an IAM Password Policy (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingPasswordPolicies.html).
//
//    // Example sending a request using GetAccountPasswordPolicyRequest.
//    req := client.GetAccountPasswordPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetAccountPasswordPolicy
func (c *Client) GetAccountPasswordPolicyRequest(input *GetAccountPasswordPolicyInput) GetAccountPasswordPolicyRequest {
	op := &aws.Operation{
		Name:       opGetAccountPasswordPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAccountPasswordPolicyInput{}
	}

	req := c.newRequest(op, input, &GetAccountPasswordPolicyOutput{})

	return GetAccountPasswordPolicyRequest{Request: req, Input: input, Copy: c.GetAccountPasswordPolicyRequest}
}

// GetAccountPasswordPolicyRequest is the request type for the
// GetAccountPasswordPolicy API operation.
type GetAccountPasswordPolicyRequest struct {
	*aws.Request
	Input *GetAccountPasswordPolicyInput
	Copy  func(*GetAccountPasswordPolicyInput) GetAccountPasswordPolicyRequest
}

// Send marshals and sends the GetAccountPasswordPolicy API request.
func (r GetAccountPasswordPolicyRequest) Send(ctx context.Context) (*GetAccountPasswordPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccountPasswordPolicyResponse{
		GetAccountPasswordPolicyOutput: r.Request.Data.(*GetAccountPasswordPolicyOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccountPasswordPolicyResponse is the response type for the
// GetAccountPasswordPolicy API operation.
type GetAccountPasswordPolicyResponse struct {
	*GetAccountPasswordPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccountPasswordPolicy request.
func (r *GetAccountPasswordPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
