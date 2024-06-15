# \BySwitchApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSwitchPortsBySwitch**](BySwitchApi.md#GetOrganizationSwitchPortsBySwitch) | **Get** /organizations/{organizationId}/switch/ports/bySwitch | List the switchports in an organization by switch



## GetOrganizationSwitchPortsBySwitch

> []InlineResponse200289 GetOrganizationSwitchPortsBySwitch(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()

List the switchports in an organization by switch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    configurationUpdatedAfter := time.Now() // time.Time | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. (optional)
    mac := "mac_example" // string | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided MAC addresses. (optional)
    name := "name_example" // string | Optional parameter to filter items to switches with names that contain the search term or are an exact match. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches in one of the provided networks. (optional)
    portProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. (optional)
    serial := "serial_example" // string | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided serials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BySwitchApi.GetOrganizationSwitchPortsBySwitch(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BySwitchApi.GetOrganizationSwitchPortsBySwitch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSwitchPortsBySwitch`: []InlineResponse200289
    fmt.Fprintf(os.Stdout, "Response from `BySwitchApi.GetOrganizationSwitchPortsBySwitch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsBySwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **time.Time** | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. | 
 **mac** | **string** | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. | 
 **macs** | **[]string** | Optional parameter to filter items to switches that have one of the provided MAC addresses. | 
 **name** | **string** | Optional parameter to filter items to switches with names that contain the search term or are an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter items to switches in one of the provided networks. | 
 **portProfileIds** | **[]string** | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. | 
 **serial** | **string** | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. | 
 **serials** | **[]string** | Optional parameter to filter items to switches that have one of the provided serials. | 

### Return type

[**[]InlineResponse200289**](InlineResponse200289.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

