# \RouteApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsRouteDetail**](RouteApi.md#CreateDeviceLiveToolsRouteDetail) | **Post** /devices/{serial}/liveTools/route/details | Enqueue a live tool job to fetch route details from a device.
[**GetDeviceLiveToolsRouteDetail**](RouteApi.md#GetDeviceLiveToolsRouteDetail) | **Get** /devices/{serial}/liveTools/route/details/{jobId} | Return a route details live tool job for a device.



## CreateDeviceLiveToolsRouteDetail

> InlineResponse20111 CreateDeviceLiveToolsRouteDetail(ctx, serial).CreateDeviceLiveToolsRouteDetail(createDeviceLiveToolsRouteDetail).Execute()

Enqueue a live tool job to fetch route details from a device.



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
    createDeviceLiveToolsRouteDetail := *openapiclient.NewInlineObject31("Type_example", *openapiclient.NewDevicesSerialLiveToolsRouteDetailsVrf("Name_example"), *openapiclient.NewDevicesSerialLiveToolsRouteDetailsDestination("Subnet_example"), *openapiclient.NewDevicesSerialLiveToolsRouteDetailsNextHop("Address_example")) // InlineObject31 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RouteApi.CreateDeviceLiveToolsRouteDetail(context.Background(), serial).CreateDeviceLiveToolsRouteDetail(createDeviceLiveToolsRouteDetail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.CreateDeviceLiveToolsRouteDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsRouteDetail`: InlineResponse20111
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.CreateDeviceLiveToolsRouteDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsRouteDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsRouteDetail** | [**InlineObject31**](InlineObject31.md) |  | 

### Return type

[**InlineResponse20111**](InlineResponse20111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsRouteDetail

> InlineResponse20043 GetDeviceLiveToolsRouteDetail(ctx, serial, jobId).Execute()

Return a route details live tool job for a device.



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
    jobId := "jobId_example" // string | Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RouteApi.GetDeviceLiveToolsRouteDetail(context.Background(), serial, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RouteApi.GetDeviceLiveToolsRouteDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsRouteDetail`: InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `RouteApi.GetDeviceLiveToolsRouteDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsRouteDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20043**](InlineResponse20043.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

