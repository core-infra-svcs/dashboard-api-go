# \SchedulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationDevicesPacketCaptureSchedule**](SchedulesApi.md#CreateOrganizationDevicesPacketCaptureSchedule) | **Post** /organizations/{organizationId}/devices/packetCapture/schedules | Create a schedule for packet capture
[**DeleteOrganizationDevicesPacketCaptureSchedule**](SchedulesApi.md#DeleteOrganizationDevicesPacketCaptureSchedule) | **Delete** /organizations/{organizationId}/devices/packetCapture/schedules/{scheduleId} | Delete schedule from cloud
[**GetNetworkCameraSchedules**](SchedulesApi.md#GetNetworkCameraSchedules) | **Get** /networks/{networkId}/camera/schedules | Returns a list of all camera recording schedules.
[**GetNetworkWirelessSsidSchedules**](SchedulesApi.md#GetNetworkWirelessSsidSchedules) | **Get** /networks/{networkId}/wireless/ssids/{number}/schedules | List the outage schedule for the SSID
[**GetOrganizationDevicesPacketCaptureSchedules**](SchedulesApi.md#GetOrganizationDevicesPacketCaptureSchedules) | **Get** /organizations/{organizationId}/devices/packetCapture/schedules | List the Packet Capture Schedules
[**ReorderOrganizationDevicesPacketCaptureSchedules**](SchedulesApi.md#ReorderOrganizationDevicesPacketCaptureSchedules) | **Post** /organizations/{organizationId}/devices/packetCapture/schedules/reorder | Bulk update priorities of pcap schedules
[**UpdateNetworkWirelessSsidSchedules**](SchedulesApi.md#UpdateNetworkWirelessSsidSchedules) | **Put** /networks/{networkId}/wireless/ssids/{number}/schedules | Update the outage schedule for the SSID
[**UpdateOrganizationDevicesPacketCaptureSchedule**](SchedulesApi.md#UpdateOrganizationDevicesPacketCaptureSchedule) | **Put** /organizations/{organizationId}/devices/packetCapture/schedules/{scheduleId} | Update a schedule for packet capture



## CreateOrganizationDevicesPacketCaptureSchedule

> InlineResponse200275Items CreateOrganizationDevicesPacketCaptureSchedule(ctx, organizationId).CreateOrganizationDevicesPacketCaptureSchedule(createOrganizationDevicesPacketCaptureSchedule).Execute()

Create a schedule for packet capture



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
    createOrganizationDevicesPacketCaptureSchedule := *openapiclient.NewInlineObject260([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices()}) // InlineObject260 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.CreateOrganizationDevicesPacketCaptureSchedule(context.Background(), organizationId).CreateOrganizationDevicesPacketCaptureSchedule(createOrganizationDevicesPacketCaptureSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.CreateOrganizationDevicesPacketCaptureSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDevicesPacketCaptureSchedule`: InlineResponse200275Items
    fmt.Fprintf(os.Stdout, "Response from `SchedulesApi.CreateOrganizationDevicesPacketCaptureSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDevicesPacketCaptureScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationDevicesPacketCaptureSchedule** | [**InlineObject260**](InlineObject260.md) |  | 

### Return type

[**InlineResponse200275Items**](InlineResponse200275Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationDevicesPacketCaptureSchedule

> DeleteOrganizationDevicesPacketCaptureSchedule(ctx, organizationId, scheduleId).DeleteOrganizationDevicesPacketCaptureSchedule(deleteOrganizationDevicesPacketCaptureSchedule).Execute()

Delete schedule from cloud



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
    scheduleId := "scheduleId_example" // string | Schedule ID
    deleteOrganizationDevicesPacketCaptureSchedule := *openapiclient.NewInlineObject263("ScheduleId_example") // InlineObject263 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.DeleteOrganizationDevicesPacketCaptureSchedule(context.Background(), organizationId, scheduleId).DeleteOrganizationDevicesPacketCaptureSchedule(deleteOrganizationDevicesPacketCaptureSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.DeleteOrganizationDevicesPacketCaptureSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**scheduleId** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationDevicesPacketCaptureScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteOrganizationDevicesPacketCaptureSchedule** | [**InlineObject263**](InlineObject263.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCameraSchedules

> []InlineResponse20080 GetNetworkCameraSchedules(ctx, networkId).Execute()

Returns a list of all camera recording schedules.



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.GetNetworkCameraSchedules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.GetNetworkCameraSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraSchedules`: []InlineResponse20080
    fmt.Fprintf(os.Stdout, "Response from `SchedulesApi.GetNetworkCameraSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20080**](InlineResponse20080.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidSchedules

> InlineResponse200205 GetNetworkWirelessSsidSchedules(ctx, networkId, number).Execute()

List the outage schedule for the SSID



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
    networkId := "networkId_example" // string | Network ID
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.GetNetworkWirelessSsidSchedules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.GetNetworkWirelessSsidSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidSchedules`: InlineResponse200205
    fmt.Fprintf(os.Stdout, "Response from `SchedulesApi.GetNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200205**](InlineResponse200205.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesPacketCaptureSchedules

> InlineResponse200275 GetOrganizationDevicesPacketCaptureSchedules(ctx, organizationId).ScheduleIds(scheduleIds).NetworkIds(networkIds).DeviceIds(deviceIds).Execute()

List the Packet Capture Schedules



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
    scheduleIds := []string{"Inner_example"} // []string | Return the packet captures schedules of the specified packet capture schedule ids (optional)
    networkIds := []string{"Inner_example"} // []string | Return the scheduled packet captures of the specified network(s) (optional)
    deviceIds := []string{"Inner_example"} // []string | Return the scheduled packet captures of the specified device(s) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.GetOrganizationDevicesPacketCaptureSchedules(context.Background(), organizationId).ScheduleIds(scheduleIds).NetworkIds(networkIds).DeviceIds(deviceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.GetOrganizationDevicesPacketCaptureSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesPacketCaptureSchedules`: InlineResponse200275
    fmt.Fprintf(os.Stdout, "Response from `SchedulesApi.GetOrganizationDevicesPacketCaptureSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesPacketCaptureSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleIds** | **[]string** | Return the packet captures schedules of the specified packet capture schedule ids | 
 **networkIds** | **[]string** | Return the scheduled packet captures of the specified network(s) | 
 **deviceIds** | **[]string** | Return the scheduled packet captures of the specified device(s) | 

### Return type

[**InlineResponse200275**](InlineResponse200275.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderOrganizationDevicesPacketCaptureSchedules

> InlineResponse200276 ReorderOrganizationDevicesPacketCaptureSchedules(ctx, organizationId).ReorderOrganizationDevicesPacketCaptureSchedules(reorderOrganizationDevicesPacketCaptureSchedules).Execute()

Bulk update priorities of pcap schedules



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
    reorderOrganizationDevicesPacketCaptureSchedules := *openapiclient.NewInlineObject261([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureSchedulesReorderOrder{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesReorderOrder()}) // InlineObject261 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.ReorderOrganizationDevicesPacketCaptureSchedules(context.Background(), organizationId).ReorderOrganizationDevicesPacketCaptureSchedules(reorderOrganizationDevicesPacketCaptureSchedules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.ReorderOrganizationDevicesPacketCaptureSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReorderOrganizationDevicesPacketCaptureSchedules`: InlineResponse200276
    fmt.Fprintf(os.Stdout, "Response from `SchedulesApi.ReorderOrganizationDevicesPacketCaptureSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderOrganizationDevicesPacketCaptureSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reorderOrganizationDevicesPacketCaptureSchedules** | [**InlineObject261**](InlineObject261.md) |  | 

### Return type

[**InlineResponse200276**](InlineResponse200276.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidSchedules

> InlineResponse200205 UpdateNetworkWirelessSsidSchedules(ctx, networkId, number).UpdateNetworkWirelessSsidSchedules(updateNetworkWirelessSsidSchedules).Execute()

Update the outage schedule for the SSID



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
    networkId := "networkId_example" // string | Network ID
    number := "number_example" // string | Number
    updateNetworkWirelessSsidSchedules := *openapiclient.NewInlineObject203() // InlineObject203 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.UpdateNetworkWirelessSsidSchedules(context.Background(), networkId, number).UpdateNetworkWirelessSsidSchedules(updateNetworkWirelessSsidSchedules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.UpdateNetworkWirelessSsidSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidSchedules`: InlineResponse200205
    fmt.Fprintf(os.Stdout, "Response from `SchedulesApi.UpdateNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidSchedules** | [**InlineObject203**](InlineObject203.md) |  | 

### Return type

[**InlineResponse200205**](InlineResponse200205.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationDevicesPacketCaptureSchedule

> InlineResponse200275Items UpdateOrganizationDevicesPacketCaptureSchedule(ctx, organizationId, scheduleId).UpdateOrganizationDevicesPacketCaptureSchedule(updateOrganizationDevicesPacketCaptureSchedule).Execute()

Update a schedule for packet capture



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
    scheduleId := "scheduleId_example" // string | Schedule ID
    updateOrganizationDevicesPacketCaptureSchedule := *openapiclient.NewInlineObject262([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureSchedulesDevices()}) // InlineObject262 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesApi.UpdateOrganizationDevicesPacketCaptureSchedule(context.Background(), organizationId, scheduleId).UpdateOrganizationDevicesPacketCaptureSchedule(updateOrganizationDevicesPacketCaptureSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesApi.UpdateOrganizationDevicesPacketCaptureSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationDevicesPacketCaptureSchedule`: InlineResponse200275Items
    fmt.Fprintf(os.Stdout, "Response from `SchedulesApi.UpdateOrganizationDevicesPacketCaptureSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**scheduleId** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationDevicesPacketCaptureScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationDevicesPacketCaptureSchedule** | [**InlineObject262**](InlineObject262.md) |  | 

### Return type

[**InlineResponse200275Items**](InlineResponse200275Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

