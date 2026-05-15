# \CycleApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsPortsCycle**](CycleApi.md#CreateDeviceLiveToolsPortsCycle) | **Post** /devices/{serial}/liveTools/ports/cycle | Enqueue a job to perform a cycle port for the device on the specified ports
[**GetDeviceLiveToolsPortsCycle**](CycleApi.md#GetDeviceLiveToolsPortsCycle) | **Get** /devices/{serial}/liveTools/ports/cycle/{id} | Return a cycle port live tool job.



## CreateDeviceLiveToolsPortsCycle

> InlineResponse2018 CreateDeviceLiveToolsPortsCycle(ctx, serial).CreateDeviceLiveToolsPortsCycle(createDeviceLiveToolsPortsCycle).Execute()

Enqueue a job to perform a cycle port for the device on the specified ports



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
    createDeviceLiveToolsPortsCycle := *openapiclient.NewInlineObject25([]string{"Ports_example"}) // InlineObject25 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CycleApi.CreateDeviceLiveToolsPortsCycle(context.Background(), serial).CreateDeviceLiveToolsPortsCycle(createDeviceLiveToolsPortsCycle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CycleApi.CreateDeviceLiveToolsPortsCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsPortsCycle`: InlineResponse2018
    fmt.Fprintf(os.Stdout, "Response from `CycleApi.CreateDeviceLiveToolsPortsCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPortsCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPortsCycle** | [**InlineObject25**](InlineObject25.md) |  | 

### Return type

[**InlineResponse2018**](InlineResponse2018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPortsCycle

> InlineResponse20038 GetDeviceLiveToolsPortsCycle(ctx, serial, id).Execute()

Return a cycle port live tool job.



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CycleApi.GetDeviceLiveToolsPortsCycle(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CycleApi.GetDeviceLiveToolsPortsCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsPortsCycle`: InlineResponse20038
    fmt.Fprintf(os.Stdout, "Response from `CycleApi.GetDeviceLiveToolsPortsCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPortsCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20038**](InlineResponse20038.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

