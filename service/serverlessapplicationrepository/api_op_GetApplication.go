// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetApplicationInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`

	SemanticVersion *string `location:"querystring" locationName:"semanticVersion" type:"string"`
}

// String returns the string representation
func (s GetApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetApplicationInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SemanticVersion != nil {
		v := *s.SemanticVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "semanticVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetApplicationOutput struct {
	_ struct{} `type:"structure"`

	ApplicationId *string `locationName:"applicationId" type:"string"`

	Author *string `locationName:"author" type:"string"`

	CreationTime *string `locationName:"creationTime" type:"string"`

	Description *string `locationName:"description" type:"string"`

	HomePageUrl *string `locationName:"homePageUrl" type:"string"`

	IsVerifiedAuthor *bool `locationName:"isVerifiedAuthor" type:"boolean"`

	Labels []string `locationName:"labels" type:"list"`

	LicenseUrl *string `locationName:"licenseUrl" type:"string"`

	Name *string `locationName:"name" type:"string"`

	ReadmeUrl *string `locationName:"readmeUrl" type:"string"`

	SpdxLicenseId *string `locationName:"spdxLicenseId" type:"string"`

	VerifiedAuthorUrl *string `locationName:"verifiedAuthorUrl" type:"string"`

	// Application version details.
	Version *Version `locationName:"version" type:"structure"`
}

// String returns the string representation
func (s GetApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Author != nil {
		v := *s.Author

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "author", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HomePageUrl != nil {
		v := *s.HomePageUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "homePageUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IsVerifiedAuthor != nil {
		v := *s.IsVerifiedAuthor

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "isVerifiedAuthor", protocol.BoolValue(v), metadata)
	}
	if s.Labels != nil {
		v := s.Labels

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "labels", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.LicenseUrl != nil {
		v := *s.LicenseUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "licenseUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReadmeUrl != nil {
		v := *s.ReadmeUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "readmeUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SpdxLicenseId != nil {
		v := *s.SpdxLicenseId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "spdxLicenseId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VerifiedAuthorUrl != nil {
		v := *s.VerifiedAuthorUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "verifiedAuthorUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := s.Version

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "version", v, metadata)
	}
	return nil
}

const opGetApplication = "GetApplication"

// GetApplicationRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Gets the specified application.
//
//    // Example sending a request using GetApplicationRequest.
//    req := client.GetApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/GetApplication
func (c *Client) GetApplicationRequest(input *GetApplicationInput) GetApplicationRequest {
	op := &aws.Operation{
		Name:       opGetApplication,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{applicationId}",
	}

	if input == nil {
		input = &GetApplicationInput{}
	}

	req := c.newRequest(op, input, &GetApplicationOutput{})

	return GetApplicationRequest{Request: req, Input: input, Copy: c.GetApplicationRequest}
}

// GetApplicationRequest is the request type for the
// GetApplication API operation.
type GetApplicationRequest struct {
	*aws.Request
	Input *GetApplicationInput
	Copy  func(*GetApplicationInput) GetApplicationRequest
}

// Send marshals and sends the GetApplication API request.
func (r GetApplicationRequest) Send(ctx context.Context) (*GetApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApplicationResponse{
		GetApplicationOutput: r.Request.Data.(*GetApplicationOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApplicationResponse is the response type for the
// GetApplication API operation.
type GetApplicationResponse struct {
	*GetApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApplication request.
func (r *GetApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
