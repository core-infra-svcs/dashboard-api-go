# \SpacesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSpacesIntegrateStatus**](SpacesApi.md#GetOrganizationSpacesIntegrateStatus) | **Get** /organizations/{organizationId}/spaces/integrate/status | Get the status of the Spaces integration in Meraki
[**RemoveOrganizationSpacesIntegration**](SpacesApi.md#RemoveOrganizationSpacesIntegration) | **Post** /organizations/{organizationId}/spaces/integration/remove | Remove the Spaces integration from Meraki



## GetOrganizationSpacesIntegrateStatus

> InlineResponse200337 GetOrganizationSpacesIntegrateStatus(ctx, organizationId).Execute()

Get the status of the Spaces integration in Meraki



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetOrganizationSpacesIntegrateStatus(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetOrganizationSpacesIntegrateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSpacesIntegrateStatus`: InlineResponse200337
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetOrganizationSpacesIntegrateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSpacesIntegrateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200337**](InlineResponse200337.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationSpacesIntegration

> InlineResponse200338 RemoveOrganizationSpacesIntegration(ctx, organizationId).Execute()

Remove the Spaces integration from Meraki



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.RemoveOrganizationSpacesIntegration(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.RemoveOrganizationSpacesIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganizationSpacesIntegration`: InlineResponse200338
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.RemoveOrganizationSpacesIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationSpacesIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200338**](InlineResponse200338.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

