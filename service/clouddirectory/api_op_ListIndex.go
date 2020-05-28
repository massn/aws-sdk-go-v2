// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListIndexInput struct {
	_ struct{} `type:"structure"`

	// The consistency level to execute the request at.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The ARN of the directory that the index exists in.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The reference to the index to list.
	//
	// IndexReference is a required field
	IndexReference *ObjectReference `type:"structure" required:"true"`

	// The maximum number of objects in a single page to retrieve from the index
	// during a request. For more information, see Amazon Cloud Directory Limits
	// (http://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html).
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// Specifies the ranges of indexed values that you want to query.
	RangesOnIndexedValues []ObjectAttributeRange `type:"list"`
}

// String returns the string representation
func (s ListIndexInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIndexInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIndexInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.IndexReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexReference"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.RangesOnIndexedValues != nil {
		for i, v := range s.RangesOnIndexedValues {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RangesOnIndexedValues", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListIndexInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IndexReference != nil {
		v := s.IndexReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "IndexReference", v, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RangesOnIndexedValues != nil {
		v := s.RangesOnIndexedValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "RangesOnIndexedValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListIndexOutput struct {
	_ struct{} `type:"structure"`

	// The objects and indexed values attached to the index.
	IndexAttachments []IndexAttachment `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListIndexOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListIndexOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IndexAttachments != nil {
		v := s.IndexAttachments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "IndexAttachments", metadata)
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

const opListIndex = "ListIndex"

// ListIndexRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists objects attached to the specified index.
//
//    // Example sending a request using ListIndexRequest.
//    req := client.ListIndexRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListIndex
func (c *Client) ListIndexRequest(input *ListIndexInput) ListIndexRequest {
	op := &aws.Operation{
		Name:       opListIndex,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/index/targets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListIndexInput{}
	}

	req := c.newRequest(op, input, &ListIndexOutput{})

	return ListIndexRequest{Request: req, Input: input, Copy: c.ListIndexRequest}
}

// ListIndexRequest is the request type for the
// ListIndex API operation.
type ListIndexRequest struct {
	*aws.Request
	Input *ListIndexInput
	Copy  func(*ListIndexInput) ListIndexRequest
}

// Send marshals and sends the ListIndex API request.
func (r ListIndexRequest) Send(ctx context.Context) (*ListIndexResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIndexResponse{
		ListIndexOutput: r.Request.Data.(*ListIndexOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListIndexRequestPaginator returns a paginator for ListIndex.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListIndexRequest(input)
//   p := clouddirectory.NewListIndexRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListIndexPaginator(req ListIndexRequest) ListIndexPaginator {
	return ListIndexPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListIndexInput
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

// ListIndexPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListIndexPaginator struct {
	aws.Pager
}

func (p *ListIndexPaginator) CurrentPage() *ListIndexOutput {
	return p.Pager.CurrentPage().(*ListIndexOutput)
}

// ListIndexResponse is the response type for the
// ListIndex API operation.
type ListIndexResponse struct {
	*ListIndexOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIndex request.
func (r *ListIndexResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
