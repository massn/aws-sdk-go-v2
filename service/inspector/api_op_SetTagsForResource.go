// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type SetTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the assessment template that you want to set tags to.
	//
	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" min:"1" type:"string" required:"true"`

	// A collection of key and value pairs that you want to set to the assessment
	// template.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s SetTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetTagsForResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetTagsForResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetTagsForResource = "SetTagsForResource"

// SetTagsForResourceRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Sets tags (key and value pairs) to the assessment template that is specified
// by the ARN of the assessment template.
//
//    // Example sending a request using SetTagsForResourceRequest.
//    req := client.SetTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/SetTagsForResource
func (c *Client) SetTagsForResourceRequest(input *SetTagsForResourceInput) SetTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opSetTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &SetTagsForResourceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return SetTagsForResourceRequest{Request: req, Input: input, Copy: c.SetTagsForResourceRequest}
}

// SetTagsForResourceRequest is the request type for the
// SetTagsForResource API operation.
type SetTagsForResourceRequest struct {
	*aws.Request
	Input *SetTagsForResourceInput
	Copy  func(*SetTagsForResourceInput) SetTagsForResourceRequest
}

// Send marshals and sends the SetTagsForResource API request.
func (r SetTagsForResourceRequest) Send(ctx context.Context) (*SetTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetTagsForResourceResponse{
		SetTagsForResourceOutput: r.Request.Data.(*SetTagsForResourceOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetTagsForResourceResponse is the response type for the
// SetTagsForResource API operation.
type SetTagsForResourceResponse struct {
	*SetTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetTagsForResource request.
func (r *SetTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
