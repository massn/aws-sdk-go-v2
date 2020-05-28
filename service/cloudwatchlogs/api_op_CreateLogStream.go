// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type CreateLogStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The name of the log stream.
	//
	// LogStreamName is a required field
	LogStreamName *string `locationName:"logStreamName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLogStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLogStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLogStreamInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.LogStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogStreamName"))
	}
	if s.LogStreamName != nil && len(*s.LogStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogStreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLogStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLogStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLogStream = "CreateLogStream"

// CreateLogStreamRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Creates a log stream for the specified log group.
//
// There is no limit on the number of log streams that you can create for a
// log group. There is a limit of 50 TPS on CreateLogStream operations, after
// which transactions are throttled.
//
// You must use the following guidelines when naming a log stream:
//
//    * Log stream names must be unique within the log group.
//
//    * Log stream names can be between 1 and 512 characters long.
//
//    * The ':' (colon) and '*' (asterisk) characters are not allowed.
//
//    // Example sending a request using CreateLogStreamRequest.
//    req := client.CreateLogStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/CreateLogStream
func (c *Client) CreateLogStreamRequest(input *CreateLogStreamInput) CreateLogStreamRequest {
	op := &aws.Operation{
		Name:       opCreateLogStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLogStreamInput{}
	}

	req := c.newRequest(op, input, &CreateLogStreamOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return CreateLogStreamRequest{Request: req, Input: input, Copy: c.CreateLogStreamRequest}
}

// CreateLogStreamRequest is the request type for the
// CreateLogStream API operation.
type CreateLogStreamRequest struct {
	*aws.Request
	Input *CreateLogStreamInput
	Copy  func(*CreateLogStreamInput) CreateLogStreamRequest
}

// Send marshals and sends the CreateLogStream API request.
func (r CreateLogStreamRequest) Send(ctx context.Context) (*CreateLogStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLogStreamResponse{
		CreateLogStreamOutput: r.Request.Data.(*CreateLogStreamOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLogStreamResponse is the response type for the
// CreateLogStream API operation.
type CreateLogStreamResponse struct {
	*CreateLogStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLogStream request.
func (r *CreateLogStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
