// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetActiveNamesInput struct {
	_ struct{} `type:"structure"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetActiveNames request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetActiveNamesInput) String() string {
	return awsutil.Prettify(s)
}

type GetActiveNamesOutput struct {
	_ struct{} `type:"structure"`

	// The list of active names returned by the get active names request.
	ActiveNames []string `locationName:"activeNames" type:"list"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetActiveNames request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetActiveNamesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetActiveNames = "GetActiveNames"

// GetActiveNamesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the names of all active (not deleted) resources.
//
//    // Example sending a request using GetActiveNamesRequest.
//    req := client.GetActiveNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetActiveNames
func (c *Client) GetActiveNamesRequest(input *GetActiveNamesInput) GetActiveNamesRequest {
	op := &aws.Operation{
		Name:       opGetActiveNames,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetActiveNamesInput{}
	}

	req := c.newRequest(op, input, &GetActiveNamesOutput{})

	return GetActiveNamesRequest{Request: req, Input: input, Copy: c.GetActiveNamesRequest}
}

// GetActiveNamesRequest is the request type for the
// GetActiveNames API operation.
type GetActiveNamesRequest struct {
	*aws.Request
	Input *GetActiveNamesInput
	Copy  func(*GetActiveNamesInput) GetActiveNamesRequest
}

// Send marshals and sends the GetActiveNames API request.
func (r GetActiveNamesRequest) Send(ctx context.Context) (*GetActiveNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetActiveNamesResponse{
		GetActiveNamesOutput: r.Request.Data.(*GetActiveNamesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetActiveNamesResponse is the response type for the
// GetActiveNames API operation.
type GetActiveNamesResponse struct {
	*GetActiveNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetActiveNames request.
func (r *GetActiveNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
