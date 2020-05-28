// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to move a dedicated IP address to a dedicated IP pool.
type PutDedicatedIpInPoolInput struct {
	_ struct{} `type:"structure"`

	// The name of the IP pool that you want to add the dedicated IP address to.
	// You have to specify an IP pool that already exists.
	//
	// DestinationPoolName is a required field
	DestinationPoolName *string `type:"string" required:"true"`

	// The IP address that you want to move to the dedicated IP pool. The value
	// you specify has to be a dedicated IP address that's associated with your
	// Amazon Pinpoint account.
	//
	// Ip is a required field
	Ip *string `location:"uri" locationName:"IP" type:"string" required:"true"`
}

// String returns the string representation
func (s PutDedicatedIpInPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDedicatedIpInPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDedicatedIpInPoolInput"}

	if s.DestinationPoolName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationPoolName"))
	}

	if s.Ip == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ip"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutDedicatedIpInPoolInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DestinationPoolName != nil {
		v := *s.DestinationPoolName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DestinationPoolName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Ip != nil {
		v := *s.Ip

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IP", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutDedicatedIpInPoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDedicatedIpInPoolOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutDedicatedIpInPoolOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutDedicatedIpInPool = "PutDedicatedIpInPool"

// PutDedicatedIpInPoolRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Move a dedicated IP address to an existing dedicated IP pool.
//
// The dedicated IP address that you specify must already exist, and must be
// associated with your Amazon Pinpoint account.
//
// The dedicated IP pool you specify must already exist. You can create a new
// pool by using the CreateDedicatedIpPool operation.
//
//    // Example sending a request using PutDedicatedIpInPoolRequest.
//    req := client.PutDedicatedIpInPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutDedicatedIpInPool
func (c *Client) PutDedicatedIpInPoolRequest(input *PutDedicatedIpInPoolInput) PutDedicatedIpInPoolRequest {
	op := &aws.Operation{
		Name:       opPutDedicatedIpInPool,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/dedicated-ips/{IP}/pool",
	}

	if input == nil {
		input = &PutDedicatedIpInPoolInput{}
	}

	req := c.newRequest(op, input, &PutDedicatedIpInPoolOutput{})

	return PutDedicatedIpInPoolRequest{Request: req, Input: input, Copy: c.PutDedicatedIpInPoolRequest}
}

// PutDedicatedIpInPoolRequest is the request type for the
// PutDedicatedIpInPool API operation.
type PutDedicatedIpInPoolRequest struct {
	*aws.Request
	Input *PutDedicatedIpInPoolInput
	Copy  func(*PutDedicatedIpInPoolInput) PutDedicatedIpInPoolRequest
}

// Send marshals and sends the PutDedicatedIpInPool API request.
func (r PutDedicatedIpInPoolRequest) Send(ctx context.Context) (*PutDedicatedIpInPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDedicatedIpInPoolResponse{
		PutDedicatedIpInPoolOutput: r.Request.Data.(*PutDedicatedIpInPoolOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDedicatedIpInPoolResponse is the response type for the
// PutDedicatedIpInPool API operation.
type PutDedicatedIpInPoolResponse struct {
	*PutDedicatedIpInPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDedicatedIpInPool request.
func (r *PutDedicatedIpInPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
