# \OverrideApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride**](OverrideApi.md#CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/{ruleGroupId}/override | Create a rule group override for an intrusion policy.
[**CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride**](OverrideApi.md#CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/{ruleId}/override | Create a rule override for an intrusion policy.
[**UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride**](OverrideApi.md#UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride) | **Put** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/{ruleGroupId}/override | Update a rule group override for an intrusion policy.
[**UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride**](OverrideApi.md#UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride) | **Put** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/{ruleId}/override | Update a rule override for an intrusion policy.



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
    resp, r, err := apiClient.OverrideApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(context.Background(), organizationId, policyId, ruleGroupId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverrideApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: InlineResponse200295Items
    fmt.Fprintf(os.Stdout, "Response from `OverrideApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: %v\n", resp)
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


## CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride

> InlineResponse200296Items CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(ctx, organizationId, policyId, ruleId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()

Create a rule override for an intrusion policy.



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
    ruleId := "ruleId_example" // string | Rule ID
    createOrganizationApplianceSecurityIntrusionPolicyRuleOverride := *openapiclient.NewInlineObject269(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesRuleIdOverrideOverride1("Action_example")) // InlineObject269 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverrideApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(context.Background(), organizationId, policyId, ruleId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverrideApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: InlineResponse200296Items
    fmt.Fprintf(os.Stdout, "Response from `OverrideApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceSecurityIntrusionPolicyRuleOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createOrganizationApplianceSecurityIntrusionPolicyRuleOverride** | [**InlineObject269**](InlineObject269.md) |  | 

### Return type

[**InlineResponse200296Items**](InlineResponse200296Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := apiClient.OverrideApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(context.Background(), organizationId, policyId, ruleGroupId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverrideApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: InlineResponse200295Items
    fmt.Fprintf(os.Stdout, "Response from `OverrideApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: %v\n", resp)
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


## UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride

> InlineResponse200296Items UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(ctx, organizationId, policyId, ruleId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()

Update a rule override for an intrusion policy.



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
    ruleId := "ruleId_example" // string | Rule ID
    updateOrganizationApplianceSecurityIntrusionPolicyRuleOverride := *openapiclient.NewInlineObject268(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesRuleIdOverrideOverride("Action_example")) // InlineObject268 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverrideApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(context.Background(), organizationId, policyId, ruleId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverrideApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: InlineResponse200296Items
    fmt.Fprintf(os.Stdout, "Response from `OverrideApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateOrganizationApplianceSecurityIntrusionPolicyRuleOverride** | [**InlineObject268**](InlineObject268.md) |  | 

### Return type

[**InlineResponse200296Items**](InlineResponse200296Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

