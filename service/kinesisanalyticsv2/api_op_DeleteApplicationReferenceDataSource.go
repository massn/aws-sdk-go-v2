// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationReferenceDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The current application version. You can use the DescribeApplication operation
	// to get the current application version. If the version specified is not the
	// current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The ID of the reference data source. When you add a reference data source
	// to your application using the AddApplicationReferenceDataSource, Kinesis
	// Data Analytics assigns an ID. You can use the DescribeApplication operation
	// to get the reference ID.
	//
	// ReferenceId is a required field
	ReferenceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationReferenceDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationReferenceDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationReferenceDataSourceInput"}

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

	if s.ReferenceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReferenceId"))
	}
	if s.ReferenceId != nil && len(*s.ReferenceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ReferenceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationReferenceDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string `min:"1" type:"string"`

	// The updated version ID of the application.
	ApplicationVersionId *int64 `min:"1" type:"long"`
}

// String returns the string representation
func (s DeleteApplicationReferenceDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApplicationReferenceDataSource = "DeleteApplicationReferenceDataSource"

// DeleteApplicationReferenceDataSourceRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Deletes a reference data source configuration from the specified SQL-based
// Amazon Kinesis Data Analytics application's configuration.
//
// If the application is running, Kinesis Data Analytics immediately removes
// the in-application table that you created using the AddApplicationReferenceDataSource
// operation.
//
//    // Example sending a request using DeleteApplicationReferenceDataSourceRequest.
//    req := client.DeleteApplicationReferenceDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/DeleteApplicationReferenceDataSource
func (c *Client) DeleteApplicationReferenceDataSourceRequest(input *DeleteApplicationReferenceDataSourceInput) DeleteApplicationReferenceDataSourceRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationReferenceDataSource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationReferenceDataSourceInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationReferenceDataSourceOutput{})

	return DeleteApplicationReferenceDataSourceRequest{Request: req, Input: input, Copy: c.DeleteApplicationReferenceDataSourceRequest}
}

// DeleteApplicationReferenceDataSourceRequest is the request type for the
// DeleteApplicationReferenceDataSource API operation.
type DeleteApplicationReferenceDataSourceRequest struct {
	*aws.Request
	Input *DeleteApplicationReferenceDataSourceInput
	Copy  func(*DeleteApplicationReferenceDataSourceInput) DeleteApplicationReferenceDataSourceRequest
}

// Send marshals and sends the DeleteApplicationReferenceDataSource API request.
func (r DeleteApplicationReferenceDataSourceRequest) Send(ctx context.Context) (*DeleteApplicationReferenceDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationReferenceDataSourceResponse{
		DeleteApplicationReferenceDataSourceOutput: r.Request.Data.(*DeleteApplicationReferenceDataSourceOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationReferenceDataSourceResponse is the response type for the
// DeleteApplicationReferenceDataSource API operation.
type DeleteApplicationReferenceDataSourceResponse struct {
	*DeleteApplicationReferenceDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationReferenceDataSource request.
func (r *DeleteApplicationReferenceDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
