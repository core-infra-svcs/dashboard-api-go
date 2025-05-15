# \HistoricalApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationAssuranceAlertsOverviewHistorical**](HistoricalApi.md#GetOrganizationAssuranceAlertsOverviewHistorical) | **Get** /organizations/{organizationId}/assurance/alerts/overview/historical | Returns historical health alert overviews



## GetOrganizationAssuranceAlertsOverviewHistorical

> InlineResponse200240 GetOrganizationAssuranceAlertsOverviewHistorical(ctx, organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).Execute()

Returns historical health alert overviews



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
    segmentDuration := int32(56) // int32 | Amount of time in seconds for each segment in the returned dataset
    tsStart := time.Now() // time.Time | Parameter to define starting timestamp of historical totals
    networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network ids. (optional)
    severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
    types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
    tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp defaults to the current time (optional)
    category := "category_example" // string | Optional parameter to filter by category. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
    deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoricalApi.GetOrganizationAssuranceAlertsOverviewHistorical(context.Background(), organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoricalApi.GetOrganizationAssuranceAlertsOverviewHistorical``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverviewHistorical`: InlineResponse200240
    fmt.Fprintf(os.Stdout, "Response from `HistoricalApi.GetOrganizationAssuranceAlertsOverviewHistorical`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentDuration** | **int32** | Amount of time in seconds for each segment in the returned dataset | 
 **tsStart** | **time.Time** | Parameter to define starting timestamp of historical totals | 
 **networkId** | **string** | Optional parameter to filter alerts overview by network ids. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp defaults to the current time | 
 **category** | **string** | Optional parameter to filter by category. | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 

### Return type

[**InlineResponse200240**](InlineResponse200240.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

