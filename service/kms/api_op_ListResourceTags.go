// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListResourceTagsInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the customer master key (CMK).
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// Use this parameter to specify the maximum number of items to return. When
	// this value is present, AWS KMS does not return more than the specified number
	// of items, but it might return fewer.
	//
	// This value is optional. If you include a value, it must be between 1 and
	// 50, inclusive. If you do not include a value, it defaults to 50.
	Limit *int64 `min:"1" type:"integer"`

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	//
	// Do not attempt to construct this value. Use only the value of NextMarker
	// from the truncated response you just received.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListResourceTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourceTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourceTagsInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResourceTagsOutput struct {
	_ struct{} `type:"structure"`

	// When Truncated is true, this element is present and contains the value to
	// use for the Marker parameter in a subsequent request.
	//
	// Do not assume or infer any information from this value.
	NextMarker *string `min:"1" type:"string"`

	// A list of tags. Each tag consists of a tag key and a tag value.
	Tags []Tag `type:"list"`

	// A flag that indicates whether there are more items in the list. When this
	// value is true, the list in this response is truncated. To get more items,
	// pass the value of the NextMarker element in thisresponse to the Marker parameter
	// in a subsequent request.
	Truncated *bool `type:"boolean"`
}

// String returns the string representation
func (s ListResourceTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListResourceTags = "ListResourceTags"

// ListResourceTagsRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Returns a list of all tags for the specified customer master key (CMK).
//
// You cannot perform this operation on a CMK in a different AWS account.
//
//    // Example sending a request using ListResourceTagsRequest.
//    req := client.ListResourceTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ListResourceTags
func (c *Client) ListResourceTagsRequest(input *ListResourceTagsInput) ListResourceTagsRequest {
	op := &aws.Operation{
		Name:       opListResourceTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListResourceTagsInput{}
	}

	req := c.newRequest(op, input, &ListResourceTagsOutput{})

	return ListResourceTagsRequest{Request: req, Input: input, Copy: c.ListResourceTagsRequest}
}

// ListResourceTagsRequest is the request type for the
// ListResourceTags API operation.
type ListResourceTagsRequest struct {
	*aws.Request
	Input *ListResourceTagsInput
	Copy  func(*ListResourceTagsInput) ListResourceTagsRequest
}

// Send marshals and sends the ListResourceTags API request.
func (r ListResourceTagsRequest) Send(ctx context.Context) (*ListResourceTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourceTagsResponse{
		ListResourceTagsOutput: r.Request.Data.(*ListResourceTagsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListResourceTagsResponse is the response type for the
// ListResourceTags API operation.
type ListResourceTagsResponse struct {
	*ListResourceTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResourceTags request.
func (r *ListResourceTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
