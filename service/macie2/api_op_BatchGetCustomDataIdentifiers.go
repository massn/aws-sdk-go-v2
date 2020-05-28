// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Specifies one or more custom data identifiers to retrieve information about.
type BatchGetCustomDataIdentifiersInput struct {
	_ struct{} `type:"structure"`

	Ids []string `locationName:"ids" type:"list"`
}

// String returns the string representation
func (s BatchGetCustomDataIdentifiersInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchGetCustomDataIdentifiersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Ids != nil {
		v := s.Ids

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ids", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Provides information about one or more custom data identifiers.
type BatchGetCustomDataIdentifiersOutput struct {
	_ struct{} `type:"structure"`

	CustomDataIdentifiers []BatchGetCustomDataIdentifierSummary `locationName:"customDataIdentifiers" type:"list"`

	NotFoundIdentifierIds []string `locationName:"notFoundIdentifierIds" type:"list"`
}

// String returns the string representation
func (s BatchGetCustomDataIdentifiersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchGetCustomDataIdentifiersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CustomDataIdentifiers != nil {
		v := s.CustomDataIdentifiers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "customDataIdentifiers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NotFoundIdentifierIds != nil {
		v := s.NotFoundIdentifierIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "notFoundIdentifierIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opBatchGetCustomDataIdentifiers = "BatchGetCustomDataIdentifiers"

// BatchGetCustomDataIdentifiersRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves information about one or more custom data identifiers.
//
//    // Example sending a request using BatchGetCustomDataIdentifiersRequest.
//    req := client.BatchGetCustomDataIdentifiersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/BatchGetCustomDataIdentifiers
func (c *Client) BatchGetCustomDataIdentifiersRequest(input *BatchGetCustomDataIdentifiersInput) BatchGetCustomDataIdentifiersRequest {
	op := &aws.Operation{
		Name:       opBatchGetCustomDataIdentifiers,
		HTTPMethod: "POST",
		HTTPPath:   "/custom-data-identifiers/get",
	}

	if input == nil {
		input = &BatchGetCustomDataIdentifiersInput{}
	}

	req := c.newRequest(op, input, &BatchGetCustomDataIdentifiersOutput{})

	return BatchGetCustomDataIdentifiersRequest{Request: req, Input: input, Copy: c.BatchGetCustomDataIdentifiersRequest}
}

// BatchGetCustomDataIdentifiersRequest is the request type for the
// BatchGetCustomDataIdentifiers API operation.
type BatchGetCustomDataIdentifiersRequest struct {
	*aws.Request
	Input *BatchGetCustomDataIdentifiersInput
	Copy  func(*BatchGetCustomDataIdentifiersInput) BatchGetCustomDataIdentifiersRequest
}

// Send marshals and sends the BatchGetCustomDataIdentifiers API request.
func (r BatchGetCustomDataIdentifiersRequest) Send(ctx context.Context) (*BatchGetCustomDataIdentifiersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetCustomDataIdentifiersResponse{
		BatchGetCustomDataIdentifiersOutput: r.Request.Data.(*BatchGetCustomDataIdentifiersOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetCustomDataIdentifiersResponse is the response type for the
// BatchGetCustomDataIdentifiers API operation.
type BatchGetCustomDataIdentifiersResponse struct {
	*BatchGetCustomDataIdentifiersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetCustomDataIdentifiers request.
func (r *BatchGetCustomDataIdentifiersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
