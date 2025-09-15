# \CellularGatewayApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCellularGatewayEsimsServiceProvidersAccount**](CellularGatewayApi.md#CreateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Post** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Add a service provider account.
[**CreateOrganizationCellularGatewayEsimsSwap**](CellularGatewayApi.md#CreateOrganizationCellularGatewayEsimsSwap) | **Post** /organizations/{organizationId}/cellularGateway/esims/swap | Swap which profile an eSIM uses.
[**DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount**](CellularGatewayApi.md#DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Delete** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Remove a service provider account&#39;s integration with the Dashboard.
[**GetDeviceCellularGatewayLan**](CellularGatewayApi.md#GetDeviceCellularGatewayLan) | **Get** /devices/{serial}/cellularGateway/lan | Show the LAN Settings of a MG
[**GetDeviceCellularGatewayPortForwardingRules**](CellularGatewayApi.md#GetDeviceCellularGatewayPortForwardingRules) | **Get** /devices/{serial}/cellularGateway/portForwardingRules | Returns the port forwarding rules for a single MG.
[**GetNetworkCellularGatewayConnectivityMonitoringDestinations**](CellularGatewayApi.md#GetNetworkCellularGatewayConnectivityMonitoringDestinations) | **Get** /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations | Return the connectivity testing destinations for an MG network
[**GetNetworkCellularGatewayDhcp**](CellularGatewayApi.md#GetNetworkCellularGatewayDhcp) | **Get** /networks/{networkId}/cellularGateway/dhcp | List common DHCP settings of MGs
[**GetNetworkCellularGatewaySubnetPool**](CellularGatewayApi.md#GetNetworkCellularGatewaySubnetPool) | **Get** /networks/{networkId}/cellularGateway/subnetPool | Return the subnet pool and mask configured for MGs in the network.
[**GetNetworkCellularGatewayUplink**](CellularGatewayApi.md#GetNetworkCellularGatewayUplink) | **Get** /networks/{networkId}/cellularGateway/uplink | Returns the uplink settings for your MG network.
[**GetOrganizationCellularGatewayEsimsInventory**](CellularGatewayApi.md#GetOrganizationCellularGatewayEsimsInventory) | **Get** /organizations/{organizationId}/cellularGateway/esims/inventory | The eSIM inventory of a given organization.
[**GetOrganizationCellularGatewayEsimsServiceProviders**](CellularGatewayApi.md#GetOrganizationCellularGatewayEsimsServiceProviders) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders | Service providers customers can add accounts for.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccounts**](CellularGatewayApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccounts) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts | Inventory of service provider accounts tied to the organization.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans**](CellularGatewayApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/communicationPlans | The communication plans available for a given provider.
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans**](CellularGatewayApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/ratePlans | The rate plans available for a given provider.
[**GetOrganizationCellularGatewayUplinkStatuses**](CellularGatewayApi.md#GetOrganizationCellularGatewayUplinkStatuses) | **Get** /organizations/{organizationId}/cellularGateway/uplink/statuses | List the uplink status of every Meraki MG cellular gateway in the organization
[**UpdateDeviceCellularGatewayLan**](CellularGatewayApi.md#UpdateDeviceCellularGatewayLan) | **Put** /devices/{serial}/cellularGateway/lan | Update the LAN Settings for a single MG.
[**UpdateDeviceCellularGatewayPortForwardingRules**](CellularGatewayApi.md#UpdateDeviceCellularGatewayPortForwardingRules) | **Put** /devices/{serial}/cellularGateway/portForwardingRules | Updates the port forwarding rules for a single MG.
[**UpdateNetworkCellularGatewayConnectivityMonitoringDestinations**](CellularGatewayApi.md#UpdateNetworkCellularGatewayConnectivityMonitoringDestinations) | **Put** /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations | Update the connectivity testing destinations for an MG network
[**UpdateNetworkCellularGatewayDhcp**](CellularGatewayApi.md#UpdateNetworkCellularGatewayDhcp) | **Put** /networks/{networkId}/cellularGateway/dhcp | Update common DHCP settings of MGs
[**UpdateNetworkCellularGatewaySubnetPool**](CellularGatewayApi.md#UpdateNetworkCellularGatewaySubnetPool) | **Put** /networks/{networkId}/cellularGateway/subnetPool | Update the subnet pool and mask configuration for MGs in the network.
[**UpdateNetworkCellularGatewayUplink**](CellularGatewayApi.md#UpdateNetworkCellularGatewayUplink) | **Put** /networks/{networkId}/cellularGateway/uplink | Updates the uplink settings for your MG network.
[**UpdateOrganizationCellularGatewayEsimsInventory**](CellularGatewayApi.md#UpdateOrganizationCellularGatewayEsimsInventory) | **Put** /organizations/{organizationId}/cellularGateway/esims/inventory/{id} | Toggle the status of an eSIM
[**UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount**](CellularGatewayApi.md#UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount) | **Put** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/{accountId} | Edit service provider account info stored in Meraki&#39;s database.
[**UpdateOrganizationCellularGatewayEsimsSwap**](CellularGatewayApi.md#UpdateOrganizationCellularGatewayEsimsSwap) | **Put** /organizations/{organizationId}/cellularGateway/esims/swap/{id} | Get the status of a profile swap.



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
    createOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject248("AccountId_example", "ApiKey_example", *openapiclient.NewOrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsServiceProvider1("Name_example"), "Title_example", "Username_example") // InlineObject248 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId).CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(createOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
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

 **createOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject248**](InlineObject248.md) |  | 

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

> InlineResponse200263 CreateOrganizationCellularGatewayEsimsSwap(ctx, organizationId).CreateOrganizationCellularGatewayEsimsSwap(createOrganizationCellularGatewayEsimsSwap).Execute()

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
    createOrganizationCellularGatewayEsimsSwap := *openapiclient.NewInlineObject250([]openapiclient.OrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps{*openapiclient.NewOrganizationsOrganizationIdCellularGatewayEsimsSwapSwaps("Eid_example")}) // InlineObject250 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.CreateOrganizationCellularGatewayEsimsSwap(context.Background(), organizationId).CreateOrganizationCellularGatewayEsimsSwap(createOrganizationCellularGatewayEsimsSwap).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.CreateOrganizationCellularGatewayEsimsSwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCellularGatewayEsimsSwap`: InlineResponse200263
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.CreateOrganizationCellularGatewayEsimsSwap`: %v\n", resp)
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

 **createOrganizationCellularGatewayEsimsSwap** | [**InlineObject250**](InlineObject250.md) |  | 

### Return type

[**InlineResponse200263**](InlineResponse200263.md)

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
    resp, r, err := apiClient.CellularGatewayApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
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


## GetDeviceCellularGatewayLan

> InlineResponse20019 GetDeviceCellularGatewayLan(ctx, serial).Execute()

Show the LAN Settings of a MG



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.GetDeviceCellularGatewayLan(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetDeviceCellularGatewayLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularGatewayLan`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCellularGatewayPortForwardingRules

> InlineResponse20020 GetDeviceCellularGatewayPortForwardingRules(ctx, serial).Execute()

Returns the port forwarding rules for a single MG.



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.GetDeviceCellularGatewayPortForwardingRules(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetDeviceCellularGatewayPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularGatewayPortForwardingRules`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewayConnectivityMonitoringDestinations

> InlineResponse20053 GetNetworkCellularGatewayConnectivityMonitoringDestinations(ctx, networkId).Execute()

Return the connectivity testing destinations for an MG network



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.GetNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetNetworkCellularGatewayConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCellularGatewayConnectivityMonitoringDestinations`: InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetNetworkCellularGatewayConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20053**](InlineResponse20053.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewayDhcp

> InlineResponse20085 GetNetworkCellularGatewayDhcp(ctx, networkId).Execute()

List common DHCP settings of MGs



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.GetNetworkCellularGatewayDhcp(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetNetworkCellularGatewayDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCellularGatewayDhcp`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetNetworkCellularGatewayDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewaySubnetPool

> InlineResponse20086 GetNetworkCellularGatewaySubnetPool(ctx, networkId).Execute()

Return the subnet pool and mask configured for MGs in the network.



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.GetNetworkCellularGatewaySubnetPool(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetNetworkCellularGatewaySubnetPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCellularGatewaySubnetPool`: InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20086**](InlineResponse20086.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewayUplink

> InlineResponse20087 GetNetworkCellularGatewayUplink(ctx, networkId).Execute()

Returns the uplink settings for your MG network.



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.GetNetworkCellularGatewayUplink(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetNetworkCellularGatewayUplink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCellularGatewayUplink`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetNetworkCellularGatewayUplink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayUplinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsInventory

> InlineResponse200258 GetOrganizationCellularGatewayEsimsInventory(ctx, organizationId).Eids(eids).Execute()

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
    resp, r, err := apiClient.CellularGatewayApi.GetOrganizationCellularGatewayEsimsInventory(context.Background(), organizationId).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetOrganizationCellularGatewayEsimsInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsInventory`: InlineResponse200258
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetOrganizationCellularGatewayEsimsInventory`: %v\n", resp)
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

[**InlineResponse200258**](InlineResponse200258.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProviders

> InlineResponse200259 GetOrganizationCellularGatewayEsimsServiceProviders(ctx, organizationId).Execute()

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
    resp, r, err := apiClient.CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProviders(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProviders`: InlineResponse200259
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProviders`: %v\n", resp)
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

[**InlineResponse200259**](InlineResponse200259.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccounts

> []InlineResponse200260 GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: []InlineResponse200260
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts`: %v\n", resp)
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

[**[]InlineResponse200260**](InlineResponse200260.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans

> InlineResponse200261 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: InlineResponse200261
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: %v\n", resp)
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

[**InlineResponse200261**](InlineResponse200261.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans

> InlineResponse200262 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(ctx, organizationId).AccountIds(accountIds).Execute()

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
    resp, r, err := apiClient.CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: InlineResponse200262
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsRatePlans`: %v\n", resp)
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

[**InlineResponse200262**](InlineResponse200262.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayUplinkStatuses

> []InlineResponse200264 GetOrganizationCellularGatewayUplinkStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()

List the uplink status of every Meraki MG cellular gateway in the organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned devices will be filtered to only include these networks. (optional)
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
    iccids := []string{"Inner_example"} // []string | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.GetOrganizationCellularGatewayUplinkStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.GetOrganizationCellularGatewayUplinkStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayUplinkStatuses`: []InlineResponse200264
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.GetOrganizationCellularGatewayUplinkStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayUplinkStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned devices will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **iccids** | **[]string** | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. | 

### Return type

[**[]InlineResponse200264**](InlineResponse200264.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGatewayLan

> InlineResponse20019 UpdateDeviceCellularGatewayLan(ctx, serial).UpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan).Execute()

Update the LAN Settings for a single MG.



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
    serial := "serial_example" // string | Serial
    updateDeviceCellularGatewayLan := *openapiclient.NewInlineObject14() // InlineObject14 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateDeviceCellularGatewayLan(context.Background(), serial).UpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateDeviceCellularGatewayLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularGatewayLan`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayLan** | [**InlineObject14**](InlineObject14.md) |  | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGatewayPortForwardingRules

> InlineResponse20020 UpdateDeviceCellularGatewayPortForwardingRules(ctx, serial).UpdateDeviceCellularGatewayPortForwardingRules(updateDeviceCellularGatewayPortForwardingRules).Execute()

Updates the port forwarding rules for a single MG.



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
    serial := "serial_example" // string | Serial
    updateDeviceCellularGatewayPortForwardingRules := *openapiclient.NewInlineObject15() // InlineObject15 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateDeviceCellularGatewayPortForwardingRules(context.Background(), serial).UpdateDeviceCellularGatewayPortForwardingRules(updateDeviceCellularGatewayPortForwardingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateDeviceCellularGatewayPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularGatewayPortForwardingRules`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayPortForwardingRules** | [**InlineObject15**](InlineObject15.md) |  | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewayConnectivityMonitoringDestinations

> InlineResponse20053 UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(ctx, networkId).UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(updateNetworkCellularGatewayConnectivityMonitoringDestinations).Execute()

Update the connectivity testing destinations for an MG network



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkCellularGatewayConnectivityMonitoringDestinations := *openapiclient.NewInlineObject88() // InlineObject88 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(updateNetworkCellularGatewayConnectivityMonitoringDestinations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCellularGatewayConnectivityMonitoringDestinations`: InlineResponse20053
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayConnectivityMonitoringDestinations** | [**InlineObject88**](InlineObject88.md) |  | 

### Return type

[**InlineResponse20053**](InlineResponse20053.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewayDhcp

> InlineResponse20085 UpdateNetworkCellularGatewayDhcp(ctx, networkId).UpdateNetworkCellularGatewayDhcp(updateNetworkCellularGatewayDhcp).Execute()

Update common DHCP settings of MGs



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkCellularGatewayDhcp := *openapiclient.NewInlineObject89() // InlineObject89 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateNetworkCellularGatewayDhcp(context.Background(), networkId).UpdateNetworkCellularGatewayDhcp(updateNetworkCellularGatewayDhcp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateNetworkCellularGatewayDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCellularGatewayDhcp`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateNetworkCellularGatewayDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayDhcp** | [**InlineObject89**](InlineObject89.md) |  | 

### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewaySubnetPool

> InlineResponse20086 UpdateNetworkCellularGatewaySubnetPool(ctx, networkId).UpdateNetworkCellularGatewaySubnetPool(updateNetworkCellularGatewaySubnetPool).Execute()

Update the subnet pool and mask configuration for MGs in the network.



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkCellularGatewaySubnetPool := *openapiclient.NewInlineObject90() // InlineObject90 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateNetworkCellularGatewaySubnetPool(context.Background(), networkId).UpdateNetworkCellularGatewaySubnetPool(updateNetworkCellularGatewaySubnetPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateNetworkCellularGatewaySubnetPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCellularGatewaySubnetPool`: InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewaySubnetPool** | [**InlineObject90**](InlineObject90.md) |  | 

### Return type

[**InlineResponse20086**](InlineResponse20086.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewayUplink

> InlineResponse20087 UpdateNetworkCellularGatewayUplink(ctx, networkId).UpdateNetworkCellularGatewayUplink(updateNetworkCellularGatewayUplink).Execute()

Updates the uplink settings for your MG network.



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkCellularGatewayUplink := *openapiclient.NewInlineObject91() // InlineObject91 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateNetworkCellularGatewayUplink(context.Background(), networkId).UpdateNetworkCellularGatewayUplink(updateNetworkCellularGatewayUplink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateNetworkCellularGatewayUplink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCellularGatewayUplink`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateNetworkCellularGatewayUplink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayUplinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayUplink** | [**InlineObject91**](InlineObject91.md) |  | 

### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCellularGatewayEsimsInventory

> InlineResponse200258Items UpdateOrganizationCellularGatewayEsimsInventory(ctx, organizationId, id).UpdateOrganizationCellularGatewayEsimsInventory(updateOrganizationCellularGatewayEsimsInventory).Execute()

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
    updateOrganizationCellularGatewayEsimsInventory := *openapiclient.NewInlineObject247() // InlineObject247 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsInventory(context.Background(), organizationId, id).UpdateOrganizationCellularGatewayEsimsInventory(updateOrganizationCellularGatewayEsimsInventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsInventory`: InlineResponse200258Items
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsInventory`: %v\n", resp)
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


 **updateOrganizationCellularGatewayEsimsInventory** | [**InlineObject247**](InlineObject247.md) |  | 

### Return type

[**InlineResponse200258Items**](InlineResponse200258Items.md)

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
    updateOrganizationCellularGatewayEsimsServiceProvidersAccount := *openapiclient.NewInlineObject249() // InlineObject249 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(context.Background(), organizationId, accountId).UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(updateOrganizationCellularGatewayEsimsServiceProvidersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: OrganizationsOrganizationIdCellularGatewayEsimsServiceProvidersAccountsItems
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount`: %v\n", resp)
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


 **updateOrganizationCellularGatewayEsimsServiceProvidersAccount** | [**InlineObject249**](InlineObject249.md) |  | 

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

> InlineResponse200263 UpdateOrganizationCellularGatewayEsimsSwap(ctx, id, organizationId).Execute()

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
    resp, r, err := apiClient.CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsSwap(context.Background(), id, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsSwap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsSwap`: InlineResponse200263
    fmt.Fprintf(os.Stdout, "Response from `CellularGatewayApi.UpdateOrganizationCellularGatewayEsimsSwap`: %v\n", resp)
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

[**InlineResponse200263**](InlineResponse200263.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

