// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteRouteSettingsInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// RouteKey is a required field
	RouteKey *string `location:"uri" locationName:"routeKey" type:"string" required:"true"`

	// StageName is a required field
	StageName *string `location:"uri" locationName:"stageName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRouteSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRouteSettingsInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.RouteKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteKey"))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteKey != nil {
		v := *s.RouteKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "stageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteRouteSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRouteSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRouteSettings = "DeleteRouteSettings"

// DeleteRouteSettingsRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Deletes the RouteSettings for a stage.
//
//    // Example sending a request using DeleteRouteSettingsRequest.
//    req := client.DeleteRouteSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/DeleteRouteSettings
func (c *Client) DeleteRouteSettingsRequest(input *DeleteRouteSettingsInput) DeleteRouteSettingsRequest {
	op := &aws.Operation{
		Name:       opDeleteRouteSettings,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v2/apis/{apiId}/stages/{stageName}/routesettings/{routeKey}",
	}

	if input == nil {
		input = &DeleteRouteSettingsInput{}
	}

	req := c.newRequest(op, input, &DeleteRouteSettingsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteRouteSettingsRequest{Request: req, Input: input, Copy: c.DeleteRouteSettingsRequest}
}

// DeleteRouteSettingsRequest is the request type for the
// DeleteRouteSettings API operation.
type DeleteRouteSettingsRequest struct {
	*aws.Request
	Input *DeleteRouteSettingsInput
	Copy  func(*DeleteRouteSettingsInput) DeleteRouteSettingsRequest
}

// Send marshals and sends the DeleteRouteSettings API request.
func (r DeleteRouteSettingsRequest) Send(ctx context.Context) (*DeleteRouteSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRouteSettingsResponse{
		DeleteRouteSettingsOutput: r.Request.Data.(*DeleteRouteSettingsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRouteSettingsResponse is the response type for the
// DeleteRouteSettings API operation.
type DeleteRouteSettingsResponse struct {
	*DeleteRouteSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRouteSettings request.
func (r *DeleteRouteSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
