// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return the details of a configuration set. Configuration
// sets enable you to publish email sending events. For information about using
// configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type DescribeConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// A list of configuration set attributes to return.
	ConfigurationSetAttributeNames []ConfigurationSetAttribute `type:"list"`

	// The name of the configuration set to describe.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConfigurationSetInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the details of a configuration set. Configuration sets enable
// you to publish email sending events. For information about using configuration
// sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type DescribeConfigurationSetOutput struct {
	_ struct{} `type:"structure"`

	// The configuration set object associated with the specified configuration
	// set.
	ConfigurationSet *ConfigurationSet `type:"structure"`

	// Specifies whether messages that use the configuration set are required to
	// use Transport Layer Security (TLS).
	DeliveryOptions *DeliveryOptions `type:"structure"`

	// A list of event destinations associated with the configuration set.
	EventDestinations []EventDestination `type:"list"`

	// An object that represents the reputation settings for the configuration set.
	ReputationOptions *ReputationOptions `type:"structure"`

	// The name of the custom open and click tracking domain associated with the
	// configuration set.
	TrackingOptions *TrackingOptions `type:"structure"`
}

// String returns the string representation
func (s DescribeConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConfigurationSet = "DescribeConfigurationSet"

// DescribeConfigurationSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns the details of the specified configuration set. For information about
// using configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DescribeConfigurationSetRequest.
//    req := client.DescribeConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DescribeConfigurationSet
func (c *Client) DescribeConfigurationSetRequest(input *DescribeConfigurationSetInput) DescribeConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opDescribeConfigurationSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &DescribeConfigurationSetOutput{})

	return DescribeConfigurationSetRequest{Request: req, Input: input, Copy: c.DescribeConfigurationSetRequest}
}

// DescribeConfigurationSetRequest is the request type for the
// DescribeConfigurationSet API operation.
type DescribeConfigurationSetRequest struct {
	*aws.Request
	Input *DescribeConfigurationSetInput
	Copy  func(*DescribeConfigurationSetInput) DescribeConfigurationSetRequest
}

// Send marshals and sends the DescribeConfigurationSet API request.
func (r DescribeConfigurationSetRequest) Send(ctx context.Context) (*DescribeConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationSetResponse{
		DescribeConfigurationSetOutput: r.Request.Data.(*DescribeConfigurationSetOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationSetResponse is the response type for the
// DescribeConfigurationSet API operation.
type DescribeConfigurationSetResponse struct {
	*DescribeConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfigurationSet request.
func (r *DescribeConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
