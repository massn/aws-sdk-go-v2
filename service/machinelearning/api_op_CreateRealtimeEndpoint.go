// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRealtimeEndpointInput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the MLModel during creation.
	//
	// MLModelId is a required field
	MLModelId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRealtimeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRealtimeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRealtimeEndpointInput"}

	if s.MLModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MLModelId"))
	}
	if s.MLModelId != nil && len(*s.MLModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MLModelId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an CreateRealtimeEndpoint operation.
//
// The result contains the MLModelId and the endpoint information for the MLModel.
//
// The endpoint information includes the URI of the MLModel; that is, the location
// to send online prediction requests for the specified MLModel.
type CreateRealtimeEndpointOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the MLModel. This value should
	// be identical to the value of the MLModelId in the request.
	MLModelId *string `min:"1" type:"string"`

	// The endpoint information of the MLModel
	RealtimeEndpointInfo *RealtimeEndpointInfo `type:"structure"`
}

// String returns the string representation
func (s CreateRealtimeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRealtimeEndpoint = "CreateRealtimeEndpoint"

// CreateRealtimeEndpointRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Creates a real-time endpoint for the MLModel. The endpoint contains the URI
// of the MLModel; that is, the location to send real-time prediction requests
// for the specified MLModel.
//
//    // Example sending a request using CreateRealtimeEndpointRequest.
//    req := client.CreateRealtimeEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateRealtimeEndpointRequest(input *CreateRealtimeEndpointInput) CreateRealtimeEndpointRequest {
	op := &aws.Operation{
		Name:       opCreateRealtimeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRealtimeEndpointInput{}
	}

	req := c.newRequest(op, input, &CreateRealtimeEndpointOutput{})

	return CreateRealtimeEndpointRequest{Request: req, Input: input, Copy: c.CreateRealtimeEndpointRequest}
}

// CreateRealtimeEndpointRequest is the request type for the
// CreateRealtimeEndpoint API operation.
type CreateRealtimeEndpointRequest struct {
	*aws.Request
	Input *CreateRealtimeEndpointInput
	Copy  func(*CreateRealtimeEndpointInput) CreateRealtimeEndpointRequest
}

// Send marshals and sends the CreateRealtimeEndpoint API request.
func (r CreateRealtimeEndpointRequest) Send(ctx context.Context) (*CreateRealtimeEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRealtimeEndpointResponse{
		CreateRealtimeEndpointOutput: r.Request.Data.(*CreateRealtimeEndpointOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRealtimeEndpointResponse is the response type for the
// CreateRealtimeEndpoint API operation.
type CreateRealtimeEndpointResponse struct {
	*CreateRealtimeEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRealtimeEndpoint request.
func (r *CreateRealtimeEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
