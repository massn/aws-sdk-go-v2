// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDominantLanguageDetectionJobsInput struct {
	_ struct{} `type:"structure"`

	// Filters that jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter
	// at a time.
	Filter *DominantLanguageDetectionJobFilter `type:"structure"`

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDominantLanguageDetectionJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDominantLanguageDetectionJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDominantLanguageDetectionJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDominantLanguageDetectionJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list containing the properties of each job that is returned.
	DominantLanguageDetectionJobPropertiesList []DominantLanguageDetectionJobProperties `type:"list"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDominantLanguageDetectionJobsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDominantLanguageDetectionJobs = "ListDominantLanguageDetectionJobs"

// ListDominantLanguageDetectionJobsRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets a list of the dominant language detection jobs that you have submitted.
//
//    // Example sending a request using ListDominantLanguageDetectionJobsRequest.
//    req := client.ListDominantLanguageDetectionJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListDominantLanguageDetectionJobs
func (c *Client) ListDominantLanguageDetectionJobsRequest(input *ListDominantLanguageDetectionJobsInput) ListDominantLanguageDetectionJobsRequest {
	op := &aws.Operation{
		Name:       opListDominantLanguageDetectionJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDominantLanguageDetectionJobsInput{}
	}

	req := c.newRequest(op, input, &ListDominantLanguageDetectionJobsOutput{})

	return ListDominantLanguageDetectionJobsRequest{Request: req, Input: input, Copy: c.ListDominantLanguageDetectionJobsRequest}
}

// ListDominantLanguageDetectionJobsRequest is the request type for the
// ListDominantLanguageDetectionJobs API operation.
type ListDominantLanguageDetectionJobsRequest struct {
	*aws.Request
	Input *ListDominantLanguageDetectionJobsInput
	Copy  func(*ListDominantLanguageDetectionJobsInput) ListDominantLanguageDetectionJobsRequest
}

// Send marshals and sends the ListDominantLanguageDetectionJobs API request.
func (r ListDominantLanguageDetectionJobsRequest) Send(ctx context.Context) (*ListDominantLanguageDetectionJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDominantLanguageDetectionJobsResponse{
		ListDominantLanguageDetectionJobsOutput: r.Request.Data.(*ListDominantLanguageDetectionJobsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDominantLanguageDetectionJobsRequestPaginator returns a paginator for ListDominantLanguageDetectionJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDominantLanguageDetectionJobsRequest(input)
//   p := comprehend.NewListDominantLanguageDetectionJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDominantLanguageDetectionJobsPaginator(req ListDominantLanguageDetectionJobsRequest) ListDominantLanguageDetectionJobsPaginator {
	return ListDominantLanguageDetectionJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDominantLanguageDetectionJobsInput
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

// ListDominantLanguageDetectionJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDominantLanguageDetectionJobsPaginator struct {
	aws.Pager
}

func (p *ListDominantLanguageDetectionJobsPaginator) CurrentPage() *ListDominantLanguageDetectionJobsOutput {
	return p.Pager.CurrentPage().(*ListDominantLanguageDetectionJobsOutput)
}

// ListDominantLanguageDetectionJobsResponse is the response type for the
// ListDominantLanguageDetectionJobs API operation.
type ListDominantLanguageDetectionJobsResponse struct {
	*ListDominantLanguageDetectionJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDominantLanguageDetectionJobs request.
func (r *ListDominantLanguageDetectionJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
