# \ClustersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchOrganizationCampusGatewayClustersTunnelingUpdate**](ClustersApi.md#BatchOrganizationCampusGatewayClustersTunnelingUpdate) | **Post** /organizations/{organizationId}/campusGateway/clusters/tunneling/batchUpdate | Update MCG cluster-network tunnel settings for multiple networks
[**CreateNetworkCampusGatewayCluster**](ClustersApi.md#CreateNetworkCampusGatewayCluster) | **Post** /networks/{networkId}/campusGateway/clusters | Create a cluster and add campus gateways to it
[**DeleteNetworkCampusGatewayCluster**](ClustersApi.md#DeleteNetworkCampusGatewayCluster) | **Delete** /networks/{networkId}/campusGateway/clusters/{clusterId} | Delete a cluster
[**GetOrganizationCampusGatewayClusters**](ClustersApi.md#GetOrganizationCampusGatewayClusters) | **Get** /organizations/{organizationId}/campusGateway/clusters | Get the details of campus gateway clusters
[**GetOrganizationCampusGatewayClustersFailoverTargets**](ClustersApi.md#GetOrganizationCampusGatewayClustersFailoverTargets) | **Get** /organizations/{organizationId}/campusGateway/clusters/failover/targets | List the details of a Failover Targets for a Campus Gateway cluster
[**GetOrganizationCampusGatewayClustersFailoverTargetsByCluster**](ClustersApi.md#GetOrganizationCampusGatewayClustersFailoverTargetsByCluster) | **Get** /organizations/{organizationId}/campusGateway/clusters/failover/targets/byCluster | Get available backup cluster targets for campus gateway failover configuration
[**GetOrganizationCampusGatewayClustersNetworksOverviews**](ClustersApi.md#GetOrganizationCampusGatewayClustersNetworksOverviews) | **Get** /organizations/{organizationId}/campusGateway/clusters/networks/overviews | List networks tunneling through Campus Gateway clusters with their AP, ssids and client counts
[**GetOrganizationCampusGatewayClustersSsids**](ClustersApi.md#GetOrganizationCampusGatewayClustersSsids) | **Get** /organizations/{organizationId}/campusGateway/clusters/ssids | List SSIDs tunneling through Campus Gateway clusters
[**GetOrganizationCampusGatewayClustersTunnelable**](ClustersApi.md#GetOrganizationCampusGatewayClustersTunnelable) | **Get** /organizations/{organizationId}/campusGateway/clusters/tunnelable | List available campus gateway clusters that can be used for wireless network tunneling
[**GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork**](ClustersApi.md#GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork) | **Get** /organizations/{organizationId}/campusGateway/clusters/tunneling/byCluster/byNetwork | List all the MCG cluster-network tunnel settings
[**ProvisionOrganizationCampusGatewayClusters**](ClustersApi.md#ProvisionOrganizationCampusGatewayClusters) | **Post** /organizations/{organizationId}/campusGateway/clusters/provision | Provisions a cluster,adds campus gateways to it and associate/dissociate failover targets.
[**UpdateNetworkCampusGatewayCluster**](ClustersApi.md#UpdateNetworkCampusGatewayCluster) | **Put** /networks/{networkId}/campusGateway/clusters/{clusterId} | Update a cluster and add/remove campus gateways to/from it



## BatchOrganizationCampusGatewayClustersTunnelingUpdate

> InlineResponse200325 BatchOrganizationCampusGatewayClustersTunnelingUpdate(ctx, organizationId).BatchOrganizationCampusGatewayClustersTunnelingUpdate(batchOrganizationCampusGatewayClustersTunnelingUpdate).Execute()

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
    batchOrganizationCampusGatewayClustersTunnelingUpdate := *openapiclient.NewInlineObject274([]openapiclient.OrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateItems{*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateItems(*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateCluster("Id_example"), *openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateNetwork("Id_example"), *openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateData(*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersTunnelingBatchUpdateDataEncryption(false)))}) // InlineObject274 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.BatchOrganizationCampusGatewayClustersTunnelingUpdate(context.Background(), organizationId).BatchOrganizationCampusGatewayClustersTunnelingUpdate(batchOrganizationCampusGatewayClustersTunnelingUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.BatchOrganizationCampusGatewayClustersTunnelingUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationCampusGatewayClustersTunnelingUpdate`: InlineResponse200325
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.BatchOrganizationCampusGatewayClustersTunnelingUpdate`: %v\n", resp)
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

 **batchOrganizationCampusGatewayClustersTunnelingUpdate** | [**InlineObject274**](InlineObject274.md) |  | 

### Return type

[**InlineResponse200325**](InlineResponse200325.md)

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
    createNetworkCampusGatewayCluster := *openapiclient.NewInlineObject104("Name_example", []openapiclient.NetworksNetworkIdCampusGatewayClustersUplinks{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersUplinks("Interface_example", int32(123), []openapiclient.NetworksNetworkIdCampusGatewayClustersAddresses{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersAddresses("AssignmentMode_example")})}, []openapiclient.NetworksNetworkIdCampusGatewayClustersTunnels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersTunnels()}, *openapiclient.NewNetworksNetworkIdCampusGatewayClustersNameservers([]string{"Addresses_example"}), []openapiclient.NetworksNetworkIdCampusGatewayClustersPortChannels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersPortChannels("Name_example", int32(123), "AllowedVlans_example")}) // InlineObject104 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.CreateNetworkCampusGatewayCluster(context.Background(), networkId).CreateNetworkCampusGatewayCluster(createNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCampusGatewayCluster`: InlineResponse20118
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateNetworkCampusGatewayCluster`: %v\n", resp)
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

 **createNetworkCampusGatewayCluster** | [**InlineObject104**](InlineObject104.md) |  | 

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
    resp, r, err := apiClient.ClustersApi.DeleteNetworkCampusGatewayCluster(context.Background(), networkId, clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteNetworkCampusGatewayCluster``: %v\n", err)
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


## GetOrganizationCampusGatewayClusters

> InlineResponse200319 GetOrganizationCampusGatewayClusters(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ClustersApi.GetOrganizationCampusGatewayClusters(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCampusGatewayClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClusters`: InlineResponse200319
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCampusGatewayClusters`: %v\n", resp)
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

[**InlineResponse200319**](InlineResponse200319.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersFailoverTargets

> []InlineResponse200320 GetOrganizationCampusGatewayClustersFailoverTargets(ctx, organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ClustersApi.GetOrganizationCampusGatewayClustersFailoverTargets(context.Background(), organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCampusGatewayClustersFailoverTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersFailoverTargets`: []InlineResponse200320
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCampusGatewayClustersFailoverTargets`: %v\n", resp)
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

[**[]InlineResponse200320**](InlineResponse200320.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersFailoverTargetsByCluster

> InlineResponse200321 GetOrganizationCampusGatewayClustersFailoverTargetsByCluster(ctx, organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ClustersApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster(context.Background(), organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: InlineResponse200321
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: %v\n", resp)
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

[**InlineResponse200321**](InlineResponse200321.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersNetworksOverviews

> InlineResponse200322 GetOrganizationCampusGatewayClustersNetworksOverviews(ctx, organizationId).ClusterIds(clusterIds).SiteIds(siteIds).NetworkIds(networkIds).TunnelingSources(tunnelingSources).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ClustersApi.GetOrganizationCampusGatewayClustersNetworksOverviews(context.Background(), organizationId).ClusterIds(clusterIds).SiteIds(siteIds).NetworkIds(networkIds).TunnelingSources(tunnelingSources).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCampusGatewayClustersNetworksOverviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersNetworksOverviews`: InlineResponse200322
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCampusGatewayClustersNetworksOverviews`: %v\n", resp)
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

[**InlineResponse200322**](InlineResponse200322.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersSsids

> InlineResponse200323 GetOrganizationCampusGatewayClustersSsids(ctx, organizationId).ClusterIds(clusterIds).NetworkIds(networkIds).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ClustersApi.GetOrganizationCampusGatewayClustersSsids(context.Background(), organizationId).ClusterIds(clusterIds).NetworkIds(networkIds).Search(search).SortBy(sortBy).SortOrder(sortOrder).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCampusGatewayClustersSsids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersSsids`: InlineResponse200323
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCampusGatewayClustersSsids`: %v\n", resp)
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

[**InlineResponse200323**](InlineResponse200323.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersTunnelable

> InlineResponse200324 GetOrganizationCampusGatewayClustersTunnelable(ctx, organizationId).FromNetworkIds(fromNetworkIds).ExcludeSsidNumbers(excludeSsidNumbers).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ClustersApi.GetOrganizationCampusGatewayClustersTunnelable(context.Background(), organizationId).FromNetworkIds(fromNetworkIds).ExcludeSsidNumbers(excludeSsidNumbers).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCampusGatewayClustersTunnelable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersTunnelable`: InlineResponse200324
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCampusGatewayClustersTunnelable`: %v\n", resp)
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

[**InlineResponse200324**](InlineResponse200324.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork

> InlineResponse200326 GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork(ctx, organizationId).ClusterIds(clusterIds).DataEncryptionEnabled(dataEncryptionEnabled).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ClustersApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork(context.Background(), organizationId).ClusterIds(clusterIds).DataEncryptionEnabled(dataEncryptionEnabled).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork`: InlineResponse200326
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork`: %v\n", resp)
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

[**InlineResponse200326**](InlineResponse200326.md)

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
    provisionOrganizationCampusGatewayClusters := *openapiclient.NewInlineObject273("ClusterId_example", *openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersProvisionNetwork("Id_example"), "Name_example", []openapiclient.NetworksNetworkIdCampusGatewayClustersUplinks{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersUplinks("Interface_example", int32(123), []openapiclient.NetworksNetworkIdCampusGatewayClustersAddresses{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersAddresses("AssignmentMode_example")})}, []openapiclient.NetworksNetworkIdCampusGatewayClustersTunnels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersTunnels()}, *openapiclient.NewNetworksNetworkIdCampusGatewayClustersNameservers([]string{"Addresses_example"}), []openapiclient.OrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels{*openapiclient.NewOrganizationsOrganizationIdCampusGatewayClustersProvisionPortChannels("Name_example", int32(123), "AllowedVlans_example")}) // InlineObject273 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.ProvisionOrganizationCampusGatewayClusters(context.Background(), organizationId).ProvisionOrganizationCampusGatewayClusters(provisionOrganizationCampusGatewayClusters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ProvisionOrganizationCampusGatewayClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionOrganizationCampusGatewayClusters`: InlineResponse2024
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ProvisionOrganizationCampusGatewayClusters`: %v\n", resp)
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

 **provisionOrganizationCampusGatewayClusters** | [**InlineObject273**](InlineObject273.md) |  | 

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
    updateNetworkCampusGatewayCluster := *openapiclient.NewInlineObject105() // InlineObject105 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.UpdateNetworkCampusGatewayCluster(context.Background(), networkId, clusterId).UpdateNetworkCampusGatewayCluster(updateNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCampusGatewayCluster`: InlineResponse20118
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateNetworkCampusGatewayCluster`: %v\n", resp)
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


 **updateNetworkCampusGatewayCluster** | [**InlineObject105**](InlineObject105.md) |  | 

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

