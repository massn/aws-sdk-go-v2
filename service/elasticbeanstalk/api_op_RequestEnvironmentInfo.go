// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Request to retrieve logs from an environment and store them in your Elastic
// Beanstalk storage bucket.
type RequestEnvironmentInfoInput struct {
	_ struct{} `type:"structure"`

	// The ID of the environment of the requested data.
	//
	// If no such environment is found, RequestEnvironmentInfo returns an InvalidParameterValue
	// error.
	//
	// Condition: You must specify either this or an EnvironmentName, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentId *string `type:"string"`

	// The name of the environment of the requested data.
	//
	// If no such environment is found, RequestEnvironmentInfo returns an InvalidParameterValue
	// error.
	//
	// Condition: You must specify either this or an EnvironmentId, or both. If
	// you do not specify either, AWS Elastic Beanstalk returns MissingRequiredParameter
	// error.
	EnvironmentName *string `min:"4" type:"string"`

	// The type of information to request.
	//
	// InfoType is a required field
	InfoType EnvironmentInfoType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s RequestEnvironmentInfoInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestEnvironmentInfoInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestEnvironmentInfoInput"}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}
	if len(s.InfoType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("InfoType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RequestEnvironmentInfoOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RequestEnvironmentInfoOutput) String() string {
	return awsutil.Prettify(s)
}

const opRequestEnvironmentInfo = "RequestEnvironmentInfo"

// RequestEnvironmentInfoRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Initiates a request to compile the specified type of information of the deployed
// environment.
//
// Setting the InfoType to tail compiles the last lines from the application
// server log files of every Amazon EC2 instance in your environment.
//
// Setting the InfoType to bundle compresses the application server log files
// for every Amazon EC2 instance into a .zip file. Legacy and .NET containers
// do not support bundle logs.
//
// Use RetrieveEnvironmentInfo to obtain the set of logs.
//
// Related Topics
//
//    * RetrieveEnvironmentInfo
//
//    // Example sending a request using RequestEnvironmentInfoRequest.
//    req := client.RequestEnvironmentInfoRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/RequestEnvironmentInfo
func (c *Client) RequestEnvironmentInfoRequest(input *RequestEnvironmentInfoInput) RequestEnvironmentInfoRequest {
	op := &aws.Operation{
		Name:       opRequestEnvironmentInfo,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RequestEnvironmentInfoInput{}
	}

	req := c.newRequest(op, input, &RequestEnvironmentInfoOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RequestEnvironmentInfoRequest{Request: req, Input: input, Copy: c.RequestEnvironmentInfoRequest}
}

// RequestEnvironmentInfoRequest is the request type for the
// RequestEnvironmentInfo API operation.
type RequestEnvironmentInfoRequest struct {
	*aws.Request
	Input *RequestEnvironmentInfoInput
	Copy  func(*RequestEnvironmentInfoInput) RequestEnvironmentInfoRequest
}

// Send marshals and sends the RequestEnvironmentInfo API request.
func (r RequestEnvironmentInfoRequest) Send(ctx context.Context) (*RequestEnvironmentInfoResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RequestEnvironmentInfoResponse{
		RequestEnvironmentInfoOutput: r.Request.Data.(*RequestEnvironmentInfoOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RequestEnvironmentInfoResponse is the response type for the
// RequestEnvironmentInfo API operation.
type RequestEnvironmentInfoResponse struct {
	*RequestEnvironmentInfoOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RequestEnvironmentInfo request.
func (r *RequestEnvironmentInfoResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
