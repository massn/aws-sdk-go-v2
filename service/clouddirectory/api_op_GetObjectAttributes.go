// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetObjectAttributesInput struct {
	_ struct{} `type:"structure"`

	// List of attribute names whose values will be retrieved.
	//
	// AttributeNames is a required field
	AttributeNames []string `type:"list" required:"true"`

	// The consistency level at which to retrieve the attributes on an object.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the object resides.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Reference that identifies the object whose attributes will be retrieved.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`

	// Identifier for the facet whose attributes will be retrieved. See SchemaFacet
	// for details.
	//
	// SchemaFacet is a required field
	SchemaFacet *SchemaFacet `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetObjectAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectAttributesInput"}

	if s.AttributeNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeNames"))
	}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if s.SchemaFacet == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaFacet"))
	}
	if s.SchemaFacet != nil {
		if err := s.SchemaFacet.Validate(); err != nil {
			invalidParams.AddNested("SchemaFacet", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AttributeNames != nil {
		v := s.AttributeNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AttributeNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if s.SchemaFacet != nil {
		v := s.SchemaFacet

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SchemaFacet", v, metadata)
	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetObjectAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The attributes that are associated with the object.
	Attributes []AttributeKeyAndValue `type:"list"`
}

// String returns the string representation
func (s GetObjectAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attributes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetObjectAttributes = "GetObjectAttributes"

// GetObjectAttributesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Retrieves attributes within a facet that are associated with an object.
//
//    // Example sending a request using GetObjectAttributesRequest.
//    req := client.GetObjectAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/GetObjectAttributes
func (c *Client) GetObjectAttributesRequest(input *GetObjectAttributesInput) GetObjectAttributesRequest {
	op := &aws.Operation{
		Name:       opGetObjectAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/attributes/get",
	}

	if input == nil {
		input = &GetObjectAttributesInput{}
	}

	req := c.newRequest(op, input, &GetObjectAttributesOutput{})

	return GetObjectAttributesRequest{Request: req, Input: input, Copy: c.GetObjectAttributesRequest}
}

// GetObjectAttributesRequest is the request type for the
// GetObjectAttributes API operation.
type GetObjectAttributesRequest struct {
	*aws.Request
	Input *GetObjectAttributesInput
	Copy  func(*GetObjectAttributesInput) GetObjectAttributesRequest
}

// Send marshals and sends the GetObjectAttributes API request.
func (r GetObjectAttributesRequest) Send(ctx context.Context) (*GetObjectAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectAttributesResponse{
		GetObjectAttributesOutput: r.Request.Data.(*GetObjectAttributesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectAttributesResponse is the response type for the
// GetObjectAttributes API operation.
type GetObjectAttributesResponse struct {
	*GetObjectAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectAttributes request.
func (r *GetObjectAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
