// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResumeSessionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the disconnected session to resume.
	//
	// SessionId is a required field
	SessionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResumeSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResumeSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResumeSessionInput"}

	if s.SessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SessionId"))
	}
	if s.SessionId != nil && len(*s.SessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SessionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResumeSessionOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the session.
	SessionId *string `min:"1" type:"string"`

	// A URL back to SSM Agent on the instance that the Session Manager client uses
	// to send commands and receive output from the instance. Format: wss://ssmmessages.region.amazonaws.com/v1/data-channel/session-id?stream=(input|output).
	//
	// region represents the Region identifier for an AWS Region supported by AWS
	// Systems Manager, such as us-east-2 for the US East (Ohio) Region. For a list
	// of supported region values, see the Region column in Systems Manager service
	// endpoints (http://docs.aws.amazon.com/general/latest/gr/ssm.html#ssm_region)
	// in the AWS General Reference.
	//
	// session-id represents the ID of a Session Manager session, such as 1a2b3c4dEXAMPLE.
	StreamUrl *string `type:"string"`

	// An encrypted token value containing session and caller information. Used
	// to authenticate the connection to the instance.
	TokenValue *string `type:"string"`
}

// String returns the string representation
func (s ResumeSessionOutput) String() string {
	return awsutil.Prettify(s)
}

const opResumeSession = "ResumeSession"

// ResumeSessionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Reconnects a session to an instance after it has been disconnected. Connections
// can be resumed for disconnected sessions, but not terminated sessions.
//
// This command is primarily for use by client machines to automatically reconnect
// during intermittent network issues. It is not intended for any other use.
//
//    // Example sending a request using ResumeSessionRequest.
//    req := client.ResumeSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ResumeSession
func (c *Client) ResumeSessionRequest(input *ResumeSessionInput) ResumeSessionRequest {
	op := &aws.Operation{
		Name:       opResumeSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResumeSessionInput{}
	}

	req := c.newRequest(op, input, &ResumeSessionOutput{})

	return ResumeSessionRequest{Request: req, Input: input, Copy: c.ResumeSessionRequest}
}

// ResumeSessionRequest is the request type for the
// ResumeSession API operation.
type ResumeSessionRequest struct {
	*aws.Request
	Input *ResumeSessionInput
	Copy  func(*ResumeSessionInput) ResumeSessionRequest
}

// Send marshals and sends the ResumeSession API request.
func (r ResumeSessionRequest) Send(ctx context.Context) (*ResumeSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResumeSessionResponse{
		ResumeSessionOutput: r.Request.Data.(*ResumeSessionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResumeSessionResponse is the response type for the
// ResumeSession API operation.
type ResumeSessionResponse struct {
	*ResumeSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResumeSession request.
func (r *ResumeSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
