// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeRobotApplicationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the robot application.
	//
	// Application is a required field
	Application *string `locationName:"application" min:"1" type:"string" required:"true"`

	// The version of the robot application to describe.
	ApplicationVersion *string `locationName:"applicationVersion" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeRobotApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRobotApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRobotApplicationInput"}

	if s.Application == nil {
		invalidParams.Add(aws.NewErrParamRequired("Application"))
	}
	if s.Application != nil && len(*s.Application) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Application", 1))
	}
	if s.ApplicationVersion != nil && len(*s.ApplicationVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRobotApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Application != nil {
		v := *s.Application

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "application", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationVersion != nil {
		v := *s.ApplicationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "applicationVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeRobotApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the robot application.
	Arn *string `locationName:"arn" min:"1" type:"string"`

	// The time, in milliseconds since the epoch, when the robot application was
	// last updated.
	LastUpdatedAt *time.Time `locationName:"lastUpdatedAt" type:"timestamp"`

	// The name of the robot application.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The revision id of the robot application.
	RevisionId *string `locationName:"revisionId" min:"1" type:"string"`

	// The robot software suite (ROS distribution) used by the robot application.
	RobotSoftwareSuite *RobotSoftwareSuite `locationName:"robotSoftwareSuite" type:"structure"`

	// The sources of the robot application.
	Sources []Source `locationName:"sources" type:"list"`

	// The list of all tags added to the specified robot application.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The version of the robot application.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeRobotApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRobotApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedAt != nil {
		v := *s.LastUpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedAt",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RobotSoftwareSuite != nil {
		v := s.RobotSoftwareSuite

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "robotSoftwareSuite", v, metadata)
	}
	if s.Sources != nil {
		v := s.Sources

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sources", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeRobotApplication = "DescribeRobotApplication"

// DescribeRobotApplicationRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Describes a robot application.
//
//    // Example sending a request using DescribeRobotApplicationRequest.
//    req := client.DescribeRobotApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/DescribeRobotApplication
func (c *Client) DescribeRobotApplicationRequest(input *DescribeRobotApplicationInput) DescribeRobotApplicationRequest {
	op := &aws.Operation{
		Name:       opDescribeRobotApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/describeRobotApplication",
	}

	if input == nil {
		input = &DescribeRobotApplicationInput{}
	}

	req := c.newRequest(op, input, &DescribeRobotApplicationOutput{})

	return DescribeRobotApplicationRequest{Request: req, Input: input, Copy: c.DescribeRobotApplicationRequest}
}

// DescribeRobotApplicationRequest is the request type for the
// DescribeRobotApplication API operation.
type DescribeRobotApplicationRequest struct {
	*aws.Request
	Input *DescribeRobotApplicationInput
	Copy  func(*DescribeRobotApplicationInput) DescribeRobotApplicationRequest
}

// Send marshals and sends the DescribeRobotApplication API request.
func (r DescribeRobotApplicationRequest) Send(ctx context.Context) (*DescribeRobotApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRobotApplicationResponse{
		DescribeRobotApplicationOutput: r.Request.Data.(*DescribeRobotApplicationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRobotApplicationResponse is the response type for the
// DescribeRobotApplication API operation.
type DescribeRobotApplicationResponse struct {
	*DescribeRobotApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRobotApplication request.
func (r *DescribeRobotApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
