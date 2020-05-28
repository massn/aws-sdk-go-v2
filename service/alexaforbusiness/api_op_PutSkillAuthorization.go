// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutSkillAuthorizationInput struct {
	_ struct{} `type:"structure"`

	// The authorization result specific to OAUTH code grant output. "Code” must
	// be populated in the AuthorizationResult map to establish the authorization.
	//
	// AuthorizationResult is a required field
	AuthorizationResult map[string]string `type:"map" required:"true" sensitive:"true"`

	// The room that the skill is authorized for.
	RoomArn *string `type:"string"`

	// The unique identifier of a skill.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutSkillAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutSkillAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutSkillAuthorizationInput"}

	if s.AuthorizationResult == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizationResult"))
	}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutSkillAuthorizationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutSkillAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutSkillAuthorization = "PutSkillAuthorization"

// PutSkillAuthorizationRequest returns a request value for making API operation for
// Alexa For Business.
//
// Links a user's account to a third-party skill provider. If this API operation
// is called by an assumed IAM role, the skill being linked must be a private
// skill. Also, the skill must be owned by the AWS account that assumed the
// IAM role.
//
//    // Example sending a request using PutSkillAuthorizationRequest.
//    req := client.PutSkillAuthorizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/PutSkillAuthorization
func (c *Client) PutSkillAuthorizationRequest(input *PutSkillAuthorizationInput) PutSkillAuthorizationRequest {
	op := &aws.Operation{
		Name:       opPutSkillAuthorization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutSkillAuthorizationInput{}
	}

	req := c.newRequest(op, input, &PutSkillAuthorizationOutput{})

	return PutSkillAuthorizationRequest{Request: req, Input: input, Copy: c.PutSkillAuthorizationRequest}
}

// PutSkillAuthorizationRequest is the request type for the
// PutSkillAuthorization API operation.
type PutSkillAuthorizationRequest struct {
	*aws.Request
	Input *PutSkillAuthorizationInput
	Copy  func(*PutSkillAuthorizationInput) PutSkillAuthorizationRequest
}

// Send marshals and sends the PutSkillAuthorization API request.
func (r PutSkillAuthorizationRequest) Send(ctx context.Context) (*PutSkillAuthorizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutSkillAuthorizationResponse{
		PutSkillAuthorizationOutput: r.Request.Data.(*PutSkillAuthorizationOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutSkillAuthorizationResponse is the response type for the
// PutSkillAuthorization API operation.
type PutSkillAuthorizationResponse struct {
	*PutSkillAuthorizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutSkillAuthorization request.
func (r *PutSkillAuthorizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
