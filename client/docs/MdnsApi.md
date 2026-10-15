# \MdnsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNetworkCampusGatewaySsidMdns**](MdnsApi.md#UpdateNetworkCampusGatewaySsidMdns) | **Put** /networks/{networkId}/campusGateway/ssids/{number}/mdns | Update the mDNS gateway settings and rules for a SSID and cluster



## UpdateNetworkCampusGatewaySsidMdns

> InlineResponse200118 UpdateNetworkCampusGatewaySsidMdns(ctx, networkId, number).UpdateNetworkCampusGatewaySsidMdns(updateNetworkCampusGatewaySsidMdns).Execute()

Update the mDNS gateway settings and rules for a SSID and cluster



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
    number := "number_example" // string | Number
    updateNetworkCampusGatewaySsidMdns := *openapiclient.NewInlineObject107() // InlineObject107 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MdnsApi.UpdateNetworkCampusGatewaySsidMdns(context.Background(), networkId, number).UpdateNetworkCampusGatewaySsidMdns(updateNetworkCampusGatewaySsidMdns).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MdnsApi.UpdateNetworkCampusGatewaySsidMdns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCampusGatewaySsidMdns`: InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `MdnsApi.UpdateNetworkCampusGatewaySsidMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCampusGatewaySsidMdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCampusGatewaySsidMdns** | [**InlineObject107**](InlineObject107.md) |  | 

### Return type

[**InlineResponse200118**](InlineResponse200118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

