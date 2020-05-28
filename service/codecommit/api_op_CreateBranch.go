// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a create branch operation.
type CreateBranchInput struct {
	_ struct{} `type:"structure"`

	// The name of the new branch to create.
	//
	// BranchName is a required field
	BranchName *string `locationName:"branchName" min:"1" type:"string" required:"true"`

	// The ID of the commit to point the new branch to.
	//
	// CommitId is a required field
	CommitId *string `locationName:"commitId" type:"string" required:"true"`

	// The name of the repository in which you want to create the new branch.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateBranchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBranchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBranchInput"}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.CommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommitId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateBranchOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateBranchOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateBranch = "CreateBranch"

// CreateBranchRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Creates a branch in a repository and points the branch to a commit.
//
// Calling the create branch operation does not set a repository's default branch.
// To do this, call the update default branch operation.
//
//    // Example sending a request using CreateBranchRequest.
//    req := client.CreateBranchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/CreateBranch
func (c *Client) CreateBranchRequest(input *CreateBranchInput) CreateBranchRequest {
	op := &aws.Operation{
		Name:       opCreateBranch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateBranchInput{}
	}

	req := c.newRequest(op, input, &CreateBranchOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return CreateBranchRequest{Request: req, Input: input, Copy: c.CreateBranchRequest}
}

// CreateBranchRequest is the request type for the
// CreateBranch API operation.
type CreateBranchRequest struct {
	*aws.Request
	Input *CreateBranchInput
	Copy  func(*CreateBranchInput) CreateBranchRequest
}

// Send marshals and sends the CreateBranch API request.
func (r CreateBranchRequest) Send(ctx context.Context) (*CreateBranchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBranchResponse{
		CreateBranchOutput: r.Request.Data.(*CreateBranchOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBranchResponse is the response type for the
// CreateBranch API operation.
type CreateBranchResponse struct {
	*CreateBranchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBranch request.
func (r *CreateBranchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
