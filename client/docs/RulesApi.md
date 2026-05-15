# \RulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWirelessAirMarshalRule**](RulesApi.md#CreateNetworkWirelessAirMarshalRule) | **Post** /networks/{networkId}/wireless/airMarshal/rules | Creates a new rule
[**CreateOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesApi.md#CreateOrganizationPoliciesGlobalFirewallRulesetsRule) | **Post** /organizations/{organizationId}/policies/global/firewall/rulesets/rules | Create an Organization-Wide Policy Firewall Rule
[**DeleteNetworkWirelessAirMarshalRule**](RulesApi.md#DeleteNetworkWirelessAirMarshalRule) | **Delete** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Delete an Air Marshal rule.
[**DeleteOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesApi.md#DeleteOrganizationPoliciesGlobalFirewallRulesetsRule) | **Delete** /organizations/{organizationId}/policies/global/firewall/rulesets/rules/{ruleId} | Delete an Organization-Wide Policy Firewall Rule
[**GetNetworkApplianceTrafficShapingRules**](RulesApi.md#GetNetworkApplianceTrafficShapingRules) | **Get** /networks/{networkId}/appliance/trafficShaping/rules | Display the traffic shaping settings rules for an MX network
[**GetNetworkWirelessSsidTrafficShapingRules**](RulesApi.md#GetNetworkWirelessSsidTrafficShapingRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Display the traffic shaping settings for a SSID on an MR network
[**GetOrganizationPoliciesGlobalFirewallRulesetsRules**](RulesApi.md#GetOrganizationPoliciesGlobalFirewallRulesetsRules) | **Get** /organizations/{organizationId}/policies/global/firewall/rulesets/rules | List Organization-Wide Policy Firewall Rules
[**GetOrganizationWirelessAirMarshalRules**](RulesApi.md#GetOrganizationWirelessAirMarshalRules) | **Get** /organizations/{organizationId}/wireless/airMarshal/rules | Returns the current Air Marshal rules for this organization
[**UpdateNetworkApplianceTrafficShapingRules**](RulesApi.md#UpdateNetworkApplianceTrafficShapingRules) | **Put** /networks/{networkId}/appliance/trafficShaping/rules | Update the traffic shaping settings rules for an MX network
[**UpdateNetworkWirelessAirMarshalRule**](RulesApi.md#UpdateNetworkWirelessAirMarshalRule) | **Put** /networks/{networkId}/wireless/airMarshal/rules/{ruleId} | Update a rule
[**UpdateNetworkWirelessSsidTrafficShapingRules**](RulesApi.md#UpdateNetworkWirelessSsidTrafficShapingRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Update the traffic shaping rules for an SSID on an MR network.
[**UpdateOrganizationPoliciesGlobalFirewallRulesetsRule**](RulesApi.md#UpdateOrganizationPoliciesGlobalFirewallRulesetsRule) | **Put** /organizations/{organizationId}/policies/global/firewall/rulesets/rules/{ruleId} | Update an Organization-Wide Policy Firewall Rule



## CreateNetworkWirelessAirMarshalRule

> InlineResponse20116 CreateNetworkWirelessAirMarshalRule(ctx, networkId).CreateNetworkWirelessAirMarshalRule(createNetworkWirelessAirMarshalRule).Execute()

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
    createNetworkWirelessAirMarshalRule := *openapiclient.NewInlineObject186("Type_example", *openapiclient.NewNetworksNetworkIdWirelessAirMarshalRulesMatch("Type_example", "String_example")) // InlineObject186 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.CreateNetworkWirelessAirMarshalRule(context.Background(), networkId).CreateNetworkWirelessAirMarshalRule(createNetworkWirelessAirMarshalRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.CreateNetworkWirelessAirMarshalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessAirMarshalRule`: InlineResponse20116
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

 **createNetworkWirelessAirMarshalRule** | [**InlineObject186**](InlineObject186.md) |  | 

### Return type

[**InlineResponse20116**](InlineResponse20116.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPoliciesGlobalFirewallRulesetsRule

> InlineResponse200374Items CreateOrganizationPoliciesGlobalFirewallRulesetsRule(ctx, organizationId).CreateOrganizationPoliciesGlobalFirewallRulesetsRule(createOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()

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
    createOrganizationPoliciesGlobalFirewallRulesetsRule := *openapiclient.NewInlineObject301("Name_example", "RulesetId_example", "Policy_example", *openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesSources([]string{"MatchCriteria_example"}), *openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinations([]string{"MatchCriteria_example"})) // InlineObject301 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId).CreateOrganizationPoliciesGlobalFirewallRulesetsRule(createOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.CreateOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalFirewallRulesetsRule`: InlineResponse200374Items
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

 **createOrganizationPoliciesGlobalFirewallRulesetsRule** | [**InlineObject301**](InlineObject301.md) |  | 

### Return type

[**InlineResponse200374Items**](InlineResponse200374Items.md)

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

> InlineResponse20092 GetNetworkApplianceTrafficShapingRules(ctx, networkId).Execute()

Display the traffic shaping settings rules for an MX network



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
    // response from `GetNetworkApplianceTrafficShapingRules`: InlineResponse20092
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

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidTrafficShapingRules

> InlineResponse200245 GetNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).Execute()

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
    // response from `GetNetworkWirelessSsidTrafficShapingRules`: InlineResponse200245
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

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalFirewallRulesetsRules

> InlineResponse200374 GetOrganizationPoliciesGlobalFirewallRulesetsRules(ctx, organizationId).RulesetIds(rulesetIds).RuleIds(ruleIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    // response from `GetOrganizationPoliciesGlobalFirewallRulesetsRules`: InlineResponse200374
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

[**InlineResponse200374**](InlineResponse200374.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessAirMarshalRules

> InlineResponse200428 GetOrganizationWirelessAirMarshalRules(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    // response from `GetOrganizationWirelessAirMarshalRules`: InlineResponse200428
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

[**InlineResponse200428**](InlineResponse200428.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingRules

> InlineResponse20092 UpdateNetworkApplianceTrafficShapingRules(ctx, networkId).UpdateNetworkApplianceTrafficShapingRules(updateNetworkApplianceTrafficShapingRules).Execute()

Update the traffic shaping settings rules for an MX network



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
    updateNetworkApplianceTrafficShapingRules := *openapiclient.NewInlineObject75() // InlineObject75 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateNetworkApplianceTrafficShapingRules(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRules(updateNetworkApplianceTrafficShapingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateNetworkApplianceTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingRules`: InlineResponse20092
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

 **updateNetworkApplianceTrafficShapingRules** | [**InlineObject75**](InlineObject75.md) |  | 

### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessAirMarshalRule

> InlineResponse20116 UpdateNetworkWirelessAirMarshalRule(ctx, networkId, ruleId).UpdateNetworkWirelessAirMarshalRule(updateNetworkWirelessAirMarshalRule).Execute()

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
    updateNetworkWirelessAirMarshalRule := *openapiclient.NewInlineObject187() // InlineObject187 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateNetworkWirelessAirMarshalRule(context.Background(), networkId, ruleId).UpdateNetworkWirelessAirMarshalRule(updateNetworkWirelessAirMarshalRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateNetworkWirelessAirMarshalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessAirMarshalRule`: InlineResponse20116
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


 **updateNetworkWirelessAirMarshalRule** | [**InlineObject187**](InlineObject187.md) |  | 

### Return type

[**InlineResponse20116**](InlineResponse20116.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidTrafficShapingRules

> InlineResponse200245 UpdateNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()

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
    updateNetworkWirelessSsidTrafficShapingRules := *openapiclient.NewInlineObject214() // InlineObject214 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidTrafficShapingRules`: InlineResponse200245
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


 **updateNetworkWirelessSsidTrafficShapingRules** | [**InlineObject214**](InlineObject214.md) |  | 

### Return type

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPoliciesGlobalFirewallRulesetsRule

> InlineResponse200374Items UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(ctx, organizationId, ruleId).UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(updateOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()

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
    updateOrganizationPoliciesGlobalFirewallRulesetsRule := *openapiclient.NewInlineObject302() // InlineObject302 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(context.Background(), organizationId, ruleId).UpdateOrganizationPoliciesGlobalFirewallRulesetsRule(updateOrganizationPoliciesGlobalFirewallRulesetsRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesApi.UpdateOrganizationPoliciesGlobalFirewallRulesetsRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalFirewallRulesetsRule`: InlineResponse200374Items
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


 **updateOrganizationPoliciesGlobalFirewallRulesetsRule** | [**InlineObject302**](InlineObject302.md) |  | 

### Return type

[**InlineResponse200374Items**](InlineResponse200374Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

