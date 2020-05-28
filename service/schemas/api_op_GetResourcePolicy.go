// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetResourcePolicyInput struct {
	_ struct{} `type:"structure"`

	RegistryName *string `location:"querystring" locationName:"registryName" type:"string"`
}

// String returns the string representation
func (s GetResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourcePolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	Policy aws.JSONValue `type:"jsonvalue"`

	RevisionId *string `type:"string"`
}

// String returns the string representation
func (s GetResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetResourcePolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policy != nil {
		v := s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Policy", protocol.JSONValue{V: v, EscapeMode: protocol.QuotedEscape}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetResourcePolicy = "GetResourcePolicy"

// GetResourcePolicyRequest returns a request value for making API operation for
// Schemas.
//
// Retrieves the resource-based policy attached to a given registry.
//
//    // Example sending a request using GetResourcePolicyRequest.
//    req := client.GetResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/GetResourcePolicy
func (c *Client) GetResourcePolicyRequest(input *GetResourcePolicyInput) GetResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opGetResourcePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/policy",
	}

	if input == nil {
		input = &GetResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &GetResourcePolicyOutput{})

	return GetResourcePolicyRequest{Request: req, Input: input, Copy: c.GetResourcePolicyRequest}
}

// GetResourcePolicyRequest is the request type for the
// GetResourcePolicy API operation.
type GetResourcePolicyRequest struct {
	*aws.Request
	Input *GetResourcePolicyInput
	Copy  func(*GetResourcePolicyInput) GetResourcePolicyRequest
}

// Send marshals and sends the GetResourcePolicy API request.
func (r GetResourcePolicyRequest) Send(ctx context.Context) (*GetResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourcePolicyResponse{
		GetResourcePolicyOutput: r.Request.Data.(*GetResourcePolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResourcePolicyResponse is the response type for the
// GetResourcePolicy API operation.
type GetResourcePolicyResponse struct {
	*GetResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourcePolicy request.
func (r *GetResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
