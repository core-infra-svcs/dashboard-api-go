# \NatApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceUplinksNatByNetwork**](NatApi.md#GetOrganizationApplianceUplinksNatByNetwork) | **Get** /organizations/{organizationId}/appliance/uplinks/nat/byNetwork | Fetch uplink NAT settings of each network in the organization
[**UpdateNetworkApplianceUplinksNat**](NatApi.md#UpdateNetworkApplianceUplinksNat) | **Put** /networks/{networkId}/appliance/uplinks/nat | Update uplink NAT settings of the specified network



## GetOrganizationApplianceUplinksNatByNetwork

> []InlineResponse200286 GetOrganizationApplianceUplinksNatByNetwork(ctx, organizationId).NetworkIds(networkIds).Interfaces(interfaces).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Fetch uplink NAT settings of each network in the organization



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by the included set of network IDs (optional)
    interfaces := []string{"Inner_example"} // []string | Optional parameter to filter the results by the included set of interfaces (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatApi.GetOrganizationApplianceUplinksNatByNetwork(context.Background(), organizationId).NetworkIds(networkIds).Interfaces(interfaces).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatApi.GetOrganizationApplianceUplinksNatByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinksNatByNetwork`: []InlineResponse200286
    fmt.Fprintf(os.Stdout, "Response from `NatApi.GetOrganizationApplianceUplinksNatByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksNatByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter the results by the included set of network IDs | 
 **interfaces** | **[]string** | Optional parameter to filter the results by the included set of interfaces | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200286**](InlineResponse200286.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceUplinksNat

> InlineResponse200105 UpdateNetworkApplianceUplinksNat(ctx, networkId).UpdateNetworkApplianceUplinksNat(updateNetworkApplianceUplinksNat).Execute()

Update uplink NAT settings of the specified network



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
    updateNetworkApplianceUplinksNat := *openapiclient.NewInlineObject91([]openapiclient.NetworksNetworkIdApplianceUplinksNatUplinks{*openapiclient.NewNetworksNetworkIdApplianceUplinksNatUplinks("Interface_example", *openapiclient.NewNetworksNetworkIdApplianceUplinksNatNat(false))}) // InlineObject91 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatApi.UpdateNetworkApplianceUplinksNat(context.Background(), networkId).UpdateNetworkApplianceUplinksNat(updateNetworkApplianceUplinksNat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatApi.UpdateNetworkApplianceUplinksNat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceUplinksNat`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `NatApi.UpdateNetworkApplianceUplinksNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceUplinksNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceUplinksNat** | [**InlineObject91**](InlineObject91.md) |  | 

### Return type

[**InlineResponse200105**](InlineResponse200105.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

