// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeletePolicyVersionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM policy from which you want to delete
	// a version.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`

	// The policy version to delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that consists of the lowercase letter 'v' followed
	// by one or two digits, and optionally followed by a period '.' and a string
	// of letters and digits.
	//
	// For more information about managed policy versions, see Versioning for Managed
	// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
	// in the IAM User Guide.
	//
	// VersionId is a required field
	VersionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePolicyVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePolicyVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePolicyVersionInput"}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePolicyVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePolicyVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePolicyVersion = "DeletePolicyVersion"

// DeletePolicyVersionRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified version from the specified managed policy.
//
// You cannot delete the default version from a policy using this API. To delete
// the default version from a policy, use DeletePolicy. To find out which version
// of a policy is marked as the default version, use ListPolicyVersions.
//
// For information about versions for managed policies, see Versioning for Managed
// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
// in the IAM User Guide.
//
//    // Example sending a request using DeletePolicyVersionRequest.
//    req := client.DeletePolicyVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeletePolicyVersion
func (c *Client) DeletePolicyVersionRequest(input *DeletePolicyVersionInput) DeletePolicyVersionRequest {
	op := &aws.Operation{
		Name:       opDeletePolicyVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePolicyVersionInput{}
	}

	req := c.newRequest(op, input, &DeletePolicyVersionOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeletePolicyVersionRequest{Request: req, Input: input, Copy: c.DeletePolicyVersionRequest}
}

// DeletePolicyVersionRequest is the request type for the
// DeletePolicyVersion API operation.
type DeletePolicyVersionRequest struct {
	*aws.Request
	Input *DeletePolicyVersionInput
	Copy  func(*DeletePolicyVersionInput) DeletePolicyVersionRequest
}

// Send marshals and sends the DeletePolicyVersion API request.
func (r DeletePolicyVersionRequest) Send(ctx context.Context) (*DeletePolicyVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePolicyVersionResponse{
		DeletePolicyVersionOutput: r.Request.Data.(*DeletePolicyVersionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePolicyVersionResponse is the response type for the
// DeletePolicyVersion API operation.
type DeletePolicyVersionResponse struct {
	*DeletePolicyVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePolicyVersion request.
func (r *DeletePolicyVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
