// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Specifies settings to update for the trail.
type UpdateTrailInput struct {
	_ struct{} `type:"structure"`

	// Specifies a log group name using an Amazon Resource Name (ARN), a unique
	// identifier that represents the log group to which CloudTrail logs will be
	// delivered. Not required unless you specify CloudWatchLogsRoleArn.
	CloudWatchLogsLogGroupArn *string `type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user's log group.
	CloudWatchLogsRoleArn *string `type:"string"`

	// Specifies whether log file validation is enabled. The default is false.
	//
	// When you disable log file integrity validation, the chain of digest files
	// is broken after one hour. CloudTrail will not create digest files for log
	// files that were delivered during a period in which log file integrity validation
	// was disabled. For example, if you enable log file integrity validation at
	// noon on January 1, disable it at noon on January 2, and re-enable it at noon
	// on January 10, digest files will not be created for the log files delivered
	// from noon on January 2 to noon on January 10. The same applies whenever you
	// stop CloudTrail logging or delete a trail.
	EnableLogFileValidation *bool `type:"boolean"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies whether the trail applies only to the current region or to all
	// regions. The default is false. If the trail exists only in the current region
	// and this value is set to true, shadow trails (replications of the trail)
	// will be created in the other regions. If the trail exists in all regions
	// and this value is set to false, the trail will remain in the region where
	// it was created, and its shadow trails in other regions will be deleted. As
	// a best practice, consider using trails that log events in all regions.
	IsMultiRegionTrail *bool `type:"boolean"`

	// Specifies whether the trail is applied to all accounts in an organization
	// in AWS Organizations, or only for the current AWS account. The default is
	// false, and cannot be true unless the call is made on behalf of an AWS account
	// that is the master account for an organization in AWS Organizations. If the
	// trail is not an organization trail and this is set to true, the trail will
	// be created in all AWS accounts that belong to the organization. If the trail
	// is an organization trail and this is set to false, the trail will remain
	// in the current AWS account but be deleted from all member accounts in the
	// organization.
	IsOrganizationTrail *bool `type:"boolean"`

	// Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail.
	// The value can be an alias name prefixed by "alias/", a fully specified ARN
	// to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// Examples:
	//
	//    * alias/MyAliasName
	//
	//    * arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	//
	//    * arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	//    * 12345678-1234-1234-1234-123456789012
	KmsKeyId *string `type:"string"`

	// Specifies the name of the trail or trail ARN. If Name is a trail name, the
	// string must meet the following requirements:
	//
	//    * Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
	//    (_), or dashes (-)
	//
	//    * Start with a letter or number, and end with a letter or number
	//
	//    * Be between 3 and 128 characters
	//
	//    * Have no adjacent periods, underscores or dashes. Names like my-_namespace
	//    and my--namespace are invalid.
	//
	//    * Not be in IP address format (for example, 192.168.5.4)
	//
	// If Name is a trail ARN, it must be in the format:
	//
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files. See Amazon S3 Bucket Naming Requirements (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).
	S3BucketName *string `type:"string"`

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket
	// you have designated for log file delivery. For more information, see Finding
	// Your CloudTrail Log Files (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).
	// The maximum length is 200 characters.
	S3KeyPrefix *string `type:"string"`

	// Specifies the name of the Amazon SNS topic defined for notification of log
	// file delivery. The maximum length is 256 characters.
	SnsTopicName *string `type:"string"`
}

// String returns the string representation
func (s UpdateTrailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTrailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTrailInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returns the objects or data listed below if successful. Otherwise, returns
// an error.
type UpdateTrailOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the Amazon Resource Name (ARN) of the log group to which CloudTrail
	// logs will be delivered.
	CloudWatchLogsLogGroupArn *string `type:"string"`

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to
	// a user's log group.
	CloudWatchLogsRoleArn *string `type:"string"`

	// Specifies whether the trail is publishing events from global services such
	// as IAM to the log files.
	IncludeGlobalServiceEvents *bool `type:"boolean"`

	// Specifies whether the trail exists in one region or in all regions.
	IsMultiRegionTrail *bool `type:"boolean"`

	// Specifies whether the trail is an organization trail.
	IsOrganizationTrail *bool `type:"boolean"`

	// Specifies the KMS key ID that encrypts the logs delivered by CloudTrail.
	// The value is a fully specified ARN to a KMS key in the format:
	//
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string `type:"string"`

	// Specifies whether log file integrity validation is enabled.
	LogFileValidationEnabled *bool `type:"boolean"`

	// Specifies the name of the trail.
	Name *string `type:"string"`

	// Specifies the name of the Amazon S3 bucket designated for publishing log
	// files.
	S3BucketName *string `type:"string"`

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket
	// you have designated for log file delivery. For more information, see Finding
	// Your CloudTrail Log Files (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).
	S3KeyPrefix *string `type:"string"`

	// Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send notifications
	// when log files are delivered. The format of a topic ARN is:
	//
	// arn:aws:sns:us-east-2:123456789012:MyTopic
	SnsTopicARN *string `type:"string"`

	// This field is no longer in use. Use SnsTopicARN.
	SnsTopicName *string `deprecated:"true" type:"string"`

	// Specifies the ARN of the trail that was updated. The format of a trail ARN
	// is:
	//
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailARN *string `type:"string"`
}

// String returns the string representation
func (s UpdateTrailOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTrail = "UpdateTrail"

// UpdateTrailRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Updates the settings that specify delivery of log files. Changes to a trail
// do not require stopping the CloudTrail service. Use this action to designate
// an existing bucket for log delivery. If the existing bucket has previously
// been a target for CloudTrail log files, an IAM policy exists for the bucket.
// UpdateTrail must be called from the region in which the trail was created;
// otherwise, an InvalidHomeRegionException is thrown.
//
//    // Example sending a request using UpdateTrailRequest.
//    req := client.UpdateTrailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/UpdateTrail
func (c *Client) UpdateTrailRequest(input *UpdateTrailInput) UpdateTrailRequest {
	op := &aws.Operation{
		Name:       opUpdateTrail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTrailInput{}
	}

	req := c.newRequest(op, input, &UpdateTrailOutput{})

	return UpdateTrailRequest{Request: req, Input: input, Copy: c.UpdateTrailRequest}
}

// UpdateTrailRequest is the request type for the
// UpdateTrail API operation.
type UpdateTrailRequest struct {
	*aws.Request
	Input *UpdateTrailInput
	Copy  func(*UpdateTrailInput) UpdateTrailRequest
}

// Send marshals and sends the UpdateTrail API request.
func (r UpdateTrailRequest) Send(ctx context.Context) (*UpdateTrailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTrailResponse{
		UpdateTrailOutput: r.Request.Data.(*UpdateTrailOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTrailResponse is the response type for the
// UpdateTrail API operation.
type UpdateTrailResponse struct {
	*UpdateTrailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrail request.
func (r *UpdateTrailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
