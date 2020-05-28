// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type MigrateWorkspaceInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the target bundle type to migrate the WorkSpace to.
	//
	// BundleId is a required field
	BundleId *string `type:"string" required:"true"`

	// The identifier of the WorkSpace to migrate from.
	//
	// SourceWorkspaceId is a required field
	SourceWorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s MigrateWorkspaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MigrateWorkspaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MigrateWorkspaceInput"}

	if s.BundleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BundleId"))
	}

	if s.SourceWorkspaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceWorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type MigrateWorkspaceOutput struct {
	_ struct{} `type:"structure"`

	// The original identifier of the WorkSpace that is being migrated.
	SourceWorkspaceId *string `type:"string"`

	// The new identifier of the WorkSpace that is being migrated. If the migration
	// does not succeed, the target WorkSpace ID will not be used, and the WorkSpace
	// will still have the original WorkSpace ID.
	TargetWorkspaceId *string `type:"string"`
}

// String returns the string representation
func (s MigrateWorkspaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opMigrateWorkspace = "MigrateWorkspace"

// MigrateWorkspaceRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Migrates a WorkSpace from one operating system or bundle type to another,
// while retaining the data on the user volume.
//
// The migration process recreates the WorkSpace by using a new root volume
// from the target bundle image and the user volume from the last available
// snapshot of the original WorkSpace. During migration, the original D:\Users\%USERNAME%
// user profile folder is renamed to D:\Users\%USERNAME%MMddyyTHHmmss%.NotMigrated.
// A new D:\Users\%USERNAME%\ folder is generated by the new OS. Certain files
// in the old user profile are moved to the new user profile.
//
// For available migration scenarios, details about what happens during migration,
// and best practices, see Migrate a WorkSpace (https://docs.aws.amazon.com/workspaces/latest/adminguide/migrate-workspaces.html).
//
//    // Example sending a request using MigrateWorkspaceRequest.
//    req := client.MigrateWorkspaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/MigrateWorkspace
func (c *Client) MigrateWorkspaceRequest(input *MigrateWorkspaceInput) MigrateWorkspaceRequest {
	op := &aws.Operation{
		Name:       opMigrateWorkspace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MigrateWorkspaceInput{}
	}

	req := c.newRequest(op, input, &MigrateWorkspaceOutput{})

	return MigrateWorkspaceRequest{Request: req, Input: input, Copy: c.MigrateWorkspaceRequest}
}

// MigrateWorkspaceRequest is the request type for the
// MigrateWorkspace API operation.
type MigrateWorkspaceRequest struct {
	*aws.Request
	Input *MigrateWorkspaceInput
	Copy  func(*MigrateWorkspaceInput) MigrateWorkspaceRequest
}

// Send marshals and sends the MigrateWorkspace API request.
func (r MigrateWorkspaceRequest) Send(ctx context.Context) (*MigrateWorkspaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MigrateWorkspaceResponse{
		MigrateWorkspaceOutput: r.Request.Data.(*MigrateWorkspaceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MigrateWorkspaceResponse is the response type for the
// MigrateWorkspace API operation.
type MigrateWorkspaceResponse struct {
	*MigrateWorkspaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MigrateWorkspace request.
func (r *MigrateWorkspaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
