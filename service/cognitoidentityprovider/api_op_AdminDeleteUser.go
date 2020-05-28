// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the request to delete a user as an administrator.
type AdminDeleteUserInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID for the user pool where you want to delete the user.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user name of the user you wish to delete.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminDeleteUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminDeleteUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminDeleteUserInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AdminDeleteUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AdminDeleteUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminDeleteUser = "AdminDeleteUser"

// AdminDeleteUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Deletes a user as an administrator. Works on any user.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminDeleteUserRequest.
//    req := client.AdminDeleteUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminDeleteUser
func (c *Client) AdminDeleteUserRequest(input *AdminDeleteUserInput) AdminDeleteUserRequest {
	op := &aws.Operation{
		Name:       opAdminDeleteUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminDeleteUserInput{}
	}

	req := c.newRequest(op, input, &AdminDeleteUserOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AdminDeleteUserRequest{Request: req, Input: input, Copy: c.AdminDeleteUserRequest}
}

// AdminDeleteUserRequest is the request type for the
// AdminDeleteUser API operation.
type AdminDeleteUserRequest struct {
	*aws.Request
	Input *AdminDeleteUserInput
	Copy  func(*AdminDeleteUserInput) AdminDeleteUserRequest
}

// Send marshals and sends the AdminDeleteUser API request.
func (r AdminDeleteUserRequest) Send(ctx context.Context) (*AdminDeleteUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminDeleteUserResponse{
		AdminDeleteUserOutput: r.Request.Data.(*AdminDeleteUserOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminDeleteUserResponse is the response type for the
// AdminDeleteUser API operation.
type AdminDeleteUserResponse struct {
	*AdminDeleteUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminDeleteUser request.
func (r *AdminDeleteUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
