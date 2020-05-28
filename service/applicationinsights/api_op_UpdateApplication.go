// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateApplicationInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether Application Insights can listen to CloudWatch events for
	// the application resources, such as instance terminated, failed deployment,
	// and others.
	CWEMonitorEnabled *bool `type:"boolean"`

	// When set to true, creates opsItems for any problems detected on an application.
	OpsCenterEnabled *bool `type:"boolean"`

	// The SNS topic provided to Application Insights that is associated to the
	// created opsItem. Allows you to receive notifications for updates to the opsItem.
	OpsItemSNSTopicArn *string `min:"20" type:"string"`

	// Disassociates the SNS topic from the opsItem created for detected problems.
	RemoveSNSTopic *bool `type:"boolean"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApplicationInput"}
	if s.OpsItemSNSTopicArn != nil && len(*s.OpsItemSNSTopicArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("OpsItemSNSTopicArn", 20))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}
	if s.ResourceGroupName != nil && len(*s.ResourceGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the application.
	ApplicationInfo *ApplicationInfo `type:"structure"`
}

// String returns the string representation
func (s UpdateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateApplication = "UpdateApplication"

// UpdateApplicationRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Updates the application.
//
//    // Example sending a request using UpdateApplicationRequest.
//    req := client.UpdateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateApplication
func (c *Client) UpdateApplicationRequest(input *UpdateApplicationInput) UpdateApplicationRequest {
	op := &aws.Operation{
		Name:       opUpdateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateApplicationInput{}
	}

	req := c.newRequest(op, input, &UpdateApplicationOutput{})

	return UpdateApplicationRequest{Request: req, Input: input, Copy: c.UpdateApplicationRequest}
}

// UpdateApplicationRequest is the request type for the
// UpdateApplication API operation.
type UpdateApplicationRequest struct {
	*aws.Request
	Input *UpdateApplicationInput
	Copy  func(*UpdateApplicationInput) UpdateApplicationRequest
}

// Send marshals and sends the UpdateApplication API request.
func (r UpdateApplicationRequest) Send(ctx context.Context) (*UpdateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApplicationResponse{
		UpdateApplicationOutput: r.Request.Data.(*UpdateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApplicationResponse is the response type for the
// UpdateApplication API operation.
type UpdateApplicationResponse struct {
	*UpdateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApplication request.
func (r *UpdateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
