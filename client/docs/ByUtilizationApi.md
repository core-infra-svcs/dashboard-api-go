# \ByUtilizationApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSummaryTopAppliancesByUtilization**](ByUtilizationApi.md#GetOrganizationSummaryTopAppliancesByUtilization) | **Get** /organizations/{organizationId}/summary/top/appliances/byUtilization | Return the top 10 appliances sorted by utilization over given time range.



## GetOrganizationSummaryTopAppliancesByUtilization

> []InlineResponse200324 GetOrganizationSummaryTopAppliancesByUtilization(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the top 10 appliances sorted by utilization over given time range.



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
    networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
    deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
    quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
    ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
    usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ByUtilizationApi.GetOrganizationSummaryTopAppliancesByUtilization(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ByUtilizationApi.GetOrganizationSummaryTopAppliancesByUtilization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopAppliancesByUtilization`: []InlineResponse200324
    fmt.Fprintf(os.Stdout, "Response from `ByUtilizationApi.GetOrganizationSummaryTopAppliancesByUtilization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopAppliancesByUtilizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200324**](InlineResponse200324.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

