# \FieldsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNetworkSmDevicesFields**](FieldsApi.md#UpdateNetworkSmDevicesFields) | **Put** /networks/{networkId}/sm/devices/fields | Modify the fields of a device



## UpdateNetworkSmDevicesFields

> []InlineResponse200118 UpdateNetworkSmDevicesFields(ctx, networkId).UpdateNetworkSmDevicesFields(updateNetworkSmDevicesFields).Execute()

Modify the fields of a device



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
    updateNetworkSmDevicesFields := *openapiclient.NewInlineObject120(*openapiclient.NewNetworksNetworkIdSmDevicesFieldsDeviceFields()) // InlineObject120 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FieldsApi.UpdateNetworkSmDevicesFields(context.Background(), networkId).UpdateNetworkSmDevicesFields(updateNetworkSmDevicesFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldsApi.UpdateNetworkSmDevicesFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSmDevicesFields`: []InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `FieldsApi.UpdateNetworkSmDevicesFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmDevicesFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSmDevicesFields** | [**InlineObject120**](InlineObject120.md) |  | 

### Return type

[**[]InlineResponse200118**](InlineResponse200118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

