# \ConnectivityApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceConnectivity**](ConnectivityApi.md#GetNetworkSmDeviceConnectivity) | **Get** /networks/{networkId}/sm/devices/{deviceId}/connectivity | Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
[**GetOrganizationSaseSitesConnectivityHistoryBySite**](ConnectivityApi.md#GetOrganizationSaseSitesConnectivityHistoryBySite) | **Get** /organizations/{organizationId}/sase/sites/connectivity/history/bySite | Get the connectivity history of SASE sites in this organization
[**GetOrganizationSaseSitesConnectivityOverview**](ConnectivityApi.md#GetOrganizationSaseSitesConnectivityOverview) | **Get** /organizations/{organizationId}/sase/sites/connectivity/overview | List high-level SASE site statuses (healthy, degraded, offline)



## GetNetworkSmDeviceConnectivity

> []InlineResponse200171 GetNetworkSmDeviceConnectivity(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns historical connectivity data (whether a device is regularly checking in to Dashboard).



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
    deviceId := "deviceId_example" // string | Device ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectivityApi.GetNetworkSmDeviceConnectivity(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityApi.GetNetworkSmDeviceConnectivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceConnectivity`: []InlineResponse200171
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityApi.GetNetworkSmDeviceConnectivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200171**](InlineResponse200171.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSitesConnectivityHistoryBySite

> InlineResponse200441 GetOrganizationSaseSitesConnectivityHistoryBySite(ctx, organizationId).SiteIds(siteIds).Timespan(timespan).Execute()

Get the connectivity history of SASE sites in this organization



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
    siteIds := []string{"Inner_example"} // []string | Array of site IDs to fetch connectivity status data (maximum 100 values) (optional)
    timespan := "timespan_example" // string | Timespan for the status data (e.g., '-2hours') (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectivityApi.GetOrganizationSaseSitesConnectivityHistoryBySite(context.Background(), organizationId).SiteIds(siteIds).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityApi.GetOrganizationSaseSitesConnectivityHistoryBySite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseSitesConnectivityHistoryBySite`: InlineResponse200441
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityApi.GetOrganizationSaseSitesConnectivityHistoryBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseSitesConnectivityHistoryBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteIds** | **[]string** | Array of site IDs to fetch connectivity status data (maximum 100 values) | 
 **timespan** | **string** | Timespan for the status data (e.g., &#39;-2hours&#39;) | 

### Return type

[**InlineResponse200441**](InlineResponse200441.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSitesConnectivityOverview

> InlineResponse200442 GetOrganizationSaseSitesConnectivityOverview(ctx, organizationId).Execute()

List high-level SASE site statuses (healthy, degraded, offline)



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
    resp, r, err := apiClient.ConnectivityApi.GetOrganizationSaseSitesConnectivityOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityApi.GetOrganizationSaseSitesConnectivityOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseSitesConnectivityOverview`: InlineResponse200442
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityApi.GetOrganizationSaseSitesConnectivityOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseSitesConnectivityOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200442**](InlineResponse200442.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

