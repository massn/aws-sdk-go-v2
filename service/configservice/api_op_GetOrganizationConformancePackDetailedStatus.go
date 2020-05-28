// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetOrganizationConformancePackDetailedStatusInput struct {
	_ struct{} `type:"structure"`

	// An OrganizationResourceDetailedStatusFilters object.
	Filters *OrganizationResourceDetailedStatusFilters `type:"structure"`

	// The maximum number of OrganizationConformancePackDetailedStatuses returned
	// on each page. If you do not specify a number, AWS Config uses the default.
	// The default is 100.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// The name of organization conformance pack for which you want status details
	// for member accounts.
	//
	// OrganizationConformancePackName is a required field
	OrganizationConformancePackName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetOrganizationConformancePackDetailedStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOrganizationConformancePackDetailedStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOrganizationConformancePackDetailedStatusInput"}

	if s.OrganizationConformancePackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationConformancePackName"))
	}
	if s.OrganizationConformancePackName != nil && len(*s.OrganizationConformancePackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationConformancePackName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetOrganizationConformancePackDetailedStatusOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`

	// A list of OrganizationConformancePackDetailedStatus objects.
	OrganizationConformancePackDetailedStatuses []OrganizationConformancePackDetailedStatus `type:"list"`
}

// String returns the string representation
func (s GetOrganizationConformancePackDetailedStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetOrganizationConformancePackDetailedStatus = "GetOrganizationConformancePackDetailedStatus"

// GetOrganizationConformancePackDetailedStatusRequest returns a request value for making API operation for
// AWS Config.
//
// Returns detailed status for each member account within an organization for
// a given organization conformance pack.
//
// Only a master account can call this API.
//
//    // Example sending a request using GetOrganizationConformancePackDetailedStatusRequest.
//    req := client.GetOrganizationConformancePackDetailedStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetOrganizationConformancePackDetailedStatus
func (c *Client) GetOrganizationConformancePackDetailedStatusRequest(input *GetOrganizationConformancePackDetailedStatusInput) GetOrganizationConformancePackDetailedStatusRequest {
	op := &aws.Operation{
		Name:       opGetOrganizationConformancePackDetailedStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetOrganizationConformancePackDetailedStatusInput{}
	}

	req := c.newRequest(op, input, &GetOrganizationConformancePackDetailedStatusOutput{})

	return GetOrganizationConformancePackDetailedStatusRequest{Request: req, Input: input, Copy: c.GetOrganizationConformancePackDetailedStatusRequest}
}

// GetOrganizationConformancePackDetailedStatusRequest is the request type for the
// GetOrganizationConformancePackDetailedStatus API operation.
type GetOrganizationConformancePackDetailedStatusRequest struct {
	*aws.Request
	Input *GetOrganizationConformancePackDetailedStatusInput
	Copy  func(*GetOrganizationConformancePackDetailedStatusInput) GetOrganizationConformancePackDetailedStatusRequest
}

// Send marshals and sends the GetOrganizationConformancePackDetailedStatus API request.
func (r GetOrganizationConformancePackDetailedStatusRequest) Send(ctx context.Context) (*GetOrganizationConformancePackDetailedStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetOrganizationConformancePackDetailedStatusResponse{
		GetOrganizationConformancePackDetailedStatusOutput: r.Request.Data.(*GetOrganizationConformancePackDetailedStatusOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetOrganizationConformancePackDetailedStatusResponse is the response type for the
// GetOrganizationConformancePackDetailedStatus API operation.
type GetOrganizationConformancePackDetailedStatusResponse struct {
	*GetOrganizationConformancePackDetailedStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetOrganizationConformancePackDetailedStatus request.
func (r *GetOrganizationConformancePackDetailedStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
