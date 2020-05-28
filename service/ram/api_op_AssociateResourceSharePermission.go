// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type AssociateResourceSharePermissionInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The ARN of the AWS RAM permission to associate with the resource share.
	//
	// PermissionArn is a required field
	PermissionArn *string `locationName:"permissionArn" type:"string" required:"true"`

	// Indicates whether the permission should replace the permissions that are
	// currently associated with the resource share. Use true to replace the current
	// permissions. Use false to add the permission to the current permission.
	Replace *bool `locationName:"replace" type:"boolean"`

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// ResourceShareArn is a required field
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateResourceSharePermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateResourceSharePermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateResourceSharePermissionInput"}

	if s.PermissionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PermissionArn"))
	}

	if s.ResourceShareArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateResourceSharePermissionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PermissionArn != nil {
		v := *s.PermissionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "permissionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Replace != nil {
		v := *s.Replace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "replace", protocol.BoolValue(v), metadata)
	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AssociateResourceSharePermissionOutput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Indicates whether the request succeeded.
	ReturnValue *bool `locationName:"returnValue" type:"boolean"`
}

// String returns the string representation
func (s AssociateResourceSharePermissionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateResourceSharePermissionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ReturnValue != nil {
		v := *s.ReturnValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "returnValue", protocol.BoolValue(v), metadata)
	}
	return nil
}

const opAssociateResourceSharePermission = "AssociateResourceSharePermission"

// AssociateResourceSharePermissionRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Associates a permission with a resource share.
//
//    // Example sending a request using AssociateResourceSharePermissionRequest.
//    req := client.AssociateResourceSharePermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AssociateResourceSharePermission
func (c *Client) AssociateResourceSharePermissionRequest(input *AssociateResourceSharePermissionInput) AssociateResourceSharePermissionRequest {
	op := &aws.Operation{
		Name:       opAssociateResourceSharePermission,
		HTTPMethod: "POST",
		HTTPPath:   "/associateresourcesharepermission",
	}

	if input == nil {
		input = &AssociateResourceSharePermissionInput{}
	}

	req := c.newRequest(op, input, &AssociateResourceSharePermissionOutput{})

	return AssociateResourceSharePermissionRequest{Request: req, Input: input, Copy: c.AssociateResourceSharePermissionRequest}
}

// AssociateResourceSharePermissionRequest is the request type for the
// AssociateResourceSharePermission API operation.
type AssociateResourceSharePermissionRequest struct {
	*aws.Request
	Input *AssociateResourceSharePermissionInput
	Copy  func(*AssociateResourceSharePermissionInput) AssociateResourceSharePermissionRequest
}

// Send marshals and sends the AssociateResourceSharePermission API request.
func (r AssociateResourceSharePermissionRequest) Send(ctx context.Context) (*AssociateResourceSharePermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateResourceSharePermissionResponse{
		AssociateResourceSharePermissionOutput: r.Request.Data.(*AssociateResourceSharePermissionOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateResourceSharePermissionResponse is the response type for the
// AssociateResourceSharePermission API operation.
type AssociateResourceSharePermissionResponse struct {
	*AssociateResourceSharePermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateResourceSharePermission request.
func (r *AssociateResourceSharePermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
