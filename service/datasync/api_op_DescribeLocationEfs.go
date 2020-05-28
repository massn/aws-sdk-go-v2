// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeLocationEfsRequest
type DescribeLocationEfsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the EFS location to describe.
	//
	// LocationArn is a required field
	LocationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLocationEfsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLocationEfsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLocationEfsInput"}

	if s.LocationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeLocationEfsResponse
type DescribeLocationEfsOutput struct {
	_ struct{} `type:"structure"`

	// The time that the EFS location was created.
	CreationTime *time.Time `type:"timestamp"`

	// The subnet and the security group that DataSync uses to access target EFS
	// file system. The subnet must have at least one mount target for that file
	// system. The security group that you provide needs to be able to communicate
	// with the security group on the mount target in the subnet specified.
	Ec2Config *Ec2Config `type:"structure"`

	// The Amazon resource Name (ARN) of the EFS location that was described.
	LocationArn *string `type:"string"`

	// The URL of the EFS location that was described.
	LocationUri *string `type:"string"`
}

// String returns the string representation
func (s DescribeLocationEfsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLocationEfs = "DescribeLocationEfs"

// DescribeLocationEfsRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns metadata, such as the path information about an Amazon EFS location.
//
//    // Example sending a request using DescribeLocationEfsRequest.
//    req := client.DescribeLocationEfsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationEfs
func (c *Client) DescribeLocationEfsRequest(input *DescribeLocationEfsInput) DescribeLocationEfsRequest {
	op := &aws.Operation{
		Name:       opDescribeLocationEfs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLocationEfsInput{}
	}

	req := c.newRequest(op, input, &DescribeLocationEfsOutput{})

	return DescribeLocationEfsRequest{Request: req, Input: input, Copy: c.DescribeLocationEfsRequest}
}

// DescribeLocationEfsRequest is the request type for the
// DescribeLocationEfs API operation.
type DescribeLocationEfsRequest struct {
	*aws.Request
	Input *DescribeLocationEfsInput
	Copy  func(*DescribeLocationEfsInput) DescribeLocationEfsRequest
}

// Send marshals and sends the DescribeLocationEfs API request.
func (r DescribeLocationEfsRequest) Send(ctx context.Context) (*DescribeLocationEfsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLocationEfsResponse{
		DescribeLocationEfsOutput: r.Request.Data.(*DescribeLocationEfsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLocationEfsResponse is the response type for the
// DescribeLocationEfs API operation.
type DescribeLocationEfsResponse struct {
	*DescribeLocationEfsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLocationEfs request.
func (r *DescribeLocationEfsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
