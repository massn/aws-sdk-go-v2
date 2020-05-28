// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateClassifierInput struct {
	_ struct{} `type:"structure"`

	// A CsvClassifier object specifying the classifier to create.
	CsvClassifier *CreateCsvClassifierRequest `type:"structure"`

	// A GrokClassifier object specifying the classifier to create.
	GrokClassifier *CreateGrokClassifierRequest `type:"structure"`

	// A JsonClassifier object specifying the classifier to create.
	JsonClassifier *CreateJsonClassifierRequest `type:"structure"`

	// An XMLClassifier object specifying the classifier to create.
	XMLClassifier *CreateXMLClassifierRequest `type:"structure"`
}

// String returns the string representation
func (s CreateClassifierInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateClassifierInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateClassifierInput"}
	if s.CsvClassifier != nil {
		if err := s.CsvClassifier.Validate(); err != nil {
			invalidParams.AddNested("CsvClassifier", err.(aws.ErrInvalidParams))
		}
	}
	if s.GrokClassifier != nil {
		if err := s.GrokClassifier.Validate(); err != nil {
			invalidParams.AddNested("GrokClassifier", err.(aws.ErrInvalidParams))
		}
	}
	if s.JsonClassifier != nil {
		if err := s.JsonClassifier.Validate(); err != nil {
			invalidParams.AddNested("JsonClassifier", err.(aws.ErrInvalidParams))
		}
	}
	if s.XMLClassifier != nil {
		if err := s.XMLClassifier.Validate(); err != nil {
			invalidParams.AddNested("XMLClassifier", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateClassifierOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateClassifierOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateClassifier = "CreateClassifier"

// CreateClassifierRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a classifier in the user's account. This can be a GrokClassifier,
// an XMLClassifier, a JsonClassifier, or a CsvClassifier, depending on which
// field of the request is present.
//
//    // Example sending a request using CreateClassifierRequest.
//    req := client.CreateClassifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateClassifier
func (c *Client) CreateClassifierRequest(input *CreateClassifierInput) CreateClassifierRequest {
	op := &aws.Operation{
		Name:       opCreateClassifier,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClassifierInput{}
	}

	req := c.newRequest(op, input, &CreateClassifierOutput{})

	return CreateClassifierRequest{Request: req, Input: input, Copy: c.CreateClassifierRequest}
}

// CreateClassifierRequest is the request type for the
// CreateClassifier API operation.
type CreateClassifierRequest struct {
	*aws.Request
	Input *CreateClassifierInput
	Copy  func(*CreateClassifierInput) CreateClassifierRequest
}

// Send marshals and sends the CreateClassifier API request.
func (r CreateClassifierRequest) Send(ctx context.Context) (*CreateClassifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateClassifierResponse{
		CreateClassifierOutput: r.Request.Data.(*CreateClassifierOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateClassifierResponse is the response type for the
// CreateClassifier API operation.
type CreateClassifierResponse struct {
	*CreateClassifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateClassifier request.
func (r *CreateClassifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
