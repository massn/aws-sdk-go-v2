// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The CheckDomainAvailability request contains the following elements.
type CheckDomainAvailabilityInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to get availability for. The top-level
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
	// Internationalized domain names are not supported for some top-level domains.
	// To determine whether the TLD that you want to use supports internationalized
	// domain names, see Domains that You Can Register with Amazon Route 53 (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html).
	// For more information, see Formatting Internationalized Domain Names (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html#domain-name-format-idns).
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// Reserved for future use.
	IdnLangCode *string `type:"string"`
}

// String returns the string representation
func (s CheckDomainAvailabilityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CheckDomainAvailabilityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CheckDomainAvailabilityInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The CheckDomainAvailability response includes the following elements.
type CheckDomainAvailabilityOutput struct {
	_ struct{} `type:"structure"`

	// Whether the domain name is available for registering.
	//
	// You can register only domains designated as AVAILABLE.
	//
	// Valid values:
	//
	// AVAILABLE
	//
	// The domain name is available.
	//
	// AVAILABLE_RESERVED
	//
	// The domain name is reserved under specific conditions.
	//
	// AVAILABLE_PREORDER
	//
	// The domain name is available and can be preordered.
	//
	// DONT_KNOW
	//
	// The TLD registry didn't reply with a definitive answer about whether the
	// domain name is available. Route 53 can return this response for a variety
	// of reasons, for example, the registry is performing maintenance. Try again
	// later.
	//
	// PENDING
	//
	// The TLD registry didn't return a response in the expected amount of time.
	// When the response is delayed, it usually takes just a few extra seconds.
	// You can resubmit the request immediately.
	//
	// RESERVED
	//
	// The domain name has been reserved for another person or organization.
	//
	// UNAVAILABLE
	//
	// The domain name is not available.
	//
	// UNAVAILABLE_PREMIUM
	//
	// The domain name is not available.
	//
	// UNAVAILABLE_RESTRICTED
	//
	// The domain name is forbidden.
	//
	// Availability is a required field
	Availability DomainAvailability `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CheckDomainAvailabilityOutput) String() string {
	return awsutil.Prettify(s)
}

const opCheckDomainAvailability = "CheckDomainAvailability"

// CheckDomainAvailabilityRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation checks the availability of one domain name. Note that if the
// availability status of a domain is pending, you must submit another request
// to determine the availability of the domain name.
//
//    // Example sending a request using CheckDomainAvailabilityRequest.
//    req := client.CheckDomainAvailabilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/CheckDomainAvailability
func (c *Client) CheckDomainAvailabilityRequest(input *CheckDomainAvailabilityInput) CheckDomainAvailabilityRequest {
	op := &aws.Operation{
		Name:       opCheckDomainAvailability,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CheckDomainAvailabilityInput{}
	}

	req := c.newRequest(op, input, &CheckDomainAvailabilityOutput{})

	return CheckDomainAvailabilityRequest{Request: req, Input: input, Copy: c.CheckDomainAvailabilityRequest}
}

// CheckDomainAvailabilityRequest is the request type for the
// CheckDomainAvailability API operation.
type CheckDomainAvailabilityRequest struct {
	*aws.Request
	Input *CheckDomainAvailabilityInput
	Copy  func(*CheckDomainAvailabilityInput) CheckDomainAvailabilityRequest
}

// Send marshals and sends the CheckDomainAvailability API request.
func (r CheckDomainAvailabilityRequest) Send(ctx context.Context) (*CheckDomainAvailabilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CheckDomainAvailabilityResponse{
		CheckDomainAvailabilityOutput: r.Request.Data.(*CheckDomainAvailabilityOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CheckDomainAvailabilityResponse is the response type for the
// CheckDomainAvailability API operation.
type CheckDomainAvailabilityResponse struct {
	*CheckDomainAvailabilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CheckDomainAvailability request.
func (r *CheckDomainAvailabilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
