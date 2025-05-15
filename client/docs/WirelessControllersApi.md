# \WirelessControllersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWirelessDevicesWirelessControllersByDevice**](WirelessControllersApi.md#GetOrganizationWirelessDevicesWirelessControllersByDevice) | **Get** /organizations/{organizationId}/wireless/devices/wirelessControllers/byDevice | List of Catalyst access points information



## GetOrganizationWirelessDevicesWirelessControllersByDevice

> InlineResponse200352 GetOrganizationWirelessDevicesWirelessControllersByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).ControllerSerials(controllerSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List of Catalyst access points information



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
    organizationId := "organizationId_example" // string | Organization ID
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter access points by network ID. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter access points by its cloud ID. This filter uses multiple exact matches. (optional)
    controllerSerials := []string{"Inner_example"} // []string | Optional parameter to filter access points by its wireless LAN controller cloud ID. This filter uses multiple exact matches. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessControllersApi.GetOrganizationWirelessDevicesWirelessControllersByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).ControllerSerials(controllerSerials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessControllersApi.GetOrganizationWirelessDevicesWirelessControllersByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesWirelessControllersByDevice`: InlineResponse200352
    fmt.Fprintf(os.Stdout, "Response from `WirelessControllersApi.GetOrganizationWirelessDevicesWirelessControllersByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesWirelessControllersByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter access points by network ID. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter access points by its cloud ID. This filter uses multiple exact matches. | 
 **controllerSerials** | **[]string** | Optional parameter to filter access points by its wireless LAN controller cloud ID. This filter uses multiple exact matches. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200352**](InlineResponse200352.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

