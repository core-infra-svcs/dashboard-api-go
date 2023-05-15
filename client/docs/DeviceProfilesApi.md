# \DeviceProfilesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceDeviceProfiles**](DeviceProfilesApi.md#GetNetworkSmDeviceDeviceProfiles) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceProfiles | Get the installed profiles associated with a device
[**GetNetworkSmUserDeviceProfiles**](DeviceProfilesApi.md#GetNetworkSmUserDeviceProfiles) | **Get** /networks/{networkId}/sm/users/{userId}/deviceProfiles | Get the profiles associated with a user



## GetNetworkSmDeviceDeviceProfiles

> []InlineResponse20053 GetNetworkSmDeviceDeviceProfiles(ctx, networkId, deviceId).Execute()

Get the installed profiles associated with a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceProfilesApi.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceProfilesApi.GetNetworkSmDeviceDeviceProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceDeviceProfiles`: []InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `DeviceProfilesApi.GetNetworkSmDeviceDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20053**](InlineResponse20053.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserDeviceProfiles

> []InlineResponse20053 GetNetworkSmUserDeviceProfiles(ctx, networkId, userId).Execute()

Get the profiles associated with a user



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
    userId := "userId_example" // string | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceProfilesApi.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceProfilesApi.GetNetworkSmUserDeviceProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmUserDeviceProfiles`: []InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `DeviceProfilesApi.GetNetworkSmUserDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20053**](InlineResponse20053.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

