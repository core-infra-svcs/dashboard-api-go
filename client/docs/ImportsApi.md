# \ImportsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInventoryOnboardingCloudMonitoringImport**](ImportsApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringImport) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.
[**GetOrganizationInventoryOnboardingCloudMonitoringImports**](ImportsApi.md#GetOrganizationInventoryOnboardingCloudMonitoringImports) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Check the status of a committed Import operation



## CreateOrganizationInventoryOnboardingCloudMonitoringImport

> []InlineResponse2017 CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.



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
    organizationId := "organizationId_example" // string | 
    createOrganizationInventoryOnboardingCloudMonitoringImport := *openapiclient.NewInlineObject200([]openapiclient.OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices{*openapiclient.NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices("DeviceId_example", "Udi_example", "NetworkId_example")}) // InlineObject200 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringImport`: []InlineResponse2017
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringImport** | [**InlineObject200**](InlineObject200.md) |  | 

### Return type

[**[]InlineResponse2017**](InlineResponse2017.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringImports

> []InlineResponse200111 GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx, organizationId).ImportIds(importIds).Execute()

Check the status of a committed Import operation



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
    organizationId := "organizationId_example" // string | 
    importIds := []string{"Inner_example"} // []string | import ids from an imports

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportsApi.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).ImportIds(importIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportsApi.GetOrganizationInventoryOnboardingCloudMonitoringImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringImports`: []InlineResponse200111
    fmt.Fprintf(os.Stdout, "Response from `ImportsApi.GetOrganizationInventoryOnboardingCloudMonitoringImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importIds** | **[]string** | import ids from an imports | 

### Return type

[**[]InlineResponse200111**](InlineResponse200111.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

