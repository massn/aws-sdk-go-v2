// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartLabelDetectionInput struct {
	_ struct{} `type:"structure"`

	// Idempotent token used to identify the start request. If you use the same
	// token with multiple StartLabelDetection requests, the same JobId is returned.
	// Use ClientRequestToken to prevent the same job from being accidently started
	// more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// An identifier you specify that's returned in the completion notification
	// that's published to your Amazon Simple Notification Service topic. For example,
	// you can use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string `min:"1" type:"string"`

	// Specifies the minimum confidence that Amazon Rekognition Video must have
	// in order to return a detected label. Confidence represents how certain Amazon
	// Rekognition is that a label is correctly identified.0 is the lowest confidence.
	// 100 is the highest confidence. Amazon Rekognition Video doesn't return any
	// labels with a confidence level lower than this specified value.
	//
	// If you don't specify MinConfidence, the operation returns labels with confidence
	// values greater than or equal to 50 percent.
	MinConfidence *float64 `type:"float"`

	// The Amazon SNS topic ARN you want Amazon Rekognition Video to publish the
	// completion status of the label detection operation to.
	NotificationChannel *NotificationChannel `type:"structure"`

	// The video in which you want to detect labels. The video must be stored in
	// an Amazon S3 bucket.
	//
	// Video is a required field
	Video *Video `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartLabelDetectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartLabelDetectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartLabelDetectionInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}

	if s.Video == nil {
		invalidParams.Add(aws.NewErrParamRequired("Video"))
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			invalidParams.AddNested("Video", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartLabelDetectionOutput struct {
	_ struct{} `type:"structure"`

	// The identifier for the label detection job. Use JobId to identify the job
	// in a subsequent call to GetLabelDetection.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartLabelDetectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartLabelDetection = "StartLabelDetection"

// StartLabelDetectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts asynchronous detection of labels in a stored video.
//
// Amazon Rekognition Video can detect labels in a video. Labels are instances
// of real-world entities. This includes objects like flower, tree, and table;
// events like wedding, graduation, and birthday party; concepts like landscape,
// evening, and nature; and activities like a person getting out of a car or
// a person skiing.
//
// The video must be stored in an Amazon S3 bucket. Use Video to specify the
// bucket name and the filename of the video. StartLabelDetection returns a
// job identifier (JobId) which you use to get the results of the operation.
// When label detection is finished, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic that you specify in
// NotificationChannel.
//
// To get the results of the label detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. If so, call
// GetLabelDetection and pass the job identifier (JobId) from the initial call
// to StartLabelDetection.
//
//    // Example sending a request using StartLabelDetectionRequest.
//    req := client.StartLabelDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartLabelDetectionRequest(input *StartLabelDetectionInput) StartLabelDetectionRequest {
	op := &aws.Operation{
		Name:       opStartLabelDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartLabelDetectionInput{}
	}

	req := c.newRequest(op, input, &StartLabelDetectionOutput{})

	return StartLabelDetectionRequest{Request: req, Input: input, Copy: c.StartLabelDetectionRequest}
}

// StartLabelDetectionRequest is the request type for the
// StartLabelDetection API operation.
type StartLabelDetectionRequest struct {
	*aws.Request
	Input *StartLabelDetectionInput
	Copy  func(*StartLabelDetectionInput) StartLabelDetectionRequest
}

// Send marshals and sends the StartLabelDetection API request.
func (r StartLabelDetectionRequest) Send(ctx context.Context) (*StartLabelDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartLabelDetectionResponse{
		StartLabelDetectionOutput: r.Request.Data.(*StartLabelDetectionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartLabelDetectionResponse is the response type for the
// StartLabelDetection API operation.
type StartLabelDetectionResponse struct {
	*StartLabelDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartLabelDetection request.
func (r *StartLabelDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
