// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateApprovalRuleTemplateDescriptionInput struct {
	_ struct{} `type:"structure"`

	// The updated description of the approval rule template.
	//
	// ApprovalRuleTemplateDescription is a required field
	ApprovalRuleTemplateDescription *string `locationName:"approvalRuleTemplateDescription" type:"string" required:"true"`

	// The name of the template for which you want to update the description.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateDescriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApprovalRuleTemplateDescriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApprovalRuleTemplateDescriptionInput"}

	if s.ApprovalRuleTemplateDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateDescription"))
	}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateApprovalRuleTemplateDescriptionOutput struct {
	_ struct{} `type:"structure"`

	// The structure and content of the updated approval rule template.
	//
	// ApprovalRuleTemplate is a required field
	ApprovalRuleTemplate *ApprovalRuleTemplate `locationName:"approvalRuleTemplate" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateDescriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateApprovalRuleTemplateDescription = "UpdateApprovalRuleTemplateDescription"

// UpdateApprovalRuleTemplateDescriptionRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Updates the description for a specified approval rule template.
//
//    // Example sending a request using UpdateApprovalRuleTemplateDescriptionRequest.
//    req := client.UpdateApprovalRuleTemplateDescriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdateApprovalRuleTemplateDescription
func (c *Client) UpdateApprovalRuleTemplateDescriptionRequest(input *UpdateApprovalRuleTemplateDescriptionInput) UpdateApprovalRuleTemplateDescriptionRequest {
	op := &aws.Operation{
		Name:       opUpdateApprovalRuleTemplateDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateApprovalRuleTemplateDescriptionInput{}
	}

	req := c.newRequest(op, input, &UpdateApprovalRuleTemplateDescriptionOutput{})

	return UpdateApprovalRuleTemplateDescriptionRequest{Request: req, Input: input, Copy: c.UpdateApprovalRuleTemplateDescriptionRequest}
}

// UpdateApprovalRuleTemplateDescriptionRequest is the request type for the
// UpdateApprovalRuleTemplateDescription API operation.
type UpdateApprovalRuleTemplateDescriptionRequest struct {
	*aws.Request
	Input *UpdateApprovalRuleTemplateDescriptionInput
	Copy  func(*UpdateApprovalRuleTemplateDescriptionInput) UpdateApprovalRuleTemplateDescriptionRequest
}

// Send marshals and sends the UpdateApprovalRuleTemplateDescription API request.
func (r UpdateApprovalRuleTemplateDescriptionRequest) Send(ctx context.Context) (*UpdateApprovalRuleTemplateDescriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApprovalRuleTemplateDescriptionResponse{
		UpdateApprovalRuleTemplateDescriptionOutput: r.Request.Data.(*UpdateApprovalRuleTemplateDescriptionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApprovalRuleTemplateDescriptionResponse is the response type for the
// UpdateApprovalRuleTemplateDescription API operation.
type UpdateApprovalRuleTemplateDescriptionResponse struct {
	*UpdateApprovalRuleTemplateDescriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApprovalRuleTemplateDescription request.
func (r *UpdateApprovalRuleTemplateDescriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
