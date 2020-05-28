// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type BatchCreateRoomMembershipInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The list of membership items.
	//
	// MembershipItemList is a required field
	MembershipItemList []MembershipItem `type:"list" required:"true"`

	// The room ID.
	//
	// RoomId is a required field
	RoomId *string `location:"uri" locationName:"roomId" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchCreateRoomMembershipInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchCreateRoomMembershipInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchCreateRoomMembershipInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.MembershipItemList == nil {
		invalidParams.Add(aws.NewErrParamRequired("MembershipItemList"))
	}

	if s.RoomId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchCreateRoomMembershipInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MembershipItemList != nil {
		v := s.MembershipItemList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "MembershipItemList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoomId != nil {
		v := *s.RoomId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "roomId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type BatchCreateRoomMembershipOutput struct {
	_ struct{} `type:"structure"`

	// If the action fails for one or more of the member IDs in the request, a list
	// of the member IDs is returned, along with error codes and error messages.
	Errors []MemberError `type:"list"`
}

// String returns the string representation
func (s BatchCreateRoomMembershipOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchCreateRoomMembershipOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Errors != nil {
		v := s.Errors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Errors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchCreateRoomMembership = "BatchCreateRoomMembership"

// BatchCreateRoomMembershipRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds up to 50 members to a chat room in an Amazon Chime Enterprise account.
// Members can be either users or bots. The member role designates whether the
// member is a chat room administrator or a general chat room member.
//
//    // Example sending a request using BatchCreateRoomMembershipRequest.
//    req := client.BatchCreateRoomMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchCreateRoomMembership
func (c *Client) BatchCreateRoomMembershipRequest(input *BatchCreateRoomMembershipInput) BatchCreateRoomMembershipRequest {
	op := &aws.Operation{
		Name:       opBatchCreateRoomMembership,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/rooms/{roomId}/memberships?operation=batch-create",
	}

	if input == nil {
		input = &BatchCreateRoomMembershipInput{}
	}

	req := c.newRequest(op, input, &BatchCreateRoomMembershipOutput{})

	return BatchCreateRoomMembershipRequest{Request: req, Input: input, Copy: c.BatchCreateRoomMembershipRequest}
}

// BatchCreateRoomMembershipRequest is the request type for the
// BatchCreateRoomMembership API operation.
type BatchCreateRoomMembershipRequest struct {
	*aws.Request
	Input *BatchCreateRoomMembershipInput
	Copy  func(*BatchCreateRoomMembershipInput) BatchCreateRoomMembershipRequest
}

// Send marshals and sends the BatchCreateRoomMembership API request.
func (r BatchCreateRoomMembershipRequest) Send(ctx context.Context) (*BatchCreateRoomMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchCreateRoomMembershipResponse{
		BatchCreateRoomMembershipOutput: r.Request.Data.(*BatchCreateRoomMembershipOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchCreateRoomMembershipResponse is the response type for the
// BatchCreateRoomMembership API operation.
type BatchCreateRoomMembershipResponse struct {
	*BatchCreateRoomMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchCreateRoomMembership request.
func (r *BatchCreateRoomMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
