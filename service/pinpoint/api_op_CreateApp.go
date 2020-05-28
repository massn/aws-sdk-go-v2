// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateAppInput struct {
	_ struct{} `type:"structure" payload:"CreateApplicationRequest"`

	// Specifies the display name of an application and the tags to associate with
	// the application.
	//
	// CreateApplicationRequest is a required field
	CreateApplicationRequest *CreateApplicationRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAppInput"}

	if s.CreateApplicationRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreateApplicationRequest"))
	}
	if s.CreateApplicationRequest != nil {
		if err := s.CreateApplicationRequest.Validate(); err != nil {
			invalidParams.AddNested("CreateApplicationRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAppInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CreateApplicationRequest != nil {
		v := s.CreateApplicationRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CreateApplicationRequest", v, metadata)
	}
	return nil
}

type CreateAppOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationResponse"`

	// Provides information about an application.
	//
	// ApplicationResponse is a required field
	ApplicationResponse *ApplicationResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateAppOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAppOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationResponse != nil {
		v := s.ApplicationResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ApplicationResponse", v, metadata)
	}
	return nil
}

const opCreateApp = "CreateApp"

// CreateAppRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates an application.
//
//    // Example sending a request using CreateAppRequest.
//    req := client.CreateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/CreateApp
func (c *Client) CreateAppRequest(input *CreateAppInput) CreateAppRequest {
	op := &aws.Operation{
		Name:       opCreateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps",
	}

	if input == nil {
		input = &CreateAppInput{}
	}

	req := c.newRequest(op, input, &CreateAppOutput{})

	return CreateAppRequest{Request: req, Input: input, Copy: c.CreateAppRequest}
}

// CreateAppRequest is the request type for the
// CreateApp API operation.
type CreateAppRequest struct {
	*aws.Request
	Input *CreateAppInput
	Copy  func(*CreateAppInput) CreateAppRequest
}

// Send marshals and sends the CreateApp API request.
func (r CreateAppRequest) Send(ctx context.Context) (*CreateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAppResponse{
		CreateAppOutput: r.Request.Data.(*CreateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAppResponse is the response type for the
// CreateApp API operation.
type CreateAppResponse struct {
	*CreateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApp request.
func (r *CreateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
