// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListProjectsInput struct {
	_ struct{} `type:"structure"`

	// During a previous call, if there are more than 100 items in the list, only
	// the first 100 items are returned, along with a unique string called a nextToken.
	// To get the next batch of items in the list, call this operation again, adding
	// the next token to the call. To get all of the items in the list, keep calling
	// this operation with each subsequent next token that is returned, until no
	// more next tokens are returned.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The criterion to be used to list build project names. Valid values include:
	//
	//    * CREATED_TIME: List based on when each build project was created.
	//
	//    * LAST_MODIFIED_TIME: List based on when information about each build
	//    project was last changed.
	//
	//    * NAME: List based on each build project's name.
	//
	// Use sortOrder to specify in what order to list the build project names based
	// on the preceding criteria.
	SortBy ProjectSortByType `locationName:"sortBy" type:"string" enum:"true"`

	// The order in which to list build projects. Valid values include:
	//
	//    * ASCENDING: List in ascending order.
	//
	//    * DESCENDING: List in descending order.
	//
	// Use sortBy to specify the criterion to be used to list build project names.
	SortOrder SortOrderType `locationName:"sortOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProjectsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListProjectsOutput struct {
	_ struct{} `type:"structure"`

	// If there are more than 100 items in the list, only the first 100 items are
	// returned, along with a unique string called a nextToken. To get the next
	// batch of items in the list, call this operation again, adding the next token
	// to the call.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of build project names, with each build project name representing
	// a single build project.
	Projects []string `locationName:"projects" min:"1" type:"list"`
}

// String returns the string representation
func (s ListProjectsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListProjects = "ListProjects"

// ListProjectsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Gets a list of build project names, with each build project name representing
// a single build project.
//
//    // Example sending a request using ListProjectsRequest.
//    req := client.ListProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/ListProjects
func (c *Client) ListProjectsRequest(input *ListProjectsInput) ListProjectsRequest {
	op := &aws.Operation{
		Name:       opListProjects,
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
		input = &ListProjectsInput{}
	}

	req := c.newRequest(op, input, &ListProjectsOutput{})

	return ListProjectsRequest{Request: req, Input: input, Copy: c.ListProjectsRequest}
}

// ListProjectsRequest is the request type for the
// ListProjects API operation.
type ListProjectsRequest struct {
	*aws.Request
	Input *ListProjectsInput
	Copy  func(*ListProjectsInput) ListProjectsRequest
}

// Send marshals and sends the ListProjects API request.
func (r ListProjectsRequest) Send(ctx context.Context) (*ListProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProjectsResponse{
		ListProjectsOutput: r.Request.Data.(*ListProjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProjectsRequestPaginator returns a paginator for ListProjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProjectsRequest(input)
//   p := codebuild.NewListProjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProjectsPaginator(req ListProjectsRequest) ListProjectsPaginator {
	return ListProjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProjectsInput
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

// ListProjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProjectsPaginator struct {
	aws.Pager
}

func (p *ListProjectsPaginator) CurrentPage() *ListProjectsOutput {
	return p.Pager.CurrentPage().(*ListProjectsOutput)
}

// ListProjectsResponse is the response type for the
// ListProjects API operation.
type ListProjectsResponse struct {
	*ListProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProjects request.
func (r *ListProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
