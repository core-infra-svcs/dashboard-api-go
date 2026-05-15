# \GroupApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups**](GroupApi.md#AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups) | **Post** /organizations/{organizationId}/policies/global/group/policies/adaptivePolicyGroups/assign | Assign adaptive policy groups to a policy
[**AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans**](GroupApi.md#AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans) | **Post** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assign | Assign VLANs to a policy
[**CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](GroupApi.md#CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Post** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments | Create an Organization-Wide Policy Ruleset Assignment
[**CreateOrganizationPoliciesGlobalGroupPolicy**](GroupApi.md#CreateOrganizationPoliciesGlobalGroupPolicy) | **Post** /organizations/{organizationId}/policies/global/group/policies | Create an Organization-Wide Policy
[**DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](GroupApi.md#DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Delete** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments/{assignmentId} | Delete an Organization-Wide Policy Ruleset Assignment
[**DeleteOrganizationPoliciesGlobalGroupPolicy**](GroupApi.md#DeleteOrganizationPoliciesGlobalGroupPolicy) | **Delete** /organizations/{organizationId}/policies/global/group/policies/{policyId} | Delete an Organization-Wide Policy
[**GetOrganizationPoliciesGlobalGroupPolicies**](GroupApi.md#GetOrganizationPoliciesGlobalGroupPolicies) | **Get** /organizations/{organizationId}/policies/global/group/policies | List Organization-Wide Policies
[**GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments**](GroupApi.md#GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/adaptivePolicyGroups/assignments | List adaptive policy group assignments
[**GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments**](GroupApi.md#GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assignments | List appliance VLAN policy assignments
[**GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan**](GroupApi.md#GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan) | **Get** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assignments/byVlan | List policies by appliance VLANs
[**GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments**](GroupApi.md#GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments | List Organization-Wide Policy Ruleset Assignments
[**RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups**](GroupApi.md#RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups) | **Post** /organizations/{organizationId}/policies/global/group/policies/adaptivePolicyGroups/remove | Remove adaptive policy groups from a policy
[**RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans**](GroupApi.md#RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans) | **Post** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/remove | Remove VLANs from a policy
[**UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](GroupApi.md#UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Put** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments/{assignmentId} | Update an Organization-Wide Policy Ruleset Assignment
[**UpdateOrganizationPoliciesGlobalGroupPolicy**](GroupApi.md#UpdateOrganizationPoliciesGlobalGroupPolicy) | **Put** /organizations/{organizationId}/policies/global/group/policies/{policyId} | Update an Organization-Wide Policy



## AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups

> InlineResponse200376 AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(ctx, organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()

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
    assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups := *openapiclient.NewInlineObject305(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignPolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignAdaptivePolicyGroups{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignAdaptivePolicyGroups("Id_example")}) // InlineObject305 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(context.Background(), organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: InlineResponse200376
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.AssignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: %v\n", resp)
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

 **assignOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups** | [**InlineObject305**](InlineObject305.md) |  | 

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


## AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans

> InlineResponse200376 AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(ctx, organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()

Assign VLANs to a policy



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
    assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans := *openapiclient.NewInlineObject307(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignPolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans("InterfaceId_example")}) // InlineObject307 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(context.Background(), organizationId).AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: InlineResponse200376
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.AssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrganizationPoliciesGlobalGroupPoliciesApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignOrganizationPoliciesGlobalGroupPoliciesApplianceVlans** | [**InlineObject307**](InlineObject307.md) |  | 

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


## CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> InlineResponse200380Items CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId).CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()

Create an Organization-Wide Policy Ruleset Assignment



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
    createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment := *openapiclient.NewInlineObject309("RulesetId_example", "PolicyId_example") // InlineObject309 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId).CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: InlineResponse200380Items
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment** | [**InlineObject309**](InlineObject309.md) |  | 

### Return type

[**InlineResponse200380Items**](InlineResponse200380Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPoliciesGlobalGroupPolicy

> InlineResponse200375Items CreateOrganizationPoliciesGlobalGroupPolicy(ctx, organizationId).CreateOrganizationPoliciesGlobalGroupPolicy(createOrganizationPoliciesGlobalGroupPolicy).Execute()

Create an Organization-Wide Policy



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
    createOrganizationPoliciesGlobalGroupPolicy := *openapiclient.NewInlineObject304("Name_example") // InlineObject304 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.CreateOrganizationPoliciesGlobalGroupPolicy(context.Background(), organizationId).CreateOrganizationPoliciesGlobalGroupPolicy(createOrganizationPoliciesGlobalGroupPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.CreateOrganizationPoliciesGlobalGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalGroupPolicy`: InlineResponse200375Items
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.CreateOrganizationPoliciesGlobalGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPoliciesGlobalGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPoliciesGlobalGroupPolicy** | [**InlineObject304**](InlineObject304.md) |  | 

### Return type

[**InlineResponse200375Items**](InlineResponse200375Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId, assignmentId).Execute()

Delete an Organization-Wide Policy Ruleset Assignment



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
    assignmentId := "assignmentId_example" // string | Assignment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId, assignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**assignmentId** | **string** | Assignment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentRequest struct via the builder pattern


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


## DeleteOrganizationPoliciesGlobalGroupPolicy

> DeleteOrganizationPoliciesGlobalGroupPolicy(ctx, organizationId, policyId).Execute()

Delete an Organization-Wide Policy



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
    policyId := "policyId_example" // string | Policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.DeleteOrganizationPoliciesGlobalGroupPolicy(context.Background(), organizationId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteOrganizationPoliciesGlobalGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPoliciesGlobalGroupPolicyRequest struct via the builder pattern


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


## GetOrganizationPoliciesGlobalGroupPolicies

> InlineResponse200375 GetOrganizationPoliciesGlobalGroupPolicies(ctx, organizationId).Name(name).PolicyIds(policyIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Organization-Wide Policies



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
    name := "name_example" // string | Filter policies by name (partial match, case-insensitive). If multiple instances are provided, only the last one is used. (optional)
    policyIds := []string{"Inner_example"} // []string | Filter policies by IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.GetOrganizationPoliciesGlobalGroupPolicies(context.Background(), organizationId).Name(name).PolicyIds(policyIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetOrganizationPoliciesGlobalGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPolicies`: InlineResponse200375
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetOrganizationPoliciesGlobalGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Filter policies by name (partial match, case-insensitive). If multiple instances are provided, only the last one is used. | 
 **policyIds** | **[]string** | Filter policies by IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200375**](InlineResponse200375.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments

> InlineResponse200377 GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments(ctx, organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).AdaptivePolicyGroupIds(adaptivePolicyGroupIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments(context.Background(), organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).AdaptivePolicyGroupIds(adaptivePolicyGroupIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments`: InlineResponse200377
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments`: %v\n", resp)
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

[**InlineResponse200377**](InlineResponse200377.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments

> InlineResponse200378 GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments(ctx, organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List appliance VLAN policy assignments



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
    interfaceIds := []string{"Inner_example"} // []string | Filter assignments by interface IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments(context.Background(), organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments`: InlineResponse200378
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignmentIds** | **[]string** | Filter assignments by assignment IDs | 
 **policyIds** | **[]string** | Filter assignments by policy IDs | 
 **interfaceIds** | **[]string** | Filter assignments by interface IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200378**](InlineResponse200378.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan

> InlineResponse200379 GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan(ctx, organizationId).Search(search).VlanIds(vlanIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List policies by appliance VLANs



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
    search := "search_example" // string | Search term for filtering policies (optional)
    vlanIds := []string{"Inner_example"} // []string | Filter by VLAN IDs (optional)
    interfaceIds := []string{"Inner_example"} // []string | Filter by interface IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan(context.Background(), organizationId).Search(search).VlanIds(vlanIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan`: InlineResponse200379
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Search term for filtering policies | 
 **vlanIds** | **[]string** | Filter by VLAN IDs | 
 **interfaceIds** | **[]string** | Filter by interface IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200379**](InlineResponse200379.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments

> InlineResponse200380 GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments(ctx, organizationId).RulesetIds(rulesetIds).PolicyIds(policyIds).AssignmentIds(assignmentIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Organization-Wide Policy Ruleset Assignments



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
    rulesetIds := []string{"Inner_example"} // []string | Filter assignments by ruleset IDs (optional)
    policyIds := []string{"Inner_example"} // []string | Filter assignments by policy IDs (optional)
    assignmentIds := []string{"Inner_example"} // []string | Filter assignments by assignment IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments(context.Background(), organizationId).RulesetIds(rulesetIds).PolicyIds(policyIds).AssignmentIds(assignmentIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments`: InlineResponse200380
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rulesetIds** | **[]string** | Filter assignments by ruleset IDs | 
 **policyIds** | **[]string** | Filter assignments by policy IDs | 
 **assignmentIds** | **[]string** | Filter assignments by assignment IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200380**](InlineResponse200380.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups

> InlineResponse200376 RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(ctx, organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()

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
    removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups := *openapiclient.NewInlineObject306(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsRemovePolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignAdaptivePolicyGroups{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignAdaptivePolicyGroups("Id_example")}) // InlineObject306 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(context.Background(), organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups(removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: InlineResponse200376
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.RemoveOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups`: %v\n", resp)
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

 **removeOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroups** | [**InlineObject306**](InlineObject306.md) |  | 

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


## RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans

> InlineResponse200376 RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(ctx, organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()

Remove VLANs from a policy



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
    removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans := *openapiclient.NewInlineObject308(*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansRemovePolicy("Id_example"), []openapiclient.OrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans{*openapiclient.NewOrganizationsOrganizationIdPoliciesGlobalGroupPoliciesApplianceVlansAssignVlans("InterfaceId_example")}) // InlineObject308 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(context.Background(), organizationId).RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans(removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: InlineResponse200376
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.RemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationPoliciesGlobalGroupPoliciesApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeOrganizationPoliciesGlobalGroupPoliciesApplianceVlans** | [**InlineObject308**](InlineObject308.md) |  | 

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


## UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> InlineResponse200380Items UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId, assignmentId).UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()

Update an Organization-Wide Policy Ruleset Assignment



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
    assignmentId := "assignmentId_example" // string | Assignment ID
    updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment := *openapiclient.NewInlineObject310() // InlineObject310 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId, assignmentId).UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: InlineResponse200380Items
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**assignmentId** | **string** | Assignment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment** | [**InlineObject310**](InlineObject310.md) |  | 

### Return type

[**InlineResponse200380Items**](InlineResponse200380Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPoliciesGlobalGroupPolicy

> InlineResponse200375Items UpdateOrganizationPoliciesGlobalGroupPolicy(ctx, organizationId, policyId).UpdateOrganizationPoliciesGlobalGroupPolicy(updateOrganizationPoliciesGlobalGroupPolicy).Execute()

Update an Organization-Wide Policy



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
    policyId := "policyId_example" // string | Policy ID
    updateOrganizationPoliciesGlobalGroupPolicy := *openapiclient.NewInlineObject311() // InlineObject311 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupApi.UpdateOrganizationPoliciesGlobalGroupPolicy(context.Background(), organizationId, policyId).UpdateOrganizationPoliciesGlobalGroupPolicy(updateOrganizationPoliciesGlobalGroupPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateOrganizationPoliciesGlobalGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalGroupPolicy`: InlineResponse200375Items
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.UpdateOrganizationPoliciesGlobalGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyId** | **string** | Policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPoliciesGlobalGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPoliciesGlobalGroupPolicy** | [**InlineObject311**](InlineObject311.md) |  | 

### Return type

[**InlineResponse200375Items**](InlineResponse200375Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

