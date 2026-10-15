# \SecurityApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationApplianceSecurityIntrusionPolicy**](SecurityApi.md#CreateOrganizationApplianceSecurityIntrusionPolicy) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies | Create a new intrusion policy for the organization.
[**CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride**](SecurityApi.md#CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/{ruleGroupId}/override | Create a rule group override for an intrusion policy.
[**CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride**](SecurityApi.md#CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/{ruleId}/override | Create a rule override for an intrusion policy.
[**DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides**](SecurityApi.md#DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/overrides/declare | Declare the desired rule group overrides for an intrusion policy.
[**DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides**](SecurityApi.md#DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/overrides/declare | Declare the desired rule overrides for an intrusion policy.
[**DeleteOrganizationApplianceSecurityIntrusionPolicy**](SecurityApi.md#DeleteOrganizationApplianceSecurityIntrusionPolicy) | **Delete** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId} | Delete a single intrusion policy for the organization.
[**DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride**](SecurityApi.md#DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride) | **Delete** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overrides/{overrideId} | Delete a rule group override for an intrusion policy.
[**DeleteOrganizationApplianceSecurityIntrusionRulesOverride**](SecurityApi.md#DeleteOrganizationApplianceSecurityIntrusionRulesOverride) | **Delete** /organizations/{organizationId}/appliance/security/intrusion/rules/overrides/{overrideId} | Delete a rule override for an intrusion policy.
[**GetNetworkApplianceClientSecurityEvents**](SecurityApi.md#GetNetworkApplianceClientSecurityEvents) | **Get** /networks/{networkId}/appliance/clients/{clientId}/security/events | List the security events for a client
[**GetNetworkApplianceSecurityEvents**](SecurityApi.md#GetNetworkApplianceSecurityEvents) | **Get** /networks/{networkId}/appliance/security/events | List the security events for a network
[**GetNetworkApplianceSecurityIntrusion**](SecurityApi.md#GetNetworkApplianceSecurityIntrusion) | **Get** /networks/{networkId}/appliance/security/intrusion | Returns all supported intrusion settings for an MX network
[**GetNetworkApplianceSecurityMalware**](SecurityApi.md#GetNetworkApplianceSecurityMalware) | **Get** /networks/{networkId}/appliance/security/malware | Returns all supported malware settings for an MX network
[**GetOrganizationApplianceSecurityEvents**](SecurityApi.md#GetOrganizationApplianceSecurityEvents) | **Get** /organizations/{organizationId}/appliance/security/events | List the security events for an organization
[**GetOrganizationApplianceSecurityIntrusion**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusion) | **Get** /organizations/{organizationId}/appliance/security/intrusion | Returns all supported intrusion settings for an organization
[**GetOrganizationApplianceSecurityIntrusionPolicies**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusionPolicies) | **Get** /organizations/{organizationId}/appliance/security/intrusion/policies | List the intrusion policies configured for an organization along with base policies.
[**GetOrganizationApplianceSecurityIntrusionPoliciesOverviews**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusionPoliciesOverviews) | **Get** /organizations/{organizationId}/appliance/security/intrusion/policies/overviews | List counts for the intrusion and base policies configured for an organization.
[**GetOrganizationApplianceSecurityIntrusionRuleGroups**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusionRuleGroups) | **Get** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups | List the rule groups that belong to a security policy.
[**GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides) | **Get** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overrides | List the rule group overrides configured for an intrusion policy.
[**GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews) | **Get** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overviews | List counts for the child rule groups and rules for each rule group in a security policy.
[**GetOrganizationApplianceSecurityIntrusionRules**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusionRules) | **Get** /organizations/{organizationId}/appliance/security/intrusion/rules | List the rules that belong to a security policy.
[**GetOrganizationApplianceSecurityIntrusionRulesOverrides**](SecurityApi.md#GetOrganizationApplianceSecurityIntrusionRulesOverrides) | **Get** /organizations/{organizationId}/appliance/security/intrusion/rules/overrides | List the rule overrides configured for an intrusion policy.
[**UpdateNetworkApplianceSecurityIntrusion**](SecurityApi.md#UpdateNetworkApplianceSecurityIntrusion) | **Put** /networks/{networkId}/appliance/security/intrusion | Set the supported intrusion settings for an MX network
[**UpdateNetworkApplianceSecurityMalware**](SecurityApi.md#UpdateNetworkApplianceSecurityMalware) | **Put** /networks/{networkId}/appliance/security/malware | Set the supported malware settings for an MX network
[**UpdateOrganizationApplianceSecurityIntrusion**](SecurityApi.md#UpdateOrganizationApplianceSecurityIntrusion) | **Put** /organizations/{organizationId}/appliance/security/intrusion | Sets supported intrusion settings for an organization
[**UpdateOrganizationApplianceSecurityIntrusionPolicy**](SecurityApi.md#UpdateOrganizationApplianceSecurityIntrusionPolicy) | **Put** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId} | Update a single intrusion policy for the organization.
[**UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride**](SecurityApi.md#UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride) | **Put** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/{ruleGroupId}/override | Update a rule group override for an intrusion policy.
[**UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride**](SecurityApi.md#UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride) | **Put** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/{ruleId}/override | Update a rule override for an intrusion policy.



## CreateOrganizationApplianceSecurityIntrusionPolicy

> InlineResponse200293Items CreateOrganizationApplianceSecurityIntrusionPolicy(ctx, organizationId).CreateOrganizationApplianceSecurityIntrusionPolicy(createOrganizationApplianceSecurityIntrusionPolicy).Execute()

Create a new intrusion policy for the organization.



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
    createOrganizationApplianceSecurityIntrusionPolicy := *openapiclient.NewInlineObject262() // InlineObject262 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicy(context.Background(), organizationId).CreateOrganizationApplianceSecurityIntrusionPolicy(createOrganizationApplianceSecurityIntrusionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceSecurityIntrusionPolicy`: InlineResponse200293Items
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceSecurityIntrusionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceSecurityIntrusionPolicy** | [**InlineObject262**](InlineObject262.md) |  | 

### Return type

[**InlineResponse200293Items**](InlineResponse200293Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(context.Background(), organizationId, policyId, ruleGroupId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: InlineResponse200295Items
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: %v\n", resp)
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
    resp, r, err := apiClient.SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(context.Background(), organizationId, policyId, ruleId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: InlineResponse200296Items
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: %v\n", resp)
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
    resp, r, err := apiClient.SecurityApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(context.Background(), organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides`: InlineResponse200295
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides`: %v\n", resp)
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


## DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides

> InlineResponse200296 DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(ctx, organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides).Execute()

Declare the desired rule overrides for an intrusion policy.



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
    declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides := *openapiclient.NewInlineObject267([]openapiclient.OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems{*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareRule("Id_example"), "Action_example")}) // InlineObject267 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(context.Background(), organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides`: InlineResponse200296
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides** | [**InlineObject267**](InlineObject267.md) |  | 

### Return type

[**InlineResponse200296**](InlineResponse200296.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApplianceSecurityIntrusionPolicy

> DeleteOrganizationApplianceSecurityIntrusionPolicy(ctx, organizationId, policyId).Execute()

Delete a single intrusion policy for the organization.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.DeleteOrganizationApplianceSecurityIntrusionPolicy(context.Background(), organizationId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.DeleteOrganizationApplianceSecurityIntrusionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplianceSecurityIntrusionPolicyRequest struct via the builder pattern


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
    resp, r, err := apiClient.SecurityApi.DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride(context.Background(), organizationId, overrideId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride``: %v\n", err)
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


## DeleteOrganizationApplianceSecurityIntrusionRulesOverride

> DeleteOrganizationApplianceSecurityIntrusionRulesOverride(ctx, organizationId, overrideId).Execute()

Delete a rule override for an intrusion policy.



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
    resp, r, err := apiClient.SecurityApi.DeleteOrganizationApplianceSecurityIntrusionRulesOverride(context.Background(), organizationId, overrideId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.DeleteOrganizationApplianceSecurityIntrusionRulesOverride``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationApplianceSecurityIntrusionRulesOverrideRequest struct via the builder pattern


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


## GetNetworkApplianceClientSecurityEvents

> []map[string]interface{} GetNetworkApplianceClientSecurityEvents(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a client



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
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetNetworkApplianceClientSecurityEvents(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetNetworkApplianceClientSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceClientSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetNetworkApplianceClientSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceClientSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityEvents

> []map[string]interface{} GetNetworkApplianceSecurityEvents(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a network



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
    networkId := "networkId_example" // string | Network ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetNetworkApplianceSecurityEvents(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetNetworkApplianceSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetNetworkApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityIntrusion

> InlineResponse20089 GetNetworkApplianceSecurityIntrusion(ctx, networkId).Execute()

Returns all supported intrusion settings for an MX network



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetNetworkApplianceSecurityIntrusion(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetNetworkApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityIntrusion`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityMalware

> InlineResponse20091 GetNetworkApplianceSecurityMalware(ctx, networkId).Execute()

Returns all supported malware settings for an MX network



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetNetworkApplianceSecurityMalware(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetNetworkApplianceSecurityMalware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSecurityMalware`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityEvents

> []map[string]interface{} GetOrganizationApplianceSecurityEvents(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for an organization



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityEvents(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusion

> InlineResponse200292 GetOrganizationApplianceSecurityIntrusion(ctx, organizationId).Execute()

Returns all supported intrusion settings for an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusion`: InlineResponse200292
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200292**](InlineResponse200292.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusionPolicies

> InlineResponse200293 GetOrganizationApplianceSecurityIntrusionPolicies(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Mode(mode).PolicyIds(policyIds).Search(search).Execute()

List the intrusion policies configured for an organization along with base policies.



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
    mode := "mode_example" // string | Controls which policy set is returned. (optional)
    policyIds := []string{"Inner_example"} // []string | Identifiers of policies to filter (optional)
    search := "search_example" // string | Filter policies by case-insensitive partial match on name or description. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusionPolicies(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Mode(mode).PolicyIds(policyIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionPolicies`: InlineResponse200293
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusionPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 25. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **mode** | **string** | Controls which policy set is returned. | 
 **policyIds** | **[]string** | Identifiers of policies to filter | 
 **search** | **string** | Filter policies by case-insensitive partial match on name or description. | 

### Return type

[**InlineResponse200293**](InlineResponse200293.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusionPoliciesOverviews

> InlineResponse200294 GetOrganizationApplianceSecurityIntrusionPoliciesOverviews(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Search(search).Execute()

List counts for the intrusion and base policies configured for an organization.



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
    policyIds := []string{"Inner_example"} // []string | Identifiers of policies to filter (optional)
    search := "search_example" // string | Filter policy overviews by case-insensitive partial match on policy name or description. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusionPoliciesOverviews(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusionPoliciesOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionPoliciesOverviews`: InlineResponse200294
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusionPoliciesOverviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionPoliciesOverviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 25. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **policyIds** | **[]string** | Identifiers of policies to filter | 
 **search** | **string** | Filter policy overviews by case-insensitive partial match on policy name or description. | 

### Return type

[**InlineResponse200294**](InlineResponse200294.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroups(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRuleGroups`: InlineResponse200297
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroups`: %v\n", resp)
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
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides`: InlineResponse200298
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides`: %v\n", resp)
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
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews`: InlineResponse200299
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverviews`: %v\n", resp)
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


## GetOrganizationApplianceSecurityIntrusionRules

> InlineResponse200300 GetOrganizationApplianceSecurityIntrusionRules(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()

List the rules that belong to a security policy.



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
    policyIds := []string{"Inner_example"} // []string | Identifiers of the base or intrusion policies to query (optional)
    parentRuleGroupIds := []string{"Inner_example"} // []string | Filter results to rules that belong to any of the specified rule groups (optional)
    search := "search_example" // string | Case-insensitive text filter applied to rule name and description (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusionRules(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRules`: InlineResponse200300
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusionRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **policyIds** | **[]string** | Identifiers of the base or intrusion policies to query | 
 **parentRuleGroupIds** | **[]string** | Filter results to rules that belong to any of the specified rule groups | 
 **search** | **string** | Case-insensitive text filter applied to rule name and description | 

### Return type

[**InlineResponse200300**](InlineResponse200300.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusionRulesOverrides

> InlineResponse200301 GetOrganizationApplianceSecurityIntrusionRulesOverrides(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()

List the rule overrides configured for an intrusion policy.



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
    policyIds := []string{"Inner_example"} // []string | Identifiers of intrusion policies to filter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRulesOverrides`: InlineResponse200301
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRulesOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 25. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **policyIds** | **[]string** | Identifiers of intrusion policies to filter | 

### Return type

[**InlineResponse200301**](InlineResponse200301.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityIntrusion

> InlineResponse20090 UpdateNetworkApplianceSecurityIntrusion(ctx, networkId).UpdateNetworkApplianceSecurityIntrusion(updateNetworkApplianceSecurityIntrusion).Execute()

Set the supported intrusion settings for an MX network



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceSecurityIntrusion := *openapiclient.NewInlineObject74() // InlineObject74 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.UpdateNetworkApplianceSecurityIntrusion(context.Background(), networkId).UpdateNetworkApplianceSecurityIntrusion(updateNetworkApplianceSecurityIntrusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.UpdateNetworkApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSecurityIntrusion`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.UpdateNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityIntrusion** | [**InlineObject74**](InlineObject74.md) |  | 

### Return type

[**InlineResponse20090**](InlineResponse20090.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityMalware

> InlineResponse20091 UpdateNetworkApplianceSecurityMalware(ctx, networkId).UpdateNetworkApplianceSecurityMalware(updateNetworkApplianceSecurityMalware).Execute()

Set the supported malware settings for an MX network



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkApplianceSecurityMalware := *openapiclient.NewInlineObject75("Mode_example") // InlineObject75 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.UpdateNetworkApplianceSecurityMalware(context.Background(), networkId).UpdateNetworkApplianceSecurityMalware(updateNetworkApplianceSecurityMalware).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.UpdateNetworkApplianceSecurityMalware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSecurityMalware`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.UpdateNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityMalware** | [**InlineObject75**](InlineObject75.md) |  | 

### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceSecurityIntrusion

> InlineResponse200292 UpdateOrganizationApplianceSecurityIntrusion(ctx, organizationId).UpdateOrganizationApplianceSecurityIntrusion(updateOrganizationApplianceSecurityIntrusion).Execute()

Sets supported intrusion settings for an organization



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
    updateOrganizationApplianceSecurityIntrusion := *openapiclient.NewInlineObject261([]openapiclient.OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules{*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules("RuleId_example")}) // InlineObject261 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.UpdateOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).UpdateOrganizationApplianceSecurityIntrusion(updateOrganizationApplianceSecurityIntrusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.UpdateOrganizationApplianceSecurityIntrusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusion`: InlineResponse200292
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.UpdateOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceSecurityIntrusion** | [**InlineObject261**](InlineObject261.md) |  | 

### Return type

[**InlineResponse200292**](InlineResponse200292.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceSecurityIntrusionPolicy

> InlineResponse200293Items UpdateOrganizationApplianceSecurityIntrusionPolicy(ctx, organizationId, policyId).UpdateOrganizationApplianceSecurityIntrusionPolicy(updateOrganizationApplianceSecurityIntrusionPolicy).Execute()

Update a single intrusion policy for the organization.



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
    updateOrganizationApplianceSecurityIntrusionPolicy := *openapiclient.NewInlineObject263(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdPolicy()) // InlineObject263 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicy(context.Background(), organizationId, policyId).UpdateOrganizationApplianceSecurityIntrusionPolicy(updateOrganizationApplianceSecurityIntrusionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusionPolicy`: InlineResponse200293Items
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationApplianceSecurityIntrusionPolicy** | [**InlineObject263**](InlineObject263.md) |  | 

### Return type

[**InlineResponse200293Items**](InlineResponse200293Items.md)

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
    resp, r, err := apiClient.SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(context.Background(), organizationId, policyId, ruleGroupId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: InlineResponse200295Items
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleGroupOverride`: %v\n", resp)
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
    resp, r, err := apiClient.SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(context.Background(), organizationId, policyId, ruleId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: InlineResponse200296Items
    fmt.Fprintf(os.Stdout, "Response from `SecurityApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: %v\n", resp)
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

