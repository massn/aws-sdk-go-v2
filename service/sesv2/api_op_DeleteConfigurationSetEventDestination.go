// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to delete an event destination from a configuration set.
type DeleteConfigurationSetEventDestinationInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that contains the event destination that
	// you want to delete.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`

	// The name of the event destination that you want to delete.
	//
	// EventDestinationName is a required field
	EventDestinationName *string `location:"uri" locationName:"EventDestinationName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetEventDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetEventDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetEventDestinationInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.EventDestinationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventDestinationName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationSetEventDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EventDestinationName != nil {
		v := *s.EventDestinationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EventDestinationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type DeleteConfigurationSetEventDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetEventDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationSetEventDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteConfigurationSetEventDestination = "DeleteConfigurationSetEventDestination"

// DeleteConfigurationSetEventDestinationRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Delete an event destination.
//
// Events include message sends, deliveries, opens, clicks, bounces, and complaints.
// Event destinations are places that you can send information about these events
// to. For example, you can send event data to Amazon SNS to receive notifications
// when you receive bounces or complaints, or you can use Amazon Kinesis Data
// Firehose to stream data to Amazon S3 for long-term storage.
//
//    // Example sending a request using DeleteConfigurationSetEventDestinationRequest.
//    req := client.DeleteConfigurationSetEventDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/DeleteConfigurationSetEventDestination
func (c *Client) DeleteConfigurationSetEventDestinationRequest(input *DeleteConfigurationSetEventDestinationInput) DeleteConfigurationSetEventDestinationRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationSetEventDestination,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/email/configuration-sets/{ConfigurationSetName}/event-destinations/{EventDestinationName}",
	}

	if input == nil {
		input = &DeleteConfigurationSetEventDestinationInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationSetEventDestinationOutput{})

	return DeleteConfigurationSetEventDestinationRequest{Request: req, Input: input, Copy: c.DeleteConfigurationSetEventDestinationRequest}
}

// DeleteConfigurationSetEventDestinationRequest is the request type for the
// DeleteConfigurationSetEventDestination API operation.
type DeleteConfigurationSetEventDestinationRequest struct {
	*aws.Request
	Input *DeleteConfigurationSetEventDestinationInput
	Copy  func(*DeleteConfigurationSetEventDestinationInput) DeleteConfigurationSetEventDestinationRequest
}

// Send marshals and sends the DeleteConfigurationSetEventDestination API request.
func (r DeleteConfigurationSetEventDestinationRequest) Send(ctx context.Context) (*DeleteConfigurationSetEventDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationSetEventDestinationResponse{
		DeleteConfigurationSetEventDestinationOutput: r.Request.Data.(*DeleteConfigurationSetEventDestinationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationSetEventDestinationResponse is the response type for the
// DeleteConfigurationSetEventDestination API operation.
type DeleteConfigurationSetEventDestinationResponse struct {
	*DeleteConfigurationSetEventDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationSetEventDestination request.
func (r *DeleteConfigurationSetEventDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
