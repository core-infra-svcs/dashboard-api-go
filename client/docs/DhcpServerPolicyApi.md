# \DhcpServerPolicyApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchDhcpServerPolicy**](DhcpServerPolicyApi.md#GetNetworkSwitchDhcpServerPolicy) | **Get** /networks/{networkId}/switch/dhcpServerPolicy | Return the DHCP server settings
[**UpdateNetworkSwitchDhcpServerPolicy**](DhcpServerPolicyApi.md#UpdateNetworkSwitchDhcpServerPolicy) | **Put** /networks/{networkId}/switch/dhcpServerPolicy | Update the DHCP server settings



## GetNetworkSwitchDhcpServerPolicy

> map[string]interface{} GetNetworkSwitchDhcpServerPolicy(ctx, networkId).Execute()

Return the DHCP server settings



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
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchDhcpServerPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDhcpServerPolicy

> map[string]interface{} UpdateNetworkSwitchDhcpServerPolicy(ctx, networkId).UpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy).Execute()

Update the DHCP server settings



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
    networkId := "networkId_example" // string | 
    updateNetworkSwitchDhcpServerPolicy := *openapiclient.NewInlineObject95() // InlineObject95 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicy(context.Background(), networkId).UpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchDhcpServerPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchDhcpServerPolicy** | [**InlineObject95**](InlineObject95.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

