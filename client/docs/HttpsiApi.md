# \HttpsiApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificatesOrganizationApplianceSecurityHttpsi**](HttpsiApi.md#CertificatesOrganizationApplianceSecurityHttpsi) | **Post** /organizations/{organizationId}/appliance/security/httpsi/certificates | Upload an HTTPS Inspection certificate to MX devices in the same organization



## CertificatesOrganizationApplianceSecurityHttpsi

> InlineResponse200297 CertificatesOrganizationApplianceSecurityHttpsi(ctx, organizationId).CertificatesOrganizationApplianceSecurityHttpsi(certificatesOrganizationApplianceSecurityHttpsi).Execute()

Upload an HTTPS Inspection certificate to MX devices in the same organization



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
    certificatesOrganizationApplianceSecurityHttpsi := *openapiclient.NewInlineObject264("Contents_example", []string{"Serials_example"}) // InlineObject264 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpsiApi.CertificatesOrganizationApplianceSecurityHttpsi(context.Background(), organizationId).CertificatesOrganizationApplianceSecurityHttpsi(certificatesOrganizationApplianceSecurityHttpsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpsiApi.CertificatesOrganizationApplianceSecurityHttpsi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificatesOrganizationApplianceSecurityHttpsi`: InlineResponse200297
    fmt.Fprintf(os.Stdout, "Response from `HttpsiApi.CertificatesOrganizationApplianceSecurityHttpsi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesOrganizationApplianceSecurityHttpsiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificatesOrganizationApplianceSecurityHttpsi** | [**InlineObject264**](InlineObject264.md) |  | 

### Return type

[**InlineResponse200297**](InlineResponse200297.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

