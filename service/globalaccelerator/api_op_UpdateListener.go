// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateListenerInput struct {
	_ struct{} `type:"structure"`

	// Client affinity lets you direct all requests from a user to the same endpoint,
	// if you have stateful applications, regardless of the port and protocol of
	// the client request. Clienty affinity gives you control over whether to always
	// route each client to the same specific endpoint.
	//
	// AWS Global Accelerator uses a consistent-flow hashing algorithm to choose
	// the optimal endpoint for a connection. If client affinity is NONE, Global
	// Accelerator uses the "five-tuple" (5-tuple) properties—source IP address,
	// source port, destination IP address, destination port, and protocol—to
	// select the hash value, and then chooses the best endpoint. However, with
	// this setting, if someone uses different ports to connect to Global Accelerator,
	// their connections might not be always routed to the same endpoint because
	// the hash value changes.
	//
	// If you want a given client to always be routed to the same endpoint, set
	// client affinity to SOURCE_IP instead. When you use the SOURCE_IP setting,
	// Global Accelerator uses the "two-tuple" (2-tuple) properties— source (client)
	// IP address and destination IP address—to select the hash value.
	//
	// The default value is NONE.
	ClientAffinity Affinity `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the listener to update.
	//
	// ListenerArn is a required field
	ListenerArn *string `type:"string" required:"true"`

	// The updated list of port ranges for the connections from clients to the accelerator.
	PortRanges []PortRange `min:"1" type:"list"`

	// The updated protocol for the connections from clients to the accelerator.
	Protocol Protocol `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateListenerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateListenerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateListenerInput"}

	if s.ListenerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ListenerArn"))
	}
	if s.PortRanges != nil && len(s.PortRanges) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortRanges", 1))
	}
	if s.PortRanges != nil {
		for i, v := range s.PortRanges {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PortRanges", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateListenerOutput struct {
	_ struct{} `type:"structure"`

	// Information for the updated listener.
	Listener *Listener `type:"structure"`
}

// String returns the string representation
func (s UpdateListenerOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateListener = "UpdateListener"

// UpdateListenerRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Update a listener. To see an AWS CLI example of updating listener, scroll
// down to Example.
//
//    // Example sending a request using UpdateListenerRequest.
//    req := client.UpdateListenerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/UpdateListener
func (c *Client) UpdateListenerRequest(input *UpdateListenerInput) UpdateListenerRequest {
	op := &aws.Operation{
		Name:       opUpdateListener,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateListenerInput{}
	}

	req := c.newRequest(op, input, &UpdateListenerOutput{})

	return UpdateListenerRequest{Request: req, Input: input, Copy: c.UpdateListenerRequest}
}

// UpdateListenerRequest is the request type for the
// UpdateListener API operation.
type UpdateListenerRequest struct {
	*aws.Request
	Input *UpdateListenerInput
	Copy  func(*UpdateListenerInput) UpdateListenerRequest
}

// Send marshals and sends the UpdateListener API request.
func (r UpdateListenerRequest) Send(ctx context.Context) (*UpdateListenerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateListenerResponse{
		UpdateListenerOutput: r.Request.Data.(*UpdateListenerOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateListenerResponse is the response type for the
// UpdateListener API operation.
type UpdateListenerResponse struct {
	*UpdateListenerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateListener request.
func (r *UpdateListenerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
