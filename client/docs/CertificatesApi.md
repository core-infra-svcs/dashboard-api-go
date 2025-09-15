# \CertificatesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationWirelessDevicesRadsecCertificatesAuthority**](CertificatesApi.md#CreateOrganizationWirelessDevicesRadsecCertificatesAuthority) | **Post** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Create an organization&#39;s RADSEC device Certificate Authority (CA)
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthorities**](CertificatesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Query for details on the organization&#39;s RADSEC device Certificate Authority certificates (CAs)
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls**](CertificatesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls | Query for certificate revocation list (CRL) for the organization&#39;s RADSEC device Certificate Authorities (CAs).
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas**](CertificatesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls/deltas | Query for all delta certificate revocation list (CRL) for the organization&#39;s RADSEC device Certificate Authority (CA) with the given id.
[**UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities**](CertificatesApi.md#UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Put** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Update an organization&#39;s RADSEC device Certificate Authority (CA) state



## CreateOrganizationWirelessDevicesRadsecCertificatesAuthority

> InlineResponse200366 CreateOrganizationWirelessDevicesRadsecCertificatesAuthority(ctx, organizationId).Execute()

Create an organization's RADSEC device Certificate Authority (CA)



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
    resp, r, err := apiClient.CertificatesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessDevicesRadsecCertificatesAuthority`: InlineResponse200366
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessDevicesRadsecCertificatesAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200366**](InlineResponse200366.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthorities

> []InlineResponse200365 GetOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for details on the organization's RADSEC device Certificate Authority certificates (CAs)



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
    certificateAuthorityIds := []string{"Inner_example"} // []string | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthorities`: []InlineResponse200365
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateAuthorityIds** | **[]string** | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. | 

### Return type

[**[]InlineResponse200365**](InlineResponse200365.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls

> InlineResponse200367 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for certificate revocation list (CRL) for the organization's RADSEC device Certificate Authorities (CAs).



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
    certificateAuthorityIds := []string{"Inner_example"} // []string | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateAuthorityIds** | **[]string** | Optional parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. | 

### Return type

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas

> InlineResponse200367 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for all delta certificate revocation list (CRL) for the organization's RADSEC device Certificate Authority (CA) with the given id.



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
    certificateAuthorityIds := []string{"Inner_example"} // []string | Parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateAuthorityIds** | **[]string** | Parameter to filter CAs by one or more CA IDs. All returned CAs will have an ID that is an exact match. | 

### Return type

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities

> InlineResponse200366 UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()

Update an organization's RADSEC device Certificate Authority (CA) state



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
    updateOrganizationWirelessDevicesRadsecCertificatesAuthorities := *openapiclient.NewInlineObject303() // InlineObject303 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(context.Background(), organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities`: InlineResponse200366
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessDevicesRadsecCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationWirelessDevicesRadsecCertificatesAuthorities** | [**InlineObject303**](InlineObject303.md) |  | 

### Return type

[**InlineResponse200366**](InlineResponse200366.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

