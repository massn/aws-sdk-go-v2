// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMemberInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time.
	// This identifier is required only if you make a service request directly using
	// an HTTP client. It is generated automatically if you use an AWS SDK or the
	// AWS CLI.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The unique identifier of the invitation that is sent to the member to join
	// the network.
	//
	// InvitationId is a required field
	InvitationId *string `min:"1" type:"string" required:"true"`

	// Member configuration parameters.
	//
	// MemberConfiguration is a required field
	MemberConfiguration *MemberConfiguration `type:"structure" required:"true"`

	// The unique identifier of the network in which the member is created.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMemberInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.InvitationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InvitationId"))
	}
	if s.InvitationId != nil && len(*s.InvitationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InvitationId", 1))
	}

	if s.MemberConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberConfiguration"))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}
	if s.MemberConfiguration != nil {
		if err := s.MemberConfiguration.Validate(); err != nil {
			invalidParams.AddNested("MemberConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMemberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InvitationId != nil {
		v := *s.InvitationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InvitationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemberConfiguration != nil {
		v := s.MemberConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "MemberConfiguration", v, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateMemberOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the member.
	MemberId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateMemberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMemberOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateMember = "CreateMember"

// CreateMemberRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Creates a member within a Managed Blockchain network.
//
//    // Example sending a request using CreateMemberRequest.
//    req := client.CreateMemberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/CreateMember
func (c *Client) CreateMemberRequest(input *CreateMemberInput) CreateMemberRequest {
	op := &aws.Operation{
		Name:       opCreateMember,
		HTTPMethod: "POST",
		HTTPPath:   "/networks/{networkId}/members",
	}

	if input == nil {
		input = &CreateMemberInput{}
	}

	req := c.newRequest(op, input, &CreateMemberOutput{})

	return CreateMemberRequest{Request: req, Input: input, Copy: c.CreateMemberRequest}
}

// CreateMemberRequest is the request type for the
// CreateMember API operation.
type CreateMemberRequest struct {
	*aws.Request
	Input *CreateMemberInput
	Copy  func(*CreateMemberInput) CreateMemberRequest
}

// Send marshals and sends the CreateMember API request.
func (r CreateMemberRequest) Send(ctx context.Context) (*CreateMemberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMemberResponse{
		CreateMemberOutput: r.Request.Data.(*CreateMemberOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMemberResponse is the response type for the
// CreateMember API operation.
type CreateMemberResponse struct {
	*CreateMemberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMember request.
func (r *CreateMemberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
