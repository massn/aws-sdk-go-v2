// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PurchaseOfferingInput struct {
	_ struct{} `type:"structure"`

	// Count is a required field
	Count *int64 `locationName:"count" min:"1" type:"integer" required:"true"`

	Name *string `locationName:"name" type:"string"`

	// OfferingId is a required field
	OfferingId *string `location:"uri" locationName:"offeringId" type:"string" required:"true"`

	RequestId *string `locationName:"requestId" type:"string" idempotencyToken:"true"`

	Start *string `locationName:"start" type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s PurchaseOfferingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseOfferingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseOfferingInput"}

	if s.Count == nil {
		invalidParams.Add(aws.NewErrParamRequired("Count"))
	}
	if s.Count != nil && *s.Count < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Count", 1))
	}

	if s.OfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PurchaseOfferingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Count != nil {
		v := *s.Count

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "count", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	var RequestId string
	if s.RequestId != nil {
		RequestId = *s.RequestId
	} else {
		RequestId = protocol.GetIdempotencyToken()
	}
	{
		v := RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Start != nil {
		v := *s.Start

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "start", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.OfferingId != nil {
		v := *s.OfferingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "offeringId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PurchaseOfferingOutput struct {
	_ struct{} `type:"structure"`

	// Reserved resources available to use
	Reservation *Reservation `locationName:"reservation" type:"structure"`
}

// String returns the string representation
func (s PurchaseOfferingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PurchaseOfferingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Reservation != nil {
		v := s.Reservation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "reservation", v, metadata)
	}
	return nil
}

const opPurchaseOffering = "PurchaseOffering"

// PurchaseOfferingRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Purchase an offering and create a reservation.
//
//    // Example sending a request using PurchaseOfferingRequest.
//    req := client.PurchaseOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/PurchaseOffering
func (c *Client) PurchaseOfferingRequest(input *PurchaseOfferingInput) PurchaseOfferingRequest {
	op := &aws.Operation{
		Name:       opPurchaseOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/prod/offerings/{offeringId}/purchase",
	}

	if input == nil {
		input = &PurchaseOfferingInput{}
	}

	req := c.newRequest(op, input, &PurchaseOfferingOutput{})

	return PurchaseOfferingRequest{Request: req, Input: input, Copy: c.PurchaseOfferingRequest}
}

// PurchaseOfferingRequest is the request type for the
// PurchaseOffering API operation.
type PurchaseOfferingRequest struct {
	*aws.Request
	Input *PurchaseOfferingInput
	Copy  func(*PurchaseOfferingInput) PurchaseOfferingRequest
}

// Send marshals and sends the PurchaseOffering API request.
func (r PurchaseOfferingRequest) Send(ctx context.Context) (*PurchaseOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseOfferingResponse{
		PurchaseOfferingOutput: r.Request.Data.(*PurchaseOfferingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseOfferingResponse is the response type for the
// PurchaseOffering API operation.
type PurchaseOfferingResponse struct {
	*PurchaseOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseOffering request.
func (r *PurchaseOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
