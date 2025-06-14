# \ElectronicShelfLabelApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceWirelessElectronicShelfLabel**](ElectronicShelfLabelApi.md#GetDeviceWirelessElectronicShelfLabel) | **Get** /devices/{serial}/wireless/electronicShelfLabel | Return the ESL settings of a device
[**GetNetworkWirelessElectronicShelfLabel**](ElectronicShelfLabelApi.md#GetNetworkWirelessElectronicShelfLabel) | **Get** /networks/{networkId}/wireless/electronicShelfLabel | Return the ESL settings of a wireless network
[**GetNetworkWirelessElectronicShelfLabelConfiguredDevices**](ElectronicShelfLabelApi.md#GetNetworkWirelessElectronicShelfLabelConfiguredDevices) | **Get** /networks/{networkId}/wireless/electronicShelfLabel/configuredDevices | Get a list of all ESL eligible devices of a network
[**UpdateDeviceWirelessElectronicShelfLabel**](ElectronicShelfLabelApi.md#UpdateDeviceWirelessElectronicShelfLabel) | **Put** /devices/{serial}/wireless/electronicShelfLabel | Update the ESL settings of a device
[**UpdateNetworkWirelessElectronicShelfLabel**](ElectronicShelfLabelApi.md#UpdateNetworkWirelessElectronicShelfLabel) | **Put** /networks/{networkId}/wireless/electronicShelfLabel | Update the ESL settings of a wireless network



## GetDeviceWirelessElectronicShelfLabel

> InlineResponse20046 GetDeviceWirelessElectronicShelfLabel(ctx, serial).Execute()

Return the ESL settings of a device



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
    resp, r, err := apiClient.ElectronicShelfLabelApi.GetDeviceWirelessElectronicShelfLabel(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelApi.GetDeviceWirelessElectronicShelfLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessElectronicShelfLabel`: InlineResponse20046
    fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelApi.GetDeviceWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20046**](InlineResponse20046.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessElectronicShelfLabel

> InlineResponse200189 GetNetworkWirelessElectronicShelfLabel(ctx, networkId).Execute()

Return the ESL settings of a wireless network



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
    resp, r, err := apiClient.ElectronicShelfLabelApi.GetNetworkWirelessElectronicShelfLabel(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelApi.GetNetworkWirelessElectronicShelfLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessElectronicShelfLabel`: InlineResponse200189
    fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelApi.GetNetworkWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200189**](InlineResponse200189.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessElectronicShelfLabelConfiguredDevices

> []InlineResponse200189 GetNetworkWirelessElectronicShelfLabelConfiguredDevices(ctx, networkId).Execute()

Get a list of all ESL eligible devices of a network



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
    resp, r, err := apiClient.ElectronicShelfLabelApi.GetNetworkWirelessElectronicShelfLabelConfiguredDevices(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelApi.GetNetworkWirelessElectronicShelfLabelConfiguredDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessElectronicShelfLabelConfiguredDevices`: []InlineResponse200189
    fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelApi.GetNetworkWirelessElectronicShelfLabelConfiguredDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessElectronicShelfLabelConfiguredDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200189**](InlineResponse200189.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessElectronicShelfLabel

> InlineResponse20046 UpdateDeviceWirelessElectronicShelfLabel(ctx, serial).UpdateDeviceWirelessElectronicShelfLabel(updateDeviceWirelessElectronicShelfLabel).Execute()

Update the ESL settings of a device



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
    updateDeviceWirelessElectronicShelfLabel := *openapiclient.NewInlineObject37() // InlineObject37 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElectronicShelfLabelApi.UpdateDeviceWirelessElectronicShelfLabel(context.Background(), serial).UpdateDeviceWirelessElectronicShelfLabel(updateDeviceWirelessElectronicShelfLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelApi.UpdateDeviceWirelessElectronicShelfLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessElectronicShelfLabel`: InlineResponse20046
    fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelApi.UpdateDeviceWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessElectronicShelfLabel** | [**InlineObject37**](InlineObject37.md) |  | 

### Return type

[**InlineResponse20046**](InlineResponse20046.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessElectronicShelfLabel

> InlineResponse200189 UpdateNetworkWirelessElectronicShelfLabel(ctx, networkId).UpdateNetworkWirelessElectronicShelfLabel(updateNetworkWirelessElectronicShelfLabel).Execute()

Update the ESL settings of a wireless network



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
    updateNetworkWirelessElectronicShelfLabel := *openapiclient.NewInlineObject185() // InlineObject185 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ElectronicShelfLabelApi.UpdateNetworkWirelessElectronicShelfLabel(context.Background(), networkId).UpdateNetworkWirelessElectronicShelfLabel(updateNetworkWirelessElectronicShelfLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelApi.UpdateNetworkWirelessElectronicShelfLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessElectronicShelfLabel`: InlineResponse200189
    fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelApi.UpdateNetworkWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessElectronicShelfLabel** | [**InlineObject185**](InlineObject185.md) |  | 

### Return type

[**InlineResponse200189**](InlineResponse200189.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

