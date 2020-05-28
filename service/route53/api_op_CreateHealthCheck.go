// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains the health check request information.
type CreateHealthCheckInput struct {
	_ struct{} `locationName:"CreateHealthCheckRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// A unique string that identifies the request and that allows you to retry
	// a failed CreateHealthCheck request without the risk of creating two identical
	// health checks:
	//
	//    * If you send a CreateHealthCheck request with the same CallerReference
	//    and settings as a previous request, and if the health check doesn't exist,
	//    Amazon Route 53 creates the health check. If the health check does exist,
	//    Route 53 returns the settings for the existing health check.
	//
	//    * If you send a CreateHealthCheck request with the same CallerReference
	//    as a deleted health check, regardless of the settings, Route 53 returns
	//    a HealthCheckAlreadyExists error.
	//
	//    * If you send a CreateHealthCheck request with the same CallerReference
	//    as an existing health check but with different settings, Route 53 returns
	//    a HealthCheckAlreadyExists error.
	//
	//    * If you send a CreateHealthCheck request with a unique CallerReference
	//    but settings identical to an existing health check, Route 53 creates the
	//    health check.
	//
	// CallerReference is a required field
	CallerReference *string `min:"1" type:"string" required:"true"`

	// A complex type that contains settings for a new health check.
	//
	// HealthCheckConfig is a required field
	HealthCheckConfig *HealthCheckConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateHealthCheckInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHealthCheckInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHealthCheckInput"}

	if s.CallerReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("CallerReference"))
	}
	if s.CallerReference != nil && len(*s.CallerReference) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CallerReference", 1))
	}

	if s.HealthCheckConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthCheckConfig"))
	}
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			invalidParams.AddNested("HealthCheckConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateHealthCheckInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "CreateHealthCheckRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.CallerReference != nil {
			v := *s.CallerReference

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "CallerReference", protocol.StringValue(v), metadata)
		}
		if s.HealthCheckConfig != nil {
			v := s.HealthCheckConfig

			metadata := protocol.Metadata{}
			e.SetFields(protocol.BodyTarget, "HealthCheckConfig", v, metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	return nil
}

// A complex type containing the response information for the new health check.
type CreateHealthCheckOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains identifying information about the health check.
	//
	// HealthCheck is a required field
	HealthCheck *HealthCheck `type:"structure" required:"true"`

	// The unique URL representing the new health check.
	//
	// Location is a required field
	Location *string `location:"header" locationName:"Location" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateHealthCheckOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateHealthCheckOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HealthCheck != nil {
		v := s.HealthCheck

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HealthCheck", v, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	return nil
}

const opCreateHealthCheck = "CreateHealthCheck"

// CreateHealthCheckRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Creates a new health check.
//
// For information about adding health checks to resource record sets, see HealthCheckId
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html#Route53-Type-ResourceRecordSet-HealthCheckId)
// in ChangeResourceRecordSets (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html).
//
// ELB Load Balancers
//
// If you're registering EC2 instances with an Elastic Load Balancing (ELB)
// load balancer, do not create Amazon Route 53 health checks for the EC2 instances.
// When you register an EC2 instance with a load balancer, you configure settings
// for an ELB health check, which performs a similar function to a Route 53
// health check.
//
// Private Hosted Zones
//
// You can associate health checks with failover resource record sets in a private
// hosted zone. Note the following:
//
//    * Route 53 health checkers are outside the VPC. To check the health of
//    an endpoint within a VPC by IP address, you must assign a public IP address
//    to the instance in the VPC.
//
//    * You can configure a health checker to check the health of an external
//    resource that the instance relies on, such as a database server.
//
//    * You can create a CloudWatch metric, associate an alarm with the metric,
//    and then create a health check that is based on the state of the alarm.
//    For example, you might create a CloudWatch metric that checks the status
//    of the Amazon EC2 StatusCheckFailed metric, add an alarm to the metric,
//    and then create a health check that is based on the state of the alarm.
//    For information about creating CloudWatch metrics and alarms by using
//    the CloudWatch console, see the Amazon CloudWatch User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/WhatIsCloudWatch.html).
//
//    // Example sending a request using CreateHealthCheckRequest.
//    req := client.CreateHealthCheckRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateHealthCheck
func (c *Client) CreateHealthCheckRequest(input *CreateHealthCheckInput) CreateHealthCheckRequest {
	op := &aws.Operation{
		Name:       opCreateHealthCheck,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/healthcheck",
	}

	if input == nil {
		input = &CreateHealthCheckInput{}
	}

	req := c.newRequest(op, input, &CreateHealthCheckOutput{})

	return CreateHealthCheckRequest{Request: req, Input: input, Copy: c.CreateHealthCheckRequest}
}

// CreateHealthCheckRequest is the request type for the
// CreateHealthCheck API operation.
type CreateHealthCheckRequest struct {
	*aws.Request
	Input *CreateHealthCheckInput
	Copy  func(*CreateHealthCheckInput) CreateHealthCheckRequest
}

// Send marshals and sends the CreateHealthCheck API request.
func (r CreateHealthCheckRequest) Send(ctx context.Context) (*CreateHealthCheckResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHealthCheckResponse{
		CreateHealthCheckOutput: r.Request.Data.(*CreateHealthCheckOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHealthCheckResponse is the response type for the
// CreateHealthCheck API operation.
type CreateHealthCheckResponse struct {
	*CreateHealthCheckOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHealthCheck request.
func (r *CreateHealthCheckResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
