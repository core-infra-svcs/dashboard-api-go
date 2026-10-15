# \SplashAuthorizationStatusApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkClientSplashAuthorizationStatus**](SplashAuthorizationStatusApi.md#GetNetworkClientSplashAuthorizationStatus) | **Get** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Return the splash authorization for a client, for each SSID they&#39;ve associated with through splash
[**UpdateNetworkClientSplashAuthorizationStatus**](SplashAuthorizationStatusApi.md#UpdateNetworkClientSplashAuthorizationStatus) | **Put** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Update a client&#39;s splash authorization



## GetNetworkClientSplashAuthorizationStatus

> InlineResponse200129 GetNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).Execute()

Return the splash authorization for a client, for each SSID they've associated with through splash



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
    clientId := "clientId_example" // string | Client ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SplashAuthorizationStatusApi.GetNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SplashAuthorizationStatusApi.GetNetworkClientSplashAuthorizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientSplashAuthorizationStatus`: InlineResponse200129
    fmt.Fprintf(os.Stdout, "Response from `SplashAuthorizationStatusApi.GetNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200129**](InlineResponse200129.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkClientSplashAuthorizationStatus

> InlineResponse200129 UpdateNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).UpdateNetworkClientSplashAuthorizationStatus(updateNetworkClientSplashAuthorizationStatus).Execute()

Update a client's splash authorization



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
    clientId := "clientId_example" // string | Client ID
    updateNetworkClientSplashAuthorizationStatus := *openapiclient.NewInlineObject114(*openapiclient.NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids()) // InlineObject114 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SplashAuthorizationStatusApi.UpdateNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).UpdateNetworkClientSplashAuthorizationStatus(updateNetworkClientSplashAuthorizationStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SplashAuthorizationStatusApi.UpdateNetworkClientSplashAuthorizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkClientSplashAuthorizationStatus`: InlineResponse200129
    fmt.Fprintf(os.Stdout, "Response from `SplashAuthorizationStatusApi.UpdateNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientSplashAuthorizationStatus** | [**InlineObject114**](InlineObject114.md) |  | 

### Return type

[**InlineResponse200129**](InlineResponse200129.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

