# \ConnectionsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCampusGatewayConnections**](ConnectionsApi.md#GetOrganizationCampusGatewayConnections) | **Get** /organizations/{organizationId}/campusGateway/connections | List the details of APs tunneling through the Campus Gateway clusters
[**GetOrganizationCampusGatewayConnectionsOverview**](ConnectionsApi.md#GetOrganizationCampusGatewayConnectionsOverview) | **Get** /organizations/{organizationId}/campusGateway/connections/overview | List the count of connections(APs) with tunneling status up and down through the Campus Gateway clusters
[**GetOrganizationSensorGatewaysConnectionsLatest**](ConnectionsApi.md#GetOrganizationSensorGatewaysConnectionsLatest) | **Get** /organizations/{organizationId}/sensor/gateways/connections/latest | Returns latest sensor-gateway connectivity data.
[**GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid**](ConnectionsApi.md#GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid) | **Get** /organizations/{organizationId}/wireless/clients/connections/impacted/byNetwork/bySsid | Summarize the number of wireless clients impacted by connection failures on network SSIDs, across an organization.
[**GetOrganizationWirelessControllerConnections**](ConnectionsApi.md#GetOrganizationWirelessControllerConnections) | **Get** /organizations/{organizationId}/wirelessController/connections | List all access points associated with wireless LAN controllers in an organization



## GetOrganizationCampusGatewayConnections

> InlineResponse200323 GetOrganizationCampusGatewayConnections(ctx, organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the details of APs tunneling through the Campus Gateway clusters



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter networks. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by its own serials. This filter uses multiple exact matches. (optional)
    campusGatewaySerials := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by MCG serials. This filter uses multiple exact matches. (optional)
    campusGatewayClusterIds := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by MCG cluster IDs. This filter uses multiple exact matches. (optional)
    campusGatewayTunnelStatuses := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by tunnel statuses. This filter uses multiple exact matches. (optional)
    search := "search_example" // string | Optional parameter to filter connections(APs) on AP name, serial, MAC address, network name, or interface IP address. This filter uses partial string matching (ILIKE). (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by device model names. This filter uses multiple exact matches. (optional)
    dataEncryptionStatuses := []string{"DataEncryptionStatuses_example"} // []string | Optional parameter to filter connections(APs) by data encryption status. This filter uses multiple exact matches. (optional)
    sortBy := "sortBy_example" // string | Optional parameter to sort results. Available options: name, serial, status, interfaces, clients, dataEncryption, networkName. Default is 'serial'. (optional)
    sortOrder := "sortOrder_example" // string | Optional parameter to specify sort direction. Default is 'asc'. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetOrganizationCampusGatewayConnections(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetOrganizationCampusGatewayConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayConnections`: InlineResponse200323
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetOrganizationCampusGatewayConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter networks. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter connections(APs) by its own serials. This filter uses multiple exact matches. | 
 **campusGatewaySerials** | **[]string** | Optional parameter to filter connections(APs) by MCG serials. This filter uses multiple exact matches. | 
 **campusGatewayClusterIds** | **[]string** | Optional parameter to filter connections(APs) by MCG cluster IDs. This filter uses multiple exact matches. | 
 **campusGatewayTunnelStatuses** | **[]string** | Optional parameter to filter connections(APs) by tunnel statuses. This filter uses multiple exact matches. | 
 **search** | **string** | Optional parameter to filter connections(APs) on AP name, serial, MAC address, network name, or interface IP address. This filter uses partial string matching (ILIKE). | 
 **models** | **[]string** | Optional parameter to filter connections(APs) by device model names. This filter uses multiple exact matches. | 
 **dataEncryptionStatuses** | **[]string** | Optional parameter to filter connections(APs) by data encryption status. This filter uses multiple exact matches. | 
 **sortBy** | **string** | Optional parameter to sort results. Available options: name, serial, status, interfaces, clients, dataEncryption, networkName. Default is &#39;serial&#39;. | 
 **sortOrder** | **string** | Optional parameter to specify sort direction. Default is &#39;asc&#39;. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200323**](InlineResponse200323.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayConnectionsOverview

> InlineResponse200324 GetOrganizationCampusGatewayConnectionsOverview(ctx, organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).Execute()

List the count of connections(APs) with tunneling status up and down through the Campus Gateway clusters



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter networks. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by its own serials. This filter uses multiple exact matches. (optional)
    campusGatewaySerials := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by Campus Gateway serials. This filter uses multiple exact matches. (optional)
    campusGatewayClusterIds := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by Campus Gateway cluster IDs. This filter uses multiple exact matches. (optional)
    campusGatewayTunnelStatuses := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by tunnel statuses. This filter uses multiple exact matches. (optional)
    search := "search_example" // string | Optional setting that lets you filter access points (APs) by name, serial number, MAC address, network name, or interface IP address. The filter matches partial text, not just exact values (uses ILIKE matching). (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by device model names. This filter uses multiple exact matches. (optional)
    dataEncryptionStatuses := []string{"DataEncryptionStatuses_example"} // []string | Optional parameter to filter connections(APs) by data encryption status. This filter uses multiple exact matches. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetOrganizationCampusGatewayConnectionsOverview(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetOrganizationCampusGatewayConnectionsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayConnectionsOverview`: InlineResponse200324
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetOrganizationCampusGatewayConnectionsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayConnectionsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter networks. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter connections(APs) by its own serials. This filter uses multiple exact matches. | 
 **campusGatewaySerials** | **[]string** | Optional parameter to filter connections(APs) by Campus Gateway serials. This filter uses multiple exact matches. | 
 **campusGatewayClusterIds** | **[]string** | Optional parameter to filter connections(APs) by Campus Gateway cluster IDs. This filter uses multiple exact matches. | 
 **campusGatewayTunnelStatuses** | **[]string** | Optional parameter to filter connections(APs) by tunnel statuses. This filter uses multiple exact matches. | 
 **search** | **string** | Optional setting that lets you filter access points (APs) by name, serial number, MAC address, network name, or interface IP address. The filter matches partial text, not just exact values (uses ILIKE matching). | 
 **models** | **[]string** | Optional parameter to filter connections(APs) by device model names. This filter uses multiple exact matches. | 
 **dataEncryptionStatuses** | **[]string** | Optional parameter to filter connections(APs) by data encryption status. This filter uses multiple exact matches. | 

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


## GetOrganizationSensorGatewaysConnectionsLatest

> InlineResponse200423 GetOrganizationSensorGatewaysConnectionsLatest(ctx, organizationId).SensorSerials(sensorSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns latest sensor-gateway connectivity data.



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
    sensorSerials := []string{"Inner_example"} // []string | List of sensor serials to filter connectivity data by sensor. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetOrganizationSensorGatewaysConnectionsLatest(context.Background(), organizationId).SensorSerials(sensorSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetOrganizationSensorGatewaysConnectionsLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSensorGatewaysConnectionsLatest`: InlineResponse200423
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetOrganizationSensorGatewaysConnectionsLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSensorGatewaysConnectionsLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensorSerials** | **[]string** | List of sensor serials to filter connectivity data by sensor. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200423**](InlineResponse200423.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid

> InlineResponse200460 GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid(ctx, organizationId).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Summarize the number of wireless clients impacted by connection failures on network SSIDs, across an organization.



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    networkGroupIds := []string{"Inner_example"} // []string | Filter results by a list of network group IDs. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 7 days. The default is 2 hours. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid(context.Background(), organizationId).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid`: InlineResponse200460
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessClientsConnectionsImpactedByNetworkBySsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **networkGroupIds** | **[]string** | Filter results by a list of network group IDs. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 8 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 7 days. The default is 2 hours. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200460**](InlineResponse200460.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerConnections

> InlineResponse200496 GetOrganizationWirelessControllerConnections(ctx, organizationId).NetworkIds(networkIds).ControllerSerials(controllerSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List all access points associated with wireless LAN controllers in an organization



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter access points by network ID. This filter uses multiple exact matches. (optional)
    controllerSerials := []string{"Inner_example"} // []string | Optional parameter to filter access points by its controller cloud ID. This filter uses multiple exact matches. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionsApi.GetOrganizationWirelessControllerConnections(context.Background(), organizationId).NetworkIds(networkIds).ControllerSerials(controllerSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsApi.GetOrganizationWirelessControllerConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerConnections`: InlineResponse200496
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsApi.GetOrganizationWirelessControllerConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter access points by network ID. This filter uses multiple exact matches. | 
 **controllerSerials** | **[]string** | Optional parameter to filter access points by its controller cloud ID. This filter uses multiple exact matches. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200496**](InlineResponse200496.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

