// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetConferencePreferenceInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetConferencePreferenceInput) String() string {
	return awsutil.Prettify(s)
}

type GetConferencePreferenceOutput struct {
	_ struct{} `type:"structure"`

	// The conference preference.
	Preference *ConferencePreference `type:"structure"`
}

// String returns the string representation
func (s GetConferencePreferenceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConferencePreference = "GetConferencePreference"

// GetConferencePreferenceRequest returns a request value for making API operation for
// Alexa For Business.
//
// Retrieves the existing conference preferences.
//
//    // Example sending a request using GetConferencePreferenceRequest.
//    req := client.GetConferencePreferenceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetConferencePreference
func (c *Client) GetConferencePreferenceRequest(input *GetConferencePreferenceInput) GetConferencePreferenceRequest {
	op := &aws.Operation{
		Name:       opGetConferencePreference,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConferencePreferenceInput{}
	}

	req := c.newRequest(op, input, &GetConferencePreferenceOutput{})

	return GetConferencePreferenceRequest{Request: req, Input: input, Copy: c.GetConferencePreferenceRequest}
}

// GetConferencePreferenceRequest is the request type for the
// GetConferencePreference API operation.
type GetConferencePreferenceRequest struct {
	*aws.Request
	Input *GetConferencePreferenceInput
	Copy  func(*GetConferencePreferenceInput) GetConferencePreferenceRequest
}

// Send marshals and sends the GetConferencePreference API request.
func (r GetConferencePreferenceRequest) Send(ctx context.Context) (*GetConferencePreferenceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConferencePreferenceResponse{
		GetConferencePreferenceOutput: r.Request.Data.(*GetConferencePreferenceOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetConferencePreferenceResponse is the response type for the
// GetConferencePreference API operation.
type GetConferencePreferenceResponse struct {
	*GetConferencePreferenceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConferencePreference request.
func (r *GetConferencePreferenceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
