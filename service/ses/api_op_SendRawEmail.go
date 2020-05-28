// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to send a single raw email using Amazon SES. For more
// information, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-raw.html).
type SendRawEmailInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set to use when you send an email using SendRawEmail.
	ConfigurationSetName *string `type:"string"`

	// A list of destinations for the message, consisting of To:, CC:, and BCC:
	// addresses.
	Destinations []string `type:"list"`

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to specify a particular "From" address in the header of the raw email.
	//
	// Instead of using this parameter, you can use the X-header X-SES-FROM-ARN
	// in the raw message of the email. If you use both the FromArn parameter and
	// the corresponding X-header, Amazon SES uses the value of the FromArn parameter.
	//
	// For information about when to use this parameter, see the description of
	// SendRawEmail in this guide, or see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-delegate-sender-tasks-email.html).
	FromArn *string `type:"string"`

	// The raw email message itself. The message has to meet the following criteria:
	//
	//    * The message has to contain a header and a body, separated by a blank
	//    line.
	//
	//    * All of the required header fields must be present in the message.
	//
	//    * Each part of a multipart MIME message must be formatted properly.
	//
	//    * Attachments must be of a content type that Amazon SES supports. For
	//    a list on unsupported content types, see Unsupported Attachment Types
	//    (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mime-types.html)
	//    in the Amazon SES Developer Guide.
	//
	//    * The entire message must be base64-encoded.
	//
	//    * If any of the MIME parts in your message contain content that is outside
	//    of the 7-bit ASCII character range, we highly recommend that you encode
	//    that content. For more information, see Sending Raw Email (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-raw.html)
	//    in the Amazon SES Developer Guide.
	//
	//    * Per RFC 5321 (https://tools.ietf.org/html/rfc5321#section-4.5.3.1.6),
	//    the maximum length of each line of text, including the <CRLF>, must not
	//    exceed 1,000 characters.
	//
	// RawMessage is a required field
	RawMessage *RawMessage `type:"structure" required:"true"`

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to use the email address specified in the ReturnPath parameter.
	//
	// For example, if the owner of example.com (which has ARN arn:aws:ses:us-east-1:123456789012:identity/example.com)
	// attaches a policy to it that authorizes you to use feedback@example.com,
	// then you would specify the ReturnPathArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com,
	// and the ReturnPath to be feedback@example.com.
	//
	// Instead of using this parameter, you can use the X-header X-SES-RETURN-PATH-ARN
	// in the raw message of the email. If you use both the ReturnPathArn parameter
	// and the corresponding X-header, Amazon SES uses the value of the ReturnPathArn
	// parameter.
	//
	// For information about when to use this parameter, see the description of
	// SendRawEmail in this guide, or see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-delegate-sender-tasks-email.html).
	ReturnPathArn *string `type:"string"`

	// The identity's email address. If you do not provide a value for this parameter,
	// you must specify a "From" address in the raw text of the message. (You can
	// also specify both.)
	//
	// Amazon SES does not support the SMTPUTF8 extension, as described inRFC6531
	// (https://tools.ietf.org/html/rfc6531). For this reason, the local part of
	// a source email address (the part of the email address that precedes the @
	// sign) may only contain 7-bit ASCII characters (https://en.wikipedia.org/wiki/Email_address#Local-part).
	// If the domain part of an address (the part after the @ sign) contains non-ASCII
	// characters, they must be encoded using Punycode, as described in RFC3492
	// (https://tools.ietf.org/html/rfc3492.html). The sender name (also known as
	// the friendly name) may contain non-ASCII characters. These characters must
	// be encoded using MIME encoded-word syntax, as described in RFC 2047 (https://tools.ietf.org/html/rfc2047).
	// MIME encoded-word syntax uses the following form: =?charset?encoding?encoded-text?=.
	//
	// If you specify the Source parameter and have feedback forwarding enabled,
	// then bounces and complaints will be sent to this email address. This takes
	// precedence over any Return-Path header that you might include in the raw
	// text of the message.
	Source *string `type:"string"`

	// This parameter is used only for sending authorization. It is the ARN of the
	// identity that is associated with the sending authorization policy that permits
	// you to send for the email address specified in the Source parameter.
	//
	// For example, if the owner of example.com (which has ARN arn:aws:ses:us-east-1:123456789012:identity/example.com)
	// attaches a policy to it that authorizes you to send from user@example.com,
	// then you would specify the SourceArn to be arn:aws:ses:us-east-1:123456789012:identity/example.com,
	// and the Source to be user@example.com.
	//
	// Instead of using this parameter, you can use the X-header X-SES-SOURCE-ARN
	// in the raw message of the email. If you use both the SourceArn parameter
	// and the corresponding X-header, Amazon SES uses the value of the SourceArn
	// parameter.
	//
	// For information about when to use this parameter, see the description of
	// SendRawEmail in this guide, or see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-delegate-sender-tasks-email.html).
	SourceArn *string `type:"string"`

	// A list of tags, in the form of name/value pairs, to apply to an email that
	// you send using SendRawEmail. Tags correspond to characteristics of the email
	// that you define, so that you can publish email sending events.
	Tags []MessageTag `type:"list"`
}

