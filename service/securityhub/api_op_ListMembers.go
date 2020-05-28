// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListMembersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return in the response.
	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	// The token that is required for pagination. On your first call to the ListMembers
	// operation, set the value of this parameter to NULL.
	//
	// For subsequent calls to the operation, to continue listing data, set the
	// value of this parameter to the value returned from the previous response.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// Specifies which member accounts to include in the response based on their
	// relationship status with the master account. The default value is TRUE.
	//
	// If OnlyAssociated is set to TRUE, the response includes member accounts whose
	// relationship status with the master is set to ENABLED or DISABLED.
	//
	// If OnlyAssociated is set to FALSE, the response includes all existing member
	// accounts.
	OnlyAssociated *bool `location:"querystring" locationName:"OnlyAssociated" type:"boolean"`
}

// String returns the string representation
func (s ListMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMembersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OnlyAssociated != nil {
		v := *s.OnlyAssociated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "OnlyAssociated", protocol.BoolValue(v), metadata)
	}
	return nil
}

type ListMembersOutput struct {
	_ struct{} `type:"structure"`

	// Member details returned by the operation.
	Members []Member `type:"list"`

	// The pagination token to use to request the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Members != nil {
		v := s.Members

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Members", metadata)
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

const opListMembers = "ListMembers"

// ListMembersRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Lists details about all member accounts for the current Security Hub master
// account.
//
//    // Example sending a request using ListMembersRequest.
//    req := client.ListMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/ListMembers
func (c *Client) ListMembersRequest(input *ListMembersInput) ListMembersRequest {
	op := &aws.Operation{
		Name:       opListMembers,
		HTTPMethod: "GET",
		HTTPPath:   "/members",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListMembersInput{}
	}

	req := c.newRequest(op, input, &ListMembersOutput{})

	return ListMembersRequest{Request: req, Input: input, Copy: c.ListMembersRequest}
}

// ListMembersRequest is the request type for the
// ListMembers API operation.
type ListMembersRequest struct {
	*aws.Request
	Input *ListMembersInput
	Copy  func(*ListMembersInput) ListMembersRequest
}

// Send marshals and sends the ListMembers API request.
func (r ListMembersRequest) Send(ctx context.Context) (*ListMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMembersResponse{
		ListMembersOutput: r.Request.Data.(*ListMembersOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMembersRequestPaginator returns a paginator for ListMembers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMembersRequest(input)
//   p := securityhub.NewListMembersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMembersPaginator(req ListMembersRequest) ListMembersPaginator {
	return ListMembersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMembersInput
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

// ListMembersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMembersPaginator struct {
	aws.Pager
}

func (p *ListMembersPaginator) CurrentPage() *ListMembersOutput {
	return p.Pager.CurrentPage().(*ListMembersOutput)
}

// ListMembersResponse is the response type for the
// ListMembers API operation.
type ListMembersResponse struct {
	*ListMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMembers request.
func (r *ListMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
