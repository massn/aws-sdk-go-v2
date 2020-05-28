// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisableLDAPSInput struct {
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
func (s DisableLDAPSInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableLDAPSInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableLDAPSInput"}

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

type DisableLDAPSOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableLDAPSOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableLDAPS = "DisableLDAPS"

// DisableLDAPSRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Deactivates LDAP secure calls for the specified directory.
//
//    // Example sending a request using DisableLDAPSRequest.
//    req := client.DisableLDAPSRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DisableLDAPS
func (c *Client) DisableLDAPSRequest(input *DisableLDAPSInput) DisableLDAPSRequest {
	op := &aws.Operation{
		Name:       opDisableLDAPS,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableLDAPSInput{}
	}

	req := c.newRequest(op, input, &DisableLDAPSOutput{})

	return DisableLDAPSRequest{Request: req, Input: input, Copy: c.DisableLDAPSRequest}
}

// DisableLDAPSRequest is the request type for the
// DisableLDAPS API operation.
type DisableLDAPSRequest struct {
	*aws.Request
	Input *DisableLDAPSInput
	Copy  func(*DisableLDAPSInput) DisableLDAPSRequest
}

// Send marshals and sends the DisableLDAPS API request.
func (r DisableLDAPSRequest) Send(ctx context.Context) (*DisableLDAPSResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableLDAPSResponse{
		DisableLDAPSOutput: r.Request.Data.(*DisableLDAPSOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableLDAPSResponse is the response type for the
// DisableLDAPS API operation.
type DisableLDAPSResponse struct {
	*DisableLDAPSOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableLDAPS request.
func (r *DisableLDAPSResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
