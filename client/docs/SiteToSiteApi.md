# \SiteToSiteApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas**](SiteToSiteApi.md#GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas) | **Get** /organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas | Get the list of available IPsec SLA policies for an organization
[**UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas**](SiteToSiteApi.md#UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas) | **Put** /organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas | Update the IPsec SLA policies for an organization



## GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas

> InlineResponse200235 GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx, organizationId).Execute()

Get the list of available IPsec SLA policies for an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteToSiteApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: InlineResponse200235
    fmt.Fprintf(os.Stdout, "Response from `SiteToSiteApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200235**](InlineResponse200235.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas

> InlineResponse200236 UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx, organizationId).UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas).Execute()

Update the IPsec SLA policies for an organization



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
    updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas := *openapiclient.NewInlineObject234() // InlineObject234 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiteToSiteApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(context.Background(), organizationId).UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: InlineResponse200236
    fmt.Fprintf(os.Stdout, "Response from `SiteToSiteApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas** | [**InlineObject234**](InlineObject234.md) |  | 

### Return type

[**InlineResponse200236**](InlineResponse200236.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

