# \BulkDeleteApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete**](BulkDeleteApi.md#CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete) | **Post** /organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkDelete | Unassign the local DNS profile to networks in the organization
[**CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete**](BulkDeleteApi.md#CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete) | **Post** /organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkDelete | Unassign the split DNS profile to networks in the organization



## CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete

> InlineResponse200223 CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(ctx, organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete).Execute()

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
    createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete := *openapiclient.NewInlineObject222([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems()}) // InlineObject222 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkDeleteApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(context.Background(), organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkDeleteApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete`: InlineResponse200223
    fmt.Fprintf(os.Stdout, "Response from `BulkDeleteApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete`: %v\n", resp)
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

 **createOrganizationApplianceDnsLocalProfilesAssignmentsBulkDelete** | [**InlineObject222**](InlineObject222.md) |  | 

### Return type

[**InlineResponse200223**](InlineResponse200223.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete

> InlineResponse200227 CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(ctx, organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete).Execute()

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
    createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete := *openapiclient.NewInlineObject228([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkDeleteItems()}) // InlineObject228 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkDeleteApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(context.Background(), organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkDeleteApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete`: InlineResponse200227
    fmt.Fprintf(os.Stdout, "Response from `BulkDeleteApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete`: %v\n", resp)
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

 **createOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete** | [**InlineObject228**](InlineObject228.md) |  | 

### Return type

[**InlineResponse200227**](InlineResponse200227.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

