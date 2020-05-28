// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListPipelines action.
type ListPipelinesInput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous list pipelines call. It
	// can be used to return the next set of pipelines in the list.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListPipelinesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPipelinesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPipelinesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListPipelines action.
type ListPipelinesOutput struct {
	_ struct{} `type:"structure"`

	// If the amount of returned information is significantly large, an identifier
	// is also returned. It can be used in a subsequent list pipelines call to return
	// the next set of pipelines in the list.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The list of pipelines.
	Pipelines []PipelineSummary `locationName:"pipelines" type:"list"`
}

// String returns the string representation
func (s ListPipelinesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPipelines = "ListPipelines"

// ListPipelinesRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Gets a summary of all of the pipelines associated with your account.
//
//    // Example sending a request using ListPipelinesRequest.
//    req := client.ListPipelinesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/ListPipelines
func (c *Client) ListPipelinesRequest(input *ListPipelinesInput) ListPipelinesRequest {
	op := &aws.Operation{
		Name:       opListPipelines,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPipelinesInput{}
	}

	req := c.newRequest(op, input, &ListPipelinesOutput{})

	return ListPipelinesRequest{Request: req, Input: input, Copy: c.ListPipelinesRequest}
}

// ListPipelinesRequest is the request type for the
// ListPipelines API operation.
type ListPipelinesRequest struct {
	*aws.Request
	Input *ListPipelinesInput
	Copy  func(*ListPipelinesInput) ListPipelinesRequest
}

// Send marshals and sends the ListPipelines API request.
func (r ListPipelinesRequest) Send(ctx context.Context) (*ListPipelinesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPipelinesResponse{
		ListPipelinesOutput: r.Request.Data.(*ListPipelinesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPipelinesRequestPaginator returns a paginator for ListPipelines.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPipelinesRequest(input)
//   p := codepipeline.NewListPipelinesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPipelinesPaginator(req ListPipelinesRequest) ListPipelinesPaginator {
	return ListPipelinesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPipelinesInput
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

// ListPipelinesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPipelinesPaginator struct {
	aws.Pager
}

func (p *ListPipelinesPaginator) CurrentPage() *ListPipelinesOutput {
	return p.Pager.CurrentPage().(*ListPipelinesOutput)
}

// ListPipelinesResponse is the response type for the
// ListPipelines API operation.
type ListPipelinesResponse struct {
	*ListPipelinesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPipelines request.
func (r *ListPipelinesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
