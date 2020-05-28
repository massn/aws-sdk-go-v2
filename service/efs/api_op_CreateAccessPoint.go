// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateAccessPointInput struct {
	_ struct{} `type:"structure"`

	// A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent
	// creation.
	//
	// ClientToken is a required field
	ClientToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The ID of the EFS file system that the access point provides access to.
	//
	// FileSystemId is a required field
	FileSystemId *string `type:"string" required:"true"`

	// The operating system user and group applied to all file system requests made
	// using the access point.
	PosixUser *PosixUser `type:"structure"`

	// Specifies the directory on the Amazon EFS file system that the access point
	// exposes as the root directory of your file system to NFS clients using the
	// access point. The clients using the access point can only access the root
	// directory and below. If the RootDirectory > Path specified does not exist,
	// EFS creates it and applies the CreationInfo settings when a client connects
	// to an access point. When specifying a RootDirectory, you need to provide
	// the Path, and the CreationInfo is optional.
	RootDirectory *RootDirectory `type:"structure"`

	// Creates tags associated with the access point. Each tag is a key-value pair.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateAccessPointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAccessPointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAccessPointInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}
	if s.PosixUser != nil {
		if err := s.PosixUser.Validate(); err != nil {
			invalidParams.AddNested("PosixUser", err.(aws.ErrInvalidParams))
		}
	}
	if s.RootDirectory != nil {
		if err := s.RootDirectory.Validate(); err != nil {
			invalidParams.AddNested("RootDirectory", err.(aws.ErrInvalidParams))
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
func (s CreateAccessPointInput) MarshalFields(e protocol.FieldEncoder) error {
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
		e.SetValue(protocol.BodyTarget, "ClientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PosixUser != nil {
		v := s.PosixUser

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PosixUser", v, metadata)
	}
	if s.RootDirectory != nil {
		v := s.RootDirectory

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RootDirectory", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Provides a description of an EFS file system access point.
type CreateAccessPointOutput struct {
	_ struct{} `type:"structure"`

	// The unique Amazon Resource Name (ARN) associated with the access point.
	AccessPointArn *string `type:"string"`

	// The ID of the access point, assigned by Amazon EFS.
	AccessPointId *string `type:"string"`

	// The opaque string specified in the request to ensure idempotent creation.
	ClientToken *string `min:"1" type:"string"`

	// The ID of the EFS file system that the access point applies to.
	FileSystemId *string `type:"string"`

	// Identifies the lifecycle phase of the access point.
	LifeCycleState LifeCycleState `type:"string" enum:"true"`

	// The name of the access point. This is the value of the Name tag.
	Name *string `type:"string"`

	// Identified the AWS account that owns the access point resource.
	OwnerId *string `type:"string"`

	// The full POSIX identity, including the user ID, group ID, and secondary group
	// IDs on the access point that is used for all file operations by NFS clients
	// using the access point.
	PosixUser *PosixUser `type:"structure"`

	// The directory on the Amazon EFS file system that the access point exposes
	// as the root directory to NFS clients using the access point.
	RootDirectory *RootDirectory `type:"structure"`

	// The tags associated with the access point, presented as an array of Tag objects.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateAccessPointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAccessPointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccessPointArn != nil {
		v := *s.AccessPointArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AccessPointArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccessPointId != nil {
		v := *s.AccessPointId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AccessPointId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LifeCycleState) > 0 {
		v := s.LifeCycleState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LifeCycleState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OwnerId != nil {
		v := *s.OwnerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OwnerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PosixUser != nil {
		v := s.PosixUser

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PosixUser", v, metadata)
	}
	if s.RootDirectory != nil {
		v := s.RootDirectory

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RootDirectory", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opCreateAccessPoint = "CreateAccessPoint"

// CreateAccessPointRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Creates an EFS access point. An access point is an application-specific view
// into an EFS file system that applies an operating system user and group,
// and a file system path, to any file system request made through the access
// point. The operating system user and group override any identity information
// provided by the NFS client. The file system path is exposed as the access
// point's root directory. Applications using the access point can only access
// data in its own directory and below. To learn more, see Mounting a File System
// Using EFS Access Points (https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html).
//
// This operation requires permissions for the elasticfilesystem:CreateAccessPoint
// action.
//
//    // Example sending a request using CreateAccessPointRequest.
//    req := client.CreateAccessPointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/CreateAccessPoint
func (c *Client) CreateAccessPointRequest(input *CreateAccessPointInput) CreateAccessPointRequest {
	op := &aws.Operation{
		Name:       opCreateAccessPoint,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-02-01/access-points",
	}

	if input == nil {
		input = &CreateAccessPointInput{}
	}

	req := c.newRequest(op, input, &CreateAccessPointOutput{})

	return CreateAccessPointRequest{Request: req, Input: input, Copy: c.CreateAccessPointRequest}
}

// CreateAccessPointRequest is the request type for the
// CreateAccessPoint API operation.
type CreateAccessPointRequest struct {
	*aws.Request
	Input *CreateAccessPointInput
	Copy  func(*CreateAccessPointInput) CreateAccessPointRequest
}

// Send marshals and sends the CreateAccessPoint API request.
func (r CreateAccessPointRequest) Send(ctx context.Context) (*CreateAccessPointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAccessPointResponse{
		CreateAccessPointOutput: r.Request.Data.(*CreateAccessPointOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAccessPointResponse is the response type for the
// CreateAccessPoint API operation.
type CreateAccessPointResponse struct {
	*CreateAccessPointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAccessPoint request.
func (r *CreateAccessPointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
