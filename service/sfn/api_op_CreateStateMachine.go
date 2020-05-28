// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateStateMachineInput struct {
	_ struct{} `type:"structure"`

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	//
	// Definition is a required field
	Definition *string `locationName:"definition" min:"1" type:"string" required:"true" sensitive:"true"`

	// Defines what execution history events are logged and where they are logged.
	//
	// By default, the level is set to OFF. For more information see Log Levels
	// (https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html)
	// in the AWS Step Functions User Guide.
	LoggingConfiguration *LoggingConfiguration `locationName:"loggingConfiguration" type:"structure"`

	// The name of the state machine.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable logging with CloudWatch Logs, the name should only contain 0-9,
	// A-Z, a-z, - and _.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"1" type:"string" required:"true"`

	// Tags to be added when creating a state machine.
	//
	// An array of key-value pairs. For more information, see Using Cost Allocation
	// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
	// in the AWS Billing and Cost Management User Guide, and Controlling Access
	// Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).
	//
	// Tags may only contain Unicode letters, digits, white space, or these symbols:
	// _ . : / = + - @.
	Tags []Tag `locationName:"tags" type:"list"`

	// Determines whether a Standard or Express state machine is created. The default
	// is STANDARD. You cannot update the type of a state machine once it has been
	// created.
	Type StateMachineType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateStateMachineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStateMachineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStateMachineInput"}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}
	if s.Definition != nil && len(*s.Definition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Definition", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
	}
	if s.LoggingConfiguration != nil {
		if err := s.LoggingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LoggingConfiguration", err.(aws.ErrInvalidParams))
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

type CreateStateMachineOutput struct {
	_ struct{} `type:"structure"`

	// The date the state machine is created.
	//
	// CreationDate is a required field
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp" required:"true"`

	// The Amazon Resource Name (ARN) that identifies the created state machine.
	//
	// StateMachineArn is a required field
	StateMachineArn *string `locationName:"stateMachineArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateStateMachineOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStateMachine = "CreateStateMachine"

// CreateStateMachineRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Creates a state machine. A state machine consists of a collection of states
// that can do work (Task states), determine to which states to transition next
// (Choice states), stop an execution with an error (Fail states), and so on.
// State machines are specified using a JSON-based, structured language. For
// more information, see Amazon States Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html)
// in the AWS Step Functions User Guide.
//
// This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
//
// CreateStateMachine is an idempotent API. Subsequent requests won’t create
// a duplicate resource if it was already created. CreateStateMachine's idempotency
// check is based on the state machine name, definition, type, and LoggingConfiguration.
// If a following request has a different roleArn or tags, Step Functions will
// ignore these differences and treat it as an idempotent request of the previous.
// In this case, roleArn and tags will not be updated, even if they are different.
//
//    // Example sending a request using CreateStateMachineRequest.
//    req := client.CreateStateMachineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/CreateStateMachine
func (c *Client) CreateStateMachineRequest(input *CreateStateMachineInput) CreateStateMachineRequest {
	op := &aws.Operation{
		Name:       opCreateStateMachine,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStateMachineInput{}
	}

	req := c.newRequest(op, input, &CreateStateMachineOutput{})

	return CreateStateMachineRequest{Request: req, Input: input, Copy: c.CreateStateMachineRequest}
}

// CreateStateMachineRequest is the request type for the
// CreateStateMachine API operation.
type CreateStateMachineRequest struct {
	*aws.Request
	Input *CreateStateMachineInput
	Copy  func(*CreateStateMachineInput) CreateStateMachineRequest
}

// Send marshals and sends the CreateStateMachine API request.
func (r CreateStateMachineRequest) Send(ctx context.Context) (*CreateStateMachineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStateMachineResponse{
		CreateStateMachineOutput: r.Request.Data.(*CreateStateMachineOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStateMachineResponse is the response type for the
// CreateStateMachine API operation.
type CreateStateMachineResponse struct {
	*CreateStateMachineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStateMachine request.
func (r *CreateStateMachineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
