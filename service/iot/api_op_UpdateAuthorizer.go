// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the authorizer's Lambda function.
	AuthorizerFunctionArn *string `locationName:"authorizerFunctionArn" type:"string"`

	// The authorizer name.
	//
	// AuthorizerName is a required field
	AuthorizerName *string `location:"uri" locationName:"authorizerName" min:"1" type:"string" required:"true"`

	// The status of the update authorizer request.
	Status AuthorizerStatus `locationName:"status" type:"string" enum:"true"`

	// The key used to extract the token from the HTTP headers.
	TokenKeyName *string `locationName:"tokenKeyName" min:"1" type:"string"`

	// The public keys used to verify the token signature.
	TokenSigningPublicKeys map[string]string `locationName:"tokenSigningPublicKeys" type:"map"`
}

// String returns the string representation
func (s UpdateAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAuthorizerInput"}

	if s.AuthorizerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerName"))
	}
	if s.AuthorizerName != nil && len(*s.AuthorizerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthorizerName", 1))
	}
	if s.TokenKeyName != nil && len(*s.TokenKeyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TokenKeyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthorizerFunctionArn != nil {
		v := *s.AuthorizerFunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerFunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.TokenKeyName != nil {
		v := *s.TokenKeyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "tokenKeyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TokenSigningPublicKeys != nil {
		v := s.TokenSigningPublicKeys

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tokenSigningPublicKeys", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.AuthorizerName != nil {
		v := *s.AuthorizerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "authorizerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateAuthorizerOutput struct {
	_ struct{} `type:"structure"`

	// The authorizer ARN.
	AuthorizerArn *string `locationName:"authorizerArn" type:"string"`

	// The authorizer name.
	AuthorizerName *string `locationName:"authorizerName" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuthorizerArn != nil {
		v := *s.AuthorizerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerName != nil {
		v := *s.AuthorizerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateAuthorizer = "UpdateAuthorizer"

// UpdateAuthorizerRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates an authorizer.
//
//    // Example sending a request using UpdateAuthorizerRequest.
//    req := client.UpdateAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateAuthorizerRequest(input *UpdateAuthorizerInput) UpdateAuthorizerRequest {
	op := &aws.Operation{
		Name:       opUpdateAuthorizer,
		HTTPMethod: "PUT",
		HTTPPath:   "/authorizer/{authorizerName}",
	}

	if input == nil {
		input = &UpdateAuthorizerInput{}
	}

	req := c.newRequest(op, input, &UpdateAuthorizerOutput{})

	return UpdateAuthorizerRequest{Request: req, Input: input, Copy: c.UpdateAuthorizerRequest}
}

// UpdateAuthorizerRequest is the request type for the
// UpdateAuthorizer API operation.
type UpdateAuthorizerRequest struct {
	*aws.Request
	Input *UpdateAuthorizerInput
	Copy  func(*UpdateAuthorizerInput) UpdateAuthorizerRequest
}

// Send marshals and sends the UpdateAuthorizer API request.
func (r UpdateAuthorizerRequest) Send(ctx context.Context) (*UpdateAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAuthorizerResponse{
		UpdateAuthorizerOutput: r.Request.Data.(*UpdateAuthorizerOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAuthorizerResponse is the response type for the
// UpdateAuthorizer API operation.
type UpdateAuthorizerResponse struct {
	*UpdateAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAuthorizer request.
func (r *UpdateAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
