// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteContainerInput struct {
	_ struct{} `type:"structure"`

	// The name of the container to delete.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteContainerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteContainerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteContainerInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteContainerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteContainerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteContainer = "DeleteContainer"

// DeleteContainerRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Deletes the specified container. Before you make a DeleteContainer request,
// delete any objects in the container or in any folders in the container. You
// can delete only empty containers.
//
//    // Example sending a request using DeleteContainerRequest.
//    req := client.DeleteContainerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainer
func (c *Client) DeleteContainerRequest(input *DeleteContainerInput) DeleteContainerRequest {
	op := &aws.Operation{
		Name:       opDeleteContainer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteContainerInput{}
	}

	req := c.newRequest(op, input, &DeleteContainerOutput{})

	return DeleteContainerRequest{Request: req, Input: input, Copy: c.DeleteContainerRequest}
}

// DeleteContainerRequest is the request type for the
// DeleteContainer API operation.
type DeleteContainerRequest struct {
	*aws.Request
	Input *DeleteContainerInput
	Copy  func(*DeleteContainerInput) DeleteContainerRequest
}

// Send marshals and sends the DeleteContainer API request.
func (r DeleteContainerRequest) Send(ctx context.Context) (*DeleteContainerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteContainerResponse{
		DeleteContainerOutput: r.Request.Data.(*DeleteContainerOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteContainerResponse is the response type for the
// DeleteContainer API operation.
type DeleteContainerResponse struct {
	*DeleteContainerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteContainer request.
func (r *DeleteContainerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
