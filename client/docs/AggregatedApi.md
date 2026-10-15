# \AggregatedApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated**](AggregatedApi.md#GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated) | **Get** /organizations/{organizationId}/summary/sustainability/summaryConsumerMetrics/deviceCategories/aggregated | Lists summary consumer metrics aggregated by device category.
[**GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated**](AggregatedApi.md#GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated) | **Get** /organizations/{organizationId}/summary/sustainability/summaryConsumerMetrics/networks/aggregated | Returns aggregated consumer metrics for all networks in the organization



## GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated

> InlineResponse200459 GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Lists summary consumer metrics aggregated by device category.



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 730 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 32 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 32 days. The default is 30 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AggregatedApi.GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AggregatedApi.GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated`: InlineResponse200459
    fmt.Fprintf(os.Stdout, "Response from `AggregatedApi.GetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummarySustainabilitySummaryConsumerMetricsDeviceCategoriesAggregatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 730 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 32 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 32 days. The default is 30 days. | 

### Return type

[**InlineResponse200459**](InlineResponse200459.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated

> InlineResponse200462 GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns aggregated consumer metrics for all networks in the organization



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 730 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 32 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 32 days. The default is 30 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AggregatedApi.GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AggregatedApi.GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated`: InlineResponse200462
    fmt.Fprintf(os.Stdout, "Response from `AggregatedApi.GetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummarySustainabilitySummaryConsumerMetricsNetworksAggregatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 730 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 32 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 32 days. The default is 30 days. | 

### Return type

[**InlineResponse200462**](InlineResponse200462.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

