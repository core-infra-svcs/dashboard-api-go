# \RevocationListsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCertificatesRevocationLists**](RevocationListsApi.md#GetOrganizationCertificatesRevocationLists) | **Get** /organizations/{organizationId}/certificates/revocationLists | Return full certificate revocation lists (CRLs) for the organization&#39;s certificate authorities



## GetOrganizationCertificatesRevocationLists

> InlineResponse200361 GetOrganizationCertificatesRevocationLists(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Return full certificate revocation lists (CRLs) for the organization's certificate authorities



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
    certificateAuthorityIds := []string{"Inner_example"} // []string | Optional filter: feature certificate authority IDs (base-10 integers). Every value must exist for this organization; otherwise the request fails. Omit to return CRLs for all feature CAs in the organization. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevocationListsApi.GetOrganizationCertificatesRevocationLists(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevocationListsApi.GetOrganizationCertificatesRevocationLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCertificatesRevocationLists`: InlineResponse200361
    fmt.Fprintf(os.Stdout, "Response from `RevocationListsApi.GetOrganizationCertificatesRevocationLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCertificatesRevocationListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateAuthorityIds** | **[]string** | Optional filter: feature certificate authority IDs (base-10 integers). Every value must exist for this organization; otherwise the request fails. Omit to return CRLs for all feature CAs in the organization. | 

### Return type

[**InlineResponse200361**](InlineResponse200361.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

