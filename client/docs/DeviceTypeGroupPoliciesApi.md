# \DeviceTypeGroupPoliciesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessSsidDeviceTypeGroupPolicies**](DeviceTypeGroupPoliciesApi.md#GetNetworkWirelessSsidDeviceTypeGroupPolicies) | **Get** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | List the device type group policies for the SSID
[**UpdateNetworkWirelessSsidDeviceTypeGroupPolicies**](DeviceTypeGroupPoliciesApi.md#UpdateNetworkWirelessSsidDeviceTypeGroupPolicies) | **Put** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | Update the device type group policies for the SSID



## GetNetworkWirelessSsidDeviceTypeGroupPolicies

> InlineResponse200250 GetNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).Execute()

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
    networkId := "networkId_example" // string | Network ID
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceTypeGroupPoliciesApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTypeGroupPoliciesApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidDeviceTypeGroupPolicies`: InlineResponse200250
    fmt.Fprintf(os.Stdout, "Response from `DeviceTypeGroupPoliciesApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200250**](InlineResponse200250.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidDeviceTypeGroupPolicies

> InlineResponse200250 UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()

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
    networkId := "networkId_example" // string | Network ID
    number := "number_example" // string | Number
    updateNetworkWirelessSsidDeviceTypeGroupPolicies := *openapiclient.NewInlineObject221() // InlineObject221 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceTypeGroupPoliciesApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTypeGroupPoliciesApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: InlineResponse200250
    fmt.Fprintf(os.Stdout, "Response from `DeviceTypeGroupPoliciesApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidDeviceTypeGroupPolicies** | [**InlineObject221**](InlineObject221.md) |  | 

### Return type

[**InlineResponse200250**](InlineResponse200250.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

