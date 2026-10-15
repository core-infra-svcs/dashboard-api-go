# \RoutingTableApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsRoutingTableLookup**](RoutingTableApi.md#CreateDeviceLiveToolsRoutingTableLookup) | **Post** /devices/{serial}/liveTools/routingTable/lookups | Enqueue a job to perform a routing table lookup request for a device
[**CreateDeviceLiveToolsRoutingTableSummary**](RoutingTableApi.md#CreateDeviceLiveToolsRoutingTableSummary) | **Post** /devices/{serial}/liveTools/routingTable/summaries | Enqueue a routing table summary job for a device
[**GetDeviceLiveToolsRoutingTableLookup**](RoutingTableApi.md#GetDeviceLiveToolsRoutingTableLookup) | **Get** /devices/{serial}/liveTools/routingTable/lookups/{id} | Return a routing table live tool lookup job for a device
[**GetDeviceLiveToolsRoutingTableSummary**](RoutingTableApi.md#GetDeviceLiveToolsRoutingTableSummary) | **Get** /devices/{serial}/liveTools/routingTable/summaries/{id} | Return the status and result of a routing table summary job



## CreateDeviceLiveToolsRoutingTableLookup

> InlineResponse20111 CreateDeviceLiveToolsRoutingTableLookup(ctx, serial).CreateDeviceLiveToolsRoutingTableLookup(createDeviceLiveToolsRoutingTableLookup).Execute()

Enqueue a job to perform a routing table lookup request for a device



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
    createDeviceLiveToolsRoutingTableLookup := *openapiclient.NewInlineObject30() // InlineObject30 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingTableApi.CreateDeviceLiveToolsRoutingTableLookup(context.Background(), serial).CreateDeviceLiveToolsRoutingTableLookup(createDeviceLiveToolsRoutingTableLookup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingTableApi.CreateDeviceLiveToolsRoutingTableLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsRoutingTableLookup`: InlineResponse20111
    fmt.Fprintf(os.Stdout, "Response from `RoutingTableApi.CreateDeviceLiveToolsRoutingTableLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsRoutingTableLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsRoutingTableLookup** | [**InlineObject30**](InlineObject30.md) |  | 

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


## CreateDeviceLiveToolsRoutingTableSummary

> InlineResponse20112 CreateDeviceLiveToolsRoutingTableSummary(ctx, serial).CreateDeviceLiveToolsRoutingTableSummary(createDeviceLiveToolsRoutingTableSummary).Execute()

Enqueue a routing table summary job for a device



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
    createDeviceLiveToolsRoutingTableSummary := *openapiclient.NewInlineObject31() // InlineObject31 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingTableApi.CreateDeviceLiveToolsRoutingTableSummary(context.Background(), serial).CreateDeviceLiveToolsRoutingTableSummary(createDeviceLiveToolsRoutingTableSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingTableApi.CreateDeviceLiveToolsRoutingTableSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsRoutingTableSummary`: InlineResponse20112
    fmt.Fprintf(os.Stdout, "Response from `RoutingTableApi.CreateDeviceLiveToolsRoutingTableSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsRoutingTableSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsRoutingTableSummary** | [**InlineObject31**](InlineObject31.md) |  | 

### Return type

[**InlineResponse20112**](InlineResponse20112.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsRoutingTableLookup

> InlineResponse20042 GetDeviceLiveToolsRoutingTableLookup(ctx, serial, id).Execute()

Return a routing table live tool lookup job for a device



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
    resp, r, err := apiClient.RoutingTableApi.GetDeviceLiveToolsRoutingTableLookup(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingTableApi.GetDeviceLiveToolsRoutingTableLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsRoutingTableLookup`: InlineResponse20042
    fmt.Fprintf(os.Stdout, "Response from `RoutingTableApi.GetDeviceLiveToolsRoutingTableLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsRoutingTableLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20042**](InlineResponse20042.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsRoutingTableSummary

> InlineResponse20043 GetDeviceLiveToolsRoutingTableSummary(ctx, serial, id).Execute()

Return the status and result of a routing table summary job



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
    resp, r, err := apiClient.RoutingTableApi.GetDeviceLiveToolsRoutingTableSummary(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingTableApi.GetDeviceLiveToolsRoutingTableSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsRoutingTableSummary`: InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `RoutingTableApi.GetDeviceLiveToolsRoutingTableSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsRoutingTableSummaryRequest struct via the builder pattern


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

