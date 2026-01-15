# \CertificatesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationNacCertificatesAuthoritiesCrl**](CertificatesApi.md#CreateOrganizationNacCertificatesAuthoritiesCrl) | **Post** /organizations/{organizationId}/nac/certificates/authorities/crls | Create a new CRL (either base or delta) for an existing CA
[**CreateOrganizationWirelessDevicesRadsecCertificatesAuthority**](CertificatesApi.md#CreateOrganizationWirelessDevicesRadsecCertificatesAuthority) | **Post** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Create an organization&#39;s RADSEC device Certificate Authority (CA)
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthorities**](CertificatesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Query for details on the organization&#39;s RADSEC device Certificate Authority certificates (CAs)
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls**](CertificatesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls | Query for certificate revocation list (CRL) for the organization&#39;s RADSEC device Certificate Authorities (CAs).
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas**](CertificatesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls/deltas | Query for all delta certificate revocation list (CRL) for the organization&#39;s RADSEC device Certificate Authority (CA) with the given id.
[**UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities**](CertificatesApi.md#UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Put** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Update an organization&#39;s RADSEC device Certificate Authority (CA) state



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
    resp, r, err := apiClient.CertificatesApi.CreateOrganizationNacCertificatesAuthoritiesCrl(context.Background(), organizationId).CreateOrganizationNacCertificatesAuthoritiesCrl(createOrganizationNacCertificatesAuthoritiesCrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CreateOrganizationNacCertificatesAuthoritiesCrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationNacCertificatesAuthoritiesCrl`: InlineResponse20124
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CreateOrganizationNacCertificatesAuthoritiesCrl`: %v\n", resp)
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


## CreateOrganizationWirelessDevicesRadsecCertificatesAuthority

> InlineResponse200376 CreateOrganizationWirelessDevicesRadsecCertificatesAuthority(ctx, organizationId).Execute()

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
    // response from `CreateOrganizationWirelessDevicesRadsecCertificatesAuthority`: InlineResponse200376
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

[**InlineResponse200376**](InlineResponse200376.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthorities

> []InlineResponse200375 GetOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

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
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthorities`: []InlineResponse200375
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

[**[]InlineResponse200375**](InlineResponse200375.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls

> InlineResponse200377 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

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
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls`: InlineResponse200377
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

[**InlineResponse200377**](InlineResponse200377.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas

> InlineResponse200377 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

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
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas`: InlineResponse200377
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

[**InlineResponse200377**](InlineResponse200377.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities

> InlineResponse200376 UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()

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
    updateOrganizationWirelessDevicesRadsecCertificatesAuthorities := *openapiclient.NewInlineObject307() // InlineObject307 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificatesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(context.Background(), organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities`: InlineResponse200376
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

 **updateOrganizationWirelessDevicesRadsecCertificatesAuthorities** | [**InlineObject307**](InlineObject307.md) |  | 

### Return type

[**InlineResponse200376**](InlineResponse200376.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

