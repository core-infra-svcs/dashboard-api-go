# \TargetsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCampusGatewayClustersFailoverTargets**](TargetsApi.md#GetOrganizationCampusGatewayClustersFailoverTargets) | **Get** /organizations/{organizationId}/campusGateway/clusters/failover/targets | List the details of a Failover Targets for a Campus Gateway cluster
[**GetOrganizationCampusGatewayClustersFailoverTargetsByCluster**](TargetsApi.md#GetOrganizationCampusGatewayClustersFailoverTargetsByCluster) | **Get** /organizations/{organizationId}/campusGateway/clusters/failover/targets/byCluster | Get available backup cluster targets for campus gateway failover configuration



## GetOrganizationCampusGatewayClustersFailoverTargets

> []InlineResponse200318 GetOrganizationCampusGatewayClustersFailoverTargets(ctx, organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.TargetsApi.GetOrganizationCampusGatewayClustersFailoverTargets(context.Background(), organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetsApi.GetOrganizationCampusGatewayClustersFailoverTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersFailoverTargets`: []InlineResponse200318
    fmt.Fprintf(os.Stdout, "Response from `TargetsApi.GetOrganizationCampusGatewayClustersFailoverTargets`: %v\n", resp)
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

[**[]InlineResponse200318**](InlineResponse200318.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCampusGatewayClustersFailoverTargetsByCluster

> InlineResponse200319 GetOrganizationCampusGatewayClustersFailoverTargetsByCluster(ctx, organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.TargetsApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster(context.Background(), organizationId).NetworkIds(networkIds).ClusterIds(clusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetsApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: InlineResponse200319
    fmt.Fprintf(os.Stdout, "Response from `TargetsApi.GetOrganizationCampusGatewayClustersFailoverTargetsByCluster`: %v\n", resp)
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

[**InlineResponse200319**](InlineResponse200319.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

