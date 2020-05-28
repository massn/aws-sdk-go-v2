// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetDeploymentGroup operation.
type GetDeploymentGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of an AWS CodeDeploy application associated with the IAM user or
	// AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// The name of a deployment group for the specified application.
	//
	// DeploymentGroupName is a required field
	DeploymentGroupName *string `locationName:"deploymentGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentGroupInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.DeploymentGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentGroupName"))
	}
	if s.DeploymentGroupName != nil && len(*s.DeploymentGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetDeploymentGroup operation.
type GetDeploymentGroupOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deployment group.
	DeploymentGroupInfo *DeploymentGroupInfo `locationName:"deploymentGroupInfo" type:"structure"`
}

// String returns the string representation
func (s GetDeploymentGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDeploymentGroup = "GetDeploymentGroup"

// GetDeploymentGroupRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about a deployment group.
//
//    // Example sending a request using GetDeploymentGroupRequest.
//    req := client.GetDeploymentGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/GetDeploymentGroup
func (c *Client) GetDeploymentGroupRequest(input *GetDeploymentGroupInput) GetDeploymentGroupRequest {
	op := &aws.Operation{
		Name:       opGetDeploymentGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeploymentGroupInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentGroupOutput{})

	return GetDeploymentGroupRequest{Request: req, Input: input, Copy: c.GetDeploymentGroupRequest}
}

// GetDeploymentGroupRequest is the request type for the
// GetDeploymentGroup API operation.
type GetDeploymentGroupRequest struct {
	*aws.Request
	Input *GetDeploymentGroupInput
	Copy  func(*GetDeploymentGroupInput) GetDeploymentGroupRequest
}

// Send marshals and sends the GetDeploymentGroup API request.
func (r GetDeploymentGroupRequest) Send(ctx context.Context) (*GetDeploymentGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentGroupResponse{
		GetDeploymentGroupOutput: r.Request.Data.(*GetDeploymentGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentGroupResponse is the response type for the
// GetDeploymentGroup API operation.
type GetDeploymentGroupResponse struct {
	*GetDeploymentGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeploymentGroup request.
func (r *GetDeploymentGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
