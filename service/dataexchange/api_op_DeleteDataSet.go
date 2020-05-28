// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteDataSetInput struct {
	_ struct{} `type:"structure"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDataSetInput"}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDataSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDataSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDataSet = "DeleteDataSet"

// DeleteDataSetRequest returns a request value for making API operation for
// AWS Data Exchange.
//
// This operation deletes a data set.
//
//    // Example sending a request using DeleteDataSetRequest.
//    req := client.DeleteDataSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dataexchange-2017-07-25/DeleteDataSet
func (c *Client) DeleteDataSetRequest(input *DeleteDataSetInput) DeleteDataSetRequest {
	op := &aws.Operation{
		Name:       opDeleteDataSet,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/data-sets/{DataSetId}",
	}

	if input == nil {
		input = &DeleteDataSetInput{}
	}

	req := c.newRequest(op, input, &DeleteDataSetOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteDataSetRequest{Request: req, Input: input, Copy: c.DeleteDataSetRequest}
}

// DeleteDataSetRequest is the request type for the
// DeleteDataSet API operation.
type DeleteDataSetRequest struct {
	*aws.Request
	Input *DeleteDataSetInput
	Copy  func(*DeleteDataSetInput) DeleteDataSetRequest
}

// Send marshals and sends the DeleteDataSet API request.
func (r DeleteDataSetRequest) Send(ctx context.Context) (*DeleteDataSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDataSetResponse{
		DeleteDataSetOutput: r.Request.Data.(*DeleteDataSetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDataSetResponse is the response type for the
// DeleteDataSet API operation.
type DeleteDataSetResponse struct {
	*DeleteDataSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDataSet request.
func (r *DeleteDataSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
