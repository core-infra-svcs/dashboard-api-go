# \UsageApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsPowerUsage**](UsageApi.md#CreateDeviceLiveToolsPowerUsage) | **Post** /devices/{serial}/liveTools/power/usage | Enqueues a live tool job that retrieves details about a device&#39;s overall power usage
[**GetDeviceLiveToolsPowerUsage**](UsageApi.md#GetDeviceLiveToolsPowerUsage) | **Get** /devices/{serial}/liveTools/power/usage/{jobId} | Retrieve the status and results of a previously created live tool job fetching details about a device&#39;s overall power usage.
[**GetOrganizationApplianceUplinksUsageByNetwork**](UsageApi.md#GetOrganizationApplianceUplinksUsageByNetwork) | **Get** /organizations/{organizationId}/appliance/uplinks/usage/byNetwork | Get the sent and received bytes for each uplink of all MX and Z networks within an organization
[**GetOrganizationCampusGatewayClientsUsageByNetworkByCluster**](UsageApi.md#GetOrganizationCampusGatewayClientsUsageByNetworkByCluster) | **Get** /organizations/{organizationId}/campusGateway/clients/usage/byNetwork/byCluster | Returns client usage details for campus gateway clusters within an organization.
[**GetOrganizationDevicesCellularDataUsageByDevice**](UsageApi.md#GetOrganizationDevicesCellularDataUsageByDevice) | **Get** /organizations/{organizationId}/devices/cellular/data/usage/byDevice | List current cellular data usage for devices in this organization
[**GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval**](UsageApi.md#GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/devices/cellular/data/usage/history/byDevice/byInterval | List historical cellular data usage grouped by device and interval in this organization
[**GetOrganizationDevicesSystemMemoryUsageHistoryByInterval**](UsageApi.md#GetOrganizationDevicesSystemMemoryUsageHistoryByInterval) | **Get** /organizations/{organizationId}/devices/system/memory/usage/history/byInterval | Return the memory utilization history in kB for devices in the organization.
[**GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval**](UsageApi.md#GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/switch/ports/usage/history/byDevice/byInterval | List the historical usage and traffic data of switchports in an organization.
[**GetOrganizationWirelessClientsUsageByNetwork**](UsageApi.md#GetOrganizationWirelessClientsUsageByNetwork) | **Get** /organizations/{organizationId}/wireless/clients/usage/byNetwork | Returns client usage details for wireless networks within an organization.
[**GetOrganizationWirelessClientsUsageByNetworkBySsid**](UsageApi.md#GetOrganizationWirelessClientsUsageByNetworkBySsid) | **Get** /organizations/{organizationId}/wireless/clients/usage/byNetwork/bySsid | Returns client usage details for wireless network SSIDs within an organization.
[**GetOrganizationWirelessClientsUsageBySsid**](UsageApi.md#GetOrganizationWirelessClientsUsageBySsid) | **Get** /organizations/{organizationId}/wireless/clients/usage/bySsid | Returns client usage details for SSIDs within an organization.
[**GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval**](UsageApi.md#GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l2/usage/history/byInterval | List wireless LAN controller layer 2 interfaces history usage in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval**](UsageApi.md#GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/l3/usage/history/byInterval | List wireless LAN controller layer 3 interfaces history usage in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval**](UsageApi.md#GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/usage/history/byInterval | Retrieve the traffic for the interfaces of a Wireless LAN controller



## CreateDeviceLiveToolsPowerUsage

> InlineResponse20110 CreateDeviceLiveToolsPowerUsage(ctx, serial).CreateDeviceLiveToolsPowerUsage(createDeviceLiveToolsPowerUsage).Execute()

Enqueues a live tool job that retrieves details about a device's overall power usage



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
    createDeviceLiveToolsPowerUsage := *openapiclient.NewInlineObject28() // InlineObject28 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.CreateDeviceLiveToolsPowerUsage(context.Background(), serial).CreateDeviceLiveToolsPowerUsage(createDeviceLiveToolsPowerUsage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.CreateDeviceLiveToolsPowerUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsPowerUsage`: InlineResponse20110
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.CreateDeviceLiveToolsPowerUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPowerUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPowerUsage** | [**InlineObject28**](InlineObject28.md) |  | 

### Return type

[**InlineResponse20110**](InlineResponse20110.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPowerUsage

> InlineResponse20041 GetDeviceLiveToolsPowerUsage(ctx, serial, jobId).Execute()

Retrieve the status and results of a previously created live tool job fetching details about a device's overall power usage.



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
    resp, r, err := apiClient.UsageApi.GetDeviceLiveToolsPowerUsage(context.Background(), serial, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetDeviceLiveToolsPowerUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsPowerUsage`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetDeviceLiveToolsPowerUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPowerUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinksUsageByNetwork

> []InlineResponse200291 GetOrganizationApplianceUplinksUsageByNetwork(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Get the sent and received bytes for each uplink of all MX and Z networks within an organization



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationApplianceUplinksUsageByNetwork(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationApplianceUplinksUsageByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinksUsageByNetwork`: []InlineResponse200291
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationApplianceUplinksUsageByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksUsageByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200291**](InlineResponse200291.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClientsUsageByNetworkByCluster

> InlineResponse200316 GetOrganizationCampusGatewayClientsUsageByNetworkByCluster(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).ClusterIds(clusterIds).UsageUnits(usageUnits).Execute()

Returns client usage details for campus gateway clusters within an organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. (optional)
    networkIds := []string{"Inner_example"} // []string | Filter results by a list of network IDs. (optional)
    networkGroupIds := []string{"Inner_example"} // []string | Limit the results to clients that belong to one of the provided network groups. (optional)
    clusterIds := []string{"Inner_example"} // []string | Filter results by a list of cluster IDs. (optional)
    usageUnits := "usageUnits_example" // string | Usage units to use in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).ClusterIds(clusterIds).UsageUnits(usageUnits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClientsUsageByNetworkByCluster`: InlineResponse200316
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClientsUsageByNetworkByClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. | 
 **networkIds** | **[]string** | Filter results by a list of network IDs. | 
 **networkGroupIds** | **[]string** | Limit the results to clients that belong to one of the provided network groups. | 
 **clusterIds** | **[]string** | Filter results by a list of cluster IDs. | 
 **usageUnits** | **string** | Usage units to use in the response. | 

### Return type

[**InlineResponse200316**](InlineResponse200316.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataUsageByDevice

> InlineResponse200349 GetOrganizationDevicesCellularDataUsageByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List current cellular data usage for devices in this organization



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
    serials := []string{"Inner_example"} // []string | Filter the results by device serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationDevicesCellularDataUsageByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationDevicesCellularDataUsageByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataUsageByDevice`: InlineResponse200349
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationDevicesCellularDataUsageByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataUsageByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Filter the results by device serials. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200349**](InlineResponse200349.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval

> InlineResponse200350 GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

List historical cellular data usage grouped by device and interval in this organization



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
    serials := []string{"Inner_example"} // []string | Required parameter to filter the results by device serials. Maximum 10 serials.
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 10. Default is 5. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 366 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 86400. Interval is calculated if time params are provided. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval`: InlineResponse200350
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataUsageHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Required parameter to filter the results by device serials. Maximum 10 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 10. Default is 5. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 366 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 86400. Interval is calculated if time params are provided. | 

### Return type

[**InlineResponse200350**](InlineResponse200350.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesSystemMemoryUsageHistoryByInterval

> InlineResponse200367 GetOrganizationDevicesSystemMemoryUsageHistoryByInterval(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).NetworkIds(networkIds).Serials(serials).ProductTypes(productTypes).Execute()

Return the memory utilization history in kB for devices in the organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 2 hours. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 3600, 14400. The default is 300. Interval is calculated if time params are provided. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the result set by the included set of network IDs (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device serial numbers (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).NetworkIds(networkIds).Serials(serials).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesSystemMemoryUsageHistoryByInterval`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationDevicesSystemMemoryUsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesSystemMemoryUsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 20. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 2 hours. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 3600, 14400. The default is 300. Interval is calculated if time params are provided. | 
 **networkIds** | **[]string** | Optional parameter to filter the result set by the included set of network IDs | 
 **serials** | **[]string** | Optional parameter to filter device availabilities history by device serial numbers | 
 **productTypes** | **[]string** | Optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | 

### Return type

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval

> InlineResponse200453 GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()

List the historical usage and traffic data of switchports in an organization.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 1200. Interval is calculated if time params are provided. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    configurationUpdatedAfter := time.Now() // time.Time | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. (optional)
    mac := "mac_example" // string | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided MAC addresses. (optional)
    name := "name_example" // string | Optional parameter to filter items to switches with names that contain the search term or are an exact match. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches in one of the provided networks. (optional)
    portProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. (optional)
    serial := "serial_example" // string | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided serials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval`: InlineResponse200453
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationSwitchPortsUsageHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsUsageHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 1200. Interval is calculated if time params are provided. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **time.Time** | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. | 
 **mac** | **string** | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. | 
 **macs** | **[]string** | Optional parameter to filter items to switches that have one of the provided MAC addresses. | 
 **name** | **string** | Optional parameter to filter items to switches with names that contain the search term or are an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter items to switches in one of the provided networks. | 
 **portProfileIds** | **[]string** | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. | 
 **serial** | **string** | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. | 
 **serials** | **[]string** | Optional parameter to filter items to switches that have one of the provided serials. | 

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


## GetOrganizationWirelessClientsUsageByNetwork

> InlineResponse200462 GetOrganizationWirelessClientsUsageByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).GatewayNetworkIds(gatewayNetworkIds).UsageUnits(usageUnits).Execute()

Returns client usage details for wireless networks within an organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. (optional)
    networkIds := []string{"Inner_example"} // []string | Filter results by a list of network IDs. (optional)
    networkGroupIds := []string{"Inner_example"} // []string | Filter results by a list of network group IDs. (optional)
    gatewayNetworkIds := []string{"Inner_example"} // []string | Limit the results to clients tunneled to campus gateways in the provided networks. (optional)
    usageUnits := "usageUnits_example" // string | Usage units to use in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationWirelessClientsUsageByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).GatewayNetworkIds(gatewayNetworkIds).UsageUnits(usageUnits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationWirelessClientsUsageByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessClientsUsageByNetwork`: InlineResponse200462
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationWirelessClientsUsageByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessClientsUsageByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. | 
 **networkIds** | **[]string** | Filter results by a list of network IDs. | 
 **networkGroupIds** | **[]string** | Filter results by a list of network group IDs. | 
 **gatewayNetworkIds** | **[]string** | Limit the results to clients tunneled to campus gateways in the provided networks. | 
 **usageUnits** | **string** | Usage units to use in the response. | 

### Return type

[**InlineResponse200462**](InlineResponse200462.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessClientsUsageByNetworkBySsid

> InlineResponse200463 GetOrganizationWirelessClientsUsageByNetworkBySsid(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).SsidIds(ssidIds).SsidNames(ssidNames).GatewayNetworkIds(gatewayNetworkIds).UsageUnits(usageUnits).Execute()

Returns client usage details for wireless network SSIDs within an organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. (optional)
    networkIds := []string{"Inner_example"} // []string | Filter results by a list of network IDs. (optional)
    networkGroupIds := []string{"Inner_example"} // []string | Filter results by a list of network group IDs. (optional)
    ssidIds := []string{"Inner_example"} // []string | Filter results by a list of SSID IDs. (optional)
    ssidNames := []string{"Inner_example"} // []string | Filter results by a list of SSID names. (optional)
    gatewayNetworkIds := []string{"Inner_example"} // []string | Limit the results to clients tunneled to campus gateways in the provided networks. (optional)
    usageUnits := "usageUnits_example" // string | Usage units to use in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationWirelessClientsUsageByNetworkBySsid(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).SsidIds(ssidIds).SsidNames(ssidNames).GatewayNetworkIds(gatewayNetworkIds).UsageUnits(usageUnits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationWirelessClientsUsageByNetworkBySsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessClientsUsageByNetworkBySsid`: InlineResponse200463
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationWirelessClientsUsageByNetworkBySsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessClientsUsageByNetworkBySsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. | 
 **networkIds** | **[]string** | Filter results by a list of network IDs. | 
 **networkGroupIds** | **[]string** | Filter results by a list of network group IDs. | 
 **ssidIds** | **[]string** | Filter results by a list of SSID IDs. | 
 **ssidNames** | **[]string** | Filter results by a list of SSID names. | 
 **gatewayNetworkIds** | **[]string** | Limit the results to clients tunneled to campus gateways in the provided networks. | 
 **usageUnits** | **string** | Usage units to use in the response. | 

### Return type

[**InlineResponse200463**](InlineResponse200463.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessClientsUsageBySsid

> InlineResponse200464 GetOrganizationWirelessClientsUsageBySsid(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).SsidNames(ssidNames).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).GatewayNetworkIds(gatewayNetworkIds).UsageUnits(usageUnits).Execute()

Returns client usage details for SSIDs within an organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. (optional)
    ssidNames := []string{"Inner_example"} // []string | Filter results by a list of SSID names. (optional)
    networkIds := []string{"Inner_example"} // []string | Limit the results to clients that belong to one of the provided networks. (optional)
    networkGroupIds := []string{"Inner_example"} // []string | Limit the results to clients that belong to one of the provided network groups. (optional)
    gatewayNetworkIds := []string{"Inner_example"} // []string | Limit the results to clients tunneled to campus gateways in the provided networks. (optional)
    usageUnits := "usageUnits_example" // string | Usage units to use in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationWirelessClientsUsageBySsid(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).SsidNames(ssidNames).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).GatewayNetworkIds(gatewayNetworkIds).UsageUnits(usageUnits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationWirelessClientsUsageBySsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessClientsUsageBySsid`: InlineResponse200464
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationWirelessClientsUsageBySsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessClientsUsageBySsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 1 hour and be less than or equal to 7 days. The default is 2 hours. | 
 **ssidNames** | **[]string** | Filter results by a list of SSID names. | 
 **networkIds** | **[]string** | Limit the results to clients that belong to one of the provided networks. | 
 **networkGroupIds** | **[]string** | Limit the results to clients that belong to one of the provided network groups. | 
 **gatewayNetworkIds** | **[]string** | Limit the results to clients tunneled to campus gateways in the provided networks. | 
 **usageUnits** | **string** | Usage units to use in the response. | 

### Return type

[**InlineResponse200464**](InlineResponse200464.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval

> InlineResponse200499 GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 2 interfaces history usage in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval`: InlineResponse200499
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200499**](InlineResponse200499.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval

> InlineResponse200502 GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(ctx, organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless LAN controller layer 3 interfaces history usage in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval(context.Background(), organizationId).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval`: InlineResponse200502
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200502**](InlineResponse200502.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval

> InlineResponse200504 GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(ctx, organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Retrieve the traffic for the interfaces of a Wireless LAN controller



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    names := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsageApi.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval(context.Background(), organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval`: InlineResponse200504
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.GetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesUsageHistoryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **names** | **[]string** | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200504**](InlineResponse200504.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

