// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The number of items that should show up on one page
	Limit *int64 `min:"1" type:"integer"`

	// A token for requesting another page of certificates if the NextToken response
	// element indicates that more certificates are available. Use the value of
	// the returned NextToken element in your request until the token comes back
	// as null. Pass null if this is the first call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCertificatesInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// A list of certificates with basic details including certificate ID, certificate
	// common name, certificate state.
	CertificatesInfo []CertificateInfo `type:"list"`

	// Indicates whether another page of certificates is available when the number
	// of available certificates exceeds the page limit.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCertificates = "ListCertificates"

// ListCertificatesRequest returns a request value for making API operation for
// AWS Directory Service.
//
// For the specified directory, lists all the certificates registered for a
// secured LDAP connection.
//
//    // Example sending a request using ListCertificatesRequest.
//    req := client.ListCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/ListCertificates
func (c *Client) ListCertificatesRequest(input *ListCertificatesInput) ListCertificatesRequest {
	op := &aws.Operation{
		Name:       opListCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCertificatesInput{}
	}

	req := c.newRequest(op, input, &ListCertificatesOutput{})

	return ListCertificatesRequest{Request: req, Input: input, Copy: c.ListCertificatesRequest}
}

// ListCertificatesRequest is the request type for the
// ListCertificates API operation.
type ListCertificatesRequest struct {
	*aws.Request
	Input *ListCertificatesInput
	Copy  func(*ListCertificatesInput) ListCertificatesRequest
}

// Send marshals and sends the ListCertificates API request.
func (r ListCertificatesRequest) Send(ctx context.Context) (*ListCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCertificatesResponse{
		ListCertificatesOutput: r.Request.Data.(*ListCertificatesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListCertificatesResponse is the response type for the
// ListCertificates API operation.
type ListCertificatesResponse struct {
	*ListCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCertificates request.
func (r *ListCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
