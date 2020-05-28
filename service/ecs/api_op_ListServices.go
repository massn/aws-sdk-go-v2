// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListServicesInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the services to list. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The launch type for the services to list.
	LaunchType LaunchType `locationName:"launchType" type:"string" enum:"true"`

	// The maximum number of service results returned by ListServices in paginated
	// output. When this parameter is used, ListServices only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListServices
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListServices returns up to 10 results
	// and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a ListServices request indicating that
	// more results are available to fulfill the request and further calls will
	// be needed. If maxResults was provided, it is possible the number of results
	// to be fewer than maxResults.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The scheduling strategy for services to list.
	SchedulingStrategy SchedulingStrategy `locationName:"schedulingStrategy" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListServicesInput) String() string {
	return awsutil.Prettify(s)
}

type ListServicesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListServices request. When the
	// results of a ListServices request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of full ARN entries for each service associated with the specified
	// cluster.
	ServiceArns []string `locationName:"serviceArns" type:"list"`
}

// String returns the string representation
func (s ListServicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListServices = "ListServices"

// ListServicesRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Lists the services that are running in a specified cluster.
//
//    // Example sending a request using ListServicesRequest.
//    req := client.ListServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListServices
func (c *Client) ListServicesRequest(input *ListServicesInput) ListServicesRequest {
	op := &aws.Operation{
		Name:       opListServices,
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
		input = &ListServicesInput{}
	}

	req := c.newRequest(op, input, &ListServicesOutput{})

	return ListServicesRequest{Request: req, Input: input, Copy: c.ListServicesRequest}
}

// ListServicesRequest is the request type for the
// ListServices API operation.
type ListServicesRequest struct {
	*aws.Request
	Input *ListServicesInput
	Copy  func(*ListServicesInput) ListServicesRequest
}

// Send marshals and sends the ListServices API request.
func (r ListServicesRequest) Send(ctx context.Context) (*ListServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServicesResponse{
		ListServicesOutput: r.Request.Data.(*ListServicesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServicesRequestPaginator returns a paginator for ListServices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServicesRequest(input)
//   p := ecs.NewListServicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServicesPaginator(req ListServicesRequest) ListServicesPaginator {
	return ListServicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListServicesInput
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

// ListServicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServicesPaginator struct {
	aws.Pager
}

func (p *ListServicesPaginator) CurrentPage() *ListServicesOutput {
	return p.Pager.CurrentPage().(*ListServicesOutput)
}

// ListServicesResponse is the response type for the
// ListServices API operation.
type ListServicesResponse struct {
	*ListServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServices request.
func (r *ListServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
