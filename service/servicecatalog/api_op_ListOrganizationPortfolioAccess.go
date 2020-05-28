// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListOrganizationPortfolioAccessInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The organization node type that will be returned in the output.
	//
	//    * ORGANIZATION - Organization that has access to the portfolio.
	//
	//    * ORGANIZATIONAL_UNIT - Organizational unit that has access to the portfolio
	//    within your organization.
	//
	//    * ACCOUNT - Account that has access to the portfolio within your organization.
	//
	// OrganizationNodeType is a required field
	OrganizationNodeType OrganizationNodeType `type:"string" required:"true" enum:"true"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The portfolio identifier. For example, port-2abcdext3y5fk.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListOrganizationPortfolioAccessInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOrganizationPortfolioAccessInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOrganizationPortfolioAccessInput"}
	if len(s.OrganizationNodeType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationNodeType"))
	}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListOrganizationPortfolioAccessOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// Displays information about the organization nodes.
	OrganizationNodes []OrganizationNode `type:"list"`
}

// String returns the string representation
func (s ListOrganizationPortfolioAccessOutput) String() string {
	return awsutil.Prettify(s)
}

const opListOrganizationPortfolioAccess = "ListOrganizationPortfolioAccess"

// ListOrganizationPortfolioAccessRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists the organization nodes that have access to the specified portfolio.
// This API can only be called by the master account in the organization.
//
//    // Example sending a request using ListOrganizationPortfolioAccessRequest.
//    req := client.ListOrganizationPortfolioAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListOrganizationPortfolioAccess
func (c *Client) ListOrganizationPortfolioAccessRequest(input *ListOrganizationPortfolioAccessInput) ListOrganizationPortfolioAccessRequest {
	op := &aws.Operation{
		Name:       opListOrganizationPortfolioAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListOrganizationPortfolioAccessInput{}
	}

	req := c.newRequest(op, input, &ListOrganizationPortfolioAccessOutput{})

	return ListOrganizationPortfolioAccessRequest{Request: req, Input: input, Copy: c.ListOrganizationPortfolioAccessRequest}
}

// ListOrganizationPortfolioAccessRequest is the request type for the
// ListOrganizationPortfolioAccess API operation.
type ListOrganizationPortfolioAccessRequest struct {
	*aws.Request
	Input *ListOrganizationPortfolioAccessInput
	Copy  func(*ListOrganizationPortfolioAccessInput) ListOrganizationPortfolioAccessRequest
}

// Send marshals and sends the ListOrganizationPortfolioAccess API request.
func (r ListOrganizationPortfolioAccessRequest) Send(ctx context.Context) (*ListOrganizationPortfolioAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOrganizationPortfolioAccessResponse{
		ListOrganizationPortfolioAccessOutput: r.Request.Data.(*ListOrganizationPortfolioAccessOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListOrganizationPortfolioAccessRequestPaginator returns a paginator for ListOrganizationPortfolioAccess.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListOrganizationPortfolioAccessRequest(input)
//   p := servicecatalog.NewListOrganizationPortfolioAccessRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListOrganizationPortfolioAccessPaginator(req ListOrganizationPortfolioAccessRequest) ListOrganizationPortfolioAccessPaginator {
	return ListOrganizationPortfolioAccessPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListOrganizationPortfolioAccessInput
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

// ListOrganizationPortfolioAccessPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListOrganizationPortfolioAccessPaginator struct {
	aws.Pager
}

func (p *ListOrganizationPortfolioAccessPaginator) CurrentPage() *ListOrganizationPortfolioAccessOutput {
	return p.Pager.CurrentPage().(*ListOrganizationPortfolioAccessOutput)
}

// ListOrganizationPortfolioAccessResponse is the response type for the
// ListOrganizationPortfolioAccess API operation.
type ListOrganizationPortfolioAccessResponse struct {
	*ListOrganizationPortfolioAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOrganizationPortfolioAccess request.
func (r *ListOrganizationPortfolioAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
