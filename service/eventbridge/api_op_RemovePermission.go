// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type RemovePermissionInput struct {
	_ struct{} `type:"structure"`

	// The name of the event bus to revoke permissions for. If you omit this, the
	// default event bus is used.
	EventBusName *string `min:"1" type:"string"`

	// The statement ID corresponding to the account that is no longer allowed to
	// put events to the default event bus.
	//
	// StatementId is a required field
	StatementId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RemovePermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemovePermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemovePermissionInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}

	if s.StatementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatementId"))
	}
	if s.StatementId != nil && len(*s.StatementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatementId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemovePermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemovePermissionOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemovePermission = "RemovePermission"

// RemovePermissionRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Revokes the permission of another AWS account to be able to put events to
// the specified event bus. Specify the account to revoke by the StatementId
// value that you associated with the account when you granted it permission
// with PutPermission. You can find the StatementId by using DescribeEventBus.
//
//    // Example sending a request using RemovePermissionRequest.
//    req := client.RemovePermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/RemovePermission
func (c *Client) RemovePermissionRequest(input *RemovePermissionInput) RemovePermissionRequest {
	op := &aws.Operation{
		Name:       opRemovePermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemovePermissionInput{}
	}

	req := c.newRequest(op, input, &RemovePermissionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RemovePermissionRequest{Request: req, Input: input, Copy: c.RemovePermissionRequest}
}

// RemovePermissionRequest is the request type for the
// RemovePermission API operation.
type RemovePermissionRequest struct {
	*aws.Request
	Input *RemovePermissionInput
	Copy  func(*RemovePermissionInput) RemovePermissionRequest
}

// Send marshals and sends the RemovePermission API request.
func (r RemovePermissionRequest) Send(ctx context.Context) (*RemovePermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemovePermissionResponse{
		RemovePermissionOutput: r.Request.Data.(*RemovePermissionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemovePermissionResponse is the response type for the
// RemovePermission API operation.
type RemovePermissionResponse struct {
	*RemovePermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemovePermission request.
func (r *RemovePermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
