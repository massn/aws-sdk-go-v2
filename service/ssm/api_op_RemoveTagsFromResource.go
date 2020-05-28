// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveTagsFromResourceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the resource from which you want to remove tags. For example:
	//
	// ManagedInstance: mi-012345abcde
	//
	// MaintenanceWindow: mw-012345abcde
	//
	// PatchBaseline: pb-012345abcde
	//
	// For the Document and Parameter values, use the name of the resource.
	//
	// The ManagedInstance type for this API action is only for on-premises managed
	// instances. Specify the name of the managed instance in the following format:
	// mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// The type of resource from which you want to remove a tag.
	//
	// The ManagedInstance type for this API action is only for on-premises managed
	// instances. Specify the name of the managed instance in the following format:
	// mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// ResourceType is a required field
	ResourceType ResourceTypeForTagging `type:"string" required:"true" enum:"true"`

	// Tag keys that you want to remove from the specified resource.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsFromResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsFromResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsFromResourceInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveTagsFromResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsFromResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveTagsFromResource = "RemoveTagsFromResource"

// RemoveTagsFromResourceRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Removes tag keys from the specified resource.
//
//    // Example sending a request using RemoveTagsFromResourceRequest.
//    req := client.RemoveTagsFromResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/RemoveTagsFromResource
func (c *Client) RemoveTagsFromResourceRequest(input *RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest {
	op := &aws.Operation{
		Name:       opRemoveTagsFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveTagsFromResourceInput{}
	}

	req := c.newRequest(op, input, &RemoveTagsFromResourceOutput{})

	return RemoveTagsFromResourceRequest{Request: req, Input: input, Copy: c.RemoveTagsFromResourceRequest}
}

// RemoveTagsFromResourceRequest is the request type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceRequest struct {
	*aws.Request
	Input *RemoveTagsFromResourceInput
	Copy  func(*RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest
}

// Send marshals and sends the RemoveTagsFromResource API request.
func (r RemoveTagsFromResourceRequest) Send(ctx context.Context) (*RemoveTagsFromResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsFromResourceResponse{
		RemoveTagsFromResourceOutput: r.Request.Data.(*RemoveTagsFromResourceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsFromResourceResponse is the response type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceResponse struct {
	*RemoveTagsFromResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTagsFromResource request.
func (r *RemoveTagsFromResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
