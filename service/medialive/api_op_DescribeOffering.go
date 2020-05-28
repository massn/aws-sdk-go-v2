// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeOfferingInput struct {
	_ struct{} `type:"structure"`

	// OfferingId is a required field
	OfferingId *string `location:"uri" locationName:"offeringId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeOfferingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOfferingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOfferingInput"}

	if s.OfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeOfferingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.OfferingId != nil {
		v := *s.OfferingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "offeringId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeOfferingOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	CurrencyCode *string `locationName:"currencyCode" type:"string"`

	Duration *int64 `locationName:"duration" type:"integer"`

	// Units for duration, e.g. 'MONTHS'
	DurationUnits OfferingDurationUnits `locationName:"durationUnits" type:"string" enum:"true"`

	FixedPrice *float64 `locationName:"fixedPrice" type:"double"`

	OfferingDescription *string `locationName:"offeringDescription" type:"string"`

	OfferingId *string `locationName:"offeringId" type:"string"`

	// Offering type, e.g. 'NO_UPFRONT'
	OfferingType OfferingType `locationName:"offeringType" type:"string" enum:"true"`

	Region *string `locationName:"region" type:"string"`

	// Resource configuration (codec, resolution, bitrate, ...)
	ResourceSpecification *ReservationResourceSpecification `locationName:"resourceSpecification" type:"structure"`

	UsagePrice *float64 `locationName:"usagePrice" type:"double"`
}

// String returns the string representation
func (s DescribeOfferingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeOfferingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CurrencyCode != nil {
		v := *s.CurrencyCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "currencyCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Duration != nil {
		v := *s.Duration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "duration", protocol.Int64Value(v), metadata)
	}
	if len(s.DurationUnits) > 0 {
		v := s.DurationUnits

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "durationUnits", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FixedPrice != nil {
		v := *s.FixedPrice

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fixedPrice", protocol.Float64Value(v), metadata)
	}
	if s.OfferingDescription != nil {
		v := *s.OfferingDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "offeringDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OfferingId != nil {
		v := *s.OfferingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "offeringId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.OfferingType) > 0 {
		v := s.OfferingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "offeringType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Region != nil {
		v := *s.Region

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "region", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceSpecification != nil {
		v := s.ResourceSpecification

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourceSpecification", v, metadata)
	}
	if s.UsagePrice != nil {
		v := *s.UsagePrice

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "usagePrice", protocol.Float64Value(v), metadata)
	}
	return nil
}

const opDescribeOffering = "DescribeOffering"

// DescribeOfferingRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Get details for an offering.
//
//    // Example sending a request using DescribeOfferingRequest.
//    req := client.DescribeOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/DescribeOffering
func (c *Client) DescribeOfferingRequest(input *DescribeOfferingInput) DescribeOfferingRequest {
	op := &aws.Operation{
		Name:       opDescribeOffering,
		HTTPMethod: "GET",
		HTTPPath:   "/prod/offerings/{offeringId}",
	}

	if input == nil {
		input = &DescribeOfferingInput{}
	}

	req := c.newRequest(op, input, &DescribeOfferingOutput{})

	return DescribeOfferingRequest{Request: req, Input: input, Copy: c.DescribeOfferingRequest}
}

// DescribeOfferingRequest is the request type for the
// DescribeOffering API operation.
type DescribeOfferingRequest struct {
	*aws.Request
	Input *DescribeOfferingInput
	Copy  func(*DescribeOfferingInput) DescribeOfferingRequest
}

// Send marshals and sends the DescribeOffering API request.
func (r DescribeOfferingRequest) Send(ctx context.Context) (*DescribeOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOfferingResponse{
		DescribeOfferingOutput: r.Request.Data.(*DescribeOfferingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOfferingResponse is the response type for the
// DescribeOffering API operation.
type DescribeOfferingResponse struct {
	*DescribeOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOffering request.
func (r *DescribeOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
