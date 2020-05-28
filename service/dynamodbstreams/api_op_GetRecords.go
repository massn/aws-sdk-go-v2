// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetRecords operation.
type GetRecordsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of records to return from the shard. The upper limit is
	// 1000.
	Limit *int64 `min:"1" type:"integer"`

	// A shard iterator that was retrieved from a previous GetShardIterator operation.
	// This iterator can be used to access the stream records in this shard.
	//
	// ShardIterator is a required field
	ShardIterator *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRecordsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRecordsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRecordsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.ShardIterator == nil {
		invalidParams.Add(aws.NewErrParamRequired("ShardIterator"))
	}
	if s.ShardIterator != nil && len(*s.ShardIterator) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ShardIterator", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetRecords operation.
type GetRecordsOutput struct {
	_ struct{} `type:"structure"`

	// The next position in the shard from which to start sequentially reading stream
	// records. If set to null, the shard has been closed and the requested iterator
	// will not return any more data.
	NextShardIterator *string `min:"1" type:"string"`

	// The stream records from the shard, which were retrieved using the shard iterator.
	Records []Record `type:"list"`
}

// String returns the string representation
func (s GetRecordsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRecords = "GetRecords"

// GetRecordsRequest returns a request value for making API operation for
// Amazon DynamoDB Streams.
//
// Retrieves the stream records from a given shard.
//
// Specify a shard iterator using the ShardIterator parameter. The shard iterator
// specifies the position in the shard from which you want to start reading
// stream records sequentially. If there are no stream records available in
// the portion of the shard that the iterator points to, GetRecords returns
// an empty list. Note that it might take multiple calls to get to a portion
// of the shard that contains stream records.
//
// GetRecords can retrieve a maximum of 1 MB of data or 1000 stream records,
// whichever comes first.
//
//    // Example sending a request using GetRecordsRequest.
//    req := client.GetRecordsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/GetRecords
func (c *Client) GetRecordsRequest(input *GetRecordsInput) GetRecordsRequest {
	op := &aws.Operation{
		Name:       opGetRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRecordsInput{}
	}

	req := c.newRequest(op, input, &GetRecordsOutput{})

	return GetRecordsRequest{Request: req, Input: input, Copy: c.GetRecordsRequest}
}

// GetRecordsRequest is the request type for the
// GetRecords API operation.
type GetRecordsRequest struct {
	*aws.Request
	Input *GetRecordsInput
	Copy  func(*GetRecordsInput) GetRecordsRequest
}

// Send marshals and sends the GetRecords API request.
func (r GetRecordsRequest) Send(ctx context.Context) (*GetRecordsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRecordsResponse{
		GetRecordsOutput: r.Request.Data.(*GetRecordsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRecordsResponse is the response type for the
// GetRecords API operation.
type GetRecordsResponse struct {
	*GetRecordsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRecords request.
func (r *GetRecordsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
