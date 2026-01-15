# \EsimsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCellularGatewayEsimsServiceProvidersAccount**](EsimsApi.md#CreateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Post** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Add a service provider account.
[**CreateOrganizationCellularGatewayEsimsSwap**](EsimsApi.md#CreateOrganizationCellularGatewayEsimsSwap) | **Post** /organizations/{organizationId}/cellularGateway/esims/swap | Swap which profile an eSIM uses.
[**DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount**](EsimsApi.md#DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Delete** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Remove a service provider account&#39;s integration with the Dashboard.
[**GetOrganizationCellularGatewayEsimsInventory**](EsimsApi.md#GetOrganizationCellularGatewayEsimsInventory) | **Get** /organizations/{organizationId}/cellularGateway/esims/inventory | The eSIM inventory of a given organization.
[**GetOrganizationCellularGatewayEsimsServiceProviders**](EsimsApi.md#GetOrganizationCellularGatewayEsimsServiceProviders) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders | Service providers customers can add accounts for.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccounts**](EsimsApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccounts) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Inventory of service provider accounts tied to the organization.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans**](EsimsApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/communicationPlans | The communication plans available for a given provider.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans**](EsimsApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/ratePlans | The rate plans available for a given provider.
[**UpdateOrganizationCellularGatewayEsimsInventory**](EsimsApi.md#UpdateOrganizationCellularGatewayEsimsInventory) | **Put** /organizations/{organizationId}/cellularGateway/esims/inventory/{id} | Toggle the status of an eSIM
[**UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount**](EsimsApi.md#UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Put** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Edit service provider account info stored in Meraki&#39;s database.
[**UpdateOrganizationCellularGatewayEsimsSwap**](EsimsApi.md#UpdateOrganizationCellularGatewayEsimsSwap) | **Put** /organizations/{organizationId}/cellularGateway/esims/swap/{id} | Get the status of a profile swap.



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
    createOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject249("AccountId_example", "ApiKey_example", *openapiclient.NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1("Name_example"), "Title_example", "Username_example") // InlineObject249 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EsimsApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId).CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(createOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
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

 **createOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject249**](InlineObject249.md) |  | 

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


## CreateOrganizationCellularGatewayEsimsSwap

> InlineResponse200271 CreateOrganizationCellularGatewayEsimsSwap(ctx, organizationId).CreateOrganizationCellularGatewayEsimsSwap(createOrganizationCellularGatewayEsimsSwap).Execute()

Swap which profile an eSIM uses.



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
    createOrganizationCellularGatewayEsimsSwap := *openapiclient.NewInlineObject251([]openapiclient.OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps{*openapiclient.NewOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps("Eid_example")}) // InlineObject251 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EsimsApi.CreateOrganizationCellularGatewayEsimsSwap(context.Background(), organizationId).CreateOrganizationCellularGatewayEsimsSwap(createOrganizationCellularGatewayEsimsSwap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.CreateOrganizationCellularGatewayEsimsSwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCellularGatewayEsimsSwap`: InlineResponse200271
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.CreateOrganizationCellularGatewayEsimsSwap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCellularGatewayEsimsSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCellularGatewayEsimsSwap** | [**InlineObject251**](InlineObject251.md) |  | 

### Return type

[**InlineResponse200271**](InlineResponse200271.md)

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
    resp, r, err := apiClient.EsimsApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
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


## GetOrganizationCellularGatewayEsimsInventory

> InlineResponse200266 GetOrganizationCellularGatewayEsimsInventory(ctx, organizationId).Eids(eids).Execute()

The eSIM inventory of a given organization.



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
    eids := []string{"Inner_example"} // []string | Optional parameter to filter the results by EID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EsimsApi.GetOrganizationCellularGatewayEsimsInventory(context.Background(), organizationId).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.GetOrganizationCellularGatewayEsimsInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsInventory`: InlineResponse200266
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.GetOrganizationCellularGatewayEsimsInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayEsimsInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eids** | **[]string** | Optional parameter to filter the results by EID. | 

### Return type

[**InlineResponse200266**](InlineResponse200266.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProviders

> InlineResponse200267 GetOrganizationCellularGatewayEsimsServiceProviders(ctx, organizationId).Execute()

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
    resp, r, err := apiClient.EsimsApi.GetOrganizationCellularGatewayEsimsServiceProviders(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProviders`: InlineResponse200267
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProviders`: %v\n", resp)
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

[**InlineResponse200267**](InlineResponse200267.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccounts

> []InlineResponse200268 GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: []InlineResponse200268
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: %v\n", resp)
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

[**[]InlineResponse200268**](InlineResponse200268.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans

> InlineResponse200269 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: InlineResponse200269
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: %v\n", resp)
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

[**InlineResponse200269**](InlineResponse200269.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans

> InlineResponse200270 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: InlineResponse200270
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: %v\n", resp)
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

[**InlineResponse200270**](InlineResponse200270.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCellularGatewayEsimsInventory

> InlineResponse200266Items UpdateOrganizationCellularGatewayEsimsInventory(ctx, organizationId, id).UpdateOrganizationCellularGatewayEsimsInventory(updateOrganizationCellularGatewayEsimsInventory).Execute()

Toggle the status of an eSIM



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
    id := "id_example" // string | ID
    updateOrganizationCellularGatewayEsimsInventory := *openapiclient.NewInlineObject248() // InlineObject248 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EsimsApi.UpdateOrganizationCellularGatewayEsimsInventory(context.Background(), organizationId, id).UpdateOrganizationCellularGatewayEsimsInventory(updateOrganizationCellularGatewayEsimsInventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.UpdateOrganizationCellularGatewayEsimsInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsInventory`: InlineResponse200266Items
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.UpdateOrganizationCellularGatewayEsimsInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCellularGatewayEsimsInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationCellularGatewayEsimsInventory** | [**InlineObject248**](InlineObject248.md) |  | 

### Return type

[**InlineResponse200266Items**](InlineResponse200266Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
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
    updateOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject250() // InlineObject250 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EsimsApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(updateOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
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


 **updateOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject250**](InlineObject250.md) |  | 

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


## UpdateOrganizationCellularGatewayEsimsSwap

> InlineResponse200271 UpdateOrganizationCellularGatewayEsimsSwap(ctx, id, organizationId).Execute()

Get the status of a profile swap.



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
    id := "id_example" // string | eSIM EID
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EsimsApi.UpdateOrganizationCellularGatewayEsimsSwap(context.Background(), id, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EsimsApi.UpdateOrganizationCellularGatewayEsimsSwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsSwap`: InlineResponse200271
    fmt.Fprintf(os.Stdout, "Response from `EsimsApi.UpdateOrganizationCellularGatewayEsimsSwap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | eSIM EID | 
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCellularGatewayEsimsSwapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200271**](InlineResponse200271.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

