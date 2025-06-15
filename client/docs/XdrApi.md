# \XdrApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableOrganizationIntegrationsXdrNetworks**](XdrApi.md#DisableOrganizationIntegrationsXdrNetworks) | **Post** /organizations/{organizationId}/integrations/xdr/networks/disable | Disable XDR on networks
[**EnableOrganizationIntegrationsXdrNetworks**](XdrApi.md#EnableOrganizationIntegrationsXdrNetworks) | **Post** /organizations/{organizationId}/integrations/xdr/networks/enable | Enable XDR on networks
[**GetOrganizationIntegrationsXdrNetworks**](XdrApi.md#GetOrganizationIntegrationsXdrNetworks) | **Get** /organizations/{organizationId}/integrations/xdr/networks | Returns the networks in the organization that have XDR enabled



## DisableOrganizationIntegrationsXdrNetworks

> InlineResponse200293 DisableOrganizationIntegrationsXdrNetworks(ctx, organizationId).DisableOrganizationIntegrationsXdrNetworks(disableOrganizationIntegrationsXdrNetworks).Execute()

Disable XDR on networks



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
    disableOrganizationIntegrationsXdrNetworks := *openapiclient.NewInlineObject268([]openapiclient.OrganizationsOrganizationIdIntegrationsXdrNetworksDisableNetworks{*openapiclient.NewOrganizationsOrganizationIdIntegrationsXdrNetworksDisableNetworks("NetworkId_example", []string{"ProductTypes_example"})}) // InlineObject268 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XdrApi.DisableOrganizationIntegrationsXdrNetworks(context.Background(), organizationId).DisableOrganizationIntegrationsXdrNetworks(disableOrganizationIntegrationsXdrNetworks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XdrApi.DisableOrganizationIntegrationsXdrNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableOrganizationIntegrationsXdrNetworks`: InlineResponse200293
    fmt.Fprintf(os.Stdout, "Response from `XdrApi.DisableOrganizationIntegrationsXdrNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableOrganizationIntegrationsXdrNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disableOrganizationIntegrationsXdrNetworks** | [**InlineObject268**](InlineObject268.md) |  | 

### Return type

[**InlineResponse200293**](InlineResponse200293.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableOrganizationIntegrationsXdrNetworks

> InlineResponse200294 EnableOrganizationIntegrationsXdrNetworks(ctx, organizationId).EnableOrganizationIntegrationsXdrNetworks(enableOrganizationIntegrationsXdrNetworks).Execute()

Enable XDR on networks



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
    enableOrganizationIntegrationsXdrNetworks := *openapiclient.NewInlineObject269([]openapiclient.OrganizationsOrganizationIdIntegrationsXdrNetworksEnableNetworks{*openapiclient.NewOrganizationsOrganizationIdIntegrationsXdrNetworksEnableNetworks("NetworkId_example", []string{"ProductTypes_example"})}) // InlineObject269 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XdrApi.EnableOrganizationIntegrationsXdrNetworks(context.Background(), organizationId).EnableOrganizationIntegrationsXdrNetworks(enableOrganizationIntegrationsXdrNetworks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XdrApi.EnableOrganizationIntegrationsXdrNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableOrganizationIntegrationsXdrNetworks`: InlineResponse200294
    fmt.Fprintf(os.Stdout, "Response from `XdrApi.EnableOrganizationIntegrationsXdrNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableOrganizationIntegrationsXdrNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableOrganizationIntegrationsXdrNetworks** | [**InlineObject269**](InlineObject269.md) |  | 

### Return type

[**InlineResponse200294**](InlineResponse200294.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationIntegrationsXdrNetworks

> InlineResponse200292 GetOrganizationIntegrationsXdrNetworks(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns the networks in the organization that have XDR enabled



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 20. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XdrApi.GetOrganizationIntegrationsXdrNetworks(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XdrApi.GetOrganizationIntegrationsXdrNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationIntegrationsXdrNetworks`: InlineResponse200292
    fmt.Fprintf(os.Stdout, "Response from `XdrApi.GetOrganizationIntegrationsXdrNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationIntegrationsXdrNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 20. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200292**](InlineResponse200292.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

