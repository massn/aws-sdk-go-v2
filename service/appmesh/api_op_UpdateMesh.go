// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateMeshInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// An object that represents the specification of a service mesh.
	Spec *MeshSpec `locationName:"spec" type:"structure"`
}

// String returns the string representation
func (s UpdateMeshInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMeshInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMeshInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			invalidParams.AddNested("Spec", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMeshInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Spec != nil {
		v := s.Spec

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spec", v, metadata)
	}
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateMeshOutput struct {
	_ struct{} `type:"structure" payload:"Mesh"`

	// An object that represents a service mesh returned by a describe operation.
	//
	// Mesh is a required field
	Mesh *MeshData `locationName:"mesh" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateMeshOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMeshOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Mesh != nil {
		v := s.Mesh

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "mesh", v, metadata)
	}
	return nil
}

const opUpdateMesh = "UpdateMesh"

// UpdateMeshRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Updates an existing service mesh.
//
//    // Example sending a request using UpdateMeshRequest.
//    req := client.UpdateMeshRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/UpdateMesh
func (c *Client) UpdateMeshRequest(input *UpdateMeshInput) UpdateMeshRequest {
	op := &aws.Operation{
		Name:       opUpdateMesh,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}",
	}

	if input == nil {
		input = &UpdateMeshInput{}
	}

	req := c.newRequest(op, input, &UpdateMeshOutput{})

	return UpdateMeshRequest{Request: req, Input: input, Copy: c.UpdateMeshRequest}
}

// UpdateMeshRequest is the request type for the
// UpdateMesh API operation.
type UpdateMeshRequest struct {
	*aws.Request
	Input *UpdateMeshInput
	Copy  func(*UpdateMeshInput) UpdateMeshRequest
}

// Send marshals and sends the UpdateMesh API request.
func (r UpdateMeshRequest) Send(ctx context.Context) (*UpdateMeshResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMeshResponse{
		UpdateMeshOutput: r.Request.Data.(*UpdateMeshOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMeshResponse is the response type for the
// UpdateMesh API operation.
type UpdateMeshResponse struct {
	*UpdateMeshOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMesh request.
func (r *UpdateMeshResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
