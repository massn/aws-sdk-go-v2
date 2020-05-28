// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateFieldLevelEncryptionConfigInput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionConfig"`

	// The request to create a new field-level encryption configuration.
	//
	// FieldLevelEncryptionConfig is a required field
	FieldLevelEncryptionConfig *FieldLevelEncryptionConfig `locationName:"FieldLevelEncryptionConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2019-03-26/"`
}

// String returns the string representation
func (s CreateFieldLevelEncryptionConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFieldLevelEncryptionConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFieldLevelEncryptionConfigInput"}

	if s.FieldLevelEncryptionConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("FieldLevelEncryptionConfig"))
	}
	if s.FieldLevelEncryptionConfig != nil {
		if err := s.FieldLevelEncryptionConfig.Validate(); err != nil {
			invalidParams.AddNested("FieldLevelEncryptionConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFieldLevelEncryptionConfigInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FieldLevelEncryptionConfig != nil {
		v := s.FieldLevelEncryptionConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2019-03-26/"}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionConfig", v, metadata)
	}
	return nil
}

type CreateFieldLevelEncryptionConfigOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryption"`

	// The current version of the field level encryption configuration. For example:
	// E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Returned when you create a new field-level encryption configuration.
	FieldLevelEncryption *FieldLevelEncryption `type:"structure"`

	// The fully qualified URI of the new configuration resource just created. For
	// example: https://cloudfront.amazonaws.com/2010-11-01/field-level-encryption-config/EDFDVBD632BHDS5.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s CreateFieldLevelEncryptionConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFieldLevelEncryptionConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	if s.FieldLevelEncryption != nil {
		v := s.FieldLevelEncryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryption", v, metadata)
	}
	return nil
}

const opCreateFieldLevelEncryptionConfig = "CreateFieldLevelEncryptionConfig2019_03_26"

// CreateFieldLevelEncryptionConfigRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Create a new field-level encryption configuration.
//
//    // Example sending a request using CreateFieldLevelEncryptionConfigRequest.
//    req := client.CreateFieldLevelEncryptionConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/CreateFieldLevelEncryptionConfig
func (c *Client) CreateFieldLevelEncryptionConfigRequest(input *CreateFieldLevelEncryptionConfigInput) CreateFieldLevelEncryptionConfigRequest {
	op := &aws.Operation{
		Name:       opCreateFieldLevelEncryptionConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-03-26/field-level-encryption",
	}

	if input == nil {
		input = &CreateFieldLevelEncryptionConfigInput{}
	}

	req := c.newRequest(op, input, &CreateFieldLevelEncryptionConfigOutput{})

	return CreateFieldLevelEncryptionConfigRequest{Request: req, Input: input, Copy: c.CreateFieldLevelEncryptionConfigRequest}
}

// CreateFieldLevelEncryptionConfigRequest is the request type for the
// CreateFieldLevelEncryptionConfig API operation.
type CreateFieldLevelEncryptionConfigRequest struct {
	*aws.Request
	Input *CreateFieldLevelEncryptionConfigInput
	Copy  func(*CreateFieldLevelEncryptionConfigInput) CreateFieldLevelEncryptionConfigRequest
}

// Send marshals and sends the CreateFieldLevelEncryptionConfig API request.
func (r CreateFieldLevelEncryptionConfigRequest) Send(ctx context.Context) (*CreateFieldLevelEncryptionConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFieldLevelEncryptionConfigResponse{
		CreateFieldLevelEncryptionConfigOutput: r.Request.Data.(*CreateFieldLevelEncryptionConfigOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFieldLevelEncryptionConfigResponse is the response type for the
// CreateFieldLevelEncryptionConfig API operation.
type CreateFieldLevelEncryptionConfigResponse struct {
	*CreateFieldLevelEncryptionConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFieldLevelEncryptionConfig request.
func (r *CreateFieldLevelEncryptionConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
