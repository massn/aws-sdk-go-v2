// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListS3ResourcesInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 250.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The Amazon Macie Classic member account ID whose associated S3 resources
	// you want to list.
	MemberAccountId *string `locationName:"memberAccountId" type:"string"`

	// Use this parameter when paginating results. Set its value to null on your
	// first call to the ListS3Resources action. Subsequent calls to the action
	// fill nextToken in the request with the value of nextToken from the previous
	// response to continue listing data.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListS3ResourcesInput) String() string {
	return awsutil.Prettify(s)
}

type ListS3ResourcesOutput struct {
	_ struct{} `type:"structure"`

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of the associated S3 resources returned by the action.
	S3Resources []S3ResourceClassification `locationName:"s3Resources" type:"list"`
}

// String returns the string representation
func (s ListS3ResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListS3Resources = "ListS3Resources"

// ListS3ResourcesRequest returns a request value for making API operation for
// Amazon Macie.
//
// Lists all the S3 resources associated with Amazon Macie Classic. If memberAccountId
// isn't specified, the action lists the S3 resources associated with Amazon
// Macie Classic for the current master account. If memberAccountId is specified,
// the action lists the S3 resources associated with Amazon Macie Classic for
// the specified member account.
//
//    // Example sending a request using ListS3ResourcesRequest.
//    req := client.ListS3ResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie-2017-12-19/ListS3Resources
func (c *Client) ListS3ResourcesRequest(input *ListS3ResourcesInput) ListS3ResourcesRequest {
	op := &aws.Operation{
		Name:       opListS3Resources,
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
		input = &ListS3ResourcesInput{}
	}

	req := c.newRequest(op, input, &ListS3ResourcesOutput{})

	return ListS3ResourcesRequest{Request: req, Input: input, Copy: c.ListS3ResourcesRequest}
}

// ListS3ResourcesRequest is the request type for the
// ListS3Resources API operation.
type ListS3ResourcesRequest struct {
	*aws.Request
	Input *ListS3ResourcesInput
	Copy  func(*ListS3ResourcesInput) ListS3ResourcesRequest
}

// Send marshals and sends the ListS3Resources API request.
func (r ListS3ResourcesRequest) Send(ctx context.Context) (*ListS3ResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListS3ResourcesResponse{
		ListS3ResourcesOutput: r.Request.Data.(*ListS3ResourcesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListS3ResourcesRequestPaginator returns a paginator for ListS3Resources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListS3ResourcesRequest(input)
//   p := macie.NewListS3ResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListS3ResourcesPaginator(req ListS3ResourcesRequest) ListS3ResourcesPaginator {
	return ListS3ResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListS3ResourcesInput
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

// ListS3ResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListS3ResourcesPaginator struct {
	aws.Pager
}

func (p *ListS3ResourcesPaginator) CurrentPage() *ListS3ResourcesOutput {
	return p.Pager.CurrentPage().(*ListS3ResourcesOutput)
}

// ListS3ResourcesResponse is the response type for the
// ListS3Resources API operation.
type ListS3ResourcesResponse struct {
	*ListS3ResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListS3Resources request.
func (r *ListS3ResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
