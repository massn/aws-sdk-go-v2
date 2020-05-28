// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateEndpointGroupInput struct {
	_ struct{} `type:"structure"`

	// The list of endpoint objects.
	EndpointConfigurations []EndpointConfiguration `type:"list"`

	// The Amazon Resource Name (ARN) of the endpoint group.
	//
	// EndpointGroupArn is a required field
	EndpointGroupArn *string `type:"string" required:"true"`

	// The time—10 seconds or 30 seconds—between each health check for an endpoint.
	// The default value is 30.
	HealthCheckIntervalSeconds *int64 `min:"10" type:"integer"`

	// If the protocol is HTTP/S, then this specifies the path that is the destination
	// for health check targets. The default value is slash (/).
	HealthCheckPath *string `type:"string"`

	// The port that AWS Global Accelerator uses to check the health of endpoints
	// that are part of this endpoint group. The default port is the listener port
	// that this endpoint group is associated with. If the listener port is a list
	// of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort *int64 `min:"1" type:"integer"`

	// The protocol that AWS Global Accelerator uses to check the health of endpoints
	// that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol HealthCheckProtocol `type:"string" enum:"true"`

	// The number of consecutive health checks required to set the state of a healthy
	// endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default
	// value is 3.
	ThresholdCount *int64 `min:"1" type:"integer"`

	// The percentage of traffic to send to an AWS Region. Additional traffic is
	// distributed to other endpoint groups for this listener.
	//
	// Use this action to increase (dial up) or decrease (dial down) traffic to
	// a specific Region. The percentage is applied to the traffic that would otherwise
	// have been routed to the Region based on optimal routing.
	//
	// The default value is 100.
	TrafficDialPercentage *float64 `type:"float"`
}

// String returns the string representation
func (s UpdateEndpointGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEndpointGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateEndpointGroupInput"}

	if s.EndpointGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointGroupArn"))
	}
	if s.HealthCheckIntervalSeconds != nil && *s.HealthCheckIntervalSeconds < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckIntervalSeconds", 10))
	}
	if s.HealthCheckPort != nil && *s.HealthCheckPort < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckPort", 1))
	}
	if s.ThresholdCount != nil && *s.ThresholdCount < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ThresholdCount", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateEndpointGroupOutput struct {
	_ struct{} `type:"structure"`

	// The information about the endpoint group that was updated.
	EndpointGroup *EndpointGroup `type:"structure"`
}

// String returns the string representation
func (s UpdateEndpointGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateEndpointGroup = "UpdateEndpointGroup"

// UpdateEndpointGroupRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Update an endpoint group. To see an AWS CLI example of updating an endpoint
// group, scroll down to Example.
//
//    // Example sending a request using UpdateEndpointGroupRequest.
//    req := client.UpdateEndpointGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/UpdateEndpointGroup
func (c *Client) UpdateEndpointGroupRequest(input *UpdateEndpointGroupInput) UpdateEndpointGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateEndpointGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateEndpointGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateEndpointGroupOutput{})

	return UpdateEndpointGroupRequest{Request: req, Input: input, Copy: c.UpdateEndpointGroupRequest}
}

// UpdateEndpointGroupRequest is the request type for the
// UpdateEndpointGroup API operation.
type UpdateEndpointGroupRequest struct {
	*aws.Request
	Input *UpdateEndpointGroupInput
	Copy  func(*UpdateEndpointGroupInput) UpdateEndpointGroupRequest
}

// Send marshals and sends the UpdateEndpointGroup API request.
func (r UpdateEndpointGroupRequest) Send(ctx context.Context) (*UpdateEndpointGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEndpointGroupResponse{
		UpdateEndpointGroupOutput: r.Request.Data.(*UpdateEndpointGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEndpointGroupResponse is the response type for the
// UpdateEndpointGroup API operation.
type UpdateEndpointGroupResponse struct {
	*UpdateEndpointGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEndpointGroup request.
func (r *UpdateEndpointGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
