# \MigrationsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationDevicesControllerMigration**](MigrationsApi.md#CreateOrganizationDevicesControllerMigration) | **Post** /organizations/{organizationId}/devices/controller/migrations | Migrate devices to another controller or management mode
[**GetOrganizationDevicesControllerMigrations**](MigrationsApi.md#GetOrganizationDevicesControllerMigrations) | **Get** /organizations/{organizationId}/devices/controller/migrations | Retrieve device migration statuses in an organization



## CreateOrganizationDevicesControllerMigration

> []InlineResponse200276Items CreateOrganizationDevicesControllerMigration(ctx, organizationId).CreateOrganizationDevicesControllerMigration(createOrganizationDevicesControllerMigration).Execute()

Migrate devices to another controller or management mode



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
    createOrganizationDevicesControllerMigration := *openapiclient.NewInlineObject257([]string{"Serials_example"}, "Target_example") // InlineObject257 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.CreateOrganizationDevicesControllerMigration(context.Background(), organizationId).CreateOrganizationDevicesControllerMigration(createOrganizationDevicesControllerMigration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.CreateOrganizationDevicesControllerMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDevicesControllerMigration`: []InlineResponse200276Items
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.CreateOrganizationDevicesControllerMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDevicesControllerMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationDevicesControllerMigration** | [**InlineObject257**](InlineObject257.md) |  | 

### Return type

[**[]InlineResponse200276Items**](InlineResponse200276Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesControllerMigrations

> InlineResponse200276 GetOrganizationDevicesControllerMigrations(ctx, organizationId).Serials(serials).NetworkIds(networkIds).Target(target).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Retrieve device migration statuses in an organization



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
    serials := []string{"Inner_example"} // []string | A list of Meraki Serials for which to retrieve migrations (optional)
    networkIds := []string{"Inner_example"} // []string | Filter device migrations by network IDs (optional)
    target := "target_example" // string | Filter device migrations by target destination (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MigrationsApi.GetOrganizationDevicesControllerMigrations(context.Background(), organizationId).Serials(serials).NetworkIds(networkIds).Target(target).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationsApi.GetOrganizationDevicesControllerMigrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesControllerMigrations`: InlineResponse200276
    fmt.Fprintf(os.Stdout, "Response from `MigrationsApi.GetOrganizationDevicesControllerMigrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesControllerMigrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of Meraki Serials for which to retrieve migrations | 
 **networkIds** | **[]string** | Filter device migrations by network IDs | 
 **target** | **string** | Filter device migrations by target destination | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200276**](InlineResponse200276.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

