// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLaunchConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// The launch configuration names. If you omit this parameter, all launch configurations
	// are described.
	LaunchConfigurationNames []string `type:"list"`

	// The maximum number of items to return with this call. The default value is
	// 50 and the maximum value is 100.
	MaxRecords *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeLaunchConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeLaunchConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// The launch configurations.
	//
	// LaunchConfigurations is a required field
	LaunchConfigurations []LaunchConfiguration `type:"list" required:"true"`

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this
	// string for the NextToken value when requesting the next set of items. This
	// value is null when there are no more items to return.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeLaunchConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLaunchConfigurations = "DescribeLaunchConfigurations"

// DescribeLaunchConfigurationsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes one or more launch configurations.
//
//    // Example sending a request using DescribeLaunchConfigurationsRequest.
//    req := client.DescribeLaunchConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeLaunchConfigurations
func (c *Client) DescribeLaunchConfigurationsRequest(input *DescribeLaunchConfigurationsInput) DescribeLaunchConfigurationsRequest {
	op := &aws.Operation{
		Name:       opDescribeLaunchConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeLaunchConfigurationsInput{}
	}

	req := c.newRequest(op, input, &DescribeLaunchConfigurationsOutput{})

	return DescribeLaunchConfigurationsRequest{Request: req, Input: input, Copy: c.DescribeLaunchConfigurationsRequest}
}

// DescribeLaunchConfigurationsRequest is the request type for the
// DescribeLaunchConfigurations API operation.
type DescribeLaunchConfigurationsRequest struct {
	*aws.Request
	Input *DescribeLaunchConfigurationsInput
	Copy  func(*DescribeLaunchConfigurationsInput) DescribeLaunchConfigurationsRequest
}

// Send marshals and sends the DescribeLaunchConfigurations API request.
func (r DescribeLaunchConfigurationsRequest) Send(ctx context.Context) (*DescribeLaunchConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLaunchConfigurationsResponse{
		DescribeLaunchConfigurationsOutput: r.Request.Data.(*DescribeLaunchConfigurationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeLaunchConfigurationsRequestPaginator returns a paginator for DescribeLaunchConfigurations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeLaunchConfigurationsRequest(input)
//   p := autoscaling.NewDescribeLaunchConfigurationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeLaunchConfigurationsPaginator(req DescribeLaunchConfigurationsRequest) DescribeLaunchConfigurationsPaginator {
	return DescribeLaunchConfigurationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeLaunchConfigurationsInput
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

// DescribeLaunchConfigurationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeLaunchConfigurationsPaginator struct {
	aws.Pager
}

func (p *DescribeLaunchConfigurationsPaginator) CurrentPage() *DescribeLaunchConfigurationsOutput {
	return p.Pager.CurrentPage().(*DescribeLaunchConfigurationsOutput)
}

// DescribeLaunchConfigurationsResponse is the response type for the
// DescribeLaunchConfigurations API operation.
type DescribeLaunchConfigurationsResponse struct {
	*DescribeLaunchConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLaunchConfigurations request.
func (r *DescribeLaunchConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
