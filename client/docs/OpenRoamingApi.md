# \OpenRoamingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWirelessSsidsOpenRoamingByNetwork**](OpenRoamingApi.md#GetOrganizationWirelessSsidsOpenRoamingByNetwork) | **Get** /organizations/{organizationId}/wireless/ssids/openRoaming/byNetwork | Returns an array of objects, each containing SSID OpenRoaming configs for the corresponding network
[**UpdateNetworkWirelessSsidOpenRoaming**](OpenRoamingApi.md#UpdateNetworkWirelessSsidOpenRoaming) | **Put** /networks/{networkId}/wireless/ssids/{number}/openRoaming | Update the OpenRoaming setting for the SSID



## GetOrganizationWirelessSsidsOpenRoamingByNetwork

> InlineResponse200390 GetOrganizationWirelessSsidsOpenRoamingByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).IncludeDisabledSsids(includeDisabledSsids).Execute()

Returns an array of objects, each containing SSID OpenRoaming configs for the corresponding network



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter OpenRoaming configuration by Network Id. (optional)
    includeDisabledSsids := true // bool | Optional parameter to include OpenRoaming configuration for disabled ssids. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenRoamingApi.GetOrganizationWirelessSsidsOpenRoamingByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).IncludeDisabledSsids(includeDisabledSsids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenRoamingApi.GetOrganizationWirelessSsidsOpenRoamingByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessSsidsOpenRoamingByNetwork`: InlineResponse200390
    fmt.Fprintf(os.Stdout, "Response from `OpenRoamingApi.GetOrganizationWirelessSsidsOpenRoamingByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessSsidsOpenRoamingByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter OpenRoaming configuration by Network Id. | 
 **includeDisabledSsids** | **bool** | Optional parameter to include OpenRoaming configuration for disabled ssids. | 

### Return type

[**InlineResponse200390**](InlineResponse200390.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidOpenRoaming

> InlineResponse200216 UpdateNetworkWirelessSsidOpenRoaming(ctx, networkId, number).UpdateNetworkWirelessSsidOpenRoaming(updateNetworkWirelessSsidOpenRoaming).Execute()

Update the OpenRoaming setting for the SSID



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
    updateNetworkWirelessSsidOpenRoaming := *openapiclient.NewInlineObject205() // InlineObject205 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenRoamingApi.UpdateNetworkWirelessSsidOpenRoaming(context.Background(), networkId, number).UpdateNetworkWirelessSsidOpenRoaming(updateNetworkWirelessSsidOpenRoaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenRoamingApi.UpdateNetworkWirelessSsidOpenRoaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidOpenRoaming`: InlineResponse200216
    fmt.Fprintf(os.Stdout, "Response from `OpenRoamingApi.UpdateNetworkWirelessSsidOpenRoaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidOpenRoamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidOpenRoaming** | [**InlineObject205**](InlineObject205.md) |  | 

### Return type

[**InlineResponse200216**](InlineResponse200216.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

