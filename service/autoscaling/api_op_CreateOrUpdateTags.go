// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type CreateOrUpdateTagsInput struct {
	_ struct{} `type:"structure"`

	// One or more tags.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s CreateOrUpdateTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOrUpdateTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOrUpdateTagsInput"}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
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

type CreateOrUpdateTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateOrUpdateTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateOrUpdateTags = "CreateOrUpdateTags"

// CreateOrUpdateTagsRequest returns a request value for making API operation for
// Auto Scaling.
//
// Creates or updates tags for the specified Auto Scaling group.
//
// When you specify a tag with a key that already exists, the operation overwrites
// the previous tag definition, and you do not get an error message.
//
// For more information, see Tagging Auto Scaling Groups and Instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using CreateOrUpdateTagsRequest.
//    req := client.CreateOrUpdateTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/CreateOrUpdateTags
func (c *Client) CreateOrUpdateTagsRequest(input *CreateOrUpdateTagsInput) CreateOrUpdateTagsRequest {
	op := &aws.Operation{
		Name:       opCreateOrUpdateTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateOrUpdateTagsInput{}
	}

	req := c.newRequest(op, input, &CreateOrUpdateTagsOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return CreateOrUpdateTagsRequest{Request: req, Input: input, Copy: c.CreateOrUpdateTagsRequest}
}

// CreateOrUpdateTagsRequest is the request type for the
// CreateOrUpdateTags API operation.
type CreateOrUpdateTagsRequest struct {
	*aws.Request
	Input *CreateOrUpdateTagsInput
	Copy  func(*CreateOrUpdateTagsInput) CreateOrUpdateTagsRequest
}

// Send marshals and sends the CreateOrUpdateTags API request.
func (r CreateOrUpdateTagsRequest) Send(ctx context.Context) (*CreateOrUpdateTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOrUpdateTagsResponse{
		CreateOrUpdateTagsOutput: r.Request.Data.(*CreateOrUpdateTagsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOrUpdateTagsResponse is the response type for the
// CreateOrUpdateTags API operation.
type CreateOrUpdateTagsResponse struct {
	*CreateOrUpdateTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOrUpdateTags request.
func (r *CreateOrUpdateTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
