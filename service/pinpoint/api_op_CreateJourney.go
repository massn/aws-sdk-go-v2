// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateJourneyInput struct {
	_ struct{} `type:"structure" payload:"WriteJourneyRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the configuration and other settings for a journey.
	//
	// WriteJourneyRequest is a required field
	WriteJourneyRequest *WriteJourneyRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateJourneyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateJourneyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateJourneyInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.WriteJourneyRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("WriteJourneyRequest"))
	}
	if s.WriteJourneyRequest != nil {
		if err := s.WriteJourneyRequest.Validate(); err != nil {
			invalidParams.AddNested("WriteJourneyRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJourneyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.WriteJourneyRequest != nil {
		v := s.WriteJourneyRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "WriteJourneyRequest", v, metadata)
	}
	return nil
}

type CreateJourneyOutput struct {
	_ struct{} `type:"structure" payload:"JourneyResponse"`

	// Provides information about the status, configuration, and other settings
	// for a journey.
	//
	// JourneyResponse is a required field
	JourneyResponse *JourneyResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateJourneyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateJourneyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JourneyResponse != nil {
		v := s.JourneyResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "JourneyResponse", v, metadata)
	}
	return nil
}

const opCreateJourney = "CreateJourney"

// CreateJourneyRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a journey for an application.
//
//    // Example sending a request using CreateJourneyRequest.
//    req := client.CreateJourneyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/CreateJourney
func (c *Client) CreateJourneyRequest(input *CreateJourneyInput) CreateJourneyRequest {
	op := &aws.Operation{
		Name:       opCreateJourney,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps/{application-id}/journeys",
	}

	if input == nil {
		input = &CreateJourneyInput{}
	}

	req := c.newRequest(op, input, &CreateJourneyOutput{})

	return CreateJourneyRequest{Request: req, Input: input, Copy: c.CreateJourneyRequest}
}

// CreateJourneyRequest is the request type for the
// CreateJourney API operation.
type CreateJourneyRequest struct {
	*aws.Request
	Input *CreateJourneyInput
	Copy  func(*CreateJourneyInput) CreateJourneyRequest
}

// Send marshals and sends the CreateJourney API request.
func (r CreateJourneyRequest) Send(ctx context.Context) (*CreateJourneyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateJourneyResponse{
		CreateJourneyOutput: r.Request.Data.(*CreateJourneyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateJourneyResponse is the response type for the
// CreateJourney API operation.
type CreateJourneyResponse struct {
	*CreateJourneyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateJourney request.
func (r *CreateJourneyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
