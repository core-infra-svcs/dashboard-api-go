# \RequestsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkPiiRequest**](RequestsApi.md#CreateNetworkPiiRequest) | **Post** /networks/{networkId}/pii/requests | Submit a new delete or restrict processing PII request
[**DeleteNetworkPiiRequest**](RequestsApi.md#DeleteNetworkPiiRequest) | **Delete** /networks/{networkId}/pii/requests/{requestId} | Delete a restrict processing PII request
[**GetNetworkPiiRequest**](RequestsApi.md#GetNetworkPiiRequest) | **Get** /networks/{networkId}/pii/requests/{requestId} | Return a PII request
[**GetNetworkPiiRequests**](RequestsApi.md#GetNetworkPiiRequests) | **Get** /networks/{networkId}/pii/requests | List the PII requests for this network or organization



## CreateNetworkPiiRequest

> InlineResponse200114 CreateNetworkPiiRequest(ctx, networkId).CreateNetworkPiiRequest(createNetworkPiiRequest).Execute()

Submit a new delete or restrict processing PII request



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
    createNetworkPiiRequest := *openapiclient.NewInlineObject116() // InlineObject116 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.CreateNetworkPiiRequest(context.Background(), networkId).CreateNetworkPiiRequest(createNetworkPiiRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.CreateNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPiiRequest`: InlineResponse200114
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.CreateNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPiiRequest** | [**InlineObject116**](InlineObject116.md) |  | 

### Return type

[**InlineResponse200114**](InlineResponse200114.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPiiRequest

> DeleteNetworkPiiRequest(ctx, networkId, requestId).Execute()

Delete a restrict processing PII request



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
    requestId := "requestId_example" // string | Request ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.DeleteNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiRequest

> InlineResponse200114 GetNetworkPiiRequest(ctx, networkId, requestId).Execute()

Return a PII request



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
    requestId := "requestId_example" // string | Request ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.GetNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiRequest`: InlineResponse200114
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.GetNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200114**](InlineResponse200114.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiRequests

> []InlineResponse200114 GetNetworkPiiRequests(ctx, networkId).Execute()

List the PII requests for this network or organization



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
    resp, r, err := apiClient.RequestsApi.GetNetworkPiiRequests(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.GetNetworkPiiRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiRequests`: []InlineResponse200114
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.GetNetworkPiiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200114**](InlineResponse200114.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

