# \OverviewsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline**](OverviewsApi.md#GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline) | **Get** /organizations/{organizationId}/api/rest/provisioning/pipelines/jobs/overviews/byPipeline | Retrieves pipeline overviews with aggregated job status counts
[**GetOrganizationApplianceInterfacesPacketsOverviewsByDevice**](OverviewsApi.md#GetOrganizationApplianceInterfacesPacketsOverviewsByDevice) | **Get** /organizations/{organizationId}/appliance/interfaces/packets/overviews/byDevice | Returns packet counter overviews for all interfaces on Secure Routers in the organization, including totals and average rates by packet type over the requested timespan.
[**GetOrganizationCampusGatewayClustersNetworksOverviews**](OverviewsApi.md#GetOrganizationCampusGatewayClustersNetworksOverviews) | **Get** /organizations/{organizationId}/campusGateway/clusters/networks/overviews | List networks tunneling through Campus Gateway clusters with their AP, ssids and client counts
[**GetOrganizationWirelessSsidsProfilesOverviews**](OverviewsApi.md#GetOrganizationWirelessSsidsProfilesOverviews) | **Get** /organizations/{organizationId}/wireless/ssids/profiles/overviews | Returns the SSID profiles&#39; overview information for an organization



## GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline

> InlineResponse200269 GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(ctx, organizationId).PipelineIds(pipelineIds).Execute()

Retrieves pipeline overviews with aggregated job status counts



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
    pipelineIds := []string{"Inner_example"} // []string | Pipeline IDs to retrieve overviews for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewsApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(context.Background(), organizationId).PipelineIds(pipelineIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewsApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: InlineResponse200269
    fmt.Fprintf(os.Stdout, "Response from `OverviewsApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipelineIds** | **[]string** | Pipeline IDs to retrieve overviews for | 

### Return type

[**InlineResponse200269**](InlineResponse200269.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceInterfacesPacketsOverviewsByDevice

> InlineResponse200284 GetOrganizationApplianceInterfacesPacketsOverviewsByDevice(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()

Returns packet counter overviews for all interfaces on Secure Routers in the organization, including totals and average rates by packet type over the requested timespan.



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 14 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter Secure Routers in the provided networks (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter Secure Routers by their serial numbers (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewsApi.GetOrganizationApplianceInterfacesPacketsOverviewsByDevice(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewsApi.GetOrganizationApplianceInterfacesPacketsOverviewsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceInterfacesPacketsOverviewsByDevice`: InlineResponse200284
    fmt.Fprintf(os.Stdout, "Response from `OverviewsApi.GetOrganizationApplianceInterfacesPacketsOverviewsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceInterfacesPacketsOverviewsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 14 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter Secure Routers in the provided networks | 
 **serials** | **[]string** | Optional parameter to filter Secure Routers by their serial numbers | 

### Return type

[**InlineResponse200284**](InlineResponse200284.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersNetworksOverviews

> InlineResponse200320 GetOrganizationCampusGatewayClustersNetworksOverviews(ctx, organizationId).ClusterIds(clusterIds).SiteIds(siteIds).NetworkIds(networkIds).TunnelingSources(tunnelingSources).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List networks tunneling through Campus Gateway clusters with their AP, ssids and client counts



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
    clusterIds := []string{"Inner_example"} // []string | Optional parameter to filter by MCG cluster IDs. This filter uses multiple exact matches. (optional)
    siteIds := []string{"Inner_example"} // []string | Optional parameter to filter by site IDs. This filter uses multiple exact matches. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter networks. This filter uses multiple exact matches. (optional)
    tunnelingSources := []string{"TunnelingSources_example"} // []string | Optional parameter to filter networks by tunneling source. 'configured' returns networks explicitly set up to tunnel through the campus gateway. 'roaming' returns networks tunneling due to AP roaming or disaster recovery. 'roaming' is only effective when 'clusterIds' is also provided; without 'clusterIds', the filter defaults to configured-only behavior. Defaults to 'configured' if omitted. (optional)
    search := "search_example" // string | Optional parameter to filter networks by wireless network name. This filter uses case-insensitive substring matching. (optional)
    sortBy := "sortBy_example" // string | Optional parameter to sort results. Default is 'name'. Use 'siteName' to sort by site name. (optional)
    sortOrder := "sortOrder_example" // string | Optional parameter to specify sort direction. Default is 'asc'. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewsApi.GetOrganizationCampusGatewayClustersNetworksOverviews(context.Background(), organizationId).ClusterIds(clusterIds).SiteIds(siteIds).NetworkIds(networkIds).TunnelingSources(tunnelingSources).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewsApi.GetOrganizationCampusGatewayClustersNetworksOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersNetworksOverviews`: InlineResponse200320
    fmt.Fprintf(os.Stdout, "Response from `OverviewsApi.GetOrganizationCampusGatewayClustersNetworksOverviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersNetworksOverviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterIds** | **[]string** | Optional parameter to filter by MCG cluster IDs. This filter uses multiple exact matches. | 
 **siteIds** | **[]string** | Optional parameter to filter by site IDs. This filter uses multiple exact matches. | 
 **networkIds** | **[]string** | Optional parameter to filter networks. This filter uses multiple exact matches. | 
 **tunnelingSources** | **[]string** | Optional parameter to filter networks by tunneling source. &#39;configured&#39; returns networks explicitly set up to tunnel through the campus gateway. &#39;roaming&#39; returns networks tunneling due to AP roaming or disaster recovery. &#39;roaming&#39; is only effective when &#39;clusterIds&#39; is also provided; without &#39;clusterIds&#39;, the filter defaults to configured-only behavior. Defaults to &#39;configured&#39; if omitted. | 
 **search** | **string** | Optional parameter to filter networks by wireless network name. This filter uses case-insensitive substring matching. | 
 **sortBy** | **string** | Optional parameter to sort results. Default is &#39;name&#39;. Use &#39;siteName&#39; to sort by site name. | 
 **sortOrder** | **string** | Optional parameter to specify sort direction. Default is &#39;asc&#39;. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200320**](InlineResponse200320.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessSsidsProfilesOverviews

> InlineResponse200488 GetOrganizationWirelessSsidsProfilesOverviews(ctx, organizationId).Name(name).SortBy(sortBy).SortOrder(sortOrder).ProfileIds(profileIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns the SSID profiles' overview information for an organization



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
    name := "name_example" // string | (Optional) Filter results by name. Case insensitive substring match. (optional)
    sortBy := "sortBy_example" // string | Column to sort results by. Default is `name`. (optional)
    sortOrder := "sortOrder_example" // string | Direction to sort results by. Default is `asc`. (optional)
    profileIds := []int32{int32(123)} // []int32 | (Optional) Filter results by a list of SSID profile IDs. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewsApi.GetOrganizationWirelessSsidsProfilesOverviews(context.Background(), organizationId).Name(name).SortBy(sortBy).SortOrder(sortOrder).ProfileIds(profileIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewsApi.GetOrganizationWirelessSsidsProfilesOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessSsidsProfilesOverviews`: InlineResponse200488
    fmt.Fprintf(os.Stdout, "Response from `OverviewsApi.GetOrganizationWirelessSsidsProfilesOverviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessSsidsProfilesOverviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | (Optional) Filter results by name. Case insensitive substring match. | 
 **sortBy** | **string** | Column to sort results by. Default is &#x60;name&#x60;. | 
 **sortOrder** | **string** | Direction to sort results by. Default is &#x60;asc&#x60;. | 
 **profileIds** | **[]int32** | (Optional) Filter results by a list of SSID profile IDs. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200488**](InlineResponse200488.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

