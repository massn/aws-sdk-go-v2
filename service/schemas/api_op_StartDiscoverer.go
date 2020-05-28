// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StartDiscovererInput struct {
	_ struct{} `type:"structure"`

	// DiscovererId is a required field
	DiscovererId *string `location:"uri" locationName:"discovererId" type:"string" required:"true"`
}

// String returns the string representation
func (s StartDiscovererInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDiscovererInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDiscovererInput"}

	if s.DiscovererId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiscovererId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartDiscovererInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DiscovererId != nil {
		v := *s.DiscovererId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "discovererId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartDiscovererOutput struct {
	_ struct{} `type:"structure"`

	DiscovererId *string `type:"string"`

	State DiscovererState `type:"string" enum:"true"`
}

// String returns the string representation
func (s StartDiscovererOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartDiscovererOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DiscovererId != nil {
		v := *s.DiscovererId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DiscovererId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opStartDiscoverer = "StartDiscoverer"

// StartDiscovererRequest returns a request value for making API operation for
// Schemas.
//
// Starts the discoverer
//
//    // Example sending a request using StartDiscovererRequest.
//    req := client.StartDiscovererRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/StartDiscoverer
func (c *Client) StartDiscovererRequest(input *StartDiscovererInput) StartDiscovererRequest {
	op := &aws.Operation{
		Name:       opStartDiscoverer,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/discoverers/id/{discovererId}/start",
	}

	if input == nil {
		input = &StartDiscovererInput{}
	}

	req := c.newRequest(op, input, &StartDiscovererOutput{})

	return StartDiscovererRequest{Request: req, Input: input, Copy: c.StartDiscovererRequest}
}

// StartDiscovererRequest is the request type for the
// StartDiscoverer API operation.
type StartDiscovererRequest struct {
	*aws.Request
	Input *StartDiscovererInput
	Copy  func(*StartDiscovererInput) StartDiscovererRequest
}

// Send marshals and sends the StartDiscoverer API request.
func (r StartDiscovererRequest) Send(ctx context.Context) (*StartDiscovererResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDiscovererResponse{
		StartDiscovererOutput: r.Request.Data.(*StartDiscovererOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDiscovererResponse is the response type for the
// StartDiscoverer API operation.
type StartDiscovererResponse struct {
	*StartDiscovererOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDiscoverer request.
func (r *StartDiscovererResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
