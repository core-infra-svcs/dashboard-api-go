# \ByClusterApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCampusGatewayClientsUsageByNetworkByCluster**](ByClusterApi.md#GetOrganizationCampusGatewayClientsUsageByNetworkByCluster) | **Get** /organizations/{organizationId}/campusGateway/clients/usage/byNetwork/byCluster | Returns client usage details for campus gateway clusters within an organization.
[**GetOrganizationCampusGatewayClustersFailoverTargetsByCluster**](ByClusterApi.md#GetOrganizationCampusGatewayClustersFailoverTargetsByCluster) | **Get** /organizations/{organizationId}/campusGateway/clusters/failover/targets/byCluster | Get available backup cluster targets for campus gateway failover configuration
[**GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork**](ByClusterApi.md#GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork) | **Get** /organizations/{organizationId}/campusGateway/clusters/tunneling/byCluster/byNetwork | List all the MCG cluster-network tunnel settings



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
    resp, r, err := apiClient.ByClusterApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).NetworkGroupIds(networkGroupIds).ClusterIds(clusterIds).UsageUnits(usageUnits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ByClusterApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClientsUsageByNetworkByCluster`: InlineResponse200331
    fmt.Fprintf(os.Stdout, "Response from `ByClusterApi.GetOrganizationCampusGatewayClientsUsageByNetworkByCluster`: %v\n", resp)
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
    resp, r, err := apiClient.ByClusterApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster(context.Background(), organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ByClusterApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: InlineResponse200334
    fmt.Fprintf(os.Stdout, "Response from `ByClusterApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: %v\n", resp)
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
    resp, r, err := apiClient.ByClusterApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork(context.Background(), organizationId).ClusterIds(clusterIds).DataEncryptionEnabled(dataEncryptionEnabled).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ByClusterApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork`: InlineResponse200339
    fmt.Fprintf(os.Stdout, "Response from `ByClusterApi.GetOrganizationCampusGatewayClustersTunnelingByClusterByNetwork`: %v\n", resp)
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

