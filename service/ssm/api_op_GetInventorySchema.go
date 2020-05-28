// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInventorySchemaInput struct {
	_ struct{} `type:"structure"`

	// Returns inventory schemas that support aggregation. For example, this call
	// returns the AWS:InstanceInformation type, because it supports aggregation
	// based on the PlatformName, PlatformType, and PlatformVersion attributes.
	Aggregator *bool `type:"boolean"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"50" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// Returns the sub-type schema for a specified inventory type.
	SubType *bool `type:"boolean"`

	// The type of inventory item to return.
	TypeName *string `type:"string"`
}

// String returns the string representation
func (s GetInventorySchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInventorySchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInventorySchemaInput"}
	if s.MaxResults != nil && *s.MaxResults < 50 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInventorySchemaOutput struct {
	_ struct{} `type:"structure"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`

	// Inventory schemas returned by the request.
	Schemas []InventoryItemSchema `type:"list"`
}

// String returns the string representation
func (s GetInventorySchemaOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInventorySchema = "GetInventorySchema"

// GetInventorySchemaRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Return a list of inventory type names for the account, or return a list of
// attribute names for a specific Inventory item type.
//
//    // Example sending a request using GetInventorySchemaRequest.
//    req := client.GetInventorySchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetInventorySchema
func (c *Client) GetInventorySchemaRequest(input *GetInventorySchemaInput) GetInventorySchemaRequest {
	op := &aws.Operation{
		Name:       opGetInventorySchema,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInventorySchemaInput{}
	}

	req := c.newRequest(op, input, &GetInventorySchemaOutput{})

	return GetInventorySchemaRequest{Request: req, Input: input, Copy: c.GetInventorySchemaRequest}
}

// GetInventorySchemaRequest is the request type for the
// GetInventorySchema API operation.
type GetInventorySchemaRequest struct {
	*aws.Request
	Input *GetInventorySchemaInput
	Copy  func(*GetInventorySchemaInput) GetInventorySchemaRequest
}

// Send marshals and sends the GetInventorySchema API request.
func (r GetInventorySchemaRequest) Send(ctx context.Context) (*GetInventorySchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInventorySchemaResponse{
		GetInventorySchemaOutput: r.Request.Data.(*GetInventorySchemaOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInventorySchemaResponse is the response type for the
// GetInventorySchema API operation.
type GetInventorySchemaResponse struct {
	*GetInventorySchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInventorySchema request.
func (r *GetInventorySchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
