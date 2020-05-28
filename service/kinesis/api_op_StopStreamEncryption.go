// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type StopStreamEncryptionInput struct {
	_ struct{} `type:"structure"`

	// The encryption type. The only valid value is KMS.
	//
	// EncryptionType is a required field
	EncryptionType EncryptionType `type:"string" required:"true" enum:"true"`

	// The GUID for the customer-managed AWS KMS key to use for encryption. This
	// value can be a globally unique identifier, a fully specified Amazon Resource
	// Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You
	// can also use a master key owned by Kinesis Data Streams by specifying the
	// alias aws/kinesis.
	//
	//    * Key ARN example: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//    * Alias ARN example: arn:aws:kms:us-east-1:123456789012:alias/MyAliasName
	//
	//    * Globally unique key ID example: 12345678-1234-1234-1234-123456789012
	//
	//    * Alias name example: alias/MyAliasName
	//
	//    * Master key owned by Kinesis Data Streams: alias/aws/kinesis
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// The name of the stream on which to stop encrypting records.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopStreamEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopStreamEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopStreamEncryptionInput"}
	if len(s.EncryptionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("EncryptionType"))
	}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopStreamEncryptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopStreamEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopStreamEncryption = "StopStreamEncryption"

// StopStreamEncryptionRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Disables server-side encryption for a specified stream.
//
// Stopping encryption is an asynchronous operation. Upon receiving the request,
// Kinesis Data Streams returns immediately and sets the status of the stream
// to UPDATING. After the update is complete, Kinesis Data Streams sets the
// status of the stream back to ACTIVE. Stopping encryption normally takes a
// few seconds to complete, but it can take minutes. You can continue to read
// and write data to your stream while its status is UPDATING. Once the status
// of the stream is ACTIVE, records written to the stream are no longer encrypted
// by Kinesis Data Streams.
//
// API Limits: You can successfully disable server-side encryption 25 times
// in a rolling 24-hour period.
//
// Note: It can take up to 5 seconds after the stream is in an ACTIVE status
// before all records written to the stream are no longer subject to encryption.
// After you disabled encryption, you can verify that encryption is not applied
// by inspecting the API response from PutRecord or PutRecords.
//
//    // Example sending a request using StopStreamEncryptionRequest.
//    req := client.StopStreamEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/StopStreamEncryption
func (c *Client) StopStreamEncryptionRequest(input *StopStreamEncryptionInput) StopStreamEncryptionRequest {
	op := &aws.Operation{
		Name:       opStopStreamEncryption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopStreamEncryptionInput{}
	}

	req := c.newRequest(op, input, &StopStreamEncryptionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return StopStreamEncryptionRequest{Request: req, Input: input, Copy: c.StopStreamEncryptionRequest}
}

// StopStreamEncryptionRequest is the request type for the
// StopStreamEncryption API operation.
type StopStreamEncryptionRequest struct {
	*aws.Request
	Input *StopStreamEncryptionInput
	Copy  func(*StopStreamEncryptionInput) StopStreamEncryptionRequest
}

// Send marshals and sends the StopStreamEncryption API request.
func (r StopStreamEncryptionRequest) Send(ctx context.Context) (*StopStreamEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopStreamEncryptionResponse{
		StopStreamEncryptionOutput: r.Request.Data.(*StopStreamEncryptionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopStreamEncryptionResponse is the response type for the
// StopStreamEncryption API operation.
type StopStreamEncryptionResponse struct {
	*StopStreamEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopStreamEncryption request.
func (r *StopStreamEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
