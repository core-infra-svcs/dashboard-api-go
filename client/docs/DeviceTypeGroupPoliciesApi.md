# \DeviceTypeGroupPoliciesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessSsidDeviceTypeGroupPolicies**](DeviceTypeGroupPoliciesApi.md#GetNetworkWirelessSsidDeviceTypeGroupPolicies) | **Get** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | List the device type group policies for the SSID
[**UpdateNetworkWirelessSsidDeviceTypeGroupPolicies**](DeviceTypeGroupPoliciesApi.md#UpdateNetworkWirelessSsidDeviceTypeGroupPolicies) | **Put** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | Update the device type group policies for the SSID



## GetNetworkWirelessSsidDeviceTypeGroupPolicies

> map[string]interface{} GetNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).Execute()

List the device type group policies for the SSID



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
    networkId := "networkId_example" // string | 
    number := "number_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceTypeGroupPoliciesApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTypeGroupPoliciesApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidDeviceTypeGroupPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceTypeGroupPoliciesApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidDeviceTypeGroupPolicies

> map[string]interface{} UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()

Update the device type group policies for the SSID



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
    networkId := "networkId_example" // string | 
    number := "number_example" // string | 
    updateNetworkWirelessSsidDeviceTypeGroupPolicies := *openapiclient.NewInlineObject155() // InlineObject155 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceTypeGroupPoliciesApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTypeGroupPoliciesApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceTypeGroupPoliciesApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidDeviceTypeGroupPolicies** | [**InlineObject155**](InlineObject155.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

