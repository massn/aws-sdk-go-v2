// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListAssetModelsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned per paginated request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to be used for the next set of paginated results.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListAssetModelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAssetModelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAssetModelsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAssetModelsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

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

type ListAssetModelsOutput struct {
	_ struct{} `type:"structure"`

	// A list that summarizes each asset model.
	//
	// AssetModelSummaries is a required field
	AssetModelSummaries []AssetModelSummary `locationName:"assetModelSummaries" type:"list" required:"true"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListAssetModelsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAssetModelsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssetModelSummaries != nil {
		v := s.AssetModelSummaries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "assetModelSummaries", metadata)
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

const opListAssetModels = "ListAssetModels"

// ListAssetModelsRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Retrieves a paginated list of summaries of all asset models.
//
//    // Example sending a request using ListAssetModelsRequest.
//    req := client.ListAssetModelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/ListAssetModels
func (c *Client) ListAssetModelsRequest(input *ListAssetModelsInput) ListAssetModelsRequest {
	op := &aws.Operation{
		Name:       opListAssetModels,
		HTTPMethod: "GET",
		HTTPPath:   "/asset-models",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAssetModelsInput{}
	}

	req := c.newRequest(op, input, &ListAssetModelsOutput{})
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("model.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return ListAssetModelsRequest{Request: req, Input: input, Copy: c.ListAssetModelsRequest}
}

// ListAssetModelsRequest is the request type for the
// ListAssetModels API operation.
type ListAssetModelsRequest struct {
	*aws.Request
	Input *ListAssetModelsInput
	Copy  func(*ListAssetModelsInput) ListAssetModelsRequest
}

// Send marshals and sends the ListAssetModels API request.
func (r ListAssetModelsRequest) Send(ctx context.Context) (*ListAssetModelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssetModelsResponse{
		ListAssetModelsOutput: r.Request.Data.(*ListAssetModelsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAssetModelsRequestPaginator returns a paginator for ListAssetModels.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAssetModelsRequest(input)
//   p := iotsitewise.NewListAssetModelsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAssetModelsPaginator(req ListAssetModelsRequest) ListAssetModelsPaginator {
	return ListAssetModelsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAssetModelsInput
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

// ListAssetModelsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAssetModelsPaginator struct {
	aws.Pager
}

func (p *ListAssetModelsPaginator) CurrentPage() *ListAssetModelsOutput {
	return p.Pager.CurrentPage().(*ListAssetModelsOutput)
}

// ListAssetModelsResponse is the response type for the
// ListAssetModels API operation.
type ListAssetModelsResponse struct {
	*ListAssetModelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssetModels request.
func (r *ListAssetModelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
