// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CancelContactInput struct {
	_ struct{} `type:"structure"`

	// UUID of a contact.
	//
	// ContactId is a required field
	ContactId *string `location:"uri" locationName:"contactId" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelContactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelContactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelContactInput"}

	if s.ContactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContactId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelContactInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ContactId != nil {
		v := *s.ContactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "contactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelContactOutput struct {
	_ struct{} `type:"structure"`

	// UUID of a contact.
	ContactId *string `locationName:"contactId" type:"string"`
}

// String returns the string representation
func (s CancelContactOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelContactOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContactId != nil {
		v := *s.ContactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCancelContact = "CancelContact"

// CancelContactRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Cancels a contact with a specified contact ID.
//
//    // Example sending a request using CancelContactRequest.
//    req := client.CancelContactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/CancelContact
func (c *Client) CancelContactRequest(input *CancelContactInput) CancelContactRequest {
	op := &aws.Operation{
		Name:       opCancelContact,
		HTTPMethod: "DELETE",
		HTTPPath:   "/contact/{contactId}",
	}

	if input == nil {
		input = &CancelContactInput{}
	}

	req := c.newRequest(op, input, &CancelContactOutput{})

	return CancelContactRequest{Request: req, Input: input, Copy: c.CancelContactRequest}
}

// CancelContactRequest is the request type for the
// CancelContact API operation.
type CancelContactRequest struct {
	*aws.Request
	Input *CancelContactInput
	Copy  func(*CancelContactInput) CancelContactRequest
}

// Send marshals and sends the CancelContact API request.
func (r CancelContactRequest) Send(ctx context.Context) (*CancelContactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelContactResponse{
		CancelContactOutput: r.Request.Data.(*CancelContactOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelContactResponse is the response type for the
// CancelContact API operation.
type CancelContactResponse struct {
	*CancelContactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelContact request.
func (r *CancelContactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
