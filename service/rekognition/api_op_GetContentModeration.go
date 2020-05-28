// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetContentModerationInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the unsafe content job. Use JobId to identify the job
	// in a subsequent call to GetContentModeration.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`

	// Maximum number of results to return per paginated call. The largest value
	// you can specify is 1000. If you specify a value greater than 1000, a maximum
	// of 1000 results is returned. The default value is 1000.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Rekognition returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of unsafe content labels.
	NextToken *string `type:"string"`

	// Sort to use for elements in the ModerationLabelDetections array. Use TIMESTAMP
	// to sort array elements by the time labels are detected. Use NAME to alphabetically
	// group elements for a label together. Within each label group, the array element
	// are sorted by detection confidence. The default sort is by TIMESTAMP.
	SortBy ContentModerationSortBy `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetContentModerationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetContentModerationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetContentModerationInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetContentModerationOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the unsafe content analysis job.
	JobStatus VideoJobStatus `type:"string" enum:"true"`

	// The detected unsafe content labels and the time(s) they were detected.
	ModerationLabels []ContentModerationDetection `type:"list"`

	// Version number of the moderation detection model that was used to detect
	// unsafe content.
	ModerationModelVersion *string `type:"string"`

	// If the response is truncated, Amazon Rekognition Video returns this token
	// that you can use in the subsequent request to retrieve the next set of unsafe
	// content labels.
	NextToken *string `type:"string"`

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string `type:"string"`

	// Information about a video that Amazon Rekognition analyzed. Videometadata
	// is returned in every page of paginated responses from GetContentModeration.
	VideoMetadata *VideoMetadata `type:"structure"`
}

// String returns the string representation
func (s GetContentModerationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetContentModeration = "GetContentModeration"

// GetContentModerationRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Gets the unsafe content analysis results for a Amazon Rekognition Video analysis
// started by StartContentModeration.
//
// Unsafe content analysis of a video is an asynchronous operation. You start
// analysis by calling StartContentModeration which returns a job identifier
// (JobId). When analysis finishes, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic registered in the
// initial call to StartContentModeration. To get the results of the unsafe
// content analysis, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED. If so, call GetContentModeration and pass the job
// identifier (JobId) from the initial call to StartContentModeration.
//
// For more information, see Working with Stored Videos in the Amazon Rekognition
// Devlopers Guide.
//
// GetContentModeration returns detected unsafe content labels, and the time
// they are detected, in an array, ModerationLabels, of ContentModerationDetection
// objects.
//
// By default, the moderated labels are returned sorted by time, in milliseconds
// from the start of the video. You can also sort them by moderated label by
// specifying NAME for the SortBy input parameter.
//
// Since video analysis can return a large number of results, use the MaxResults
// parameter to limit the number of labels returned in a single call to GetContentModeration.
// If there are more results than specified in MaxResults, the value of NextToken
// in the operation response contains a pagination token for getting the next
// set of results. To get the next page of results, call GetContentModeration
// and populate the NextToken request parameter with the value of NextToken
// returned from the previous call to GetContentModeration.
//
// For more information, see Detecting Unsafe Content in the Amazon Rekognition
// Developer Guide.
//
//    // Example sending a request using GetContentModerationRequest.
//    req := client.GetContentModerationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetContentModerationRequest(input *GetContentModerationInput) GetContentModerationRequest {
	op := &aws.Operation{
		Name:       opGetContentModeration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetContentModerationInput{}
	}

	req := c.newRequest(op, input, &GetContentModerationOutput{})

	return GetContentModerationRequest{Request: req, Input: input, Copy: c.GetContentModerationRequest}
}

// GetContentModerationRequest is the request type for the
// GetContentModeration API operation.
type GetContentModerationRequest struct {
	*aws.Request
	Input *GetContentModerationInput
	Copy  func(*GetContentModerationInput) GetContentModerationRequest
}

// Send marshals and sends the GetContentModeration API request.
func (r GetContentModerationRequest) Send(ctx context.Context) (*GetContentModerationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetContentModerationResponse{
		GetContentModerationOutput: r.Request.Data.(*GetContentModerationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetContentModerationRequestPaginator returns a paginator for GetContentModeration.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetContentModerationRequest(input)
//   p := rekognition.NewGetContentModerationRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetContentModerationPaginator(req GetContentModerationRequest) GetContentModerationPaginator {
	return GetContentModerationPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetContentModerationInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetContentModerationPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetContentModerationPaginator struct {
	aws.Pager
}

func (p *GetContentModerationPaginator) CurrentPage() *GetContentModerationOutput {
	return p.Pager.CurrentPage().(*GetContentModerationOutput)
}

// GetContentModerationResponse is the response type for the
// GetContentModeration API operation.
type GetContentModerationResponse struct {
	*GetContentModerationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetContentModeration request.
func (r *GetContentModerationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
