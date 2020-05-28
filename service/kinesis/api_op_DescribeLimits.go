// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLimitsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeLimitsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeLimitsOutput struct {
	_ struct{} `type:"structure"`

	// The number of open shards.
	//
	// OpenShardCount is a required field
	OpenShardCount *int64 `type:"integer" required:"true"`

	// The maximum number of shards.
	//
	// ShardLimit is a required field
	ShardLimit *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s DescribeLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLimits = "DescribeLimits"

// DescribeLimitsRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Describes the shard limits and usage for the account.
//
// If you update your account limits, the old limits might be returned for a
// few minutes.
//
// This operation has a limit of one transaction per second per account.
//
//    // Example sending a request using DescribeLimitsRequest.
//    req := client.DescribeLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/DescribeLimits
func (c *Client) DescribeLimitsRequest(input *DescribeLimitsInput) DescribeLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLimitsInput{}
	}

	req := c.newRequest(op, input, &DescribeLimitsOutput{})

	return DescribeLimitsRequest{Request: req, Input: input, Copy: c.DescribeLimitsRequest}
}

// DescribeLimitsRequest is the request type for the
// DescribeLimits API operation.
type DescribeLimitsRequest struct {
	*aws.Request
	Input *DescribeLimitsInput
	Copy  func(*DescribeLimitsInput) DescribeLimitsRequest
}

// Send marshals and sends the DescribeLimits API request.
func (r DescribeLimitsRequest) Send(ctx context.Context) (*DescribeLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLimitsResponse{
		DescribeLimitsOutput: r.Request.Data.(*DescribeLimitsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLimitsResponse is the response type for the
// DescribeLimits API operation.
type DescribeLimitsResponse struct {
	*DescribeLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLimits request.
func (r *DescribeLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
