# \IpsecApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas**](IpsecApi.md#GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas) | **Get** /organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas | Get the list of available IPsec SLA policies for an organization
[**UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas**](IpsecApi.md#UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas) | **Put** /organizations/{organizationId}/appliance/vpn/siteToSite/ipsec/peers/slas | Update the IPsec SLA policies for an organization



## GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas

> InlineResponse200233 GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx, organizationId).Execute()

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
    resp, r, err := apiClient.IpsecApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpsecApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: InlineResponse200233
    fmt.Fprintf(os.Stdout, "Response from `IpsecApi.GetOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: %v\n", resp)
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

[**InlineResponse200233**](InlineResponse200233.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas

> InlineResponse200234 UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(ctx, organizationId).UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas).Execute()

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
    updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas := *openapiclient.NewInlineObject233() // InlineObject233 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IpsecApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(context.Background(), organizationId).UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas(updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpsecApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: InlineResponse200234
    fmt.Fprintf(os.Stdout, "Response from `IpsecApi.UpdateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas`: %v\n", resp)
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

 **updateOrganizationApplianceVpnSiteToSiteIpsecPeersSlas** | [**InlineObject233**](InlineObject233.md) |  | 

### Return type

[**InlineResponse200234**](InlineResponse200234.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

