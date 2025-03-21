# \SimsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularSims**](SimsApi.md#GetDeviceCellularSims) | **Get** /devices/{serial}/cellular/sims | Return the SIM and APN configurations for a cellular device.
[**UpdateDeviceCellularSims**](SimsApi.md#UpdateDeviceCellularSims) | **Put** /devices/{serial}/cellular/sims | Updates the SIM and APN configurations for a cellular device.



## GetDeviceCellularSims

> InlineResponse20018 GetDeviceCellularSims(ctx, serial).Execute()

Return the SIM and APN configurations for a cellular device.



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
    resp, r, err := apiClient.SimsApi.GetDeviceCellularSims(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimsApi.GetDeviceCellularSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularSims`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `SimsApi.GetDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularSims

> InlineResponse20018 UpdateDeviceCellularSims(ctx, serial).UpdateDeviceCellularSims(updateDeviceCellularSims).Execute()

Updates the SIM and APN configurations for a cellular device.



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
    updateDeviceCellularSims := *openapiclient.NewInlineObject13() // InlineObject13 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimsApi.UpdateDeviceCellularSims(context.Background(), serial).UpdateDeviceCellularSims(updateDeviceCellularSims).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimsApi.UpdateDeviceCellularSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularSims`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `SimsApi.UpdateDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularSims** | [**InlineObject13**](InlineObject13.md) |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

