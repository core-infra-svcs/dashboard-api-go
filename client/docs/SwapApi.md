# \SwapApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCellularGatewayEsimsSwap**](SwapApi.md#CreateOrganizationCellularGatewayEsimsSwap) | **Post** /organizations/{organizationId}/cellularGateway/esims/swap | Swap which profile an eSIM uses.
[**UpdateOrganizationCellularGatewayEsimsSwap**](SwapApi.md#UpdateOrganizationCellularGatewayEsimsSwap) | **Put** /organizations/{organizationId}/cellularGateway/esims/swap/{id} | Get the status of a profile swap.



## CreateOrganizationCellularGatewayEsimsSwap

> InlineResponse200259 CreateOrganizationCellularGatewayEsimsSwap(ctx, organizationId).CreateOrganizationCellularGatewayEsimsSwap(createOrganizationCellularGatewayEsimsSwap).Execute()

Swap which profile an eSIM uses.



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
    createOrganizationCellularGatewayEsimsSwap := *openapiclient.NewInlineObject248([]openapiclient.OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps{*openapiclient.NewOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps("Eid_example")}) // InlineObject248 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwapApi.CreateOrganizationCellularGatewayEsimsSwap(context.Background(), organizationId).CreateOrganizationCellularGatewayEsimsSwap(createOrganizationCellularGatewayEsimsSwap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwapApi.CreateOrganizationCellularGatewayEsimsSwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCellularGatewayEsimsSwap`: InlineResponse200259
    fmt.Fprintf(os.Stdout, "Response from `SwapApi.CreateOrganizationCellularGatewayEsimsSwap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCellularGatewayEsimsSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCellularGatewayEsimsSwap** | [**InlineObject248**](InlineObject248.md) |  | 

### Return type

[**InlineResponse200259**](InlineResponse200259.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCellularGatewayEsimsSwap

> InlineResponse200259 UpdateOrganizationCellularGatewayEsimsSwap(ctx, id, organizationId).Execute()

Get the status of a profile swap.



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
    id := "id_example" // string | eSIM EID
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SwapApi.UpdateOrganizationCellularGatewayEsimsSwap(context.Background(), id, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwapApi.UpdateOrganizationCellularGatewayEsimsSwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsSwap`: InlineResponse200259
    fmt.Fprintf(os.Stdout, "Response from `SwapApi.UpdateOrganizationCellularGatewayEsimsSwap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | eSIM EID | 
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCellularGatewayEsimsSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200259**](InlineResponse200259.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

