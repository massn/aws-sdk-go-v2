// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the GetDirectoryLimits operation.
type GetDirectoryLimitsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetDirectoryLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the results of the GetDirectoryLimits operation.
type GetDirectoryLimitsOutput struct {
	_ struct{} `type:"structure"`

	// A DirectoryLimits object that contains the directory limits for the current
	// rRegion.
	DirectoryLimits *DirectoryLimits `type:"structure"`
}

// String returns the string representation
func (s GetDirectoryLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDirectoryLimits = "GetDirectoryLimits"

// GetDirectoryLimitsRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Obtains directory limit information for the current Region.
//
//    // Example sending a request using GetDirectoryLimitsRequest.
//    req := client.GetDirectoryLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/GetDirectoryLimits
func (c *Client) GetDirectoryLimitsRequest(input *GetDirectoryLimitsInput) GetDirectoryLimitsRequest {
	op := &aws.Operation{
		Name:       opGetDirectoryLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDirectoryLimitsInput{}
	}

	req := c.newRequest(op, input, &GetDirectoryLimitsOutput{})

	return GetDirectoryLimitsRequest{Request: req, Input: input, Copy: c.GetDirectoryLimitsRequest}
}

// GetDirectoryLimitsRequest is the request type for the
// GetDirectoryLimits API operation.
type GetDirectoryLimitsRequest struct {
	*aws.Request
	Input *GetDirectoryLimitsInput
	Copy  func(*GetDirectoryLimitsInput) GetDirectoryLimitsRequest
}

// Send marshals and sends the GetDirectoryLimits API request.
func (r GetDirectoryLimitsRequest) Send(ctx context.Context) (*GetDirectoryLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDirectoryLimitsResponse{
		GetDirectoryLimitsOutput: r.Request.Data.(*GetDirectoryLimitsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDirectoryLimitsResponse is the response type for the
// GetDirectoryLimits API operation.
type GetDirectoryLimitsResponse struct {
	*GetDirectoryLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDirectoryLimits request.
func (r *GetDirectoryLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
