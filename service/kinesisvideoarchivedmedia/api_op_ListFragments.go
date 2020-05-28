// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListFragmentsInput struct {
	_ struct{} `type:"structure"`

	// Describes the timestamp range and timestamp origin for the range of fragments
	// to return.
	FragmentSelector *FragmentSelector `type:"structure"`

	// The total number of fragments to return. If the total number of fragments
	// available is more than the value specified in max-results, then a ListFragmentsOutput$NextToken
	// is provided in the output that you can use to resume pagination.
	MaxResults *int64 `min:"1" type:"long"`

	// A token to specify where to start paginating. This is the ListFragmentsOutput$NextToken
	// from a previously truncated response.
	NextToken *string `min:"1" type:"string"`

	// The name of the stream from which to retrieve a fragment list.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListFragmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFragmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFragmentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}
	if s.FragmentSelector != nil {
		if err := s.FragmentSelector.Validate(); err != nil {
			invalidParams.AddNested("FragmentSelector", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFragmentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FragmentSelector != nil {
		v := s.FragmentSelector

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "FragmentSelector", v, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListFragmentsOutput struct {
	_ struct{} `type:"structure"`

	// A list of archived Fragment objects from the stream that meet the selector
	// criteria. Results are in no specific order, even across pages.
	Fragments []Fragment `type:"list"`

	// If the returned list is truncated, the operation returns this token to use
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFragmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFragmentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Fragments != nil {
		v := s.Fragments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Fragments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFragments = "ListFragments"

// ListFragmentsRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams Archived Media.
//
// Returns a list of Fragment objects from the specified stream and timestamp
// range within the archived data.
//
// Listing fragments is eventually consistent. This means that even if the producer
// receives an acknowledgment that a fragment is persisted, the result might
// not be returned immediately from a request to ListFragments. However, results
// are typically available in less than one second.
//
// You must first call the GetDataEndpoint API to get an endpoint. Then send
// the ListFragments requests to this endpoint using the --endpoint-url parameter
// (https://docs.aws.amazon.com/cli/latest/reference/).
//
// If an error is thrown after invoking a Kinesis Video Streams archived media
// API, in addition to the HTTP status code and the response body, it includes
// the following pieces of information:
//
//    * x-amz-ErrorType HTTP header – contains a more specific error type
//    in addition to what the HTTP status code provides.
//
//    * x-amz-RequestId HTTP header – if you want to report an issue to AWS,
//    the support team can better diagnose the problem if given the Request
//    Id.
//
// Both the HTTP status code and the ErrorType header can be utilized to make
// programmatic decisions about whether errors are retry-able and under what
// conditions, as well as provide information on what actions the client programmer
// might need to take in order to successfully try again.
//
// For more information, see the Errors section at the bottom of this topic,
// as well as Common Errors (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html).
//
//    // Example sending a request using ListFragmentsRequest.
//    req := client.ListFragmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/ListFragments
func (c *Client) ListFragmentsRequest(input *ListFragmentsInput) ListFragmentsRequest {
	op := &aws.Operation{
		Name:       opListFragments,
		HTTPMethod: "POST",
		HTTPPath:   "/listFragments",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFragmentsInput{}
	}

	req := c.newRequest(op, input, &ListFragmentsOutput{})

	return ListFragmentsRequest{Request: req, Input: input, Copy: c.ListFragmentsRequest}
}

// ListFragmentsRequest is the request type for the
// ListFragments API operation.
type ListFragmentsRequest struct {
	*aws.Request
	Input *ListFragmentsInput
	Copy  func(*ListFragmentsInput) ListFragmentsRequest
}

// Send marshals and sends the ListFragments API request.
func (r ListFragmentsRequest) Send(ctx context.Context) (*ListFragmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFragmentsResponse{
		ListFragmentsOutput: r.Request.Data.(*ListFragmentsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFragmentsRequestPaginator returns a paginator for ListFragments.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFragmentsRequest(input)
//   p := kinesisvideoarchivedmedia.NewListFragmentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFragmentsPaginator(req ListFragmentsRequest) ListFragmentsPaginator {
	return ListFragmentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFragmentsInput
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

// ListFragmentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFragmentsPaginator struct {
	aws.Pager
}

func (p *ListFragmentsPaginator) CurrentPage() *ListFragmentsOutput {
	return p.Pager.CurrentPage().(*ListFragmentsOutput)
}

// ListFragmentsResponse is the response type for the
// ListFragments API operation.
type ListFragmentsResponse struct {
	*ListFragmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFragments request.
func (r *ListFragmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
