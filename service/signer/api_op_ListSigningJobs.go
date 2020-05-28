// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package signer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListSigningJobsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the maximum number of items to return in the response. Use this
	// parameter when paginating results. If additional items exist beyond the number
	// you specify, the nextToken element is set in the response. Use the nextToken
	// value in a subsequent request to retrieve additional items.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// String for specifying the next set of paginated results to return. After
	// you receive a response with truncated results, use this parameter in a subsequent
	// request. Set it to the value of nextToken from the response that you just
	// received.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The ID of microcontroller platform that you specified for the distribution
	// of your code image.
	PlatformId *string `location:"querystring" locationName:"platformId" type:"string"`

	// The IAM principal that requested the signing job.
	RequestedBy *string `location:"querystring" locationName:"requestedBy" type:"string"`

	// A status value with which to filter your results.
	Status SigningStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListSigningJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSigningJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSigningJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSigningJobsInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.PlatformId != nil {
		v := *s.PlatformId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "platformId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestedBy != nil {
		v := *s.RequestedBy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "requestedBy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListSigningJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list of your signing jobs.
	Jobs []SigningJob `locationName:"jobs" type:"list"`

	// String for specifying the next set of paginated results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListSigningJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSigningJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Jobs != nil {
		v := s.Jobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "jobs", metadata)
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

const opListSigningJobs = "ListSigningJobs"

// ListSigningJobsRequest returns a request value for making API operation for
// AWS Signer.
//
// Lists all your signing jobs. You can use the maxResults parameter to limit
// the number of signing jobs that are returned in the response. If additional
// jobs remain to be listed, code signing returns a nextToken value. Use this
// value in subsequent calls to ListSigningJobs to fetch the remaining values.
// You can continue calling ListSigningJobs with your maxResults parameter and
// with new values that code signing returns in the nextToken parameter until
// all of your signing jobs have been returned.
//
//    // Example sending a request using ListSigningJobsRequest.
//    req := client.ListSigningJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/signer-2017-08-25/ListSigningJobs
func (c *Client) ListSigningJobsRequest(input *ListSigningJobsInput) ListSigningJobsRequest {
	op := &aws.Operation{
		Name:       opListSigningJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/signing-jobs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSigningJobsInput{}
	}

	req := c.newRequest(op, input, &ListSigningJobsOutput{})

	return ListSigningJobsRequest{Request: req, Input: input, Copy: c.ListSigningJobsRequest}
}

// ListSigningJobsRequest is the request type for the
// ListSigningJobs API operation.
type ListSigningJobsRequest struct {
	*aws.Request
	Input *ListSigningJobsInput
	Copy  func(*ListSigningJobsInput) ListSigningJobsRequest
}

// Send marshals and sends the ListSigningJobs API request.
func (r ListSigningJobsRequest) Send(ctx context.Context) (*ListSigningJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSigningJobsResponse{
		ListSigningJobsOutput: r.Request.Data.(*ListSigningJobsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSigningJobsRequestPaginator returns a paginator for ListSigningJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSigningJobsRequest(input)
//   p := signer.NewListSigningJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSigningJobsPaginator(req ListSigningJobsRequest) ListSigningJobsPaginator {
	return ListSigningJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSigningJobsInput
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

// ListSigningJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSigningJobsPaginator struct {
	aws.Pager
}

func (p *ListSigningJobsPaginator) CurrentPage() *ListSigningJobsOutput {
	return p.Pager.CurrentPage().(*ListSigningJobsOutput)
}

// ListSigningJobsResponse is the response type for the
// ListSigningJobs API operation.
type ListSigningJobsResponse struct {
	*ListSigningJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSigningJobs request.
func (r *ListSigningJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
