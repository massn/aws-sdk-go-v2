// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchRoomsInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to list a specified set of rooms. The supported filter
	// keys are RoomName and ProfileName.
	Filters []Filter `type:"list"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `min:"1" type:"string"`

	// The sort order to use in listing the specified set of rooms. The supported
	// sort keys are RoomName and ProfileName.
	SortCriteria []Sort `type:"list"`
}

// String returns the string representation
func (s SearchRoomsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchRoomsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchRoomsInput"}
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

type SearchRoomsOutput struct {
	_ struct{} `type:"structure"`

	// The token returned to indicate that there is more data available.
	NextToken *string `min:"1" type:"string"`

	// The rooms that meet the specified set of filter criteria, in sort order.
	Rooms []RoomData `type:"list"`

	// The total number of rooms returned.
	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s SearchRoomsOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchRooms = "SearchRooms"

// SearchRoomsRequest returns a request value for making API operation for
// Alexa For Business.
//
// Searches rooms and lists the ones that meet a set of filter and sort criteria.
//
//    // Example sending a request using SearchRoomsRequest.
//    req := client.SearchRoomsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SearchRooms
func (c *Client) SearchRoomsRequest(input *SearchRoomsInput) SearchRoomsRequest {
	op := &aws.Operation{
		Name:       opSearchRooms,
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
		input = &SearchRoomsInput{}
	}

	req := c.newRequest(op, input, &SearchRoomsOutput{})

	return SearchRoomsRequest{Request: req, Input: input, Copy: c.SearchRoomsRequest}
}

// SearchRoomsRequest is the request type for the
// SearchRooms API operation.
type SearchRoomsRequest struct {
	*aws.Request
	Input *SearchRoomsInput
	Copy  func(*SearchRoomsInput) SearchRoomsRequest
}

// Send marshals and sends the SearchRooms API request.
func (r SearchRoomsRequest) Send(ctx context.Context) (*SearchRoomsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchRoomsResponse{
		SearchRoomsOutput: r.Request.Data.(*SearchRoomsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchRoomsRequestPaginator returns a paginator for SearchRooms.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchRoomsRequest(input)
//   p := alexaforbusiness.NewSearchRoomsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchRoomsPaginator(req SearchRoomsRequest) SearchRoomsPaginator {
	return SearchRoomsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchRoomsInput
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

// SearchRoomsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchRoomsPaginator struct {
	aws.Pager
}

func (p *SearchRoomsPaginator) CurrentPage() *SearchRoomsOutput {
	return p.Pager.CurrentPage().(*SearchRoomsOutput)
}

// SearchRoomsResponse is the response type for the
// SearchRooms API operation.
type SearchRoomsResponse struct {
	*SearchRoomsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchRooms request.
func (r *SearchRoomsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
