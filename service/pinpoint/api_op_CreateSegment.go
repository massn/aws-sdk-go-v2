// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateSegmentInput struct {
	_ struct{} `type:"structure" payload:"WriteSegmentRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the configuration, dimension, and other settings for a segment.
	// A WriteSegmentRequest object can include a Dimensions object or a SegmentGroups
	// object, but not both.
	//
	// WriteSegmentRequest is a required field
	WriteSegmentRequest *WriteSegmentRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateSegmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSegmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSegmentInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.WriteSegmentRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("WriteSegmentRequest"))
	}
	if s.WriteSegmentRequest != nil {
		if err := s.WriteSegmentRequest.Validate(); err != nil {
			invalidParams.AddNested("WriteSegmentRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSegmentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.WriteSegmentRequest != nil {
		v := s.WriteSegmentRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "WriteSegmentRequest", v, metadata)
	}
	return nil
}

type CreateSegmentOutput struct {
	_ struct{} `type:"structure" payload:"SegmentResponse"`

	// Provides information about the configuration, dimension, and other settings
	// for a segment.
	//
	// SegmentResponse is a required field
	SegmentResponse *SegmentResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateSegmentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSegmentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SegmentResponse != nil {
		v := s.SegmentResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SegmentResponse", v, metadata)
	}
	return nil
}

const opCreateSegment = "CreateSegment"

// CreateSegmentRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a new segment for an application or updates the configuration, dimension,
// and other settings for an existing segment that's associated with an application.
//
//    // Example sending a request using CreateSegmentRequest.
//    req := client.CreateSegmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/CreateSegment
func (c *Client) CreateSegmentRequest(input *CreateSegmentInput) CreateSegmentRequest {
	op := &aws.Operation{
		Name:       opCreateSegment,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apps/{application-id}/segments",
	}

	if input == nil {
		input = &CreateSegmentInput{}
	}

	req := c.newRequest(op, input, &CreateSegmentOutput{})

	return CreateSegmentRequest{Request: req, Input: input, Copy: c.CreateSegmentRequest}
}

// CreateSegmentRequest is the request type for the
// CreateSegment API operation.
type CreateSegmentRequest struct {
	*aws.Request
	Input *CreateSegmentInput
	Copy  func(*CreateSegmentInput) CreateSegmentRequest
}

// Send marshals and sends the CreateSegment API request.
func (r CreateSegmentRequest) Send(ctx context.Context) (*CreateSegmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSegmentResponse{
		CreateSegmentOutput: r.Request.Data.(*CreateSegmentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSegmentResponse is the response type for the
// CreateSegment API operation.
type CreateSegmentResponse struct {
	*CreateSegmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSegment request.
func (r *CreateSegmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
