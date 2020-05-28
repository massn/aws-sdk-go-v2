// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteDatasetGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the dataset group to delete.
	//
	// DatasetGroupArn is a required field
	DatasetGroupArn *string `locationName:"datasetGroupArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDatasetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDatasetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDatasetGroupInput"}

	if s.DatasetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetGroupArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDatasetGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDatasetGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDatasetGroup = "DeleteDatasetGroup"

// DeleteDatasetGroupRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Deletes a dataset group. Before you delete a dataset group, you must delete
// the following:
//
//    * All associated event trackers.
//
//    * All associated solutions.
//
//    * All datasets in the dataset group.
//
//    // Example sending a request using DeleteDatasetGroupRequest.
//    req := client.DeleteDatasetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DeleteDatasetGroup
func (c *Client) DeleteDatasetGroupRequest(input *DeleteDatasetGroupInput) DeleteDatasetGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDatasetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDatasetGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDatasetGroupOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteDatasetGroupRequest{Request: req, Input: input, Copy: c.DeleteDatasetGroupRequest}
}

// DeleteDatasetGroupRequest is the request type for the
// DeleteDatasetGroup API operation.
type DeleteDatasetGroupRequest struct {
	*aws.Request
	Input *DeleteDatasetGroupInput
	Copy  func(*DeleteDatasetGroupInput) DeleteDatasetGroupRequest
}

// Send marshals and sends the DeleteDatasetGroup API request.
func (r DeleteDatasetGroupRequest) Send(ctx context.Context) (*DeleteDatasetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDatasetGroupResponse{
		DeleteDatasetGroupOutput: r.Request.Data.(*DeleteDatasetGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDatasetGroupResponse is the response type for the
// DeleteDatasetGroup API operation.
type DeleteDatasetGroupResponse struct {
	*DeleteDatasetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDatasetGroup request.
func (r *DeleteDatasetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
