// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteNetworkProfileInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the network profile to delete.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNetworkProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNetworkProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNetworkProfileInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNetworkProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNetworkProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteNetworkProfile = "DeleteNetworkProfile"

// DeleteNetworkProfileRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Deletes a network profile.
//
//    // Example sending a request using DeleteNetworkProfileRequest.
//    req := client.DeleteNetworkProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/DeleteNetworkProfile
func (c *Client) DeleteNetworkProfileRequest(input *DeleteNetworkProfileInput) DeleteNetworkProfileRequest {
	op := &aws.Operation{
		Name:       opDeleteNetworkProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNetworkProfileInput{}
	}

	req := c.newRequest(op, input, &DeleteNetworkProfileOutput{})

	return DeleteNetworkProfileRequest{Request: req, Input: input, Copy: c.DeleteNetworkProfileRequest}
}

// DeleteNetworkProfileRequest is the request type for the
// DeleteNetworkProfile API operation.
type DeleteNetworkProfileRequest struct {
	*aws.Request
	Input *DeleteNetworkProfileInput
	Copy  func(*DeleteNetworkProfileInput) DeleteNetworkProfileRequest
}

// Send marshals and sends the DeleteNetworkProfile API request.
func (r DeleteNetworkProfileRequest) Send(ctx context.Context) (*DeleteNetworkProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNetworkProfileResponse{
		DeleteNetworkProfileOutput: r.Request.Data.(*DeleteNetworkProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNetworkProfileResponse is the response type for the
// DeleteNetworkProfile API operation.
type DeleteNetworkProfileResponse struct {
	*DeleteNetworkProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNetworkProfile request.
func (r *DeleteNetworkProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
