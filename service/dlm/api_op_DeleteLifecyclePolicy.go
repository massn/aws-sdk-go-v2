// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dlm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteLifecyclePolicyInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the lifecycle policy.
	//
	// PolicyId is a required field
	PolicyId *string `location:"uri" locationName:"policyId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLifecyclePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLifecyclePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLifecyclePolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLifecyclePolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PolicyId != nil {
		v := *s.PolicyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "policyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteLifecyclePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLifecyclePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLifecyclePolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteLifecyclePolicy = "DeleteLifecyclePolicy"

// DeleteLifecyclePolicyRequest returns a request value for making API operation for
// Amazon Data Lifecycle Manager.
//
// Deletes the specified lifecycle policy and halts the automated operations
// that the policy specified.
//
//    // Example sending a request using DeleteLifecyclePolicyRequest.
//    req := client.DeleteLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dlm-2018-01-12/DeleteLifecyclePolicy
func (c *Client) DeleteLifecyclePolicyRequest(input *DeleteLifecyclePolicyInput) DeleteLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteLifecyclePolicy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/policies/{policyId}/",
	}

	if input == nil {
		input = &DeleteLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteLifecyclePolicyOutput{})

	return DeleteLifecyclePolicyRequest{Request: req, Input: input, Copy: c.DeleteLifecyclePolicyRequest}
}

// DeleteLifecyclePolicyRequest is the request type for the
// DeleteLifecyclePolicy API operation.
type DeleteLifecyclePolicyRequest struct {
	*aws.Request
	Input *DeleteLifecyclePolicyInput
	Copy  func(*DeleteLifecyclePolicyInput) DeleteLifecyclePolicyRequest
}

// Send marshals and sends the DeleteLifecyclePolicy API request.
func (r DeleteLifecyclePolicyRequest) Send(ctx context.Context) (*DeleteLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLifecyclePolicyResponse{
		DeleteLifecyclePolicyOutput: r.Request.Data.(*DeleteLifecyclePolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLifecyclePolicyResponse is the response type for the
// DeleteLifecyclePolicy API operation.
type DeleteLifecyclePolicyResponse struct {
	*DeleteLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLifecyclePolicy request.
func (r *DeleteLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
