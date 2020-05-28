// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eventbridge

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutRuleInput struct {
	_ struct{} `type:"structure"`

	// A description of the rule.
	Description *string `type:"string"`

	// The event bus to associate with this rule. If you omit this, the default
	// event bus is used.
	EventBusName *string `min:"1" type:"string"`

	// The event pattern. For more information, see Events and Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	EventPattern *string `type:"string"`

	// The name of the rule that you are creating or updating.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role associated with the rule.
	RoleArn *string `min:"1" type:"string"`

	// The scheduling expression. For example, "cron(0 20 * * ? *)" or "rate(5 minutes)".
	ScheduleExpression *string `type:"string"`

	// Indicates whether the rule is enabled or disabled.
	State RuleState `type:"string" enum:"true"`

	// The list of key-value pairs to associate with the rule.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s PutRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRuleInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
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

type PutRuleOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the rule.
	RuleArn *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutRule = "PutRule"

// PutRuleRequest returns a request value for making API operation for
// Amazon EventBridge.
//
// Creates or updates the specified rule. Rules are enabled by default, or based
// on value of the state. You can disable a rule using DisableRule.
//
// A single rule watches for events from a single event bus. Events generated
// by AWS services go to your account's default event bus. Events generated
// by SaaS partner services or applications go to the matching partner event
// bus. If you have custom applications or services, you can specify whether
// their events go to your default event bus or a custom event bus that you
// have created. For more information, see CreateEventBus.
//
// If you are updating an existing rule, the rule is replaced with what you
// specify in this PutRule command. If you omit arguments in PutRule, the old
// values for those arguments are not kept. Instead, they are replaced with
// null values.
//
// When you create or update a rule, incoming events might not immediately start
// matching to new or updated rules. Allow a short period of time for changes
// to take effect.
//
// A rule must contain at least an EventPattern or ScheduleExpression. Rules
// with EventPatterns are triggered when a matching event is observed. Rules
// with ScheduleExpressions self-trigger based on the given schedule. A rule
// can have both an EventPattern and a ScheduleExpression, in which case the
// rule triggers on matching events as well as on a schedule.
//
// When you initially create a rule, you can optionally assign one or more tags
// to the rule. Tags can help you organize and categorize your resources. You
// can also use them to scope user permissions, by granting a user permission
// to access or change only rules with certain tag values. To use the PutRule
// operation and assign tags, you must have both the events:PutRule and events:TagResource
// permissions.
//
// If you are updating an existing rule, any tags you specify in the PutRule
// operation are ignored. To update the tags of an existing rule, use TagResource
// and UntagResource.
//
// Most services in AWS treat : or / as the same character in Amazon Resource
// Names (ARNs). However, EventBridge uses an exact match in event patterns
// and rules. Be sure to use the correct ARN characters when creating event
// patterns so that they match the ARN syntax in the event you want to match.
//
// In EventBridge, it is possible to create rules that lead to infinite loops,
// where a rule is fired repeatedly. For example, a rule might detect that ACLs
// have changed on an S3 bucket, and trigger software to change them to the
// desired state. If the rule is not written carefully, the subsequent change
// to the ACLs fires the rule again, creating an infinite loop.
//
// To prevent this, write the rules so that the triggered actions do not re-fire
// the same rule. For example, your rule could fire only if ACLs are found to
// be in a bad state, instead of after any change.
//
// An infinite loop can quickly cause higher than expected charges. We recommend
// that you use budgeting, which alerts you when charges exceed your specified
// limit. For more information, see Managing Your Costs with Budgets (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html).
//
//    // Example sending a request using PutRuleRequest.
//    req := client.PutRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eventbridge-2015-10-07/PutRule
func (c *Client) PutRuleRequest(input *PutRuleInput) PutRuleRequest {
	op := &aws.Operation{
		Name:       opPutRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutRuleInput{}
	}

	req := c.newRequest(op, input, &PutRuleOutput{})

	return PutRuleRequest{Request: req, Input: input, Copy: c.PutRuleRequest}
}

// PutRuleRequest is the request type for the
// PutRule API operation.
type PutRuleRequest struct {
	*aws.Request
	Input *PutRuleInput
	Copy  func(*PutRuleInput) PutRuleRequest
}

// Send marshals and sends the PutRule API request.
func (r PutRuleRequest) Send(ctx context.Context) (*PutRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRuleResponse{
		PutRuleOutput: r.Request.Data.(*PutRuleOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRuleResponse is the response type for the
// PutRule API operation.
type PutRuleResponse struct {
	*PutRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRule request.
func (r *PutRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
