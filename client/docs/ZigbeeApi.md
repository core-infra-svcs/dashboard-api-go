# \ZigbeeApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceWirelessZigbeeEnrollment**](ZigbeeApi.md#CreateDeviceWirelessZigbeeEnrollment) | **Post** /devices/{serial}/wireless/zigbee/enrollments | Enqueue a job to start enrolling door locks on zigbee configured wireless devices
[**CreateOrganizationWirelessZigbeeDisenrollment**](ZigbeeApi.md#CreateOrganizationWirelessZigbeeDisenrollment) | **Post** /organizations/{organizationId}/wireless/zigbee/disenrollments | Enqueue a job to start disenrolling door locks on zigbee configured wireless devices
[**GetDeviceWirelessZigbeeEnrollment**](ZigbeeApi.md#GetDeviceWirelessZigbeeEnrollment) | **Get** /devices/{serial}/wireless/zigbee/enrollments/{enrollmentId} | Return an enrollment
[**GetOrganizationWirelessZigbeeByNetwork**](ZigbeeApi.md#GetOrganizationWirelessZigbeeByNetwork) | **Get** /organizations/{organizationId}/wireless/zigbee/byNetwork | Return list of Zigbee configs
[**GetOrganizationWirelessZigbeeDevices**](ZigbeeApi.md#GetOrganizationWirelessZigbeeDevices) | **Get** /organizations/{organizationId}/wireless/zigbee/devices | List the Zigbee wireless devices for an organization or the supplied network(s)
[**GetOrganizationWirelessZigbeeDisenrollment**](ZigbeeApi.md#GetOrganizationWirelessZigbeeDisenrollment) | **Get** /organizations/{organizationId}/wireless/zigbee/disenrollments/{disenrollmentId} | Return a disenrollment
[**GetOrganizationWirelessZigbeeDoorLocks**](ZigbeeApi.md#GetOrganizationWirelessZigbeeDoorLocks) | **Get** /organizations/{organizationId}/wireless/zigbee/doorLocks | Return the list of door locks for a network
[**UpdateNetworkWirelessZigbee**](ZigbeeApi.md#UpdateNetworkWirelessZigbee) | **Put** /networks/{networkId}/wireless/zigbee | Update Zigbee Configs for specified network
[**UpdateOrganizationWirelessZigbeeDevice**](ZigbeeApi.md#UpdateOrganizationWirelessZigbeeDevice) | **Put** /organizations/{organizationId}/wireless/zigbee/devices/{id} | Endpoint to update zigbee gateways
[**UpdateOrganizationWirelessZigbeeDoorLock**](ZigbeeApi.md#UpdateOrganizationWirelessZigbeeDoorLock) | **Put** /organizations/{organizationId}/wireless/zigbee/doorLocks/{doorLockId} | Endpoint to batch update door locks params



## CreateDeviceWirelessZigbeeEnrollment

> InlineResponse20110 CreateDeviceWirelessZigbeeEnrollment(ctx, serial).Execute()

Enqueue a job to start enrolling door locks on zigbee configured wireless devices



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
    resp, r, err := apiClient.ZigbeeApi.CreateDeviceWirelessZigbeeEnrollment(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.CreateDeviceWirelessZigbeeEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceWirelessZigbeeEnrollment`: InlineResponse20110
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.CreateDeviceWirelessZigbeeEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceWirelessZigbeeEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20110**](InlineResponse20110.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationWirelessZigbeeDisenrollment

> InlineResponse20124 CreateOrganizationWirelessZigbeeDisenrollment(ctx, organizationId).CreateOrganizationWirelessZigbeeDisenrollment(createOrganizationWirelessZigbeeDisenrollment).Execute()

Enqueue a job to start disenrolling door locks on zigbee configured wireless devices



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
    createOrganizationWirelessZigbeeDisenrollment := *openapiclient.NewInlineObject312() // InlineObject312 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.CreateOrganizationWirelessZigbeeDisenrollment(context.Background(), organizationId).CreateOrganizationWirelessZigbeeDisenrollment(createOrganizationWirelessZigbeeDisenrollment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.CreateOrganizationWirelessZigbeeDisenrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessZigbeeDisenrollment`: InlineResponse20124
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.CreateOrganizationWirelessZigbeeDisenrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessZigbeeDisenrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationWirelessZigbeeDisenrollment** | [**InlineObject312**](InlineObject312.md) |  | 

### Return type

[**InlineResponse20124**](InlineResponse20124.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessZigbeeEnrollment

> InlineResponse20049 GetDeviceWirelessZigbeeEnrollment(ctx, serial, enrollmentId).Execute()

Return an enrollment



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
    enrollmentId := "enrollmentId_example" // string | Enrollment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.GetDeviceWirelessZigbeeEnrollment(context.Background(), serial, enrollmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.GetDeviceWirelessZigbeeEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessZigbeeEnrollment`: InlineResponse20049
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.GetDeviceWirelessZigbeeEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**enrollmentId** | **string** | Enrollment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessZigbeeEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20049**](InlineResponse20049.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessZigbeeByNetwork

> []InlineResponse20117 GetOrganizationWirelessZigbeeByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Return list of Zigbee configs



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Filter by specified Network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.GetOrganizationWirelessZigbeeByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.GetOrganizationWirelessZigbeeByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessZigbeeByNetwork`: []InlineResponse20117
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.GetOrganizationWirelessZigbeeByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessZigbeeByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Filter by specified Network IDs | 

### Return type

[**[]InlineResponse20117**](InlineResponse20117.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessZigbeeDevices

> []InlineResponse200379 GetOrganizationWirelessZigbeeDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).IsEnrolled(isEnrolled).Search(search).Execute()

List the Zigbee wireless devices for an organization or the supplied network(s)



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Parameter of networks you want the zigbee devices for. E.g.: networkIds[]=N_12345678&networkIds[]=N_3456 (optional)
    isEnrolled := true // bool | Filter devices based on if they are enrolled or not (optional)
    search := "search_example" // string | Filter devices by their name, tag or serial (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.GetOrganizationWirelessZigbeeDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).IsEnrolled(isEnrolled).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.GetOrganizationWirelessZigbeeDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessZigbeeDevices`: []InlineResponse200379
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.GetOrganizationWirelessZigbeeDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessZigbeeDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Parameter of networks you want the zigbee devices for. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;N_3456 | 
 **isEnrolled** | **bool** | Filter devices based on if they are enrolled or not | 
 **search** | **string** | Filter devices by their name, tag or serial | 

### Return type

[**[]InlineResponse200379**](InlineResponse200379.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessZigbeeDisenrollment

> InlineResponse200380 GetOrganizationWirelessZigbeeDisenrollment(ctx, organizationId, disenrollmentId).Execute()

Return a disenrollment



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
    disenrollmentId := "disenrollmentId_example" // string | Disenrollment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.GetOrganizationWirelessZigbeeDisenrollment(context.Background(), organizationId, disenrollmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.GetOrganizationWirelessZigbeeDisenrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessZigbeeDisenrollment`: InlineResponse200380
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.GetOrganizationWirelessZigbeeDisenrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**disenrollmentId** | **string** | Disenrollment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessZigbeeDisenrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200380**](InlineResponse200380.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessZigbeeDoorLocks

> []InlineResponse20049DoorLocks GetOrganizationWirelessZigbeeDoorLocks(ctx, organizationId).NetworkIds(networkIds).Serial(serial).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.ZigbeeApi.GetOrganizationWirelessZigbeeDoorLocks(context.Background(), organizationId).NetworkIds(networkIds).Serial(serial).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.GetOrganizationWirelessZigbeeDoorLocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessZigbeeDoorLocks`: []InlineResponse20049DoorLocks
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.GetOrganizationWirelessZigbeeDoorLocks`: %v\n", resp)
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

[**[]InlineResponse20049DoorLocks**](InlineResponse20049DoorLocks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessZigbee

> InlineResponse20117 UpdateNetworkWirelessZigbee(ctx, networkId).UpdateNetworkWirelessZigbee(updateNetworkWirelessZigbee).Execute()

Update Zigbee Configs for specified network



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
    updateNetworkWirelessZigbee := *openapiclient.NewInlineObject209() // InlineObject209 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.UpdateNetworkWirelessZigbee(context.Background(), networkId).UpdateNetworkWirelessZigbee(updateNetworkWirelessZigbee).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.UpdateNetworkWirelessZigbee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessZigbee`: InlineResponse20117
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.UpdateNetworkWirelessZigbee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessZigbeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessZigbee** | [**InlineObject209**](InlineObject209.md) |  | 

### Return type

[**InlineResponse20117**](InlineResponse20117.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessZigbeeDevice

> InlineResponse200379 UpdateOrganizationWirelessZigbeeDevice(ctx, organizationId, id).UpdateOrganizationWirelessZigbeeDevice(updateOrganizationWirelessZigbeeDevice).Execute()

Endpoint to update zigbee gateways



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
    id := "id_example" // string | ID
    updateOrganizationWirelessZigbeeDevice := *openapiclient.NewInlineObject311(false) // InlineObject311 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.UpdateOrganizationWirelessZigbeeDevice(context.Background(), organizationId, id).UpdateOrganizationWirelessZigbeeDevice(updateOrganizationWirelessZigbeeDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.UpdateOrganizationWirelessZigbeeDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessZigbeeDevice`: InlineResponse200379
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.UpdateOrganizationWirelessZigbeeDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessZigbeeDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationWirelessZigbeeDevice** | [**InlineObject311**](InlineObject311.md) |  | 

### Return type

[**InlineResponse200379**](InlineResponse200379.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessZigbeeDoorLock

> InlineResponse20049DoorLocks UpdateOrganizationWirelessZigbeeDoorLock(ctx, organizationId, doorLockId).UpdateOrganizationWirelessZigbeeDoorLock(updateOrganizationWirelessZigbeeDoorLock).Execute()

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
    updateOrganizationWirelessZigbeeDoorLock := *openapiclient.NewInlineObject313() // InlineObject313 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZigbeeApi.UpdateOrganizationWirelessZigbeeDoorLock(context.Background(), organizationId, doorLockId).UpdateOrganizationWirelessZigbeeDoorLock(updateOrganizationWirelessZigbeeDoorLock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZigbeeApi.UpdateOrganizationWirelessZigbeeDoorLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessZigbeeDoorLock`: InlineResponse20049DoorLocks
    fmt.Fprintf(os.Stdout, "Response from `ZigbeeApi.UpdateOrganizationWirelessZigbeeDoorLock`: %v\n", resp)
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


 **updateOrganizationWirelessZigbeeDoorLock** | [**InlineObject313**](InlineObject313.md) |  | 

### Return type

[**InlineResponse20049DoorLocks**](InlineResponse20049DoorLocks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

