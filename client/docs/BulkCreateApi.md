# \BulkCreateApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate**](BulkCreateApi.md#CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate) | **Post** /organizations/{organizationId}/appliance/dns/local/profiles/assignments/bulkCreate | Assign the local DNS profile to networks in the organization
[**CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate**](BulkCreateApi.md#CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate) | **Post** /organizations/{organizationId}/appliance/dns/split/profiles/assignments/bulkCreate | Assign the split DNS profile to networks in the organization



## CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate

> InlineResponse200220 CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate(ctx, organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate).Execute()

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
    createOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate := *openapiclient.NewInlineObject218([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems()}) // InlineObject218 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkCreateApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate(context.Background(), organizationId).CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkCreateApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate`: InlineResponse200220
    fmt.Fprintf(os.Stdout, "Response from `BulkCreateApi.CreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsLocalProfilesAssignmentsBulkCreate** | [**InlineObject218**](InlineObject218.md) |  | 

### Return type

[**InlineResponse200220**](InlineResponse200220.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate

> InlineResponse200224 CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(ctx, organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate).Execute()

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
    createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate := *openapiclient.NewInlineObject224([]openapiclient.OrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems{*openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalProfilesAssignmentsBulkCreateItems()}) // InlineObject224 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkCreateApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(context.Background(), organizationId).CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate(createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkCreateApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate`: InlineResponse200224
    fmt.Fprintf(os.Stdout, "Response from `BulkCreateApi.CreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate`: %v\n", resp)
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

 **createOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate** | [**InlineObject224**](InlineObject224.md) |  | 

### Return type

[**InlineResponse200224**](InlineResponse200224.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

