# \MasksApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceCellularUplinksBandsMasksUpdate**](MasksApi.md#CreateDeviceCellularUplinksBandsMasksUpdate) | **Post** /devices/{serial}/cellular/uplinks/bands/masks/update | Update the cellular band masks for a device



## CreateDeviceCellularUplinksBandsMasksUpdate

> InlineResponse20027 CreateDeviceCellularUplinksBandsMasksUpdate(ctx, serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()

Update the cellular band masks for a device



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
    createDeviceCellularUplinksBandsMasksUpdate := *openapiclient.NewInlineObject15("Slot_example", "Type_example", []string{"Masked_example"}) // InlineObject15 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MasksApi.CreateDeviceCellularUplinksBandsMasksUpdate(context.Background(), serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MasksApi.CreateDeviceCellularUplinksBandsMasksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceCellularUplinksBandsMasksUpdate`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `MasksApi.CreateDeviceCellularUplinksBandsMasksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceCellularUplinksBandsMasksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceCellularUplinksBandsMasksUpdate** | [**InlineObject15**](InlineObject15.md) |  | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

