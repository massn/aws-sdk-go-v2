// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarconnections

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// The AWS::CodeStarConnections::Connection resource can be used to connect
// external source providers with services like AWS CodePipeline.
//
// Note: A connection created through CloudFormation is in `PENDING` status
// by default. You can make its status `AVAILABLE` by editing the connection
// in the CodePipeline console.
type Connection struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the connection. The ARN is used as the
	// connection reference when the connection is shared between AWS services.
	//
	// The ARN is never reused if the connection is deleted.
	ConnectionArn *string `type:"string"`

	// The name of the connection. Connection names must be unique in an AWS user
	// account.
	ConnectionName *string `min:"1" type:"string"`

	// The current status of the connection.
	ConnectionStatus ConnectionStatus `type:"string" enum:"true"`

	// The identifier of the external provider where your third-party code repository
	// is configured. For Bitbucket, this is the account ID of the owner of the
	// Bitbucket repository.
	OwnerAccountId *string `min:"12" type:"string"`

	// The name of the external provider where your third-party code repository
	// is configured. Currently, the valid provider type is Bitbucket.
	ProviderType ProviderType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Connection) String() string {
	return awsutil.Prettify(s)
}

// A tag is a key-value pair that is used to manage the resource.
//
// This tag is available for use by AWS services that support tags.
type Tag struct {
	_ struct{} `type:"structure"`

	// The tag's key.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The tag's value.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
