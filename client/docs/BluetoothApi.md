# \BluetoothApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceWirelessBluetoothSettings**](BluetoothApi.md#GetDeviceWirelessBluetoothSettings) | **Get** /devices/{serial}/wireless/bluetooth/settings | Return the bluetooth settings for a wireless device
[**GetNetworkWirelessBluetoothSettings**](BluetoothApi.md#GetNetworkWirelessBluetoothSettings) | **Get** /networks/{networkId}/wireless/bluetooth/settings | Return the Bluetooth settings for a network. &lt;a href&#x3D;\&quot;https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\&quot;&gt;Bluetooth settings&lt;/a&gt; must be enabled on the network.
[**UpdateDeviceWirelessBluetoothSettings**](BluetoothApi.md#UpdateDeviceWirelessBluetoothSettings) | **Put** /devices/{serial}/wireless/bluetooth/settings | Update the bluetooth settings for a wireless device
[**UpdateNetworkWirelessBluetoothSettings**](BluetoothApi.md#UpdateNetworkWirelessBluetoothSettings) | **Put** /networks/{networkId}/wireless/bluetooth/settings | Update the Bluetooth settings for a network



## GetDeviceWirelessBluetoothSettings

> map[string]interface{} GetDeviceWirelessBluetoothSettings(ctx, serial).Execute()

Return the bluetooth settings for a wireless device



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
    serial := "serial_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BluetoothApi.GetDeviceWirelessBluetoothSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.GetDeviceWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessBluetoothSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.GetDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


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


## GetNetworkWirelessBluetoothSettings

> map[string]interface{} GetNetworkWirelessBluetoothSettings(ctx, networkId).Execute()

Return the Bluetooth settings for a network. <a href=\"https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\">Bluetooth settings</a> must be enabled on the network.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BluetoothApi.GetNetworkWirelessBluetoothSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.GetNetworkWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessBluetoothSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.GetNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


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


## UpdateDeviceWirelessBluetoothSettings

> map[string]interface{} UpdateDeviceWirelessBluetoothSettings(ctx, serial).UpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings).Execute()

Update the bluetooth settings for a wireless device



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
    serial := "serial_example" // string | 
    updateDeviceWirelessBluetoothSettings := *openapiclient.NewInlineObject21() // InlineObject21 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BluetoothApi.UpdateDeviceWirelessBluetoothSettings(context.Background(), serial).UpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.UpdateDeviceWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessBluetoothSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.UpdateDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessBluetoothSettings** | [**InlineObject21**](InlineObject21.md) |  | 

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


## UpdateNetworkWirelessBluetoothSettings

> map[string]interface{} UpdateNetworkWirelessBluetoothSettings(ctx, networkId).UpdateNetworkWirelessBluetoothSettings(updateNetworkWirelessBluetoothSettings).Execute()

Update the Bluetooth settings for a network



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
    updateNetworkWirelessBluetoothSettings := *openapiclient.NewInlineObject131() // InlineObject131 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BluetoothApi.UpdateNetworkWirelessBluetoothSettings(context.Background(), networkId).UpdateNetworkWirelessBluetoothSettings(updateNetworkWirelessBluetoothSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.UpdateNetworkWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessBluetoothSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.UpdateNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessBluetoothSettings** | [**InlineObject131**](InlineObject131.md) |  | 

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

