# \FieldsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNetworkSmDevicesFields**](FieldsApi.md#UpdateNetworkSmDevicesFields) | **Put** /networks/{networkId}/sm/devices/fields | Modify the fields of a device



## UpdateNetworkSmDevicesFields

> map[string]interface{} UpdateNetworkSmDevicesFields(ctx, networkId).UpdateNetworkSmDevicesFields(updateNetworkSmDevicesFields).Execute()

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
    networkId := "networkId_example" // string | 
    updateNetworkSmDevicesFields := *openapiclient.NewInlineObject83(*openapiclient.NewNetworksNetworkIdSmDevicesFieldsDeviceFields()) // InlineObject83 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FieldsApi.UpdateNetworkSmDevicesFields(context.Background(), networkId).UpdateNetworkSmDevicesFields(updateNetworkSmDevicesFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldsApi.UpdateNetworkSmDevicesFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSmDevicesFields`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FieldsApi.UpdateNetworkSmDevicesFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmDevicesFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSmDevicesFields** | [**InlineObject83**](InlineObject83.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

