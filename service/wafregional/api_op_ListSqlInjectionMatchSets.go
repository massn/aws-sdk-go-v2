// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to list the SqlInjectionMatchSet objects created by the current
// AWS account.
type ListSqlInjectionMatchSetsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of SqlInjectionMatchSet objects that you want AWS WAF
	// to return for this request. If you have more SqlInjectionMatchSet objects
	// than the number you specify for Limit, the response includes a NextMarker
	// value that you can use to get another batch of Rules.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more SqlInjectionMatchSet objects
	// than the value of Limit, AWS WAF returns a NextMarker value in the response
	// that allows you to list another group of SqlInjectionMatchSets. For the second
	// and subsequent ListSqlInjectionMatchSets requests, specify the value of NextMarker
	// from the previous response to get information about another batch of SqlInjectionMatchSets.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListSqlInjectionMatchSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSqlInjectionMatchSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSqlInjectionMatchSetsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response to a ListSqlInjectionMatchSets request.
type ListSqlInjectionMatchSetsOutput struct {
	_ struct{} `type:"structure"`

	// If you have more SqlInjectionMatchSet objects than the number that you specified
	// for Limit in the request, the response includes a NextMarker value. To list
	// more SqlInjectionMatchSet objects, submit another ListSqlInjectionMatchSets
	// request, and specify the NextMarker value from the response in the NextMarker
	// value in the next request.
	NextMarker *string `min:"1" type:"string"`

	// An array of SqlInjectionMatchSetSummary objects.
	SqlInjectionMatchSets []SqlInjectionMatchSetSummary `type:"list"`
}

// String returns the string representation
func (s ListSqlInjectionMatchSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSqlInjectionMatchSets = "ListSqlInjectionMatchSets"

// ListSqlInjectionMatchSetsRequest returns a request value for making API operation for
// AWS WAF Regional.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Returns an array of SqlInjectionMatchSet objects.
//
//    // Example sending a request using ListSqlInjectionMatchSetsRequest.
//    req := client.ListSqlInjectionMatchSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListSqlInjectionMatchSets
func (c *Client) ListSqlInjectionMatchSetsRequest(input *ListSqlInjectionMatchSetsInput) ListSqlInjectionMatchSetsRequest {
	op := &aws.Operation{
		Name:       opListSqlInjectionMatchSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSqlInjectionMatchSetsInput{}
	}

	req := c.newRequest(op, input, &ListSqlInjectionMatchSetsOutput{})

	return ListSqlInjectionMatchSetsRequest{Request: req, Input: input, Copy: c.ListSqlInjectionMatchSetsRequest}
}

// ListSqlInjectionMatchSetsRequest is the request type for the
// ListSqlInjectionMatchSets API operation.
type ListSqlInjectionMatchSetsRequest struct {
	*aws.Request
	Input *ListSqlInjectionMatchSetsInput
	Copy  func(*ListSqlInjectionMatchSetsInput) ListSqlInjectionMatchSetsRequest
}

// Send marshals and sends the ListSqlInjectionMatchSets API request.
func (r ListSqlInjectionMatchSetsRequest) Send(ctx context.Context) (*ListSqlInjectionMatchSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSqlInjectionMatchSetsResponse{
		ListSqlInjectionMatchSetsOutput: r.Request.Data.(*ListSqlInjectionMatchSetsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListSqlInjectionMatchSetsResponse is the response type for the
// ListSqlInjectionMatchSets API operation.
type ListSqlInjectionMatchSetsResponse struct {
	*ListSqlInjectionMatchSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSqlInjectionMatchSets request.
func (r *ListSqlInjectionMatchSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
