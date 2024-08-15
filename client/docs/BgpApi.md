# \BgpApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceVpnBgp**](BgpApi.md#GetNetworkApplianceVpnBgp) | **Get** /networks/{networkId}/appliance/vpn/bgp | Return a Hub BGP Configuration
[**UpdateNetworkApplianceVpnBgp**](BgpApi.md#UpdateNetworkApplianceVpnBgp) | **Put** /networks/{networkId}/appliance/vpn/bgp | Update a Hub BGP Configuration



## GetNetworkApplianceVpnBgp

> InlineResponse20069 GetNetworkApplianceVpnBgp(ctx, networkId).Execute()

Return a Hub BGP Configuration



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
    resp, r, err := apiClient.BgpApi.GetNetworkApplianceVpnBgp(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BgpApi.GetNetworkApplianceVpnBgp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVpnBgp`: InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `BgpApi.GetNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20069**](InlineResponse20069.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVpnBgp

> InlineResponse20069 UpdateNetworkApplianceVpnBgp(ctx, networkId).UpdateNetworkApplianceVpnBgp(updateNetworkApplianceVpnBgp).Execute()

Update a Hub BGP Configuration



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
    updateNetworkApplianceVpnBgp := *openapiclient.NewInlineObject74(false) // InlineObject74 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BgpApi.UpdateNetworkApplianceVpnBgp(context.Background(), networkId).UpdateNetworkApplianceVpnBgp(updateNetworkApplianceVpnBgp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BgpApi.UpdateNetworkApplianceVpnBgp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVpnBgp`: InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `BgpApi.UpdateNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnBgp** | [**InlineObject74**](InlineObject74.md) |  | 

### Return type

[**InlineResponse20069**](InlineResponse20069.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

