// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRegexPatternSetsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of objects that you want AWS WAF to return for this request.
	// If more objects are available, in the response, AWS WAF provides a NextMarker
	// value that you can use in a subsequent call to get the next batch of objects.
	Limit *int64 `min:"1" type:"integer"`

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, AWS WAF
	// returns a NextMarker value in the response. To retrieve the next batch of
	// objects, provide the marker from the prior call in your next request.
	NextMarker *string `min:"1" type:"string"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ListRegexPatternSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRegexPatternSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRegexPatternSetsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRegexPatternSetsOutput struct {
	_ struct{} `type:"structure"`

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, AWS WAF
	// returns a NextMarker value in the response. To retrieve the next batch of
	// objects, provide the marker from the prior call in your next request.
	NextMarker *string `min:"1" type:"string"`

	RegexPatternSets []RegexPatternSetSummary `type:"list"`
}

// String returns the string representation
func (s ListRegexPatternSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRegexPatternSets = "ListRegexPatternSets"

// ListRegexPatternSetsRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Retrieves an array of RegexPatternSetSummary objects for the regex pattern
// sets that you manage.
//
//    // Example sending a request using ListRegexPatternSetsRequest.
//    req := client.ListRegexPatternSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/ListRegexPatternSets
func (c *Client) ListRegexPatternSetsRequest(input *ListRegexPatternSetsInput) ListRegexPatternSetsRequest {
	op := &aws.Operation{
		Name:       opListRegexPatternSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRegexPatternSetsInput{}
	}

	req := c.newRequest(op, input, &ListRegexPatternSetsOutput{})

	return ListRegexPatternSetsRequest{Request: req, Input: input, Copy: c.ListRegexPatternSetsRequest}
}

// ListRegexPatternSetsRequest is the request type for the
// ListRegexPatternSets API operation.
type ListRegexPatternSetsRequest struct {
	*aws.Request
	Input *ListRegexPatternSetsInput
	Copy  func(*ListRegexPatternSetsInput) ListRegexPatternSetsRequest
}

// Send marshals and sends the ListRegexPatternSets API request.
func (r ListRegexPatternSetsRequest) Send(ctx context.Context) (*ListRegexPatternSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRegexPatternSetsResponse{
		ListRegexPatternSetsOutput: r.Request.Data.(*ListRegexPatternSetsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRegexPatternSetsResponse is the response type for the
// ListRegexPatternSets API operation.
type ListRegexPatternSetsResponse struct {
	*ListRegexPatternSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRegexPatternSets request.
func (r *ListRegexPatternSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
