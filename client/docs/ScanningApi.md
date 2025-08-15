# \ScanningApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationWirelessLocationScanningReceiver**](ScanningApi.md#CreateOrganizationWirelessLocationScanningReceiver) | **Post** /organizations/{organizationId}/wireless/location/scanning/receivers | Add new receiver for scanning API
[**DeleteOrganizationWirelessLocationScanningReceiver**](ScanningApi.md#DeleteOrganizationWirelessLocationScanningReceiver) | **Delete** /organizations/{organizationId}/wireless/location/scanning/receivers/{receiverId} | Delete a scanning API receiver
[**GetOrganizationWirelessLocationScanningByNetwork**](ScanningApi.md#GetOrganizationWirelessLocationScanningByNetwork) | **Get** /organizations/{organizationId}/wireless/location/scanning/byNetwork | Return scanning API settings
[**GetOrganizationWirelessLocationScanningReceivers**](ScanningApi.md#GetOrganizationWirelessLocationScanningReceivers) | **Get** /organizations/{organizationId}/wireless/location/scanning/receivers | Return scanning API receivers
[**UpdateNetworkWirelessLocationScanning**](ScanningApi.md#UpdateNetworkWirelessLocationScanning) | **Put** /networks/{networkId}/wireless/location/scanning | Change scanning API settings
[**UpdateOrganizationWirelessLocationScanningReceiver**](ScanningApi.md#UpdateOrganizationWirelessLocationScanningReceiver) | **Put** /organizations/{organizationId}/wireless/location/scanning/receivers/{receiverId} | Change scanning API receiver settings



## CreateOrganizationWirelessLocationScanningReceiver

> InlineResponse200368Items CreateOrganizationWirelessLocationScanningReceiver(ctx, organizationId).CreateOrganizationWirelessLocationScanningReceiver(createOrganizationWirelessLocationScanningReceiver).Execute()

Add new receiver for scanning API



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
    createOrganizationWirelessLocationScanningReceiver := *openapiclient.NewInlineObject303(*openapiclient.NewOrganizationsOrganizationIdWirelessLocationScanningReceiversNetwork(), "Url_example", "Version_example", *openapiclient.NewOrganizationsOrganizationIdWirelessLocationScanningReceiversRadio(), "SharedSecret_example") // InlineObject303 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanningApi.CreateOrganizationWirelessLocationScanningReceiver(context.Background(), organizationId).CreateOrganizationWirelessLocationScanningReceiver(createOrganizationWirelessLocationScanningReceiver).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.CreateOrganizationWirelessLocationScanningReceiver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessLocationScanningReceiver`: InlineResponse200368Items
    fmt.Fprintf(os.Stdout, "Response from `ScanningApi.CreateOrganizationWirelessLocationScanningReceiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessLocationScanningReceiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationWirelessLocationScanningReceiver** | [**InlineObject303**](InlineObject303.md) |  | 

### Return type

[**InlineResponse200368Items**](InlineResponse200368Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationWirelessLocationScanningReceiver

> DeleteOrganizationWirelessLocationScanningReceiver(ctx, organizationId, receiverId).Execute()

Delete a scanning API receiver



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
    receiverId := "receiverId_example" // string | Receiver ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanningApi.DeleteOrganizationWirelessLocationScanningReceiver(context.Background(), organizationId, receiverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.DeleteOrganizationWirelessLocationScanningReceiver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**receiverId** | **string** | Receiver ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationWirelessLocationScanningReceiverRequest struct via the builder pattern


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


## GetOrganizationWirelessLocationScanningByNetwork

> InlineResponse200367 GetOrganizationWirelessLocationScanningByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Return scanning API settings



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 250. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter scanning settings by network ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanningApi.GetOrganizationWirelessLocationScanningByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.GetOrganizationWirelessLocationScanningByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessLocationScanningByNetwork`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `ScanningApi.GetOrganizationWirelessLocationScanningByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessLocationScanningByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 250. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter scanning settings by network ID. | 

### Return type

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessLocationScanningReceivers

> InlineResponse200368 GetOrganizationWirelessLocationScanningReceivers(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Return scanning API receivers



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 250. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter scanning API receivers by network ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanningApi.GetOrganizationWirelessLocationScanningReceivers(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.GetOrganizationWirelessLocationScanningReceivers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessLocationScanningReceivers`: InlineResponse200368
    fmt.Fprintf(os.Stdout, "Response from `ScanningApi.GetOrganizationWirelessLocationScanningReceivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessLocationScanningReceiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 250. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter scanning API receivers by network ID. | 

### Return type

[**InlineResponse200368**](InlineResponse200368.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessLocationScanning

> InlineResponse200196 UpdateNetworkWirelessLocationScanning(ctx, networkId).UpdateNetworkWirelessLocationScanning(updateNetworkWirelessLocationScanning).Execute()

Change scanning API settings



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
    updateNetworkWirelessLocationScanning := *openapiclient.NewInlineObject190() // InlineObject190 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanningApi.UpdateNetworkWirelessLocationScanning(context.Background(), networkId).UpdateNetworkWirelessLocationScanning(updateNetworkWirelessLocationScanning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.UpdateNetworkWirelessLocationScanning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessLocationScanning`: InlineResponse200196
    fmt.Fprintf(os.Stdout, "Response from `ScanningApi.UpdateNetworkWirelessLocationScanning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessLocationScanningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessLocationScanning** | [**InlineObject190**](InlineObject190.md) |  | 

### Return type

[**InlineResponse200196**](InlineResponse200196.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessLocationScanningReceiver

> InlineResponse200368Items UpdateOrganizationWirelessLocationScanningReceiver(ctx, organizationId, receiverId).UpdateOrganizationWirelessLocationScanningReceiver(updateOrganizationWirelessLocationScanningReceiver).Execute()

Change scanning API receiver settings



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
    receiverId := "receiverId_example" // string | Receiver ID
    updateOrganizationWirelessLocationScanningReceiver := *openapiclient.NewInlineObject304() // InlineObject304 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanningApi.UpdateOrganizationWirelessLocationScanningReceiver(context.Background(), organizationId, receiverId).UpdateOrganizationWirelessLocationScanningReceiver(updateOrganizationWirelessLocationScanningReceiver).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanningApi.UpdateOrganizationWirelessLocationScanningReceiver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessLocationScanningReceiver`: InlineResponse200368Items
    fmt.Fprintf(os.Stdout, "Response from `ScanningApi.UpdateOrganizationWirelessLocationScanningReceiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**receiverId** | **string** | Receiver ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessLocationScanningReceiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationWirelessLocationScanningReceiver** | [**InlineObject304**](InlineObject304.md) |  | 

### Return type

[**InlineResponse200368Items**](InlineResponse200368Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

