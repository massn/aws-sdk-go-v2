// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetUserDefinedFunctionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the functions to be retrieved are located.
	// If none is provided, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database where the functions are located.
	DatabaseName *string `min:"1" type:"string"`

	// The maximum number of functions to return in one response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`

	// An optional function-name pattern string that filters the function definitions
	// returned.
	//
	// Pattern is a required field
	Pattern *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserDefinedFunctionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserDefinedFunctionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserDefinedFunctionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Pattern == nil {
		invalidParams.Add(aws.NewErrParamRequired("Pattern"))
	}
	if s.Pattern != nil && len(*s.Pattern) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Pattern", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetUserDefinedFunctionsOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if the list of functions returned does not include
	// the last requested function.
	NextToken *string `type:"string"`

	// A list of requested function definitions.
	UserDefinedFunctions []UserDefinedFunction `type:"list"`
}

// String returns the string representation
func (s GetUserDefinedFunctionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUserDefinedFunctions = "GetUserDefinedFunctions"

// GetUserDefinedFunctionsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves multiple function definitions from the Data Catalog.
//
//    // Example sending a request using GetUserDefinedFunctionsRequest.
//    req := client.GetUserDefinedFunctionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetUserDefinedFunctions
func (c *Client) GetUserDefinedFunctionsRequest(input *GetUserDefinedFunctionsInput) GetUserDefinedFunctionsRequest {
	op := &aws.Operation{
		Name:       opGetUserDefinedFunctions,
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
		input = &GetUserDefinedFunctionsInput{}
	}

	req := c.newRequest(op, input, &GetUserDefinedFunctionsOutput{})

	return GetUserDefinedFunctionsRequest{Request: req, Input: input, Copy: c.GetUserDefinedFunctionsRequest}
}

// GetUserDefinedFunctionsRequest is the request type for the
// GetUserDefinedFunctions API operation.
type GetUserDefinedFunctionsRequest struct {
	*aws.Request
	Input *GetUserDefinedFunctionsInput
	Copy  func(*GetUserDefinedFunctionsInput) GetUserDefinedFunctionsRequest
}

// Send marshals and sends the GetUserDefinedFunctions API request.
func (r GetUserDefinedFunctionsRequest) Send(ctx context.Context) (*GetUserDefinedFunctionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserDefinedFunctionsResponse{
		GetUserDefinedFunctionsOutput: r.Request.Data.(*GetUserDefinedFunctionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetUserDefinedFunctionsRequestPaginator returns a paginator for GetUserDefinedFunctions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetUserDefinedFunctionsRequest(input)
//   p := glue.NewGetUserDefinedFunctionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetUserDefinedFunctionsPaginator(req GetUserDefinedFunctionsRequest) GetUserDefinedFunctionsPaginator {
	return GetUserDefinedFunctionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetUserDefinedFunctionsInput
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

// GetUserDefinedFunctionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetUserDefinedFunctionsPaginator struct {
	aws.Pager
}

func (p *GetUserDefinedFunctionsPaginator) CurrentPage() *GetUserDefinedFunctionsOutput {
	return p.Pager.CurrentPage().(*GetUserDefinedFunctionsOutput)
}

// GetUserDefinedFunctionsResponse is the response type for the
// GetUserDefinedFunctions API operation.
type GetUserDefinedFunctionsResponse struct {
	*GetUserDefinedFunctionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUserDefinedFunctions request.
func (r *GetUserDefinedFunctionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
