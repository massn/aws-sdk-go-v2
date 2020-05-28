// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutRoomSkillParameterInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room associated with the room skill parameter. Required.
	RoomArn *string `type:"string"`

	// The updated room skill parameter. Required.
	//
	// RoomSkillParameter is a required field
	RoomSkillParameter *RoomSkillParameter `type:"structure" required:"true"`

	// The ARN of the skill associated with the room skill parameter. Required.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutRoomSkillParameterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRoomSkillParameterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRoomSkillParameterInput"}

	if s.RoomSkillParameter == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomSkillParameter"))
	}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}
	if s.RoomSkillParameter != nil {
		if err := s.RoomSkillParameter.Validate(); err != nil {
			invalidParams.AddNested("RoomSkillParameter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutRoomSkillParameterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutRoomSkillParameterOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutRoomSkillParameter = "PutRoomSkillParameter"

// PutRoomSkillParameterRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates room skill parameter details by room, skill, and parameter key ID.
// Not all skills have a room skill parameter.
//
//    // Example sending a request using PutRoomSkillParameterRequest.
//    req := client.PutRoomSkillParameterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/PutRoomSkillParameter
func (c *Client) PutRoomSkillParameterRequest(input *PutRoomSkillParameterInput) PutRoomSkillParameterRequest {
	op := &aws.Operation{
		Name:       opPutRoomSkillParameter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutRoomSkillParameterInput{}
	}

	req := c.newRequest(op, input, &PutRoomSkillParameterOutput{})

	return PutRoomSkillParameterRequest{Request: req, Input: input, Copy: c.PutRoomSkillParameterRequest}
}

// PutRoomSkillParameterRequest is the request type for the
// PutRoomSkillParameter API operation.
type PutRoomSkillParameterRequest struct {
	*aws.Request
	Input *PutRoomSkillParameterInput
	Copy  func(*PutRoomSkillParameterInput) PutRoomSkillParameterRequest
}

// Send marshals and sends the PutRoomSkillParameter API request.
func (r PutRoomSkillParameterRequest) Send(ctx context.Context) (*PutRoomSkillParameterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRoomSkillParameterResponse{
		PutRoomSkillParameterOutput: r.Request.Data.(*PutRoomSkillParameterOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRoomSkillParameterResponse is the response type for the
// PutRoomSkillParameter API operation.
type PutRoomSkillParameterResponse struct {
	*PutRoomSkillParameterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRoomSkillParameter request.
func (r *PutRoomSkillParameterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
