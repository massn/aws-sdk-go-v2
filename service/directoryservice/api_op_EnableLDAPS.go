// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableLDAPSInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The type of LDAP security to enable. Currently only the value Client is supported.
	//
	// Type is a required field
	Type LDAPSType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s EnableLDAPSInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableLDAPSInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableLDAPSInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableLDAPSOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableLDAPSOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableLDAPS = "EnableLDAPS"

// EnableLDAPSRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Activates the switch for the specific directory to always use LDAP secure
// calls.
//
//    // Example sending a request using EnableLDAPSRequest.
//    req := client.EnableLDAPSRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/EnableLDAPS
func (c *Client) EnableLDAPSRequest(input *EnableLDAPSInput) EnableLDAPSRequest {
	op := &aws.Operation{
		Name:       opEnableLDAPS,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableLDAPSInput{}
	}

	req := c.newRequest(op, input, &EnableLDAPSOutput{})

	return EnableLDAPSRequest{Request: req, Input: input, Copy: c.EnableLDAPSRequest}
}

// EnableLDAPSRequest is the request type for the
// EnableLDAPS API operation.
type EnableLDAPSRequest struct {
	*aws.Request
	Input *EnableLDAPSInput
	Copy  func(*EnableLDAPSInput) EnableLDAPSRequest
}

// Send marshals and sends the EnableLDAPS API request.
func (r EnableLDAPSRequest) Send(ctx context.Context) (*EnableLDAPSResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableLDAPSResponse{
		EnableLDAPSOutput: r.Request.Data.(*EnableLDAPSOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableLDAPSResponse is the response type for the
// EnableLDAPS API operation.
type EnableLDAPSResponse struct {
	*EnableLDAPSOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableLDAPS request.
func (r *EnableLDAPSResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
