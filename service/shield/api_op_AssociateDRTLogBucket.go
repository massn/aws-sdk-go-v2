// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateDRTLogBucketInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket that contains your AWS WAF logs.
	//
	// LogBucket is a required field
	LogBucket *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDRTLogBucketInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDRTLogBucketInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDRTLogBucketInput"}

	if s.LogBucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogBucket"))
	}
	if s.LogBucket != nil && len(*s.LogBucket) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LogBucket", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateDRTLogBucketOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDRTLogBucketOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateDRTLogBucket = "AssociateDRTLogBucket"

// AssociateDRTLogBucketRequest returns a request value for making API operation for
// AWS Shield.
//
// Authorizes the DDoS Response team (DRT) to access the specified Amazon S3
// bucket containing your AWS WAF logs. You can associate up to 10 Amazon S3
// buckets with your subscription.
//
// To use the services of the DRT and make an AssociateDRTLogBucket request,
// you must be subscribed to the Business Support plan (https://aws.amazon.com/premiumsupport/business-support/)
// or the Enterprise Support plan (https://aws.amazon.com/premiumsupport/enterprise-support/).
//
//    // Example sending a request using AssociateDRTLogBucketRequest.
//    req := client.AssociateDRTLogBucketRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/AssociateDRTLogBucket
func (c *Client) AssociateDRTLogBucketRequest(input *AssociateDRTLogBucketInput) AssociateDRTLogBucketRequest {
	op := &aws.Operation{
		Name:       opAssociateDRTLogBucket,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateDRTLogBucketInput{}
	}

	req := c.newRequest(op, input, &AssociateDRTLogBucketOutput{})

	return AssociateDRTLogBucketRequest{Request: req, Input: input, Copy: c.AssociateDRTLogBucketRequest}
}

// AssociateDRTLogBucketRequest is the request type for the
// AssociateDRTLogBucket API operation.
type AssociateDRTLogBucketRequest struct {
	*aws.Request
	Input *AssociateDRTLogBucketInput
	Copy  func(*AssociateDRTLogBucketInput) AssociateDRTLogBucketRequest
}

// Send marshals and sends the AssociateDRTLogBucket API request.
func (r AssociateDRTLogBucketRequest) Send(ctx context.Context) (*AssociateDRTLogBucketResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDRTLogBucketResponse{
		AssociateDRTLogBucketOutput: r.Request.Data.(*AssociateDRTLogBucketOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDRTLogBucketResponse is the response type for the
// AssociateDRTLogBucket API operation.
type AssociateDRTLogBucketResponse struct {
	*AssociateDRTLogBucketOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDRTLogBucket request.
func (r *AssociateDRTLogBucketResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
