// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TestIdentityProviderInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned identifier for a specific file transfer protocol-enabled
	// server. That server's user authentication method is tested with a user name
	// and password.
	//
	// ServerId is a required field
	ServerId *string `min:"19" type:"string" required:"true"`

	// The type of file transfer protocol to be tested.
	//
	// The available protocols are:
	//
	//    * Secure Shell (SSH) File Transfer Protocol (SFTP)
	//
	//    * File Transfer Protocol Secure (FTPS)
	//
	//    * File Transfer Protocol (FTP)
	ServerProtocol Protocol `type:"string" enum:"true"`

	// The name of the user account to be tested.
	//
	// UserName is a required field
	UserName *string `min:"3" type:"string" required:"true"`

	// The password of the user account to be tested.
	UserPassword *string `type:"string" sensitive:"true"`
}

// String returns the string representation
func (s TestIdentityProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestIdentityProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestIdentityProviderInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}
	if s.ServerId != nil && len(*s.ServerId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerId", 19))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TestIdentityProviderOutput struct {
	_ struct{} `type:"structure"`

	// A message that indicates whether the test was successful or not.
	Message *string `type:"string"`

	// The response that is returned from your API Gateway.
	Response *string `type:"string"`

	// The HTTP status code that is the response from your API Gateway.
	//
	// StatusCode is a required field
	StatusCode *int64 `type:"integer" required:"true"`

	// The endpoint of the service used to authenticate a user.
	//
	// Url is a required field
	Url *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TestIdentityProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opTestIdentityProvider = "TestIdentityProvider"

// TestIdentityProviderRequest returns a request value for making API operation for
// AWS Transfer Family.
//
// If the IdentityProviderType of a file transfer protocol-enabled server is
// API_Gateway, tests whether your API Gateway is set up successfully. We highly
// recommend that you call this operation to test your authentication method
// as soon as you create your server. By doing so, you can troubleshoot issues
// with the API Gateway integration to ensure that your users can successfully
// use the service.
//
//    // Example sending a request using TestIdentityProviderRequest.
//    req := client.TestIdentityProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/TestIdentityProvider
func (c *Client) TestIdentityProviderRequest(input *TestIdentityProviderInput) TestIdentityProviderRequest {
	op := &aws.Operation{
		Name:       opTestIdentityProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestIdentityProviderInput{}
	}

	req := c.newRequest(op, input, &TestIdentityProviderOutput{})

	return TestIdentityProviderRequest{Request: req, Input: input, Copy: c.TestIdentityProviderRequest}
}

// TestIdentityProviderRequest is the request type for the
// TestIdentityProvider API operation.
type TestIdentityProviderRequest struct {
	*aws.Request
	Input *TestIdentityProviderInput
	Copy  func(*TestIdentityProviderInput) TestIdentityProviderRequest
}

// Send marshals and sends the TestIdentityProvider API request.
func (r TestIdentityProviderRequest) Send(ctx context.Context) (*TestIdentityProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestIdentityProviderResponse{
		TestIdentityProviderOutput: r.Request.Data.(*TestIdentityProviderOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestIdentityProviderResponse is the response type for the
// TestIdentityProvider API operation.
type TestIdentityProviderResponse struct {
	*TestIdentityProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestIdentityProvider request.
func (r *TestIdentityProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
