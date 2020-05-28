// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateModelInput struct {
	_ struct{} `type:"structure"`

	// Specifies the containers in the inference pipeline.
	Containers []ContainerDefinition `type:"list"`

	// Isolates the model container. No inbound or outbound network calls can be
	// made to or from the model container.
	EnableNetworkIsolation *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the IAM role that Amazon SageMaker can
	// assume to access model artifacts and docker image for deployment on ML compute
	// instances or for batch transform jobs. Deploying on ML compute instances
	// is part of model hosting. For more information, see Amazon SageMaker Roles
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).
	//
	// To be able to pass this role to Amazon SageMaker, the caller of this API
	// must have the iam:PassRole permission.
	//
	// ExecutionRoleArn is a required field
	ExecutionRoleArn *string `min:"20" type:"string" required:"true"`

	// The name of the new model.
	//
	// ModelName is a required field
	ModelName *string `type:"string" required:"true"`

	// The location of the primary docker image containing inference code, associated
	// artifacts, and custom environment map that the inference code uses when the
	// model is deployed for predictions.
	PrimaryContainer *ContainerDefinition `type:"structure"`

	// An array of key-value pairs. For more information, see Using Cost Allocation
	// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []Tag `type:"list"`

	// A VpcConfig object that specifies the VPC that you want your model to connect
	// to. Control access to and from your model container by configuring the VPC.
	// VpcConfig is used in hosting services and in batch transform. For more information,
	// see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html)
	// and Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private
	// Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html).
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s CreateModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateModelInput"}

	if s.ExecutionRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExecutionRoleArn"))
	}
	if s.ExecutionRoleArn != nil && len(*s.ExecutionRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ExecutionRoleArn", 20))
	}

	if s.ModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelName"))
	}
	if s.Containers != nil {
		for i, v := range s.Containers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Containers", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.PrimaryContainer != nil {
		if err := s.PrimaryContainer.Validate(); err != nil {
			invalidParams.AddNested("PrimaryContainer", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			invalidParams.AddNested("VpcConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateModelOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the model created in Amazon SageMaker.
	//
	// ModelArn is a required field
	ModelArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateModelOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateModel = "CreateModel"

// CreateModelRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a model in Amazon SageMaker. In the request, you name the model and
// describe a primary container. For the primary container, you specify the
// Docker image that contains inference code, artifacts (from prior training),
// and a custom environment map that the inference code uses when you deploy
// the model for predictions.
//
// Use this API to create a model if you want to use Amazon SageMaker hosting
// services or run a batch transform job.
//
// To host your model, you create an endpoint configuration with the CreateEndpointConfig
// API, and then create an endpoint with the CreateEndpoint API. Amazon SageMaker
// then deploys all of the containers that you defined for the model in the
// hosting environment.
//
// For an example that calls this method when deploying a model to Amazon SageMaker
// hosting services, see Deploy the Model to Amazon SageMaker Hosting Services
// (AWS SDK for Python (Boto 3)). (https://docs.aws.amazon.com/sagemaker/latest/dg/ex1-deploy-model.html#ex1-deploy-model-boto)
//
// To run a batch transform using your model, you start a job with the CreateTransformJob
// API. Amazon SageMaker uses your model and your dataset to get inferences
// which are then saved to a specified S3 location.
//
// In the CreateModel request, you must define a container with the PrimaryContainer
// parameter.
//
// In the request, you also provide an IAM role that Amazon SageMaker can assume
// to access model artifacts and docker image for deployment on ML compute hosting
// instances or for batch transform jobs. In addition, you also use the IAM
// role to manage permissions the inference code needs. For example, if the
// inference code access any other AWS resources, you grant necessary permissions
// via this role.
//
//    // Example sending a request using CreateModelRequest.
//    req := client.CreateModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateModel
func (c *Client) CreateModelRequest(input *CreateModelInput) CreateModelRequest {
	op := &aws.Operation{
		Name:       opCreateModel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateModelInput{}
	}

	req := c.newRequest(op, input, &CreateModelOutput{})

	return CreateModelRequest{Request: req, Input: input, Copy: c.CreateModelRequest}
}

// CreateModelRequest is the request type for the
// CreateModel API operation.
type CreateModelRequest struct {
	*aws.Request
	Input *CreateModelInput
	Copy  func(*CreateModelInput) CreateModelRequest
}

// Send marshals and sends the CreateModel API request.
func (r CreateModelRequest) Send(ctx context.Context) (*CreateModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateModelResponse{
		CreateModelOutput: r.Request.Data.(*CreateModelOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateModelResponse is the response type for the
// CreateModel API operation.
type CreateModelResponse struct {
	*CreateModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateModel request.
func (r *CreateModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
