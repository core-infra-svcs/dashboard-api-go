# \EoxApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationInventoryDevicesEoxOverview**](EoxApi.md#GetOrganizationInventoryDevicesEoxOverview) | **Get** /organizations/{organizationId}/inventory/devices/eox/overview | Fetch the EOX summary for an organization, including counts of devices that are end-of-sale, end-of-support, and end-of-support-soon.



## GetOrganizationInventoryDevicesEoxOverview

> InlineResponse200310 GetOrganizationInventoryDevicesEoxOverview(ctx, organizationId).Execute()

Fetch the EOX summary for an organization, including counts of devices that are end-of-sale, end-of-support, and end-of-support-soon.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EoxApi.GetOrganizationInventoryDevicesEoxOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EoxApi.GetOrganizationInventoryDevicesEoxOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevicesEoxOverview`: InlineResponse200310
    fmt.Fprintf(os.Stdout, "Response from `EoxApi.GetOrganizationInventoryDevicesEoxOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesEoxOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200310**](InlineResponse200310.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

