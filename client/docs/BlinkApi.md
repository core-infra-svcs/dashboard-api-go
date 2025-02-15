# \BlinkApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsLedsBlink**](BlinkApi.md#CreateDeviceLiveToolsLedsBlink) | **Post** /devices/{serial}/liveTools/leds/blink | Enqueue a job to blink LEDs on a device
[**GetDeviceLiveToolsLedsBlink**](BlinkApi.md#GetDeviceLiveToolsLedsBlink) | **Get** /devices/{serial}/liveTools/leds/blink/{ledsBlinkId} | Return a blink LEDs job



## CreateDeviceLiveToolsLedsBlink

> InlineResponse2013 CreateDeviceLiveToolsLedsBlink(ctx, serial).CreateDeviceLiveToolsLedsBlink(createDeviceLiveToolsLedsBlink).Execute()

Enqueue a job to blink LEDs on a device



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
    createDeviceLiveToolsLedsBlink := *openapiclient.NewInlineObject18(int32(123)) // InlineObject18 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlinkApi.CreateDeviceLiveToolsLedsBlink(context.Background(), serial).CreateDeviceLiveToolsLedsBlink(createDeviceLiveToolsLedsBlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlinkApi.CreateDeviceLiveToolsLedsBlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsLedsBlink`: InlineResponse2013
    fmt.Fprintf(os.Stdout, "Response from `BlinkApi.CreateDeviceLiveToolsLedsBlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsLedsBlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsLedsBlink** | [**InlineObject18**](InlineObject18.md) |  | 

### Return type

[**InlineResponse2013**](InlineResponse2013.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsLedsBlink

> InlineResponse20024 GetDeviceLiveToolsLedsBlink(ctx, serial, ledsBlinkId).Execute()

Return a blink LEDs job



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
    ledsBlinkId := "ledsBlinkId_example" // string | Leds blink ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlinkApi.GetDeviceLiveToolsLedsBlink(context.Background(), serial, ledsBlinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlinkApi.GetDeviceLiveToolsLedsBlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsLedsBlink`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `BlinkApi.GetDeviceLiveToolsLedsBlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**ledsBlinkId** | **string** | Leds blink ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsLedsBlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

