# \OweApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWirelessSsidsOweByNetwork**](OweApi.md#GetOrganizationWirelessSsidsOweByNetwork) | **Get** /organizations/{organizationId}/wireless/ssids/owe/byNetwork | Returns an array of objects, each containing OWE transition pairs for the corresponding network
[**UpdateNetworkWirelessSsidsOwe**](OweApi.md#UpdateNetworkWirelessSsidsOwe) | **Put** /networks/{networkId}/wireless/ssids/owe | Update the OWE transition pairs for a network



## GetOrganizationWirelessSsidsOweByNetwork

> InlineResponse200516 GetOrganizationWirelessSsidsOweByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Returns an array of objects, each containing OWE transition pairs for the corresponding network



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter OWE transition configuration by Network Id. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OweApi.GetOrganizationWirelessSsidsOweByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OweApi.GetOrganizationWirelessSsidsOweByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessSsidsOweByNetwork`: InlineResponse200516
    fmt.Fprintf(os.Stdout, "Response from `OweApi.GetOrganizationWirelessSsidsOweByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessSsidsOweByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter OWE transition configuration by Network Id. | 

### Return type

[**InlineResponse200516**](InlineResponse200516.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidsOwe

> InlineResponse200248 UpdateNetworkWirelessSsidsOwe(ctx, networkId).UpdateNetworkWirelessSsidsOwe(updateNetworkWirelessSsidsOwe).Execute()

Update the OWE transition pairs for a network



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
    updateNetworkWirelessSsidsOwe := *openapiclient.NewInlineObject218([]openapiclient.NetworksNetworkIdWirelessSsidsOweTransitions{*openapiclient.NewNetworksNetworkIdWirelessSsidsOweTransitions(*openapiclient.NewNetworksNetworkIdWirelessSsidsOweSsids(*openapiclient.NewNetworksNetworkIdWirelessSsidsOweSsidsEncrypted(int32(123))))}) // InlineObject218 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OweApi.UpdateNetworkWirelessSsidsOwe(context.Background(), networkId).UpdateNetworkWirelessSsidsOwe(updateNetworkWirelessSsidsOwe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OweApi.UpdateNetworkWirelessSsidsOwe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidsOwe`: InlineResponse200248
    fmt.Fprintf(os.Stdout, "Response from `OweApi.UpdateNetworkWirelessSsidsOwe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidsOweRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessSsidsOwe** | [**InlineObject218**](InlineObject218.md) |  | 

### Return type

[**InlineResponse200248**](InlineResponse200248.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

