// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutAppLaunchConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ID of the application associated with the launch configuration.
	AppId *string `locationName:"appId" type:"string"`

	// Name of service role in the customer's account that Amazon CloudFormation
	// uses to launch the application.
	RoleName *string `locationName:"roleName" type:"string"`

	// Launch configurations for server groups in the application.
	ServerGroupLaunchConfigurations []ServerGroupLaunchConfiguration `locationName:"serverGroupLaunchConfigurations" type:"list"`
}

// String returns the string representation
func (s PutAppLaunchConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

type PutAppLaunchConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutAppLaunchConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutAppLaunchConfiguration = "PutAppLaunchConfiguration"

// PutAppLaunchConfigurationRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Creates a launch configuration for an application.
//
//    // Example sending a request using PutAppLaunchConfigurationRequest.
//    req := client.PutAppLaunchConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/PutAppLaunchConfiguration
func (c *Client) PutAppLaunchConfigurationRequest(input *PutAppLaunchConfigurationInput) PutAppLaunchConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutAppLaunchConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutAppLaunchConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutAppLaunchConfigurationOutput{})

	return PutAppLaunchConfigurationRequest{Request: req, Input: input, Copy: c.PutAppLaunchConfigurationRequest}
}

// PutAppLaunchConfigurationRequest is the request type for the
// PutAppLaunchConfiguration API operation.
type PutAppLaunchConfigurationRequest struct {
	*aws.Request
	Input *PutAppLaunchConfigurationInput
	Copy  func(*PutAppLaunchConfigurationInput) PutAppLaunchConfigurationRequest
}

// Send marshals and sends the PutAppLaunchConfiguration API request.
func (r PutAppLaunchConfigurationRequest) Send(ctx context.Context) (*PutAppLaunchConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAppLaunchConfigurationResponse{
		PutAppLaunchConfigurationOutput: r.Request.Data.(*PutAppLaunchConfigurationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAppLaunchConfigurationResponse is the response type for the
// PutAppLaunchConfiguration API operation.
type PutAppLaunchConfigurationResponse struct {
	*PutAppLaunchConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAppLaunchConfiguration request.
func (r *PutAppLaunchConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
