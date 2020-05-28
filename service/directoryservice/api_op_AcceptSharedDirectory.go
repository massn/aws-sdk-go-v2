// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AcceptSharedDirectoryInput struct {
	_ struct{} `type:"structure"`

	// Identifier of the shared directory in the directory consumer account. This
	// identifier is different for each directory owner account.
	//
	// SharedDirectoryId is a required field
	SharedDirectoryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptSharedDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptSharedDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptSharedDirectoryInput"}

	if s.SharedDirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SharedDirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AcceptSharedDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// The shared directory in the directory consumer account.
	SharedDirectory *SharedDirectory `type:"structure"`
}

// String returns the string representation
func (s AcceptSharedDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptSharedDirectory = "AcceptSharedDirectory"

// AcceptSharedDirectoryRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Accepts a directory sharing request that was sent from the directory owner
// account.
//
//    // Example sending a request using AcceptSharedDirectoryRequest.
//    req := client.AcceptSharedDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/AcceptSharedDirectory
func (c *Client) AcceptSharedDirectoryRequest(input *AcceptSharedDirectoryInput) AcceptSharedDirectoryRequest {
	op := &aws.Operation{
		Name:       opAcceptSharedDirectory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptSharedDirectoryInput{}
	}

	req := c.newRequest(op, input, &AcceptSharedDirectoryOutput{})

	return AcceptSharedDirectoryRequest{Request: req, Input: input, Copy: c.AcceptSharedDirectoryRequest}
}

// AcceptSharedDirectoryRequest is the request type for the
// AcceptSharedDirectory API operation.
type AcceptSharedDirectoryRequest struct {
	*aws.Request
	Input *AcceptSharedDirectoryInput
	Copy  func(*AcceptSharedDirectoryInput) AcceptSharedDirectoryRequest
}

// Send marshals and sends the AcceptSharedDirectory API request.
func (r AcceptSharedDirectoryRequest) Send(ctx context.Context) (*AcceptSharedDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptSharedDirectoryResponse{
		AcceptSharedDirectoryOutput: r.Request.Data.(*AcceptSharedDirectoryOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptSharedDirectoryResponse is the response type for the
// AcceptSharedDirectory API operation.
type AcceptSharedDirectoryResponse struct {
	*AcceptSharedDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptSharedDirectory request.
func (r *AcceptSharedDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
