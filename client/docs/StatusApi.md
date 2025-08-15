# \StatusApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceWirelessStatus**](StatusApi.md#GetDeviceWirelessStatus) | **Get** /devices/{serial}/wireless/status | Return the SSID statuses of an access point
[**GetOrganizationSpacesIntegrateStatus**](StatusApi.md#GetOrganizationSpacesIntegrateStatus) | **Get** /organizations/{organizationId}/spaces/integrate/status | Get the status of the Spaces integration in Meraki



## GetDeviceWirelessStatus

> InlineResponse20047 GetDeviceWirelessStatus(ctx, serial).Execute()

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
    // response from `GetDeviceWirelessStatus`: InlineResponse20047
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

[**InlineResponse20047**](InlineResponse20047.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSpacesIntegrateStatus

> InlineResponse200324 GetOrganizationSpacesIntegrateStatus(ctx, organizationId).Execute()

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
    // response from `GetOrganizationSpacesIntegrateStatus`: InlineResponse200324
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

[**InlineResponse200324**](InlineResponse200324.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

