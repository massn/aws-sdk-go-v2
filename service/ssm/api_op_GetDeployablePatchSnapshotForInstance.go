// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDeployablePatchSnapshotForInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the instance for which the appropriate patch snapshot should be
	// retrieved.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The user-defined snapshot ID.
	//
	// SnapshotId is a required field
	SnapshotId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeployablePatchSnapshotForInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeployablePatchSnapshotForInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeployablePatchSnapshotForInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.SnapshotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotId"))
	}
	if s.SnapshotId != nil && len(*s.SnapshotId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("SnapshotId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDeployablePatchSnapshotForInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the instance.
	InstanceId *string `type:"string"`

	// Returns the specific operating system (for example Windows Server 2012 or
	// Amazon Linux 2015.09) on the instance for the specified patch snapshot.
	Product *string `type:"string"`

	// A pre-signed Amazon S3 URL that can be used to download the patch snapshot.
	SnapshotDownloadUrl *string `type:"string"`

	// The user-defined snapshot ID.
	SnapshotId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s GetDeployablePatchSnapshotForInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDeployablePatchSnapshotForInstance = "GetDeployablePatchSnapshotForInstance"

// GetDeployablePatchSnapshotForInstanceRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the current snapshot for the patch baseline the instance uses.
// This API is primarily used by the AWS-RunPatchBaseline Systems Manager document.
//
//    // Example sending a request using GetDeployablePatchSnapshotForInstanceRequest.
//    req := client.GetDeployablePatchSnapshotForInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetDeployablePatchSnapshotForInstance
func (c *Client) GetDeployablePatchSnapshotForInstanceRequest(input *GetDeployablePatchSnapshotForInstanceInput) GetDeployablePatchSnapshotForInstanceRequest {
	op := &aws.Operation{
		Name:       opGetDeployablePatchSnapshotForInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDeployablePatchSnapshotForInstanceInput{}
	}

	req := c.newRequest(op, input, &GetDeployablePatchSnapshotForInstanceOutput{})

	return GetDeployablePatchSnapshotForInstanceRequest{Request: req, Input: input, Copy: c.GetDeployablePatchSnapshotForInstanceRequest}
}

// GetDeployablePatchSnapshotForInstanceRequest is the request type for the
// GetDeployablePatchSnapshotForInstance API operation.
type GetDeployablePatchSnapshotForInstanceRequest struct {
	*aws.Request
	Input *GetDeployablePatchSnapshotForInstanceInput
	Copy  func(*GetDeployablePatchSnapshotForInstanceInput) GetDeployablePatchSnapshotForInstanceRequest
}

// Send marshals and sends the GetDeployablePatchSnapshotForInstance API request.
func (r GetDeployablePatchSnapshotForInstanceRequest) Send(ctx context.Context) (*GetDeployablePatchSnapshotForInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeployablePatchSnapshotForInstanceResponse{
		GetDeployablePatchSnapshotForInstanceOutput: r.Request.Data.(*GetDeployablePatchSnapshotForInstanceOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeployablePatchSnapshotForInstanceResponse is the response type for the
// GetDeployablePatchSnapshotForInstance API operation.
type GetDeployablePatchSnapshotForInstanceResponse struct {
	*GetDeployablePatchSnapshotForInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeployablePatchSnapshotForInstance request.
func (r *GetDeployablePatchSnapshotForInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
