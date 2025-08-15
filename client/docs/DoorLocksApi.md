# \DoorLocksApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWirelessZigbeeDoorLocks**](DoorLocksApi.md#GetOrganizationWirelessZigbeeDoorLocks) | **Get** /organizations/{organizationId}/wireless/zigbee/doorLocks | Return the list of door locks for a network
[**UpdateOrganizationWirelessZigbeeDoorLock**](DoorLocksApi.md#UpdateOrganizationWirelessZigbeeDoorLock) | **Put** /organizations/{organizationId}/wireless/zigbee/doorLocks/{doorLockId} | Endpoint to batch update door locks params



## GetOrganizationWirelessZigbeeDoorLocks

> []InlineResponse20048DoorLocks GetOrganizationWirelessZigbeeDoorLocks(ctx, organizationId).NetworkIds(networkIds).Serial(serial).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the list of door locks for a network



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
    networkIds := []string{"Inner_example"} // []string | Filter by specified Network IDs (optional)
    serial := "serial_example" // string | Filter by device serial (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DoorLocksApi.GetOrganizationWirelessZigbeeDoorLocks(context.Background(), organizationId).NetworkIds(networkIds).Serial(serial).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DoorLocksApi.GetOrganizationWirelessZigbeeDoorLocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessZigbeeDoorLocks`: []InlineResponse20048DoorLocks
    fmt.Fprintf(os.Stdout, "Response from `DoorLocksApi.GetOrganizationWirelessZigbeeDoorLocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessZigbeeDoorLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter by specified Network IDs | 
 **serial** | **string** | Filter by device serial | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse20048DoorLocks**](InlineResponse20048DoorLocks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessZigbeeDoorLock

> InlineResponse20048DoorLocks UpdateOrganizationWirelessZigbeeDoorLock(ctx, organizationId, doorLockId).UpdateOrganizationWirelessZigbeeDoorLock(updateOrganizationWirelessZigbeeDoorLock).Execute()

Endpoint to batch update door locks params



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
    doorLockId := "doorLockId_example" // string | Door lock ID
    updateOrganizationWirelessZigbeeDoorLock := *openapiclient.NewInlineObject310() // InlineObject310 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DoorLocksApi.UpdateOrganizationWirelessZigbeeDoorLock(context.Background(), organizationId, doorLockId).UpdateOrganizationWirelessZigbeeDoorLock(updateOrganizationWirelessZigbeeDoorLock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DoorLocksApi.UpdateOrganizationWirelessZigbeeDoorLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessZigbeeDoorLock`: InlineResponse20048DoorLocks
    fmt.Fprintf(os.Stdout, "Response from `DoorLocksApi.UpdateOrganizationWirelessZigbeeDoorLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**doorLockId** | **string** | Door lock ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessZigbeeDoorLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationWirelessZigbeeDoorLock** | [**InlineObject310**](InlineObject310.md) |  | 

### Return type

[**InlineResponse20048DoorLocks**](InlineResponse20048DoorLocks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

