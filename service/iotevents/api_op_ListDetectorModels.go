// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDetectorModelsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorModelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDetectorModelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDetectorModelsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorModelsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDetectorModelsOutput struct {
	_ struct{} `type:"structure"`

	// Summary information about the detector models.
	DetectorModelSummaries []DetectorModelSummary `locationName:"detectorModelSummaries" type:"list"`

	// A token to retrieve the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDetectorModelsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDetectorModelsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DetectorModelSummaries != nil {
		v := s.DetectorModelSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "detectorModelSummaries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDetectorModels = "ListDetectorModels"

// ListDetectorModelsRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Lists the detector models you have created. Only the metadata associated
// with each detector model is returned.
//
//    // Example sending a request using ListDetectorModelsRequest.
//    req := client.ListDetectorModelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/ListDetectorModels
func (c *Client) ListDetectorModelsRequest(input *ListDetectorModelsInput) ListDetectorModelsRequest {
	op := &aws.Operation{
		Name:       opListDetectorModels,
		HTTPMethod: "GET",
		HTTPPath:   "/detector-models",
	}

	if input == nil {
		input = &ListDetectorModelsInput{}
	}

	req := c.newRequest(op, input, &ListDetectorModelsOutput{})

	return ListDetectorModelsRequest{Request: req, Input: input, Copy: c.ListDetectorModelsRequest}
}

// ListDetectorModelsRequest is the request type for the
// ListDetectorModels API operation.
type ListDetectorModelsRequest struct {
	*aws.Request
	Input *ListDetectorModelsInput
	Copy  func(*ListDetectorModelsInput) ListDetectorModelsRequest
}

// Send marshals and sends the ListDetectorModels API request.
func (r ListDetectorModelsRequest) Send(ctx context.Context) (*ListDetectorModelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDetectorModelsResponse{
		ListDetectorModelsOutput: r.Request.Data.(*ListDetectorModelsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDetectorModelsResponse is the response type for the
// ListDetectorModels API operation.
type ListDetectorModelsResponse struct {
	*ListDetectorModelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDetectorModels request.
func (r *ListDetectorModelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
