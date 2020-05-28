// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteSchemaVersionInput struct {
	_ struct{} `type:"structure"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// SchemaName is a required field
	SchemaName *string `location:"uri" locationName:"schemaName" type:"string" required:"true"`

	// SchemaVersion is a required field
	SchemaVersion *string `location:"uri" locationName:"schemaVersion" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSchemaVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSchemaVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSchemaVersionInput"}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if s.SchemaName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaName"))
	}

	if s.SchemaVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaVersion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSchemaVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "schemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "schemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSchemaVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSchemaVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSchemaVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteSchemaVersion = "DeleteSchemaVersion"

// DeleteSchemaVersionRequest returns a request value for making API operation for
// Schemas.
//
// Delete the schema version definition
//
//    // Example sending a request using DeleteSchemaVersionRequest.
//    req := client.DeleteSchemaVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/DeleteSchemaVersion
func (c *Client) DeleteSchemaVersionRequest(input *DeleteSchemaVersionInput) DeleteSchemaVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteSchemaVersion,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas/name/{schemaName}/version/{schemaVersion}",
	}

	if input == nil {
		input = &DeleteSchemaVersionInput{}
	}

	req := c.newRequest(op, input, &DeleteSchemaVersionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteSchemaVersionRequest{Request: req, Input: input, Copy: c.DeleteSchemaVersionRequest}
}

// DeleteSchemaVersionRequest is the request type for the
// DeleteSchemaVersion API operation.
type DeleteSchemaVersionRequest struct {
	*aws.Request
	Input *DeleteSchemaVersionInput
	Copy  func(*DeleteSchemaVersionInput) DeleteSchemaVersionRequest
}

// Send marshals and sends the DeleteSchemaVersion API request.
func (r DeleteSchemaVersionRequest) Send(ctx context.Context) (*DeleteSchemaVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSchemaVersionResponse{
		DeleteSchemaVersionOutput: r.Request.Data.(*DeleteSchemaVersionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSchemaVersionResponse is the response type for the
// DeleteSchemaVersion API operation.
type DeleteSchemaVersionResponse struct {
	*DeleteSchemaVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSchemaVersion request.
func (r *DeleteSchemaVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
