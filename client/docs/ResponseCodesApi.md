# \ResponseCodesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApiRequestsOverviewResponseCodesByInterval**](ResponseCodesApi.md#GetOrganizationApiRequestsOverviewResponseCodesByInterval) | **Get** /organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval | Tracks organizations&#39; API requests by response code across a given time period



## GetOrganizationApiRequestsOverviewResponseCodesByInterval

> []InlineResponse200220 GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()

Tracks organizations' API requests by response code across a given time period



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. (optional)
    version := int32(56) // int32 | Filter by API version of the endpoint. Allowable values are: [0, 1] (optional)
    operationIds := []string{"Inner_example"} // []string | Filter by operation ID of the endpoint (optional)
    sourceIps := []string{"Inner_example"} // []string | Filter by source IP that made the API request (optional)
    adminIds := []string{"Inner_example"} // []string | Filter by admin ID of user that made the API request (optional)
    userAgent := "userAgent_example" // string | Filter by user agent string for API request. This will filter by a complete or partial match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResponseCodesApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResponseCodesApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRequestsOverviewResponseCodesByInterval`: []InlineResponse200220
    fmt.Fprintf(os.Stdout, "Response from `ResponseCodesApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. | 
 **version** | **int32** | Filter by API version of the endpoint. Allowable values are: [0, 1] | 
 **operationIds** | **[]string** | Filter by operation ID of the endpoint | 
 **sourceIps** | **[]string** | Filter by source IP that made the API request | 
 **adminIds** | **[]string** | Filter by admin ID of user that made the API request | 
 **userAgent** | **string** | Filter by user agent string for API request. This will filter by a complete or partial match. | 

### Return type

[**[]InlineResponse200220**](InlineResponse200220.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

