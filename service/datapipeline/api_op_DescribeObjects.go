// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeObjects.
type DescribeObjectsInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether any expressions in the object should be evaluated when
	// the object descriptions are returned.
	EvaluateExpressions *bool `locationName:"evaluateExpressions" type:"boolean"`

	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// DescribeObjects with the marker value from the previous call to retrieve
	// the next set of results.
	Marker *string `locationName:"marker" type:"string"`

	// The IDs of the pipeline objects that contain the definitions to be described.
	// You can pass as many as 25 identifiers in a single call to DescribeObjects.
	//
	// ObjectIds is a required field
	ObjectIds []string `locationName:"objectIds" type:"list" required:"true"`

	// The ID of the pipeline that contains the object definitions.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeObjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeObjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeObjectsInput"}

	if s.ObjectIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectIds"))
	}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DescribeObjects.
type DescribeObjectsOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether there are more results to return.
	HasMoreResults *bool `locationName:"hasMoreResults" type:"boolean"`

	// The starting point for the next page of results. To view the next page of
	// results, call DescribeObjects again with this marker value. If the value
	// is null, there are no more results.
	Marker *string `locationName:"marker" type:"string"`

	// An array of object definitions.
	//
	// PipelineObjects is a required field
	PipelineObjects []PipelineObject `locationName:"pipelineObjects" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeObjectsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeObjects = "DescribeObjects"

// DescribeObjectsRequest returns a request value for making API operation for
// AWS Data Pipeline.
//
// Gets the object definitions for a set of objects associated with the pipeline.
// Object definitions are composed of a set of fields that define the properties
// of the object.
//
//    // Example sending a request using DescribeObjectsRequest.
//    req := client.DescribeObjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/DescribeObjects
func (c *Client) DescribeObjectsRequest(input *DescribeObjectsInput) DescribeObjectsRequest {
	op := &aws.Operation{
		Name:       opDescribeObjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"marker"},
			OutputTokens:    []string{"marker"},
			LimitToken:      "",
			TruncationToken: "hasMoreResults",
		},
	}

	if input == nil {
		input = &DescribeObjectsInput{}
	}

	req := c.newRequest(op, input, &DescribeObjectsOutput{})

	return DescribeObjectsRequest{Request: req, Input: input, Copy: c.DescribeObjectsRequest}
}

// DescribeObjectsRequest is the request type for the
// DescribeObjects API operation.
type DescribeObjectsRequest struct {
	*aws.Request
	Input *DescribeObjectsInput
	Copy  func(*DescribeObjectsInput) DescribeObjectsRequest
}

// Send marshals and sends the DescribeObjects API request.
func (r DescribeObjectsRequest) Send(ctx context.Context) (*DescribeObjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeObjectsResponse{
		DescribeObjectsOutput: r.Request.Data.(*DescribeObjectsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeObjectsRequestPaginator returns a paginator for DescribeObjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeObjectsRequest(input)
//   p := datapipeline.NewDescribeObjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeObjectsPaginator(req DescribeObjectsRequest) DescribeObjectsPaginator {
	return DescribeObjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeObjectsInput
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

// DescribeObjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeObjectsPaginator struct {
	aws.Pager
}

func (p *DescribeObjectsPaginator) CurrentPage() *DescribeObjectsOutput {
	return p.Pager.CurrentPage().(*DescribeObjectsOutput)
}

// DescribeObjectsResponse is the response type for the
// DescribeObjects API operation.
type DescribeObjectsResponse struct {
	*DescribeObjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeObjects request.
func (r *DescribeObjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
