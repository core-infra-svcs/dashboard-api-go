# \SaseApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachOrganizationSaseSites**](SaseApi.md#AttachOrganizationSaseSites) | **Post** /organizations/{organizationId}/sase/sites/attach | Attach sites in this organization to Secure Access
[**BatchOrganizationSaseConnectorsCreate**](SaseApi.md#BatchOrganizationSaseConnectorsCreate) | **Post** /organizations/{organizationId}/sase/connectors/batchCreate | Deploy SSE Connectors for specified regions
[**BatchOrganizationSaseConnectorsDelete**](SaseApi.md#BatchOrganizationSaseConnectorsDelete) | **Post** /organizations/{organizationId}/sase/connectors/batchDelete | Delete SSE Connectors by ID
[**DetachOrganizationSaseSites**](SaseApi.md#DetachOrganizationSaseSites) | **Delete** /organizations/{organizationId}/sase/sites/detach | Detach sites in this organization from Secure Access
[**GetOrganizationSaseConnectors**](SaseApi.md#GetOrganizationSaseConnectors) | **Get** /organizations/{organizationId}/sase/connectors | List SSE Connectors for an organization
[**GetOrganizationSaseNetworksEligible**](SaseApi.md#GetOrganizationSaseNetworksEligible) | **Get** /organizations/{organizationId}/sase/networks/eligible | List of MX networks or templates that can be enrolled into Secure Access
[**GetOrganizationSaseRegions**](SaseApi.md#GetOrganizationSaseRegions) | **Get** /organizations/{organizationId}/sase/regions | List regions
[**GetOrganizationSaseSites**](SaseApi.md#GetOrganizationSaseSites) | **Get** /organizations/{organizationId}/sase/sites | List of enrolled SASE sites in this organization
[**GetOrganizationSaseSitesConnectivityHistoryBySite**](SaseApi.md#GetOrganizationSaseSitesConnectivityHistoryBySite) | **Get** /organizations/{organizationId}/sase/sites/connectivity/history/bySite | Get the connectivity history of SASE sites in this organization
[**GetOrganizationSaseSitesConnectivityOverview**](SaseApi.md#GetOrganizationSaseSitesConnectivityOverview) | **Get** /organizations/{organizationId}/sase/sites/connectivity/overview | List high-level SASE site statuses (healthy, degraded, offline)
[**UpdateOrganizationSaseSite**](SaseApi.md#UpdateOrganizationSaseSite) | **Put** /organizations/{organizationId}/sase/sites/{siteId} | Update the configuration for a site



## AttachOrganizationSaseSites

> InlineResponse200245 AttachOrganizationSaseSites(ctx, organizationId).AttachOrganizationSaseSites(attachOrganizationSaseSites).Execute()

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
    attachOrganizationSaseSites := *openapiclient.NewInlineObject317() // InlineObject317 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.AttachOrganizationSaseSites(context.Background(), organizationId).AttachOrganizationSaseSites(attachOrganizationSaseSites).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.AttachOrganizationSaseSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachOrganizationSaseSites`: InlineResponse200245
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

 **attachOrganizationSaseSites** | [**InlineObject317**](InlineObject317.md) |  | 

### Return type

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchOrganizationSaseConnectorsCreate

> InlineResponse200245 BatchOrganizationSaseConnectorsCreate(ctx, organizationId).BatchOrganizationSaseConnectorsCreate(batchOrganizationSaseConnectorsCreate).Execute()

Deploy SSE Connectors for specified regions



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
    batchOrganizationSaseConnectorsCreate := *openapiclient.NewInlineObject315() // InlineObject315 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.BatchOrganizationSaseConnectorsCreate(context.Background(), organizationId).BatchOrganizationSaseConnectorsCreate(batchOrganizationSaseConnectorsCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.BatchOrganizationSaseConnectorsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationSaseConnectorsCreate`: InlineResponse200245
    fmt.Fprintf(os.Stdout, "Response from `SaseApi.BatchOrganizationSaseConnectorsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchOrganizationSaseConnectorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchOrganizationSaseConnectorsCreate** | [**InlineObject315**](InlineObject315.md) |  | 

### Return type

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchOrganizationSaseConnectorsDelete

> InlineResponse200245 BatchOrganizationSaseConnectorsDelete(ctx, organizationId).BatchOrganizationSaseConnectorsDelete(batchOrganizationSaseConnectorsDelete).Execute()

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
    batchOrganizationSaseConnectorsDelete := *openapiclient.NewInlineObject316() // InlineObject316 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.BatchOrganizationSaseConnectorsDelete(context.Background(), organizationId).BatchOrganizationSaseConnectorsDelete(batchOrganizationSaseConnectorsDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.BatchOrganizationSaseConnectorsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationSaseConnectorsDelete`: InlineResponse200245
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

 **batchOrganizationSaseConnectorsDelete** | [**InlineObject316**](InlineObject316.md) |  | 

### Return type

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachOrganizationSaseSites

> InlineResponse200245 DetachOrganizationSaseSites(ctx, organizationId).DetachOrganizationSaseSites(detachOrganizationSaseSites).Execute()

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
    detachOrganizationSaseSites := *openapiclient.NewInlineObject318() // InlineObject318 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.DetachOrganizationSaseSites(context.Background(), organizationId).DetachOrganizationSaseSites(detachOrganizationSaseSites).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.DetachOrganizationSaseSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachOrganizationSaseSites`: InlineResponse200245
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

 **detachOrganizationSaseSites** | [**InlineObject318**](InlineObject318.md) |  | 

### Return type

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseConnectors

> InlineResponse200362 GetOrganizationSaseConnectors(ctx, organizationId).Execute()

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
    // response from `GetOrganizationSaseConnectors`: InlineResponse200362
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

[**InlineResponse200362**](InlineResponse200362.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseNetworksEligible

> InlineResponse200363 GetOrganizationSaseNetworksEligible(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).Execute()

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
    // response from `GetOrganizationSaseNetworksEligible`: InlineResponse200363
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

[**InlineResponse200363**](InlineResponse200363.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseRegions

> InlineResponse200364 GetOrganizationSaseRegions(ctx, organizationId).Execute()

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
    // response from `GetOrganizationSaseRegions`: InlineResponse200364
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

[**InlineResponse200364**](InlineResponse200364.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSites

> InlineResponse200365 GetOrganizationSaseSites(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).Status(status).SiteId(siteId).Execute()

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
    // response from `GetOrganizationSaseSites`: InlineResponse200365
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

[**InlineResponse200365**](InlineResponse200365.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSitesConnectivityHistoryBySite

> InlineResponse200366 GetOrganizationSaseSitesConnectivityHistoryBySite(ctx, organizationId).SiteIds(siteIds).Timespan(timespan).Execute()

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
    // response from `GetOrganizationSaseSitesConnectivityHistoryBySite`: InlineResponse200366
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

[**InlineResponse200366**](InlineResponse200366.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSaseSitesConnectivityOverview

> InlineResponse200367 GetOrganizationSaseSitesConnectivityOverview(ctx, organizationId).Execute()

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
    // response from `GetOrganizationSaseSitesConnectivityOverview`: InlineResponse200367
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

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSaseSite

> InlineResponse200368 UpdateOrganizationSaseSite(ctx, organizationId, siteId).UpdateOrganizationSaseSite(updateOrganizationSaseSite).Execute()

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
    updateOrganizationSaseSite := *openapiclient.NewInlineObject319("SiteId_example") // InlineObject319 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SaseApi.UpdateOrganizationSaseSite(context.Background(), organizationId, siteId).UpdateOrganizationSaseSite(updateOrganizationSaseSite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaseApi.UpdateOrganizationSaseSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSaseSite`: InlineResponse200368
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


 **updateOrganizationSaseSite** | [**InlineObject319**](InlineObject319.md) |  | 

### Return type

[**InlineResponse200368**](InlineResponse200368.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

