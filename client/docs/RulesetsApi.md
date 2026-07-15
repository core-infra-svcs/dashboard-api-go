# \RulesetsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationPoliciesGlobalFirewallRuleset**](RulesetsApi.md#CreateOrganizationPoliciesGlobalFirewallRuleset) | **Post** /organizations/{organizationId}/policies/global/firewall/rulesets | Create an Organization-Wide Policy Firewall Ruleset
[**CreateOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesetsApi.md#CreateOrganizationPoliciesGlobalFirewallRulesetsRule) | **Post** /organizations/{organizationId}/policies/global/firewall/rulesets/rules | Create an Organization-Wide Policy Firewall Rule
[**CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](RulesetsApi.md#CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Post** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments | Create an Organization-Wide Policy Ruleset Assignment
[**DeleteOrganizationPoliciesGlobalFirewallRuleset**](RulesetsApi.md#DeleteOrganizationPoliciesGlobalFirewallRuleset) | **Delete** /organizations/{organizationId}/policies/global/firewall/rulesets/{rulesetId} | Delete an Organization-Wide Policy Firewall Ruleset
[**DeleteOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesetsApi.md#DeleteOrganizationPoliciesGlobalFirewallRulesetsRule) | **Delete** /organizations/{organizationId}/policies/global/firewall/rulesets/rules/{ruleId} | Delete an Organization-Wide Policy Firewall Rule
[**DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](RulesetsApi.md#DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Delete** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments/{assignmentId} | Delete an Organization-Wide Policy Ruleset Assignment
[**GetOrganizationPoliciesGlobalFirewallRulesets**](RulesetsApi.md#GetOrganizationPoliciesGlobalFirewallRulesets) | **Get** /organizations/{organizationId}/policies/global/firewall/rulesets | List Organization-Wide Policy Firewall Rulesets
[**GetOrganizationPoliciesGlobalFirewallRulesetsRules**](RulesetsApi.md#GetOrganizationPoliciesGlobalFirewallRulesetsRules) | **Get** /organizations/{organizationId}/policies/global/firewall/rulesets/rules | List Organization-Wide Policy Firewall Rules
[**GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments**](RulesetsApi.md#GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments | List Organization-Wide Policy Ruleset Assignments
[**UpdateOrganizationPoliciesGlobalFirewallRuleset**](RulesetsApi.md#UpdateOrganizationPoliciesGlobalFirewallRuleset) | **Put** /organizations/{organizationId}/policies/global/firewall/rulesets/{rulesetId} | Update an Organization-Wide Policy Firewall Ruleset
[**UpdateOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesetsApi.md#UpdateOrganizationPoliciesGlobalFirewallRulesetsRule) | **Put** /organizations/{organizationId}/policies/global/firewall/rulesets/rules/{ruleId} | Update an Organization-Wide Policy Firewall Rule
[**UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](RulesetsApi.md#UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Put** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments/{assignmentId} | Update an Organization-Wide Policy Ruleset Assignment



## CreateOrganizationPoliciesGlobalFirewallRuleset

> InlineResponse200388Items CreateOrganizationPoliciesGlobalFirewallRuleset(ctx, organizationId).CreateOrganizationPoliciesGlobalFirewallRuleset(createOrganizationPoliciesGlobalFirewallRuleset).Execute()

Create an Organization-Wide Policy Firewall Ruleset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationPoliciesGlobalFirewallRuleset := *openapiclient.NewInlineObject313("Name_example") // InlineObject313 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.CreateOrganizationPoliciesGlobalFirewallRuleset(context.Background(), organizationId).CreateOrganizationPoliciesGlobalFirewallRuleset(createOrganizationPoliciesGlobalFirewallRuleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.CreateOrganizationPoliciesGlobalFirewallRuleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalFirewallRuleset`: InlineResponse200388Items
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.CreateOrganizationPoliciesGlobalFirewallRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPoliciesGlobalFirewallRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPoliciesGlobalFirewallRuleset** | [**InlineObject313**](InlineObject313.md) |  | 

### Return type

[**InlineResponse200388Items**](InlineResponse200388Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPoliciesGlobalFirewallRulesetsRule

> InlineResponse200389Items CreateOrganizationPoliciesGlobalFirewallRulesetsRule(ctx, organizationId).CreateOrganizationPoliciesGlobalFirewallRulesetsRule(createOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()

Create an Organization-Wide Policy Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationPoliciesGlobalFirewallRulesetsRule := *openapiclient.NewInlineObject314("Name_example", "RulesetId_example", "Policy_example", *openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources([]string{"MatchCriteria_example"}), *openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations([]string{"MatchCriteria_example"})) // InlineObject314 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId).CreateOrganizationPoliciesGlobalFirewallRulesetsRule(createOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalFirewallRulesetsRule`: InlineResponse200389Items
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPoliciesGlobalFirewallRulesetsRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPoliciesGlobalFirewallRulesetsRule** | [**InlineObject314**](InlineObject314.md) |  | 

### Return type

[**InlineResponse200389Items**](InlineResponse200389Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> InlineResponse200395Items CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId).CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()

Create an Organization-Wide Policy Ruleset Assignment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment := *openapiclient.NewInlineObject322("RulesetId_example", "PolicyId_example") // InlineObject322 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId).CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: InlineResponse200395Items
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment** | [**InlineObject322**](InlineObject322.md) |  | 

### Return type

[**InlineResponse200395Items**](InlineResponse200395Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPoliciesGlobalFirewallRuleset

> DeleteOrganizationPoliciesGlobalFirewallRuleset(ctx, organizationId, rulesetId).Execute()

Delete an Organization-Wide Policy Firewall Ruleset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    rulesetId := "rulesetId_example" // string | Ruleset ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.DeleteOrganizationPoliciesGlobalFirewallRuleset(context.Background(), organizationId, rulesetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.DeleteOrganizationPoliciesGlobalFirewallRuleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**rulesetId** | **string** | Ruleset ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPoliciesGlobalFirewallRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPoliciesGlobalFirewallRulesetsRule

> DeleteOrganizationPoliciesGlobalFirewallRulesetsRule(ctx, organizationId, ruleId).Execute()

Delete an Organization-Wide Policy Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    ruleId := "ruleId_example" // string | Rule ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.DeleteOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.DeleteOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPoliciesGlobalFirewallRulesetsRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId, assignmentId).Execute()

Delete an Organization-Wide Policy Ruleset Assignment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    assignmentId := "assignmentId_example" // string | Assignment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId, assignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**assignmentId** | **string** | Assignment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalFirewallRulesets

> InlineResponse200388 GetOrganizationPoliciesGlobalFirewallRulesets(ctx, organizationId).RulesetIds(rulesetIds).Name(name).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Organization-Wide Policy Firewall Rulesets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    rulesetIds := []string{"Inner_example"} // []string | Filter rulesets by IDs (optional)
    name := "name_example" // string | Filter rulesets by name (partial match, case-insensitive). If multiple instances are provided, only the last one is used. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.GetOrganizationPoliciesGlobalFirewallRulesets(context.Background(), organizationId).RulesetIds(rulesetIds).Name(name).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.GetOrganizationPoliciesGlobalFirewallRulesets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalFirewallRulesets`: InlineResponse200388
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.GetOrganizationPoliciesGlobalFirewallRulesets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalFirewallRulesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rulesetIds** | **[]string** | Filter rulesets by IDs | 
 **name** | **string** | Filter rulesets by name (partial match, case-insensitive). If multiple instances are provided, only the last one is used. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200388**](InlineResponse200388.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalFirewallRulesetsRules

> InlineResponse200389 GetOrganizationPoliciesGlobalFirewallRulesetsRules(ctx, organizationId).RulesetIds(rulesetIds).RuleIds(ruleIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Organization-Wide Policy Firewall Rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    rulesetIds := []string{"Inner_example"} // []string | Filter rules by firewall ruleset IDs (optional)
    ruleIds := []string{"Inner_example"} // []string | Filter rules by rule IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.GetOrganizationPoliciesGlobalFirewallRulesetsRules(context.Background(), organizationId).RulesetIds(rulesetIds).RuleIds(ruleIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.GetOrganizationPoliciesGlobalFirewallRulesetsRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalFirewallRulesetsRules`: InlineResponse200389
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.GetOrganizationPoliciesGlobalFirewallRulesetsRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalFirewallRulesetsRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rulesetIds** | **[]string** | Filter rules by firewall ruleset IDs | 
 **ruleIds** | **[]string** | Filter rules by rule IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200389**](InlineResponse200389.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments

> InlineResponse200395 GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments(ctx, organizationId).RulesetIds(rulesetIds).PolicyIds(policyIds).AssignmentIds(assignmentIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Organization-Wide Policy Ruleset Assignments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    rulesetIds := []string{"Inner_example"} // []string | Filter assignments by ruleset IDs (optional)
    policyIds := []string{"Inner_example"} // []string | Filter assignments by policy IDs (optional)
    assignmentIds := []string{"Inner_example"} // []string | Filter assignments by assignment IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments(context.Background(), organizationId).RulesetIds(rulesetIds).PolicyIds(policyIds).AssignmentIds(assignmentIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments`: InlineResponse200395
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rulesetIds** | **[]string** | Filter assignments by ruleset IDs | 
 **policyIds** | **[]string** | Filter assignments by policy IDs | 
 **assignmentIds** | **[]string** | Filter assignments by assignment IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200395**](InlineResponse200395.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPoliciesGlobalFirewallRuleset

> InlineResponse200388Items UpdateOrganizationPoliciesGlobalFirewallRuleset(ctx, organizationId, rulesetId).UpdateOrganizationPoliciesGlobalFirewallRuleset(updateOrganizationPoliciesGlobalFirewallRuleset).Execute()

Update an Organization-Wide Policy Firewall Ruleset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    rulesetId := "rulesetId_example" // string | Ruleset ID
    updateOrganizationPoliciesGlobalFirewallRuleset := *openapiclient.NewInlineObject316() // InlineObject316 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.UpdateOrganizationPoliciesGlobalFirewallRuleset(context.Background(), organizationId, rulesetId).UpdateOrganizationPoliciesGlobalFirewallRuleset(updateOrganizationPoliciesGlobalFirewallRuleset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.UpdateOrganizationPoliciesGlobalFirewallRuleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalFirewallRuleset`: InlineResponse200388Items
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.UpdateOrganizationPoliciesGlobalFirewallRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**rulesetId** | **string** | Ruleset ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPoliciesGlobalFirewallRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPoliciesGlobalFirewallRuleset** | [**InlineObject316**](InlineObject316.md) |  | 

### Return type

[**InlineResponse200388Items**](InlineResponse200388Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPoliciesGlobalFirewallRulesetsRule

> InlineResponse200389Items UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(ctx, organizationId, ruleId).UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(updateOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()

Update an Organization-Wide Policy Firewall Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    ruleId := "ruleId_example" // string | Rule ID
    updateOrganizationPoliciesGlobalFirewallRulesetsRule := *openapiclient.NewInlineObject315() // InlineObject315 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId, ruleId).UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(updateOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalFirewallRulesetsRule`: InlineResponse200389Items
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPoliciesGlobalFirewallRulesetsRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPoliciesGlobalFirewallRulesetsRule** | [**InlineObject315**](InlineObject315.md) |  | 

### Return type

[**InlineResponse200389Items**](InlineResponse200389Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> InlineResponse200395Items UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId, assignmentId).UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()

Update an Organization-Wide Policy Ruleset Assignment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    assignmentId := "assignmentId_example" // string | Assignment ID
    updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment := *openapiclient.NewInlineObject323() // InlineObject323 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesetsApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId, assignmentId).UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesetsApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: InlineResponse200395Items
    fmt.Fprintf(os.Stdout, "Response from `RulesetsApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**assignmentId** | **string** | Assignment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment** | [**InlineObject323**](InlineObject323.md) |  | 

### Return type

[**InlineResponse200395Items**](InlineResponse200395Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

