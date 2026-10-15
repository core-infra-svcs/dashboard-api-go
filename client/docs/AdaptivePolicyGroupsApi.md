# \AdaptivePolicyGroupsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups**](AdaptivePolicyGroupsApi.md#AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups) | **Post** /organizations/{organizationId}/policies/global/group/policies/adaptivePolicyGroups/assign | Assign adaptive policy groups to a policy
[**GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments**](AdaptivePolicyGroupsApi.md#GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/adaptivePolicyGroups/assignments | List adaptive policy group assignments
[**RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups**](AdaptivePolicyGroupsApi.md#RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups) | **Post** /organizations/{organizationId}/policies/global/group/policies/adaptivePolicyGroups/remove | Remove adaptive policy groups from a policy



## AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups

> InlineResponse200424 AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(ctx, organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()

Assign adaptive policy groups to a policy



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
    assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups := *openapiclient.NewInlineObject339(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignPolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaAdaptivePolicyGroups{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaAdaptivePolicyGroups("Id_example")}) // InlineObject339 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyGroupsApi.AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(context.Background(), organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyGroupsApi.AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: InlineResponse200424
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyGroupsApi.AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups** | [**InlineObject339**](InlineObject339.md) |  | 

### Return type

[**InlineResponse200424**](InlineResponse200424.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments

> InlineResponse200425 GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments(ctx, organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).AdaptivePolicyGroupIds(adaptivePolicyGroupIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List adaptive policy group assignments



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
    assignmentIds := []string{"Inner_example"} // []string | Filter assignments by assignment IDs (optional)
    policyIds := []string{"Inner_example"} // []string | Filter assignments by policy IDs (optional)
    adaptivePolicyGroupIds := []string{"Inner_example"} // []string | Filter assignments by adaptive policy group IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyGroupsApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments(context.Background(), organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).AdaptivePolicyGroupIds(adaptivePolicyGroupIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyGroupsApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments`: InlineResponse200425
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyGroupsApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignmentIds** | **[]string** | Filter assignments by assignment IDs | 
 **policyIds** | **[]string** | Filter assignments by policy IDs | 
 **adaptivePolicyGroupIds** | **[]string** | Filter assignments by adaptive policy group IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200425**](InlineResponse200425.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups

> InlineResponse200424 RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(ctx, organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()

Remove adaptive policy groups from a policy



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
    removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups := *openapiclient.NewInlineObject340(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsRemovePolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaAdaptivePolicyGroups{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalFirewallRulesetsRulesDestinationsCriteriaAdaptivePolicyGroups("Id_example")}) // InlineObject340 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyGroupsApi.RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(context.Background(), organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyGroupsApi.RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: InlineResponse200424
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyGroupsApi.RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups** | [**InlineObject340**](InlineObject340.md) |  | 

### Return type

[**InlineResponse200424**](InlineResponse200424.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

