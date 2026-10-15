# \ByGroupApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationNetworksGroupsOverviewByGroup**](ByGroupApi.md#GetOrganizationNetworksGroupsOverviewByGroup) | **Get** /organizations/{organizationId}/networks/groups/overview/byGroup | List the client and status overview information for the network groups in an organization



## GetOrganizationNetworksGroupsOverviewByGroup

> InlineResponse200416 GetOrganizationNetworksGroupsOverviewByGroup(ctx, organizationId).SortBy(sortBy).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the client and status overview information for the network groups in an organization



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
    sortBy := "sortBy_example" // string | Field by which to sort the results (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ByGroupApi.GetOrganizationNetworksGroupsOverviewByGroup(context.Background(), organizationId).SortBy(sortBy).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ByGroupApi.GetOrganizationNetworksGroupsOverviewByGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationNetworksGroupsOverviewByGroup`: InlineResponse200416
    fmt.Fprintf(os.Stdout, "Response from `ByGroupApi.GetOrganizationNetworksGroupsOverviewByGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationNetworksGroupsOverviewByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** | Field by which to sort the results | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200416**](InlineResponse200416.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

