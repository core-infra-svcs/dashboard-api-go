# \RuleGroupsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride**](RuleGroupsApi.md#CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/{ruleGroupId}/override | Create a rule group override for an intrusion policy.
[**DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides**](RuleGroupsApi.md#DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/overrides/declare | Declare the desired rule group overrides for an intrusion policy.
[**DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride**](RuleGroupsApi.md#DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride) | **Delete** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overrides/{overrideId} | Delete a rule group override for an intrusion policy.
[**GetOrganizationApplianceSecurityIntrusionRuleGroups**](RuleGroupsApi.md#GetOrganizationApplianceSecurityIntrusionRuleGroups) | **Get** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups | List the rule groups that belong to a security policy.
[**GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides**](RuleGroupsApi.md#GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides) | **Get** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overrides | List the rule group overrides configured for an intrusion policy.
[**GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews**](RuleGroupsApi.md#GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews) | **Get** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overviews | List counts for the child rule groups and rules for each rule group in a security policy.
[**UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride**](RuleGroupsApi.md#UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride) | **Put** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/{ruleGroupId}/override | Update a rule group override for an intrusion policy.



## CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride

> InlineResponse200295Items CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(ctx, organizationId, policyId, ruleGroupId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()

Create a rule group override for an intrusion policy.



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
    policyId := "policyId_example" // string | Policy ID
    ruleGroupId := "ruleGroupId_example" // string | Rule group ID
    createOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride := *openapiclient.NewInlineObject266(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsRuleGroupIdOverrideOverride(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareSecurity(int32(123)))) // InlineObject266 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(context.Background(), organizationId, policyId, ruleGroupId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: InlineResponse200295Items
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 
**ruleGroupId** | **string** | Rule group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride** | [**InlineObject266**](InlineObject266.md) |  | 

### Return type

[**InlineResponse200295Items**](InlineResponse200295Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides

> InlineResponse200295 DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(ctx, organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides).Execute()

Declare the desired rule group overrides for an intrusion policy.



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
    policyId := "policyId_example" // string | Policy ID
    declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides := *openapiclient.NewInlineObject264([]openapiclient.OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems{*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareGroup("Id_example"), *openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareSecurity(int32(123)))}) // InlineObject264 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(context.Background(), organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides`: InlineResponse200295
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides** | [**InlineObject264**](InlineObject264.md) |  | 

### Return type

[**InlineResponse200295**](InlineResponse200295.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride

> DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride(ctx, organizationId, overrideId).Execute()

Delete a rule group override for an intrusion policy.



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
    overrideId := "overrideId_example" // string | Override ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride(context.Background(), organizationId, overrideId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**overrideId** | **string** | Override ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverrideRequest struct via the builder pattern


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


## GetOrganizationApplianceSecurityIntrusionRuleGroups

> InlineResponse200297 GetOrganizationApplianceSecurityIntrusionRuleGroups(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()

List the rule groups that belong to a security policy.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    policyIds := []string{"Inner_example"} // []string | Collection of base or intrusion policy identifiers to filter results by (optional)
    parentRuleGroupIds := []string{"Inner_example"} // []string | Filter results to rule groups whose parent matches any of the provided identifiers (optional)
    search := "search_example" // string | Case-insensitive text filter applied to rule group name and description (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroups(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRuleGroups`: InlineResponse200297
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRuleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **policyIds** | **[]string** | Collection of base or intrusion policy identifiers to filter results by | 
 **parentRuleGroupIds** | **[]string** | Filter results to rule groups whose parent matches any of the provided identifiers | 
 **search** | **string** | Case-insensitive text filter applied to rule group name and description | 

### Return type

[**InlineResponse200297**](InlineResponse200297.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides

> InlineResponse200298 GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()

List the rule group overrides configured for an intrusion policy.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 25. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    policyIds := []string{"Inner_example"} // []string | Collection of intrusion policy identifiers to filter the overrides by. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides`: InlineResponse200298
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRuleGroupsOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 25. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **policyIds** | **[]string** | Collection of intrusion policy identifiers to filter the overrides by. | 

### Return type

[**InlineResponse200298**](InlineResponse200298.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews

> InlineResponse200299 GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()

List counts for the child rule groups and rules for each rule group in a security policy.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    policyIds := []string{"Inner_example"} // []string | Collection of base or intrusion policy identifiers to filter results by (optional)
    parentRuleGroupIds := []string{"Inner_example"} // []string | Filter results to rule groups whose parent matches any of the provided identifiers (optional)
    search := "search_example" // string | Case-insensitive text filter applied to rule group name and description (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews`: InlineResponse200299
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRuleGroupsOverviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **policyIds** | **[]string** | Collection of base or intrusion policy identifiers to filter results by | 
 **parentRuleGroupIds** | **[]string** | Filter results to rule groups whose parent matches any of the provided identifiers | 
 **search** | **string** | Case-insensitive text filter applied to rule group name and description | 

### Return type

[**InlineResponse200299**](InlineResponse200299.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride

> InlineResponse200295Items UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(ctx, organizationId, policyId, ruleGroupId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()

Update a rule group override for an intrusion policy.



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
    policyId := "policyId_example" // string | Policy ID
    ruleGroupId := "ruleGroupId_example" // string | Rule group ID
    updateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride := *openapiclient.NewInlineObject265(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsRuleGroupIdOverrideOverride(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareSecurity(int32(123)))) // InlineObject265 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RuleGroupsApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(context.Background(), organizationId, policyId, ruleGroupId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuleGroupsApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: InlineResponse200295Items
    fmt.Fprintf(os.Stdout, "Response from `RuleGroupsApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 
**ruleGroupId** | **string** | Rule group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride** | [**InlineObject265**](InlineObject265.md) |  | 

### Return type

[**InlineResponse200295Items**](InlineResponse200295Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

