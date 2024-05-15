# \ActionBatchesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationActionBatch**](ActionBatchesApi.md#CreateOrganizationActionBatch) | **Post** /organizations/{organizationId}/actionBatches | Create an action batch
[**DeleteOrganizationActionBatch**](ActionBatchesApi.md#DeleteOrganizationActionBatch) | **Delete** /organizations/{organizationId}/actionBatches/{actionBatchId} | Delete an action batch
[**GetOrganizationActionBatch**](ActionBatchesApi.md#GetOrganizationActionBatch) | **Get** /organizations/{organizationId}/actionBatches/{actionBatchId} | Return an action batch
[**GetOrganizationActionBatches**](ActionBatchesApi.md#GetOrganizationActionBatches) | **Get** /organizations/{organizationId}/actionBatches | Return the list of action batches in the organization
[**UpdateOrganizationActionBatch**](ActionBatchesApi.md#UpdateOrganizationActionBatch) | **Put** /organizations/{organizationId}/actionBatches/{actionBatchId} | Update an action batch



## CreateOrganizationActionBatch

> InlineResponse20111 CreateOrganizationActionBatch(ctx, organizationId).CreateOrganizationActionBatch(createOrganizationActionBatch).Execute()

Create an action batch



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
    createOrganizationActionBatch := *openapiclient.NewInlineObject197([]openapiclient.OrganizationsOrganizationIdActionBatchesActions1{*openapiclient.NewOrganizationsOrganizationIdActionBatchesActions1("Resource_example", "Operation_example")}) // InlineObject197 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.CreateOrganizationActionBatch(context.Background(), organizationId).CreateOrganizationActionBatch(createOrganizationActionBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.CreateOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationActionBatch`: InlineResponse20111
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.CreateOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationActionBatch** | [**InlineObject197**](InlineObject197.md) |  | 

### Return type

[**InlineResponse20111**](InlineResponse20111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationActionBatch

> DeleteOrganizationActionBatch(ctx, organizationId, actionBatchId).Execute()

Delete an action batch



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
    actionBatchId := "actionBatchId_example" // string | Action batch ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.DeleteOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.DeleteOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationActionBatchRequest struct via the builder pattern


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


## GetOrganizationActionBatch

> InlineResponse20111 GetOrganizationActionBatch(ctx, organizationId, actionBatchId).Execute()

Return an action batch



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
    actionBatchId := "actionBatchId_example" // string | Action batch ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.GetOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.GetOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationActionBatch`: InlineResponse20111
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.GetOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20111**](InlineResponse20111.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationActionBatches

> []InlineResponse200195 GetOrganizationActionBatches(ctx, organizationId).Status(status).Execute()

Return the list of action batches in the organization



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
    status := "status_example" // string | Filter batches by status. Valid types are pending, completed, and failed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.GetOrganizationActionBatches(context.Background(), organizationId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.GetOrganizationActionBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationActionBatches`: []InlineResponse200195
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.GetOrganizationActionBatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationActionBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter batches by status. Valid types are pending, completed, and failed. | 

### Return type

[**[]InlineResponse200195**](InlineResponse200195.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationActionBatch

> InlineResponse200195 UpdateOrganizationActionBatch(ctx, organizationId, actionBatchId).UpdateOrganizationActionBatch(updateOrganizationActionBatch).Execute()

Update an action batch



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
    actionBatchId := "actionBatchId_example" // string | Action batch ID
    updateOrganizationActionBatch := *openapiclient.NewInlineObject198() // InlineObject198 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.UpdateOrganizationActionBatch(context.Background(), organizationId, actionBatchId).UpdateOrganizationActionBatch(updateOrganizationActionBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.UpdateOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationActionBatch`: InlineResponse200195
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.UpdateOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationActionBatch** | [**InlineObject198**](InlineObject198.md) |  | 

### Return type

[**InlineResponse200195**](InlineResponse200195.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

