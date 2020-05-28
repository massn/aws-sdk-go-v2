// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateDomainConfigurationInput struct {
	_ struct{} `type:"structure"`

	// An object that specifies the authorization service for a domain.
	AuthorizerConfig *AuthorizerConfig `locationName:"authorizerConfig" type:"structure"`

	// The name of the domain configuration. This value must be unique to a region.
	//
	// DomainConfigurationName is a required field
	DomainConfigurationName *string `location:"uri" locationName:"domainConfigurationName" min:"1" type:"string" required:"true"`

	// The name of the domain.
	DomainName *string `locationName:"domainName" min:"1" type:"string"`

	// The ARNs of the certificates that AWS IoT passes to the device during the
	// TLS handshake. Currently you can specify only one certificate ARN. This value
	// is not required for AWS-managed domains.
	ServerCertificateArns []string `locationName:"serverCertificateArns" type:"list"`

	// The type of service delivered by the endpoint.
	//
	// AWS IoT Core currently supports only the DATA service type.
	ServiceType ServiceType `locationName:"serviceType" type:"string" enum:"true"`

	// Metadata which can be used to manage the domain configuration.
	//
	// For URI Request parameters use format: ...key1=value1&key2=value2...
	//
	// For the CLI command-line parameter use format: &&tags "key1=value1&key2=value2..."
	//
	// For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags []Tag `locationName:"tags" type:"list"`

	// The certificate used to validate the server certificate and prove domain
	// name ownership. This certificate must be signed by a public certificate authority.
	// This value is not required for AWS-managed domains.
	ValidationCertificateArn *string `locationName:"validationCertificateArn" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDomainConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDomainConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDomainConfigurationInput"}

	if s.DomainConfigurationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainConfigurationName"))
	}
	if s.DomainConfigurationName != nil && len(*s.DomainConfigurationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainConfigurationName", 1))
	}
	if s.DomainName != nil && len(*s.DomainName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 1))
	}
	if s.ValidationCertificateArn != nil && len(*s.ValidationCertificateArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ValidationCertificateArn", 1))
	}
	if s.AuthorizerConfig != nil {
		if err := s.AuthorizerConfig.Validate(); err != nil {
			invalidParams.AddNested("AuthorizerConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthorizerConfig != nil {
		v := s.AuthorizerConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "authorizerConfig", v, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ServerCertificateArns != nil {
		v := s.ServerCertificateArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "serverCertificateArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.ServiceType) > 0 {
		v := s.ServiceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serviceType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ValidationCertificateArn != nil {
		v := *s.ValidationCertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "validationCertificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainConfigurationName != nil {
		v := *s.DomainConfigurationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainConfigurationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateDomainConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the domain configuration.
	DomainConfigurationArn *string `locationName:"domainConfigurationArn" type:"string"`

	// The name of the domain configuration.
	DomainConfigurationName *string `locationName:"domainConfigurationName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDomainConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDomainConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainConfigurationArn != nil {
		v := *s.DomainConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainConfigurationName != nil {
		v := *s.DomainConfigurationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainConfigurationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateDomainConfiguration = "CreateDomainConfiguration"

// CreateDomainConfigurationRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a domain configuration.
//
// The domain configuration feature is in public preview and is subject to change.
//
//    // Example sending a request using CreateDomainConfigurationRequest.
//    req := client.CreateDomainConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDomainConfigurationRequest(input *CreateDomainConfigurationInput) CreateDomainConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateDomainConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/domainConfigurations/{domainConfigurationName}",
	}

	if input == nil {
		input = &CreateDomainConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateDomainConfigurationOutput{})

	return CreateDomainConfigurationRequest{Request: req, Input: input, Copy: c.CreateDomainConfigurationRequest}
}

// CreateDomainConfigurationRequest is the request type for the
// CreateDomainConfiguration API operation.
type CreateDomainConfigurationRequest struct {
	*aws.Request
	Input *CreateDomainConfigurationInput
	Copy  func(*CreateDomainConfigurationInput) CreateDomainConfigurationRequest
}

// Send marshals and sends the CreateDomainConfiguration API request.
func (r CreateDomainConfigurationRequest) Send(ctx context.Context) (*CreateDomainConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainConfigurationResponse{
		CreateDomainConfigurationOutput: r.Request.Data.(*CreateDomainConfigurationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDomainConfigurationResponse is the response type for the
// CreateDomainConfiguration API operation.
type CreateDomainConfigurationResponse struct {
	*CreateDomainConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDomainConfiguration request.
func (r *CreateDomainConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
