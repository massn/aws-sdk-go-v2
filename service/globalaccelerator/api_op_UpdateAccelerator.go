// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateAcceleratorInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the accelerator to update.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`

	// Indicates whether an accelerator is enabled. The value is true or false.
	// The default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to
	// false, the accelerator can be deleted.
	Enabled *bool `type:"boolean"`

	// The value for the address type must be IPv4.
	IpAddressType IpAddressType `type:"string" enum:"true"`

	// The name of the accelerator. The name can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens (-), and must not begin
	// or end with a hyphen.
	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateAcceleratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAcceleratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAcceleratorInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAcceleratorOutput struct {
	_ struct{} `type:"structure"`

	// Information about the updated accelerator.
	Accelerator *Accelerator `type:"structure"`
}

// String returns the string representation
func (s UpdateAcceleratorOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAccelerator = "UpdateAccelerator"

// UpdateAcceleratorRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Update an accelerator. To see an AWS CLI example of updating an accelerator,
// scroll down to Example.
//
// You must specify the US West (Oregon) Region to create or update accelerators.
//
//    // Example sending a request using UpdateAcceleratorRequest.
//    req := client.UpdateAcceleratorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/UpdateAccelerator
func (c *Client) UpdateAcceleratorRequest(input *UpdateAcceleratorInput) UpdateAcceleratorRequest {
	op := &aws.Operation{
		Name:       opUpdateAccelerator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAcceleratorInput{}
	}

	req := c.newRequest(op, input, &UpdateAcceleratorOutput{})

	return UpdateAcceleratorRequest{Request: req, Input: input, Copy: c.UpdateAcceleratorRequest}
}

// UpdateAcceleratorRequest is the request type for the
// UpdateAccelerator API operation.
type UpdateAcceleratorRequest struct {
	*aws.Request
	Input *UpdateAcceleratorInput
	Copy  func(*UpdateAcceleratorInput) UpdateAcceleratorRequest
}

// Send marshals and sends the UpdateAccelerator API request.
func (r UpdateAcceleratorRequest) Send(ctx context.Context) (*UpdateAcceleratorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAcceleratorResponse{
		UpdateAcceleratorOutput: r.Request.Data.(*UpdateAcceleratorOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAcceleratorResponse is the response type for the
// UpdateAccelerator API operation.
type UpdateAcceleratorResponse struct {
	*UpdateAcceleratorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccelerator request.
func (r *UpdateAcceleratorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
