// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateDiscoveredResourceInput struct {
	_ struct{} `type:"structure"`

	// Object representing a Resource.
	//
	// DiscoveredResource is a required field
	DiscoveredResource *DiscoveredResource `type:"structure" required:"true"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// The identifier given to the MigrationTask. Do not store personal data in
	// this field.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// The name of the ProgressUpdateStream.
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDiscoveredResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDiscoveredResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDiscoveredResourceInput"}

	if s.DiscoveredResource == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiscoveredResource"))
	}

	if s.MigrationTaskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MigrationTaskName"))
	}
	if s.MigrationTaskName != nil && len(*s.MigrationTaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MigrationTaskName", 1))
	}

	if s.ProgressUpdateStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStream"))
	}
	if s.ProgressUpdateStream != nil && len(*s.ProgressUpdateStream) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStream", 1))
	}
	if s.DiscoveredResource != nil {
		if err := s.DiscoveredResource.Validate(); err != nil {
			invalidParams.AddNested("DiscoveredResource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateDiscoveredResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDiscoveredResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateDiscoveredResource = "AssociateDiscoveredResource"

// AssociateDiscoveredResourceRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Associates a discovered resource ID from Application Discovery Service with
// a migration task.
//
//    // Example sending a request using AssociateDiscoveredResourceRequest.
//    req := client.AssociateDiscoveredResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/AssociateDiscoveredResource
func (c *Client) AssociateDiscoveredResourceRequest(input *AssociateDiscoveredResourceInput) AssociateDiscoveredResourceRequest {
	op := &aws.Operation{
		Name:       opAssociateDiscoveredResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateDiscoveredResourceInput{}
	}

	req := c.newRequest(op, input, &AssociateDiscoveredResourceOutput{})

	return AssociateDiscoveredResourceRequest{Request: req, Input: input, Copy: c.AssociateDiscoveredResourceRequest}
}

// AssociateDiscoveredResourceRequest is the request type for the
// AssociateDiscoveredResource API operation.
type AssociateDiscoveredResourceRequest struct {
	*aws.Request
	Input *AssociateDiscoveredResourceInput
	Copy  func(*AssociateDiscoveredResourceInput) AssociateDiscoveredResourceRequest
}

// Send marshals and sends the AssociateDiscoveredResource API request.
func (r AssociateDiscoveredResourceRequest) Send(ctx context.Context) (*AssociateDiscoveredResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDiscoveredResourceResponse{
		AssociateDiscoveredResourceOutput: r.Request.Data.(*AssociateDiscoveredResourceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDiscoveredResourceResponse is the response type for the
// AssociateDiscoveredResource API operation.
type AssociateDiscoveredResourceResponse struct {
	*AssociateDiscoveredResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDiscoveredResource request.
func (r *AssociateDiscoveredResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
