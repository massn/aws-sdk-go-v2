// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateGatewayInformationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that you
	// want to use to monitor and log events in the gateway.
	//
	// For more information, see What Is Amazon CloudWatch Logs? (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/WhatIsCloudWatchLogs.html).
	CloudWatchLogGroupARN *string `type:"string"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The name you configured for your gateway.
	GatewayName *string `min:"2" type:"string"`

	// A value that indicates the time zone of the gateway.
	GatewayTimezone *string `min:"3" type:"string"`
}

// String returns the string representation
func (s UpdateGatewayInformationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGatewayInformationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGatewayInformationInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.GatewayName != nil && len(*s.GatewayName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayName", 2))
	}
	if s.GatewayTimezone != nil && len(*s.GatewayTimezone) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayTimezone", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the ARN of the gateway that was updated.
type UpdateGatewayInformationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// The name you configured for your gateway.
	GatewayName *string `type:"string"`
}

// String returns the string representation
func (s UpdateGatewayInformationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateGatewayInformation = "UpdateGatewayInformation"

// UpdateGatewayInformationRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates a gateway's metadata, which includes the gateway's name and time
// zone. To specify which gateway to update, use the Amazon Resource Name (ARN)
// of the gateway in your request.
//
// For Gateways activated after September 2, 2015, the gateway's ARN contains
// the gateway ID rather than the gateway name. However, changing the name of
// the gateway has no effect on the gateway's ARN.
//
//    // Example sending a request using UpdateGatewayInformationRequest.
//    req := client.UpdateGatewayInformationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateGatewayInformation
func (c *Client) UpdateGatewayInformationRequest(input *UpdateGatewayInformationInput) UpdateGatewayInformationRequest {
	op := &aws.Operation{
		Name:       opUpdateGatewayInformation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateGatewayInformationInput{}
	}

	req := c.newRequest(op, input, &UpdateGatewayInformationOutput{})

	return UpdateGatewayInformationRequest{Request: req, Input: input, Copy: c.UpdateGatewayInformationRequest}
}

// UpdateGatewayInformationRequest is the request type for the
// UpdateGatewayInformation API operation.
type UpdateGatewayInformationRequest struct {
	*aws.Request
	Input *UpdateGatewayInformationInput
	Copy  func(*UpdateGatewayInformationInput) UpdateGatewayInformationRequest
}

// Send marshals and sends the UpdateGatewayInformation API request.
func (r UpdateGatewayInformationRequest) Send(ctx context.Context) (*UpdateGatewayInformationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGatewayInformationResponse{
		UpdateGatewayInformationOutput: r.Request.Data.(*UpdateGatewayInformationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGatewayInformationResponse is the response type for the
// UpdateGatewayInformation API operation.
type UpdateGatewayInformationResponse struct {
	*UpdateGatewayInformationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGatewayInformation request.
func (r *UpdateGatewayInformationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
