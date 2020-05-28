// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListUpdatesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of update results returned by ListUpdates in paginated
	// output. When you use this parameter, ListUpdates returns only maxResults
	// results in a single page along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another ListUpdates
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If you don't use this parameter, ListUpdates returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The name of the Amazon EKS cluster to list updates for.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// The nextToken value returned from a previous paginated ListUpdates request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The name of the Amazon EKS managed node group to list updates for.
	NodegroupName *string `location:"querystring" locationName:"nodegroupName" type:"string"`
}

// String returns the string representation
func (s ListUpdatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUpdatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUpdatesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUpdatesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NodegroupName != nil {
		v := *s.NodegroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nodegroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListUpdatesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListUpdates request. When the
	// results of a ListUpdates request exceed maxResults, you can use this value
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of all the updates for the specified cluster and Region.
	UpdateIds []string `locationName:"updateIds" type:"list"`
}

// String returns the string representation
func (s ListUpdatesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUpdatesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UpdateIds != nil {
		v := s.UpdateIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "updateIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListUpdates = "ListUpdates"

// ListUpdatesRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Lists the updates associated with an Amazon EKS cluster or managed node group
// in your AWS account, in the specified Region.
//
//    // Example sending a request using ListUpdatesRequest.
//    req := client.ListUpdatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/ListUpdates
func (c *Client) ListUpdatesRequest(input *ListUpdatesInput) ListUpdatesRequest {
	op := &aws.Operation{
		Name:       opListUpdates,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}/updates",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListUpdatesInput{}
	}

	req := c.newRequest(op, input, &ListUpdatesOutput{})

	return ListUpdatesRequest{Request: req, Input: input, Copy: c.ListUpdatesRequest}
}

// ListUpdatesRequest is the request type for the
// ListUpdates API operation.
type ListUpdatesRequest struct {
	*aws.Request
	Input *ListUpdatesInput
	Copy  func(*ListUpdatesInput) ListUpdatesRequest
}

// Send marshals and sends the ListUpdates API request.
func (r ListUpdatesRequest) Send(ctx context.Context) (*ListUpdatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUpdatesResponse{
		ListUpdatesOutput: r.Request.Data.(*ListUpdatesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListUpdatesRequestPaginator returns a paginator for ListUpdates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListUpdatesRequest(input)
//   p := eks.NewListUpdatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListUpdatesPaginator(req ListUpdatesRequest) ListUpdatesPaginator {
	return ListUpdatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListUpdatesInput
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

// ListUpdatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListUpdatesPaginator struct {
	aws.Pager
}

func (p *ListUpdatesPaginator) CurrentPage() *ListUpdatesOutput {
	return p.Pager.CurrentPage().(*ListUpdatesOutput)
}

// ListUpdatesResponse is the response type for the
// ListUpdates API operation.
type ListUpdatesResponse struct {
	*ListUpdatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUpdates request.
func (r *ListUpdatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
