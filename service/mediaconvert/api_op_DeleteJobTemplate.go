// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Delete a job template by sending a request with the job template name
type DeleteJobTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the job template to be deleted.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteJobTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteJobTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteJobTemplateInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteJobTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Delete job template requests will return an OK message or error message with
// an empty body.
type DeleteJobTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteJobTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteJobTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteJobTemplate = "DeleteJobTemplate"

// DeleteJobTemplateRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Permanently delete a job template you have created.
//
//    // Example sending a request using DeleteJobTemplateRequest.
//    req := client.DeleteJobTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/DeleteJobTemplate
func (c *Client) DeleteJobTemplateRequest(input *DeleteJobTemplateInput) DeleteJobTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteJobTemplate,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2017-08-29/jobTemplates/{name}",
	}

	if input == nil {
		input = &DeleteJobTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteJobTemplateOutput{})

	return DeleteJobTemplateRequest{Request: req, Input: input, Copy: c.DeleteJobTemplateRequest}
}

// DeleteJobTemplateRequest is the request type for the
// DeleteJobTemplate API operation.
type DeleteJobTemplateRequest struct {
	*aws.Request
	Input *DeleteJobTemplateInput
	Copy  func(*DeleteJobTemplateInput) DeleteJobTemplateRequest
}

// Send marshals and sends the DeleteJobTemplate API request.
func (r DeleteJobTemplateRequest) Send(ctx context.Context) (*DeleteJobTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteJobTemplateResponse{
		DeleteJobTemplateOutput: r.Request.Data.(*DeleteJobTemplateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteJobTemplateResponse is the response type for the
// DeleteJobTemplate API operation.
type DeleteJobTemplateResponse struct {
	*DeleteJobTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteJobTemplate request.
func (r *DeleteJobTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
