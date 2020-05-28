// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateConfigInput struct {
	_ struct{} `type:"structure"`

	// Parameters of a Config.
	//
	// ConfigData is a required field
	ConfigData *ConfigTypeData `locationName:"configData" type:"structure" required:"true"`

	// Name of a Config.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// Tags assigned to a Config.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigInput"}

	if s.ConfigData == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigData"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.ConfigData != nil {
		if err := s.ConfigData.Validate(); err != nil {
			invalidParams.AddNested("ConfigData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ConfigData != nil {
		v := s.ConfigData

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configData", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

type CreateConfigOutput struct {
	_ struct{} `type:"structure"`

	// ARN of a Config.
	ConfigArn *string `locationName:"configArn" type:"string"`

	// UUID of a Config.
	ConfigId *string `locationName:"configId" type:"string"`

	// Type of a Config.
	ConfigType ConfigCapabilityType `locationName:"configType" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ConfigArn != nil {
		v := *s.ConfigArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigId != nil {
		v := *s.ConfigId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConfigType) > 0 {
		v := s.ConfigType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "configType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opCreateConfig = "CreateConfig"

// CreateConfigRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Creates a Config with the specified configData parameters.
//
// Only one type of configData can be specified.
//
//    // Example sending a request using CreateConfigRequest.
//    req := client.CreateConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/CreateConfig
func (c *Client) CreateConfigRequest(input *CreateConfigInput) CreateConfigRequest {
	op := &aws.Operation{
		Name:       opCreateConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/config",
	}

	if input == nil {
		input = &CreateConfigInput{}
	}

	req := c.newRequest(op, input, &CreateConfigOutput{})

	return CreateConfigRequest{Request: req, Input: input, Copy: c.CreateConfigRequest}
}

// CreateConfigRequest is the request type for the
// CreateConfig API operation.
type CreateConfigRequest struct {
	*aws.Request
	Input *CreateConfigInput
	Copy  func(*CreateConfigInput) CreateConfigRequest
}

// Send marshals and sends the CreateConfig API request.
func (r CreateConfigRequest) Send(ctx context.Context) (*CreateConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConfigResponse{
		CreateConfigOutput: r.Request.Data.(*CreateConfigOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConfigResponse is the response type for the
// CreateConfig API operation.
type CreateConfigResponse struct {
	*CreateConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConfig request.
func (r *CreateConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
