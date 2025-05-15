# \MacTableApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsMacTable**](MacTableApi.md#CreateDeviceLiveToolsMacTable) | **Post** /devices/{serial}/liveTools/macTable | Enqueue a job to request the MAC table from the device
[**GetDeviceLiveToolsMacTable**](MacTableApi.md#GetDeviceLiveToolsMacTable) | **Get** /devices/{serial}/liveTools/macTable/{macTableId} | Return a MAC table live tool job.



## CreateDeviceLiveToolsMacTable

> InlineResponse2014 CreateDeviceLiveToolsMacTable(ctx, serial).CreateDeviceLiveToolsMacTable(createDeviceLiveToolsMacTable).Execute()

Enqueue a job to request the MAC table from the device



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
    createDeviceLiveToolsMacTable := *openapiclient.NewInlineObject19() // InlineObject19 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacTableApi.CreateDeviceLiveToolsMacTable(context.Background(), serial).CreateDeviceLiveToolsMacTable(createDeviceLiveToolsMacTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacTableApi.CreateDeviceLiveToolsMacTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsMacTable`: InlineResponse2014
    fmt.Fprintf(os.Stdout, "Response from `MacTableApi.CreateDeviceLiveToolsMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsMacTable** | [**InlineObject19**](InlineObject19.md) |  | 

### Return type

[**InlineResponse2014**](InlineResponse2014.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsMacTable

> InlineResponse20025 GetDeviceLiveToolsMacTable(ctx, serial, macTableId).Execute()

Return a MAC table live tool job.



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
    macTableId := "macTableId_example" // string | Mac table ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MacTableApi.GetDeviceLiveToolsMacTable(context.Background(), serial, macTableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MacTableApi.GetDeviceLiveToolsMacTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsMacTable`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `MacTableApi.GetDeviceLiveToolsMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**macTableId** | **string** | Mac table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

