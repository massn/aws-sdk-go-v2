// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResolveRoomInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the skill that was requested. Required.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`

	// The ARN of the user. Required.
	//
	// UserId is a required field
	UserId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResolveRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResolveRoomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResolveRoomInput"}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResolveRoomOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room from which the skill request was invoked.
	RoomArn *string `type:"string"`

	// The name of the room from which the skill request was invoked.
	RoomName *string `min:"1" type:"string"`

	// Response to get the room profile request. Required.
	RoomSkillParameters []RoomSkillParameter `type:"list"`
}

// String returns the string representation
func (s ResolveRoomOutput) String() string {
	return awsutil.Prettify(s)
}

const opResolveRoom = "ResolveRoom"

// ResolveRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Determines the details for the room from which a skill request was invoked.
// This operation is used by skill developers.
//
//    // Example sending a request using ResolveRoomRequest.
//    req := client.ResolveRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ResolveRoom
func (c *Client) ResolveRoomRequest(input *ResolveRoomInput) ResolveRoomRequest {
	op := &aws.Operation{
		Name:       opResolveRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResolveRoomInput{}
	}

	req := c.newRequest(op, input, &ResolveRoomOutput{})

	return ResolveRoomRequest{Request: req, Input: input, Copy: c.ResolveRoomRequest}
}

// ResolveRoomRequest is the request type for the
// ResolveRoom API operation.
type ResolveRoomRequest struct {
	*aws.Request
	Input *ResolveRoomInput
	Copy  func(*ResolveRoomInput) ResolveRoomRequest
}

// Send marshals and sends the ResolveRoom API request.
func (r ResolveRoomRequest) Send(ctx context.Context) (*ResolveRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResolveRoomResponse{
		ResolveRoomOutput: r.Request.Data.(*ResolveRoomOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResolveRoomResponse is the response type for the
// ResolveRoom API operation.
type ResolveRoomResponse struct {
	*ResolveRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResolveRoom request.
func (r *ResolveRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
