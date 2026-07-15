# \SaseApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachOrganizationSaseSites**](SaseApi.md#AttachOrganizationSaseSites) | **Post** /organizations/{organizationId}/sase/sites/attach | Attach sites in this organization to Secure Access
[**BatchOrganizationSaseConnectorsDelete**](SaseApi.md#BatchOrganizationSaseConnectorsDelete) | **Post** /organizations/{organizationId}/sase/connectors/batchDelete | Delete SSE Connectors by ID
[**CreateOrganizationSaseIntegration**](SaseApi.md#CreateOrganizationSaseIntegration) | **Post** /organizations/{organizationId}/sase/integrations | Create a new Secure Access integration
[**DeleteOrganizationSaseIntegration**](SaseApi.md#DeleteOrganizationSaseIntegration) | **Delete** /organizations/{organizationId}/sase/integrations/{integrationId} | Remove a Secure Access integration
[**DetachOrganizationSaseSites**](SaseApi.md#DetachOrganizationSaseSites) | **Delete** /organizations/{organizationId}/sase/sites/detach | Detach sites in this organization from Secure Access
[**GetOrganizationSaseConnectors**](SaseApi.md#GetOrganizationSaseConnectors) | **Get** /organizations/{organizationId}/sase/connectors | List SSE Connectors for an organization
[**GetOrganizationSaseIntegration**](SaseApi.md#GetOrganizationSaseIntegration) | **Get** /organizations/{organizationId}/sase/integrations/{integrationId} | Get details of a Secure Access integration
[**GetOrganizationSaseNetworksEligible**](SaseApi.md#GetOrganizationSaseNetworksEligible) | **Get** /organizations/{organizationId}/sase/networks/eligible | List of MX networks or templates that can be enrolled into Secure Access
[**GetOrganizationSaseRegions**](SaseApi.md#GetOrganizationSaseRegions) | **Get** /organizations/{organizationId}/sase/regions | List regions
[**GetOrganizationSaseSites**](SaseApi.md#GetOrganizationSaseSites) | **Get** /organizations/{organizationId}/sase/sites | List of enrolled SASE sites in this organization
[**GetOrganizationSaseSitesConnectivityHistoryBySite**](SaseApi.md#GetOrganizationSaseSitesConnectivityHistoryBySite) | **Get** /organizations/{organizationId}/sase/sites/connectivity/history/bySite | Get the connectivity history of SASE sites in this organization
[**GetOrganizationSaseSitesConnectivityOverview**](SaseApi.md#GetOrganizationSaseSitesConnectivityOverview) | **Get** /organizations/{organizationId}/sase/sites/connectivity/overview | List high-level SASE site statuses (healthy, degraded, offline)
[**UpdateOrganizationSaseSite**](SaseApi.md#UpdateOrganizationSaseSite) | **Put** /organizations/{organizationId}/sase/sites/{siteId} | Update the configuration for a site



## AttachOrganizationSaseSites

> InlineResponse2024 AttachOrganizationSaseSites(ctx, organizationId).AttachOrganizationSaseSites(attachOrganizationSaseSites).Execute()

Attach sites in this organization to Secure Access



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
    attachOrganizationSaseSites := *openapiclient.NewInlineObject336() // InlineObject336 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.AttachOrganizationSaseSites(context.Background(), organizationId).AttachOrganizationSaseSites(attachOrganizationSaseSites).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.AttachOrganizationSaseSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachOrganizationSaseSites`: InlineResponse2024
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.AttachOrganizationSaseSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachOrganizationSaseSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachOrganizationSaseSites** | [**InlineObject336**](InlineObject336.md) |  | 

### Return type

[**InlineResponse2024**](InlineResponse2024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchOrganizationSaseConnectorsDelete

> InlineResponse2024 BatchOrganizationSaseConnectorsDelete(ctx, organizationId).BatchOrganizationSaseConnectorsDelete(batchOrganizationSaseConnectorsDelete).Execute()

Delete SSE Connectors by ID



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
    batchOrganizationSaseConnectorsDelete := *openapiclient.NewInlineObject334() // InlineObject334 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.BatchOrganizationSaseConnectorsDelete(context.Background(), organizationId).BatchOrganizationSaseConnectorsDelete(batchOrganizationSaseConnectorsDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.BatchOrganizationSaseConnectorsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationSaseConnectorsDelete`: InlineResponse2024
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.BatchOrganizationSaseConnectorsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchOrganizationSaseConnectorsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchOrganizationSaseConnectorsDelete** | [**InlineObject334**](InlineObject334.md) |  | 

### Return type

[**InlineResponse2024**](InlineResponse2024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSaseIntegration

> InlineResponse20128 CreateOrganizationSaseIntegration(ctx, organizationId).CreateOrganizationSaseIntegration(createOrganizationSaseIntegration).Execute()

Create a new Secure Access integration



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
    createOrganizationSaseIntegration := *openapiclient.NewInlineObject335(*openapiclient.NewOrganizationsOrganizationIdSaseIntegrationsApi("Key_example", "Secret_example")) // InlineObject335 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.CreateOrganizationSaseIntegration(context.Background(), organizationId).CreateOrganizationSaseIntegration(createOrganizationSaseIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.CreateOrganizationSaseIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSaseIntegration`: InlineResponse20128
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.CreateOrganizationSaseIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSaseIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSaseIntegration** | [**InlineObject335**](InlineObject335.md) |  | 

### Return type

[**InlineResponse20128**](InlineResponse20128.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSaseIntegration

> DeleteOrganizationSaseIntegration(ctx, organizationId, integrationId).Execute()

Remove a Secure Access integration



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
    integrationId := "integrationId_example" // string | Integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.DeleteOrganizationSaseIntegration(context.Background(), organizationId, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.DeleteOrganizationSaseIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**integrationId** | **string** | Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSaseIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachOrganizationSaseSites

> InlineResponse2024 DetachOrganizationSaseSites(ctx, organizationId).DetachOrganizationSaseSites(detachOrganizationSaseSites).Execute()

Detach sites in this organization from Secure Access



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
    detachOrganizationSaseSites := *openapiclient.NewInlineObject337() // InlineObject337 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.DetachOrganizationSaseSites(context.Background(), organizationId).DetachOrganizationSaseSites(detachOrganizationSaseSites).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.DetachOrganizationSaseSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachOrganizationSaseSites`: InlineResponse2024
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.DetachOrganizationSaseSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachOrganizationSaseSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detachOrganizationSaseSites** | [**InlineObject337**](InlineObject337.md) |  | 

### Return type

[**InlineResponse2024**](InlineResponse2024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseConnectors

> InlineResponse200401 GetOrganizationSaseConnectors(ctx, organizationId).Execute()

List SSE Connectors for an organization



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
    resp, r, err := apiClient.SaseApi.GetOrganizationSaseConnectors(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.GetOrganizationSaseConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseConnectors`: InlineResponse200401
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.GetOrganizationSaseConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200401**](InlineResponse200401.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseIntegration

> InlineResponse20128 GetOrganizationSaseIntegration(ctx, organizationId, integrationId).Execute()

Get details of a Secure Access integration



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
    integrationId := "integrationId_example" // string | Integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.GetOrganizationSaseIntegration(context.Background(), organizationId, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.GetOrganizationSaseIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseIntegration`: InlineResponse20128
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.GetOrganizationSaseIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**integrationId** | **string** | Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20128**](InlineResponse20128.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseNetworksEligible

> InlineResponse200402 GetOrganizationSaseNetworksEligible(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).Execute()

List of MX networks or templates that can be enrolled into Secure Access



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 5. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    search := "search_example" // string | If provided, filters results by network name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.GetOrganizationSaseNetworksEligible(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.GetOrganizationSaseNetworksEligible``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseNetworksEligible`: InlineResponse200402
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.GetOrganizationSaseNetworksEligible`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseNetworksEligibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 5. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **search** | **string** | If provided, filters results by network name | 

### Return type

[**InlineResponse200402**](InlineResponse200402.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseRegions

> InlineResponse200403 GetOrganizationSaseRegions(ctx, organizationId).Execute()

List regions



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
    resp, r, err := apiClient.SaseApi.GetOrganizationSaseRegions(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.GetOrganizationSaseRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseRegions`: InlineResponse200403
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.GetOrganizationSaseRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200403**](InlineResponse200403.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSites

> InlineResponse200404 GetOrganizationSaseSites(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).Status(status).SiteId(siteId).Execute()

List of enrolled SASE sites in this organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    search := "search_example" // string | If provided, filters results by search string (optional)
    status := "status_example" // string | If provided, filters results by site status label (e.g., 'good') (optional)
    siteId := "siteId_example" // string | If provided, returns only the site matching this ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.GetOrganizationSaseSites(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).Status(status).SiteId(siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.GetOrganizationSaseSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseSites`: InlineResponse200404
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.GetOrganizationSaseSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **search** | **string** | If provided, filters results by search string | 
 **status** | **string** | If provided, filters results by site status label (e.g., &#39;good&#39;) | 
 **siteId** | **string** | If provided, returns only the site matching this ID | 

### Return type

[**InlineResponse200404**](InlineResponse200404.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSitesConnectivityHistoryBySite

> InlineResponse200405 GetOrganizationSaseSitesConnectivityHistoryBySite(ctx, organizationId).SiteIds(siteIds).Timespan(timespan).Execute()

Get the connectivity history of SASE sites in this organization



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
    siteIds := []string{"Inner_example"} // []string | Array of site IDs to fetch connectivity status data (maximum 100 values) (optional)
    timespan := "timespan_example" // string | Timespan for the status data (e.g., '-2hours') (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.GetOrganizationSaseSitesConnectivityHistoryBySite(context.Background(), organizationId).SiteIds(siteIds).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.GetOrganizationSaseSitesConnectivityHistoryBySite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseSitesConnectivityHistoryBySite`: InlineResponse200405
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.GetOrganizationSaseSitesConnectivityHistoryBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseSitesConnectivityHistoryBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteIds** | **[]string** | Array of site IDs to fetch connectivity status data (maximum 100 values) | 
 **timespan** | **string** | Timespan for the status data (e.g., &#39;-2hours&#39;) | 

### Return type

[**InlineResponse200405**](InlineResponse200405.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSitesConnectivityOverview

> InlineResponse200406 GetOrganizationSaseSitesConnectivityOverview(ctx, organizationId).Execute()

List high-level SASE site statuses (healthy, degraded, offline)



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
    resp, r, err := apiClient.SaseApi.GetOrganizationSaseSitesConnectivityOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.GetOrganizationSaseSitesConnectivityOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseSitesConnectivityOverview`: InlineResponse200406
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.GetOrganizationSaseSitesConnectivityOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseSitesConnectivityOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200406**](InlineResponse200406.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSaseSite

> InlineResponse200407 UpdateOrganizationSaseSite(ctx, organizationId, siteId).UpdateOrganizationSaseSite(updateOrganizationSaseSite).Execute()

Update the configuration for a site



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
    siteId := "siteId_example" // string | Site ID
    updateOrganizationSaseSite := *openapiclient.NewInlineObject338("SiteId_example") // InlineObject338 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.UpdateOrganizationSaseSite(context.Background(), organizationId, siteId).UpdateOrganizationSaseSite(updateOrganizationSaseSite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.UpdateOrganizationSaseSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSaseSite`: InlineResponse200407
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.UpdateOrganizationSaseSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSaseSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSaseSite** | [**InlineObject338**](InlineObject338.md) |  | 

### Return type

[**InlineResponse200407**](InlineResponse200407.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

