// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteUserInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the user's enrollment in the organization. Required.
	//
	// EnrollmentId is a required field
	EnrollmentId *string `type:"string" required:"true"`

	// The ARN of the user to delete in the organization. Required.
	UserArn *string `type:"string"`
}

// String returns the string representation
func (s DeleteUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUserInput"}

	if s.EnrollmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnrollmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteUser = "DeleteUser"

// DeleteUserRequest returns a request value for making API operation for
// Alexa For Business.
//
// Deletes a specified user by user ARN and enrollment ARN.
//
//    // Example sending a request using DeleteUserRequest.
//    req := client.DeleteUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/DeleteUser
func (c *Client) DeleteUserRequest(input *DeleteUserInput) DeleteUserRequest {
	op := &aws.Operation{
		Name:       opDeleteUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserInput{}
	}

	req := c.newRequest(op, input, &DeleteUserOutput{})

	return DeleteUserRequest{Request: req, Input: input, Copy: c.DeleteUserRequest}
}

// DeleteUserRequest is the request type for the
// DeleteUser API operation.
type DeleteUserRequest struct {
	*aws.Request
	Input *DeleteUserInput
	Copy  func(*DeleteUserInput) DeleteUserRequest
}

// Send marshals and sends the DeleteUser API request.
func (r DeleteUserRequest) Send(ctx context.Context) (*DeleteUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUserResponse{
		DeleteUserOutput: r.Request.Data.(*DeleteUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUserResponse is the response type for the
// DeleteUser API operation.
type DeleteUserResponse struct {
	*DeleteUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUser request.
func (r *DeleteUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
