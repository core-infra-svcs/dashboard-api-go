# \DetailsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateOrganizationDevicesDetails**](DetailsApi.md#BulkUpdateOrganizationDevicesDetails) | **Post** /organizations/{organizationId}/devices/details/bulkUpdate | Updating device details (currently only used for Catalyst devices)
[**CreateDeviceLiveToolsRouteDetail**](DetailsApi.md#CreateDeviceLiveToolsRouteDetail) | **Post** /devices/{serial}/liveTools/route/details | Enqueue a live tool job to fetch route details from a device.
[**GetDeviceLiveToolsRouteDetail**](DetailsApi.md#GetDeviceLiveToolsRouteDetail) | **Get** /devices/{serial}/liveTools/route/details/{jobId} | Return a route details live tool job for a device.



## BulkUpdateOrganizationDevicesDetails

> InlineResponse200385 BulkUpdateOrganizationDevicesDetails(ctx, organizationId).BulkUpdateOrganizationDevicesDetails(bulkUpdateOrganizationDevicesDetails).Execute()

Updating device details (currently only used for Catalyst devices)



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
    bulkUpdateOrganizationDevicesDetails := *openapiclient.NewInlineObject308([]string{"Serials_example"}, []openapiclient.NetworksNetworkIdDevicesClaimDetails{*openapiclient.NewNetworksNetworkIdDevicesClaimDetails("Name_example")}) // InlineObject308 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DetailsApi.BulkUpdateOrganizationDevicesDetails(context.Background(), organizationId).BulkUpdateOrganizationDevicesDetails(bulkUpdateOrganizationDevicesDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetailsApi.BulkUpdateOrganizationDevicesDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateOrganizationDevicesDetails`: InlineResponse200385
    fmt.Fprintf(os.Stdout, "Response from `DetailsApi.BulkUpdateOrganizationDevicesDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateOrganizationDevicesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateOrganizationDevicesDetails** | [**InlineObject308**](InlineObject308.md) |  | 

### Return type

[**InlineResponse200385**](InlineResponse200385.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.DetailsApi.CreateDeviceLiveToolsRouteDetail(context.Background(), serial).CreateDeviceLiveToolsRouteDetail(createDeviceLiveToolsRouteDetail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetailsApi.CreateDeviceLiveToolsRouteDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsRouteDetail`: InlineResponse20111
    fmt.Fprintf(os.Stdout, "Response from `DetailsApi.CreateDeviceLiveToolsRouteDetail`: %v\n", resp)
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
    resp, r, err := apiClient.DetailsApi.GetDeviceLiveToolsRouteDetail(context.Background(), serial, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DetailsApi.GetDeviceLiveToolsRouteDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsRouteDetail`: InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `DetailsApi.GetDeviceLiveToolsRouteDetail`: %v\n", resp)
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

