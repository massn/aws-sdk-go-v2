// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateResourceInput struct {
	_ struct{} `type:"structure"`

	// The resource ARN.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`

	// The new role to use for the given resource registered in AWS Lake Formation.
	//
	// RoleArn is a required field
	RoleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateResource = "UpdateResource"

// UpdateResourceRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Updates the data access role used for vending access to the given (registered)
// resource in AWS Lake Formation.
//
//    // Example sending a request using UpdateResourceRequest.
//    req := client.UpdateResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/UpdateResource
func (c *Client) UpdateResourceRequest(input *UpdateResourceInput) UpdateResourceRequest {
	op := &aws.Operation{
		Name:       opUpdateResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateResourceInput{}
	}

	req := c.newRequest(op, input, &UpdateResourceOutput{})

	return UpdateResourceRequest{Request: req, Input: input, Copy: c.UpdateResourceRequest}
}

// UpdateResourceRequest is the request type for the
// UpdateResource API operation.
type UpdateResourceRequest struct {
	*aws.Request
	Input *UpdateResourceInput
	Copy  func(*UpdateResourceInput) UpdateResourceRequest
}

// Send marshals and sends the UpdateResource API request.
func (r UpdateResourceRequest) Send(ctx context.Context) (*UpdateResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateResourceResponse{
		UpdateResourceOutput: r.Request.Data.(*UpdateResourceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateResourceResponse is the response type for the
// UpdateResource API operation.
type UpdateResourceResponse struct {
	*UpdateResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateResource request.
func (r *UpdateResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
