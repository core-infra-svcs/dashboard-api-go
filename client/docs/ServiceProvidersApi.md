# \ServiceProvidersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCellularGatewayEsimsServiceProvidersAccount**](ServiceProvidersApi.md#CreateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Post** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Add a service provider account.
[**DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount**](ServiceProvidersApi.md#DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Delete** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Remove a service provider account&#39;s integration with the Dashboard.
[**GetOrganizationCellularGatewayEsimsServiceProviders**](ServiceProvidersApi.md#GetOrganizationCellularGatewayEsimsServiceProviders) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders | Service providers customers can add accounts for.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccounts**](ServiceProvidersApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccounts) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Inventory of service provider accounts tied to the organization.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans**](ServiceProvidersApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/communicationPlans | The communication plans available for a given provider.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans**](ServiceProvidersApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/ratePlans | The rate plans available for a given provider.
[**UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount**](ServiceProvidersApi.md#UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Put** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Edit service provider account info stored in Meraki&#39;s database.



## CreateOrganizationCellularGatewayEsimsServiceProvidersAccount

> OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(ctx, organizationId).CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(createOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()

Add a service provider account.



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
    createOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject243("AccountId_example", "ApiKey_example", *openapiclient.NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1(), "Title_example", "Username_example") // InlineObject243 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceProvidersApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId).CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(createOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvidersApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `ServiceProvidersApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCellularGatewayEsimsServiceProvidersAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject243**](InlineObject243.md) |  | 

### Return type

[**OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems**](OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount

> DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount(ctx, organizationId, accountId).Execute()

Remove a service provider account's integration with the Dashboard.



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
    accountId := "accountId_example" // string | Account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceProvidersApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvidersApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCellularGatewayEsimsServiceProvidersAccountRequest struct via the builder pattern


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


## GetOrganizationCellularGatewayEsimsServiceProviders

> InlineResponse200249 GetOrganizationCellularGatewayEsimsServiceProviders(ctx, organizationId).Execute()

Service providers customers can add accounts for.



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
    resp, r, err := apiClient.ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProviders(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProviders`: InlineResponse200249
    fmt.Fprintf(os.Stdout, "Response from `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayEsimsServiceProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200249**](InlineResponse200249.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccounts

> []InlineResponse200250 GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(ctx, organizationId).AccountIds(accountIds).Execute()

Inventory of service provider accounts tied to the organization.



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
    accountIds := []int32{int32(123)} // []int32 | Optional parameter to filter the results by service provider account IDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: []InlineResponse200250
    fmt.Fprintf(os.Stdout, "Response from `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIds** | **[]int32** | Optional parameter to filter the results by service provider account IDs. | 

### Return type

[**[]InlineResponse200250**](InlineResponse200250.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans

> InlineResponse200251 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(ctx, organizationId).AccountIds(accountIds).Execute()

The communication plans available for a given provider.



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
    accountIds := []string{"Inner_example"} // []string | Account IDs that communication plans will be fetched for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: InlineResponse200251
    fmt.Fprintf(os.Stdout, "Response from `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIds** | **[]string** | Account IDs that communication plans will be fetched for | 

### Return type

[**InlineResponse200251**](InlineResponse200251.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans

> InlineResponse200252 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(ctx, organizationId).AccountIds(accountIds).Execute()

The rate plans available for a given provider.



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
    accountIds := []string{"Inner_example"} // []string | Account IDs that rate plans will be fetched for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: InlineResponse200252
    fmt.Fprintf(os.Stdout, "Response from `ServiceProvidersApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIds** | **[]string** | Account IDs that rate plans will be fetched for | 

### Return type

[**InlineResponse200252**](InlineResponse200252.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount

> OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(ctx, organizationId, accountId).UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(updateOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()

Edit service provider account info stored in Meraki's database.



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
    accountId := "accountId_example" // string | Account ID
    updateOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject244() // InlineObject244 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceProvidersApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(updateOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvidersApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `ServiceProvidersApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**accountId** | **string** | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCellularGatewayEsimsServiceProvidersAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject244**](InlineObject244.md) |  | 

### Return type

[**OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems**](OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

