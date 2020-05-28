// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateLoadBalancerListeners.
type CreateLoadBalancerListenersInput struct {
	_ struct{} `type:"structure"`

	// The listeners.
	//
	// Listeners is a required field
	Listeners []Listener `type:"list" required:"true"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLoadBalancerListenersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLoadBalancerListenersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLoadBalancerListenersInput"}

	if s.Listeners == nil {
		invalidParams.Add(aws.NewErrParamRequired("Listeners"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}
	if s.Listeners != nil {
		for i, v := range s.Listeners {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Listeners", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the parameters for CreateLoadBalancerListener.
type CreateLoadBalancerListenersOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLoadBalancerListenersOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLoadBalancerListeners = "CreateLoadBalancerListeners"

// CreateLoadBalancerListenersRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates one or more listeners for the specified load balancer. If a listener
// with the specified port does not already exist, it is created; otherwise,
// the properties of the new listener must match the properties of the existing
// listener.
//
// For more information, see Listeners for Your Classic Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using CreateLoadBalancerListenersRequest.
//    req := client.CreateLoadBalancerListenersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/CreateLoadBalancerListeners
func (c *Client) CreateLoadBalancerListenersRequest(input *CreateLoadBalancerListenersInput) CreateLoadBalancerListenersRequest {
	op := &aws.Operation{
		Name:       opCreateLoadBalancerListeners,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLoadBalancerListenersInput{}
	}

	req := c.newRequest(op, input, &CreateLoadBalancerListenersOutput{})

	return CreateLoadBalancerListenersRequest{Request: req, Input: input, Copy: c.CreateLoadBalancerListenersRequest}
}

// CreateLoadBalancerListenersRequest is the request type for the
// CreateLoadBalancerListeners API operation.
type CreateLoadBalancerListenersRequest struct {
	*aws.Request
	Input *CreateLoadBalancerListenersInput
	Copy  func(*CreateLoadBalancerListenersInput) CreateLoadBalancerListenersRequest
}

// Send marshals and sends the CreateLoadBalancerListeners API request.
func (r CreateLoadBalancerListenersRequest) Send(ctx context.Context) (*CreateLoadBalancerListenersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLoadBalancerListenersResponse{
		CreateLoadBalancerListenersOutput: r.Request.Data.(*CreateLoadBalancerListenersOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLoadBalancerListenersResponse is the response type for the
// CreateLoadBalancerListeners API operation.
type CreateLoadBalancerListenersResponse struct {
	*CreateLoadBalancerListenersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLoadBalancerListeners request.
func (r *CreateLoadBalancerListenersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
