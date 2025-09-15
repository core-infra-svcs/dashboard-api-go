# \MulticastRoutingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsMulticastRouting**](MulticastRoutingApi.md#CreateDeviceLiveToolsMulticastRouting) | **Post** /devices/{serial}/liveTools/multicastRouting | Enqueue a job to perform a Multicast routing request for the device
[**GetDeviceLiveToolsMulticastRouting**](MulticastRoutingApi.md#GetDeviceLiveToolsMulticastRouting) | **Get** /devices/{serial}/liveTools/multicastRouting/{multicastRoutingId} | Return a Multicast routing live tool job.



## CreateDeviceLiveToolsMulticastRouting

> InlineResponse2015 CreateDeviceLiveToolsMulticastRouting(ctx, serial).CreateDeviceLiveToolsMulticastRouting(createDeviceLiveToolsMulticastRouting).Execute()

Enqueue a job to perform a Multicast routing request for the device



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
    createDeviceLiveToolsMulticastRouting := *openapiclient.NewInlineObject20() // InlineObject20 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MulticastRoutingApi.CreateDeviceLiveToolsMulticastRouting(context.Background(), serial).CreateDeviceLiveToolsMulticastRouting(createDeviceLiveToolsMulticastRouting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MulticastRoutingApi.CreateDeviceLiveToolsMulticastRouting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsMulticastRouting`: InlineResponse2015
    fmt.Fprintf(os.Stdout, "Response from `MulticastRoutingApi.CreateDeviceLiveToolsMulticastRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsMulticastRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsMulticastRouting** | [**InlineObject20**](InlineObject20.md) |  | 

### Return type

[**InlineResponse2015**](InlineResponse2015.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsMulticastRouting

> InlineResponse20026 GetDeviceLiveToolsMulticastRouting(ctx, serial, multicastRoutingId).Execute()

Return a Multicast routing live tool job.



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
    multicastRoutingId := "multicastRoutingId_example" // string | Multicast routing ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MulticastRoutingApi.GetDeviceLiveToolsMulticastRouting(context.Background(), serial, multicastRoutingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MulticastRoutingApi.GetDeviceLiveToolsMulticastRouting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsMulticastRouting`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `MulticastRoutingApi.GetDeviceLiveToolsMulticastRouting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**multicastRoutingId** | **string** | Multicast routing ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsMulticastRoutingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

