# \ImpactedDeviceApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationAssuranceImpactedDeviceWirelessByNetwork**](ImpactedDeviceApi.md#GetOrganizationAssuranceImpactedDeviceWirelessByNetwork) | **Get** /organizations/{organizationId}/assurance/impactedDevice/wireless/byNetwork | Returns count of impacted wireless devices per network on a given organization and time range.



## GetOrganizationAssuranceImpactedDeviceWirelessByNetwork

> InlineResponse200307 GetOrganizationAssuranceImpactedDeviceWirelessByNetwork(ctx, organizationId).NetworkGroupIds(networkGroupIds).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns count of impacted wireless devices per network on a given organization and time range.



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
    networkGroupIds := []string{"Inner_example"} // []string | Filter results by a list of network group IDs. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 14 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 2 hours and be less than or equal to 14 days. The default is 2 hours. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImpactedDeviceApi.GetOrganizationAssuranceImpactedDeviceWirelessByNetwork(context.Background(), organizationId).NetworkGroupIds(networkGroupIds).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImpactedDeviceApi.GetOrganizationAssuranceImpactedDeviceWirelessByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceImpactedDeviceWirelessByNetwork`: InlineResponse200307
    fmt.Fprintf(os.Stdout, "Response from `ImpactedDeviceApi.GetOrganizationAssuranceImpactedDeviceWirelessByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceImpactedDeviceWirelessByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkGroupIds** | **[]string** | Filter results by a list of network group IDs. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 14 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 2 hours and be less than or equal to 14 days. The default is 2 hours. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200307**](InlineResponse200307.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

