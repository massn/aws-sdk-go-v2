// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResumeClusterInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the cluster to be resumed.
	//
	// ClusterIdentifier is a required field
	ClusterIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResumeClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResumeClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResumeClusterInput"}

	if s.ClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResumeClusterOutput struct {
	_ struct{} `type:"structure"`

	// Describes a cluster.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s ResumeClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opResumeCluster = "ResumeCluster"

// ResumeClusterRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Resumes a paused cluster.
//
//    // Example sending a request using ResumeClusterRequest.
//    req := client.ResumeClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ResumeCluster
func (c *Client) ResumeClusterRequest(input *ResumeClusterInput) ResumeClusterRequest {
	op := &aws.Operation{
		Name:       opResumeCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResumeClusterInput{}
	}

	req := c.newRequest(op, input, &ResumeClusterOutput{})

	return ResumeClusterRequest{Request: req, Input: input, Copy: c.ResumeClusterRequest}
}

// ResumeClusterRequest is the request type for the
// ResumeCluster API operation.
type ResumeClusterRequest struct {
	*aws.Request
	Input *ResumeClusterInput
	Copy  func(*ResumeClusterInput) ResumeClusterRequest
}

// Send marshals and sends the ResumeCluster API request.
func (r ResumeClusterRequest) Send(ctx context.Context) (*ResumeClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResumeClusterResponse{
		ResumeClusterOutput: r.Request.Data.(*ResumeClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResumeClusterResponse is the response type for the
// ResumeCluster API operation.
type ResumeClusterResponse struct {
	*ResumeClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResumeCluster request.
func (r *ResumeClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
