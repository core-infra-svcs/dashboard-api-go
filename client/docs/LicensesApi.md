# \LicensesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrganizationLicensesSeats**](LicensesApi.md#AssignOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/assignSeats | Assign SM seats to a network
[**GetOrganizationLicense**](LicensesApi.md#GetOrganizationLicense) | **Get** /organizations/{organizationId}/licenses/{licenseId} | Display a license
[**GetOrganizationLicenses**](LicensesApi.md#GetOrganizationLicenses) | **Get** /organizations/{organizationId}/licenses | List the licenses for an organization
[**GetOrganizationLicensesOverview**](LicensesApi.md#GetOrganizationLicensesOverview) | **Get** /organizations/{organizationId}/licenses/overview | Return an overview of the license state for an organization
[**GetOrganizationLicensingCotermLicenses**](LicensesApi.md#GetOrganizationLicensingCotermLicenses) | **Get** /organizations/{organizationId}/licensing/coterm/licenses | List the licenses in a coterm organization
[**MoveOrganizationLicenses**](LicensesApi.md#MoveOrganizationLicenses) | **Post** /organizations/{organizationId}/licenses/move | Move licenses to another organization
[**MoveOrganizationLicensesSeats**](LicensesApi.md#MoveOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/moveSeats | Move SM seats to another organization
[**MoveOrganizationLicensingCotermLicenses**](LicensesApi.md#MoveOrganizationLicensingCotermLicenses) | **Post** /organizations/{organizationId}/licensing/coterm/licenses/move | Moves a license to a different organization (coterm only)
[**RenewOrganizationLicensesSeats**](LicensesApi.md#RenewOrganizationLicensesSeats) | **Post** /organizations/{organizationId}/licenses/renewSeats | Renew SM seats of a license
[**UpdateOrganizationLicense**](LicensesApi.md#UpdateOrganizationLicense) | **Put** /organizations/{organizationId}/licenses/{licenseId} | Update a license



## AssignOrganizationLicensesSeats

> InlineResponse200165 AssignOrganizationLicensesSeats(ctx, organizationId).AssignOrganizationLicensesSeats(assignOrganizationLicensesSeats).Execute()

Assign SM seats to a network



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
    assignOrganizationLicensesSeats := *openapiclient.NewInlineObject228("LicenseId_example", "NetworkId_example", int32(123)) // InlineObject228 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.AssignOrganizationLicensesSeats(context.Background(), organizationId).AssignOrganizationLicensesSeats(assignOrganizationLicensesSeats).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.AssignOrganizationLicensesSeats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOrganizationLicensesSeats`: InlineResponse200165
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.AssignOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationLicensesSeats** | [**InlineObject228**](InlineObject228.md) |  | 

### Return type

[**InlineResponse200165**](InlineResponse200165.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicense

> InlineResponse200164 GetOrganizationLicense(ctx, organizationId, licenseId).Execute()

Display a license



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
    licenseId := "licenseId_example" // string | License ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.GetOrganizationLicense(context.Background(), organizationId, licenseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.GetOrganizationLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicense`: InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.GetOrganizationLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**licenseId** | **string** | License ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicenses

> []InlineResponse200164 GetOrganizationLicenses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).DeviceSerial(deviceSerial).NetworkId(networkId).State(state).Execute()

List the licenses for an organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter the licenses to those assigned to a particular device. Returned in the same order that they are queued to the device. (optional)
    networkId := "networkId_example" // string | Filter the licenses to those assigned in a particular network (optional)
    state := "state_example" // string | Filter the licenses to those in a particular state. Can be one of 'active', 'expired', 'expiring', 'recentlyQueued', 'unused' or 'unusedActive' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.GetOrganizationLicenses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).DeviceSerial(deviceSerial).NetworkId(networkId).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.GetOrganizationLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicenses`: []InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.GetOrganizationLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **deviceSerial** | **string** | Filter the licenses to those assigned to a particular device. Returned in the same order that they are queued to the device. | 
 **networkId** | **string** | Filter the licenses to those assigned in a particular network | 
 **state** | **string** | Filter the licenses to those in a particular state. Can be one of &#39;active&#39;, &#39;expired&#39;, &#39;expiring&#39;, &#39;recentlyQueued&#39;, &#39;unused&#39; or &#39;unusedActive&#39; | 

### Return type

[**[]InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicensesOverview

> map[string]interface{} GetOrganizationLicensesOverview(ctx, organizationId).Execute()

Return an overview of the license state for an organization



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
    resp, r, err := apiClient.LicensesApi.GetOrganizationLicensesOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.GetOrganizationLicensesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicensesOverview`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.GetOrganizationLicensesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicensingCotermLicenses

> []InlineResponse200168 GetOrganizationLicensingCotermLicenses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Invalidated(invalidated).Expired(expired).Execute()

List the licenses in a coterm organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    invalidated := true // bool | Filter for licenses that are invalidated (optional)
    expired := true // bool | Filter for licenses that are expired (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.GetOrganizationLicensingCotermLicenses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Invalidated(invalidated).Expired(expired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.GetOrganizationLicensingCotermLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicensingCotermLicenses`: []InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.GetOrganizationLicensingCotermLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensingCotermLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **invalidated** | **bool** | Filter for licenses that are invalidated | 
 **expired** | **bool** | Filter for licenses that are expired | 

### Return type

[**[]InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicenses

> InlineResponse200166 MoveOrganizationLicenses(ctx, organizationId).MoveOrganizationLicenses(moveOrganizationLicenses).Execute()

Move licenses to another organization



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
    moveOrganizationLicenses := *openapiclient.NewInlineObject229("DestOrganizationId_example", []string{"LicenseIds_example"}) // InlineObject229 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.MoveOrganizationLicenses(context.Background(), organizationId).MoveOrganizationLicenses(moveOrganizationLicenses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.MoveOrganizationLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveOrganizationLicenses`: InlineResponse200166
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.MoveOrganizationLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicenses** | [**InlineObject229**](InlineObject229.md) |  | 

### Return type

[**InlineResponse200166**](InlineResponse200166.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicensesSeats

> InlineResponse200167 MoveOrganizationLicensesSeats(ctx, organizationId).MoveOrganizationLicensesSeats(moveOrganizationLicensesSeats).Execute()

Move SM seats to another organization



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
    moveOrganizationLicensesSeats := *openapiclient.NewInlineObject230("DestOrganizationId_example", "LicenseId_example", int32(123)) // InlineObject230 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.MoveOrganizationLicensesSeats(context.Background(), organizationId).MoveOrganizationLicensesSeats(moveOrganizationLicensesSeats).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.MoveOrganizationLicensesSeats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveOrganizationLicensesSeats`: InlineResponse200167
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.MoveOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicensesSeats** | [**InlineObject230**](InlineObject230.md) |  | 

### Return type

[**InlineResponse200167**](InlineResponse200167.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicensingCotermLicenses

> InlineResponse200169 MoveOrganizationLicensingCotermLicenses(ctx, organizationId).MoveOrganizationLicensingCotermLicenses(moveOrganizationLicensingCotermLicenses).Execute()

Moves a license to a different organization (coterm only)



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
    moveOrganizationLicensingCotermLicenses := *openapiclient.NewInlineObject233(*openapiclient.NewOrganizationsOrganizationIdLicensingCotermLicensesMoveDestination(), []openapiclient.OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses{*openapiclient.NewOrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses("Key_example", []openapiclient.OrganizationsOrganizationIdLicensingCotermLicensesMoveCounts{*openapiclient.NewOrganizationsOrganizationIdLicensingCotermLicensesMoveCounts("Model_example", int32(123))})}) // InlineObject233 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.MoveOrganizationLicensingCotermLicenses(context.Background(), organizationId).MoveOrganizationLicensingCotermLicenses(moveOrganizationLicensingCotermLicenses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.MoveOrganizationLicensingCotermLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveOrganizationLicensingCotermLicenses`: InlineResponse200169
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.MoveOrganizationLicensingCotermLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensingCotermLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicensingCotermLicenses** | [**InlineObject233**](InlineObject233.md) |  | 

### Return type

[**InlineResponse200169**](InlineResponse200169.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewOrganizationLicensesSeats

> InlineResponse200165 RenewOrganizationLicensesSeats(ctx, organizationId).RenewOrganizationLicensesSeats(renewOrganizationLicensesSeats).Execute()

Renew SM seats of a license



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
    renewOrganizationLicensesSeats := *openapiclient.NewInlineObject231("LicenseIdToRenew_example", "UnusedLicenseId_example") // InlineObject231 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.RenewOrganizationLicensesSeats(context.Background(), organizationId).RenewOrganizationLicensesSeats(renewOrganizationLicensesSeats).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.RenewOrganizationLicensesSeats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenewOrganizationLicensesSeats`: InlineResponse200165
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.RenewOrganizationLicensesSeats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewOrganizationLicensesSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renewOrganizationLicensesSeats** | [**InlineObject231**](InlineObject231.md) |  | 

### Return type

[**InlineResponse200165**](InlineResponse200165.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationLicense

> InlineResponse200164 UpdateOrganizationLicense(ctx, organizationId, licenseId).UpdateOrganizationLicense(updateOrganizationLicense).Execute()

Update a license



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
    licenseId := "licenseId_example" // string | License ID
    updateOrganizationLicense := *openapiclient.NewInlineObject232() // InlineObject232 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.UpdateOrganizationLicense(context.Background(), organizationId, licenseId).UpdateOrganizationLicense(updateOrganizationLicense).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.UpdateOrganizationLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationLicense`: InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.UpdateOrganizationLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**licenseId** | **string** | License ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationLicense** | [**InlineObject232**](InlineObject232.md) |  | 

### Return type

[**InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

