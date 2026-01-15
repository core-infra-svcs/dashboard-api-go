# \SdwanApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNetworkApplianceSdwanInternetPolicies**](SdwanApi.md#UpdateNetworkApplianceSdwanInternetPolicies) | **Put** /networks/{networkId}/appliance/sdwan/internetPolicies | Update SDWAN internet traffic preferences for an MX network



## UpdateNetworkApplianceSdwanInternetPolicies

> InlineResponse20065 UpdateNetworkApplianceSdwanInternetPolicies(ctx, networkId).UpdateNetworkApplianceSdwanInternetPolicies(updateNetworkApplianceSdwanInternetPolicies).Execute()

Update SDWAN internet traffic preferences for an MX network



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
    updateNetworkApplianceSdwanInternetPolicies := *openapiclient.NewInlineObject60() // InlineObject60 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SdwanApi.UpdateNetworkApplianceSdwanInternetPolicies(context.Background(), networkId).UpdateNetworkApplianceSdwanInternetPolicies(updateNetworkApplianceSdwanInternetPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SdwanApi.UpdateNetworkApplianceSdwanInternetPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSdwanInternetPolicies`: InlineResponse20065
    fmt.Fprintf(os.Stdout, "Response from `SdwanApi.UpdateNetworkApplianceSdwanInternetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSdwanInternetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSdwanInternetPolicies** | [**InlineObject60**](InlineObject60.md) |  | 

### Return type

[**InlineResponse20065**](InlineResponse20065.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

