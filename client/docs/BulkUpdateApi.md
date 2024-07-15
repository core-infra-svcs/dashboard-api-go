# \BulkUpdateApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateOrganizationDevicesDetails**](BulkUpdateApi.md#BulkUpdateOrganizationDevicesDetails) | **Post** /organizations/{organizationId}/devices/details/bulkUpdate | Updating device details (currently only used for Catalyst devices)



## BulkUpdateOrganizationDevicesDetails

> InlineResponse200237 BulkUpdateOrganizationDevicesDetails(ctx, organizationId).BulkUpdateOrganizationDevicesDetails(bulkUpdateOrganizationDevicesDetails).Execute()

Updating device details (currently only used for Catalyst devices)



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
    bulkUpdateOrganizationDevicesDetails := *openapiclient.NewInlineObject229([]string{"Serials_example"}, []openapiclient.OrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails{*openapiclient.NewOrganizationsOrganizationIdDevicesDetailsBulkUpdateDetails("Name_example")}) // InlineObject229 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkUpdateApi.BulkUpdateOrganizationDevicesDetails(context.Background(), organizationId).BulkUpdateOrganizationDevicesDetails(bulkUpdateOrganizationDevicesDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateApi.BulkUpdateOrganizationDevicesDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateOrganizationDevicesDetails`: InlineResponse200237
    fmt.Fprintf(os.Stdout, "Response from `BulkUpdateApi.BulkUpdateOrganizationDevicesDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateOrganizationDevicesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateOrganizationDevicesDetails** | [**InlineObject229**](InlineObject229.md) |  | 

### Return type

[**InlineResponse200237**](InlineResponse200237.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

