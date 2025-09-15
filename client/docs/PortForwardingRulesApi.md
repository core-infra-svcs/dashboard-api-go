# \PortForwardingRulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularGatewayPortForwardingRules**](PortForwardingRulesApi.md#GetDeviceCellularGatewayPortForwardingRules) | **Get** /devices/{serial}/cellularGateway/portForwardingRules | Returns the port forwarding rules for a single MG.
[**GetNetworkApplianceFirewallPortForwardingRules**](PortForwardingRulesApi.md#GetNetworkApplianceFirewallPortForwardingRules) | **Get** /networks/{networkId}/appliance/firewall/portForwardingRules | Return the port forwarding rules for an MX network
[**UpdateDeviceCellularGatewayPortForwardingRules**](PortForwardingRulesApi.md#UpdateDeviceCellularGatewayPortForwardingRules) | **Put** /devices/{serial}/cellularGateway/portForwardingRules | Updates the port forwarding rules for a single MG.
[**UpdateNetworkApplianceFirewallPortForwardingRules**](PortForwardingRulesApi.md#UpdateNetworkApplianceFirewallPortForwardingRules) | **Put** /networks/{networkId}/appliance/firewall/portForwardingRules | Update the port forwarding rules for an MX network



## GetDeviceCellularGatewayPortForwardingRules

> InlineResponse20020 GetDeviceCellularGatewayPortForwardingRules(ctx, serial).Execute()

Returns the port forwarding rules for a single MG.



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortForwardingRulesApi.GetDeviceCellularGatewayPortForwardingRules(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.GetDeviceCellularGatewayPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularGatewayPortForwardingRules`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.GetDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallPortForwardingRules

> InlineResponse20060 GetNetworkApplianceFirewallPortForwardingRules(ctx, networkId).Execute()

Return the port forwarding rules for an MX network



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortForwardingRulesApi.GetNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.GetNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallPortForwardingRules`: InlineResponse20060
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.GetNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20060**](InlineResponse20060.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGatewayPortForwardingRules

> InlineResponse20020 UpdateDeviceCellularGatewayPortForwardingRules(ctx, serial).UpdateDeviceCellularGatewayPortForwardingRules(updateDeviceCellularGatewayPortForwardingRules).Execute()

Updates the port forwarding rules for a single MG.



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
    serial := "serial_example" // string | Serial
    updateDeviceCellularGatewayPortForwardingRules := *openapiclient.NewInlineObject15() // InlineObject15 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortForwardingRulesApi.UpdateDeviceCellularGatewayPortForwardingRules(context.Background(), serial).UpdateDeviceCellularGatewayPortForwardingRules(updateDeviceCellularGatewayPortForwardingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.UpdateDeviceCellularGatewayPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularGatewayPortForwardingRules`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.UpdateDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayPortForwardingRules** | [**InlineObject15**](InlineObject15.md) |  | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallPortForwardingRules

> InlineResponse20060 UpdateNetworkApplianceFirewallPortForwardingRules(ctx, networkId).UpdateNetworkApplianceFirewallPortForwardingRules(updateNetworkApplianceFirewallPortForwardingRules).Execute()

Update the port forwarding rules for an MX network



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
    updateNetworkApplianceFirewallPortForwardingRules := *openapiclient.NewInlineObject53([]openapiclient.NetworksNetworkIdApplianceFirewallPortForwardingRulesRules{*openapiclient.NewNetworksNetworkIdApplianceFirewallPortForwardingRulesRules("LanIp_example", "PublicPort_example", "LocalPort_example", []string{"AllowedIps_example"}, "Protocol_example")}) // InlineObject53 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortForwardingRulesApi.UpdateNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).UpdateNetworkApplianceFirewallPortForwardingRules(updateNetworkApplianceFirewallPortForwardingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.UpdateNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallPortForwardingRules`: InlineResponse20060
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.UpdateNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallPortForwardingRules** | [**InlineObject53**](InlineObject53.md) |  | 

### Return type

[**InlineResponse20060**](InlineResponse20060.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

