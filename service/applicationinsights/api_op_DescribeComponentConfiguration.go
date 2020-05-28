// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeComponentConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the component.
	//
	// ComponentName is a required field
	ComponentName *string `type:"string" required:"true"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeComponentConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeComponentConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeComponentConfigurationInput"}

	if s.ComponentName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComponentName"))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}
	if s.ResourceGroupName != nil && len(*s.ResourceGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeComponentConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The configuration settings of the component. The value is the escaped JSON
	// of the configuration.
	ComponentConfiguration *string `min:"1" type:"string"`

	// Indicates whether the application component is monitored.
	Monitor *bool `type:"boolean"`

	// The tier of the application component. Supported tiers include DOT_NET_CORE,
	// DOT_NET_WORKER, DOT_NET_WEB, SQL_SERVER, and DEFAULT
	Tier Tier `min:"1" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeComponentConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeComponentConfiguration = "DescribeComponentConfiguration"

// DescribeComponentConfigurationRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Describes the monitoring configuration of the component.
//
//    // Example sending a request using DescribeComponentConfigurationRequest.
//    req := client.DescribeComponentConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/DescribeComponentConfiguration
func (c *Client) DescribeComponentConfigurationRequest(input *DescribeComponentConfigurationInput) DescribeComponentConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeComponentConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeComponentConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeComponentConfigurationOutput{})

	return DescribeComponentConfigurationRequest{Request: req, Input: input, Copy: c.DescribeComponentConfigurationRequest}
}

// DescribeComponentConfigurationRequest is the request type for the
// DescribeComponentConfiguration API operation.
type DescribeComponentConfigurationRequest struct {
	*aws.Request
	Input *DescribeComponentConfigurationInput
	Copy  func(*DescribeComponentConfigurationInput) DescribeComponentConfigurationRequest
}

// Send marshals and sends the DescribeComponentConfiguration API request.
func (r DescribeComponentConfigurationRequest) Send(ctx context.Context) (*DescribeComponentConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeComponentConfigurationResponse{
		DescribeComponentConfigurationOutput: r.Request.Data.(*DescribeComponentConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeComponentConfigurationResponse is the response type for the
// DescribeComponentConfiguration API operation.
type DescribeComponentConfigurationResponse struct {
	*DescribeComponentConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeComponentConfiguration request.
func (r *DescribeComponentConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
