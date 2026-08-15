# \ConnectorsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchOrganizationSaseConnectorsDelete**](ConnectorsApi.md#BatchOrganizationSaseConnectorsDelete) | **Post** /organizations/{organizationId}/sase/connectors/batchDelete | Delete SSE Connectors by ID
[**GetOrganizationSaseConnectors**](ConnectorsApi.md#GetOrganizationSaseConnectors) | **Get** /organizations/{organizationId}/sase/connectors | List SSE Connectors for an organization



## BatchOrganizationSaseConnectorsDelete

> InlineResponse2025 BatchOrganizationSaseConnectorsDelete(ctx, organizationId).BatchOrganizationSaseConnectorsDelete(batchOrganizationSaseConnectorsDelete).Execute()

Delete SSE Connectors by ID



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
    batchOrganizationSaseConnectorsDelete := *openapiclient.NewInlineObject341() // InlineObject341 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectorsApi.BatchOrganizationSaseConnectorsDelete(context.Background(), organizationId).BatchOrganizationSaseConnectorsDelete(batchOrganizationSaseConnectorsDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsApi.BatchOrganizationSaseConnectorsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationSaseConnectorsDelete`: InlineResponse2025
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsApi.BatchOrganizationSaseConnectorsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchOrganizationSaseConnectorsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchOrganizationSaseConnectorsDelete** | [**InlineObject341**](InlineObject341.md) |  | 

### Return type

[**InlineResponse2025**](InlineResponse2025.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseConnectors

> InlineResponse200416 GetOrganizationSaseConnectors(ctx, organizationId).Execute()

List SSE Connectors for an organization



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
    resp, r, err := apiClient.ConnectorsApi.GetOrganizationSaseConnectors(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsApi.GetOrganizationSaseConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseConnectors`: InlineResponse200416
    fmt.Fprintf(os.Stdout, "Response from `ConnectorsApi.GetOrganizationSaseConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200416**](InlineResponse200416.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

