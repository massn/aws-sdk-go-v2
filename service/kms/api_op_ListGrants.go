// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListGrantsInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the customer master key (CMK).
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify
	// a CMK in a different AWS account, you must use the key ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// Use this parameter to specify the maximum number of items to return. When
	// this value is present, AWS KMS does not return more than the specified number
	// of items, but it might return fewer.
	//
	// This value is optional. If you include a value, it must be between 1 and
	// 100, inclusive. If you do not include a value, it defaults to 50.
	Limit *int64 `min:"1" type:"integer"`

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListGrantsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGrantsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGrantsInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListGrantsOutput struct {
	_ struct{} `type:"structure"`

	// A list of grants.
	Grants []GrantListEntry `type:"list"`

	// When Truncated is true, this element is present and contains the value to
	// use for the Marker parameter in a subsequent request.
	NextMarker *string `min:"1" type:"string"`

	// A flag that indicates whether there are more items in the list. When this
	// value is true, the list in this response is truncated. To get more items,
	// pass the value of the NextMarker element in thisresponse to the Marker parameter
	// in a subsequent request.
	Truncated *bool `type:"boolean"`
}

// String returns the string representation
func (s ListGrantsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListGrants = "ListGrants"

// ListGrantsRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Gets a list of all grants for the specified customer master key (CMK).
//
// To perform this operation on a CMK in a different AWS account, specify the
// key ARN in the value of the KeyId parameter.
//
//    // Example sending a request using ListGrantsRequest.
//    req := client.ListGrantsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ListGrants
func (c *Client) ListGrantsRequest(input *ListGrantsInput) ListGrantsRequest {
	op := &aws.Operation{
		Name:       opListGrants,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "Limit",
			TruncationToken: "Truncated",
		},
	}

	if input == nil {
		input = &ListGrantsInput{}
	}

	req := c.newRequest(op, input, &ListGrantsOutput{})

	return ListGrantsRequest{Request: req, Input: input, Copy: c.ListGrantsRequest}
}

// ListGrantsRequest is the request type for the
// ListGrants API operation.
type ListGrantsRequest struct {
	*aws.Request
	Input *ListGrantsInput
	Copy  func(*ListGrantsInput) ListGrantsRequest
}

// Send marshals and sends the ListGrants API request.
func (r ListGrantsRequest) Send(ctx context.Context) (*ListGrantsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGrantsResponse{
		ListGrantsOutput: r.Request.Data.(*ListGrantsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGrantsRequestPaginator returns a paginator for ListGrants.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGrantsRequest(input)
//   p := kms.NewListGrantsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGrantsPaginator(req ListGrantsRequest) ListGrantsPaginator {
	return ListGrantsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListGrantsInput
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

// ListGrantsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGrantsPaginator struct {
	aws.Pager
}

func (p *ListGrantsPaginator) CurrentPage() *ListGrantsOutput {
	return p.Pager.CurrentPage().(*ListGrantsOutput)
}

// ListGrantsResponse is the response type for the
// ListGrants API operation.
type ListGrantsResponse struct {
	*ListGrantsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGrants request.
func (r *ListGrantsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
