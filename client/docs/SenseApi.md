# \SenseApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraSense**](SenseApi.md#GetDeviceCameraSense) | **Get** /devices/{serial}/camera/sense | Returns sense settings for a given camera
[**GetDeviceCameraSenseObjectDetectionModels**](SenseApi.md#GetDeviceCameraSenseObjectDetectionModels) | **Get** /devices/{serial}/camera/sense/objectDetectionModels | Returns the MV Sense object detection model list for the given camera
[**UpdateDeviceCameraSense**](SenseApi.md#UpdateDeviceCameraSense) | **Put** /devices/{serial}/camera/sense | Update sense settings for the given camera



## GetDeviceCameraSense

> InlineResponse20021 GetDeviceCameraSense(ctx, serial).Execute()

Returns sense settings for a given camera



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
    resp, r, err := apiClient.SenseApi.GetDeviceCameraSense(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SenseApi.GetDeviceCameraSense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraSense`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `SenseApi.GetDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraSenseObjectDetectionModels

> []InlineResponse20022 GetDeviceCameraSenseObjectDetectionModels(ctx, serial).Execute()

Returns the MV Sense object detection model list for the given camera



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
    resp, r, err := apiClient.SenseApi.GetDeviceCameraSenseObjectDetectionModels(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SenseApi.GetDeviceCameraSenseObjectDetectionModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraSenseObjectDetectionModels`: []InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `SenseApi.GetDeviceCameraSenseObjectDetectionModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseObjectDetectionModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20022**](InlineResponse20022.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraSense

> InlineResponse20021 UpdateDeviceCameraSense(ctx, serial).UpdateDeviceCameraSense(updateDeviceCameraSense).Execute()

Update sense settings for the given camera



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
    updateDeviceCameraSense := *openapiclient.NewInlineObject12() // InlineObject12 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SenseApi.UpdateDeviceCameraSense(context.Background(), serial).UpdateDeviceCameraSense(updateDeviceCameraSense).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SenseApi.UpdateDeviceCameraSense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraSense`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `SenseApi.UpdateDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraSenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraSense** | [**InlineObject12**](InlineObject12.md) |  | 

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

