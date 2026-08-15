# \IntegrationsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSaseIntegration**](IntegrationsApi.md#CreateOrganizationSaseIntegration) | **Post** /organizations/{organizationId}/sase/integrations | Create a new Secure Access integration
[**DeleteOrganizationSaseIntegration**](IntegrationsApi.md#DeleteOrganizationSaseIntegration) | **Delete** /organizations/{organizationId}/sase/integrations/{integrationId} | Remove a Secure Access integration
[**DisableOrganizationIntegrationsXdrNetworks**](IntegrationsApi.md#DisableOrganizationIntegrationsXdrNetworks) | **Post** /organizations/{organizationId}/integrations/xdr/networks/disable | Disable XDR on networks
[**EnableOrganizationIntegrationsXdrNetworks**](IntegrationsApi.md#EnableOrganizationIntegrationsXdrNetworks) | **Post** /organizations/{organizationId}/integrations/xdr/networks/enable | Enable XDR on networks
[**GetOrganizationIntegrationsDeployable**](IntegrationsApi.md#GetOrganizationIntegrationsDeployable) | **Get** /organizations/{organizationId}/integrations/deployable | Provides a list of integrations that can be enabled for an Organization.
[**GetOrganizationIntegrationsDeployed**](IntegrationsApi.md#GetOrganizationIntegrationsDeployed) | **Get** /organizations/{organizationId}/integrations/deployed | Provides a list of integrations enabled for an Organization.
[**GetOrganizationIntegrationsXdrNetworks**](IntegrationsApi.md#GetOrganizationIntegrationsXdrNetworks) | **Get** /organizations/{organizationId}/integrations/xdr/networks | Returns the networks in the organization that have XDR enabled
[**GetOrganizationSaseIntegration**](IntegrationsApi.md#GetOrganizationSaseIntegration) | **Get** /organizations/{organizationId}/sase/integrations/{integrationId} | Get details of a Secure Access integration



## CreateOrganizationSaseIntegration

> InlineResponse20128 CreateOrganizationSaseIntegration(ctx, organizationId).CreateOrganizationSaseIntegration(createOrganizationSaseIntegration).Execute()

Create a new Secure Access integration



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
    createOrganizationSaseIntegration := *openapiclient.NewInlineObject342(*openapiclient.NewOrganizationsOrganizationIdSaseIntegrationsApi("Key_example", "Secret_example")) // InlineObject342 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.CreateOrganizationSaseIntegration(context.Background(), organizationId).CreateOrganizationSaseIntegration(createOrganizationSaseIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.CreateOrganizationSaseIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSaseIntegration`: InlineResponse20128
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.CreateOrganizationSaseIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSaseIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSaseIntegration** | [**InlineObject342**](InlineObject342.md) |  | 

### Return type

[**InlineResponse20128**](InlineResponse20128.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSaseIntegration

> DeleteOrganizationSaseIntegration(ctx, organizationId, integrationId).Execute()

Remove a Secure Access integration



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
    integrationId := "integrationId_example" // string | Integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.DeleteOrganizationSaseIntegration(context.Background(), organizationId, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.DeleteOrganizationSaseIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**integrationId** | **string** | Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSaseIntegrationRequest struct via the builder pattern


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


## DisableOrganizationIntegrationsXdrNetworks

> InlineResponse200381 DisableOrganizationIntegrationsXdrNetworks(ctx, organizationId).DisableOrganizationIntegrationsXdrNetworks(disableOrganizationIntegrationsXdrNetworks).Execute()

Disable XDR on networks



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
    disableOrganizationIntegrationsXdrNetworks := *openapiclient.NewInlineObject299([]openapiclient.OrganizationsOrganizationIdIntegrationsXdrNetworksDisableNetworks{*openapiclient.NewOrganizationsOrganizationIdIntegrationsXdrNetworksDisableNetworks("NetworkId_example", []string{"ProductTypes_example"})}) // InlineObject299 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.DisableOrganizationIntegrationsXdrNetworks(context.Background(), organizationId).DisableOrganizationIntegrationsXdrNetworks(disableOrganizationIntegrationsXdrNetworks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.DisableOrganizationIntegrationsXdrNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableOrganizationIntegrationsXdrNetworks`: InlineResponse200381
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.DisableOrganizationIntegrationsXdrNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableOrganizationIntegrationsXdrNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disableOrganizationIntegrationsXdrNetworks** | [**InlineObject299**](InlineObject299.md) |  | 

### Return type

[**InlineResponse200381**](InlineResponse200381.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableOrganizationIntegrationsXdrNetworks

> InlineResponse200382 EnableOrganizationIntegrationsXdrNetworks(ctx, organizationId).EnableOrganizationIntegrationsXdrNetworks(enableOrganizationIntegrationsXdrNetworks).Execute()

Enable XDR on networks



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
    enableOrganizationIntegrationsXdrNetworks := *openapiclient.NewInlineObject300([]openapiclient.OrganizationsOrganizationIdIntegrationsXdrNetworksEnableNetworks{*openapiclient.NewOrganizationsOrganizationIdIntegrationsXdrNetworksEnableNetworks("NetworkId_example", []string{"ProductTypes_example"})}) // InlineObject300 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.EnableOrganizationIntegrationsXdrNetworks(context.Background(), organizationId).EnableOrganizationIntegrationsXdrNetworks(enableOrganizationIntegrationsXdrNetworks).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.EnableOrganizationIntegrationsXdrNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableOrganizationIntegrationsXdrNetworks`: InlineResponse200382
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.EnableOrganizationIntegrationsXdrNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableOrganizationIntegrationsXdrNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableOrganizationIntegrationsXdrNetworks** | [**InlineObject300**](InlineObject300.md) |  | 

### Return type

[**InlineResponse200382**](InlineResponse200382.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationIntegrationsDeployable

> InlineResponse200378 GetOrganizationIntegrationsDeployable(ctx, organizationId).Execute()

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
    resp, r, err := apiClient.IntegrationsApi.GetOrganizationIntegrationsDeployable(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetOrganizationIntegrationsDeployable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationIntegrationsDeployable`: InlineResponse200378
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetOrganizationIntegrationsDeployable`: %v\n", resp)
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

[**InlineResponse200378**](InlineResponse200378.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationIntegrationsDeployed

> InlineResponse200379 GetOrganizationIntegrationsDeployed(ctx, organizationId).Execute()

Provides a list of integrations enabled for an Organization.



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
    resp, r, err := apiClient.IntegrationsApi.GetOrganizationIntegrationsDeployed(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetOrganizationIntegrationsDeployed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationIntegrationsDeployed`: InlineResponse200379
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetOrganizationIntegrationsDeployed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationIntegrationsDeployedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetOrganizationIntegrationsXdrNetworks

> InlineResponse200380 GetOrganizationIntegrationsXdrNetworks(ctx, organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns the networks in the organization that have XDR enabled



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 20. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GetOrganizationIntegrationsXdrNetworks(context.Background(), organizationId).NetworkIds(networkIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetOrganizationIntegrationsXdrNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationIntegrationsXdrNetworks`: InlineResponse200380
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetOrganizationIntegrationsXdrNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationIntegrationsXdrNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 20. | 
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


## GetOrganizationSaseIntegration

> InlineResponse20128 GetOrganizationSaseIntegration(ctx, organizationId, integrationId).Execute()

Get details of a Secure Access integration



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
    integrationId := "integrationId_example" // string | Integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GetOrganizationSaseIntegration(context.Background(), organizationId, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetOrganizationSaseIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSaseIntegration`: InlineResponse20128
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetOrganizationSaseIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**integrationId** | **string** | Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSaseIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20128**](InlineResponse20128.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

