# \NetflowApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkNetflow**](NetflowApi.md#GetNetworkNetflow) | **Get** /networks/{networkId}/netflow | Return the NetFlow traffic reporting settings for a network
[**UpdateNetworkNetflow**](NetflowApi.md#UpdateNetworkNetflow) | **Put** /networks/{networkId}/netflow | Update the NetFlow traffic reporting settings for a network



## GetNetworkNetflow

> InlineResponse200111 GetNetworkNetflow(ctx, networkId).Execute()

Return the NetFlow traffic reporting settings for a network



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
    resp, r, err := apiClient.NetflowApi.GetNetworkNetflow(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetflowApi.GetNetworkNetflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkNetflow`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `NetflowApi.GetNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkNetflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkNetflow

> InlineResponse200111 UpdateNetworkNetflow(ctx, networkId).UpdateNetworkNetflow(updateNetworkNetflow).Execute()

Update the NetFlow traffic reporting settings for a network



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
    updateNetworkNetflow := *openapiclient.NewInlineObject117() // InlineObject117 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetflowApi.UpdateNetworkNetflow(context.Background(), networkId).UpdateNetworkNetflow(updateNetworkNetflow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetflowApi.UpdateNetworkNetflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkNetflow`: InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `NetflowApi.UpdateNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkNetflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkNetflow** | [**InlineObject117**](InlineObject117.md) |  | 

### Return type

[**InlineResponse200111**](InlineResponse200111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

