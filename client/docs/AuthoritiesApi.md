# \AuthoritiesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCertificatesAuthority**](AuthoritiesApi.md#CreateOrganizationCertificatesAuthority) | **Post** /organizations/{organizationId}/certificates/authorities | Create a certificate authority for an organization
[**CreateOrganizationWirelessDevicesRadsecCertificatesAuthority**](AuthoritiesApi.md#CreateOrganizationWirelessDevicesRadsecCertificatesAuthority) | **Post** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Create an organization&#39;s RadSec device Certificate Authority (CA)
[**DeleteOrganizationCertificatesAuthorities**](AuthoritiesApi.md#DeleteOrganizationCertificatesAuthorities) | **Delete** /organizations/{organizationId}/certificates/authorities | Delete a certificate authority
[**GetOrganizationCertificatesAuthorities**](AuthoritiesApi.md#GetOrganizationCertificatesAuthorities) | **Get** /organizations/{organizationId}/certificates/authorities | List certificate authorities for an organization
[**GetOrganizationCertificatesAuthoritiesJob**](AuthoritiesApi.md#GetOrganizationCertificatesAuthoritiesJob) | **Get** /organizations/{organizationId}/certificates/authorities/jobs/{jobId} | Return the status and result of a certificate authority job.
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthorities**](AuthoritiesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Query for details on the organization&#39;s RadSec device Certificate Authority certificates (CAs)
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls**](AuthoritiesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls | Query for certificate revocation list (CRL) for the organization&#39;s RadSec device Certificate Authorities (CAs).
[**GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas**](AuthoritiesApi.md#GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas) | **Get** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities/crls/deltas | Query for all delta certificate revocation list (CRL) for the organization&#39;s RadSec device Certificate Authority (CA) with the given id.
[**RevokeOrganizationCertificatesAuthorities**](AuthoritiesApi.md#RevokeOrganizationCertificatesAuthorities) | **Post** /organizations/{organizationId}/certificates/authorities/revoke | Revoke a trusted feature certificate authority.
[**UpdateOrganizationCertificatesAuthorities**](AuthoritiesApi.md#UpdateOrganizationCertificatesAuthorities) | **Put** /organizations/{organizationId}/certificates/authorities | Trust a newly created certificate authority (transition from untrusted to trusted).
[**UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities**](AuthoritiesApi.md#UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities) | **Put** /organizations/{organizationId}/wireless/devices/radsec/certificates/authorities | Update an organization&#39;s RadSec device Certificate Authority (CA) state



## CreateOrganizationCertificatesAuthority

> InlineResponse2025 CreateOrganizationCertificatesAuthority(ctx, organizationId).CreateOrganizationCertificatesAuthority(createOrganizationCertificatesAuthority).Execute()

Create a certificate authority for an organization



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
    createOrganizationCertificatesAuthority := *openapiclient.NewInlineObject295("FeatureType_example") // InlineObject295 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthoritiesApi.CreateOrganizationCertificatesAuthority(context.Background(), organizationId).CreateOrganizationCertificatesAuthority(createOrganizationCertificatesAuthority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.CreateOrganizationCertificatesAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCertificatesAuthority`: InlineResponse2025
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.CreateOrganizationCertificatesAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCertificatesAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCertificatesAuthority** | [**InlineObject295**](InlineObject295.md) |  | 

### Return type

[**InlineResponse2025**](InlineResponse2025.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationWirelessDevicesRadsecCertificatesAuthority

> InlineResponse200515 CreateOrganizationWirelessDevicesRadsecCertificatesAuthority(ctx, organizationId).Execute()

Create an organization's RadSec device Certificate Authority (CA)



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
    resp, r, err := apiClient.AuthoritiesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessDevicesRadsecCertificatesAuthority`: InlineResponse200515
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.CreateOrganizationWirelessDevicesRadsecCertificatesAuthority`: %v\n", resp)
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

[**InlineResponse200515**](InlineResponse200515.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationCertificatesAuthorities

> InlineResponse200360 DeleteOrganizationCertificatesAuthorities(ctx, organizationId).DeleteOrganizationCertificatesAuthorities(deleteOrganizationCertificatesAuthorities).Execute()

Delete a certificate authority



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
    deleteOrganizationCertificatesAuthorities := *openapiclient.NewInlineObject296("AuthorityId_example", "Name_example") // InlineObject296 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthoritiesApi.DeleteOrganizationCertificatesAuthorities(context.Background(), organizationId).DeleteOrganizationCertificatesAuthorities(deleteOrganizationCertificatesAuthorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.DeleteOrganizationCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganizationCertificatesAuthorities`: InlineResponse200360
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.DeleteOrganizationCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteOrganizationCertificatesAuthorities** | [**InlineObject296**](InlineObject296.md) |  | 

### Return type

[**InlineResponse200360**](InlineResponse200360.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCertificatesAuthorities

> InlineResponse200358 GetOrganizationCertificatesAuthorities(ctx, organizationId).AuthorityIds(authorityIds).SortBy(sortBy).Direction(direction).IncludeChain(includeChain).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List certificate authorities for an organization



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
    authorityIds := []string{"Inner_example"} // []string | Feature certificate authority IDs to filter by (exact match on each id; duplicates are ignored) (optional)
    sortBy := "sortBy_example" // string | Field to sort by (default: authorityId) (optional)
    direction := "direction_example" // string | Sort direction (default: asc) (optional)
    includeChain := true // bool | When true, include the certification path from each certificate authority up to the self signed root. Required by validators that cannot anchor trust at a non self signed certificate authority. Defaults to false. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthoritiesApi.GetOrganizationCertificatesAuthorities(context.Background(), organizationId).AuthorityIds(authorityIds).SortBy(sortBy).Direction(direction).IncludeChain(includeChain).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.GetOrganizationCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCertificatesAuthorities`: InlineResponse200358
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.GetOrganizationCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorityIds** | **[]string** | Feature certificate authority IDs to filter by (exact match on each id; duplicates are ignored) | 
 **sortBy** | **string** | Field to sort by (default: authorityId) | 
 **direction** | **string** | Sort direction (default: asc) | 
 **includeChain** | **bool** | When true, include the certification path from each certificate authority up to the self signed root. Required by validators that cannot anchor trust at a non self signed certificate authority. Defaults to false. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200358**](InlineResponse200358.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCertificatesAuthoritiesJob

> InlineResponse2025 GetOrganizationCertificatesAuthoritiesJob(ctx, organizationId, jobId).Execute()

Return the status and result of a certificate authority job.



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
    jobId := "jobId_example" // string | Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthoritiesApi.GetOrganizationCertificatesAuthoritiesJob(context.Background(), organizationId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.GetOrganizationCertificatesAuthoritiesJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCertificatesAuthoritiesJob`: InlineResponse2025
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.GetOrganizationCertificatesAuthoritiesJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCertificatesAuthoritiesJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2025**](InlineResponse2025.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthorities

> []InlineResponse200514 GetOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for details on the organization's RadSec device Certificate Authority certificates (CAs)



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
    resp, r, err := apiClient.AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthorities`: []InlineResponse200514
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthorities`: %v\n", resp)
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

[**[]InlineResponse200514**](InlineResponse200514.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls

> InlineResponse200516 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for certificate revocation list (CRL) for the organization's RadSec device Certificate Authorities (CAs).



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
    resp, r, err := apiClient.AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls`: InlineResponse200516
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrls`: %v\n", resp)
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

[**InlineResponse200516**](InlineResponse200516.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas

> InlineResponse200516 GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas(ctx, organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()

Query for all delta certificate revocation list (CRL) for the organization's RadSec device Certificate Authority (CA) with the given id.



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
    resp, r, err := apiClient.AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas(context.Background(), organizationId).CertificateAuthorityIds(certificateAuthorityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas`: InlineResponse200516
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.GetOrganizationWirelessDevicesRadsecCertificatesAuthoritiesCrlsDeltas`: %v\n", resp)
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

[**InlineResponse200516**](InlineResponse200516.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOrganizationCertificatesAuthorities

> InlineResponse200359 RevokeOrganizationCertificatesAuthorities(ctx, organizationId).RevokeOrganizationCertificatesAuthorities(revokeOrganizationCertificatesAuthorities).Execute()

Revoke a trusted feature certificate authority.



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
    revokeOrganizationCertificatesAuthorities := *openapiclient.NewInlineObject297("AuthorityId_example") // InlineObject297 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthoritiesApi.RevokeOrganizationCertificatesAuthorities(context.Background(), organizationId).RevokeOrganizationCertificatesAuthorities(revokeOrganizationCertificatesAuthorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.RevokeOrganizationCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokeOrganizationCertificatesAuthorities`: InlineResponse200359
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.RevokeOrganizationCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOrganizationCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revokeOrganizationCertificatesAuthorities** | [**InlineObject297**](InlineObject297.md) |  | 

### Return type

[**InlineResponse200359**](InlineResponse200359.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCertificatesAuthorities

> InlineResponse200359 UpdateOrganizationCertificatesAuthorities(ctx, organizationId).UpdateOrganizationCertificatesAuthorities(updateOrganizationCertificatesAuthorities).Execute()

Trust a newly created certificate authority (transition from untrusted to trusted).



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
    updateOrganizationCertificatesAuthorities := *openapiclient.NewInlineObject294("AuthorityId_example") // InlineObject294 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthoritiesApi.UpdateOrganizationCertificatesAuthorities(context.Background(), organizationId).UpdateOrganizationCertificatesAuthorities(updateOrganizationCertificatesAuthorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.UpdateOrganizationCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCertificatesAuthorities`: InlineResponse200359
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.UpdateOrganizationCertificatesAuthorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCertificatesAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationCertificatesAuthorities** | [**InlineObject294**](InlineObject294.md) |  | 

### Return type

[**InlineResponse200359**](InlineResponse200359.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities

> InlineResponse200515 UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(ctx, organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()

Update an organization's RadSec device Certificate Authority (CA) state



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
    updateOrganizationWirelessDevicesRadsecCertificatesAuthorities := *openapiclient.NewInlineObject380() // InlineObject380 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthoritiesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(context.Background(), organizationId).UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities(updateOrganizationWirelessDevicesRadsecCertificatesAuthorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthoritiesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities`: InlineResponse200515
    fmt.Fprintf(os.Stdout, "Response from `AuthoritiesApi.UpdateOrganizationWirelessDevicesRadsecCertificatesAuthorities`: %v\n", resp)
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

 **updateOrganizationWirelessDevicesRadsecCertificatesAuthorities** | [**InlineObject380**](InlineObject380.md) |  | 

### Return type

[**InlineResponse200515**](InlineResponse200515.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

