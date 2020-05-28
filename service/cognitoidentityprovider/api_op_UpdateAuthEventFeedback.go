// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateAuthEventFeedbackInput struct {
	_ struct{} `type:"structure"`

	// The event ID.
	//
	// EventId is a required field
	EventId *string `min:"1" type:"string" required:"true"`

	// The feedback token.
	//
	// FeedbackToken is a required field
	FeedbackToken *string `type:"string" required:"true" sensitive:"true"`

	// The authentication event feedback value.
	//
	// FeedbackValue is a required field
	FeedbackValue FeedbackValueType `type:"string" required:"true" enum:"true"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user pool username.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s UpdateAuthEventFeedbackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAuthEventFeedbackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAuthEventFeedbackInput"}

	if s.EventId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventId"))
	}
	if s.EventId != nil && len(*s.EventId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventId", 1))
	}

	if s.FeedbackToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("FeedbackToken"))
	}
	if len(s.FeedbackValue) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("FeedbackValue"))
	}

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

type UpdateAuthEventFeedbackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAuthEventFeedbackOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAuthEventFeedback = "UpdateAuthEventFeedback"

// UpdateAuthEventFeedbackRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Provides the feedback for an authentication event whether it was from a valid
// user or not. This feedback is used for improving the risk evaluation decision
// for the user pool as part of Amazon Cognito advanced security.
//
//    // Example sending a request using UpdateAuthEventFeedbackRequest.
//    req := client.UpdateAuthEventFeedbackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateAuthEventFeedback
func (c *Client) UpdateAuthEventFeedbackRequest(input *UpdateAuthEventFeedbackInput) UpdateAuthEventFeedbackRequest {
	op := &aws.Operation{
		Name:       opUpdateAuthEventFeedback,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAuthEventFeedbackInput{}
	}

	req := c.newRequest(op, input, &UpdateAuthEventFeedbackOutput{})

	return UpdateAuthEventFeedbackRequest{Request: req, Input: input, Copy: c.UpdateAuthEventFeedbackRequest}
}

// UpdateAuthEventFeedbackRequest is the request type for the
// UpdateAuthEventFeedback API operation.
type UpdateAuthEventFeedbackRequest struct {
	*aws.Request
	Input *UpdateAuthEventFeedbackInput
	Copy  func(*UpdateAuthEventFeedbackInput) UpdateAuthEventFeedbackRequest
}

// Send marshals and sends the UpdateAuthEventFeedback API request.
func (r UpdateAuthEventFeedbackRequest) Send(ctx context.Context) (*UpdateAuthEventFeedbackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAuthEventFeedbackResponse{
		UpdateAuthEventFeedbackOutput: r.Request.Data.(*UpdateAuthEventFeedbackOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAuthEventFeedbackResponse is the response type for the
// UpdateAuthEventFeedback API operation.
type UpdateAuthEventFeedbackResponse struct {
	*UpdateAuthEventFeedbackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAuthEventFeedback request.
func (r *UpdateAuthEventFeedbackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
