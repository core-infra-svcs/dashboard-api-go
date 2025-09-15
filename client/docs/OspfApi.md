# \OspfApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchRoutingOspf**](OspfApi.md#GetNetworkSwitchRoutingOspf) | **Get** /networks/{networkId}/switch/routing/ospf | Return layer 3 OSPF routing configuration
[**UpdateNetworkSwitchRoutingOspf**](OspfApi.md#UpdateNetworkSwitchRoutingOspf) | **Put** /networks/{networkId}/switch/routing/ospf | Update layer 3 OSPF routing configuration



## GetNetworkSwitchRoutingOspf

> InlineResponse200168 GetNetworkSwitchRoutingOspf(ctx, networkId).Vrf(vrf).Execute()

Return layer 3 OSPF routing configuration



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
    vrf := "vrf_example" // string | The VRF to return the OSPF routing configuration for. When not provided, the default VRF is used. Included on networks with IOS XE 17.18 or higher (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OspfApi.GetNetworkSwitchRoutingOspf(context.Background(), networkId).Vrf(vrf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OspfApi.GetNetworkSwitchRoutingOspf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchRoutingOspf`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `OspfApi.GetNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vrf** | **string** | The VRF to return the OSPF routing configuration for. When not provided, the default VRF is used. Included on networks with IOS XE 17.18 or higher | 

### Return type

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingOspf

> InlineResponse200168 UpdateNetworkSwitchRoutingOspf(ctx, networkId).Vrf(vrf).UpdateNetworkSwitchRoutingOspf(updateNetworkSwitchRoutingOspf).Execute()

Update layer 3 OSPF routing configuration



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
    vrf := "vrf_example" // string | The VRF to return the OSPF routing configuration for. When not provided, the default VRF is used. Requires IOS XE 17.18 or higher (optional)
    updateNetworkSwitchRoutingOspf := *openapiclient.NewInlineObject157() // InlineObject157 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OspfApi.UpdateNetworkSwitchRoutingOspf(context.Background(), networkId).Vrf(vrf).UpdateNetworkSwitchRoutingOspf(updateNetworkSwitchRoutingOspf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OspfApi.UpdateNetworkSwitchRoutingOspf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchRoutingOspf`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `OspfApi.UpdateNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vrf** | **string** | The VRF to return the OSPF routing configuration for. When not provided, the default VRF is used. Requires IOS XE 17.18 or higher | 
 **updateNetworkSwitchRoutingOspf** | [**InlineObject157**](InlineObject157.md) |  | 

### Return type

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

