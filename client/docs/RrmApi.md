# \RrmApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWirelessRadioRrmByNetwork**](RrmApi.md#GetOrganizationWirelessRadioRrmByNetwork) | **Get** /organizations/{organizationId}/wireless/radio/rrm/byNetwork | List the AutoRF settings of an organization by network
[**UpdateNetworkWirelessRadioRrm**](RrmApi.md#UpdateNetworkWirelessRadioRrm) | **Put** /networks/{networkId}/wireless/radio/rrm | Update the AutoRF settings for a wireless network



## GetOrganizationWirelessRadioRrmByNetwork

> InlineResponse200387 GetOrganizationWirelessRadioRrmByNetwork(ctx, organizationId).NetworkIds(networkIds).StartingAfter(startingAfter).EndingBefore(endingBefore).PerPage(perPage).SortOrder(sortOrder).Execute()

List the AutoRF settings of an organization by network



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter results by network. (optional)
    startingAfter := "startingAfter_example" // string | Retrieving items after this network ID (optional)
    endingBefore := "endingBefore_example" // string | Retrieving items before this network ID (optional)
    perPage := int32(56) // int32 | Number of items per page (optional)
    sortOrder := "sortOrder_example" // string | The sort order of items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RrmApi.GetOrganizationWirelessRadioRrmByNetwork(context.Background(), organizationId).NetworkIds(networkIds).StartingAfter(startingAfter).EndingBefore(endingBefore).PerPage(perPage).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RrmApi.GetOrganizationWirelessRadioRrmByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessRadioRrmByNetwork`: InlineResponse200387
    fmt.Fprintf(os.Stdout, "Response from `RrmApi.GetOrganizationWirelessRadioRrmByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessRadioRrmByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter results by network. | 
 **startingAfter** | **string** | Retrieving items after this network ID | 
 **endingBefore** | **string** | Retrieving items before this network ID | 
 **perPage** | **int32** | Number of items per page | 
 **sortOrder** | **string** | The sort order of items | 

### Return type

[**InlineResponse200387**](InlineResponse200387.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessRadioRrm

> InlineResponse200205 UpdateNetworkWirelessRadioRrm(ctx, networkId).UpdateNetworkWirelessRadioRrm(updateNetworkWirelessRadioRrm).Execute()

Update the AutoRF settings for a wireless network



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
    updateNetworkWirelessRadioRrm := *openapiclient.NewInlineObject192() // InlineObject192 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RrmApi.UpdateNetworkWirelessRadioRrm(context.Background(), networkId).UpdateNetworkWirelessRadioRrm(updateNetworkWirelessRadioRrm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RrmApi.UpdateNetworkWirelessRadioRrm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessRadioRrm`: InlineResponse200205
    fmt.Fprintf(os.Stdout, "Response from `RrmApi.UpdateNetworkWirelessRadioRrm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessRadioRrmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessRadioRrm** | [**InlineObject192**](InlineObject192.md) |  | 

### Return type

[**InlineResponse200205**](InlineResponse200205.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

