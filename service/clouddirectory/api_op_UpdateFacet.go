// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateFacetInput struct {
	_ struct{} `type:"structure"`

	// List of attributes that need to be updated in a given schema Facet. Each
	// attribute is followed by AttributeAction, which specifies the type of update
	// operation to perform.
	AttributeUpdates []FacetAttributeUpdate `type:"list"`

	// The name of the facet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The object type that is associated with the facet. See CreateFacetRequest$ObjectType
	// for more details.
	ObjectType ObjectType `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Facet. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateFacetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFacetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFacetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}
	if s.AttributeUpdates != nil {
		for i, v := range s.AttributeUpdates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AttributeUpdates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFacetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AttributeUpdates != nil {
		v := s.AttributeUpdates

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AttributeUpdates", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ObjectType) > 0 {
		v := s.ObjectType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ObjectType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateFacetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateFacetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFacetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateFacet = "UpdateFacet"

// UpdateFacetRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Does the following:
//
// Adds new Attributes, Rules, or ObjectTypes.
//
// Updates existing Attributes, Rules, or ObjectTypes.
//
// Deletes existing Attributes, Rules, or ObjectTypes.
//
//    // Example sending a request using UpdateFacetRequest.
//    req := client.UpdateFacetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/UpdateFacet
func (c *Client) UpdateFacetRequest(input *UpdateFacetInput) UpdateFacetRequest {
	op := &aws.Operation{
		Name:       opUpdateFacet,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/facet",
	}

	if input == nil {
		input = &UpdateFacetInput{}
	}

	req := c.newRequest(op, input, &UpdateFacetOutput{})

	return UpdateFacetRequest{Request: req, Input: input, Copy: c.UpdateFacetRequest}
}

// UpdateFacetRequest is the request type for the
// UpdateFacet API operation.
type UpdateFacetRequest struct {
	*aws.Request
	Input *UpdateFacetInput
	Copy  func(*UpdateFacetInput) UpdateFacetRequest
}

// Send marshals and sends the UpdateFacet API request.
func (r UpdateFacetRequest) Send(ctx context.Context) (*UpdateFacetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFacetResponse{
		UpdateFacetOutput: r.Request.Data.(*UpdateFacetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFacetResponse is the response type for the
// UpdateFacet API operation.
type UpdateFacetResponse struct {
	*UpdateFacetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFacet request.
func (r *UpdateFacetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
