# \SentryApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSmSentryPoliciesAssignmentsByNetwork**](SentryApi.md#GetOrganizationSmSentryPoliciesAssignmentsByNetwork) | **Get** /organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork | List the Sentry Policies for an organization ordered in ascending order of priority
[**UpdateOrganizationSmSentryPoliciesAssignments**](SentryApi.md#UpdateOrganizationSmSentryPoliciesAssignments) | **Put** /organizations/{organizationId}/sm/sentry/policies/assignments | Update an Organizations Sentry Policies using the provided list



## GetOrganizationSmSentryPoliciesAssignmentsByNetwork

> []InlineResponse200262 GetOrganizationSmSentryPoliciesAssignmentsByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List the Sentry Policies for an organization ordered in ascending order of priority



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter Sentry Policies by Network Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SentryApi.GetOrganizationSmSentryPoliciesAssignmentsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SentryApi.GetOrganizationSmSentryPoliciesAssignmentsByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: []InlineResponse200262
    fmt.Fprintf(os.Stdout, "Response from `SentryApi.GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmSentryPoliciesAssignmentsByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter Sentry Policies by Network Id | 

### Return type

[**[]InlineResponse200262**](InlineResponse200262.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSmSentryPoliciesAssignments

> InlineResponse200261 UpdateOrganizationSmSentryPoliciesAssignments(ctx, organizationId).UpdateOrganizationSmSentryPoliciesAssignments(updateOrganizationSmSentryPoliciesAssignments).Execute()

Update an Organizations Sentry Policies using the provided list



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
    updateOrganizationSmSentryPoliciesAssignments := *openapiclient.NewInlineObject252([]openapiclient.OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems{*openapiclient.NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems("NetworkId_example")}) // InlineObject252 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SentryApi.UpdateOrganizationSmSentryPoliciesAssignments(context.Background(), organizationId).UpdateOrganizationSmSentryPoliciesAssignments(updateOrganizationSmSentryPoliciesAssignments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SentryApi.UpdateOrganizationSmSentryPoliciesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSmSentryPoliciesAssignments`: InlineResponse200261
    fmt.Fprintf(os.Stdout, "Response from `SentryApi.UpdateOrganizationSmSentryPoliciesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSmSentryPoliciesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationSmSentryPoliciesAssignments** | [**InlineObject252**](InlineObject252.md) |  | 

### Return type

[**InlineResponse200261**](InlineResponse200261.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

