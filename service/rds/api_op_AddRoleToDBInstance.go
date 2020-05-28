// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type AddRoleToDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB instance to associate the IAM role with.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The name of the feature for the DB instance that the IAM role is to be associated
	// with. For the list of supported feature names, see DBEngineVersion.
	//
	// FeatureName is a required field
	FeatureName *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role to associate with the DB instance,
	// for example arn:aws:iam::123456789012:role/AccessRole.
	//
	// RoleArn is a required field
	RoleArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddRoleToDBInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddRoleToDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddRoleToDBInstanceInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if s.FeatureName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FeatureName"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddRoleToDBInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddRoleToDBInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddRoleToDBInstance = "AddRoleToDBInstance"

// AddRoleToDBInstanceRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Associates an AWS Identity and Access Management (IAM) role with a DB instance.
//
// To add a role to a DB instance, the status of the DB instance must be available.
//
//    // Example sending a request using AddRoleToDBInstanceRequest.
//    req := client.AddRoleToDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/AddRoleToDBInstance
func (c *Client) AddRoleToDBInstanceRequest(input *AddRoleToDBInstanceInput) AddRoleToDBInstanceRequest {
	op := &aws.Operation{
		Name:       opAddRoleToDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddRoleToDBInstanceInput{}
	}

	req := c.newRequest(op, input, &AddRoleToDBInstanceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return AddRoleToDBInstanceRequest{Request: req, Input: input, Copy: c.AddRoleToDBInstanceRequest}
}

// AddRoleToDBInstanceRequest is the request type for the
// AddRoleToDBInstance API operation.
type AddRoleToDBInstanceRequest struct {
	*aws.Request
	Input *AddRoleToDBInstanceInput
	Copy  func(*AddRoleToDBInstanceInput) AddRoleToDBInstanceRequest
}

// Send marshals and sends the AddRoleToDBInstance API request.
func (r AddRoleToDBInstanceRequest) Send(ctx context.Context) (*AddRoleToDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddRoleToDBInstanceResponse{
		AddRoleToDBInstanceOutput: r.Request.Data.(*AddRoleToDBInstanceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddRoleToDBInstanceResponse is the response type for the
// AddRoleToDBInstance API operation.
type AddRoleToDBInstanceResponse struct {
	*AddRoleToDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddRoleToDBInstance request.
func (r *AddRoleToDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
