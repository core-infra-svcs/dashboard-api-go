# \AccountApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectNetworkApplianceUmbrellaAccount**](AccountApi.md#ConnectNetworkApplianceUmbrellaAccount) | **Post** /networks/{networkId}/appliance/umbrella/account/connect | Connect a Cisco Umbrella account to this network
[**DisconnectNetworkApplianceUmbrellaAccount**](AccountApi.md#DisconnectNetworkApplianceUmbrellaAccount) | **Post** /networks/{networkId}/appliance/umbrella/account/disconnect | Disconnect Umbrella account from this network



## ConnectNetworkApplianceUmbrellaAccount

> InlineResponse20081 ConnectNetworkApplianceUmbrellaAccount(ctx, networkId).ConnectNetworkApplianceUmbrellaAccount(connectNetworkApplianceUmbrellaAccount).Execute()

Connect a Cisco Umbrella account to this network



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
    connectNetworkApplianceUmbrellaAccount := *openapiclient.NewInlineObject75(*openapiclient.NewNetworksNetworkIdApplianceUmbrellaAccountConnectApi("Key_example", "Secret_example")) // InlineObject75 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.ConnectNetworkApplianceUmbrellaAccount(context.Background(), networkId).ConnectNetworkApplianceUmbrellaAccount(connectNetworkApplianceUmbrellaAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ConnectNetworkApplianceUmbrellaAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectNetworkApplianceUmbrellaAccount`: InlineResponse20081
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ConnectNetworkApplianceUmbrellaAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectNetworkApplianceUmbrellaAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectNetworkApplianceUmbrellaAccount** | [**InlineObject75**](InlineObject75.md) |  | 

### Return type

[**InlineResponse20081**](InlineResponse20081.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectNetworkApplianceUmbrellaAccount

> DisconnectNetworkApplianceUmbrellaAccount(ctx, networkId).Execute()

Disconnect Umbrella account from this network



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
    resp, r, err := apiClient.AccountApi.DisconnectNetworkApplianceUmbrellaAccount(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.DisconnectNetworkApplianceUmbrellaAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectNetworkApplianceUmbrellaAccountRequest struct via the builder pattern


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

