# \OverridesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides**](OverridesApi.md#DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/ruleGroups/overrides/declare | Declare the desired rule group overrides for an intrusion policy.
[**DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides**](OverridesApi.md#DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides) | **Post** /organizations/{organizationId}/appliance/security/intrusion/policies/{policyId}/rules/overrides/declare | Declare the desired rule overrides for an intrusion policy.
[**DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride**](OverridesApi.md#DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride) | **Delete** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overrides/{overrideId} | Delete a rule group override for an intrusion policy.
[**DeleteOrganizationApplianceSecurityIntrusionRulesOverride**](OverridesApi.md#DeleteOrganizationApplianceSecurityIntrusionRulesOverride) | **Delete** /organizations/{organizationId}/appliance/security/intrusion/rules/overrides/{overrideId} | Delete a rule override for an intrusion policy.
[**GetDeviceWirelessRadioOverrides**](OverridesApi.md#GetDeviceWirelessRadioOverrides) | **Get** /devices/{serial}/wireless/radio/overrides | Return the radio overrides of a device
[**GetNetworkWirelessSsidOverrides**](OverridesApi.md#GetNetworkWirelessSsidOverrides) | **Get** /networks/{networkId}/wireless/ssids/{number}/overrides | Display the overrides for this SSID
[**GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides**](OverridesApi.md#GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides) | **Get** /organizations/{organizationId}/appliance/security/intrusion/ruleGroups/overrides | List the rule group overrides configured for an intrusion policy.
[**GetOrganizationApplianceSecurityIntrusionRulesOverrides**](OverridesApi.md#GetOrganizationApplianceSecurityIntrusionRulesOverrides) | **Get** /organizations/{organizationId}/appliance/security/intrusion/rules/overrides | List the rule overrides configured for an intrusion policy.
[**GetOrganizationWirelessRadioOverridesByDevice**](OverridesApi.md#GetOrganizationWirelessRadioOverridesByDevice) | **Get** /organizations/{organizationId}/wireless/radio/overrides/byDevice | Return a list of radio overrides
[**UpdateDeviceWirelessRadioOverrides**](OverridesApi.md#UpdateDeviceWirelessRadioOverrides) | **Put** /devices/{serial}/wireless/radio/overrides | Update 2.4 GHz, 5 GHz, and 6 GHz radio settings (channel, channel width, power, and enable/disable) that override RF profiles
[**UpdateNetworkWirelessSsidOverrides**](OverridesApi.md#UpdateNetworkWirelessSsidOverrides) | **Put** /networks/{networkId}/wireless/ssids/{number}/overrides | Update the overrides for this SSID



## DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides

> InlineResponse200301 DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(ctx, organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides).Execute()

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
    declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides := *openapiclient.NewInlineObject268([]openapiclient.OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems{*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareItems(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareGroup("Id_example"), *openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRuleGroupsOverridesDeclareSecurity(int32(123)))}) // InlineObject268 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(context.Background(), organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides`: InlineResponse200301
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides`: %v\n", resp)
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


 **declareOrganizationApplianceSecurityIntrusionPolicyRuleGroupsOverrides** | [**InlineObject268**](InlineObject268.md) |  | 

### Return type

[**InlineResponse200301**](InlineResponse200301.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides

> InlineResponse200302 DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(ctx, organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides).Execute()

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
    declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides := *openapiclient.NewInlineObject271([]openapiclient.OrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems{*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareItems(*openapiclient.NewOrganizationsOrganizationIdApplianceSecurityIntrusionPoliciesPolicyIdRulesOverridesDeclareRule("Id_example"), "Action_example")}) // InlineObject271 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(context.Background(), organizationId, policyId).DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides(declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides`: InlineResponse200302
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.DeclareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides`: %v\n", resp)
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


 **declareOrganizationApplianceSecurityIntrusionPolicyRulesOverrides** | [**InlineObject271**](InlineObject271.md) |  | 

### Return type

[**InlineResponse200302**](InlineResponse200302.md)

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
    resp, r, err := apiClient.OverridesApi.DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride(context.Background(), organizationId, overrideId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.DeleteOrganizationApplianceSecurityIntrusionRuleGroupsOverride``: %v\n", err)
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
    resp, r, err := apiClient.OverridesApi.DeleteOrganizationApplianceSecurityIntrusionRulesOverride(context.Background(), organizationId, overrideId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.DeleteOrganizationApplianceSecurityIntrusionRulesOverride``: %v\n", err)
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


## GetDeviceWirelessRadioOverrides

> InlineResponse20066 GetDeviceWirelessRadioOverrides(ctx, serial).Execute()

Return the radio overrides of a device



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.GetDeviceWirelessRadioOverrides(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.GetDeviceWirelessRadioOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessRadioOverrides`: InlineResponse20066
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.GetDeviceWirelessRadioOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessRadioOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20066**](InlineResponse20066.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidOverrides

> InlineResponse200260 GetNetworkWirelessSsidOverrides(ctx, networkId, number).Execute()

Display the overrides for this SSID



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
    resp, r, err := apiClient.OverridesApi.GetNetworkWirelessSsidOverrides(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.GetNetworkWirelessSsidOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidOverrides`: InlineResponse200260
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.GetNetworkWirelessSsidOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidOverridesRequest struct via the builder pattern


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


## GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides

> InlineResponse200304 GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()

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
    resp, r, err := apiClient.OverridesApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides`: InlineResponse200304
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.GetOrganizationApplianceSecurityIntrusionRuleGroupsOverrides`: %v\n", resp)
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

[**InlineResponse200304**](InlineResponse200304.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusionRulesOverrides

> InlineResponse200307 GetOrganizationApplianceSecurityIntrusionRulesOverrides(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()

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
    resp, r, err := apiClient.OverridesApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PolicyIds(policyIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityIntrusionRulesOverrides`: InlineResponse200307
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.GetOrganizationApplianceSecurityIntrusionRulesOverrides`: %v\n", resp)
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

[**InlineResponse200307**](InlineResponse200307.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessRadioOverridesByDevice

> InlineResponse200523 GetOrganizationWirelessRadioOverridesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()

Return a list of radio overrides



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned radio overrides will be filtered to only include these networks. (optional)
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned radio overrides will be filtered to only include these serials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.GetOrganizationWirelessRadioOverridesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.GetOrganizationWirelessRadioOverridesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessRadioOverridesByDevice`: InlineResponse200523
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.GetOrganizationWirelessRadioOverridesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessRadioOverridesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned radio overrides will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned radio overrides will be filtered to only include these serials. | 

### Return type

[**InlineResponse200523**](InlineResponse200523.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessRadioOverrides

> InlineResponse20066 UpdateDeviceWirelessRadioOverrides(ctx, serial).UpdateDeviceWirelessRadioOverrides(updateDeviceWirelessRadioOverrides).Execute()

Update 2.4 GHz, 5 GHz, and 6 GHz radio settings (channel, channel width, power, and enable/disable) that override RF profiles



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
    serial := "serial_example" // string | Serial
    updateDeviceWirelessRadioOverrides := *openapiclient.NewInlineObject50() // InlineObject50 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.UpdateDeviceWirelessRadioOverrides(context.Background(), serial).UpdateDeviceWirelessRadioOverrides(updateDeviceWirelessRadioOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.UpdateDeviceWirelessRadioOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessRadioOverrides`: InlineResponse20066
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.UpdateDeviceWirelessRadioOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessRadioOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessRadioOverrides** | [**InlineObject50**](InlineObject50.md) |  | 

### Return type

[**InlineResponse20066**](InlineResponse20066.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidOverrides

> InlineResponse200260 UpdateNetworkWirelessSsidOverrides(ctx, networkId, number).UpdateNetworkWirelessSsidOverrides(updateNetworkWirelessSsidOverrides).Execute()

Update the overrides for this SSID



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
    updateNetworkWirelessSsidOverrides := *openapiclient.NewInlineObject232() // InlineObject232 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.UpdateNetworkWirelessSsidOverrides(context.Background(), networkId, number).UpdateNetworkWirelessSsidOverrides(updateNetworkWirelessSsidOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.UpdateNetworkWirelessSsidOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidOverrides`: InlineResponse200260
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.UpdateNetworkWirelessSsidOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidOverrides** | [**InlineObject232**](InlineObject232.md) |  | 

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

