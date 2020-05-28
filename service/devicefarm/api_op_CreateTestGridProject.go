// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTestGridProjectInput struct {
	_ struct{} `type:"structure"`

	// Human-readable description of the project.
	Description *string `locationName:"description" min:"1" type:"string"`

	// Human-readable name of the Selenium testing project.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTestGridProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTestGridProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTestGridProjectInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTestGridProjectOutput struct {
	_ struct{} `type:"structure"`

	// ARN of the Selenium testing project that was created.
	TestGridProject *TestGridProject `locationName:"testGridProject" type:"structure"`
}

// String returns the string representation
func (s CreateTestGridProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTestGridProject = "CreateTestGridProject"

// CreateTestGridProjectRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Creates a Selenium testing project. Projects are used to track TestGridSession
// instances.
//
//    // Example sending a request using CreateTestGridProjectRequest.
//    req := client.CreateTestGridProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateTestGridProject
func (c *Client) CreateTestGridProjectRequest(input *CreateTestGridProjectInput) CreateTestGridProjectRequest {
	op := &aws.Operation{
		Name:       opCreateTestGridProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTestGridProjectInput{}
	}

	req := c.newRequest(op, input, &CreateTestGridProjectOutput{})

	return CreateTestGridProjectRequest{Request: req, Input: input, Copy: c.CreateTestGridProjectRequest}
}

// CreateTestGridProjectRequest is the request type for the
// CreateTestGridProject API operation.
type CreateTestGridProjectRequest struct {
	*aws.Request
	Input *CreateTestGridProjectInput
	Copy  func(*CreateTestGridProjectInput) CreateTestGridProjectRequest
}

// Send marshals and sends the CreateTestGridProject API request.
func (r CreateTestGridProjectRequest) Send(ctx context.Context) (*CreateTestGridProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTestGridProjectResponse{
		CreateTestGridProjectOutput: r.Request.Data.(*CreateTestGridProjectOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTestGridProjectResponse is the response type for the
// CreateTestGridProject API operation.
type CreateTestGridProjectResponse struct {
	*CreateTestGridProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTestGridProject request.
func (r *CreateTestGridProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
