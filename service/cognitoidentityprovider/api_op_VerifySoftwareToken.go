// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type VerifySoftwareTokenInput struct {
	_ struct{} `type:"structure"`

	// The access token.
	AccessToken *string `type:"string" sensitive:"true"`

	// The friendly device name.
	FriendlyDeviceName *string `type:"string"`

	// The session which should be passed both ways in challenge-response calls
	// to the service.
	Session *string `min:"20" type:"string"`

	// The one time password computed using the secret code returned by
	//
	// UserCode is a required field
	UserCode *string `min:"6" type:"string" required:"true"`
}

// String returns the string representation
func (s VerifySoftwareTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifySoftwareTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "VerifySoftwareTokenInput"}
	if s.Session != nil && len(*s.Session) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("Session", 20))
	}

	if s.UserCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserCode"))
	}
	if s.UserCode != nil && len(*s.UserCode) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("UserCode", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type VerifySoftwareTokenOutput struct {
	_ struct{} `type:"structure"`

	// The session which should be passed both ways in challenge-response calls
	// to the service.
	Session *string `min:"20" type:"string"`

	// The status of the verify software token.
	Status VerifySoftwareTokenResponseType `type:"string" enum:"true"`
}

// String returns the string representation
func (s VerifySoftwareTokenOutput) String() string {
	return awsutil.Prettify(s)
}

const opVerifySoftwareToken = "VerifySoftwareToken"

// VerifySoftwareTokenRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Use this API to register a user's entered TOTP code and mark the user's software
// token MFA status as "verified" if successful. The request takes an access
// token or a session string, but not both.
//
//    // Example sending a request using VerifySoftwareTokenRequest.
//    req := client.VerifySoftwareTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/VerifySoftwareToken
func (c *Client) VerifySoftwareTokenRequest(input *VerifySoftwareTokenInput) VerifySoftwareTokenRequest {
	op := &aws.Operation{
		Name:       opVerifySoftwareToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &VerifySoftwareTokenInput{}
	}

	req := c.newRequest(op, input, &VerifySoftwareTokenOutput{})

	return VerifySoftwareTokenRequest{Request: req, Input: input, Copy: c.VerifySoftwareTokenRequest}
}

// VerifySoftwareTokenRequest is the request type for the
// VerifySoftwareToken API operation.
type VerifySoftwareTokenRequest struct {
	*aws.Request
	Input *VerifySoftwareTokenInput
	Copy  func(*VerifySoftwareTokenInput) VerifySoftwareTokenRequest
}

// Send marshals and sends the VerifySoftwareToken API request.
func (r VerifySoftwareTokenRequest) Send(ctx context.Context) (*VerifySoftwareTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &VerifySoftwareTokenResponse{
		VerifySoftwareTokenOutput: r.Request.Data.(*VerifySoftwareTokenOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// VerifySoftwareTokenResponse is the response type for the
// VerifySoftwareToken API operation.
type VerifySoftwareTokenResponse struct {
	*VerifySoftwareTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// VerifySoftwareToken request.
func (r *VerifySoftwareTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
