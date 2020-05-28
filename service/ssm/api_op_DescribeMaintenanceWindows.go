// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeMaintenanceWindowsInput struct {
	_ struct{} `type:"structure"`

	// Optional filters used to narrow down the scope of the returned maintenance
	// windows. Supported filter keys are Name and Enabled.
	Filters []MaintenanceWindowFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"10" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMaintenanceWindowsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMaintenanceWindowsInput"}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeMaintenanceWindowsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`

	// Information about the maintenance windows.
	WindowIdentities []MaintenanceWindowIdentity `type:"list"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMaintenanceWindows = "DescribeMaintenanceWindows"

// DescribeMaintenanceWindowsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the maintenance windows in an AWS account.
//
//    // Example sending a request using DescribeMaintenanceWindowsRequest.
//    req := client.DescribeMaintenanceWindowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeMaintenanceWindows
func (c *Client) DescribeMaintenanceWindowsRequest(input *DescribeMaintenanceWindowsInput) DescribeMaintenanceWindowsRequest {
	op := &aws.Operation{
		Name:       opDescribeMaintenanceWindows,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMaintenanceWindowsInput{}
	}

	req := c.newRequest(op, input, &DescribeMaintenanceWindowsOutput{})

	return DescribeMaintenanceWindowsRequest{Request: req, Input: input, Copy: c.DescribeMaintenanceWindowsRequest}
}

// DescribeMaintenanceWindowsRequest is the request type for the
// DescribeMaintenanceWindows API operation.
type DescribeMaintenanceWindowsRequest struct {
	*aws.Request
	Input *DescribeMaintenanceWindowsInput
	Copy  func(*DescribeMaintenanceWindowsInput) DescribeMaintenanceWindowsRequest
}

// Send marshals and sends the DescribeMaintenanceWindows API request.
func (r DescribeMaintenanceWindowsRequest) Send(ctx context.Context) (*DescribeMaintenanceWindowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMaintenanceWindowsResponse{
		DescribeMaintenanceWindowsOutput: r.Request.Data.(*DescribeMaintenanceWindowsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMaintenanceWindowsResponse is the response type for the
// DescribeMaintenanceWindows API operation.
type DescribeMaintenanceWindowsResponse struct {
	*DescribeMaintenanceWindowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMaintenanceWindows request.
func (r *DescribeMaintenanceWindowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
