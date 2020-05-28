// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListConfigurationsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListConfigurationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {
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

// The response contains an array of Configuration and a next token if the response
// is truncated.
type ListConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// An array of MSK configurations.
	Configurations []Configuration `locationName:"configurations" type:"list"`

	// The paginated results marker. When the result of a ListConfigurations operation
	// is truncated, the call returns NextToken in the response. To get another
	// batch of configurations, provide this token in your next request.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Configurations != nil {
		v := s.Configurations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "configurations", metadata)
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

const opListConfigurations = "ListConfigurations"

// ListConfigurationsRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a list of all the MSK configurations in this Region.
//
//    // Example sending a request using ListConfigurationsRequest.
//    req := client.ListConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListConfigurations
func (c *Client) ListConfigurationsRequest(input *ListConfigurationsInput) ListConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/configurations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListConfigurationsOutput{})

	return ListConfigurationsRequest{Request: req, Input: input, Copy: c.ListConfigurationsRequest}
}

// ListConfigurationsRequest is the request type for the
// ListConfigurations API operation.
type ListConfigurationsRequest struct {
	*aws.Request
	Input *ListConfigurationsInput
	Copy  func(*ListConfigurationsInput) ListConfigurationsRequest
}

// Send marshals and sends the ListConfigurations API request.
func (r ListConfigurationsRequest) Send(ctx context.Context) (*ListConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConfigurationsResponse{
		ListConfigurationsOutput: r.Request.Data.(*ListConfigurationsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListConfigurationsRequestPaginator returns a paginator for ListConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListConfigurationsRequest(input)
//   p := kafka.NewListConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListConfigurationsPaginator(req ListConfigurationsRequest) ListConfigurationsPaginator {
	return ListConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListConfigurationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListConfigurationsPaginator struct {
	aws.Pager
}

func (p *ListConfigurationsPaginator) CurrentPage() *ListConfigurationsOutput {
	return p.Pager.CurrentPage().(*ListConfigurationsOutput)
}

// ListConfigurationsResponse is the response type for the
// ListConfigurations API operation.
type ListConfigurationsResponse struct {
	*ListConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConfigurations request.
func (r *ListConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
