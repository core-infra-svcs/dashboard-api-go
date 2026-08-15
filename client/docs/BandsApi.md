# \BandsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceCellularUplinksBandsMasksUpdate**](BandsApi.md#CreateDeviceCellularUplinksBandsMasksUpdate) | **Post** /devices/{serial}/cellular/uplinks/bands/masks/update | Update the cellular band masks for a device
[**GetOrganizationDevicesCellularUplinksBandsByDevice**](BandsApi.md#GetOrganizationDevicesCellularUplinksBandsByDevice) | **Get** /organizations/{organizationId}/devices/cellular/uplinks/bands/byDevice | List the latest cellular uplink signal information for devices in an organization



## CreateDeviceCellularUplinksBandsMasksUpdate

> InlineResponse20028 CreateDeviceCellularUplinksBandsMasksUpdate(ctx, serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()

Update the cellular band masks for a device



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
    createDeviceCellularUplinksBandsMasksUpdate := *openapiclient.NewInlineObject16("Slot_example", "Type_example", []string{"Masked_example"}) // InlineObject16 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BandsApi.CreateDeviceCellularUplinksBandsMasksUpdate(context.Background(), serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BandsApi.CreateDeviceCellularUplinksBandsMasksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceCellularUplinksBandsMasksUpdate`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `BandsApi.CreateDeviceCellularUplinksBandsMasksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceCellularUplinksBandsMasksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceCellularUplinksBandsMasksUpdate** | [**InlineObject16**](InlineObject16.md) |  | 

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularUplinksBandsByDevice

> InlineResponse200352 GetOrganizationDevicesCellularUplinksBandsByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the latest cellular uplink signal information for devices in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter the results by device serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BandsApi.GetOrganizationDevicesCellularUplinksBandsByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BandsApi.GetOrganizationDevicesCellularUplinksBandsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularUplinksBandsByDevice`: InlineResponse200352
    fmt.Fprintf(os.Stdout, "Response from `BandsApi.GetOrganizationDevicesCellularUplinksBandsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularUplinksBandsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter the results by device serials. Maximum 1000 serials. | 
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

