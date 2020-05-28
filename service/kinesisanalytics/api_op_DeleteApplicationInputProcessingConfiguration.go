// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationInputProcessingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Kinesis Analytics application name.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The version ID of the Kinesis Analytics application.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The ID of the input configuration from which to delete the input processing
	// configuration. You can get a list of the input IDs for an application by
	// using the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation.
	//
	// InputId is a required field
	InputId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationInputProcessingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationInputProcessingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationInputProcessingConfigurationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.InputId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputId"))
	}
	if s.InputId != nil && len(*s.InputId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationInputProcessingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationInputProcessingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApplicationInputProcessingConfiguration = "DeleteApplicationInputProcessingConfiguration"

// DeleteApplicationInputProcessingConfigurationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
//
// This documentation is for version 1 of the Amazon Kinesis Data Analytics
// API, which only supports SQL applications. Version 2 of the API supports
// SQL and Java applications. For more information about version 2, see Amazon
// Kinesis Data Analytics API V2 Documentation (/kinesisanalytics/latest/apiv2/Welcome.html).
//
// Deletes an InputProcessingConfiguration (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html)
// from an input.
//
//    // Example sending a request using DeleteApplicationInputProcessingConfigurationRequest.
//    req := client.DeleteApplicationInputProcessingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationInputProcessingConfiguration
func (c *Client) DeleteApplicationInputProcessingConfigurationRequest(input *DeleteApplicationInputProcessingConfigurationInput) DeleteApplicationInputProcessingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationInputProcessingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationInputProcessingConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationInputProcessingConfigurationOutput{})

	return DeleteApplicationInputProcessingConfigurationRequest{Request: req, Input: input, Copy: c.DeleteApplicationInputProcessingConfigurationRequest}
}

// DeleteApplicationInputProcessingConfigurationRequest is the request type for the
// DeleteApplicationInputProcessingConfiguration API operation.
type DeleteApplicationInputProcessingConfigurationRequest struct {
	*aws.Request
	Input *DeleteApplicationInputProcessingConfigurationInput
	Copy  func(*DeleteApplicationInputProcessingConfigurationInput) DeleteApplicationInputProcessingConfigurationRequest
}

// Send marshals and sends the DeleteApplicationInputProcessingConfiguration API request.
func (r DeleteApplicationInputProcessingConfigurationRequest) Send(ctx context.Context) (*DeleteApplicationInputProcessingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationInputProcessingConfigurationResponse{
		DeleteApplicationInputProcessingConfigurationOutput: r.Request.Data.(*DeleteApplicationInputProcessingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationInputProcessingConfigurationResponse is the response type for the
// DeleteApplicationInputProcessingConfiguration API operation.
type DeleteApplicationInputProcessingConfigurationResponse struct {
	*DeleteApplicationInputProcessingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationInputProcessingConfiguration request.
func (r *DeleteApplicationInputProcessingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
