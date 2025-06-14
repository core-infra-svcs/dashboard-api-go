# \LinkAggregationsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchLinkAggregation**](LinkAggregationsApi.md#CreateNetworkSwitchLinkAggregation) | **Post** /networks/{networkId}/switch/linkAggregations | Create a link aggregation group
[**DeleteNetworkSwitchLinkAggregation**](LinkAggregationsApi.md#DeleteNetworkSwitchLinkAggregation) | **Delete** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Split a link aggregation group into separate ports
[**GetNetworkSwitchLinkAggregations**](LinkAggregationsApi.md#GetNetworkSwitchLinkAggregations) | **Get** /networks/{networkId}/switch/linkAggregations | List link aggregation groups
[**UpdateNetworkSwitchLinkAggregation**](LinkAggregationsApi.md#UpdateNetworkSwitchLinkAggregation) | **Put** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Update a link aggregation group



## CreateNetworkSwitchLinkAggregation

> InlineResponse200158 CreateNetworkSwitchLinkAggregation(ctx, networkId).CreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation).Execute()

Create a link aggregation group



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
    createNetworkSwitchLinkAggregation := *openapiclient.NewInlineObject145() // InlineObject145 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkAggregationsApi.CreateNetworkSwitchLinkAggregation(context.Background(), networkId).CreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.CreateNetworkSwitchLinkAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchLinkAggregation`: InlineResponse200158
    fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsApi.CreateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchLinkAggregation** | [**InlineObject145**](InlineObject145.md) |  | 

### Return type

[**InlineResponse200158**](InlineResponse200158.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchLinkAggregation

> DeleteNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).Execute()

Split a link aggregation group into separate ports



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
    linkAggregationId := "linkAggregationId_example" // string | Link aggregation ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkAggregationsApi.DeleteNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.DeleteNetworkSwitchLinkAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**linkAggregationId** | **string** | Link aggregation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchLinkAggregationRequest struct via the builder pattern


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


## GetNetworkSwitchLinkAggregations

> []InlineResponse200158 GetNetworkSwitchLinkAggregations(ctx, networkId).Execute()

List link aggregation groups



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
    resp, r, err := apiClient.LinkAggregationsApi.GetNetworkSwitchLinkAggregations(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.GetNetworkSwitchLinkAggregations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchLinkAggregations`: []InlineResponse200158
    fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsApi.GetNetworkSwitchLinkAggregations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchLinkAggregationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200158**](InlineResponse200158.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchLinkAggregation

> InlineResponse200158 UpdateNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregation(updateNetworkSwitchLinkAggregation).Execute()

Update a link aggregation group



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
    linkAggregationId := "linkAggregationId_example" // string | Link aggregation ID
    updateNetworkSwitchLinkAggregation := *openapiclient.NewInlineObject146() // InlineObject146 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkAggregationsApi.UpdateNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregation(updateNetworkSwitchLinkAggregation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.UpdateNetworkSwitchLinkAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchLinkAggregation`: InlineResponse200158
    fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsApi.UpdateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**linkAggregationId** | **string** | Link aggregation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchLinkAggregation** | [**InlineObject146**](InlineObject146.md) |  | 

### Return type

[**InlineResponse200158**](InlineResponse200158.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

