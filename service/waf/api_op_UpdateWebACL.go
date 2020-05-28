// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateWebACLInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A default action for the web ACL, either ALLOW or BLOCK. AWS WAF performs
	// the default action if a request doesn't match the criteria in any of the
	// rules in a web ACL.
	DefaultAction *WafAction `type:"structure"`

	// An array of updates to make to the WebACL.
	//
	// An array of WebACLUpdate objects that you want to insert into or delete from
	// a WebACL. For more information, see the applicable data types:
	//
	//    * WebACLUpdate: Contains Action and ActivatedRule
	//
	//    * ActivatedRule: Contains Action, OverrideAction, Priority, RuleId, and
	//    Type. ActivatedRule|OverrideAction applies only when updating or adding
	//    a RuleGroup to a WebACL. In this case, you do not use ActivatedRule|Action.
	//    For all other update requests, ActivatedRule|Action is used instead of
	//    ActivatedRule|OverrideAction.
	//
	//    * WafAction: Contains Type
	Updates []WebACLUpdate `type:"list"`

	// The WebACLId of the WebACL that you want to update. WebACLId is returned
	// by CreateWebACL and by ListWebACLs.
	//
	// WebACLId is a required field
	WebACLId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateWebACLInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.WebACLId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLId"))
	}
	if s.WebACLId != nil && len(*s.WebACLId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLId", 1))
	}
	if s.DefaultAction != nil {
		if err := s.DefaultAction.Validate(); err != nil {
			invalidParams.AddNested("DefaultAction", err.(aws.ErrInvalidParams))
		}
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateWebACLOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateWebACL request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateWebACLOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateWebACL = "UpdateWebACL"

// UpdateWebACLRequest returns a request value for making API operation for
// AWS WAF.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Inserts or deletes ActivatedRule objects in a WebACL. Each Rule identifies
// web requests that you want to allow, block, or count. When you update a WebACL,
// you specify the following values:
//
//    * A default action for the WebACL, either ALLOW or BLOCK. AWS WAF performs
//    the default action if a request doesn't match the criteria in any of the
//    Rules in a WebACL.
//
//    * The Rules that you want to add or delete. If you want to replace one
//    Rule with another, you delete the existing Rule and add the new one.
//
//    * For each Rule, whether you want AWS WAF to allow requests, block requests,
//    or count requests that match the conditions in the Rule.
//
//    * The order in which you want AWS WAF to evaluate the Rules in a WebACL.
//    If you add more than one Rule to a WebACL, AWS WAF evaluates each request
//    against the Rules in order based on the value of Priority. (The Rule that
//    has the lowest value for Priority is evaluated first.) When a web request
//    matches all the predicates (such as ByteMatchSets and IPSets) in a Rule,
//    AWS WAF immediately takes the corresponding action, allow or block, and
//    doesn't evaluate the request against the remaining Rules in the WebACL,
//    if any.
//
// To create and configure a WebACL, perform the following steps:
//
// Create and update the predicates that you want to include in Rules. For more
// information, see CreateByteMatchSet, UpdateByteMatchSet, CreateIPSet, UpdateIPSet,
// CreateSqlInjectionMatchSet, and UpdateSqlInjectionMatchSet.
//
// Create and update the Rules that you want to include in the WebACL. For more
// information, see CreateRule and UpdateRule.
//
// Create a WebACL. See CreateWebACL.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateWebACL request.
//
// Submit an UpdateWebACL request to specify the Rules that you want to include
// in the WebACL, to specify the default action, and to associate the WebACL
// with a CloudFront distribution.
//
// The ActivatedRule can be a rule group. If you specify a rule group as your
// ActivatedRule , you can exclude specific rules from that rule group.
//
// If you already have a rule group associated with a web ACL and want to submit
// an UpdateWebACL request to exclude certain rules from that rule group, you
// must first remove the rule group from the web ACL, the re-insert it again,
// specifying the excluded rules. For details, see ActivatedRule$ExcludedRules .
//
// Be aware that if you try to add a RATE_BASED rule to a web ACL without setting
// the rule type when first creating the rule, the UpdateWebACL request will
// fail because the request tries to add a REGULAR rule (the default rule type)
// with the specified ID, which does not exist.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateWebACLRequest.
//    req := client.UpdateWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateWebACL
func (c *Client) UpdateWebACLRequest(input *UpdateWebACLInput) UpdateWebACLRequest {
	op := &aws.Operation{
		Name:       opUpdateWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateWebACLInput{}
	}

	req := c.newRequest(op, input, &UpdateWebACLOutput{})

	return UpdateWebACLRequest{Request: req, Input: input, Copy: c.UpdateWebACLRequest}
}

// UpdateWebACLRequest is the request type for the
// UpdateWebACL API operation.
type UpdateWebACLRequest struct {
	*aws.Request
	Input *UpdateWebACLInput
	Copy  func(*UpdateWebACLInput) UpdateWebACLRequest
}

// Send marshals and sends the UpdateWebACL API request.
func (r UpdateWebACLRequest) Send(ctx context.Context) (*UpdateWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateWebACLResponse{
		UpdateWebACLOutput: r.Request.Data.(*UpdateWebACLOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateWebACLResponse is the response type for the
// UpdateWebACL API operation.
type UpdateWebACLResponse struct {
	*UpdateWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateWebACL request.
func (r *UpdateWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
