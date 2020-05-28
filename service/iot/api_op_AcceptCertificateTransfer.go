// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the AcceptCertificateTransfer operation.
type AcceptCertificateTransferInput struct {
	_ struct{} `type:"structure"`

	// The ID of the certificate. (The last part of the certificate ARN contains
	// the certificate ID.)
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"certificateId" min:"64" type:"string" required:"true"`

	// Specifies whether the certificate is active.
	SetAsActive *bool `location:"querystring" locationName:"setAsActive" type:"boolean"`
}

// String returns the string representation
func (s AcceptCertificateTransferInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptCertificateTransferInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptCertificateTransferInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 64))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptCertificateTransferInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SetAsActive != nil {
		v := *s.SetAsActive

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "setAsActive", protocol.BoolValue(v), metadata)
	}
	return nil
}

type AcceptCertificateTransferOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AcceptCertificateTransferOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptCertificateTransferOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAcceptCertificateTransfer = "AcceptCertificateTransfer"

// AcceptCertificateTransferRequest returns a request value for making API operation for
// AWS IoT.
//
// Accepts a pending certificate transfer. The default state of the certificate
// is INACTIVE.
//
// To check for pending certificate transfers, call ListCertificates to enumerate
// your certificates.
//
//    // Example sending a request using AcceptCertificateTransferRequest.
//    req := client.AcceptCertificateTransferRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AcceptCertificateTransferRequest(input *AcceptCertificateTransferInput) AcceptCertificateTransferRequest {
	op := &aws.Operation{
		Name:       opAcceptCertificateTransfer,
		HTTPMethod: "PATCH",
		HTTPPath:   "/accept-certificate-transfer/{certificateId}",
	}

	if input == nil {
		input = &AcceptCertificateTransferInput{}
	}

	req := c.newRequest(op, input, &AcceptCertificateTransferOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AcceptCertificateTransferRequest{Request: req, Input: input, Copy: c.AcceptCertificateTransferRequest}
}

// AcceptCertificateTransferRequest is the request type for the
// AcceptCertificateTransfer API operation.
type AcceptCertificateTransferRequest struct {
	*aws.Request
	Input *AcceptCertificateTransferInput
	Copy  func(*AcceptCertificateTransferInput) AcceptCertificateTransferRequest
}

// Send marshals and sends the AcceptCertificateTransfer API request.
func (r AcceptCertificateTransferRequest) Send(ctx context.Context) (*AcceptCertificateTransferResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptCertificateTransferResponse{
		AcceptCertificateTransferOutput: r.Request.Data.(*AcceptCertificateTransferOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptCertificateTransferResponse is the response type for the
// AcceptCertificateTransfer API operation.
type AcceptCertificateTransferResponse struct {
	*AcceptCertificateTransferOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptCertificateTransfer request.
func (r *AcceptCertificateTransferResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
