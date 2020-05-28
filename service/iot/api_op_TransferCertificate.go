// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the TransferCertificate operation.
type TransferCertificateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the certificate. (The last part of the certificate ARN contains
	// the certificate ID.)
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"certificateId" min:"64" type:"string" required:"true"`

	// The AWS account.
	//
	// TargetAwsAccount is a required field
	TargetAwsAccount *string `location:"querystring" locationName:"targetAwsAccount" min:"12" type:"string" required:"true"`

	// The transfer message.
	TransferMessage *string `locationName:"transferMessage" type:"string"`
}

// String returns the string representation
func (s TransferCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TransferCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TransferCertificateInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 64))
	}

	if s.TargetAwsAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetAwsAccount"))
	}
	if s.TargetAwsAccount != nil && len(*s.TargetAwsAccount) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetAwsAccount", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TransferCertificateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TransferMessage != nil {
		v := *s.TransferMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "transferMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TargetAwsAccount != nil {
		v := *s.TargetAwsAccount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "targetAwsAccount", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the TransferCertificate operation.
type TransferCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the certificate.
	TransferredCertificateArn *string `locationName:"transferredCertificateArn" type:"string"`
}

// String returns the string representation
func (s TransferCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TransferCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TransferredCertificateArn != nil {
		v := *s.TransferredCertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "transferredCertificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opTransferCertificate = "TransferCertificate"

// TransferCertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Transfers the specified certificate to the specified AWS account.
//
// You can cancel the transfer until it is acknowledged by the recipient.
//
// No notification is sent to the transfer destination's account. It is up to
// the caller to notify the transfer target.
//
// The certificate being transferred must not be in the ACTIVE state. You can
// use the UpdateCertificate API to deactivate it.
//
// The certificate must not have any policies attached to it. You can use the
// DetachPrincipalPolicy API to detach them.
//
//    // Example sending a request using TransferCertificateRequest.
//    req := client.TransferCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) TransferCertificateRequest(input *TransferCertificateInput) TransferCertificateRequest {
	op := &aws.Operation{
		Name:       opTransferCertificate,
		HTTPMethod: "PATCH",
		HTTPPath:   "/transfer-certificate/{certificateId}",
	}

	if input == nil {
		input = &TransferCertificateInput{}
	}

	req := c.newRequest(op, input, &TransferCertificateOutput{})

	return TransferCertificateRequest{Request: req, Input: input, Copy: c.TransferCertificateRequest}
}

// TransferCertificateRequest is the request type for the
// TransferCertificate API operation.
type TransferCertificateRequest struct {
	*aws.Request
	Input *TransferCertificateInput
	Copy  func(*TransferCertificateInput) TransferCertificateRequest
}

// Send marshals and sends the TransferCertificate API request.
func (r TransferCertificateRequest) Send(ctx context.Context) (*TransferCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TransferCertificateResponse{
		TransferCertificateOutput: r.Request.Data.(*TransferCertificateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TransferCertificateResponse is the response type for the
// TransferCertificate API operation.
type TransferCertificateResponse struct {
	*TransferCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TransferCertificate request.
func (r *TransferCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
