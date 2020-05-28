// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UploadLayerPartInput struct {
	_ struct{} `type:"structure"`

	// The base64-encoded layer part payload.
	//
	// LayerPartBlob is automatically base64 encoded/decoded by the SDK.
	//
	// LayerPartBlob is a required field
	LayerPartBlob []byte `locationName:"layerPartBlob" type:"blob" required:"true"`

	// The position of the first byte of the layer part witin the overall image
	// layer.
	//
	// PartFirstByte is a required field
	PartFirstByte *int64 `locationName:"partFirstByte" type:"long" required:"true"`

	// The position of the last byte of the layer part within the overall image
	// layer.
	//
	// PartLastByte is a required field
	PartLastByte *int64 `locationName:"partLastByte" type:"long" required:"true"`

	// The AWS account ID associated with the registry to which you are uploading
	// layer parts. If you do not specify a registry, the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository to which you are uploading layer parts.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`

	// The upload ID from a previous InitiateLayerUpload operation to associate
	// with the layer part upload.
	//
	// UploadId is a required field
	UploadId *string `locationName:"uploadId" type:"string" required:"true"`
}

// String returns the string representation
func (s UploadLayerPartInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadLayerPartInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadLayerPartInput"}

	if s.LayerPartBlob == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerPartBlob"))
	}

	if s.PartFirstByte == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartFirstByte"))
	}

	if s.PartLastByte == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartLastByte"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UploadLayerPartOutput struct {
	_ struct{} `type:"structure"`

	// The integer value of the last byte received in the request.
	LastByteReceived *int64 `locationName:"lastByteReceived" type:"long"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`

	// The upload ID associated with the request.
	UploadId *string `locationName:"uploadId" type:"string"`
}

// String returns the string representation
func (s UploadLayerPartOutput) String() string {
	return awsutil.Prettify(s)
}

const opUploadLayerPart = "UploadLayerPart"

// UploadLayerPartRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Uploads an image layer part to Amazon ECR.
//
// When an image is pushed, each new image layer is uploaded in parts. The maximum
// size of each image layer part can be 20971520 bytes (or about 20MB). The
// UploadLayerPart API is called once per each new image layer part.
//
// This operation is used by the Amazon ECR proxy and is not generally used
// by customers for pulling and pushing images. In most cases, you should use
// the docker CLI to pull, tag, and push images.
//
//    // Example sending a request using UploadLayerPartRequest.
//    req := client.UploadLayerPartRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/UploadLayerPart
func (c *Client) UploadLayerPartRequest(input *UploadLayerPartInput) UploadLayerPartRequest {
	op := &aws.Operation{
		Name:       opUploadLayerPart,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UploadLayerPartInput{}
	}

	req := c.newRequest(op, input, &UploadLayerPartOutput{})

	return UploadLayerPartRequest{Request: req, Input: input, Copy: c.UploadLayerPartRequest}
}

// UploadLayerPartRequest is the request type for the
// UploadLayerPart API operation.
type UploadLayerPartRequest struct {
	*aws.Request
	Input *UploadLayerPartInput
	Copy  func(*UploadLayerPartInput) UploadLayerPartRequest
}

// Send marshals and sends the UploadLayerPart API request.
func (r UploadLayerPartRequest) Send(ctx context.Context) (*UploadLayerPartResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadLayerPartResponse{
		UploadLayerPartOutput: r.Request.Data.(*UploadLayerPartOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadLayerPartResponse is the response type for the
// UploadLayerPart API operation.
type UploadLayerPartResponse struct {
	*UploadLayerPartOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadLayerPart request.
func (r *UploadLayerPartResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
