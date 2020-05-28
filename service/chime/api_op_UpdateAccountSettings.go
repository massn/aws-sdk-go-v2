// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateAccountSettingsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The Amazon Chime account settings to update.
	//
	// AccountSettings is a required field
	AccountSettings *AccountSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateAccountSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAccountSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAccountSettingsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.AccountSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountSettings"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountSettings != nil {
		v := s.AccountSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "AccountSettings", v, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateAccountSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAccountSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateAccountSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateAccountSettings = "UpdateAccountSettings"

// UpdateAccountSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates the settings for the specified Amazon Chime account. You can update
// settings for remote control of shared screens, or for the dial-out option.
// For more information about these settings, see Use the Policies Page (https://docs.aws.amazon.com/chime/latest/ag/policies.html)
// in the Amazon Chime Administration Guide.
//
//    // Example sending a request using UpdateAccountSettingsRequest.
//    req := client.UpdateAccountSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateAccountSettings
func (c *Client) UpdateAccountSettingsRequest(input *UpdateAccountSettingsInput) UpdateAccountSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateAccountSettings,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{accountId}/settings",
	}

	if input == nil {
		input = &UpdateAccountSettingsInput{}
	}

	req := c.newRequest(op, input, &UpdateAccountSettingsOutput{})

	return UpdateAccountSettingsRequest{Request: req, Input: input, Copy: c.UpdateAccountSettingsRequest}
}

// UpdateAccountSettingsRequest is the request type for the
// UpdateAccountSettings API operation.
type UpdateAccountSettingsRequest struct {
	*aws.Request
	Input *UpdateAccountSettingsInput
	Copy  func(*UpdateAccountSettingsInput) UpdateAccountSettingsRequest
}

// Send marshals and sends the UpdateAccountSettings API request.
func (r UpdateAccountSettingsRequest) Send(ctx context.Context) (*UpdateAccountSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAccountSettingsResponse{
		UpdateAccountSettingsOutput: r.Request.Data.(*UpdateAccountSettingsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAccountSettingsResponse is the response type for the
// UpdateAccountSettings API operation.
type UpdateAccountSettingsResponse struct {
	*UpdateAccountSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAccountSettings request.
func (r *UpdateAccountSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
