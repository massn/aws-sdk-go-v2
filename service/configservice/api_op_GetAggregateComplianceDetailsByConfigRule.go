// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAggregateComplianceDetailsByConfigRuleInput struct {
	_ struct{} `type:"structure"`

	// The 12-digit account ID of the source account.
	//
	// AccountId is a required field
	AccountId *string `type:"string" required:"true"`

	// The source region from where the data is aggregated.
	//
	// AwsRegion is a required field
	AwsRegion *string `min:"1" type:"string" required:"true"`

	// The resource compliance status.
	//
	// For the GetAggregateComplianceDetailsByConfigRuleRequest data type, AWS Config
	// supports only the COMPLIANT and NON_COMPLIANT. AWS Config does not support
	// the NOT_APPLICABLE and INSUFFICIENT_DATA values.
	ComplianceType ComplianceType `type:"string" enum:"true"`

	// The name of the AWS Config rule for which you want compliance information.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// The name of the configuration aggregator.
	//
	// ConfigurationAggregatorName is a required field
	ConfigurationAggregatorName *string `min:"1" type:"string" required:"true"`

	// The maximum number of evaluation results returned on each page. The default
	// is 50. You cannot specify a number greater than 100. If you specify 0, AWS
	// Config uses the default.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetAggregateComplianceDetailsByConfigRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAggregateComplianceDetailsByConfigRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAggregateComplianceDetailsByConfigRuleInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.AwsRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsRegion"))
	}
	if s.AwsRegion != nil && len(*s.AwsRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsRegion", 1))
	}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}

	if s.ConfigurationAggregatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationAggregatorName"))
	}
	if s.ConfigurationAggregatorName != nil && len(*s.ConfigurationAggregatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationAggregatorName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAggregateComplianceDetailsByConfigRuleOutput struct {
	_ struct{} `type:"structure"`

	// Returns an AggregateEvaluationResults object.
	AggregateEvaluationResults []AggregateEvaluationResult `type:"list"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetAggregateComplianceDetailsByConfigRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAggregateComplianceDetailsByConfigRule = "GetAggregateComplianceDetailsByConfigRule"

// GetAggregateComplianceDetailsByConfigRuleRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the evaluation results for the specified AWS Config rule for a specific
// resource in a rule. The results indicate which AWS resources were evaluated
// by the rule, when each resource was last evaluated, and whether each resource
// complies with the rule.
//
// The results can return an empty result page. But if you have a nextToken,
// the results are displayed on the next page.
//
//    // Example sending a request using GetAggregateComplianceDetailsByConfigRuleRequest.
//    req := client.GetAggregateComplianceDetailsByConfigRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetAggregateComplianceDetailsByConfigRule
func (c *Client) GetAggregateComplianceDetailsByConfigRuleRequest(input *GetAggregateComplianceDetailsByConfigRuleInput) GetAggregateComplianceDetailsByConfigRuleRequest {
	op := &aws.Operation{
		Name:       opGetAggregateComplianceDetailsByConfigRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAggregateComplianceDetailsByConfigRuleInput{}
	}

	req := c.newRequest(op, input, &GetAggregateComplianceDetailsByConfigRuleOutput{})

	return GetAggregateComplianceDetailsByConfigRuleRequest{Request: req, Input: input, Copy: c.GetAggregateComplianceDetailsByConfigRuleRequest}
}

// GetAggregateComplianceDetailsByConfigRuleRequest is the request type for the
// GetAggregateComplianceDetailsByConfigRule API operation.
type GetAggregateComplianceDetailsByConfigRuleRequest struct {
	*aws.Request
	Input *GetAggregateComplianceDetailsByConfigRuleInput
	Copy  func(*GetAggregateComplianceDetailsByConfigRuleInput) GetAggregateComplianceDetailsByConfigRuleRequest
}

// Send marshals and sends the GetAggregateComplianceDetailsByConfigRule API request.
func (r GetAggregateComplianceDetailsByConfigRuleRequest) Send(ctx context.Context) (*GetAggregateComplianceDetailsByConfigRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAggregateComplianceDetailsByConfigRuleResponse{
		GetAggregateComplianceDetailsByConfigRuleOutput: r.Request.Data.(*GetAggregateComplianceDetailsByConfigRuleOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAggregateComplianceDetailsByConfigRuleResponse is the response type for the
// GetAggregateComplianceDetailsByConfigRule API operation.
type GetAggregateComplianceDetailsByConfigRuleResponse struct {
	*GetAggregateComplianceDetailsByConfigRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAggregateComplianceDetailsByConfigRule request.
func (r *GetAggregateComplianceDetailsByConfigRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
