// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeDevicePolicyConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDevicePolicyConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDevicePolicyConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDevicePolicyConfigurationInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDevicePolicyConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeDevicePolicyConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The certificate chain, including intermediate certificates and the root certificate
	// authority certificate used to issue device certificates.
	DeviceCaCertificate *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDevicePolicyConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDevicePolicyConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeviceCaCertificate != nil {
		v := *s.DeviceCaCertificate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeviceCaCertificate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeDevicePolicyConfiguration = "DescribeDevicePolicyConfiguration"

// DescribeDevicePolicyConfigurationRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Describes the device policy configuration for the specified fleet.
//
//    // Example sending a request using DescribeDevicePolicyConfigurationRequest.
//    req := client.DescribeDevicePolicyConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeDevicePolicyConfiguration
func (c *Client) DescribeDevicePolicyConfigurationRequest(input *DescribeDevicePolicyConfigurationInput) DescribeDevicePolicyConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeDevicePolicyConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/describeDevicePolicyConfiguration",
	}

	if input == nil {
		input = &DescribeDevicePolicyConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeDevicePolicyConfigurationOutput{})

	return DescribeDevicePolicyConfigurationRequest{Request: req, Input: input, Copy: c.DescribeDevicePolicyConfigurationRequest}
}

// DescribeDevicePolicyConfigurationRequest is the request type for the
// DescribeDevicePolicyConfiguration API operation.
type DescribeDevicePolicyConfigurationRequest struct {
	*aws.Request
	Input *DescribeDevicePolicyConfigurationInput
	Copy  func(*DescribeDevicePolicyConfigurationInput) DescribeDevicePolicyConfigurationRequest
}

// Send marshals and sends the DescribeDevicePolicyConfiguration API request.
func (r DescribeDevicePolicyConfigurationRequest) Send(ctx context.Context) (*DescribeDevicePolicyConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDevicePolicyConfigurationResponse{
		DescribeDevicePolicyConfigurationOutput: r.Request.Data.(*DescribeDevicePolicyConfigurationOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDevicePolicyConfigurationResponse is the response type for the
// DescribeDevicePolicyConfiguration API operation.
type DescribeDevicePolicyConfigurationResponse struct {
	*DescribeDevicePolicyConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDevicePolicyConfiguration request.
func (r *DescribeDevicePolicyConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
