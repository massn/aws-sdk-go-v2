// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account ID and email address pairs of the accounts that you want
	// to associate with the master GuardDuty account.
	//
	// AccountDetails is a required field
	AccountDetails []AccountDetail `locationName:"accountDetails" min:"1" type:"list" required:"true"`

	// The unique ID of the detector of the GuardDuty account that you want to associate
	// member accounts with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMembersInput"}

	if s.AccountDetails == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountDetails"))
	}
	if s.AccountDetails != nil && len(s.AccountDetails) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountDetails", 1))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if s.AccountDetails != nil {
		for i, v := range s.AccountDetails {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AccountDetails", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountDetails != nil {
		v := s.AccountDetails

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountDetails", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects that include the accountIds of the unprocessed accounts
	// and a result string that explains why each was unprocessed.
	//
	// UnprocessedAccounts is a required field
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opCreateMembers = "CreateMembers"

// CreateMembersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Creates member accounts of the current AWS account by specifying a list of
// AWS account IDs. The current AWS account can then invite these members to
// manage GuardDuty in their accounts.
//
//    // Example sending a request using CreateMembersRequest.
//    req := client.CreateMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreateMembers
func (c *Client) CreateMembersRequest(input *CreateMembersInput) CreateMembersRequest {
	op := &aws.Operation{
		Name:       opCreateMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/member",
	}

	if input == nil {
		input = &CreateMembersInput{}
	}

	req := c.newRequest(op, input, &CreateMembersOutput{})

	return CreateMembersRequest{Request: req, Input: input, Copy: c.CreateMembersRequest}
}

// CreateMembersRequest is the request type for the
// CreateMembers API operation.
type CreateMembersRequest struct {
	*aws.Request
	Input *CreateMembersInput
	Copy  func(*CreateMembersInput) CreateMembersRequest
}

// Send marshals and sends the CreateMembers API request.
func (r CreateMembersRequest) Send(ctx context.Context) (*CreateMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMembersResponse{
		CreateMembersOutput: r.Request.Data.(*CreateMembersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMembersResponse is the response type for the
// CreateMembers API operation.
type CreateMembersResponse struct {
	*CreateMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMembers request.
func (r *CreateMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
