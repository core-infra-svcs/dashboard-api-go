# \PolicyApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkClientPolicy**](PolicyApi.md#GetNetworkClientPolicy) | **Get** /networks/{networkId}/clients/{clientId}/policy | Return the policy assigned to a client on the network
[**UpdateNetworkClientPolicy**](PolicyApi.md#UpdateNetworkClientPolicy) | **Put** /networks/{networkId}/clients/{clientId}/policy | Update the policy assigned to a client on the network



## GetNetworkClientPolicy

> InlineResponse20089 GetNetworkClientPolicy(ctx, networkId, clientId).Execute()

Return the policy assigned to a client on the network



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
    resp, r, err := apiClient.PolicyApi.GetNetworkClientPolicy(context.Background(), networkId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetNetworkClientPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientPolicy`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkClientPolicy

> InlineResponse20089 UpdateNetworkClientPolicy(ctx, networkId, clientId).UpdateNetworkClientPolicy(updateNetworkClientPolicy).Execute()

Update the policy assigned to a client on the network



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
    updateNetworkClientPolicy := *openapiclient.NewInlineObject92("DevicePolicy_example") // InlineObject92 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.UpdateNetworkClientPolicy(context.Background(), networkId, clientId).UpdateNetworkClientPolicy(updateNetworkClientPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.UpdateNetworkClientPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkClientPolicy`: InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.UpdateNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientPolicy** | [**InlineObject92**](InlineObject92.md) |  | 

### Return type

[**InlineResponse20089**](InlineResponse20089.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

