# \AllowlistApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry**](AllowlistApi.md#CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry) | **Post** /organizations/{organizationId}/wireless/ssids/firewall/isolation/allowlist/entries | Create isolation allow list MAC entry for this organization
[**DeleteOrganizationWirelessSsidsFirewallIsolationAllowlistEntry**](AllowlistApi.md#DeleteOrganizationWirelessSsidsFirewallIsolationAllowlistEntry) | **Delete** /organizations/{organizationId}/wireless/ssids/firewall/isolation/allowlist/entries/{entryId} | Destroy isolation allow list MAC entry for this organization
[**GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries**](AllowlistApi.md#GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries) | **Get** /organizations/{organizationId}/wireless/ssids/firewall/isolation/allowlist/entries | List the L2 isolation allow list MAC entry in an organization
[**UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry**](AllowlistApi.md#UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry) | **Put** /organizations/{organizationId}/wireless/ssids/firewall/isolation/allowlist/entries/{entryId} | Update isolation allow list MAC entry info



## CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry

> InlineResponse200367Items CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(ctx, organizationId).CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(createOrganizationWirelessSsidsFirewallIsolationAllowlistEntry).Execute()

Create isolation allow list MAC entry for this organization



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
    createOrganizationWirelessSsidsFirewallIsolationAllowlistEntry := *openapiclient.NewInlineObject305(*openapiclient.NewOrganizationsOrganizationIdWirelessSsidsFirewallIsolationAllowlistEntriesClient(), *openapiclient.NewOrganizationsOrganizationIdWirelessSsidsFirewallIsolationAllowlistEntriesSsid(), *openapiclient.NewOrganizationsOrganizationIdWirelessSsidsFirewallIsolationAllowlistEntriesNetwork()) // InlineObject305 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowlistApi.CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(context.Background(), organizationId).CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(createOrganizationWirelessSsidsFirewallIsolationAllowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowlistApi.CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry`: InlineResponse200367Items
    fmt.Fprintf(os.Stdout, "Response from `AllowlistApi.CreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationWirelessSsidsFirewallIsolationAllowlistEntry** | [**InlineObject305**](InlineObject305.md) |  | 

### Return type

[**InlineResponse200367Items**](InlineResponse200367Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationWirelessSsidsFirewallIsolationAllowlistEntry

> DeleteOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(ctx, organizationId, entryId).Execute()

Destroy isolation allow list MAC entry for this organization



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
    entryId := "entryId_example" // string | Entry ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowlistApi.DeleteOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(context.Background(), organizationId, entryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowlistApi.DeleteOrganizationWirelessSsidsFirewallIsolationAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**entryId** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationWirelessSsidsFirewallIsolationAllowlistEntryRequest struct via the builder pattern


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


## GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries

> InlineResponse200367 GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Ssids(ssids).Execute()

List the L2 isolation allow list MAC entry in an organization



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
    networkIds := []string{"Inner_example"} // []string | networkIds array to filter out results (optional)
    ssids := []int32{int32(123)} // []int32 | ssids number array to filter out results (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowlistApi.GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Ssids(ssids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowlistApi.GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `AllowlistApi.GetOrganizationWirelessSsidsFirewallIsolationAllowlistEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessSsidsFirewallIsolationAllowlistEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | networkIds array to filter out results | 
 **ssids** | **[]int32** | ssids number array to filter out results | 

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


## UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry

> InlineResponse200367Items UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(ctx, organizationId, entryId).UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(updateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry).Execute()

Update isolation allow list MAC entry info



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
    entryId := "entryId_example" // string | Entry ID
    updateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry := *openapiclient.NewInlineObject306() // InlineObject306 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AllowlistApi.UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(context.Background(), organizationId, entryId).UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry(updateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AllowlistApi.UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry`: InlineResponse200367Items
    fmt.Fprintf(os.Stdout, "Response from `AllowlistApi.UpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**entryId** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry** | [**InlineObject306**](InlineObject306.md) |  | 

### Return type

[**InlineResponse200367Items**](InlineResponse200367Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

