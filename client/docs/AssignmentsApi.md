# \AssignmentsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate**](AssignmentsApi.md#BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate) | **Post** /organizations/{organizationId}/devices/cellular/data/profiles/assignments/batchCreate | Assign devices to a Cellular Data Management Profile in batch
[**BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate**](AssignmentsApi.md#BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate) | **Post** /organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkCreate | Assign the local DNS profile to networks in the organization
[**BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete**](AssignmentsApi.md#BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete) | **Post** /organizations/{organizationId}/devices/cellular/data/profiles/assignments/bulkDelete | Unassign devices from a Cellular Data Management Profile in batch
[**CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete**](AssignmentsApi.md#CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete) | **Post** /organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkDelete | Unassign the local DNS profile to networks in the organization
[**CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate**](AssignmentsApi.md#CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate) | **Post** /organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkCreate | Assign the split DNS profile to networks in the organization
[**CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete**](AssignmentsApi.md#CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete) | **Post** /organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkDelete | Unassign the split DNS profile to networks in the organization
[**CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](AssignmentsApi.md#CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Post** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments | Create an Organization-Wide Policy Ruleset Assignment
[**CreateOrganizationWirelessSsidsProfilesAssignment**](AssignmentsApi.md#CreateOrganizationWirelessSsidsProfilesAssignment) | **Post** /organizations/{organizationId}/wireless/ssids/profiles/assignments | Assigns an SSID profile to an SSID in the organization
[**DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](AssignmentsApi.md#DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Delete** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments/{assignmentId} | Delete an Organization-Wide Policy Ruleset Assignment
[**DeleteOrganizationWirelessSsidsProfilesAssignments**](AssignmentsApi.md#DeleteOrganizationWirelessSsidsProfilesAssignments) | **Delete** /organizations/{organizationId}/wireless/ssids/profiles/assignments | Unassigns the SSID profile assigned to an SSID
[**GetNetworkVlanProfilesAssignmentsByDevice**](AssignmentsApi.md#GetNetworkVlanProfilesAssignmentsByDevice) | **Get** /networks/{networkId}/vlanProfiles/assignments/byDevice | Get the assigned VLAN Profiles for devices in a network
[**GetOrganizationApplianceDnsLocalProfilesAssignments**](AssignmentsApi.md#GetOrganizationApplianceDnsLocalProfilesAssignments) | **Get** /organizations/{organizationId}/appliance/dns/local/profiles/assignments | Fetch the local DNS profile assignments in the organization
[**GetOrganizationApplianceDnsSplitProfilesAssignments**](AssignmentsApi.md#GetOrganizationApplianceDnsSplitProfilesAssignments) | **Get** /organizations/{organizationId}/appliance/dns/split/profiles/assignments | Fetch the split DNS profile assignments in the organization
[**GetOrganizationDevicesCellularDataProfilesAssignments**](AssignmentsApi.md#GetOrganizationDevicesCellularDataProfilesAssignments) | **Get** /organizations/{organizationId}/devices/cellular/data/profiles/assignments | List Cellular Data Management Profile assignments in this organization
[**GetOrganizationPoliciesAssignmentsByClient**](AssignmentsApi.md#GetOrganizationPoliciesAssignmentsByClient) | **Get** /organizations/{organizationId}/policies/assignments/byClient | Get policies for all clients with policies
[**GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments**](AssignmentsApi.md#GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/adaptivePolicyGroups/assignments | List adaptive policy group assignments
[**GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments**](AssignmentsApi.md#GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assignments | List appliance VLAN policy assignments
[**GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan**](AssignmentsApi.md#GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan) | **Get** /organizations/{organizationId}/policies/global/group/policies/appliance/vlans/assignments/byVlan | List policies by appliance VLANs
[**GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments**](AssignmentsApi.md#GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments) | **Get** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments | List Organization-Wide Policy Ruleset Assignments
[**GetOrganizationSmSentryPoliciesAssignmentsByNetwork**](AssignmentsApi.md#GetOrganizationSmSentryPoliciesAssignmentsByNetwork) | **Get** /organizations/{organizationId}/sm/sentry/policies/assignments/byNetwork | List the Sentry Policies for an organization ordered in ascending order of priority
[**GetOrganizationWirelessRfProfilesAssignmentsByDevice**](AssignmentsApi.md#GetOrganizationWirelessRfProfilesAssignmentsByDevice) | **Get** /organizations/{organizationId}/wireless/rfProfiles/assignments/byDevice | List the RF profiles of an organization by device
[**GetOrganizationWirelessSsidsProfilesAssignments**](AssignmentsApi.md#GetOrganizationWirelessSsidsProfilesAssignments) | **Get** /organizations/{organizationId}/wireless/ssids/profiles/assignments | List the SSID profile assignments in an organization
[**GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork**](AssignmentsApi.md#GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork) | **Get** /organizations/{organizationId}/wireless/ssids/profiles/assignments/byNetwork | List the SSID profile assignments in an organization, grouped by network
[**ReassignNetworkVlanProfilesAssignments**](AssignmentsApi.md#ReassignNetworkVlanProfilesAssignments) | **Post** /networks/{networkId}/vlanProfiles/assignments/reassign | Update the assigned VLAN Profile for devices in a network
[**UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment**](AssignmentsApi.md#UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment) | **Put** /organizations/{organizationId}/policies/global/group/policies/firewall/rulesets/assignments/{assignmentId} | Update an Organization-Wide Policy Ruleset Assignment
[**UpdateOrganizationSmSentryPoliciesAssignments**](AssignmentsApi.md#UpdateOrganizationSmSentryPoliciesAssignments) | **Put** /organizations/{organizationId}/sm/sentry/policies/assignments | Update an Organizations Sentry Policies using the provided list



## BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate

> InlineResponse200347 BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(ctx, organizationId).BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(batchOrganizationDevicesCellularDataProfilesAssignmentsCreate).Execute()

Assign devices to a Cellular Data Management Profile in batch



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
    batchOrganizationDevicesCellularDataProfilesAssignmentsCreate := *openapiclient.NewInlineObject282([]openapiclient.OrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateItems{*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateItems(*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateProfile("Id_example"), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateDevice("Serial_example"))}) // InlineObject282 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(context.Background(), organizationId).BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(batchOrganizationDevicesCellularDataProfilesAssignmentsCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate`: InlineResponse200347
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchOrganizationDevicesCellularDataProfilesAssignmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchOrganizationDevicesCellularDataProfilesAssignmentsCreate** | [**InlineObject282**](InlineObject282.md) |  | 

### Return type

[**InlineResponse200347**](InlineResponse200347.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate

> InlineResponse200278 BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(ctx, organizationId).BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate).Execute()

Assign the local DNS profile to networks in the organization



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
    bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate := *openapiclient.NewInlineObject246([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems()}) // InlineObject246 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(context.Background(), organizationId).BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate(bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate`: InlineResponse200278
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.BulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate** | [**InlineObject246**](InlineObject246.md) |  | 

### Return type

[**InlineResponse200278**](InlineResponse200278.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete

> BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(ctx, organizationId).BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete).Execute()

Unassign devices from a Cellular Data Management Profile in batch



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
    bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete := *openapiclient.NewInlineObject283([]openapiclient.OrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteItems(*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteProfile("Id_example"), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteDevice("Serial_example"))}) // InlineObject283 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(context.Background(), organizationId).BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationDevicesCellularDataProfilesAssignmentsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete** | [**InlineObject283**](InlineObject283.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete

> InlineResponse200278 CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(ctx, organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete).Execute()

Unassign the local DNS profile to networks in the organization



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
    createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete := *openapiclient.NewInlineObject247([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems()}) // InlineObject247 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(context.Background(), organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete`: InlineResponse200278
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete** | [**InlineObject247**](InlineObject247.md) |  | 

### Return type

[**InlineResponse200278**](InlineResponse200278.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate

> InlineResponse200282 CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(ctx, organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate).Execute()

Assign the split DNS profile to networks in the organization



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
    createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate := *openapiclient.NewInlineObject252([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems()}) // InlineObject252 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(context.Background(), organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate`: InlineResponse200282
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate** | [**InlineObject252**](InlineObject252.md) |  | 

### Return type

[**InlineResponse200282**](InlineResponse200282.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete

> InlineResponse200282 CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(ctx, organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete).Execute()

Unassign the split DNS profile to networks in the organization



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
    createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete := *openapiclient.NewInlineObject253([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems()}) // InlineObject253 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(context.Background(), organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete`: InlineResponse200282
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete** | [**InlineObject253**](InlineObject253.md) |  | 

### Return type

[**InlineResponse200282**](InlineResponse200282.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> InlineResponse200410Items CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId).CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()

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
    createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment := *openapiclient.NewInlineObject329("RulesetId_example", "PolicyId_example") // InlineObject329 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId).CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: InlineResponse200410Items
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.CreateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: %v\n", resp)
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

 **createOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment** | [**InlineObject329**](InlineObject329.md) |  | 

### Return type

[**InlineResponse200410Items**](InlineResponse200410Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationWirelessSsidsProfilesAssignment

> InlineResponse200489Items CreateOrganizationWirelessSsidsProfilesAssignment(ctx, organizationId).CreateOrganizationWirelessSsidsProfilesAssignment(createOrganizationWirelessSsidsProfilesAssignment).Execute()

Assigns an SSID profile to an SSID in the organization



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
    createOrganizationWirelessSsidsProfilesAssignment := *openapiclient.NewInlineObject363(*openapiclient.NewOrganizationsOrganizationIdWirelessSsidsProfilesAssignmentsProfile("Id_example"), *openapiclient.NewOrganizationsOrganizationIdWirelessSsidsProfilesAssignmentsSsid()) // InlineObject363 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.CreateOrganizationWirelessSsidsProfilesAssignment(context.Background(), organizationId).CreateOrganizationWirelessSsidsProfilesAssignment(createOrganizationWirelessSsidsProfilesAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.CreateOrganizationWirelessSsidsProfilesAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessSsidsProfilesAssignment`: InlineResponse200489Items
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.CreateOrganizationWirelessSsidsProfilesAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessSsidsProfilesAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationWirelessSsidsProfilesAssignment** | [**InlineObject363**](InlineObject363.md) |  | 

### Return type

[**InlineResponse200489Items**](InlineResponse200489Items.md)

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
    resp, r, err := apiClient.AssignmentsApi.DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId, assignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.DeleteOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
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


## DeleteOrganizationWirelessSsidsProfilesAssignments

> DeleteOrganizationWirelessSsidsProfilesAssignments(ctx, organizationId).DeleteOrganizationWirelessSsidsProfilesAssignments(deleteOrganizationWirelessSsidsProfilesAssignments).Execute()

Unassigns the SSID profile assigned to an SSID



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
    deleteOrganizationWirelessSsidsProfilesAssignments := *openapiclient.NewInlineObject364(*openapiclient.NewOrganizationsOrganizationIdWirelessSsidsProfilesAssignmentsSsid1()) // InlineObject364 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.DeleteOrganizationWirelessSsidsProfilesAssignments(context.Background(), organizationId).DeleteOrganizationWirelessSsidsProfilesAssignments(deleteOrganizationWirelessSsidsProfilesAssignments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.DeleteOrganizationWirelessSsidsProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationWirelessSsidsProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteOrganizationWirelessSsidsProfilesAssignments** | [**InlineObject364**](InlineObject364.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkVlanProfilesAssignmentsByDevice

> []InlineResponse200216 GetNetworkVlanProfilesAssignmentsByDevice(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()

Get the assigned VLAN Profiles for devices in a network



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
    networkId := "networkId_example" // string | Network ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product types. (optional)
    stackIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by Switch Stack ids. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetNetworkVlanProfilesAssignmentsByDevice(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetNetworkVlanProfilesAssignmentsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkVlanProfilesAssignmentsByDevice`: []InlineResponse200216
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetNetworkVlanProfilesAssignmentsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfilesAssignmentsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **serials** | **[]string** | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product types. | 
 **stackIds** | **[]string** | Optional parameter to filter devices by Switch Stack ids. | 

### Return type

[**[]InlineResponse200216**](InlineResponse200216.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDnsLocalProfilesAssignments

> InlineResponse200277 GetOrganizationApplianceDnsLocalProfilesAssignments(ctx, organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()

Fetch the local DNS profile assignments in the organization



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
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by profile IDs (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationApplianceDnsLocalProfilesAssignments(context.Background(), organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationApplianceDnsLocalProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsLocalProfilesAssignments`: InlineResponse200277
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationApplianceDnsLocalProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsLocalProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**InlineResponse200277**](InlineResponse200277.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceDnsSplitProfilesAssignments

> InlineResponse200281 GetOrganizationApplianceDnsSplitProfilesAssignments(ctx, organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()

Fetch the split DNS profile assignments in the organization



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
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by profile IDs (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationApplianceDnsSplitProfilesAssignments(context.Background(), organizationId).ProfileIds(profileIds).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationApplianceDnsSplitProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsSplitProfilesAssignments`: InlineResponse200281
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationApplianceDnsSplitProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsSplitProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**InlineResponse200281**](InlineResponse200281.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataProfilesAssignments

> InlineResponse200346 GetOrganizationDevicesCellularDataProfilesAssignments(ctx, organizationId).ProfileIds(profileIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Cellular Data Management Profile assignments in this organization



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
    profileIds := []string{"Inner_example"} // []string | Optional parameter to find assignments by Profile IDs. Maximum 1000 profile IDs. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to find assignments by Device Serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationDevicesCellularDataProfilesAssignments(context.Background(), organizationId).ProfileIds(profileIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationDevicesCellularDataProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataProfilesAssignments`: InlineResponse200346
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationDevicesCellularDataProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to find assignments by Profile IDs. Maximum 1000 profile IDs. | 
 **serials** | **[]string** | Optional parameter to find assignments by Device Serials. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200346**](InlineResponse200346.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesAssignmentsByClient

> []InlineResponse200401 GetOrganizationPoliciesAssignmentsByClient(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).IncludeUndetectedClients(includeUndetectedClients).Execute()

Get policies for all clients with policies



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
    networkIds := []string{"Inner_example"} // []string | Network Ids (minimum: 1, maximum: 30)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    includeUndetectedClients := true // bool | Include provisioned clients that have not associated to the network. Default: false (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationPoliciesAssignmentsByClient(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).IncludeUndetectedClients(includeUndetectedClients).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationPoliciesAssignmentsByClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesAssignmentsByClient`: []InlineResponse200401
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationPoliciesAssignmentsByClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPoliciesAssignmentsByClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Network Ids (minimum: 1, maximum: 30) | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **includeUndetectedClients** | **bool** | Include provisioned clients that have not associated to the network. Default: false | 

### Return type

[**[]InlineResponse200401**](InlineResponse200401.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments

> InlineResponse200407 GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments(ctx, organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).AdaptivePolicyGroupIds(adaptivePolicyGroupIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments(context.Background(), organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).AdaptivePolicyGroupIds(adaptivePolicyGroupIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments`: InlineResponse200407
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesAdaptivePolicyGroupsAssignments`: %v\n", resp)
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

[**InlineResponse200407**](InlineResponse200407.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments

> InlineResponse200408 GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments(ctx, organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments(context.Background(), organizationId).AssignmentIds(assignmentIds).PolicyIds(policyIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments`: InlineResponse200408
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignments`: %v\n", resp)
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

[**InlineResponse200408**](InlineResponse200408.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan

> InlineResponse200409 GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan(ctx, organizationId).Search(search).VlanIds(vlanIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan(context.Background(), organizationId).Search(search).VlanIds(vlanIds).InterfaceIds(interfaceIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan`: InlineResponse200409
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesApplianceVlansAssignmentsByVlan`: %v\n", resp)
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

[**InlineResponse200409**](InlineResponse200409.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments

> InlineResponse200410 GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments(ctx, organizationId).RulesetIds(rulesetIds).PolicyIds(policyIds).AssignmentIds(assignmentIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

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
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments(context.Background(), organizationId).RulesetIds(rulesetIds).PolicyIds(policyIds).AssignmentIds(assignmentIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments`: InlineResponse200410
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignments`: %v\n", resp)
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

[**InlineResponse200410**](InlineResponse200410.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmSentryPoliciesAssignmentsByNetwork

> []InlineResponse200429 GetOrganizationSmSentryPoliciesAssignmentsByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

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
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationSmSentryPoliciesAssignmentsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationSmSentryPoliciesAssignmentsByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: []InlineResponse200429
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationSmSentryPoliciesAssignmentsByNetwork`: %v\n", resp)
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

[**[]InlineResponse200429**](InlineResponse200429.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessRfProfilesAssignmentsByDevice

> []InlineResponse200485 GetOrganizationWirelessRfProfilesAssignmentsByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).Models(models).Execute()

List the RF profiles of an organization by device



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. (optional)
    name := "name_example" // string | Optional parameter to filter RF profiles by device name. All returned devices will have a name that contains the search term or is an exact match. (optional)
    mac := "mac_example" // string | Optional parameter to filter RF profiles by device MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. (optional)
    serial := "serial_example" // string | Optional parameter to filter RF profiles by device serial number. All returned devices will have a serial number that contains the search term or is an exact match. (optional)
    model := "model_example" // string | Optional parameter to filter RF profiles by device model. All returned devices will have a model that contains the search term or is an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter RF profiles by one or more device MAC addresses. All returned devices will have a MAC address that is an exact match. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter RF profiles by one or more device serial numbers. All returned devices will have a serial number that is an exact match. (optional)
    models := []string{"Inner_example"} // []string | Optional parameter to filter RF profiles by one or more device models. All returned devices will have a model that is an exact match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationWirelessRfProfilesAssignmentsByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).Models(models).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationWirelessRfProfilesAssignmentsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessRfProfilesAssignmentsByDevice`: []InlineResponse200485
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationWirelessRfProfilesAssignmentsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessRfProfilesAssignmentsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | 
 **name** | **string** | Optional parameter to filter RF profiles by device name. All returned devices will have a name that contains the search term or is an exact match. | 
 **mac** | **string** | Optional parameter to filter RF profiles by device MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. | 
 **serial** | **string** | Optional parameter to filter RF profiles by device serial number. All returned devices will have a serial number that contains the search term or is an exact match. | 
 **model** | **string** | Optional parameter to filter RF profiles by device model. All returned devices will have a model that contains the search term or is an exact match. | 
 **macs** | **[]string** | Optional parameter to filter RF profiles by one or more device MAC addresses. All returned devices will have a MAC address that is an exact match. | 
 **serials** | **[]string** | Optional parameter to filter RF profiles by one or more device serial numbers. All returned devices will have a serial number that is an exact match. | 
 **models** | **[]string** | Optional parameter to filter RF profiles by one or more device models. All returned devices will have a model that is an exact match. | 

### Return type

[**[]InlineResponse200485**](InlineResponse200485.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessSsidsProfilesAssignments

> InlineResponse200489 GetOrganizationWirelessSsidsProfilesAssignments(ctx, organizationId).NetworkIds(networkIds).SsidIds(ssidIds).ProfileIds(profileIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the SSID profile assignments in an organization



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
    networkIds := []string{"Inner_example"} // []string | The network IDs to include in the result set. (optional)
    ssidIds := []string{"Inner_example"} // []string | The SSID IDs to include in the result set. (optional)
    profileIds := []string{"Inner_example"} // []string | The SSID profile IDs to include in the result set. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationWirelessSsidsProfilesAssignments(context.Background(), organizationId).NetworkIds(networkIds).SsidIds(ssidIds).ProfileIds(profileIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationWirelessSsidsProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessSsidsProfilesAssignments`: InlineResponse200489
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationWirelessSsidsProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessSsidsProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | The network IDs to include in the result set. | 
 **ssidIds** | **[]string** | The SSID IDs to include in the result set. | 
 **profileIds** | **[]string** | The SSID profile IDs to include in the result set. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200489**](InlineResponse200489.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork

> InlineResponse200490 GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork(ctx, organizationId).NetworkIds(networkIds).ProfileIds(profileIds).NetworkGroupIds(networkGroupIds).IncludeAllNetworks(includeAllNetworks).ExcludeProfileIds(excludeProfileIds).SortBy(sortBy).SortOrder(sortOrder).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the SSID profile assignments in an organization, grouped by network



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
    networkIds := []string{"Inner_example"} // []string | The network IDs to include in the result set. (optional)
    profileIds := []string{"Inner_example"} // []string | The SSID profile IDs to include in the result set. (optional)
    networkGroupIds := []string{"Inner_example"} // []string | The network group IDs to include in the result set. (optional)
    includeAllNetworks := true // bool | When set to true, include all networks in the organization, even those without any SSID profile assignments. Defaults to false. (optional)
    excludeProfileIds := []string{"Inner_example"} // []string | The SSID profile IDs to exclude from the result set. (optional)
    sortBy := "sortBy_example" // string | Optional parameter to specify the field used to sort results. (default: network) (optional)
    sortOrder := "sortOrder_example" // string | Optional parameter to specify the sort order. Default value is asc. (optional)
    search := "search_example" // string | Optional parameter to search on network name or network group name. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork(context.Background(), organizationId).NetworkIds(networkIds).ProfileIds(profileIds).NetworkGroupIds(networkGroupIds).IncludeAllNetworks(includeAllNetworks).ExcludeProfileIds(excludeProfileIds).SortBy(sortBy).SortOrder(sortOrder).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork`: InlineResponse200490
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.GetOrganizationWirelessSsidsProfilesAssignmentsByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessSsidsProfilesAssignmentsByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | The network IDs to include in the result set. | 
 **profileIds** | **[]string** | The SSID profile IDs to include in the result set. | 
 **networkGroupIds** | **[]string** | The network group IDs to include in the result set. | 
 **includeAllNetworks** | **bool** | When set to true, include all networks in the organization, even those without any SSID profile assignments. Defaults to false. | 
 **excludeProfileIds** | **[]string** | The SSID profile IDs to exclude from the result set. | 
 **sortBy** | **string** | Optional parameter to specify the field used to sort results. (default: network) | 
 **sortOrder** | **string** | Optional parameter to specify the sort order. Default value is asc. | 
 **search** | **string** | Optional parameter to search on network name or network group name. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200490**](InlineResponse200490.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReassignNetworkVlanProfilesAssignments

> InlineResponse200217 ReassignNetworkVlanProfilesAssignments(ctx, networkId).ReassignNetworkVlanProfilesAssignments(reassignNetworkVlanProfilesAssignments).Execute()

Update the assigned VLAN Profile for devices in a network



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
    networkId := "networkId_example" // string | Network ID
    reassignNetworkVlanProfilesAssignments := *openapiclient.NewInlineObject192([]string{"Serials_example"}, []string{"StackIds_example"}) // InlineObject192 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.ReassignNetworkVlanProfilesAssignments(context.Background(), networkId).ReassignNetworkVlanProfilesAssignments(reassignNetworkVlanProfilesAssignments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.ReassignNetworkVlanProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReassignNetworkVlanProfilesAssignments`: InlineResponse200217
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.ReassignNetworkVlanProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReassignNetworkVlanProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reassignNetworkVlanProfilesAssignments** | [**InlineObject192**](InlineObject192.md) |  | 

### Return type

[**InlineResponse200217**](InlineResponse200217.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment

> InlineResponse200410Items UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(ctx, organizationId, assignmentId).UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()

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
    updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment := *openapiclient.NewInlineObject330() // InlineObject330 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(context.Background(), organizationId, assignmentId).UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment(updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: InlineResponse200410Items
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.UpdateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment`: %v\n", resp)
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


 **updateOrganizationPoliciesGlobalGroupPoliciesFirewallRulesetsAssignment** | [**InlineObject330**](InlineObject330.md) |  | 

### Return type

[**InlineResponse200410Items**](InlineResponse200410Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSmSentryPoliciesAssignments

> InlineResponse200428 UpdateOrganizationSmSentryPoliciesAssignments(ctx, organizationId).UpdateOrganizationSmSentryPoliciesAssignments(updateOrganizationSmSentryPoliciesAssignments).Execute()

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
    updateOrganizationSmSentryPoliciesAssignments := *openapiclient.NewInlineObject348([]openapiclient.OrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems{*openapiclient.NewOrganizationsOrganizationIdSmSentryPoliciesAssignmentsItems("NetworkId_example")}) // InlineObject348 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssignmentsApi.UpdateOrganizationSmSentryPoliciesAssignments(context.Background(), organizationId).UpdateOrganizationSmSentryPoliciesAssignments(updateOrganizationSmSentryPoliciesAssignments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssignmentsApi.UpdateOrganizationSmSentryPoliciesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSmSentryPoliciesAssignments`: InlineResponse200428
    fmt.Fprintf(os.Stdout, "Response from `AssignmentsApi.UpdateOrganizationSmSentryPoliciesAssignments`: %v\n", resp)
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

 **updateOrganizationSmSentryPoliciesAssignments** | [**InlineObject348**](InlineObject348.md) |  | 

### Return type

[**InlineResponse200428**](InlineResponse200428.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

