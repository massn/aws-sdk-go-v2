// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateSigninDelegateGroupsFromAccountInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The sign-in delegate group names.
	//
	// GroupNames is a required field
	GroupNames []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DisassociateSigninDelegateGroupsFromAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateSigninDelegateGroupsFromAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateSigninDelegateGroupsFromAccountInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.GroupNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupNames"))
	}
	if s.GroupNames != nil && len(s.GroupNames) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupNames", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateSigninDelegateGroupsFromAccountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GroupNames != nil {
		v := s.GroupNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "GroupNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisassociateSigninDelegateGroupsFromAccountOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateSigninDelegateGroupsFromAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateSigninDelegateGroupsFromAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisassociateSigninDelegateGroupsFromAccount = "DisassociateSigninDelegateGroupsFromAccount"

// DisassociateSigninDelegateGroupsFromAccountRequest returns a request value for making API operation for
// Amazon Chime.
//
// Disassociates the specified sign-in delegate groups from the specified Amazon
// Chime account.
//
//    // Example sending a request using DisassociateSigninDelegateGroupsFromAccountRequest.
//    req := client.DisassociateSigninDelegateGroupsFromAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DisassociateSigninDelegateGroupsFromAccount
func (c *Client) DisassociateSigninDelegateGroupsFromAccountRequest(input *DisassociateSigninDelegateGroupsFromAccountInput) DisassociateSigninDelegateGroupsFromAccountRequest {
	op := &aws.Operation{
		Name:       opDisassociateSigninDelegateGroupsFromAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}?operation=disassociate-signin-delegate-groups",
	}

	if input == nil {
		input = &DisassociateSigninDelegateGroupsFromAccountInput{}
	}

	req := c.newRequest(op, input, &DisassociateSigninDelegateGroupsFromAccountOutput{})

	return DisassociateSigninDelegateGroupsFromAccountRequest{Request: req, Input: input, Copy: c.DisassociateSigninDelegateGroupsFromAccountRequest}
}

// DisassociateSigninDelegateGroupsFromAccountRequest is the request type for the
// DisassociateSigninDelegateGroupsFromAccount API operation.
type DisassociateSigninDelegateGroupsFromAccountRequest struct {
	*aws.Request
	Input *DisassociateSigninDelegateGroupsFromAccountInput
	Copy  func(*DisassociateSigninDelegateGroupsFromAccountInput) DisassociateSigninDelegateGroupsFromAccountRequest
}

// Send marshals and sends the DisassociateSigninDelegateGroupsFromAccount API request.
func (r DisassociateSigninDelegateGroupsFromAccountRequest) Send(ctx context.Context) (*DisassociateSigninDelegateGroupsFromAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateSigninDelegateGroupsFromAccountResponse{
		DisassociateSigninDelegateGroupsFromAccountOutput: r.Request.Data.(*DisassociateSigninDelegateGroupsFromAccountOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateSigninDelegateGroupsFromAccountResponse is the response type for the
// DisassociateSigninDelegateGroupsFromAccount API operation.
type DisassociateSigninDelegateGroupsFromAccountResponse struct {
	*DisassociateSigninDelegateGroupsFromAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateSigninDelegateGroupsFromAccount request.
func (r *DisassociateSigninDelegateGroupsFromAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
