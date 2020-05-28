// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteLinkInput struct {
	_ struct{} `type:"structure"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`

	// The ID of the link.
	//
	// LinkId is a required field
	LinkId *string `location:"uri" locationName:"linkId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLinkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLinkInput"}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}

	if s.LinkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LinkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLinkInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LinkId != nil {
		v := *s.LinkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "linkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteLinkOutput struct {
	_ struct{} `type:"structure"`

	// Information about the link.
	Link *Link `type:"structure"`
}

// String returns the string representation
func (s DeleteLinkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteLinkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Link != nil {
		v := s.Link

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Link", v, metadata)
	}
	return nil
}

const opDeleteLink = "DeleteLink"

// DeleteLinkRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Deletes an existing link. You must first disassociate the link from any devices
// and customer gateways.
//
//    // Example sending a request using DeleteLinkRequest.
//    req := client.DeleteLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/DeleteLink
func (c *Client) DeleteLinkRequest(input *DeleteLinkInput) DeleteLinkRequest {
	op := &aws.Operation{
		Name:       opDeleteLink,
		HTTPMethod: "DELETE",
		HTTPPath:   "/global-networks/{globalNetworkId}/links/{linkId}",
	}

	if input == nil {
		input = &DeleteLinkInput{}
	}

	req := c.newRequest(op, input, &DeleteLinkOutput{})

	return DeleteLinkRequest{Request: req, Input: input, Copy: c.DeleteLinkRequest}
}

// DeleteLinkRequest is the request type for the
// DeleteLink API operation.
type DeleteLinkRequest struct {
	*aws.Request
	Input *DeleteLinkInput
	Copy  func(*DeleteLinkInput) DeleteLinkRequest
}

// Send marshals and sends the DeleteLink API request.
func (r DeleteLinkRequest) Send(ctx context.Context) (*DeleteLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLinkResponse{
		DeleteLinkOutput: r.Request.Data.(*DeleteLinkOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLinkResponse is the response type for the
// DeleteLink API operation.
type DeleteLinkResponse struct {
	*DeleteLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLink request.
func (r *DeleteLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
