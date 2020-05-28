// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateStreamingURLInput struct {
	_ struct{} `type:"structure"`

	// The name of the application to launch after the session starts. This is the
	// name that you specified as Name in the Image Assistant.
	ApplicationId *string `min:"1" type:"string"`

	// The name of the fleet.
	//
	// FleetName is a required field
	FleetName *string `min:"1" type:"string" required:"true"`

	// The session context. For more information, see Session Context (https://docs.aws.amazon.com/appstream2/latest/developerguide/managing-stacks-fleets.html#managing-stacks-fleets-parameters)
	// in the Amazon AppStream 2.0 Administration Guide.
	SessionContext *string `min:"1" type:"string"`

	// The name of the stack.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`

	// The identifier of the user.
	//
	// UserId is a required field
	UserId *string `min:"2" type:"string" required:"true"`

	// The time that the streaming URL will be valid, in seconds. Specify a value
	// between 1 and 604800 seconds. The default is 60 seconds.
	Validity *int64 `type:"long"`
}

// String returns the string representation
func (s CreateStreamingURLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStreamingURLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStreamingURLInput"}
	if s.ApplicationId != nil && len(*s.ApplicationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationId", 1))
	}

	if s.FleetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetName"))
	}
	if s.FleetName != nil && len(*s.FleetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetName", 1))
	}
	if s.SessionContext != nil && len(*s.SessionContext) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SessionContext", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateStreamingURLOutput struct {
	_ struct{} `type:"structure"`

	// The elapsed time, in seconds after the Unix epoch, when this URL expires.
	Expires *time.Time `type:"timestamp"`

	// The URL to start the AppStream 2.0 streaming session.
	StreamingURL *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateStreamingURLOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStreamingURL = "CreateStreamingURL"

// CreateStreamingURLRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Creates a temporary URL to start an AppStream 2.0 streaming session for the
// specified user. A streaming URL enables application streaming to be tested
// without user setup.
//
//    // Example sending a request using CreateStreamingURLRequest.
//    req := client.CreateStreamingURLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateStreamingURL
func (c *Client) CreateStreamingURLRequest(input *CreateStreamingURLInput) CreateStreamingURLRequest {
	op := &aws.Operation{
		Name:       opCreateStreamingURL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStreamingURLInput{}
	}

	req := c.newRequest(op, input, &CreateStreamingURLOutput{})

	return CreateStreamingURLRequest{Request: req, Input: input, Copy: c.CreateStreamingURLRequest}
}

// CreateStreamingURLRequest is the request type for the
// CreateStreamingURL API operation.
type CreateStreamingURLRequest struct {
	*aws.Request
	Input *CreateStreamingURLInput
	Copy  func(*CreateStreamingURLInput) CreateStreamingURLRequest
}

// Send marshals and sends the CreateStreamingURL API request.
func (r CreateStreamingURLRequest) Send(ctx context.Context) (*CreateStreamingURLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStreamingURLResponse{
		CreateStreamingURLOutput: r.Request.Data.(*CreateStreamingURLOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStreamingURLResponse is the response type for the
// CreateStreamingURL API operation.
type CreateStreamingURLResponse struct {
	*CreateStreamingURLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStreamingURL request.
func (r *CreateStreamingURLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
