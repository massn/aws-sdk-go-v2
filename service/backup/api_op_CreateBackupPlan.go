// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateBackupPlanInput struct {
	_ struct{} `type:"structure"`

	// Specifies the body of a backup plan. Includes a BackupPlanName and one or
	// more sets of Rules.
	//
	// BackupPlan is a required field
	BackupPlan *BackupPlanInput `type:"structure" required:"true"`

	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair. The specified tags
	// are assigned to all backups created with this plan.
	BackupPlanTags map[string]string `type:"map" sensitive:"true"`

	// Identifies the request and allows failed requests to be retried without the
	// risk of executing the operation twice. If the request includes a CreatorRequestId
	// that matches an existing backup plan, that plan is returned. This parameter
	// is optional.
	CreatorRequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateBackupPlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBackupPlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBackupPlanInput"}

	if s.BackupPlan == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlan"))
	}
	if s.BackupPlan != nil {
		if err := s.BackupPlan.Validate(); err != nil {
			invalidParams.AddNested("BackupPlan", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBackupPlanInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPlan != nil {
		v := s.BackupPlan

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "BackupPlan", v, metadata)
	}
	if s.BackupPlanTags != nil {
		v := s.BackupPlanTags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "BackupPlanTags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.CreatorRequestId != nil {
		v := *s.CreatorRequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatorRequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateBackupPlanOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example, arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string `type:"string"`

	// Uniquely identifies a backup plan.
	BackupPlanId *string `type:"string"`

	// The date and time that a backup plan is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. They cannot be edited.
	VersionId *string `type:"string"`
}

// String returns the string representation
func (s CreateBackupPlanOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBackupPlanOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPlanArn != nil {
		v := *s.BackupPlanArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateBackupPlan = "CreateBackupPlan"

// CreateBackupPlanRequest returns a request value for making API operation for
// AWS Backup.
//
// Backup plans are documents that contain information that AWS Backup uses
// to schedule tasks that create recovery points of resources.
//
// If you call CreateBackupPlan with a plan that already exists, an AlreadyExistsException
// is returned.
//
//    // Example sending a request using CreateBackupPlanRequest.
//    req := client.CreateBackupPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/CreateBackupPlan
func (c *Client) CreateBackupPlanRequest(input *CreateBackupPlanInput) CreateBackupPlanRequest {
	op := &aws.Operation{
		Name:       opCreateBackupPlan,
		HTTPMethod: "PUT",
		HTTPPath:   "/backup/plans/",
	}

	if input == nil {
		input = &CreateBackupPlanInput{}
	}

	req := c.newRequest(op, input, &CreateBackupPlanOutput{})

	return CreateBackupPlanRequest{Request: req, Input: input, Copy: c.CreateBackupPlanRequest}
}

// CreateBackupPlanRequest is the request type for the
// CreateBackupPlan API operation.
type CreateBackupPlanRequest struct {
	*aws.Request
	Input *CreateBackupPlanInput
	Copy  func(*CreateBackupPlanInput) CreateBackupPlanRequest
}

// Send marshals and sends the CreateBackupPlan API request.
func (r CreateBackupPlanRequest) Send(ctx context.Context) (*CreateBackupPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBackupPlanResponse{
		CreateBackupPlanOutput: r.Request.Data.(*CreateBackupPlanOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBackupPlanResponse is the response type for the
// CreateBackupPlan API operation.
type CreateBackupPlanResponse struct {
	*CreateBackupPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBackupPlan request.
func (r *CreateBackupPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
