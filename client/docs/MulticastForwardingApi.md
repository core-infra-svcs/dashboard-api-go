# \MulticastForwardingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceFirewallMulticastForwardingByNetwork**](MulticastForwardingApi.md#GetOrganizationApplianceFirewallMulticastForwardingByNetwork) | **Get** /organizations/{organizationId}/appliance/firewall/multicastForwarding/byNetwork | List Static Multicasting forwarding settings for MX networks
[**UpdateNetworkApplianceFirewallMulticastForwarding**](MulticastForwardingApi.md#UpdateNetworkApplianceFirewallMulticastForwarding) | **Put** /networks/{networkId}/appliance/firewall/multicastForwarding | Update static multicast forward rules for a network



## GetOrganizationApplianceFirewallMulticastForwardingByNetwork

> InlineResponse200233 GetOrganizationApplianceFirewallMulticastForwardingByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List Static Multicasting forwarding settings for MX networks



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MulticastForwardingApi.GetOrganizationApplianceFirewallMulticastForwardingByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MulticastForwardingApi.GetOrganizationApplianceFirewallMulticastForwardingByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceFirewallMulticastForwardingByNetwork`: InlineResponse200233
    fmt.Fprintf(os.Stdout, "Response from `MulticastForwardingApi.GetOrganizationApplianceFirewallMulticastForwardingByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceFirewallMulticastForwardingByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**InlineResponse200233**](InlineResponse200233.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallMulticastForwarding

> InlineResponse20059 UpdateNetworkApplianceFirewallMulticastForwarding(ctx, networkId).UpdateNetworkApplianceFirewallMulticastForwarding(updateNetworkApplianceFirewallMulticastForwarding).Execute()

Update static multicast forward rules for a network



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
    updateNetworkApplianceFirewallMulticastForwarding := *openapiclient.NewInlineObject50([]openapiclient.NetworksNetworkIdApplianceFirewallMulticastForwardingRules{*openapiclient.NewNetworksNetworkIdApplianceFirewallMulticastForwardingRules("Description_example", "Address_example", []string{"VlanIds_example"})}) // InlineObject50 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MulticastForwardingApi.UpdateNetworkApplianceFirewallMulticastForwarding(context.Background(), networkId).UpdateNetworkApplianceFirewallMulticastForwarding(updateNetworkApplianceFirewallMulticastForwarding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MulticastForwardingApi.UpdateNetworkApplianceFirewallMulticastForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallMulticastForwarding`: InlineResponse20059
    fmt.Fprintf(os.Stdout, "Response from `MulticastForwardingApi.UpdateNetworkApplianceFirewallMulticastForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallMulticastForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallMulticastForwarding** | [**InlineObject50**](InlineObject50.md) |  | 

### Return type

[**InlineResponse20059**](InlineResponse20059.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

