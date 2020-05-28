// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request that represents an offering renewal.
type RenewOfferingInput struct {
	_ struct{} `type:"structure"`

	// The ID of a request to renew an offering.
	OfferingId *string `locationName:"offeringId" min:"32" type:"string"`

	// The quantity requested in an offering renewal.
	Quantity *int64 `locationName:"quantity" type:"integer"`
}

// String returns the string representation
func (s RenewOfferingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RenewOfferingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RenewOfferingInput"}
	if s.OfferingId != nil && len(*s.OfferingId) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("OfferingId", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a renewal offering.
type RenewOfferingOutput struct {
	_ struct{} `type:"structure"`

	// Represents the status of the offering transaction for the renewal.
	OfferingTransaction *OfferingTransaction `locationName:"offeringTransaction" type:"structure"`
}

// String returns the string representation
func (s RenewOfferingOutput) String() string {
	return awsutil.Prettify(s)
}

const opRenewOffering = "RenewOffering"

// RenewOfferingRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Explicitly sets the quantity of devices to renew for an offering, starting
// from the effectiveDate of the next period. The API returns a NotEligible
// error if the user is not permitted to invoke the operation. If you must be
// able to invoke this operation, contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com).
//
//    // Example sending a request using RenewOfferingRequest.
//    req := client.RenewOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/RenewOffering
func (c *Client) RenewOfferingRequest(input *RenewOfferingInput) RenewOfferingRequest {
	op := &aws.Operation{
		Name:       opRenewOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RenewOfferingInput{}
	}

	req := c.newRequest(op, input, &RenewOfferingOutput{})

	return RenewOfferingRequest{Request: req, Input: input, Copy: c.RenewOfferingRequest}
}

// RenewOfferingRequest is the request type for the
// RenewOffering API operation.
type RenewOfferingRequest struct {
	*aws.Request
	Input *RenewOfferingInput
	Copy  func(*RenewOfferingInput) RenewOfferingRequest
}

// Send marshals and sends the RenewOffering API request.
func (r RenewOfferingRequest) Send(ctx context.Context) (*RenewOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RenewOfferingResponse{
		RenewOfferingOutput: r.Request.Data.(*RenewOfferingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RenewOfferingResponse is the response type for the
// RenewOffering API operation.
type RenewOfferingResponse struct {
	*RenewOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RenewOffering request.
func (r *RenewOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
