# \ConnectivityMonitoringDestinationsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceConnectivityMonitoringDestinations**](ConnectivityMonitoringDestinationsApi.md#GetNetworkApplianceConnectivityMonitoringDestinations) | **Get** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Return the connectivity testing destinations for an MX network
[**GetNetworkCellularGatewayConnectivityMonitoringDestinations**](ConnectivityMonitoringDestinationsApi.md#GetNetworkCellularGatewayConnectivityMonitoringDestinations) | **Get** /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations | Return the connectivity testing destinations for an MG network
[**UpdateNetworkApplianceConnectivityMonitoringDestinations**](ConnectivityMonitoringDestinationsApi.md#UpdateNetworkApplianceConnectivityMonitoringDestinations) | **Put** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Update the connectivity testing destinations for an MX network
[**UpdateNetworkCellularGatewayConnectivityMonitoringDestinations**](ConnectivityMonitoringDestinationsApi.md#UpdateNetworkCellularGatewayConnectivityMonitoringDestinations) | **Put** /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations | Update the connectivity testing destinations for an MG network



## GetNetworkApplianceConnectivityMonitoringDestinations

> InlineResponse20051 GetNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).Execute()

Return the connectivity testing destinations for an MX network



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
    resp, r, err := apiClient.ConnectivityMonitoringDestinationsApi.GetNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityMonitoringDestinationsApi.GetNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceConnectivityMonitoringDestinations`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityMonitoringDestinationsApi.GetNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewayConnectivityMonitoringDestinations

> InlineResponse20051 GetNetworkCellularGatewayConnectivityMonitoringDestinations(ctx, networkId).Execute()

Return the connectivity testing destinations for an MG network



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
    resp, r, err := apiClient.ConnectivityMonitoringDestinationsApi.GetNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityMonitoringDestinationsApi.GetNetworkCellularGatewayConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCellularGatewayConnectivityMonitoringDestinations`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityMonitoringDestinationsApi.GetNetworkCellularGatewayConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceConnectivityMonitoringDestinations

> InlineResponse20051 UpdateNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).UpdateNetworkApplianceConnectivityMonitoringDestinations(updateNetworkApplianceConnectivityMonitoringDestinations).Execute()

Update the connectivity testing destinations for an MX network



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
    updateNetworkApplianceConnectivityMonitoringDestinations := *openapiclient.NewInlineObject41() // InlineObject41 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectivityMonitoringDestinationsApi.UpdateNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).UpdateNetworkApplianceConnectivityMonitoringDestinations(updateNetworkApplianceConnectivityMonitoringDestinations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityMonitoringDestinationsApi.UpdateNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceConnectivityMonitoringDestinations`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityMonitoringDestinationsApi.UpdateNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceConnectivityMonitoringDestinations** | [**InlineObject41**](InlineObject41.md) |  | 

### Return type

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewayConnectivityMonitoringDestinations

> InlineResponse20051 UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(ctx, networkId).UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(updateNetworkCellularGatewayConnectivityMonitoringDestinations).Execute()

Update the connectivity testing destinations for an MG network



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
    updateNetworkCellularGatewayConnectivityMonitoringDestinations := *openapiclient.NewInlineObject87() // InlineObject87 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectivityMonitoringDestinationsApi.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(updateNetworkCellularGatewayConnectivityMonitoringDestinations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectivityMonitoringDestinationsApi.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCellularGatewayConnectivityMonitoringDestinations`: InlineResponse20051
    fmt.Fprintf(os.Stdout, "Response from `ConnectivityMonitoringDestinationsApi.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayConnectivityMonitoringDestinations** | [**InlineObject87**](InlineObject87.md) |  | 

### Return type

[**InlineResponse20051**](InlineResponse20051.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

