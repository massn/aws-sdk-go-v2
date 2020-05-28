// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListQualificationTypesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `min:"1" type:"integer"`

	// Specifies that only Qualification types that the Requester created are returned.
	// If false, the operation returns all Qualification types.
	MustBeOwnedByCaller *bool `type:"boolean"`

	// Specifies that only Qualification types that a user can request through the
	// Amazon Mechanical Turk web site, such as by taking a Qualification test,
	// are returned as results of the search. Some Qualification types, such as
	// those assigned automatically by the system, cannot be requested directly
	// by users. If false, all Qualification types, including those managed by the
	// system, are considered. Valid values are True | False.
	//
	// MustBeRequestable is a required field
	MustBeRequestable *bool `type:"boolean" required:"true"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// A text query against all of the searchable attributes of Qualification types.
	Query *string `type:"string"`
}

// String returns the string representation
func (s ListQualificationTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListQualificationTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListQualificationTypesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.MustBeRequestable == nil {
		invalidParams.Add(aws.NewErrParamRequired("MustBeRequestable"))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListQualificationTypesOutput struct {
	_ struct{} `type:"structure"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The number of Qualification types on this page in the filtered results list,
	// equivalent to the number of types this operation returns.
	NumResults *int64 `type:"integer"`

	// The list of QualificationType elements returned by the query.
	QualificationTypes []QualificationType `type:"list"`
}

// String returns the string representation
func (s ListQualificationTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListQualificationTypes = "ListQualificationTypes"

// ListQualificationTypesRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The ListQualificationTypes operation returns a list of Qualification types,
// filtered by an optional search term.
//
//    // Example sending a request using ListQualificationTypesRequest.
//    req := client.ListQualificationTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListQualificationTypes
func (c *Client) ListQualificationTypesRequest(input *ListQualificationTypesInput) ListQualificationTypesRequest {
	op := &aws.Operation{
		Name:       opListQualificationTypes,
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
		input = &ListQualificationTypesInput{}
	}

	req := c.newRequest(op, input, &ListQualificationTypesOutput{})

	return ListQualificationTypesRequest{Request: req, Input: input, Copy: c.ListQualificationTypesRequest}
}

// ListQualificationTypesRequest is the request type for the
// ListQualificationTypes API operation.
type ListQualificationTypesRequest struct {
	*aws.Request
	Input *ListQualificationTypesInput
	Copy  func(*ListQualificationTypesInput) ListQualificationTypesRequest
}

// Send marshals and sends the ListQualificationTypes API request.
func (r ListQualificationTypesRequest) Send(ctx context.Context) (*ListQualificationTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListQualificationTypesResponse{
		ListQualificationTypesOutput: r.Request.Data.(*ListQualificationTypesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListQualificationTypesRequestPaginator returns a paginator for ListQualificationTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListQualificationTypesRequest(input)
//   p := mturk.NewListQualificationTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListQualificationTypesPaginator(req ListQualificationTypesRequest) ListQualificationTypesPaginator {
	return ListQualificationTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListQualificationTypesInput
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

// ListQualificationTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListQualificationTypesPaginator struct {
	aws.Pager
}

func (p *ListQualificationTypesPaginator) CurrentPage() *ListQualificationTypesOutput {
	return p.Pager.CurrentPage().(*ListQualificationTypesOutput)
}

// ListQualificationTypesResponse is the response type for the
// ListQualificationTypes API operation.
type ListQualificationTypesResponse struct {
	*ListQualificationTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListQualificationTypes request.
func (r *ListQualificationTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
