// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateNetworkProfileInput struct {
	_ struct{} `type:"structure"`

	// The description of the network profile.
	Description *string `locationName:"description" type:"string"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	DownlinkBandwidthBits *int64 `locationName:"downlinkBandwidthBits" type:"long"`

	// Delay time for all packets to destination in milliseconds as an integer from
	// 0 to 2000.
	DownlinkDelayMs *int64 `locationName:"downlinkDelayMs" type:"long"`

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	DownlinkJitterMs *int64 `locationName:"downlinkJitterMs" type:"long"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *int64 `locationName:"downlinkLossPercent" type:"integer"`

	// The name for the new network profile.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the project for which you want to create
	// a network profile.
	//
	// ProjectArn is a required field
	ProjectArn *string `locationName:"projectArn" min:"32" type:"string" required:"true"`

	// The type of network profile to create. Valid values are listed here.
	Type NetworkProfileType `locationName:"type" type:"string" enum:"true"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	UplinkBandwidthBits *int64 `locationName:"uplinkBandwidthBits" type:"long"`

	// Delay time for all packets to destination in milliseconds as an integer from
	// 0 to 2000.
	UplinkDelayMs *int64 `locationName:"uplinkDelayMs" type:"long"`

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	UplinkJitterMs *int64 `locationName:"uplinkJitterMs" type:"long"`

	// Proportion of transmitted packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *int64 `locationName:"uplinkLossPercent" type:"integer"`
}

// String returns the string representation
func (s CreateNetworkProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNetworkProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNetworkProfileInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateNetworkProfileOutput struct {
	_ struct{} `type:"structure"`

	// The network profile that is returned by the create network profile request.
	NetworkProfile *NetworkProfile `locationName:"networkProfile" type:"structure"`
}

// String returns the string representation
func (s CreateNetworkProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNetworkProfile = "CreateNetworkProfile"

// CreateNetworkProfileRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Creates a network profile.
//
//    // Example sending a request using CreateNetworkProfileRequest.
//    req := client.CreateNetworkProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateNetworkProfile
func (c *Client) CreateNetworkProfileRequest(input *CreateNetworkProfileInput) CreateNetworkProfileRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNetworkProfileInput{}
	}

	req := c.newRequest(op, input, &CreateNetworkProfileOutput{})

	return CreateNetworkProfileRequest{Request: req, Input: input, Copy: c.CreateNetworkProfileRequest}
}

// CreateNetworkProfileRequest is the request type for the
// CreateNetworkProfile API operation.
type CreateNetworkProfileRequest struct {
	*aws.Request
	Input *CreateNetworkProfileInput
	Copy  func(*CreateNetworkProfileInput) CreateNetworkProfileRequest
}

// Send marshals and sends the CreateNetworkProfile API request.
func (r CreateNetworkProfileRequest) Send(ctx context.Context) (*CreateNetworkProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkProfileResponse{
		CreateNetworkProfileOutput: r.Request.Data.(*CreateNetworkProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkProfileResponse is the response type for the
// CreateNetworkProfile API operation.
type CreateNetworkProfileResponse struct {
	*CreateNetworkProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkProfile request.
func (r *CreateNetworkProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
