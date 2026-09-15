# \OverridesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceWirelessRadioOverrides**](OverridesApi.md#GetDeviceWirelessRadioOverrides) | **Get** /devices/{serial}/wireless/radio/overrides | Return the radio overrides of a device
[**GetOrganizationWirelessRadioOverridesByDevice**](OverridesApi.md#GetOrganizationWirelessRadioOverridesByDevice) | **Get** /organizations/{organizationId}/wireless/radio/overrides/byDevice | Return a list of radio overrides
[**UpdateDeviceWirelessRadioOverrides**](OverridesApi.md#UpdateDeviceWirelessRadioOverrides) | **Put** /devices/{serial}/wireless/radio/overrides | Update 2.4 GHz, 5 GHz, and 6 GHz radio settings (channel, channel width, power, and enable/disable) that override RF profiles



## GetDeviceWirelessRadioOverrides

> InlineResponse20064 GetDeviceWirelessRadioOverrides(ctx, serial).Execute()

Return the radio overrides of a device



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.GetDeviceWirelessRadioOverrides(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.GetDeviceWirelessRadioOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessRadioOverrides`: InlineResponse20064
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.GetDeviceWirelessRadioOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessRadioOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20064**](InlineResponse20064.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessRadioOverridesByDevice

> InlineResponse200488 GetOrganizationWirelessRadioOverridesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()

Return a list of radio overrides



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned radio overrides will be filtered to only include these networks. (optional)
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned radio overrides will be filtered to only include these serials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.GetOrganizationWirelessRadioOverridesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.GetOrganizationWirelessRadioOverridesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessRadioOverridesByDevice`: InlineResponse200488
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.GetOrganizationWirelessRadioOverridesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessRadioOverridesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned radio overrides will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned radio overrides will be filtered to only include these serials. | 

### Return type

[**InlineResponse200488**](InlineResponse200488.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessRadioOverrides

> InlineResponse20064 UpdateDeviceWirelessRadioOverrides(ctx, serial).UpdateDeviceWirelessRadioOverrides(updateDeviceWirelessRadioOverrides).Execute()

Update 2.4 GHz, 5 GHz, and 6 GHz radio settings (channel, channel width, power, and enable/disable) that override RF profiles



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
    serial := "serial_example" // string | Serial
    updateDeviceWirelessRadioOverrides := *openapiclient.NewInlineObject47() // InlineObject47 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverridesApi.UpdateDeviceWirelessRadioOverrides(context.Background(), serial).UpdateDeviceWirelessRadioOverrides(updateDeviceWirelessRadioOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverridesApi.UpdateDeviceWirelessRadioOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessRadioOverrides`: InlineResponse20064
    fmt.Fprintf(os.Stdout, "Response from `OverridesApi.UpdateDeviceWirelessRadioOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessRadioOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessRadioOverrides** | [**InlineObject47**](InlineObject47.md) |  | 

### Return type

[**InlineResponse20064**](InlineResponse20064.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

