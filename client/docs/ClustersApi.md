# \ClustersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCampusGatewayCluster**](ClustersApi.md#CreateNetworkCampusGatewayCluster) | **Post** /networks/{networkId}/campusGateway/clusters | Create a cluster and add campus gateways to it
[**UpdateNetworkCampusGatewayCluster**](ClustersApi.md#UpdateNetworkCampusGatewayCluster) | **Put** /networks/{networkId}/campusGateway/clusters/{clusterId} | Update a cluster and add/remove campus gateways to/from it



## CreateNetworkCampusGatewayCluster

> InlineResponse20110 CreateNetworkCampusGatewayCluster(ctx, networkId).CreateNetworkCampusGatewayCluster(createNetworkCampusGatewayCluster).Execute()

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
    createNetworkCampusGatewayCluster := *openapiclient.NewInlineObject85("Name_example", []openapiclient.NetworksNetworkIdCampusGatewayClustersUplinks{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersUplinks("Interface_example", int32(123), []openapiclient.NetworksNetworkIdCampusGatewayClustersAddresses{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersAddresses("AssignmentMode_example")})}, []openapiclient.NetworksNetworkIdCampusGatewayClustersTunnels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersTunnels()}, *openapiclient.NewNetworksNetworkIdCampusGatewayClustersNameservers(), []openapiclient.NetworksNetworkIdCampusGatewayClustersPortChannels{*openapiclient.NewNetworksNetworkIdCampusGatewayClustersPortChannels("Name_example", int32(123), "AllowedVlans_example")}) // InlineObject85 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.CreateNetworkCampusGatewayCluster(context.Background(), networkId).CreateNetworkCampusGatewayCluster(createNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCampusGatewayCluster`: InlineResponse20110
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

 **createNetworkCampusGatewayCluster** | [**InlineObject85**](InlineObject85.md) |  | 

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


## UpdateNetworkCampusGatewayCluster

> InlineResponse20110 UpdateNetworkCampusGatewayCluster(ctx, networkId, clusterId).UpdateNetworkCampusGatewayCluster(updateNetworkCampusGatewayCluster).Execute()

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
    updateNetworkCampusGatewayCluster := *openapiclient.NewInlineObject86() // InlineObject86 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClustersApi.UpdateNetworkCampusGatewayCluster(context.Background(), networkId, clusterId).UpdateNetworkCampusGatewayCluster(updateNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCampusGatewayCluster`: InlineResponse20110
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


 **updateNetworkCampusGatewayCluster** | [**InlineObject86**](InlineObject86.md) |  | 

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

