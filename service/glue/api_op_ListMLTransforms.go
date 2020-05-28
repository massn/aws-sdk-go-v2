// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListMLTransformsInput struct {
	_ struct{} `type:"structure"`

	// A TransformFilterCriteria used to filter the machine learning transforms.
	Filter *TransformFilterCriteria `type:"structure"`

	// The maximum size of a list to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation request.
	NextToken *string `type:"string"`

	// A TransformSortCriteria used to sort the machine learning transforms.
	Sort *TransformSortCriteria `type:"structure"`

	// Specifies to return only these tagged resources.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s ListMLTransformsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMLTransformsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMLTransformsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}
	if s.Sort != nil {
		if err := s.Sort.Validate(); err != nil {
			invalidParams.AddNested("Sort", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListMLTransformsOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string `type:"string"`

	// The identifiers of all the machine learning transforms in the account, or
	// the machine learning transforms with the specified tags.
	//
	// TransformIds is a required field
	TransformIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s ListMLTransformsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListMLTransforms = "ListMLTransforms"

// ListMLTransformsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a sortable, filterable list of existing AWS Glue machine learning
// transforms in this AWS account, or the resources with the specified tag.
// This operation takes the optional Tags field, which you can use as a filter
// of the responses so that tagged resources can be retrieved as a group. If
// you choose to use tag filtering, only resources with the tags are retrieved.
//
//    // Example sending a request using ListMLTransformsRequest.
//    req := client.ListMLTransformsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/ListMLTransforms
func (c *Client) ListMLTransformsRequest(input *ListMLTransformsInput) ListMLTransformsRequest {
	op := &aws.Operation{
		Name:       opListMLTransforms,
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
		input = &ListMLTransformsInput{}
	}

	req := c.newRequest(op, input, &ListMLTransformsOutput{})

	return ListMLTransformsRequest{Request: req, Input: input, Copy: c.ListMLTransformsRequest}
}

// ListMLTransformsRequest is the request type for the
// ListMLTransforms API operation.
type ListMLTransformsRequest struct {
	*aws.Request
	Input *ListMLTransformsInput
	Copy  func(*ListMLTransformsInput) ListMLTransformsRequest
}

// Send marshals and sends the ListMLTransforms API request.
func (r ListMLTransformsRequest) Send(ctx context.Context) (*ListMLTransformsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMLTransformsResponse{
		ListMLTransformsOutput: r.Request.Data.(*ListMLTransformsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMLTransformsRequestPaginator returns a paginator for ListMLTransforms.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMLTransformsRequest(input)
//   p := glue.NewListMLTransformsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMLTransformsPaginator(req ListMLTransformsRequest) ListMLTransformsPaginator {
	return ListMLTransformsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMLTransformsInput
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

// ListMLTransformsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMLTransformsPaginator struct {
	aws.Pager
}

func (p *ListMLTransformsPaginator) CurrentPage() *ListMLTransformsOutput {
	return p.Pager.CurrentPage().(*ListMLTransformsOutput)
}

// ListMLTransformsResponse is the response type for the
// ListMLTransforms API operation.
type ListMLTransformsResponse struct {
	*ListMLTransformsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMLTransforms request.
func (r *ListMLTransformsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
