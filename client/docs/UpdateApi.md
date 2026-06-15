# \UpdateApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceApplianceInterfacesPortsUpdate**](UpdateApi.md#CreateDeviceApplianceInterfacesPortsUpdate) | **Post** /devices/{serial}/appliance/interfaces/ports/update | Update configurations for an appliance&#39;s specified port
[**CreateDeviceCellularUplinksBandsMasksUpdate**](UpdateApi.md#CreateDeviceCellularUplinksBandsMasksUpdate) | **Post** /devices/{serial}/cellular/uplinks/bands/masks/update | Update the cellular band masks for a device



## CreateDeviceApplianceInterfacesPortsUpdate

> InlineResponse2008 CreateDeviceApplianceInterfacesPortsUpdate(ctx, serial).CreateDeviceApplianceInterfacesPortsUpdate(createDeviceApplianceInterfacesPortsUpdate).Execute()

Update configurations for an appliance's specified port



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
    createDeviceApplianceInterfacesPortsUpdate := *openapiclient.NewInlineObject4() // InlineObject4 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateApi.CreateDeviceApplianceInterfacesPortsUpdate(context.Background(), serial).CreateDeviceApplianceInterfacesPortsUpdate(createDeviceApplianceInterfacesPortsUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateApi.CreateDeviceApplianceInterfacesPortsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceApplianceInterfacesPortsUpdate`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `UpdateApi.CreateDeviceApplianceInterfacesPortsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceApplianceInterfacesPortsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceApplianceInterfacesPortsUpdate** | [**InlineObject4**](InlineObject4.md) |  | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceCellularUplinksBandsMasksUpdate

> InlineResponse20028 CreateDeviceCellularUplinksBandsMasksUpdate(ctx, serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()

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
    createDeviceCellularUplinksBandsMasksUpdate := *openapiclient.NewInlineObject16("Slot_example", "Type_example", []string{"Masked_example"}) // InlineObject16 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateApi.CreateDeviceCellularUplinksBandsMasksUpdate(context.Background(), serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateApi.CreateDeviceCellularUplinksBandsMasksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceCellularUplinksBandsMasksUpdate`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `UpdateApi.CreateDeviceCellularUplinksBandsMasksUpdate`: %v\n", resp)
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

 **createDeviceCellularUplinksBandsMasksUpdate** | [**InlineObject16**](InlineObject16.md) |  | 

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

