// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to retrieve information about an email address that's on the suppression
// list for your account.
type GetSuppressedDestinationInput struct {
	_ struct{} `type:"structure"`

	// The email address that's on the account suppression list.
	//
	// EmailAddress is a required field
	EmailAddress *string `location:"uri" locationName:"EmailAddress" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSuppressedDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSuppressedDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSuppressedDestinationInput"}

	if s.EmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSuppressedDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EmailAddress != nil {
		v := *s.EmailAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EmailAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about the suppressed email address.
type GetSuppressedDestinationOutput struct {
	_ struct{} `type:"structure"`

	// An object containing information about the suppressed email address.
	//
	// SuppressedDestination is a required field
	SuppressedDestination *SuppressedDestination `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSuppressedDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSuppressedDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SuppressedDestination != nil {
		v := s.SuppressedDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SuppressedDestination", v, metadata)
	}
	return nil
}

const opGetSuppressedDestination = "GetSuppressedDestination"

// GetSuppressedDestinationRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Retrieves information about a specific email address that's on the suppression
// list for your account.
//
//    // Example sending a request using GetSuppressedDestinationRequest.
//    req := client.GetSuppressedDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/GetSuppressedDestination
func (c *Client) GetSuppressedDestinationRequest(input *GetSuppressedDestinationInput) GetSuppressedDestinationRequest {
	op := &aws.Operation{
		Name:       opGetSuppressedDestination,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/email/suppression/addresses/{EmailAddress}",
	}

	if input == nil {
		input = &GetSuppressedDestinationInput{}
	}

	req := c.newRequest(op, input, &GetSuppressedDestinationOutput{})

	return GetSuppressedDestinationRequest{Request: req, Input: input, Copy: c.GetSuppressedDestinationRequest}
}

// GetSuppressedDestinationRequest is the request type for the
// GetSuppressedDestination API operation.
type GetSuppressedDestinationRequest struct {
	*aws.Request
	Input *GetSuppressedDestinationInput
	Copy  func(*GetSuppressedDestinationInput) GetSuppressedDestinationRequest
}

// Send marshals and sends the GetSuppressedDestination API request.
func (r GetSuppressedDestinationRequest) Send(ctx context.Context) (*GetSuppressedDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSuppressedDestinationResponse{
		GetSuppressedDestinationOutput: r.Request.Data.(*GetSuppressedDestinationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSuppressedDestinationResponse is the response type for the
// GetSuppressedDestination API operation.
type GetSuppressedDestinationResponse struct {
	*GetSuppressedDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSuppressedDestination request.
func (r *GetSuppressedDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
