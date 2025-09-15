# \LocalOverridesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice**](LocalOverridesApi.md#GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice) | **Get** /organizations/{organizationId}/campusGateway/devices/uplinks/localOverrides/byDevice | Uplink overrides configured locally on Campus Gateway devices in an organization.



## GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice

> InlineResponse200257 GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Uplink overrides configured locally on Campus Gateway devices in an organization.



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
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocalOverridesApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocalOverridesApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice`: InlineResponse200257
    fmt.Fprintf(os.Stdout, "Response from `LocalOverridesApi.GetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayDevicesUplinksLocalOverridesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200257**](InlineResponse200257.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

