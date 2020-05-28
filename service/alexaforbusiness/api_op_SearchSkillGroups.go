// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchSkillGroupsInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to list a specified set of skill groups. The supported
	// filter key is SkillGroupName.
	Filters []Filter `type:"list"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	// Required.
	NextToken *string `min:"1" type:"string"`

	// The sort order to use in listing the specified set of skill groups. The supported
	// sort key is SkillGroupName.
	SortCriteria []Sort `type:"list"`
}

// String returns the string representation
func (s SearchSkillGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchSkillGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchSkillGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SortCriteria != nil {
		for i, v := range s.SortCriteria {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SortCriteria", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchSkillGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The token returned to indicate that there is more data available.
	NextToken *string `min:"1" type:"string"`

	// The skill groups that meet the filter criteria, in sort order.
	SkillGroups []SkillGroupData `type:"list"`

	// The total number of skill groups returned.
	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s SearchSkillGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchSkillGroups = "SearchSkillGroups"

// SearchSkillGroupsRequest returns a request value for making API operation for
// Alexa For Business.
//
// Searches skill groups and lists the ones that meet a set of filter and sort
// criteria.
//
//    // Example sending a request using SearchSkillGroupsRequest.
//    req := client.SearchSkillGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SearchSkillGroups
func (c *Client) SearchSkillGroupsRequest(input *SearchSkillGroupsInput) SearchSkillGroupsRequest {
	op := &aws.Operation{
		Name:       opSearchSkillGroups,
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
		input = &SearchSkillGroupsInput{}
	}

	req := c.newRequest(op, input, &SearchSkillGroupsOutput{})

	return SearchSkillGroupsRequest{Request: req, Input: input, Copy: c.SearchSkillGroupsRequest}
}

// SearchSkillGroupsRequest is the request type for the
// SearchSkillGroups API operation.
type SearchSkillGroupsRequest struct {
	*aws.Request
	Input *SearchSkillGroupsInput
	Copy  func(*SearchSkillGroupsInput) SearchSkillGroupsRequest
}

// Send marshals and sends the SearchSkillGroups API request.
func (r SearchSkillGroupsRequest) Send(ctx context.Context) (*SearchSkillGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchSkillGroupsResponse{
		SearchSkillGroupsOutput: r.Request.Data.(*SearchSkillGroupsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchSkillGroupsRequestPaginator returns a paginator for SearchSkillGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchSkillGroupsRequest(input)
//   p := alexaforbusiness.NewSearchSkillGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchSkillGroupsPaginator(req SearchSkillGroupsRequest) SearchSkillGroupsPaginator {
	return SearchSkillGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchSkillGroupsInput
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

// SearchSkillGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchSkillGroupsPaginator struct {
	aws.Pager
}

func (p *SearchSkillGroupsPaginator) CurrentPage() *SearchSkillGroupsOutput {
	return p.Pager.CurrentPage().(*SearchSkillGroupsOutput)
}

// SearchSkillGroupsResponse is the response type for the
// SearchSkillGroups API operation.
type SearchSkillGroupsResponse struct {
	*SearchSkillGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchSkillGroups request.
func (r *SearchSkillGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
