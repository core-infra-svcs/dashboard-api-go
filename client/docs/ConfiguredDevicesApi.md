# \ConfiguredDevicesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessElectronicShelfLabelConfiguredDevices**](ConfiguredDevicesApi.md#GetNetworkWirelessElectronicShelfLabelConfiguredDevices) | **Get** /networks/{networkId}/wireless/electronicShelfLabel/configuredDevices | Get a list of all ESL eligible devices of a network



## GetNetworkWirelessElectronicShelfLabelConfiguredDevices

> []InlineResponse200198 GetNetworkWirelessElectronicShelfLabelConfiguredDevices(ctx, networkId).Execute()

Get a list of all ESL eligible devices of a network



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
    resp, r, err := apiClient.ConfiguredDevicesApi.GetNetworkWirelessElectronicShelfLabelConfiguredDevices(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguredDevicesApi.GetNetworkWirelessElectronicShelfLabelConfiguredDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessElectronicShelfLabelConfiguredDevices`: []InlineResponse200198
    fmt.Fprintf(os.Stdout, "Response from `ConfiguredDevicesApi.GetNetworkWirelessElectronicShelfLabelConfiguredDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessElectronicShelfLabelConfiguredDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200198**](InlineResponse200198.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

