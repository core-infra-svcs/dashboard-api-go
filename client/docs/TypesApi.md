# \TypesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationAssuranceAlertsTaxonomyTypes**](TypesApi.md#GetOrganizationAssuranceAlertsTaxonomyTypes) | **Get** /organizations/{organizationId}/assurance/alerts/taxonomy/types | Return a list of alert types



## GetOrganizationAssuranceAlertsTaxonomyTypes

> []InlineResponse200255 GetOrganizationAssuranceAlertsTaxonomyTypes(ctx, organizationId).Execute()

Return a list of alert types



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
    resp, r, err := apiClient.TypesApi.GetOrganizationAssuranceAlertsTaxonomyTypes(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesApi.GetOrganizationAssuranceAlertsTaxonomyTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsTaxonomyTypes`: []InlineResponse200255
    fmt.Fprintf(os.Stdout, "Response from `TypesApi.GetOrganizationAssuranceAlertsTaxonomyTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsTaxonomyTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200255**](InlineResponse200255.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

