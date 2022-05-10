# \SiteToSiteVpnApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceVpnSiteToSiteVpn**](SiteToSiteVpnApi.md#GetNetworkApplianceVpnSiteToSiteVpn) | **Get** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Return the site-to-site VPN settings of a network
[**UpdateNetworkApplianceVpnSiteToSiteVpn**](SiteToSiteVpnApi.md#UpdateNetworkApplianceVpnSiteToSiteVpn) | **Put** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Update the site-to-site VPN settings of a network



## GetNetworkApplianceVpnSiteToSiteVpn

> map[string]interface{} GetNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).Execute()

Return the site-to-site VPN settings of a network



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
    resp, r, err := apiClient.SiteToSiteVpnApi.GetNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteVpnApi.GetNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVpnSiteToSiteVpn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SiteToSiteVpnApi.GetNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


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


## UpdateNetworkApplianceVpnSiteToSiteVpn

> map[string]interface{} UpdateNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).UpdateNetworkApplianceVpnSiteToSiteVpn(updateNetworkApplianceVpnSiteToSiteVpn).Execute()

Update the site-to-site VPN settings of a network



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
    updateNetworkApplianceVpnSiteToSiteVpn := *openapiclient.NewInlineObject51("Mode_example") // InlineObject51 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteToSiteVpnApi.UpdateNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).UpdateNetworkApplianceVpnSiteToSiteVpn(updateNetworkApplianceVpnSiteToSiteVpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteVpnApi.UpdateNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVpnSiteToSiteVpn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SiteToSiteVpnApi.UpdateNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnSiteToSiteVpn** | [**InlineObject51**](InlineObject51.md) |  | 

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