// String returns the string representation
func (s SendRawEmailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendRawEmailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendRawEmailInput"}

	if s.RawMessage == nil {
		invalidParams.Add(aws.NewErrParamRequired("RawMessage"))
	}
	if s.RawMessage != nil {
		if err := s.RawMessage.Validate(); err != nil {
			invalidParams.AddNested("RawMessage", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a unique message ID.
type SendRawEmailOutput struct {
	_ struct{} `type:"structure"`

	// The unique message identifier returned from the SendRawEmail action.
	//
	// MessageId is a required field
	MessageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SendRawEmailOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendRawEmail = "SendRawEmail"

// SendRawEmailRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Composes an email message and immediately queues it for sending.
//
// This operation is more flexible than the SendEmail API operation. When you
// use the SendRawEmail operation, you can specify the headers of the message
// as well as its content. This flexibility is useful, for example, when you
// want to send a multipart MIME email (such a message that contains both a
// text and an HTML version). You can also use this operation to send messages
// that include attachments.
//
// The SendRawEmail operation has the following requirements:
//
//    * You can only send email from verified email addresses or domains (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html).
//    If you try to send email from an address that isn't verified, the operation
//    results in an "Email address not verified" error.
//
//    * If your account is still in the Amazon SES sandbox (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/request-production-access.html),
//    you can only send email to other verified addresses in your account, or
//    to addresses that are associated with the Amazon SES mailbox simulator
//    (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mailbox-simulator.html).
//
//    * The maximum message size, including attachments, is 10 MB.
//
//    * Each message has to include at least one recipient address. A recipient
//    address includes any address on the To:, CC:, or BCC: lines.
//
//    * If you send a single message to more than one recipient address, and
//    one of the recipient addresses isn't in a valid format (that is, it's
//    not in the format UserName@[SubDomain.]Domain.TopLevelDomain), Amazon
//    SES rejects the entire message, even if the other addresses are valid.
//
//    * Each message can include up to 50 recipient addresses across the To:,
//    CC:, or BCC: lines. If you need to send a single message to more than
//    50 recipients, you have to split the list of recipient addresses into
//    groups of less than 50 recipients, and send separate messages to each
//    group.
//
//    * Amazon SES allows you to specify 8-bit Content-Transfer-Encoding for
//    MIME message parts. However, if Amazon SES has to modify the contents
//    of your message (for example, if you use open and click tracking), 8-bit
//    content isn't preserved. For this reason, we highly recommend that you
//    encode all content that isn't 7-bit ASCII. For more information, see MIME
//    Encoding (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-raw.html#send-email-mime-encoding)
//    in the Amazon SES Developer Guide.
//
// Additionally, keep the following considerations in mind when using the SendRawEmail
// operation:
//
//    * Although you can customize the message headers when using the SendRawEmail
//    operation, Amazon SES will automatically apply its own Message-ID and
//    Date headers; if you passed these headers when creating the message, they
//    will be overwritten by the values that Amazon SES provides.
//
//    * If you are using sending authorization to send on behalf of another
//    user, SendRawEmail enables you to specify the cross-account identity for
//    the email's Source, From, and Return-Path parameters in one of two ways:
//    you can pass optional parameters SourceArn, FromArn, and/or ReturnPathArn
//    to the API, or you can include the following X-headers in the header of
//    your raw email: X-SES-SOURCE-ARN X-SES-FROM-ARN X-SES-RETURN-PATH-ARN
//    Don't include these X-headers in the DKIM signature. Amazon SES removes
//    these before it sends the email. If you only specify the SourceIdentityArn
//    parameter, Amazon SES sets the From and Return-Path addresses to the same
//    identity that you specified. For more information about sending authorization,
//    see the Using Sending Authorization with Amazon SES (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html)
//    in the Amazon SES Developer Guide.
//
//    * For every message that you send, the total number of recipients (including
//    each recipient in the To:, CC: and BCC: fields) is counted against the
//    maximum number of emails you can send in a 24-hour period (your sending
//    quota). For more information about sending quotas in Amazon SES, see Managing
//    Your Amazon SES Sending Limits (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/manage-sending-limits.html)
//    in the Amazon SES Developer Guide.
//
//    // Example sending a request using SendRawEmailRequest.
//    req := client.SendRawEmailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/SendRawEmail
func (c *Client) SendRawEmailRequest(input *SendRawEmailInput) SendRawEmailRequest {
	op := &aws.Operation{
		Name:       opSendRawEmail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendRawEmailInput{}
	}

	req := c.newRequest(op, input, &SendRawEmailOutput{})

	return SendRawEmailRequest{Request: req, Input: input, Copy: c.SendRawEmailRequest}
}

// SendRawEmailRequest is the request type for the
// SendRawEmail API operation.
type SendRawEmailRequest struct {
	*aws.Request
	Input *SendRawEmailInput
	Copy  func(*SendRawEmailInput) SendRawEmailRequest
}

// Send marshals and sends the SendRawEmail API request.
func (r SendRawEmailRequest) Send(ctx context.Context) (*SendRawEmailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendRawEmailResponse{
		SendRawEmailOutput: r.Request.Data.(*SendRawEmailOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendRawEmailResponse is the response type for the
// SendRawEmail API operation.
type SendRawEmailResponse struct {
	*SendRawEmailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendRawEmail request.
func (r *SendRawEmailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
