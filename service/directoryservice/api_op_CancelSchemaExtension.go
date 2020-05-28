// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CancelSchemaExtensionInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory whose schema extension will be canceled.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The identifier of the schema extension that will be canceled.
	//
	// SchemaExtensionId is a required field
	SchemaExtensionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CancelSchemaExtensionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelSchemaExtensionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelSchemaExtensionInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.SchemaExtensionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaExtensionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelSchemaExtensionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelSchemaExtensionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelSchemaExtension = "CancelSchemaExtension"

// CancelSchemaExtensionRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Cancels an in-progress schema extension to a Microsoft AD directory. Once
// a schema extension has started replicating to all domain controllers, the
// task can no longer be canceled. A schema extension can be canceled during
// any of the following states; Initializing, CreatingSnapshot, and UpdatingSchema.
//
//    // Example sending a request using CancelSchemaExtensionRequest.
//    req := client.CancelSchemaExtensionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CancelSchemaExtension
func (c *Client) CancelSchemaExtensionRequest(input *CancelSchemaExtensionInput) CancelSchemaExtensionRequest {
	op := &aws.Operation{
		Name:       opCancelSchemaExtension,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelSchemaExtensionInput{}
	}

	req := c.newRequest(op, input, &CancelSchemaExtensionOutput{})

	return CancelSchemaExtensionRequest{Request: req, Input: input, Copy: c.CancelSchemaExtensionRequest}
}

// CancelSchemaExtensionRequest is the request type for the
// CancelSchemaExtension API operation.
type CancelSchemaExtensionRequest struct {
	*aws.Request
	Input *CancelSchemaExtensionInput
	Copy  func(*CancelSchemaExtensionInput) CancelSchemaExtensionRequest
}

// Send marshals and sends the CancelSchemaExtension API request.
func (r CancelSchemaExtensionRequest) Send(ctx context.Context) (*CancelSchemaExtensionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelSchemaExtensionResponse{
		CancelSchemaExtensionOutput: r.Request.Data.(*CancelSchemaExtensionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelSchemaExtensionResponse is the response type for the
// CancelSchemaExtension API operation.
type CancelSchemaExtensionResponse struct {
	*CancelSchemaExtensionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelSchemaExtension request.
func (r *CancelSchemaExtensionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
