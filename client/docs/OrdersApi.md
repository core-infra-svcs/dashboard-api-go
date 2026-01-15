# \OrdersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimOrganizationInventoryOrders**](OrdersApi.md#ClaimOrganizationInventoryOrders) | **Post** /organizations/{organizationId}/inventory/orders/claim | Claim an order by the secure unique order claim number, the order claim id
[**PreviewOrganizationInventoryOrders**](OrdersApi.md#PreviewOrganizationInventoryOrders) | **Post** /organizations/{organizationId}/inventory/orders/preview | Preview the results and status of an order claim by the secure order id



## ClaimOrganizationInventoryOrders

> InlineResponse200310 ClaimOrganizationInventoryOrders(ctx, organizationId).ClaimOrganizationInventoryOrders(claimOrganizationInventoryOrders).Execute()

Claim an order by the secure unique order claim number, the order claim id



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
    claimOrganizationInventoryOrders := *openapiclient.NewInlineObject278("ClaimId_example") // InlineObject278 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.ClaimOrganizationInventoryOrders(context.Background(), organizationId).ClaimOrganizationInventoryOrders(claimOrganizationInventoryOrders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ClaimOrganizationInventoryOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimOrganizationInventoryOrders`: InlineResponse200310
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ClaimOrganizationInventoryOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimOrganizationInventoryOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimOrganizationInventoryOrders** | [**InlineObject278**](InlineObject278.md) |  | 

### Return type

[**InlineResponse200310**](InlineResponse200310.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewOrganizationInventoryOrders

> InlineResponse200311 PreviewOrganizationInventoryOrders(ctx, organizationId).PreviewOrganizationInventoryOrders(previewOrganizationInventoryOrders).Execute()

Preview the results and status of an order claim by the secure order id



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
    previewOrganizationInventoryOrders := *openapiclient.NewInlineObject279("ClaimId_example") // InlineObject279 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.PreviewOrganizationInventoryOrders(context.Background(), organizationId).PreviewOrganizationInventoryOrders(previewOrganizationInventoryOrders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.PreviewOrganizationInventoryOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewOrganizationInventoryOrders`: InlineResponse200311
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.PreviewOrganizationInventoryOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewOrganizationInventoryOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewOrganizationInventoryOrders** | [**InlineObject279**](InlineObject279.md) |  | 

### Return type

[**InlineResponse200311**](InlineResponse200311.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

