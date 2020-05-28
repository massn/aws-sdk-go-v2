// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateWorkteamInput struct {
	_ struct{} `type:"structure"`

	// A description of the work team.
	//
	// Description is a required field
	Description *string `min:"1" type:"string" required:"true"`

	// A list of MemberDefinition objects that contains objects that identify the
	// Amazon Cognito user pool that makes up the work team. For more information,
	// see Amazon Cognito User Pools (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools.html).
	//
	// All of the CognitoMemberDefinition objects that make up the member definition
	// must have the same ClientId and UserPool values.
	//
	// MemberDefinitions is a required field
	MemberDefinitions []MemberDefinition `min:"1" type:"list" required:"true"`

	// Configures notification of workers regarding available or expiring work items.
	NotificationConfiguration *NotificationConfiguration `type:"structure"`

	// An array of key-value pairs.
	//
	// For more information, see Resource Tag (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html)
	// and Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []Tag `type:"list"`

	// The name of the work team. Use this name to identify the work team.
	//
	// WorkteamName is a required field
	WorkteamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateWorkteamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWorkteamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWorkteamInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.MemberDefinitions == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberDefinitions"))
	}
	if s.MemberDefinitions != nil && len(s.MemberDefinitions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberDefinitions", 1))
	}

	if s.WorkteamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkteamName"))
	}
	if s.WorkteamName != nil && len(*s.WorkteamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkteamName", 1))
	}
	if s.MemberDefinitions != nil {
		for i, v := range s.MemberDefinitions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MemberDefinitions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateWorkteamOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the work team. You can use this ARN to
	// identify the work team.
	WorkteamArn *string `type:"string"`
}

// String returns the string representation
func (s CreateWorkteamOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateWorkteam = "CreateWorkteam"

// CreateWorkteamRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a new work team for labeling your data. A work team is defined by
// one or more Amazon Cognito user pools. You must first create the user pools
// before you can create a work team.
//
// You cannot create more than 25 work teams in an account and region.
//
//    // Example sending a request using CreateWorkteamRequest.
//    req := client.CreateWorkteamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateWorkteam
func (c *Client) CreateWorkteamRequest(input *CreateWorkteamInput) CreateWorkteamRequest {
	op := &aws.Operation{
		Name:       opCreateWorkteam,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWorkteamInput{}
	}

	req := c.newRequest(op, input, &CreateWorkteamOutput{})

	return CreateWorkteamRequest{Request: req, Input: input, Copy: c.CreateWorkteamRequest}
}

// CreateWorkteamRequest is the request type for the
// CreateWorkteam API operation.
type CreateWorkteamRequest struct {
	*aws.Request
	Input *CreateWorkteamInput
	Copy  func(*CreateWorkteamInput) CreateWorkteamRequest
}

// Send marshals and sends the CreateWorkteam API request.
func (r CreateWorkteamRequest) Send(ctx context.Context) (*CreateWorkteamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWorkteamResponse{
		CreateWorkteamOutput: r.Request.Data.(*CreateWorkteamOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWorkteamResponse is the response type for the
// CreateWorkteam API operation.
type CreateWorkteamResponse struct {
	*CreateWorkteamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWorkteam request.
func (r *CreateWorkteamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
