// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCommentsForComparedCommitInput struct {
	_ struct{} `type:"structure"`

	// To establish the directionality of the comparison, the full commit ID of
	// the after commit.
	//
	// AfterCommitId is a required field
	AfterCommitId *string `locationName:"afterCommitId" type:"string" required:"true"`

	// To establish the directionality of the comparison, the full commit ID of
	// the before commit.
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string"`

	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments, but you can configure up to 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// An enumeration token that when provided in a request, returns the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the repository where you want to compare commits.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCommentsForComparedCommitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCommentsForComparedCommitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCommentsForComparedCommitInput"}

	if s.AfterCommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AfterCommitId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCommentsForComparedCommitOutput struct {
	_ struct{} `type:"structure"`

	// A list of comment objects on the compared commit.
	CommentsForComparedCommitData []CommentsForComparedCommit `locationName:"commentsForComparedCommitData" type:"list"`

	// An enumeration token that can be used in a request to return the next batch
	// of the results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetCommentsForComparedCommitOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCommentsForComparedCommit = "GetCommentsForComparedCommit"

// GetCommentsForComparedCommitRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about comments made on the comparison between two commits.
//
//    // Example sending a request using GetCommentsForComparedCommitRequest.
//    req := client.GetCommentsForComparedCommitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetCommentsForComparedCommit
func (c *Client) GetCommentsForComparedCommitRequest(input *GetCommentsForComparedCommitInput) GetCommentsForComparedCommitRequest {
	op := &aws.Operation{
		Name:       opGetCommentsForComparedCommit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetCommentsForComparedCommitInput{}
	}

	req := c.newRequest(op, input, &GetCommentsForComparedCommitOutput{})

	return GetCommentsForComparedCommitRequest{Request: req, Input: input, Copy: c.GetCommentsForComparedCommitRequest}
}

// GetCommentsForComparedCommitRequest is the request type for the
// GetCommentsForComparedCommit API operation.
type GetCommentsForComparedCommitRequest struct {
	*aws.Request
	Input *GetCommentsForComparedCommitInput
	Copy  func(*GetCommentsForComparedCommitInput) GetCommentsForComparedCommitRequest
}

// Send marshals and sends the GetCommentsForComparedCommit API request.
func (r GetCommentsForComparedCommitRequest) Send(ctx context.Context) (*GetCommentsForComparedCommitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCommentsForComparedCommitResponse{
		GetCommentsForComparedCommitOutput: r.Request.Data.(*GetCommentsForComparedCommitOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetCommentsForComparedCommitRequestPaginator returns a paginator for GetCommentsForComparedCommit.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetCommentsForComparedCommitRequest(input)
//   p := codecommit.NewGetCommentsForComparedCommitRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetCommentsForComparedCommitPaginator(req GetCommentsForComparedCommitRequest) GetCommentsForComparedCommitPaginator {
	return GetCommentsForComparedCommitPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetCommentsForComparedCommitInput
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

// GetCommentsForComparedCommitPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetCommentsForComparedCommitPaginator struct {
	aws.Pager
}

func (p *GetCommentsForComparedCommitPaginator) CurrentPage() *GetCommentsForComparedCommitOutput {
	return p.Pager.CurrentPage().(*GetCommentsForComparedCommitOutput)
}

// GetCommentsForComparedCommitResponse is the response type for the
// GetCommentsForComparedCommit API operation.
type GetCommentsForComparedCommitResponse struct {
	*GetCommentsForComparedCommitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCommentsForComparedCommit request.
func (r *GetCommentsForComparedCommitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
