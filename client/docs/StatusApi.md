# \StatusApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsPortsStatus**](StatusApi.md#CreateDeviceLiveToolsPortsStatus) | **Post** /devices/{serial}/liveTools/ports/status | Enqueue a job to retrieve port status for a device
[**GetDeviceLiveToolsPortsStatus**](StatusApi.md#GetDeviceLiveToolsPortsStatus) | **Get** /devices/{serial}/liveTools/ports/status/{jobId} | Return a port status live tool job.
[**GetDeviceWirelessStatus**](StatusApi.md#GetDeviceWirelessStatus) | **Get** /devices/{serial}/wireless/status | Return the SSID statuses of an access point
[**GetOrganizationSpacesIntegrateStatus**](StatusApi.md#GetOrganizationSpacesIntegrateStatus) | **Get** /organizations/{organizationId}/spaces/integrate/status | Get the status of the Spaces integration in Meraki



## CreateDeviceLiveToolsPortsStatus

> InlineResponse2019 CreateDeviceLiveToolsPortsStatus(ctx, serial).CreateDeviceLiveToolsPortsStatus(createDeviceLiveToolsPortsStatus).Execute()

Enqueue a job to retrieve port status for a device



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
    createDeviceLiveToolsPortsStatus := *openapiclient.NewInlineObject28() // InlineObject28 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatusApi.CreateDeviceLiveToolsPortsStatus(context.Background(), serial).CreateDeviceLiveToolsPortsStatus(createDeviceLiveToolsPortsStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.CreateDeviceLiveToolsPortsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsPortsStatus`: InlineResponse2019
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.CreateDeviceLiveToolsPortsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPortsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPortsStatus** | [**InlineObject28**](InlineObject28.md) |  | 

### Return type

[**InlineResponse2019**](InlineResponse2019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPortsStatus

> InlineResponse20040 GetDeviceLiveToolsPortsStatus(ctx, serial, jobId).Execute()

Return a port status live tool job.



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
    resp, r, err := apiClient.StatusApi.GetDeviceLiveToolsPortsStatus(context.Background(), serial, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.GetDeviceLiveToolsPortsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsPortsStatus`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.GetDeviceLiveToolsPortsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPortsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessStatus

> InlineResponse20065 GetDeviceWirelessStatus(ctx, serial).Execute()

Return the SSID statuses of an access point



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
    resp, r, err := apiClient.StatusApi.GetDeviceWirelessStatus(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.GetDeviceWirelessStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessStatus`: InlineResponse20065
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.GetDeviceWirelessStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20065**](InlineResponse20065.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSpacesIntegrateStatus

> InlineResponse200453 GetOrganizationSpacesIntegrateStatus(ctx, organizationId).Execute()

Get the status of the Spaces integration in Meraki



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
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatusApi.GetOrganizationSpacesIntegrateStatus(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.GetOrganizationSpacesIntegrateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSpacesIntegrateStatus`: InlineResponse200453
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.GetOrganizationSpacesIntegrateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSpacesIntegrateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200453**](InlineResponse200453.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

