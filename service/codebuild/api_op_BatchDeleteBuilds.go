// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDeleteBuildsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the builds to delete.
	//
	// Ids is a required field
	Ids []string `locationName:"ids" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDeleteBuildsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeleteBuildsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeleteBuildsInput"}

	if s.Ids == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ids"))
	}
	if s.Ids != nil && len(s.Ids) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Ids", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDeleteBuildsOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the builds that were successfully deleted.
	BuildsDeleted []string `locationName:"buildsDeleted" min:"1" type:"list"`

	// Information about any builds that could not be successfully deleted.
	BuildsNotDeleted []BuildNotDeleted `locationName:"buildsNotDeleted" type:"list"`
}

// String returns the string representation
func (s BatchDeleteBuildsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDeleteBuilds = "BatchDeleteBuilds"

// BatchDeleteBuildsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Deletes one or more builds.
//
//    // Example sending a request using BatchDeleteBuildsRequest.
//    req := client.BatchDeleteBuildsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/BatchDeleteBuilds
func (c *Client) BatchDeleteBuildsRequest(input *BatchDeleteBuildsInput) BatchDeleteBuildsRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteBuilds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeleteBuildsInput{}
	}

	req := c.newRequest(op, input, &BatchDeleteBuildsOutput{})

	return BatchDeleteBuildsRequest{Request: req, Input: input, Copy: c.BatchDeleteBuildsRequest}
}

// BatchDeleteBuildsRequest is the request type for the
// BatchDeleteBuilds API operation.
type BatchDeleteBuildsRequest struct {
	*aws.Request
	Input *BatchDeleteBuildsInput
	Copy  func(*BatchDeleteBuildsInput) BatchDeleteBuildsRequest
}

// Send marshals and sends the BatchDeleteBuilds API request.
func (r BatchDeleteBuildsRequest) Send(ctx context.Context) (*BatchDeleteBuildsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteBuildsResponse{
		BatchDeleteBuildsOutput: r.Request.Data.(*BatchDeleteBuildsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteBuildsResponse is the response type for the
// BatchDeleteBuilds API operation.
type BatchDeleteBuildsResponse struct {
	*BatchDeleteBuildsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteBuilds request.
func (r *BatchDeleteBuildsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
