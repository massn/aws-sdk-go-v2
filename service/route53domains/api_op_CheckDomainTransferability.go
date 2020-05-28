// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The CheckDomainTransferability request contains the following elements.
type CheckDomainTransferabilityInput struct {
	_ struct{} `type:"structure"`

	// If the registrar for the top-level domain (TLD) requires an authorization
	// code to transfer the domain, the code that you got from the current registrar
	// for the domain.
	AuthCode *string `type:"string" sensitive:"true"`

	// The name of the domain that you want to transfer to Route 53. The top-level
	// domain (TLD), such as .com, must be a TLD that Route 53 supports. For a list
	// of supported TLDs, see Domains that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide.
	//
	// The domain name can contain only the following characters:
	//
	//    * Letters a through z. Domain names are not case sensitive.
	//
	//    * Numbers 0 through 9.
	//
	//    * Hyphen (-). You can't specify a hyphen at the beginning or end of a
	//    label.
	//
	//    * Period (.) to separate the labels in the name, such as the . in example.com.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CheckDomainTransferabilityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CheckDomainTransferabilityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CheckDomainTransferabilityInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The CheckDomainTransferability response includes the following elements.
type CheckDomainTransferabilityOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about whether the specified domain
	// can be transferred to Route 53.
	//
	// Transferability is a required field
	Transferability *DomainTransferability `type:"structure" required:"true"`
}

// String returns the string representation
func (s CheckDomainTransferabilityOutput) String() string {
	return awsutil.Prettify(s)
}

const opCheckDomainTransferability = "CheckDomainTransferability"

// CheckDomainTransferabilityRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// Checks whether a domain name can be transferred to Amazon Route 53.
//
//    // Example sending a request using CheckDomainTransferabilityRequest.
//    req := client.CheckDomainTransferabilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/CheckDomainTransferability
func (c *Client) CheckDomainTransferabilityRequest(input *CheckDomainTransferabilityInput) CheckDomainTransferabilityRequest {
	op := &aws.Operation{
		Name:       opCheckDomainTransferability,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CheckDomainTransferabilityInput{}
	}

	req := c.newRequest(op, input, &CheckDomainTransferabilityOutput{})

	return CheckDomainTransferabilityRequest{Request: req, Input: input, Copy: c.CheckDomainTransferabilityRequest}
}

// CheckDomainTransferabilityRequest is the request type for the
// CheckDomainTransferability API operation.
type CheckDomainTransferabilityRequest struct {
	*aws.Request
	Input *CheckDomainTransferabilityInput
	Copy  func(*CheckDomainTransferabilityInput) CheckDomainTransferabilityRequest
}

// Send marshals and sends the CheckDomainTransferability API request.
func (r CheckDomainTransferabilityRequest) Send(ctx context.Context) (*CheckDomainTransferabilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CheckDomainTransferabilityResponse{
		CheckDomainTransferabilityOutput: r.Request.Data.(*CheckDomainTransferabilityOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CheckDomainTransferabilityResponse is the response type for the
// CheckDomainTransferability API operation.
type CheckDomainTransferabilityResponse struct {
	*CheckDomainTransferabilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CheckDomainTransferability request.
func (r *CheckDomainTransferabilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
