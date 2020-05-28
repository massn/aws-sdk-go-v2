// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type SetInstanceHealthInput struct {
	_ struct{} `type:"structure"`

	// The health status of the instance. Set to Healthy to have the instance remain
	// in service. Set to Unhealthy to have the instance be out of service. Amazon
	// EC2 Auto Scaling terminates and replaces the unhealthy instance.
	//
	// HealthStatus is a required field
	HealthStatus *string `min:"1" type:"string" required:"true"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"1" type:"string" required:"true"`

	// If the Auto Scaling group of the specified instance has a HealthCheckGracePeriod
	// specified for the group, by default, this call respects the grace period.
	// Set this to False, to have the call not respect the grace period associated
	// with the group.
	//
	// For more information about the health check grace period, see CreateAutoScalingGroup
	// (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_CreateAutoScalingGroup.html)
	// in the Amazon EC2 Auto Scaling API Reference.
	ShouldRespectGracePeriod *bool `type:"boolean"`
}

// String returns the string representation
func (s SetInstanceHealthInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetInstanceHealthInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetInstanceHealthInput"}

	if s.HealthStatus == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthStatus"))
	}
	if s.HealthStatus != nil && len(*s.HealthStatus) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HealthStatus", 1))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetInstanceHealthOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetInstanceHealthOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetInstanceHealth = "SetInstanceHealth"

// SetInstanceHealthRequest returns a request value for making API operation for
// Auto Scaling.
//
// Sets the health status of the specified instance.
//
// For more information, see Health Checks for Auto Scaling Instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using SetInstanceHealthRequest.
//    req := client.SetInstanceHealthRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/SetInstanceHealth
func (c *Client) SetInstanceHealthRequest(input *SetInstanceHealthInput) SetInstanceHealthRequest {
	op := &aws.Operation{
		Name:       opSetInstanceHealth,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetInstanceHealthInput{}
	}

	req := c.newRequest(op, input, &SetInstanceHealthOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return SetInstanceHealthRequest{Request: req, Input: input, Copy: c.SetInstanceHealthRequest}
}

// SetInstanceHealthRequest is the request type for the
// SetInstanceHealth API operation.
type SetInstanceHealthRequest struct {
	*aws.Request
	Input *SetInstanceHealthInput
	Copy  func(*SetInstanceHealthInput) SetInstanceHealthRequest
}

// Send marshals and sends the SetInstanceHealth API request.
func (r SetInstanceHealthRequest) Send(ctx context.Context) (*SetInstanceHealthResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetInstanceHealthResponse{
		SetInstanceHealthOutput: r.Request.Data.(*SetInstanceHealthOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetInstanceHealthResponse is the response type for the
// SetInstanceHealth API operation.
type SetInstanceHealthResponse struct {
	*SetInstanceHealthOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetInstanceHealth request.
func (r *SetInstanceHealthResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
