// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type RemoveAllResourcePermissionsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"ResourceId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveAllResourcePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveAllResourcePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveAllResourcePermissionsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveAllResourcePermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ResourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RemoveAllResourcePermissionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveAllResourcePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveAllResourcePermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRemoveAllResourcePermissions = "RemoveAllResourcePermissions"

// RemoveAllResourcePermissionsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Removes all the permissions from the specified resource.
//
//    // Example sending a request using RemoveAllResourcePermissionsRequest.
//    req := client.RemoveAllResourcePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/RemoveAllResourcePermissions
func (c *Client) RemoveAllResourcePermissionsRequest(input *RemoveAllResourcePermissionsInput) RemoveAllResourcePermissionsRequest {
	op := &aws.Operation{
		Name:       opRemoveAllResourcePermissions,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/resources/{ResourceId}/permissions",
	}

	if input == nil {
		input = &RemoveAllResourcePermissionsInput{}
	}

	req := c.newRequest(op, input, &RemoveAllResourcePermissionsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return RemoveAllResourcePermissionsRequest{Request: req, Input: input, Copy: c.RemoveAllResourcePermissionsRequest}
}

// RemoveAllResourcePermissionsRequest is the request type for the
// RemoveAllResourcePermissions API operation.
type RemoveAllResourcePermissionsRequest struct {
	*aws.Request
	Input *RemoveAllResourcePermissionsInput
	Copy  func(*RemoveAllResourcePermissionsInput) RemoveAllResourcePermissionsRequest
}

// Send marshals and sends the RemoveAllResourcePermissions API request.
func (r RemoveAllResourcePermissionsRequest) Send(ctx context.Context) (*RemoveAllResourcePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveAllResourcePermissionsResponse{
		RemoveAllResourcePermissionsOutput: r.Request.Data.(*RemoveAllResourcePermissionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveAllResourcePermissionsResponse is the response type for the
// RemoveAllResourcePermissions API operation.
type RemoveAllResourcePermissionsResponse struct {
	*RemoveAllResourcePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveAllResourcePermissions request.
func (r *RemoveAllResourcePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
