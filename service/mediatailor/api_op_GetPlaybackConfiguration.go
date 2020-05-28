// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediatailor

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetPlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `location:"uri" locationName:"Name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPlaybackConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPlaybackConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPlaybackConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetPlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The URL for the ad decision server (ADS). This includes the specification
	// of static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing, you can provide a
	// static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `type:"string"`

	// The configuration for Avail Suppression.
	AvailSuppression *AvailSuppression `type:"structure"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The configuration for DASH content.
	DashConfiguration *DashConfiguration `type:"structure"`

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration `type:"structure"`

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *LivePreRollConfiguration `type:"structure"`

	// The identifier for the playback configuration.
	Name *string `type:"string"`

	PersonalizationThresholdSeconds *int64 `min:"1" type:"integer"`

	// The Amazon Resource Name (ARN) for the playback configuration.
	PlaybackConfigurationArn *string `type:"string"`

	// The URL that the player accesses to get a manifest from AWS Elemental MediaTailor.
	// This session will use server-side reporting.
	PlaybackEndpointPrefix *string `type:"string"`

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string `type:"string"`

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill
	// in gaps in media content. Configuring the slate is optional for non-VPAID
	// playback configurations. For VPAID, the slate is required because MediaTailor
	// provides it in the slots designated for dynamic ad content. The slate must
	// be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `type:"string"`

	// The tags assigned to the playback configuration.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of MediaTailor.
	// Use this only if you have already set up custom profiles with the help of
	// AWS Support.
	TranscodeProfileName *string `type:"string"`

	// The URL prefix for the master playlist for the stream, minus the asset ID.
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s GetPlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetPlaybackConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AvailSuppression != nil {
		v := s.AvailSuppression

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "AvailSuppression", v, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.DashConfiguration != nil {
		v := s.DashConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DashConfiguration", v, metadata)
	}
	if s.HlsConfiguration != nil {
		v := s.HlsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HlsConfiguration", v, metadata)
	}
	if s.LivePreRollConfiguration != nil {
		v := s.LivePreRollConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LivePreRollConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PersonalizationThresholdSeconds != nil {
		v := *s.PersonalizationThresholdSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PersonalizationThresholdSeconds", protocol.Int64Value(v), metadata)
	}
	if s.PlaybackConfigurationArn != nil {
		v := *s.PlaybackConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlaybackEndpointPrefix != nil {
		v := *s.PlaybackEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SessionInitializationEndpointPrefix != nil {
		v := *s.SessionInitializationEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SessionInitializationEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TranscodeProfileName != nil {
		v := *s.TranscodeProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TranscodeProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetPlaybackConfiguration = "GetPlaybackConfiguration"

// GetPlaybackConfigurationRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Returns the playback configuration for the specified name.
//
//    // Example sending a request using GetPlaybackConfigurationRequest.
//    req := client.GetPlaybackConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfiguration
func (c *Client) GetPlaybackConfigurationRequest(input *GetPlaybackConfigurationInput) GetPlaybackConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetPlaybackConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfiguration/{Name}",
	}

	if input == nil {
		input = &GetPlaybackConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetPlaybackConfigurationOutput{})

	return GetPlaybackConfigurationRequest{Request: req, Input: input, Copy: c.GetPlaybackConfigurationRequest}
}

// GetPlaybackConfigurationRequest is the request type for the
// GetPlaybackConfiguration API operation.
type GetPlaybackConfigurationRequest struct {
	*aws.Request
	Input *GetPlaybackConfigurationInput
	Copy  func(*GetPlaybackConfigurationInput) GetPlaybackConfigurationRequest
}

// Send marshals and sends the GetPlaybackConfiguration API request.
func (r GetPlaybackConfigurationRequest) Send(ctx context.Context) (*GetPlaybackConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPlaybackConfigurationResponse{
		GetPlaybackConfigurationOutput: r.Request.Data.(*GetPlaybackConfigurationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPlaybackConfigurationResponse is the response type for the
// GetPlaybackConfiguration API operation.
type GetPlaybackConfigurationResponse struct {
	*GetPlaybackConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPlaybackConfiguration request.
func (r *GetPlaybackConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
