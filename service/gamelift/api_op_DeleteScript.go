// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteScriptInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a Realtime script to delete. You can use either the
	// script ID or ARN value.
	//
	// ScriptId is a required field
	ScriptId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteScriptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteScriptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteScriptInput"}

	if s.ScriptId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScriptId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteScriptOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteScriptOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteScript = "DeleteScript"

// DeleteScriptRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Deletes a Realtime script. This action permanently deletes the script record.
// If script files were uploaded, they are also deleted (files stored in an
// S3 bucket are not deleted).
//
// To delete a script, specify the script ID. Before deleting a script, be sure
// to terminate all fleets that are deployed with the script being deleted.
// Fleet instances periodically check for script updates, and if the script
// record no longer exists, the instance will go into an error state and be
// unable to host game sessions.
//
// Learn more
//
// Amazon GameLift Realtime Servers (https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html)
//
// Related operations
//
//    * CreateScript
//
//    * ListScripts
//
//    * DescribeScript
//
//    * UpdateScript
//
//    * DeleteScript
//
//    // Example sending a request using DeleteScriptRequest.
//    req := client.DeleteScriptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteScript
func (c *Client) DeleteScriptRequest(input *DeleteScriptInput) DeleteScriptRequest {
	op := &aws.Operation{
		Name:       opDeleteScript,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteScriptInput{}
	}

	req := c.newRequest(op, input, &DeleteScriptOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteScriptRequest{Request: req, Input: input, Copy: c.DeleteScriptRequest}
}

// DeleteScriptRequest is the request type for the
// DeleteScript API operation.
type DeleteScriptRequest struct {
	*aws.Request
	Input *DeleteScriptInput
	Copy  func(*DeleteScriptInput) DeleteScriptRequest
}

// Send marshals and sends the DeleteScript API request.
func (r DeleteScriptRequest) Send(ctx context.Context) (*DeleteScriptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteScriptResponse{
		DeleteScriptOutput: r.Request.Data.(*DeleteScriptOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteScriptResponse is the response type for the
// DeleteScript API operation.
type DeleteScriptResponse struct {
	*DeleteScriptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteScript request.
func (r *DeleteScriptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
