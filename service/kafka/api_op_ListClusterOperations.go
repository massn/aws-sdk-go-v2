// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListClusterOperationsInput struct {
	_ struct{} `type:"structure"`

	// ClusterArn is a required field
	ClusterArn *string `location:"uri" locationName:"clusterArn" type:"string" required:"true"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClusterOperationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListClusterOperationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListClusterOperationsInput"}

	if s.ClusterArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClusterOperationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterArn != nil {
		v := *s.ClusterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "clusterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	return nil
}

// The response contains an array containing cluster operation information and
// a next token if the response is truncated.
type ListClusterOperationsOutput struct {
	_ struct{} `type:"structure"`

	// An array of cluster operation information objects.
	ClusterOperationInfoList []ClusterOperationInfo `locationName:"clusterOperationInfoList" type:"list"`

	// If the response of ListClusterOperations is truncated, it returns a NextToken
	// in the response. This Nexttoken should be sent in the subsequent request
	// to ListClusterOperations.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClusterOperationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListClusterOperationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClusterOperationInfoList != nil {
		v := s.ClusterOperationInfoList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "clusterOperationInfoList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListClusterOperations = "ListClusterOperations"

// ListClusterOperationsRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Returns a list of all the operations that have been performed on the specified
// MSK cluster.
//
//    // Example sending a request using ListClusterOperationsRequest.
//    req := client.ListClusterOperationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/ListClusterOperations
func (c *Client) ListClusterOperationsRequest(input *ListClusterOperationsInput) ListClusterOperationsRequest {
	op := &aws.Operation{
		Name:       opListClusterOperations,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{clusterArn}/operations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClusterOperationsInput{}
	}

	req := c.newRequest(op, input, &ListClusterOperationsOutput{})

	return ListClusterOperationsRequest{Request: req, Input: input, Copy: c.ListClusterOperationsRequest}
}

// ListClusterOperationsRequest is the request type for the
// ListClusterOperations API operation.
type ListClusterOperationsRequest struct {
	*aws.Request
	Input *ListClusterOperationsInput
	Copy  func(*ListClusterOperationsInput) ListClusterOperationsRequest
}

// Send marshals and sends the ListClusterOperations API request.
func (r ListClusterOperationsRequest) Send(ctx context.Context) (*ListClusterOperationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListClusterOperationsResponse{
		ListClusterOperationsOutput: r.Request.Data.(*ListClusterOperationsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListClusterOperationsRequestPaginator returns a paginator for ListClusterOperations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListClusterOperationsRequest(input)
//   p := kafka.NewListClusterOperationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListClusterOperationsPaginator(req ListClusterOperationsRequest) ListClusterOperationsPaginator {
	return ListClusterOperationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListClusterOperationsInput
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

// ListClusterOperationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListClusterOperationsPaginator struct {
	aws.Pager
}

func (p *ListClusterOperationsPaginator) CurrentPage() *ListClusterOperationsOutput {
	return p.Pager.CurrentPage().(*ListClusterOperationsOutput)
}

// ListClusterOperationsResponse is the response type for the
// ListClusterOperations API operation.
type ListClusterOperationsResponse struct {
	*ListClusterOperationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListClusterOperations request.
func (r *ListClusterOperationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
