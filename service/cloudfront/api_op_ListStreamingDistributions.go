// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to list your streaming distributions.
type ListStreamingDistributionsInput struct {
	_ struct{} `type:"structure"`

	// The value that you provided for the Marker request parameter.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// The value that you provided for the MaxItems request parameter.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`
}

// String returns the string representation
func (s ListStreamingDistributionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListStreamingDistributionsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// The returned result of the corresponding request.
type ListStreamingDistributionsOutput struct {
	_ struct{} `type:"structure" payload:"StreamingDistributionList"`

	// The StreamingDistributionList type.
	StreamingDistributionList *StreamingDistributionList `type:"structure"`
}

// String returns the string representation
func (s ListStreamingDistributionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListStreamingDistributionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.StreamingDistributionList != nil {
		v := s.StreamingDistributionList

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "StreamingDistributionList", v, metadata)
	}
	return nil
}

const opListStreamingDistributions = "ListStreamingDistributions2019_03_26"

// ListStreamingDistributionsRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// List streaming distributions.
//
//    // Example sending a request using ListStreamingDistributionsRequest.
//    req := client.ListStreamingDistributionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/ListStreamingDistributions
func (c *Client) ListStreamingDistributionsRequest(input *ListStreamingDistributionsInput) ListStreamingDistributionsRequest {
	op := &aws.Operation{
		Name:       opListStreamingDistributions,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-03-26/streaming-distribution",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"StreamingDistributionList.NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "StreamingDistributionList.IsTruncated",
		},
	}

	if input == nil {
		input = &ListStreamingDistributionsInput{}
	}

	req := c.newRequest(op, input, &ListStreamingDistributionsOutput{})

	return ListStreamingDistributionsRequest{Request: req, Input: input, Copy: c.ListStreamingDistributionsRequest}
}

// ListStreamingDistributionsRequest is the request type for the
// ListStreamingDistributions API operation.
type ListStreamingDistributionsRequest struct {
	*aws.Request
	Input *ListStreamingDistributionsInput
	Copy  func(*ListStreamingDistributionsInput) ListStreamingDistributionsRequest
}

// Send marshals and sends the ListStreamingDistributions API request.
func (r ListStreamingDistributionsRequest) Send(ctx context.Context) (*ListStreamingDistributionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStreamingDistributionsResponse{
		ListStreamingDistributionsOutput: r.Request.Data.(*ListStreamingDistributionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListStreamingDistributionsRequestPaginator returns a paginator for ListStreamingDistributions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListStreamingDistributionsRequest(input)
//   p := cloudfront.NewListStreamingDistributionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListStreamingDistributionsPaginator(req ListStreamingDistributionsRequest) ListStreamingDistributionsPaginator {
	return ListStreamingDistributionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListStreamingDistributionsInput
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

// ListStreamingDistributionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListStreamingDistributionsPaginator struct {
	aws.Pager
}

func (p *ListStreamingDistributionsPaginator) CurrentPage() *ListStreamingDistributionsOutput {
	return p.Pager.CurrentPage().(*ListStreamingDistributionsOutput)
}

// ListStreamingDistributionsResponse is the response type for the
// ListStreamingDistributions API operation.
type ListStreamingDistributionsResponse struct {
	*ListStreamingDistributionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStreamingDistributions request.
func (r *ListStreamingDistributionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
