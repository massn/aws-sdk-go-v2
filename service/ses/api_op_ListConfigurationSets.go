// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to list the configuration sets associated with your
// AWS account. Configuration sets enable you to publish email sending events.
// For information about using configuration sets, see the Amazon SES Developer
// Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsInput struct {
	_ struct{} `type:"structure"`

	// The number of configuration sets to return.
	MaxItems *int64 `type:"integer"`

	// A token returned from a previous call to ListConfigurationSets to indicate
	// the position of the configuration set in the configuration set list.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListConfigurationSetsInput) String() string {
	return awsutil.Prettify(s)
}

// A list of configuration sets associated with your AWS account. Configuration
// sets enable you to publish email sending events. For information about using
// configuration sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsOutput struct {
	_ struct{} `type:"structure"`

	// A list of configuration sets.
	ConfigurationSets []ConfigurationSet `type:"list"`

	// A token indicating that there are additional configuration sets available
	// to be listed. Pass this token to successive calls of ListConfigurationSets.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListConfigurationSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListConfigurationSets = "ListConfigurationSets"

// ListConfigurationSetsRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Provides a list of the configuration sets associated with your Amazon SES
// account in the current AWS Region. For information about using configuration
// sets, see Monitoring Your Amazon SES Sending Activity (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second. This operation
// will return up to 1,000 configuration sets each time it is run. If your Amazon
// SES account has more than 1,000 configuration sets, this operation will also
// return a NextToken element. You can then execute the ListConfigurationSets
// operation again, passing the NextToken parameter and the value of the NextToken
// element to retrieve additional results.
//
//    // Example sending a request using ListConfigurationSetsRequest.
//    req := client.ListConfigurationSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/ListConfigurationSets
func (c *Client) ListConfigurationSetsRequest(input *ListConfigurationSetsInput) ListConfigurationSetsRequest {
	op := &aws.Operation{
		Name:       opListConfigurationSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListConfigurationSetsInput{}
	}

	req := c.newRequest(op, input, &ListConfigurationSetsOutput{})

	return ListConfigurationSetsRequest{Request: req, Input: input, Copy: c.ListConfigurationSetsRequest}
}

// ListConfigurationSetsRequest is the request type for the
// ListConfigurationSets API operation.
type ListConfigurationSetsRequest struct {
	*aws.Request
	Input *ListConfigurationSetsInput
	Copy  func(*ListConfigurationSetsInput) ListConfigurationSetsRequest
}

// Send marshals and sends the ListConfigurationSets API request.
func (r ListConfigurationSetsRequest) Send(ctx context.Context) (*ListConfigurationSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConfigurationSetsResponse{
		ListConfigurationSetsOutput: r.Request.Data.(*ListConfigurationSetsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListConfigurationSetsResponse is the response type for the
// ListConfigurationSets API operation.
type ListConfigurationSetsResponse struct {
	*ListConfigurationSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConfigurationSets request.
func (r *ListConfigurationSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
