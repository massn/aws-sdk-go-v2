// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateInputSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	Tags map[string]string `locationName:"tags" type:"map"`

	WhitelistRules []InputWhitelistRuleCidr `locationName:"whitelistRules" type:"list"`
}

// String returns the string representation
func (s CreateInputSecurityGroupInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateInputSecurityGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.WhitelistRules != nil {
		v := s.WhitelistRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "whitelistRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type CreateInputSecurityGroupOutput struct {
	_ struct{} `type:"structure"`

	// An Input Security Group
	SecurityGroup *InputSecurityGroup `locationName:"securityGroup" type:"structure"`
}

// String returns the string representation
func (s CreateInputSecurityGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateInputSecurityGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SecurityGroup != nil {
		v := s.SecurityGroup

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "securityGroup", v, metadata)
	}
	return nil
}

const opCreateInputSecurityGroup = "CreateInputSecurityGroup"

// CreateInputSecurityGroupRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Creates a Input Security Group
//
//    // Example sending a request using CreateInputSecurityGroupRequest.
//    req := client.CreateInputSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/CreateInputSecurityGroup
func (c *Client) CreateInputSecurityGroupRequest(input *CreateInputSecurityGroupInput) CreateInputSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opCreateInputSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/prod/inputSecurityGroups",
	}

	if input == nil {
		input = &CreateInputSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &CreateInputSecurityGroupOutput{})

	return CreateInputSecurityGroupRequest{Request: req, Input: input, Copy: c.CreateInputSecurityGroupRequest}
}

// CreateInputSecurityGroupRequest is the request type for the
// CreateInputSecurityGroup API operation.
type CreateInputSecurityGroupRequest struct {
	*aws.Request
	Input *CreateInputSecurityGroupInput
	Copy  func(*CreateInputSecurityGroupInput) CreateInputSecurityGroupRequest
}

// Send marshals and sends the CreateInputSecurityGroup API request.
func (r CreateInputSecurityGroupRequest) Send(ctx context.Context) (*CreateInputSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInputSecurityGroupResponse{
		CreateInputSecurityGroupOutput: r.Request.Data.(*CreateInputSecurityGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInputSecurityGroupResponse is the response type for the
// CreateInputSecurityGroup API operation.
type CreateInputSecurityGroupResponse struct {
	*CreateInputSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInputSecurityGroup request.
func (r *CreateInputSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
