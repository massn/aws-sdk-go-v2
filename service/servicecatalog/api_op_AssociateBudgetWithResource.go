// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateBudgetWithResourceInput struct {
	_ struct{} `type:"structure"`

	// The name of the budget you want to associate.
	//
	// BudgetName is a required field
	BudgetName *string `min:"1" type:"string" required:"true"`

	// The resource identifier. Either a portfolio-id or a product-id.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateBudgetWithResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateBudgetWithResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateBudgetWithResourceInput"}

	if s.BudgetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BudgetName"))
	}
	if s.BudgetName != nil && len(*s.BudgetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BudgetName", 1))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateBudgetWithResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateBudgetWithResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateBudgetWithResource = "AssociateBudgetWithResource"

// AssociateBudgetWithResourceRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Associates the specified budget with the specified resource.
//
//    // Example sending a request using AssociateBudgetWithResourceRequest.
//    req := client.AssociateBudgetWithResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/AssociateBudgetWithResource
func (c *Client) AssociateBudgetWithResourceRequest(input *AssociateBudgetWithResourceInput) AssociateBudgetWithResourceRequest {
	op := &aws.Operation{
		Name:       opAssociateBudgetWithResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateBudgetWithResourceInput{}
	}

	req := c.newRequest(op, input, &AssociateBudgetWithResourceOutput{})

	return AssociateBudgetWithResourceRequest{Request: req, Input: input, Copy: c.AssociateBudgetWithResourceRequest}
}

// AssociateBudgetWithResourceRequest is the request type for the
// AssociateBudgetWithResource API operation.
type AssociateBudgetWithResourceRequest struct {
	*aws.Request
	Input *AssociateBudgetWithResourceInput
	Copy  func(*AssociateBudgetWithResourceInput) AssociateBudgetWithResourceRequest
}

// Send marshals and sends the AssociateBudgetWithResource API request.
func (r AssociateBudgetWithResourceRequest) Send(ctx context.Context) (*AssociateBudgetWithResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateBudgetWithResourceResponse{
		AssociateBudgetWithResourceOutput: r.Request.Data.(*AssociateBudgetWithResourceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateBudgetWithResourceResponse is the response type for the
// AssociateBudgetWithResource API operation.
type AssociateBudgetWithResourceResponse struct {
	*AssociateBudgetWithResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateBudgetWithResource request.
func (r *AssociateBudgetWithResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
