# \ByModelApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationDevicesOverviewByModel**](ByModelApi.md#GetOrganizationDevicesOverviewByModel) | **Get** /organizations/{organizationId}/devices/overview/byModel | Lists the count for each device model



## GetOrganizationDevicesOverviewByModel

> InlineResponse200285 GetOrganizationDevicesOverviewByModel(ctx, organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()

Lists the count for each device model



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
    models := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by networkId. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ByModelApi.GetOrganizationDevicesOverviewByModel(context.Background(), organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ByModelApi.GetOrganizationDevicesOverviewByModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesOverviewByModel`: InlineResponse200285
    fmt.Fprintf(os.Stdout, "Response from `ByModelApi.GetOrganizationDevicesOverviewByModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesOverviewByModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **models** | **[]string** | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by networkId. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 

### Return type

[**InlineResponse200285**](InlineResponse200285.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

