# \StormControlApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchStormControl**](StormControlApi.md#GetNetworkSwitchStormControl) | **Get** /networks/{networkId}/switch/stormControl | Return the storm control configuration for a switch network
[**UpdateNetworkSwitchStormControl**](StormControlApi.md#UpdateNetworkSwitchStormControl) | **Put** /networks/{networkId}/switch/stormControl | Update the storm control configuration for a switch network



## GetNetworkSwitchStormControl

> InlineResponse200169 GetNetworkSwitchStormControl(ctx, networkId).Execute()

Return the storm control configuration for a switch network



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
    resp, r, err := apiClient.StormControlApi.GetNetworkSwitchStormControl(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StormControlApi.GetNetworkSwitchStormControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStormControl`: InlineResponse200169
    fmt.Fprintf(os.Stdout, "Response from `StormControlApi.GetNetworkSwitchStormControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStormControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200169**](InlineResponse200169.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStormControl

> InlineResponse200169 UpdateNetworkSwitchStormControl(ctx, networkId).UpdateNetworkSwitchStormControl(updateNetworkSwitchStormControl).Execute()

Update the storm control configuration for a switch network



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
    updateNetworkSwitchStormControl := *openapiclient.NewInlineObject164() // InlineObject164 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StormControlApi.UpdateNetworkSwitchStormControl(context.Background(), networkId).UpdateNetworkSwitchStormControl(updateNetworkSwitchStormControl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StormControlApi.UpdateNetworkSwitchStormControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStormControl`: InlineResponse200169
    fmt.Fprintf(os.Stdout, "Response from `StormControlApi.UpdateNetworkSwitchStormControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStormControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchStormControl** | [**InlineObject164**](InlineObject164.md) |  | 

### Return type

[**InlineResponse200169**](InlineResponse200169.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

