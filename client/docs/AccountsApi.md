# \AccountsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCellularGatewayEsimsServiceProvidersAccount**](AccountsApi.md#CreateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Post** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Add a service provider account.
[**DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount**](AccountsApi.md#DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Delete** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Remove a service provider account&#39;s integration with the Dashboard.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccounts**](AccountsApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccounts) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Inventory of service provider accounts tied to the organization.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans**](AccountsApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/communicationPlans | The communication plans available for a given provider.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans**](AccountsApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/ratePlans | The rate plans available for a given provider.
[**UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount**](AccountsApi.md#UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Put** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Edit service provider account info stored in Meraki&#39;s database.



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
    createOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject229("AccountId_example", "ApiKey_example", *openapiclient.NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1(), "Title_example", "Username_example") // InlineObject229 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId).CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(createOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
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

 **createOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject229**](InlineObject229.md) |  | 

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
    resp, r, err := apiClient.AccountsApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
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


## GetOrganizationCellularGatewayEsimsServiceProvidersAccounts

> []InlineResponse200234 GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: []InlineResponse200234
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: %v\n", resp)
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

[**[]InlineResponse200234**](InlineResponse200234.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans

> InlineResponse200235 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: InlineResponse200235
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: %v\n", resp)
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

[**InlineResponse200235**](InlineResponse200235.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans

> InlineResponse200236 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: InlineResponse200236
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: %v\n", resp)
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

[**InlineResponse200236**](InlineResponse200236.md)

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
    updateOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject230() // InlineObject230 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(updateOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
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


 **updateOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject230**](InlineObject230.md) |  | 

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

