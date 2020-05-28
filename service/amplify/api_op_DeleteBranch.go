// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for delete branch request.
type DeleteBranchInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name for the branch.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBranchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBranchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBranchInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBranchInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for delete branch request.
type DeleteBranchOutput struct {
	_ struct{} `type:"structure"`

	// Branch structure for an Amplify App.
	//
	// Branch is a required field
	Branch *Branch `locationName:"branch" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteBranchOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBranchOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Branch != nil {
		v := s.Branch

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "branch", v, metadata)
	}
	return nil
}

const opDeleteBranch = "DeleteBranch"

// DeleteBranchRequest returns a request value for making API operation for
// AWS Amplify.
//
// Deletes a branch for an Amplify App.
//
//    // Example sending a request using DeleteBranchRequest.
//    req := client.DeleteBranchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/DeleteBranch
func (c *Client) DeleteBranchRequest(input *DeleteBranchInput) DeleteBranchRequest {
	op := &aws.Operation{
		Name:       opDeleteBranch,
		HTTPMethod: "DELETE",
		HTTPPath:   "/apps/{appId}/branches/{branchName}",
	}

	if input == nil {
		input = &DeleteBranchInput{}
	}

	req := c.newRequest(op, input, &DeleteBranchOutput{})

	return DeleteBranchRequest{Request: req, Input: input, Copy: c.DeleteBranchRequest}
}

// DeleteBranchRequest is the request type for the
// DeleteBranch API operation.
type DeleteBranchRequest struct {
	*aws.Request
	Input *DeleteBranchInput
	Copy  func(*DeleteBranchInput) DeleteBranchRequest
}

// Send marshals and sends the DeleteBranch API request.
func (r DeleteBranchRequest) Send(ctx context.Context) (*DeleteBranchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBranchResponse{
		DeleteBranchOutput: r.Request.Data.(*DeleteBranchOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBranchResponse is the response type for the
// DeleteBranch API operation.
type DeleteBranchResponse struct {
	*DeleteBranchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBranch request.
func (r *DeleteBranchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
