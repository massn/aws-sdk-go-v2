// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The structure representing the describeProfilingGroupRequest.
type DescribeProfilingGroupInput struct {
	_ struct{} `type:"structure"`

	// The profiling group name.
	//
	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProfilingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProfilingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProfilingGroupInput"}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProfilingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The structure representing the describeProfilingGroupResponse.
type DescribeProfilingGroupOutput struct {
	_ struct{} `type:"structure" payload:"ProfilingGroup"`

	// Information about a profiling group.
	//
	// ProfilingGroup is a required field
	ProfilingGroup *ProfilingGroupDescription `locationName:"profilingGroup" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeProfilingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProfilingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ProfilingGroup != nil {
		v := s.ProfilingGroup

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "profilingGroup", v, metadata)
	}
	return nil
}

const opDescribeProfilingGroup = "DescribeProfilingGroup"

// DescribeProfilingGroupRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Describes a profiling group.
//
//    // Example sending a request using DescribeProfilingGroupRequest.
//    req := client.DescribeProfilingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/DescribeProfilingGroup
func (c *Client) DescribeProfilingGroupRequest(input *DescribeProfilingGroupInput) DescribeProfilingGroupRequest {
	op := &aws.Operation{
		Name:       opDescribeProfilingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/profilingGroups/{profilingGroupName}",
	}

	if input == nil {
		input = &DescribeProfilingGroupInput{}
	}

	req := c.newRequest(op, input, &DescribeProfilingGroupOutput{})

	return DescribeProfilingGroupRequest{Request: req, Input: input, Copy: c.DescribeProfilingGroupRequest}
}

// DescribeProfilingGroupRequest is the request type for the
// DescribeProfilingGroup API operation.
type DescribeProfilingGroupRequest struct {
	*aws.Request
	Input *DescribeProfilingGroupInput
	Copy  func(*DescribeProfilingGroupInput) DescribeProfilingGroupRequest
}

// Send marshals and sends the DescribeProfilingGroup API request.
func (r DescribeProfilingGroupRequest) Send(ctx context.Context) (*DescribeProfilingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProfilingGroupResponse{
		DescribeProfilingGroupOutput: r.Request.Data.(*DescribeProfilingGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProfilingGroupResponse is the response type for the
// DescribeProfilingGroup API operation.
type DescribeProfilingGroupResponse struct {
	*DescribeProfilingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProfilingGroup request.
func (r *DescribeProfilingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
