# \LanApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularGatewayLan**](LanApi.md#GetDeviceCellularGatewayLan) | **Get** /devices/{serial}/cellularGateway/lan | Show the LAN Settings of a MG
[**UpdateDeviceCellularGatewayLan**](LanApi.md#UpdateDeviceCellularGatewayLan) | **Put** /devices/{serial}/cellularGateway/lan | Update the LAN Settings for a single MG.



## GetDeviceCellularGatewayLan

> InlineResponse20019 GetDeviceCellularGatewayLan(ctx, serial).Execute()

Show the LAN Settings of a MG



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
    resp, r, err := apiClient.LanApi.GetDeviceCellularGatewayLan(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanApi.GetDeviceCellularGatewayLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularGatewayLan`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `LanApi.GetDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGatewayLan

> InlineResponse20019 UpdateDeviceCellularGatewayLan(ctx, serial).UpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan).Execute()

Update the LAN Settings for a single MG.



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
    updateDeviceCellularGatewayLan := *openapiclient.NewInlineObject14() // InlineObject14 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanApi.UpdateDeviceCellularGatewayLan(context.Background(), serial).UpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanApi.UpdateDeviceCellularGatewayLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularGatewayLan`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `LanApi.UpdateDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayLan** | [**InlineObject14**](InlineObject14.md) |  | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

