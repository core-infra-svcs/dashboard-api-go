# \BySiteApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSaseSitesConnectivityHistoryBySite**](BySiteApi.md#GetOrganizationSaseSitesConnectivityHistoryBySite) | **Get** /organizations/{organizationId}/sase/sites/connectivity/history/bySite | Get the connectivity history of SASE sites in this organization



## GetOrganizationSaseSitesConnectivityHistoryBySite

> InlineResponse200424 GetOrganizationSaseSitesConnectivityHistoryBySite(ctx, organizationId).SiteIds(siteIds).Timespan(timespan).Execute()

Get the connectivity history of SASE sites in this organization



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
    siteIds := []string{"Inner_example"} // []string | Array of site IDs to fetch connectivity status data (maximum 100 values) (optional)
    timespan := "timespan_example" // string | Timespan for the status data (e.g., '-2hours') (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BySiteApi.GetOrganizationSaseSitesConnectivityHistoryBySite(context.Background(), organizationId).SiteIds(siteIds).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BySiteApi.GetOrganizationSaseSitesConnectivityHistoryBySite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseSitesConnectivityHistoryBySite`: InlineResponse200424
    fmt.Fprintf(os.Stdout, "Response from `BySiteApi.GetOrganizationSaseSitesConnectivityHistoryBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseSitesConnectivityHistoryBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteIds** | **[]string** | Array of site IDs to fetch connectivity status data (maximum 100 values) | 
 **timespan** | **string** | Timespan for the status data (e.g., &#39;-2hours&#39;) | 

### Return type

[**InlineResponse200424**](InlineResponse200424.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

