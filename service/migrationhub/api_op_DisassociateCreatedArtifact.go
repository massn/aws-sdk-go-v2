// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateCreatedArtifactInput struct {
	_ struct{} `type:"structure"`

	// An ARN of the AWS resource related to the migration (e.g., AMI, EC2 instance,
	// RDS instance, etc.)
	//
	// CreatedArtifactName is a required field
	CreatedArtifactName *string `min:"1" type:"string" required:"true"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// Unique identifier that references the migration task to be disassociated
	// with the artifact. Do not store personal data in this field.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// The name of the ProgressUpdateStream.
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateCreatedArtifactInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateCreatedArtifactInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateCreatedArtifactInput"}

	if s.CreatedArtifactName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreatedArtifactName"))
	}
	if s.CreatedArtifactName != nil && len(*s.CreatedArtifactName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CreatedArtifactName", 1))
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

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateCreatedArtifactOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateCreatedArtifactOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateCreatedArtifact = "DisassociateCreatedArtifact"

// DisassociateCreatedArtifactRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Disassociates a created artifact of an AWS resource with a migration task
// performed by a migration tool that was previously associated. This API has
// the following traits:
//
//    * A migration user can call the DisassociateCreatedArtifacts operation
//    to disassociate a created AWS Artifact from a migration task.
//
//    * The created artifact name must be provided in ARN (Amazon Resource Name)
//    format which will contain information about type and region; for example:
//    arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b.
//
//    * Examples of the AWS resource behind the created artifact are, AMI's,
//    EC2 instance, or RDS instance, etc.
//
//    // Example sending a request using DisassociateCreatedArtifactRequest.
//    req := client.DisassociateCreatedArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/DisassociateCreatedArtifact
func (c *Client) DisassociateCreatedArtifactRequest(input *DisassociateCreatedArtifactInput) DisassociateCreatedArtifactRequest {
	op := &aws.Operation{
		Name:       opDisassociateCreatedArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateCreatedArtifactInput{}
	}

	req := c.newRequest(op, input, &DisassociateCreatedArtifactOutput{})

	return DisassociateCreatedArtifactRequest{Request: req, Input: input, Copy: c.DisassociateCreatedArtifactRequest}
}

// DisassociateCreatedArtifactRequest is the request type for the
// DisassociateCreatedArtifact API operation.
type DisassociateCreatedArtifactRequest struct {
	*aws.Request
	Input *DisassociateCreatedArtifactInput
	Copy  func(*DisassociateCreatedArtifactInput) DisassociateCreatedArtifactRequest
}

// Send marshals and sends the DisassociateCreatedArtifact API request.
func (r DisassociateCreatedArtifactRequest) Send(ctx context.Context) (*DisassociateCreatedArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateCreatedArtifactResponse{
		DisassociateCreatedArtifactOutput: r.Request.Data.(*DisassociateCreatedArtifactOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateCreatedArtifactResponse is the response type for the
// DisassociateCreatedArtifact API operation.
type DisassociateCreatedArtifactResponse struct {
	*DisassociateCreatedArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateCreatedArtifact request.
func (r *DisassociateCreatedArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
