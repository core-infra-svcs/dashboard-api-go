# \NacApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationNacCertificatesAuthoritiesCrl**](NacApi.md#CreateOrganizationNacCertificatesAuthoritiesCrl) | **Post** /organizations/{organizationId}/nac/certificates/authorities/crls | Create a new CRL (either base or delta) for an existing CA



## CreateOrganizationNacCertificatesAuthoritiesCrl

> InlineResponse20124 CreateOrganizationNacCertificatesAuthoritiesCrl(ctx, organizationId).CreateOrganizationNacCertificatesAuthoritiesCrl(createOrganizationNacCertificatesAuthoritiesCrl).Execute()

Create a new CRL (either base or delta) for an existing CA



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
    createOrganizationNacCertificatesAuthoritiesCrl := *openapiclient.NewInlineObject288("CaId_example", "Content_example", false) // InlineObject288 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NacApi.CreateOrganizationNacCertificatesAuthoritiesCrl(context.Background(), organizationId).CreateOrganizationNacCertificatesAuthoritiesCrl(createOrganizationNacCertificatesAuthoritiesCrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NacApi.CreateOrganizationNacCertificatesAuthoritiesCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationNacCertificatesAuthoritiesCrl`: InlineResponse20124
    fmt.Fprintf(os.Stdout, "Response from `NacApi.CreateOrganizationNacCertificatesAuthoritiesCrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationNacCertificatesAuthoritiesCrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationNacCertificatesAuthoritiesCrl** | [**InlineObject288**](InlineObject288.md) |  | 

### Return type

[**InlineResponse20124**](InlineResponse20124.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

