// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetRecommenderConfigurationInput struct {
	_ struct{} `type:"structure"`

	// RecommenderId is a required field
	RecommenderId *string `location:"uri" locationName:"recommender-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRecommenderConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRecommenderConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRecommenderConfigurationInput"}

	if s.RecommenderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RecommenderId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRecommenderConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RecommenderId != nil {
		v := *s.RecommenderId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "recommender-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetRecommenderConfigurationOutput struct {
	_ struct{} `type:"structure" payload:"RecommenderConfigurationResponse"`

	// Provides information about Amazon Pinpoint configuration settings for retrieving
	// and processing data from a recommender model.
	//
	// RecommenderConfigurationResponse is a required field
	RecommenderConfigurationResponse *RecommenderConfigurationResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetRecommenderConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetRecommenderConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RecommenderConfigurationResponse != nil {
		v := s.RecommenderConfigurationResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "RecommenderConfigurationResponse", v, metadata)
	}
	return nil
}

const opGetRecommenderConfiguration = "GetRecommenderConfiguration"

// GetRecommenderConfigurationRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about an Amazon Pinpoint configuration for a recommender
// model.
//
//    // Example sending a request using GetRecommenderConfigurationRequest.
//    req := client.GetRecommenderConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetRecommenderConfiguration
func (c *Client) GetRecommenderConfigurationRequest(input *GetRecommenderConfigurationInput) GetRecommenderConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetRecommenderConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/recommenders/{recommender-id}",
	}

	if input == nil {
		input = &GetRecommenderConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetRecommenderConfigurationOutput{})

	return GetRecommenderConfigurationRequest{Request: req, Input: input, Copy: c.GetRecommenderConfigurationRequest}
}

// GetRecommenderConfigurationRequest is the request type for the
// GetRecommenderConfiguration API operation.
type GetRecommenderConfigurationRequest struct {
	*aws.Request
	Input *GetRecommenderConfigurationInput
	Copy  func(*GetRecommenderConfigurationInput) GetRecommenderConfigurationRequest
}

// Send marshals and sends the GetRecommenderConfiguration API request.
func (r GetRecommenderConfigurationRequest) Send(ctx context.Context) (*GetRecommenderConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRecommenderConfigurationResponse{
		GetRecommenderConfigurationOutput: r.Request.Data.(*GetRecommenderConfigurationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRecommenderConfigurationResponse is the response type for the
// GetRecommenderConfiguration API operation.
type GetRecommenderConfigurationResponse struct {
	*GetRecommenderConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRecommenderConfiguration request.
func (r *GetRecommenderConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
