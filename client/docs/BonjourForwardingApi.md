# \BonjourForwardingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessSsidBonjourForwarding**](BonjourForwardingApi.md#GetNetworkWirelessSsidBonjourForwarding) | **Get** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | List the Bonjour forwarding setting and rules for the SSID
[**UpdateNetworkWirelessSsidBonjourForwarding**](BonjourForwardingApi.md#UpdateNetworkWirelessSsidBonjourForwarding) | **Put** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | Update the bonjour forwarding setting and rules for the SSID



## GetNetworkWirelessSsidBonjourForwarding

> InlineResponse200209 GetNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).Execute()

List the Bonjour forwarding setting and rules for the SSID



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
    resp, r, err := apiClient.BonjourForwardingApi.GetNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BonjourForwardingApi.GetNetworkWirelessSsidBonjourForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidBonjourForwarding`: InlineResponse200209
    fmt.Fprintf(os.Stdout, "Response from `BonjourForwardingApi.GetNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200209**](InlineResponse200209.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidBonjourForwarding

> InlineResponse200209 UpdateNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).UpdateNetworkWirelessSsidBonjourForwarding(updateNetworkWirelessSsidBonjourForwarding).Execute()

Update the bonjour forwarding setting and rules for the SSID



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
    updateNetworkWirelessSsidBonjourForwarding := *openapiclient.NewInlineObject196() // InlineObject196 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BonjourForwardingApi.UpdateNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).UpdateNetworkWirelessSsidBonjourForwarding(updateNetworkWirelessSsidBonjourForwarding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BonjourForwardingApi.UpdateNetworkWirelessSsidBonjourForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidBonjourForwarding`: InlineResponse200209
    fmt.Fprintf(os.Stdout, "Response from `BonjourForwardingApi.UpdateNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidBonjourForwarding** | [**InlineObject196**](InlineObject196.md) |  | 

### Return type

[**InlineResponse200209**](InlineResponse200209.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

