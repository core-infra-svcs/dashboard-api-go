# \RecordsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationApplianceDnsLocalRecord**](RecordsApi.md#CreateOrganizationApplianceDnsLocalRecord) | **Post** /organizations/{organizationId}/appliance/dns/local/records | Create a new local DNS record
[**DeleteOrganizationApplianceDnsLocalRecord**](RecordsApi.md#DeleteOrganizationApplianceDnsLocalRecord) | **Delete** /organizations/{organizationId}/appliance/dns/local/records/{recordId} | Deletes a local DNS record
[**GetOrganizationApplianceDnsLocalRecords**](RecordsApi.md#GetOrganizationApplianceDnsLocalRecords) | **Get** /organizations/{organizationId}/appliance/dns/local/records | Fetch the DNS records used in local DNS profiles
[**UpdateOrganizationApplianceDnsLocalRecord**](RecordsApi.md#UpdateOrganizationApplianceDnsLocalRecord) | **Put** /organizations/{organizationId}/appliance/dns/local/records/{recordId} | Updates a local DNS record



## CreateOrganizationApplianceDnsLocalRecord

> []InlineResponse200226 CreateOrganizationApplianceDnsLocalRecord(ctx, organizationId).CreateOrganizationApplianceDnsLocalRecord(createOrganizationApplianceDnsLocalRecord).Execute()

Create a new local DNS record



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
    createOrganizationApplianceDnsLocalRecord := *openapiclient.NewInlineObject227("Hostname_example", "Address_example", *openapiclient.NewOrganizationsOrganizationIdApplianceDnsLocalRecordsProfile()) // InlineObject227 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsApi.CreateOrganizationApplianceDnsLocalRecord(context.Background(), organizationId).CreateOrganizationApplianceDnsLocalRecord(createOrganizationApplianceDnsLocalRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.CreateOrganizationApplianceDnsLocalRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationApplianceDnsLocalRecord`: []InlineResponse200226
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.CreateOrganizationApplianceDnsLocalRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationApplianceDnsLocalRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationApplianceDnsLocalRecord** | [**InlineObject227**](InlineObject227.md) |  | 

### Return type

[**[]InlineResponse200226**](InlineResponse200226.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationApplianceDnsLocalRecord

> DeleteOrganizationApplianceDnsLocalRecord(ctx, organizationId, recordId).Execute()

Deletes a local DNS record



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
    recordId := "recordId_example" // string | Record ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsApi.DeleteOrganizationApplianceDnsLocalRecord(context.Background(), organizationId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.DeleteOrganizationApplianceDnsLocalRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**recordId** | **string** | Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationApplianceDnsLocalRecordRequest struct via the builder pattern


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


## GetOrganizationApplianceDnsLocalRecords

> []InlineResponse200226 GetOrganizationApplianceDnsLocalRecords(ctx, organizationId).ProfileIds(profileIds).Execute()

Fetch the DNS records used in local DNS profiles



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsApi.GetOrganizationApplianceDnsLocalRecords(context.Background(), organizationId).ProfileIds(profileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.GetOrganizationApplianceDnsLocalRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceDnsLocalRecords`: []InlineResponse200226
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.GetOrganizationApplianceDnsLocalRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceDnsLocalRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by profile IDs | 

### Return type

[**[]InlineResponse200226**](InlineResponse200226.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceDnsLocalRecord

> InlineResponse200226 UpdateOrganizationApplianceDnsLocalRecord(ctx, organizationId, recordId).UpdateOrganizationApplianceDnsLocalRecord(updateOrganizationApplianceDnsLocalRecord).Execute()

Updates a local DNS record



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
    recordId := "recordId_example" // string | Record ID
    updateOrganizationApplianceDnsLocalRecord := *openapiclient.NewInlineObject228() // InlineObject228 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsApi.UpdateOrganizationApplianceDnsLocalRecord(context.Background(), organizationId, recordId).UpdateOrganizationApplianceDnsLocalRecord(updateOrganizationApplianceDnsLocalRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.UpdateOrganizationApplianceDnsLocalRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceDnsLocalRecord`: InlineResponse200226
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.UpdateOrganizationApplianceDnsLocalRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**recordId** | **string** | Record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceDnsLocalRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationApplianceDnsLocalRecord** | [**InlineObject228**](InlineObject228.md) |  | 

### Return type

[**InlineResponse200226**](InlineResponse200226.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

