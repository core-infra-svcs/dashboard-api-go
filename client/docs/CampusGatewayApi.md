# \CampusGatewayApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchOrganizationCampusGatewayClustersTunnelingUpdate**](CampusGatewayApi.md#BatchOrganizationCampusGatewayClustersTunnelingUpdate) | **Post** /organizations/{organizationId}/campusGateway/clusters/tunneling/batchUpdate | Update MCG cluster-network tunnel settings for multiple networks
[**CreateNetworkCampusGatewayCluster**](CampusGatewayApi.md#CreateNetworkCampusGatewayCluster) | **Post** /networks/{networkId}/campusGateway/clusters | Create a cluster and add campus gateways to it
[**DeleteNetworkCampusGatewayCluster**](CampusGatewayApi.md#DeleteNetworkCampusGatewayCluster) | **Delete** /networks/{networkId}/campusGateway/clusters/{clusterId} | Delete a cluster
[**GetOrganizationCampusGatewayClientsUsageByNetworkByCluster**](CampusGatewayApi.md#GetOrganizationCampusGatewayClientsUsageByNetworkByCluster) | **Get** /organizations/{organizationId}/campusGateway/clients/usage/byNetwork/byCluster | Returns client usage details for campus gateway clusters within an organization.
[**GetOrganizationCampusGatewayClusters**](CampusGatewayApi.md#GetOrganizationCampusGatewayClusters) | **Get** /organizations/{organizationId}/campusGateway/clusters | Get the details of campus gateway clusters
[**GetOrganizationCampusGatewayClustersFailoverTargets**](CampusGatewayApi.md#GetOrganizationCampusGatewayClustersFailoverTargets) | **Get** /organizations/{organizationId}/campusGateway/clusters/failover/targets | List the details of a Failover Targets for a Campus Gateway cluster
[**GetOrganizationCampusGatewayClustersFailoverTargetsByCluster**](CampusGatewayApi.md#GetOrganizationCampusGatewayClustersFailoverTargetsByCluster) | **Get** /organizations/{organizationId}/campusGateway/clusters/failover/targets/byCluster | Get available backup cluster targets for campus gateway failover configuration
[**GetOrganizationCampusGatewayClustersNetworksOverviews**](CampusGatewayApi.md#GetOrganizationCampusGatewayClustersNetworksOverviews) | **Get** /organizations/{organizationId}/campusGateway/clusters/networks/overviews | List networks tunneling through Campus Gateway clusters with their AP, ssids and client counts
[**GetOrganizationCampusGatewayClustersSsids**](CampusGatewayApi.md#GetOrganizationCampusGatewayClustersSsids) | **Get** /organizations/{organizationId}/campusGateway/clusters/ssids | List SSIDs tunneling through Campus Gateway clusters
[**GetOrganizationCampusGatewayClustersTunnelable**](CampusGatewayApi.md#GetOrganizationCampusGatewayClustersTunnelable) | **Get** /organizations/{organizationId}/campusGateway/clusters/tunnelable | List available campus gateway clusters that can be used for wireless network tunneling
[**GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork**](CampusGatewayApi.md#GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork) | **Get** /organizations/{organizationId}/campusGateway/clusters/tunneling/byCluster/byNetwork | List all the MCG cluster-network tunnel settings
[**GetOrganizationCampusGatewayConnections**](CampusGatewayApi.md#GetOrganizationCampusGatewayConnections) | **Get** /organizations/{organizationId}/campusGateway/connections | List the details of APs tunneling through the Campus Gateway clusters
[**GetOrganizationCampusGatewayConnectionsOverview**](CampusGatewayApi.md#GetOrganizationCampusGatewayConnectionsOverview) | **Get** /organizations/{organizationId}/campusGateway/connections/overview | List the count of connections(APs) with tunneling status up and down through the Campus Gateway clusters
[**GetOrganizationCampusGatewayDevicesSystemCpuLoad**](CampusGatewayApi.md#GetOrganizationCampusGatewayDevicesSystemCpuLoad) | **Get** /organizations/{organizationId}/campusGateway/devices/system/cpu/load | Return the latest CPU load averages reported by Campus Gateway devices.
[**GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice**](CampusGatewayApi.md#GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice) | **Get** /organizations/{organizationId}/campusGateway/devices/uplinks/localOverrides/byDevice | Uplink overrides configured locally on Campus Gateway devices in an organization.
[**ProvisionOrganizationCampusGatewayClusters**](CampusGatewayApi.md#ProvisionOrganizationCampusGatewayClusters) | **Post** /organizations/{organizationId}/campusGateway/clusters/provision | Provisions a cluster,adds campus gateways to it and associate/dissociate failover targets.
[**UpdateNetworkCampusGatewayCluster**](CampusGatewayApi.md#UpdateNetworkCampusGatewayCluster) | **Put** /networks/{networkId}/campusGateway/clusters/{clusterId} | Update a cluster and add/remove campus gateways to/from it
[**UpdateNetworkCampusGatewaySsidMdns**](CampusGatewayApi.md#UpdateNetworkCampusGatewaySsidMdns) | **Put** /networks/{networkId}/campusGateway/ssids/{number}/mdns | Update the mDNS gateway settings and rules for a SSID and cluster



## BatchOrganizationCampusGatewayClustersTunnelingUpdate

> InlineResponse200338 BatchOrganizationCampusGatewayClustersTunnelingUpdate(ctx, organizationId).BatchOrganizationCampusGatewayClustersTunnelingUpdate(batchOrganizationCampusGatewayClustersTunnelingUpdate).Execute()

Update MCG cluster-network tunnel settings for multiple networks



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
    batchOrganizationCampusGatewayClustersTunnelingUpdate := *openapiclient.NewInlineObject285([]openapiclient.OrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateItems{*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateItems(*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateCluster("Id_example"), *openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateNetwork("Id_example"), *openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateData(*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateDataEncryption(false)))}) // InlineObject285 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.BatchOrganizationCampusGatewayClustersTunnelingUpdate(context.Background(), organizationId).BatchOrganizationCampusGatewayClustersTunnelingUpdate(batchOrganizationCampusGatewayClustersTunnelingUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.BatchOrganizationCampusGatewayClustersTunnelingUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationCampusGatewayClustersTunnelingUpdate`: InlineResponse200338
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.BatchOrganizationCampusGatewayClustersTunnelingUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchOrganizationCampusGatewayClustersTunnelingUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchOrganizationCampusGatewayClustersTunnelingUpdate** | [**InlineObject285**](InlineObject285.md) |  | 

### Return type

[**InlineResponse200338**](InlineResponse200338.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkCampusGatewayCluster

> InlineResponse20118 CreateNetworkCampusGatewayCluster(ctx, networkId).CreateNetworkCampusGatewayCluster(createNetworkCampusGatewayCluster).Execute()

Create a cluster and add campus gateways to it



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
    createNetworkCampusGatewayCluster := *openapiclient.NewInlineObject105("Name_example", []openapiclient.NetworksNetworkIdCampusGatewayClustersUplinks{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersUplinks("Interface_example", int32(123), []openapiclient.NetworksNetworkIdCampusGatewayClustersAddresses{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersAddresses("AssignmentMode_example")})}, []openapiclient.NetworksNetworkIdCampusGatewayClustersTunnels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersTunnels()}, *openapiclient.NewNetworksNetworkIdCampusGatewayClustersNameservers([]string{"Addresses_example"}), []openapiclient.NetworksNetworkIdCampusGatewayClustersPortChannels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersPortChannels("Name_example", int32(123), "AllowedVlans_example")}) // InlineObject105 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.CreateNetworkCampusGatewayCluster(context.Background(), networkId).CreateNetworkCampusGatewayCluster(createNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.CreateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCampusGatewayCluster`: InlineResponse20118
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.CreateNetworkCampusGatewayCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCampusGatewayClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCampusGatewayCluster** | [**InlineObject105**](InlineObject105.md) |  | 

### Return type

[**InlineResponse20118**](InlineResponse20118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkCampusGatewayCluster

> DeleteNetworkCampusGatewayCluster(ctx, networkId, clusterId).Execute()

Delete a cluster



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
    clusterId := "clusterId_example" // string | Cluster ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.DeleteNetworkCampusGatewayCluster(context.Background(), networkId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.DeleteNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkCampusGatewayClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClientsUsageByNetworkByCluster

> InlineResponse200331 GetOrganizationCampusGatewayClientsUsageByNetworkByCluster(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).ClusterIds(clusterIds).UsageUnits(usageUnits).Execute()

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
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).ClusterIds(clusterIds).UsageUnits(usageUnits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClientsUsageByNetworkByCluster`: InlineResponse200331
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster`: %v\n", resp)
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

[**InlineResponse200331**](InlineResponse200331.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClusters

> InlineResponse200332 GetOrganizationCampusGatewayClusters(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Get the details of campus gateway clusters



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
    networkIds := []string{"Inner_example"} // []string | Networks for which information should be gathered. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClusters(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClusters`: InlineResponse200332
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Networks for which information should be gathered. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200332**](InlineResponse200332.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersFailoverTargets

> []InlineResponse200333 GetOrganizationCampusGatewayClustersFailoverTargets(ctx, organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the details of a Failover Targets for a Campus Gateway cluster



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
    clusterIds := []string{"Inner_example"} // []string | Optional parameter to filter clusters. This filter uses multiple exact matches. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClustersFailoverTargets(context.Background(), organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClustersFailoverTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersFailoverTargets`: []InlineResponse200333
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClustersFailoverTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersFailoverTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter networks. This filter uses multiple exact matches. | 
 **clusterIds** | **[]string** | Optional parameter to filter clusters. This filter uses multiple exact matches. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse200333**](InlineResponse200333.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersFailoverTargetsByCluster

> InlineResponse200334 GetOrganizationCampusGatewayClustersFailoverTargetsByCluster(ctx, organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Get available backup cluster targets for campus gateway failover configuration



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
    networkIds := []string{"Inner_example"} // []string | Networks for which backup cluster targets should be gathered. (optional)
    clusterIds := []string{"Inner_example"} // []string | Cluster IDs to filter backup cluster targets. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster(context.Background(), organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: InlineResponse200334
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersFailoverTargetsByClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Networks for which backup cluster targets should be gathered. | 
 **clusterIds** | **[]string** | Cluster IDs to filter backup cluster targets. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200334**](InlineResponse200334.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersNetworksOverviews

> InlineResponse200335 GetOrganizationCampusGatewayClustersNetworksOverviews(ctx, organizationId).ClusterIds(clusterIds).SiteIds(siteIds).NetworkIds(networkIds).TunnelingSources(tunnelingSources).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClustersNetworksOverviews(context.Background(), organizationId).ClusterIds(clusterIds).SiteIds(siteIds).NetworkIds(networkIds).TunnelingSources(tunnelingSources).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClustersNetworksOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersNetworksOverviews`: InlineResponse200335
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClustersNetworksOverviews`: %v\n", resp)
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

[**InlineResponse200335**](InlineResponse200335.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersSsids

> InlineResponse200336 GetOrganizationCampusGatewayClustersSsids(ctx, organizationId).ClusterIds(clusterIds).NetworkIds(networkIds).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List SSIDs tunneling through Campus Gateway clusters



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter networks. This filter uses multiple exact matches. (optional)
    search := "search_example" // string | Optional parameter to filter SSIDs by name. This filter uses case-insensitive starts with string matching. (optional)
    sortBy := "sortBy_example" // string | Optional parameter to sort results. Default is 'networkId'. (optional)
    sortOrder := "sortOrder_example" // string | Optional parameter to specify sort direction. Default is 'asc'. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClustersSsids(context.Background(), organizationId).ClusterIds(clusterIds).NetworkIds(networkIds).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClustersSsids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersSsids`: InlineResponse200336
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClustersSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersSsidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterIds** | **[]string** | Optional parameter to filter by MCG cluster IDs. This filter uses multiple exact matches. | 
 **networkIds** | **[]string** | Optional parameter to filter networks. This filter uses multiple exact matches. | 
 **search** | **string** | Optional parameter to filter SSIDs by name. This filter uses case-insensitive starts with string matching. | 
 **sortBy** | **string** | Optional parameter to sort results. Default is &#39;networkId&#39;. | 
 **sortOrder** | **string** | Optional parameter to specify sort direction. Default is &#39;asc&#39;. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200336**](InlineResponse200336.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersTunnelable

> InlineResponse200337 GetOrganizationCampusGatewayClustersTunnelable(ctx, organizationId).FromNetworkIds(fromNetworkIds).ExcludeSsidNumbers(excludeSsidNumbers).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List available campus gateway clusters that can be used for wireless network tunneling



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
    fromNetworkIds := []string{"Inner_example"} // []string | Filter results by wireless network IDs.
    excludeSsidNumbers := []string{"Inner_example"} // []string | Optional SSID numbers (0-15) to exclude from restriction checks, corresponding to each fromNetworkIds entry. Only one SSID per network can be excluded. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 200. Default is 200. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClustersTunnelable(context.Background(), organizationId).FromNetworkIds(fromNetworkIds).ExcludeSsidNumbers(excludeSsidNumbers).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClustersTunnelable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersTunnelable`: InlineResponse200337
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClustersTunnelable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersTunnelableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromNetworkIds** | **[]string** | Filter results by wireless network IDs. | 
 **excludeSsidNumbers** | **[]string** | Optional SSID numbers (0-15) to exclude from restriction checks, corresponding to each fromNetworkIds entry. Only one SSID per network can be excluded. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 200. Default is 200. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200337**](InlineResponse200337.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork

> InlineResponse200339 GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork(ctx, organizationId).ClusterIds(clusterIds).DataEncryptionEnabled(dataEncryptionEnabled).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List all the MCG cluster-network tunnel settings



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
    clusterIds := []string{"Inner_example"} // []string | Optional parameter to filter MCG clusters. This filter uses multiple exact matches (optional)
    dataEncryptionEnabled := true // bool | Optional parameter to filter cluster-network tunnel settings by data encryption configuration (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter networks. This filter uses multiple exact matches (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork(context.Background(), organizationId).ClusterIds(clusterIds).DataEncryptionEnabled(dataEncryptionEnabled).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork`: InlineResponse200339
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersTunnelingByClusterByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterIds** | **[]string** | Optional parameter to filter MCG clusters. This filter uses multiple exact matches | 
 **dataEncryptionEnabled** | **bool** | Optional parameter to filter cluster-network tunnel settings by data encryption configuration | 
 **networkIds** | **[]string** | Optional parameter to filter networks. This filter uses multiple exact matches | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200339**](InlineResponse200339.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayConnections

> InlineResponse200340 GetOrganizationCampusGatewayConnections(ctx, organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).CampusGatewayPriorities(campusGatewayPriorities).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    campusGatewayPriorities := []int32{int32(123)} // []int32 | Optional parameter to filter connections(APs) by Campus Gateway tunnel priority. Use 0 for primary or 1 for backup. This filter uses multiple exact matches. (optional)
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
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayConnections(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).CampusGatewayPriorities(campusGatewayPriorities).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayConnections`: InlineResponse200340
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayConnections`: %v\n", resp)
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
 **campusGatewayPriorities** | **[]int32** | Optional parameter to filter connections(APs) by Campus Gateway tunnel priority. Use 0 for primary or 1 for backup. This filter uses multiple exact matches. | 
 **search** | **string** | Optional parameter to filter connections(APs) on AP name, serial, MAC address, network name, or interface IP address. This filter uses partial string matching (ILIKE). | 
 **models** | **[]string** | Optional parameter to filter connections(APs) by device model names. This filter uses multiple exact matches. | 
 **dataEncryptionStatuses** | **[]string** | Optional parameter to filter connections(APs) by data encryption status. This filter uses multiple exact matches. | 
 **sortBy** | **string** | Optional parameter to sort results. Available options: name, serial, status, interfaces, clients, dataEncryption, networkName. Default is &#39;serial&#39;. | 
 **sortOrder** | **string** | Optional parameter to specify sort direction. Default is &#39;asc&#39;. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200340**](InlineResponse200340.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayConnectionsOverview

> InlineResponse200341 GetOrganizationCampusGatewayConnectionsOverview(ctx, organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).CampusGatewayPriorities(campusGatewayPriorities).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).Execute()

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
    campusGatewayPriorities := []int32{int32(123)} // []int32 | Optional parameter to filter connections(APs) by Campus Gateway tunnel priority. Use 0 for primary or 1 for backup. This filter uses multiple exact matches. (optional)
    search := "search_example" // string | Optional setting that lets you filter access points (APs) by name, serial number, MAC address, network name, or interface IP address. The filter matches partial text, not just exact values (uses ILIKE matching). (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter connections(APs) by device model names. This filter uses multiple exact matches. (optional)
    dataEncryptionStatuses := []string{"DataEncryptionStatuses_example"} // []string | Optional parameter to filter connections(APs) by data encryption status. This filter uses multiple exact matches. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayConnectionsOverview(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewaySerials(campusGatewaySerials).CampusGatewayClusterIds(campusGatewayClusterIds).CampusGatewayTunnelStatuses(campusGatewayTunnelStatuses).CampusGatewayPriorities(campusGatewayPriorities).Search(search).Models(models).DataEncryptionStatuses(dataEncryptionStatuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayConnectionsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayConnectionsOverview`: InlineResponse200341
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayConnectionsOverview`: %v\n", resp)
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
 **campusGatewayPriorities** | **[]int32** | Optional parameter to filter connections(APs) by Campus Gateway tunnel priority. Use 0 for primary or 1 for backup. This filter uses multiple exact matches. | 
 **search** | **string** | Optional setting that lets you filter access points (APs) by name, serial number, MAC address, network name, or interface IP address. The filter matches partial text, not just exact values (uses ILIKE matching). | 
 **models** | **[]string** | Optional parameter to filter connections(APs) by device model names. This filter uses multiple exact matches. | 
 **dataEncryptionStatuses** | **[]string** | Optional parameter to filter connections(APs) by data encryption status. This filter uses multiple exact matches. | 

### Return type

[**InlineResponse200341**](InlineResponse200341.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayDevicesSystemCpuLoad

> InlineResponse200342 GetOrganizationCampusGatewayDevicesSystemCpuLoad(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the latest CPU load averages reported by Campus Gateway devices.



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
    networkIds := []string{"Inner_example"} // []string | Optional exact-match network filters. Accepts both node group (N_) and locale (L_) network IDs. (optional)
    serials := []string{"Inner_example"} // []string | Optional exact-match Campus Gateway serial filters. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayDevicesSystemCpuLoad(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayDevicesSystemCpuLoad``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayDevicesSystemCpuLoad`: InlineResponse200342
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayDevicesSystemCpuLoad`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayDevicesSystemCpuLoadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional exact-match network filters. Accepts both node group (N_) and locale (L_) network IDs. | 
 **serials** | **[]string** | Optional exact-match Campus Gateway serial filters. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200342**](InlineResponse200342.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice

> InlineResponse200343 GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Uplink overrides configured locally on Campus Gateway devices in an organization.



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
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice`: InlineResponse200343
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200343**](InlineResponse200343.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionOrganizationCampusGatewayClusters

> InlineResponse2024 ProvisionOrganizationCampusGatewayClusters(ctx, organizationId).ProvisionOrganizationCampusGatewayClusters(provisionOrganizationCampusGatewayClusters).Execute()

Provisions a cluster,adds campus gateways to it and associate/dissociate failover targets.



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
    provisionOrganizationCampusGatewayClusters := *openapiclient.NewInlineObject284("ClusterId_example", *openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork("Id_example"), "Name_example", []openapiclient.NetworksNetworkIdCampusGatewayClustersUplinks{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersUplinks("Interface_example", int32(123), []openapiclient.NetworksNetworkIdCampusGatewayClustersAddresses{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersAddresses("AssignmentMode_example")})}, []openapiclient.NetworksNetworkIdCampusGatewayClustersTunnels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersTunnels()}, *openapiclient.NewNetworksNetworkIdCampusGatewayClustersNameservers([]string{"Addresses_example"}), []openapiclient.OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels{*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels("Name_example", int32(123), "AllowedVlans_example")}) // InlineObject284 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.ProvisionOrganizationCampusGatewayClusters(context.Background(), organizationId).ProvisionOrganizationCampusGatewayClusters(provisionOrganizationCampusGatewayClusters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.ProvisionOrganizationCampusGatewayClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionOrganizationCampusGatewayClusters`: InlineResponse2024
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.ProvisionOrganizationCampusGatewayClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionOrganizationCampusGatewayClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisionOrganizationCampusGatewayClusters** | [**InlineObject284**](InlineObject284.md) |  | 

### Return type

[**InlineResponse2024**](InlineResponse2024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCampusGatewayCluster

> InlineResponse20118 UpdateNetworkCampusGatewayCluster(ctx, networkId, clusterId).UpdateNetworkCampusGatewayCluster(updateNetworkCampusGatewayCluster).Execute()

Update a cluster and add/remove campus gateways to/from it



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
    clusterId := "clusterId_example" // string | Cluster ID
    updateNetworkCampusGatewayCluster := *openapiclient.NewInlineObject106() // InlineObject106 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.UpdateNetworkCampusGatewayCluster(context.Background(), networkId, clusterId).UpdateNetworkCampusGatewayCluster(updateNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.UpdateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCampusGatewayCluster`: InlineResponse20118
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.UpdateNetworkCampusGatewayCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCampusGatewayClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCampusGatewayCluster** | [**InlineObject106**](InlineObject106.md) |  | 

### Return type

[**InlineResponse20118**](InlineResponse20118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCampusGatewaySsidMdns

> InlineResponse200118 UpdateNetworkCampusGatewaySsidMdns(ctx, networkId, number).UpdateNetworkCampusGatewaySsidMdns(updateNetworkCampusGatewaySsidMdns).Execute()

Update the mDNS gateway settings and rules for a SSID and cluster



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
    updateNetworkCampusGatewaySsidMdns := *openapiclient.NewInlineObject107() // InlineObject107 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampusGatewayApi.UpdateNetworkCampusGatewaySsidMdns(context.Background(), networkId, number).UpdateNetworkCampusGatewaySsidMdns(updateNetworkCampusGatewaySsidMdns).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.UpdateNetworkCampusGatewaySsidMdns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCampusGatewaySsidMdns`: InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `CampusGatewayApi.UpdateNetworkCampusGatewaySsidMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCampusGatewaySsidMdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCampusGatewaySsidMdns** | [**InlineObject107**](InlineObject107.md) |  | 

### Return type

[**InlineResponse200118**](InlineResponse200118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

