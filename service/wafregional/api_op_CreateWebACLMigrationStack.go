// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateWebACLMigrationStackInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether to exclude entities that can't be migrated or to stop the
	// migration. Set this to true to ignore unsupported entities in the web ACL
	// during the migration. Otherwise, if AWS WAF encounters unsupported entities,
	// it stops the process and throws an exception.
	//
	// IgnoreUnsupportedType is a required field
	IgnoreUnsupportedType *bool `type:"boolean" required:"true"`

	// The name of the Amazon S3 bucket to store the CloudFormation template in.
	// The S3 bucket must be configured as follows for the migration:
	//
	//    * The bucket name must start with aws-waf-migration-. For example, aws-waf-migration-my-web-acl.
	//
	//    * The bucket must be in the Region where you are deploying the template.
	//    For example, for a web ACL in us-west-2, you must use an Amazon S3 bucket
	//    in us-west-2 and you must deploy the template stack to us-west-2.
	//
	//    * The bucket policies must permit the migration process to write data.
	//    For listings of the bucket policies, see the Examples section.
	//
	// S3BucketName is a required field
	S3BucketName *string `min:"3" type:"string" required:"true"`

	// The UUID of the WAF Classic web ACL that you want to migrate to WAF v2.
	//
	// WebACLId is a required field
	WebACLId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateWebACLMigrationStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWebACLMigrationStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWebACLMigrationStackInput"}

	if s.IgnoreUnsupportedType == nil {
		invalidParams.Add(aws.NewErrParamRequired("IgnoreUnsupportedType"))
	}

	if s.S3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3BucketName"))
	}
	if s.S3BucketName != nil && len(*s.S3BucketName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("S3BucketName", 3))
	}

	if s.WebACLId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLId"))
	}
	if s.WebACLId != nil && len(*s.WebACLId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateWebACLMigrationStackOutput struct {
	_ struct{} `type:"structure"`

	// The URL of the template created in Amazon S3.
	//
	// S3ObjectUrl is a required field
	S3ObjectUrl *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateWebACLMigrationStackOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateWebACLMigrationStack = "CreateWebACLMigrationStack"

// CreateWebACLMigrationStackRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Creates an AWS CloudFormation WAFV2 template for the specified web ACL in
// the specified Amazon S3 bucket. Then, in CloudFormation, you create a stack
// from the template, to create the web ACL and its resources in AWS WAFV2.
// Use this to migrate your AWS WAF Classic web ACL to the latest version of
// AWS WAF.
//
// This is part of a larger migration procedure for web ACLs from AWS WAF Classic
// to the latest version of AWS WAF. For the full procedure, including caveats
// and manual steps to complete the migration and switch over to the new web
// ACL, see Migrating your AWS WAF Classic resources to AWS WAF (https://docs.aws.amazon.com/waf/latest/developerguide/waf-migrating-from-classic.html)
// in the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
//    // Example sending a request using CreateWebACLMigrationStackRequest.
//    req := client.CreateWebACLMigrationStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/CreateWebACLMigrationStack
func (c *Client) CreateWebACLMigrationStackRequest(input *CreateWebACLMigrationStackInput) CreateWebACLMigrationStackRequest {
	op := &aws.Operation{
		Name:       opCreateWebACLMigrationStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWebACLMigrationStackInput{}
	}

	req := c.newRequest(op, input, &CreateWebACLMigrationStackOutput{})

	return CreateWebACLMigrationStackRequest{Request: req, Input: input, Copy: c.CreateWebACLMigrationStackRequest}
}

// CreateWebACLMigrationStackRequest is the request type for the
// CreateWebACLMigrationStack API operation.
type CreateWebACLMigrationStackRequest struct {
	*aws.Request
	Input *CreateWebACLMigrationStackInput
	Copy  func(*CreateWebACLMigrationStackInput) CreateWebACLMigrationStackRequest
}

// Send marshals and sends the CreateWebACLMigrationStack API request.
func (r CreateWebACLMigrationStackRequest) Send(ctx context.Context) (*CreateWebACLMigrationStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWebACLMigrationStackResponse{
		CreateWebACLMigrationStackOutput: r.Request.Data.(*CreateWebACLMigrationStackOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWebACLMigrationStackResponse is the response type for the
// CreateWebACLMigrationStack API operation.
type CreateWebACLMigrationStackResponse struct {
	*CreateWebACLMigrationStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWebACLMigrationStack request.
func (r *CreateWebACLMigrationStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
