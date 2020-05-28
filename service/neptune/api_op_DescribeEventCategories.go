// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEventCategoriesInput struct {
	_ struct{} `type:"structure"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// The type of source that is generating the events.
	//
	// Valid values: db-instance | db-parameter-group | db-security-group | db-snapshot
	SourceType *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventCategoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventCategoriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventCategoriesInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEventCategoriesOutput struct {
	_ struct{} `type:"structure"`

	// A list of EventCategoriesMap data types.
	EventCategoriesMapList []EventCategoriesMap `locationNameList:"EventCategoriesMap" type:"list"`
}

// String returns the string representation
func (s DescribeEventCategoriesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEventCategories = "DescribeEventCategories"

// DescribeEventCategoriesRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Displays a list of categories for all event source types, or, if specified,
// for a specified source type.
//
//    // Example sending a request using DescribeEventCategoriesRequest.
//    req := client.DescribeEventCategoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeEventCategories
func (c *Client) DescribeEventCategoriesRequest(input *DescribeEventCategoriesInput) DescribeEventCategoriesRequest {
	op := &aws.Operation{
		Name:       opDescribeEventCategories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEventCategoriesInput{}
	}

	req := c.newRequest(op, input, &DescribeEventCategoriesOutput{})

	return DescribeEventCategoriesRequest{Request: req, Input: input, Copy: c.DescribeEventCategoriesRequest}
}

// DescribeEventCategoriesRequest is the request type for the
// DescribeEventCategories API operation.
type DescribeEventCategoriesRequest struct {
	*aws.Request
	Input *DescribeEventCategoriesInput
	Copy  func(*DescribeEventCategoriesInput) DescribeEventCategoriesRequest
}

// Send marshals and sends the DescribeEventCategories API request.
func (r DescribeEventCategoriesRequest) Send(ctx context.Context) (*DescribeEventCategoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventCategoriesResponse{
		DescribeEventCategoriesOutput: r.Request.Data.(*DescribeEventCategoriesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventCategoriesResponse is the response type for the
// DescribeEventCategories API operation.
type DescribeEventCategoriesResponse struct {
	*DescribeEventCategoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventCategories request.
func (r *DescribeEventCategoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
