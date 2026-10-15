# \RulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWirelessAirMarshalRule**](RulesApi.md#CreateNetworkWirelessAirMarshalRule) | **Post** /networks/{networkId}/wireless/airMarshal/rules | Creates a new rule
[**CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride**](RulesApi.md#CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/{ruleId}/override | Create a rule override for an intrusion policy.
[**CreateOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesApi.md#CreateOrganizationPoliciesGlobalFirewallRulesetsRule) | **Post** /organizations/{organizationId}/policies/global/firewall/rulesets/rules | Create an Organization-Wide Policy Firewall Rule
[**DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides**](RulesApi.md#DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/overrides/declare | Declare the desired rule overrides for an intrusion policy.
[**DeleteNetworkWirelessAirMarshalRule**](RulesApi.md#DeleteNetworkWirelessAirMarshalRule) | **Delete** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Delete an Air Marshal rule.
[**DeleteOrganizationApplianceSecurityIntrusionRulesOverride**](RulesApi.md#DeleteOrganizationApplianceSecurityIntrusionRulesOverride) | **Delete** /organizations/{organizationId}/appliance/security/intrusion/rules/overrides/{overrideId} | Delete a rule override for an intrusion policy.
[**DeleteOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesApi.md#DeleteOrganizationPoliciesGlobalFirewallRulesetsRule) | **Delete** /organizations/{organizationId}/policies/global/firewall/rulesets/rules/{ruleId} | Delete an Organization-Wide Policy Firewall Rule
[**GetNetworkApplianceTrafficShapingRules**](RulesApi.md#GetNetworkApplianceTrafficShapingRules) | **Get** /networks/{networkId}/appliance/trafficShaping/rules | Display the traffic shaping settings rules for a Security Appliance or Secure Router network
[**GetNetworkWirelessSsidTrafficShapingRules**](RulesApi.md#GetNetworkWirelessSsidTrafficShapingRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Display the traffic shaping settings for a SSID on an MR network
[**GetOrganizationApplianceSecurityIntrusionRules**](RulesApi.md#GetOrganizationApplianceSecurityIntrusionRules) | **Get** /organizations/{organizationId}/appliance/security/intrusion/rules | List the rules that belong to a security policy.
[**GetOrganizationApplianceSecurityIntrusionRulesOverrides**](RulesApi.md#GetOrganizationApplianceSecurityIntrusionRulesOverrides) | **Get** /organizations/{organizationId}/appliance/security/intrusion/rules/overrides | List the rule overrides configured for an intrusion policy.
[**GetOrganizationPoliciesGlobalFirewallRulesetsRules**](RulesApi.md#GetOrganizationPoliciesGlobalFirewallRulesetsRules) | **Get** /organizations/{organizationId}/policies/global/firewall/rulesets/rules | List Organization-Wide Policy Firewall Rules
[**GetOrganizationWirelessAirMarshalRules**](RulesApi.md#GetOrganizationWirelessAirMarshalRules) | **Get** /organizations/{organizationId}/wireless/airMarshal/rules | Returns the current Air Marshal rules for this organization
[**UpdateNetworkApplianceTrafficShapingRules**](RulesApi.md#UpdateNetworkApplianceTrafficShapingRules) | **Put** /networks/{networkId}/appliance/trafficShaping/rules | Update the traffic shaping settings rules for a Security Appliance or Secure Router network
[**UpdateNetworkWirelessAirMarshalRule**](RulesApi.md#UpdateNetworkWirelessAirMarshalRule) | **Put** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Update a rule
[**UpdateNetworkWirelessSsidTrafficShapingRules**](RulesApi.md#UpdateNetworkWirelessSsidTrafficShapingRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Update the traffic shaping rules for an SSID on an MR network.
[**UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride**](RulesApi.md#UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride) | **Put** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/{ruleId}/override | Update a rule override for an intrusion policy.
[**UpdateOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesApi.md#UpdateOrganizationPoliciesGlobalFirewallRulesetsRule) | **Put** /organizations/{organizationId}/policies/global/firewall/rulesets/rules/{ruleId} | Update an Organization-Wide Policy Firewall Rule



## CreateNetworkWirelessAirMarshalRule

> InlineResponse20121 CreateNetworkWirelessAirMarshalRule(ctx, networkId).CreateNetworkWirelessAirMarshalRule(createNetworkWirelessAirMarshalRule).Execute()

Creates a new rule



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
    createNetworkWirelessAirMarshalRule := *openapiclient.NewInlineObject202("Type_example", *openapiclient.NewNetworksNetworkIdWirelessAirMarshalRulesMatch("Type_example", "String_example")) // InlineObject202 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.CreateNetworkWirelessAirMarshalRule(context.Background(), networkId).CreateNetworkWirelessAirMarshalRule(createNetworkWirelessAirMarshalRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.CreateNetworkWirelessAirMarshalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessAirMarshalRule`: InlineResponse20121
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.CreateNetworkWirelessAirMarshalRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessAirMarshalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWirelessAirMarshalRule** | [**InlineObject202**](InlineObject202.md) |  | 

### Return type

[**InlineResponse20121**](InlineResponse20121.md)

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
    resp, r, err := apiClient.RulesApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(context.Background(), organizationId, policyId, ruleId).CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(createOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: InlineResponse200296Items
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.CreateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: %v\n", resp)
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


## CreateOrganizationPoliciesGlobalFirewallRulesetsRule

> InlineResponse200422Items CreateOrganizationPoliciesGlobalFirewallRulesetsRule(ctx, organizationId).CreateOrganizationPoliciesGlobalFirewallRulesetsRule(createOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()

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
    createOrganizationPoliciesGlobalFirewallRulesetsRule := *openapiclient.NewInlineObject335("Name_example", "RulesetId_example", "Policy_example", *openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources([]string{"MatchCriteria_example"}), *openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations([]string{"MatchCriteria_example"})) // InlineObject335 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId).CreateOrganizationPoliciesGlobalFirewallRulesetsRule(createOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalFirewallRulesetsRule`: InlineResponse200422Items
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule`: %v\n", resp)
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

 **createOrganizationPoliciesGlobalFirewallRulesetsRule** | [**InlineObject335**](InlineObject335.md) |  | 

### Return type

[**InlineResponse200422Items**](InlineResponse200422Items.md)

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
    resp, r, err := apiClient.RulesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(context.Background(), organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides`: InlineResponse200296
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides`: %v\n", resp)
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


## DeleteNetworkWirelessAirMarshalRule

> DeleteNetworkWirelessAirMarshalRule(ctx, networkId, ruleId).Execute()

Delete an Air Marshal rule.



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
    ruleId := "ruleId_example" // string | Rule ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.DeleteNetworkWirelessAirMarshalRule(context.Background(), networkId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DeleteNetworkWirelessAirMarshalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessAirMarshalRuleRequest struct via the builder pattern


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
    resp, r, err := apiClient.RulesApi.DeleteOrganizationApplianceSecurityIntrusionRulesOverride(context.Background(), organizationId, overrideId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DeleteOrganizationApplianceSecurityIntrusionRulesOverride``: %v\n", err)
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
    resp, r, err := apiClient.RulesApi.DeleteOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.DeleteOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
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


## GetNetworkApplianceTrafficShapingRules

> InlineResponse20098 GetNetworkApplianceTrafficShapingRules(ctx, networkId).Execute()

Display the traffic shaping settings rules for a Security Appliance or Secure Router network



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
    resp, r, err := apiClient.RulesApi.GetNetworkApplianceTrafficShapingRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetNetworkApplianceTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingRules`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidTrafficShapingRules

> InlineResponse200260 GetNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).Execute()

Display the traffic shaping settings for a SSID on an MR network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.GetNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidTrafficShapingRules`: InlineResponse200260
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200260**](InlineResponse200260.md)

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
    resp, r, err := apiClient.RulesApi.GetOrganizationApplianceSecurityIntrusionRules(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).ParentRuleGroupIds(parentRuleGroupIds).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetOrganizationApplianceSecurityIntrusionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRules`: InlineResponse200300
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetOrganizationApplianceSecurityIntrusionRules`: %v\n", resp)
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
    resp, r, err := apiClient.RulesApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRulesOverrides`: InlineResponse200301
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides`: %v\n", resp)
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


## GetOrganizationPoliciesGlobalFirewallRulesetsRules

> InlineResponse200422 GetOrganizationPoliciesGlobalFirewallRulesetsRules(ctx, organizationId).RulesetIds(rulesetIds).RuleIds(ruleIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.RulesApi.GetOrganizationPoliciesGlobalFirewallRulesetsRules(context.Background(), organizationId).RulesetIds(rulesetIds).RuleIds(ruleIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetOrganizationPoliciesGlobalFirewallRulesetsRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalFirewallRulesetsRules`: InlineResponse200422
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetOrganizationPoliciesGlobalFirewallRulesetsRules`: %v\n", resp)
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

[**InlineResponse200422**](InlineResponse200422.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessAirMarshalRules

> InlineResponse200485 GetOrganizationWirelessAirMarshalRules(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns the current Air Marshal rules for this organization



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
    networkIds := []string{"Inner_example"} // []string | (optional) The set of network IDs to include. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.GetOrganizationWirelessAirMarshalRules(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.GetOrganizationWirelessAirMarshalRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessAirMarshalRules`: InlineResponse200485
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.GetOrganizationWirelessAirMarshalRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessAirMarshalRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | (optional) The set of network IDs to include. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200485**](InlineResponse200485.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingRules

> InlineResponse20098 UpdateNetworkApplianceTrafficShapingRules(ctx, networkId).UpdateNetworkApplianceTrafficShapingRules(updateNetworkApplianceTrafficShapingRules).Execute()

Update the traffic shaping settings rules for a Security Appliance or Secure Router network



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
    updateNetworkApplianceTrafficShapingRules := *openapiclient.NewInlineObject84() // InlineObject84 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateNetworkApplianceTrafficShapingRules(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRules(updateNetworkApplianceTrafficShapingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateNetworkApplianceTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingRules`: InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.UpdateNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingRules** | [**InlineObject84**](InlineObject84.md) |  | 

### Return type

[**InlineResponse20098**](InlineResponse20098.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessAirMarshalRule

> InlineResponse20121 UpdateNetworkWirelessAirMarshalRule(ctx, networkId, ruleId).UpdateNetworkWirelessAirMarshalRule(updateNetworkWirelessAirMarshalRule).Execute()

Update a rule



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
    ruleId := "ruleId_example" // string | Rule ID
    updateNetworkWirelessAirMarshalRule := *openapiclient.NewInlineObject203() // InlineObject203 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateNetworkWirelessAirMarshalRule(context.Background(), networkId, ruleId).UpdateNetworkWirelessAirMarshalRule(updateNetworkWirelessAirMarshalRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateNetworkWirelessAirMarshalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessAirMarshalRule`: InlineResponse20121
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.UpdateNetworkWirelessAirMarshalRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessAirMarshalRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessAirMarshalRule** | [**InlineObject203**](InlineObject203.md) |  | 

### Return type

[**InlineResponse20121**](InlineResponse20121.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidTrafficShapingRules

> InlineResponse200260 UpdateNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()

Update the traffic shaping rules for an SSID on an MR network.



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidTrafficShapingRules := *openapiclient.NewInlineObject232() // InlineObject232 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidTrafficShapingRules`: InlineResponse200260
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.UpdateNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidTrafficShapingRules** | [**InlineObject232**](InlineObject232.md) |  | 

### Return type

[**InlineResponse200260**](InlineResponse200260.md)

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
    resp, r, err := apiClient.RulesApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(context.Background(), organizationId, policyId, ruleId).UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride(updateOrganizationApplianceSecurityIntrusionPolicyRuleOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: InlineResponse200296Items
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.UpdateOrganizationApplianceSecurityIntrusionPolicyRuleOverride`: %v\n", resp)
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


## UpdateOrganizationPoliciesGlobalFirewallRulesetsRule

> InlineResponse200422Items UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(ctx, organizationId, ruleId).UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(updateOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()

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
    updateOrganizationPoliciesGlobalFirewallRulesetsRule := *openapiclient.NewInlineObject336() // InlineObject336 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId, ruleId).UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(updateOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalFirewallRulesetsRule`: InlineResponse200422Items
    fmt.Fprintf(os.Stdout, "Response from `RulesApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule`: %v\n", resp)
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


 **updateOrganizationPoliciesGlobalFirewallRulesetsRule** | [**InlineObject336**](InlineObject336.md) |  | 

### Return type

[**InlineResponse200422Items**](InlineResponse200422Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

