# \DeployableApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationIntegrationsDeployable**](DeployableApi.md#GetOrganizationIntegrationsDeployable) | **Get** /organizations/{organizationId}/integrations/deployable | Provides a list of integrations that can be enabled for an Organization.



## GetOrganizationIntegrationsDeployable

> InlineResponse200314 GetOrganizationIntegrationsDeployable(ctx, organizationId).Execute()

Provides a list of integrations that can be enabled for an Organization.



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
    resp, r, err := apiClient.DeployableApi.GetOrganizationIntegrationsDeployable(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeployableApi.GetOrganizationIntegrationsDeployable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationIntegrationsDeployable`: InlineResponse200314
    fmt.Fprintf(os.Stdout, "Response from `DeployableApi.GetOrganizationIntegrationsDeployable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationIntegrationsDeployableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200314**](InlineResponse200314.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

