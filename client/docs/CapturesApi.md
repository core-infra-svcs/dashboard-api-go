# \CapturesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkOrganizationDevicesPacketCaptureCapturesCreate**](CapturesApi.md#BulkOrganizationDevicesPacketCaptureCapturesCreate) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/bulkCreate | Perform a packet capture on multiple devices and store in Meraki Cloud.
[**BulkOrganizationDevicesPacketCaptureCapturesDelete**](CapturesApi.md#BulkOrganizationDevicesPacketCaptureCapturesDelete) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/bulkDelete | BulkDelete packet captures from cloud
[**CreateOrganizationDevicesPacketCaptureCapture**](CapturesApi.md#CreateOrganizationDevicesPacketCaptureCapture) | **Post** /organizations/{organizationId}/devices/packetCapture/captures | Perform a packet capture on a device and store in Meraki Cloud
[**DeleteOrganizationDevicesPacketCaptureCapture**](CapturesApi.md#DeleteOrganizationDevicesPacketCaptureCapture) | **Delete** /organizations/{organizationId}/devices/packetCapture/captures/{captureId} | Delete a single packet capture from cloud using captureId
[**GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl**](CapturesApi.md#GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/{captureId}/downloadUrl/generate | Get presigned download URL for given packet capture id
[**GetOrganizationDevicesPacketCaptureCaptures**](CapturesApi.md#GetOrganizationDevicesPacketCaptureCaptures) | **Get** /organizations/{organizationId}/devices/packetCapture/captures | List Packet Captures
[**StopOrganizationDevicesPacketCaptureCapture**](CapturesApi.md#StopOrganizationDevicesPacketCaptureCapture) | **Post** /organizations/{organizationId}/devices/packetCapture/captures/{captureId}/stop | Stop a specific packet capture (not supported for Catalyst devices)



## BulkOrganizationDevicesPacketCaptureCapturesCreate

> InlineResponse20118 BulkOrganizationDevicesPacketCaptureCapturesCreate(ctx, organizationId).BulkOrganizationDevicesPacketCaptureCapturesCreate(bulkOrganizationDevicesPacketCaptureCapturesCreate).Execute()

Perform a packet capture on multiple devices and store in Meraki Cloud.



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
    bulkOrganizationDevicesPacketCaptureCapturesCreate := *openapiclient.NewInlineObject257([]openapiclient.OrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices{*openapiclient.NewOrganizationsOrganizationIdDevicesPacketCaptureCapturesBulkCreateDevices()}, "Name_example") // InlineObject257 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.BulkOrganizationDevicesPacketCaptureCapturesCreate(context.Background(), organizationId).BulkOrganizationDevicesPacketCaptureCapturesCreate(bulkOrganizationDevicesPacketCaptureCapturesCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.BulkOrganizationDevicesPacketCaptureCapturesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkOrganizationDevicesPacketCaptureCapturesCreate`: InlineResponse20118
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.BulkOrganizationDevicesPacketCaptureCapturesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationDevicesPacketCaptureCapturesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationDevicesPacketCaptureCapturesCreate** | [**InlineObject257**](InlineObject257.md) |  | 

### Return type

[**InlineResponse20118**](InlineResponse20118.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkOrganizationDevicesPacketCaptureCapturesDelete

> BulkOrganizationDevicesPacketCaptureCapturesDelete(ctx, organizationId).BulkOrganizationDevicesPacketCaptureCapturesDelete(bulkOrganizationDevicesPacketCaptureCapturesDelete).Execute()

BulkDelete packet captures from cloud



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
    bulkOrganizationDevicesPacketCaptureCapturesDelete := *openapiclient.NewInlineObject258([]string{"CaptureIds_example"}) // InlineObject258 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.BulkOrganizationDevicesPacketCaptureCapturesDelete(context.Background(), organizationId).BulkOrganizationDevicesPacketCaptureCapturesDelete(bulkOrganizationDevicesPacketCaptureCapturesDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.BulkOrganizationDevicesPacketCaptureCapturesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationDevicesPacketCaptureCapturesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationDevicesPacketCaptureCapturesDelete** | [**InlineObject258**](InlineObject258.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationDevicesPacketCaptureCapture

> InlineResponse200274Items CreateOrganizationDevicesPacketCaptureCapture(ctx, organizationId).CreateOrganizationDevicesPacketCaptureCapture(createOrganizationDevicesPacketCaptureCapture).Execute()

Perform a packet capture on a device and store in Meraki Cloud



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
    createOrganizationDevicesPacketCaptureCapture := *openapiclient.NewInlineObject256([]string{"Serials_example"}, "Name_example") // InlineObject256 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.CreateOrganizationDevicesPacketCaptureCapture(context.Background(), organizationId).CreateOrganizationDevicesPacketCaptureCapture(createOrganizationDevicesPacketCaptureCapture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.CreateOrganizationDevicesPacketCaptureCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDevicesPacketCaptureCapture`: InlineResponse200274Items
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.CreateOrganizationDevicesPacketCaptureCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDevicesPacketCaptureCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationDevicesPacketCaptureCapture** | [**InlineObject256**](InlineObject256.md) |  | 

### Return type

[**InlineResponse200274Items**](InlineResponse200274Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationDevicesPacketCaptureCapture

> DeleteOrganizationDevicesPacketCaptureCapture(ctx, organizationId, captureId).Execute()

Delete a single packet capture from cloud using captureId



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
    captureId := "captureId_example" // string | Capture ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.DeleteOrganizationDevicesPacketCaptureCapture(context.Background(), organizationId, captureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.DeleteOrganizationDevicesPacketCaptureCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**captureId** | **string** | Capture ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationDevicesPacketCaptureCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl

> InlineResponse200275 GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl(ctx, organizationId, captureId).Execute()

Get presigned download URL for given packet capture id



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
    captureId := "captureId_example" // string | Capture ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl(context.Background(), organizationId, captureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl`: InlineResponse200275
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.GenerateOrganizationDevicesPacketCaptureCaptureDownloadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**captureId** | **string** | Capture ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOrganizationDevicesPacketCaptureCaptureDownloadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200275**](InlineResponse200275.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesPacketCaptureCaptures

> InlineResponse200274 GetOrganizationDevicesPacketCaptureCaptures(ctx, organizationId).CaptureIds(captureIds).NetworkIds(networkIds).Serials(serials).Process(process).CaptureStatus(captureStatus).Name(name).ClientMac(clientMac).Notes(notes).DeviceName(deviceName).AdminName(adminName).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List Packet Captures



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
    captureIds := []string{"Inner_example"} // []string | Return the packet captures of the specified capture ids (optional)
    networkIds := []string{"Inner_example"} // []string | Return the packet captures of the specified network(s) (optional)
    serials := []string{"Inner_example"} // []string | Return the packet captures of the specified device(s) (optional)
    process := []string{"Inner_example"} // []string | Return the packet captures of the specified process (optional)
    captureStatus := []string{"Inner_example"} // []string | Return the packet captures of the specified capture status (optional)
    name := []string{"Inner_example"} // []string | Return the packet captures matching the specified name (optional)
    clientMac := []string{"Inner_example"} // []string | Return the packet captures matching the specified client macs (optional)
    notes := "notes_example" // string | Return the packet captures matching the specified notes (optional)
    deviceName := "deviceName_example" // string | Return the packet captures matching the specified device name (optional)
    adminName := "adminName_example" // string | Return the packet captures matching the admin name (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.GetOrganizationDevicesPacketCaptureCaptures(context.Background(), organizationId).CaptureIds(captureIds).NetworkIds(networkIds).Serials(serials).Process(process).CaptureStatus(captureStatus).Name(name).ClientMac(clientMac).Notes(notes).DeviceName(deviceName).AdminName(adminName).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.GetOrganizationDevicesPacketCaptureCaptures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesPacketCaptureCaptures`: InlineResponse200274
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.GetOrganizationDevicesPacketCaptureCaptures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesPacketCaptureCapturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captureIds** | **[]string** | Return the packet captures of the specified capture ids | 
 **networkIds** | **[]string** | Return the packet captures of the specified network(s) | 
 **serials** | **[]string** | Return the packet captures of the specified device(s) | 
 **process** | **[]string** | Return the packet captures of the specified process | 
 **captureStatus** | **[]string** | Return the packet captures of the specified capture status | 
 **name** | **[]string** | Return the packet captures matching the specified name | 
 **clientMac** | **[]string** | Return the packet captures matching the specified client macs | 
 **notes** | **string** | Return the packet captures matching the specified notes | 
 **deviceName** | **string** | Return the packet captures matching the specified device name | 
 **adminName** | **string** | Return the packet captures matching the admin name | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;descending&#39;. | 

### Return type

[**InlineResponse200274**](InlineResponse200274.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopOrganizationDevicesPacketCaptureCapture

> InlineResponse200274Items StopOrganizationDevicesPacketCaptureCapture(ctx, organizationId, captureId).StopOrganizationDevicesPacketCaptureCapture(stopOrganizationDevicesPacketCaptureCapture).Execute()

Stop a specific packet capture (not supported for Catalyst devices)



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
    captureId := "captureId_example" // string | Capture ID
    stopOrganizationDevicesPacketCaptureCapture := *openapiclient.NewInlineObject259([]string{"Serials_example"}) // InlineObject259 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapturesApi.StopOrganizationDevicesPacketCaptureCapture(context.Background(), organizationId, captureId).StopOrganizationDevicesPacketCaptureCapture(stopOrganizationDevicesPacketCaptureCapture).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapturesApi.StopOrganizationDevicesPacketCaptureCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopOrganizationDevicesPacketCaptureCapture`: InlineResponse200274Items
    fmt.Fprintf(os.Stdout, "Response from `CapturesApi.StopOrganizationDevicesPacketCaptureCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**captureId** | **string** | Capture ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopOrganizationDevicesPacketCaptureCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stopOrganizationDevicesPacketCaptureCapture** | [**InlineObject259**](InlineObject259.md) |  | 

### Return type

[**InlineResponse200274Items**](InlineResponse200274Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

