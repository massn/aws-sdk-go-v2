// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteIndexInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the index to delete.
	//
	// Id is a required field
	Id *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIndexInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIndexInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIndexInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteIndexOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteIndexOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteIndex = "DeleteIndex"

// DeleteIndexRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Deletes an existing Amazon Kendra index. An exception is not thrown if the
// index is already being deleted. While the index is being deleted, the Status
// field returned by a call to the DescribeIndex operation is set to DELETING.
//
//    // Example sending a request using DeleteIndexRequest.
//    req := client.DeleteIndexRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/DeleteIndex
func (c *Client) DeleteIndexRequest(input *DeleteIndexInput) DeleteIndexRequest {
	op := &aws.Operation{
		Name:       opDeleteIndex,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteIndexInput{}
	}

	req := c.newRequest(op, input, &DeleteIndexOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteIndexRequest{Request: req, Input: input, Copy: c.DeleteIndexRequest}
}

// DeleteIndexRequest is the request type for the
// DeleteIndex API operation.
type DeleteIndexRequest struct {
	*aws.Request
	Input *DeleteIndexInput
	Copy  func(*DeleteIndexInput) DeleteIndexRequest
}

// Send marshals and sends the DeleteIndex API request.
func (r DeleteIndexRequest) Send(ctx context.Context) (*DeleteIndexResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIndexResponse{
		DeleteIndexOutput: r.Request.Data.(*DeleteIndexOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIndexResponse is the response type for the
// DeleteIndex API operation.
type DeleteIndexResponse struct {
	*DeleteIndexOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIndex request.
func (r *DeleteIndexResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
