// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDatasetsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in this request.
	//
	// The default value is 100.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDatasetsInput) MarshalFields(e protocol.FieldEncoder) error {
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

type ListDatasetsOutput struct {
	_ struct{} `type:"structure"`

	// A list of "DatasetSummary" objects.
	DatasetSummaries []DatasetSummary `locationName:"datasetSummaries" type:"list"`

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatasetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDatasetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DatasetSummaries != nil {
		v := s.DatasetSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "datasetSummaries", metadata)
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

const opListDatasets = "ListDatasets"

// ListDatasetsRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about data sets.
//
//    // Example sending a request using ListDatasetsRequest.
//    req := client.ListDatasetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/ListDatasets
func (c *Client) ListDatasetsRequest(input *ListDatasetsInput) ListDatasetsRequest {
	op := &aws.Operation{
		Name:       opListDatasets,
		HTTPMethod: "GET",
		HTTPPath:   "/datasets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDatasetsInput{}
	}

	req := c.newRequest(op, input, &ListDatasetsOutput{})

	return ListDatasetsRequest{Request: req, Input: input, Copy: c.ListDatasetsRequest}
}

// ListDatasetsRequest is the request type for the
// ListDatasets API operation.
type ListDatasetsRequest struct {
	*aws.Request
	Input *ListDatasetsInput
	Copy  func(*ListDatasetsInput) ListDatasetsRequest
}

// Send marshals and sends the ListDatasets API request.
func (r ListDatasetsRequest) Send(ctx context.Context) (*ListDatasetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDatasetsResponse{
		ListDatasetsOutput: r.Request.Data.(*ListDatasetsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDatasetsRequestPaginator returns a paginator for ListDatasets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDatasetsRequest(input)
//   p := iotanalytics.NewListDatasetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDatasetsPaginator(req ListDatasetsRequest) ListDatasetsPaginator {
	return ListDatasetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDatasetsInput
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

// ListDatasetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDatasetsPaginator struct {
	aws.Pager
}

func (p *ListDatasetsPaginator) CurrentPage() *ListDatasetsOutput {
	return p.Pager.CurrentPage().(*ListDatasetsOutput)
}

// ListDatasetsResponse is the response type for the
// ListDatasets API operation.
type ListDatasetsResponse struct {
	*ListDatasetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDatasets request.
func (r *ListDatasetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
