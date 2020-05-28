// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package wafregionaliface provides an interface to enable mocking the AWS WAF Regional service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package wafregionaliface

import (
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

// ClientAPI provides an interface to enable mocking the
// wafregional.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // WAF Regional.
//    func myFunc(svc wafregionaliface.ClientAPI) bool {
//        // Make svc.AssociateWebACL request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := wafregional.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        wafregionaliface.ClientPI
//    }
//    func (m *mockClientClient) AssociateWebACL(input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AssociateWebACLRequest(*wafregional.AssociateWebACLInput) wafregional.AssociateWebACLRequest

	CreateByteMatchSetRequest(*wafregional.CreateByteMatchSetInput) wafregional.CreateByteMatchSetRequest

	CreateGeoMatchSetRequest(*wafregional.CreateGeoMatchSetInput) wafregional.CreateGeoMatchSetRequest

	CreateIPSetRequest(*wafregional.CreateIPSetInput) wafregional.CreateIPSetRequest

	CreateRateBasedRuleRequest(*wafregional.CreateRateBasedRuleInput) wafregional.CreateRateBasedRuleRequest

	CreateRegexMatchSetRequest(*wafregional.CreateRegexMatchSetInput) wafregional.CreateRegexMatchSetRequest

	CreateRegexPatternSetRequest(*wafregional.CreateRegexPatternSetInput) wafregional.CreateRegexPatternSetRequest

	CreateRuleRequest(*wafregional.CreateRuleInput) wafregional.CreateRuleRequest

	CreateRuleGroupRequest(*wafregional.CreateRuleGroupInput) wafregional.CreateRuleGroupRequest

	CreateSizeConstraintSetRequest(*wafregional.CreateSizeConstraintSetInput) wafregional.CreateSizeConstraintSetRequest

	CreateSqlInjectionMatchSetRequest(*wafregional.CreateSqlInjectionMatchSetInput) wafregional.CreateSqlInjectionMatchSetRequest

	CreateWebACLRequest(*wafregional.CreateWebACLInput) wafregional.CreateWebACLRequest

	CreateWebACLMigrationStackRequest(*wafregional.CreateWebACLMigrationStackInput) wafregional.CreateWebACLMigrationStackRequest

	CreateXssMatchSetRequest(*wafregional.CreateXssMatchSetInput) wafregional.CreateXssMatchSetRequest

	DeleteByteMatchSetRequest(*wafregional.DeleteByteMatchSetInput) wafregional.DeleteByteMatchSetRequest

	DeleteGeoMatchSetRequest(*wafregional.DeleteGeoMatchSetInput) wafregional.DeleteGeoMatchSetRequest

	DeleteIPSetRequest(*wafregional.DeleteIPSetInput) wafregional.DeleteIPSetRequest

	DeleteLoggingConfigurationRequest(*wafregional.DeleteLoggingConfigurationInput) wafregional.DeleteLoggingConfigurationRequest

	DeletePermissionPolicyRequest(*wafregional.DeletePermissionPolicyInput) wafregional.DeletePermissionPolicyRequest

	DeleteRateBasedRuleRequest(*wafregional.DeleteRateBasedRuleInput) wafregional.DeleteRateBasedRuleRequest

	DeleteRegexMatchSetRequest(*wafregional.DeleteRegexMatchSetInput) wafregional.DeleteRegexMatchSetRequest

	DeleteRegexPatternSetRequest(*wafregional.DeleteRegexPatternSetInput) wafregional.DeleteRegexPatternSetRequest

	DeleteRuleRequest(*wafregional.DeleteRuleInput) wafregional.DeleteRuleRequest

	DeleteRuleGroupRequest(*wafregional.DeleteRuleGroupInput) wafregional.DeleteRuleGroupRequest

	DeleteSizeConstraintSetRequest(*wafregional.DeleteSizeConstraintSetInput) wafregional.DeleteSizeConstraintSetRequest

	DeleteSqlInjectionMatchSetRequest(*wafregional.DeleteSqlInjectionMatchSetInput) wafregional.DeleteSqlInjectionMatchSetRequest

	DeleteWebACLRequest(*wafregional.DeleteWebACLInput) wafregional.DeleteWebACLRequest

	DeleteXssMatchSetRequest(*wafregional.DeleteXssMatchSetInput) wafregional.DeleteXssMatchSetRequest

	DisassociateWebACLRequest(*wafregional.DisassociateWebACLInput) wafregional.DisassociateWebACLRequest

	GetByteMatchSetRequest(*wafregional.GetByteMatchSetInput) wafregional.GetByteMatchSetRequest

	GetChangeTokenRequest(*wafregional.GetChangeTokenInput) wafregional.GetChangeTokenRequest

	GetChangeTokenStatusRequest(*wafregional.GetChangeTokenStatusInput) wafregional.GetChangeTokenStatusRequest

	GetGeoMatchSetRequest(*wafregional.GetGeoMatchSetInput) wafregional.GetGeoMatchSetRequest

	GetIPSetRequest(*wafregional.GetIPSetInput) wafregional.GetIPSetRequest

	GetLoggingConfigurationRequest(*wafregional.GetLoggingConfigurationInput) wafregional.GetLoggingConfigurationRequest

	GetPermissionPolicyRequest(*wafregional.GetPermissionPolicyInput) wafregional.GetPermissionPolicyRequest

	GetRateBasedRuleRequest(*wafregional.GetRateBasedRuleInput) wafregional.GetRateBasedRuleRequest

	GetRateBasedRuleManagedKeysRequest(*wafregional.GetRateBasedRuleManagedKeysInput) wafregional.GetRateBasedRuleManagedKeysRequest

	GetRegexMatchSetRequest(*wafregional.GetRegexMatchSetInput) wafregional.GetRegexMatchSetRequest

	GetRegexPatternSetRequest(*wafregional.GetRegexPatternSetInput) wafregional.GetRegexPatternSetRequest

	GetRuleRequest(*wafregional.GetRuleInput) wafregional.GetRuleRequest

	GetRuleGroupRequest(*wafregional.GetRuleGroupInput) wafregional.GetRuleGroupRequest

	GetSampledRequestsRequest(*wafregional.GetSampledRequestsInput) wafregional.GetSampledRequestsRequest

	GetSizeConstraintSetRequest(*wafregional.GetSizeConstraintSetInput) wafregional.GetSizeConstraintSetRequest

	GetSqlInjectionMatchSetRequest(*wafregional.GetSqlInjectionMatchSetInput) wafregional.GetSqlInjectionMatchSetRequest

	GetWebACLRequest(*wafregional.GetWebACLInput) wafregional.GetWebACLRequest

	GetWebACLForResourceRequest(*wafregional.GetWebACLForResourceInput) wafregional.GetWebACLForResourceRequest

	GetXssMatchSetRequest(*wafregional.GetXssMatchSetInput) wafregional.GetXssMatchSetRequest

	ListActivatedRulesInRuleGroupRequest(*wafregional.ListActivatedRulesInRuleGroupInput) wafregional.ListActivatedRulesInRuleGroupRequest

	ListByteMatchSetsRequest(*wafregional.ListByteMatchSetsInput) wafregional.ListByteMatchSetsRequest

	ListGeoMatchSetsRequest(*wafregional.ListGeoMatchSetsInput) wafregional.ListGeoMatchSetsRequest

	ListIPSetsRequest(*wafregional.ListIPSetsInput) wafregional.ListIPSetsRequest

	ListLoggingConfigurationsRequest(*wafregional.ListLoggingConfigurationsInput) wafregional.ListLoggingConfigurationsRequest

	ListRateBasedRulesRequest(*wafregional.ListRateBasedRulesInput) wafregional.ListRateBasedRulesRequest

	ListRegexMatchSetsRequest(*wafregional.ListRegexMatchSetsInput) wafregional.ListRegexMatchSetsRequest

	ListRegexPatternSetsRequest(*wafregional.ListRegexPatternSetsInput) wafregional.ListRegexPatternSetsRequest

	ListResourcesForWebACLRequest(*wafregional.ListResourcesForWebACLInput) wafregional.ListResourcesForWebACLRequest

	ListRuleGroupsRequest(*wafregional.ListRuleGroupsInput) wafregional.ListRuleGroupsRequest

	ListRulesRequest(*wafregional.ListRulesInput) wafregional.ListRulesRequest

	ListSizeConstraintSetsRequest(*wafregional.ListSizeConstraintSetsInput) wafregional.ListSizeConstraintSetsRequest

	ListSqlInjectionMatchSetsRequest(*wafregional.ListSqlInjectionMatchSetsInput) wafregional.ListSqlInjectionMatchSetsRequest

	ListSubscribedRuleGroupsRequest(*wafregional.ListSubscribedRuleGroupsInput) wafregional.ListSubscribedRuleGroupsRequest

	ListTagsForResourceRequest(*wafregional.ListTagsForResourceInput) wafregional.ListTagsForResourceRequest

	ListWebACLsRequest(*wafregional.ListWebACLsInput) wafregional.ListWebACLsRequest

	ListXssMatchSetsRequest(*wafregional.ListXssMatchSetsInput) wafregional.ListXssMatchSetsRequest

	PutLoggingConfigurationRequest(*wafregional.PutLoggingConfigurationInput) wafregional.PutLoggingConfigurationRequest

	PutPermissionPolicyRequest(*wafregional.PutPermissionPolicyInput) wafregional.PutPermissionPolicyRequest

	TagResourceRequest(*wafregional.TagResourceInput) wafregional.TagResourceRequest

	UntagResourceRequest(*wafregional.UntagResourceInput) wafregional.UntagResourceRequest

	UpdateByteMatchSetRequest(*wafregional.UpdateByteMatchSetInput) wafregional.UpdateByteMatchSetRequest

	UpdateGeoMatchSetRequest(*wafregional.UpdateGeoMatchSetInput) wafregional.UpdateGeoMatchSetRequest

	UpdateIPSetRequest(*wafregional.UpdateIPSetInput) wafregional.UpdateIPSetRequest

	UpdateRateBasedRuleRequest(*wafregional.UpdateRateBasedRuleInput) wafregional.UpdateRateBasedRuleRequest

	UpdateRegexMatchSetRequest(*wafregional.UpdateRegexMatchSetInput) wafregional.UpdateRegexMatchSetRequest

	UpdateRegexPatternSetRequest(*wafregional.UpdateRegexPatternSetInput) wafregional.UpdateRegexPatternSetRequest

	UpdateRuleRequest(*wafregional.UpdateRuleInput) wafregional.UpdateRuleRequest

	UpdateRuleGroupRequest(*wafregional.UpdateRuleGroupInput) wafregional.UpdateRuleGroupRequest

	UpdateSizeConstraintSetRequest(*wafregional.UpdateSizeConstraintSetInput) wafregional.UpdateSizeConstraintSetRequest

	UpdateSqlInjectionMatchSetRequest(*wafregional.UpdateSqlInjectionMatchSetInput) wafregional.UpdateSqlInjectionMatchSetRequest

	UpdateWebACLRequest(*wafregional.UpdateWebACLInput) wafregional.UpdateWebACLRequest

	UpdateXssMatchSetRequest(*wafregional.UpdateXssMatchSetInput) wafregional.UpdateXssMatchSetRequest
}

var _ ClientAPI = (*wafregional.Client)(nil)
