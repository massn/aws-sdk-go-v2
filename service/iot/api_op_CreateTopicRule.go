// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the CreateTopicRule operation.
type CreateTopicRuleInput struct {
	_ struct{} `type:"structure" payload:"TopicRulePayload"`

	// The name of the rule.
	//
	// RuleName is a required field
	RuleName *string `location:"uri" locationName:"ruleName" min:"1" type:"string" required:"true"`

	// Metadata which can be used to manage the topic rule.
	//
	// For URI Request parameters use format: ...key1=value1&key2=value2...
	//
	// For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	//
	// For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// The rule payload.
	//
	// TopicRulePayload is a required field
	TopicRulePayload *TopicRulePayload `locationName:"topicRulePayload" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateTopicRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTopicRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTopicRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if s.TopicRulePayload == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicRulePayload"))
	}
	if s.TopicRulePayload != nil {
		if err := s.TopicRulePayload.Validate(); err != nil {
			invalidParams.AddNested("TopicRulePayload", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTopicRuleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Tags != nil {
		v := *s.Tags

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-tagging", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RuleName != nil {
		v := *s.RuleName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ruleName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TopicRulePayload != nil {
		v := s.TopicRulePayload

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "topicRulePayload", v, metadata)
	}
	return nil
}

type CreateTopicRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTopicRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTopicRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateTopicRule = "CreateTopicRule"

// CreateTopicRuleRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a rule. Creating rules is an administrator-level action. Any user
// who has permission to create rules will be able to access data processed
// by the rule.
//
//    // Example sending a request using CreateTopicRuleRequest.
//    req := client.CreateTopicRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateTopicRuleRequest(input *CreateTopicRuleInput) CreateTopicRuleRequest {
	op := &aws.Operation{
		Name:       opCreateTopicRule,
		HTTPMethod: "POST",
		HTTPPath:   "/rules/{ruleName}",
	}

	if input == nil {
		input = &CreateTopicRuleInput{}
	}

	req := c.newRequest(op, input, &CreateTopicRuleOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return CreateTopicRuleRequest{Request: req, Input: input, Copy: c.CreateTopicRuleRequest}
}

// CreateTopicRuleRequest is the request type for the
// CreateTopicRule API operation.
type CreateTopicRuleRequest struct {
	*aws.Request
	Input *CreateTopicRuleInput
	Copy  func(*CreateTopicRuleInput) CreateTopicRuleRequest
}

// Send marshals and sends the CreateTopicRule API request.
func (r CreateTopicRuleRequest) Send(ctx context.Context) (*CreateTopicRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTopicRuleResponse{
		CreateTopicRuleOutput: r.Request.Data.(*CreateTopicRuleOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTopicRuleResponse is the response type for the
// CreateTopicRule API operation.
type CreateTopicRuleResponse struct {
	*CreateTopicRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTopicRule request.
func (r *CreateTopicRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
