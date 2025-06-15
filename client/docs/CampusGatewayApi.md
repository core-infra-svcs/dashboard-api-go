# \CampusGatewayApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCampusGatewayCluster**](CampusGatewayApi.md#CreateNetworkCampusGatewayCluster) | **Post** /networks/{networkId}/campusGateway/clusters | Create a cluster and add campus gateways to it
[**GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice**](CampusGatewayApi.md#GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice) | **Get** /organizations/{organizationId}/campusGateway/devices/uplinks/localOverrides/byDevice | Uplink overrides configured locally on Campus Gateway devices in an organization.
[**UpdateNetworkCampusGatewayCluster**](CampusGatewayApi.md#UpdateNetworkCampusGatewayCluster) | **Put** /networks/{networkId}/campusGateway/clusters/{clusterId} | Update a cluster and add/remove campus gateways to/from it



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
    resp, r, err := apiClient.CampusGatewayApi.CreateNetworkCampusGatewayCluster(context.Background(), networkId).CreateNetworkCampusGatewayCluster(createNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.CreateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCampusGatewayCluster`: InlineResponse20110
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


## GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice

> InlineResponse200252 GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    // response from `GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice`: InlineResponse200252
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

[**InlineResponse200252**](InlineResponse200252.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := apiClient.CampusGatewayApi.UpdateNetworkCampusGatewayCluster(context.Background(), networkId, clusterId).UpdateNetworkCampusGatewayCluster(updateNetworkCampusGatewayCluster).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampusGatewayApi.UpdateNetworkCampusGatewayCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCampusGatewayCluster`: InlineResponse20110
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

