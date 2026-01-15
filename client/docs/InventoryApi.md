# \InventoryApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimIntoOrganizationInventory**](InventoryApi.md#ClaimIntoOrganizationInventory) | **Post** /organizations/{organizationId}/inventory/claim | Claim a list of devices, licenses, and/or orders into an organization inventory
[**ClaimOrganizationInventoryOrders**](InventoryApi.md#ClaimOrganizationInventoryOrders) | **Post** /organizations/{organizationId}/inventory/orders/claim | Claim an order by the secure unique order claim number, the order claim id
[**CreateOrganizationInventoryDevicesSwapsBulk**](InventoryApi.md#CreateOrganizationInventoryDevicesSwapsBulk) | **Post** /organizations/{organizationId}/inventory/devices/swaps/bulk | Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.
[**CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent**](InventoryApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/exportEvents | Imports event logs related to the onboarding app into elastisearch
[**CreateOrganizationInventoryOnboardingCloudMonitoringImport**](InventoryApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringImport) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.
[**CreateOrganizationInventoryOnboardingCloudMonitoringPrepare**](InventoryApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringPrepare) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare | Initiates or updates an import session
[**GetOrganizationCellularGatewayEsimsInventory**](InventoryApi.md#GetOrganizationCellularGatewayEsimsInventory) | **Get** /organizations/{organizationId}/cellularGateway/esims/inventory | The eSIM inventory of a given organization.
[**GetOrganizationInventoryDevice**](InventoryApi.md#GetOrganizationInventoryDevice) | **Get** /organizations/{organizationId}/inventory/devices/{serial} | Return a single device from the inventory of an organization
[**GetOrganizationInventoryDevices**](InventoryApi.md#GetOrganizationInventoryDevices) | **Get** /organizations/{organizationId}/inventory/devices | Return the device inventory for an organization
[**GetOrganizationInventoryDevicesSwapsBulk**](InventoryApi.md#GetOrganizationInventoryDevicesSwapsBulk) | **Get** /organizations/{organizationId}/inventory/devices/swaps/bulk/{id} | List of device swaps for a given request ID ({id}).
[**GetOrganizationInventoryOnboardingCloudMonitoringImports**](InventoryApi.md#GetOrganizationInventoryOnboardingCloudMonitoringImports) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Check the status of a committed Import operation
[**GetOrganizationInventoryOnboardingCloudMonitoringNetworks**](InventoryApi.md#GetOrganizationInventoryOnboardingCloudMonitoringNetworks) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks | Returns list of networks eligible for adding cloud monitored device
[**PreviewOrganizationInventoryOrders**](InventoryApi.md#PreviewOrganizationInventoryOrders) | **Post** /organizations/{organizationId}/inventory/orders/preview | Preview the results and status of an order claim by the secure order id
[**ReleaseFromOrganizationInventory**](InventoryApi.md#ReleaseFromOrganizationInventory) | **Post** /organizations/{organizationId}/inventory/release | Release a list of claimed devices from an organization.
[**UpdateOrganizationCellularGatewayEsimsInventory**](InventoryApi.md#UpdateOrganizationCellularGatewayEsimsInventory) | **Put** /organizations/{organizationId}/cellularGateway/esims/inventory/{id} | Toggle the status of an eSIM



## ClaimIntoOrganizationInventory

> InlineResponse200273 ClaimIntoOrganizationInventory(ctx, organizationId).ClaimIntoOrganizationInventory(claimIntoOrganizationInventory).Execute()

Claim a list of devices, licenses, and/or orders into an organization inventory



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
    claimIntoOrganizationInventory := *openapiclient.NewInlineObject273() // InlineObject273 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.ClaimIntoOrganizationInventory(context.Background(), organizationId).ClaimIntoOrganizationInventory(claimIntoOrganizationInventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ClaimIntoOrganizationInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimIntoOrganizationInventory`: InlineResponse200273
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ClaimIntoOrganizationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimIntoOrganizationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimIntoOrganizationInventory** | [**InlineObject273**](InlineObject273.md) |  | 

### Return type

[**InlineResponse200273**](InlineResponse200273.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimOrganizationInventoryOrders

> InlineResponse200310 ClaimOrganizationInventoryOrders(ctx, organizationId).ClaimOrganizationInventoryOrders(claimOrganizationInventoryOrders).Execute()

Claim an order by the secure unique order claim number, the order claim id



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
    claimOrganizationInventoryOrders := *openapiclient.NewInlineObject278("ClaimId_example") // InlineObject278 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.ClaimOrganizationInventoryOrders(context.Background(), organizationId).ClaimOrganizationInventoryOrders(claimOrganizationInventoryOrders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ClaimOrganizationInventoryOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimOrganizationInventoryOrders`: InlineResponse200310
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ClaimOrganizationInventoryOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimOrganizationInventoryOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimOrganizationInventoryOrders** | [**InlineObject278**](InlineObject278.md) |  | 

### Return type

[**InlineResponse200310**](InlineResponse200310.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryDevicesSwapsBulk

> InlineResponse207 CreateOrganizationInventoryDevicesSwapsBulk(ctx, organizationId).CreateOrganizationInventoryDevicesSwapsBulk(createOrganizationInventoryDevicesSwapsBulk).Execute()

Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.



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
    createOrganizationInventoryDevicesSwapsBulk := *openapiclient.NewInlineObject274([]openapiclient.OrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps{*openapiclient.NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkSwaps(*openapiclient.NewOrganizationsOrganizationIdInventoryDevicesSwapsBulkDevices("Old_example", "New_example"), "AfterAction_example")}) // InlineObject274 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.CreateOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId).CreateOrganizationInventoryDevicesSwapsBulk(createOrganizationInventoryDevicesSwapsBulk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryDevicesSwapsBulk`: InlineResponse207
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CreateOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryDevicesSwapsBulk** | [**InlineObject274**](InlineObject274.md) |  | 

### Return type

[**InlineResponse207**](InlineResponse207.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent

> map[string]interface{} CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(createOrganizationInventoryOnboardingCloudMonitoringExportEvent).Execute()

Imports event logs related to the onboarding app into elastisearch



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
    createOrganizationInventoryOnboardingCloudMonitoringExportEvent := *openapiclient.NewInlineObject275("LogEvent_example", int32(123)) // InlineObject275 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(createOrganizationInventoryOnboardingCloudMonitoringExportEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringExportEvent** | [**InlineObject275**](InlineObject275.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringImport

> []InlineResponse20122 CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()

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
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationInventoryOnboardingCloudMonitoringImport := *openapiclient.NewInlineObject276([]openapiclient.OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices{*openapiclient.NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringImportsDevices("DeviceId_example", "Udi_example", "NetworkId_example")}) // InlineObject276 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringImport`: []InlineResponse20122
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringImport** | [**InlineObject276**](InlineObject276.md) |  | 

### Return type

[**[]InlineResponse20122**](InlineResponse20122.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringPrepare

> []InlineResponse20123 CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare).Execute()

Initiates or updates an import session



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
    createOrganizationInventoryOnboardingCloudMonitoringPrepare := *openapiclient.NewInlineObject277([]openapiclient.OrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices{*openapiclient.NewOrganizationsOrganizationIdInventoryOnboardingCloudMonitoringPrepareDevices("Sudi_example")}) // InlineObject277 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: []InlineResponse20123
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringPrepare** | [**InlineObject277**](InlineObject277.md) |  | 

### Return type

[**[]InlineResponse20123**](InlineResponse20123.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    resp, r, err := apiClient.InventoryApi.GetOrganizationCellularGatewayEsimsInventory(context.Background(), organizationId).Eids(eids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetOrganizationCellularGatewayEsimsInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsInventory`: InlineResponse200266
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetOrganizationCellularGatewayEsimsInventory`: %v\n", resp)
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


## GetOrganizationInventoryDevice

> InlineResponse200308 GetOrganizationInventoryDevice(ctx, organizationId, serial).Execute()

Return a single device from the inventory of an organization



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetOrganizationInventoryDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevice`: InlineResponse200308
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetOrganizationInventoryDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200308**](InlineResponse200308.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevices

> []InlineResponse200308 GetOrganizationInventoryDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()

Return the device inventory for an organization



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
    usedState := "usedState_example" // string | Filter results by used or unused inventory. Accepted values are 'used' or 'unused'. (optional)
    search := "search_example" // string | Search for devices in inventory based on serial number, mac address, or model. (optional)
    macs := []string{"Inner_example"} // []string | Search for devices in inventory based on mac addresses. (optional)
    networkIds := []string{"Inner_example"} // []string | Search for devices in inventory based on network ids. Use explicit 'null' value to get available devices only. (optional)
    serials := []string{"Inner_example"} // []string | Search for devices in inventory based on serials. (optional)
    models := []string{"Inner_example"} // []string | Search for devices in inventory based on model. (optional)
    orderNumbers := []string{"Inner_example"} // []string | Search for devices in inventory based on order numbers. (optional)
    tags := []string{"Inner_example"} // []string | Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
    tagsFilterType := "tagsFilterType_example" // string | To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Filter devices by product type. Accepted values are appliance, camera, campusGateway, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetOrganizationInventoryDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetOrganizationInventoryDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevices`: []InlineResponse200308
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetOrganizationInventoryDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **usedState** | **string** | Filter results by used or unused inventory. Accepted values are &#39;used&#39; or &#39;unused&#39;. | 
 **search** | **string** | Search for devices in inventory based on serial number, mac address, or model. | 
 **macs** | **[]string** | Search for devices in inventory based on mac addresses. | 
 **networkIds** | **[]string** | Search for devices in inventory based on network ids. Use explicit &#39;null&#39; value to get available devices only. | 
 **serials** | **[]string** | Search for devices in inventory based on serials. | 
 **models** | **[]string** | Search for devices in inventory based on model. | 
 **orderNumbers** | **[]string** | Search for devices in inventory based on order numbers. | 
 **tags** | **[]string** | Filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | To use with &#39;tags&#39; parameter, to filter devices which contain ANY or ALL given tags. Accepted values are &#39;withAnyTags&#39; or &#39;withAllTags&#39;, default is &#39;withAnyTags&#39;. | 
 **productTypes** | **[]string** | Filter devices by product type. Accepted values are appliance, camera, campusGateway, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. | 

### Return type

[**[]InlineResponse200308**](InlineResponse200308.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevicesSwapsBulk

> InlineResponse207 GetOrganizationInventoryDevicesSwapsBulk(ctx, organizationId, id).Execute()

List of device swaps for a given request ID ({id}).



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryDevicesSwapsBulk`: InlineResponse207
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse207**](InlineResponse207.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringImports

> []InlineResponse200309 GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx, organizationId).ImportIds(importIds).Execute()

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
    organizationId := "organizationId_example" // string | Organization ID
    importIds := []string{"Inner_example"} // []string | import ids from an imports

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).ImportIds(importIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringImports`: []InlineResponse200309
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importIds** | **[]string** | import ids from an imports | 

### Return type

[**[]InlineResponse200309**](InlineResponse200309.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringNetworks

> []InlineResponse20050 GetOrganizationInventoryOnboardingCloudMonitoringNetworks(ctx, organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns list of networks eligible for adding cloud monitored device



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
    deviceType := "deviceType_example" // string | Device Type switch or wireless controller
    search := "search_example" // string | Optional parameter to search on network name (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: []InlineResponse20050
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceType** | **string** | Device Type switch or wireless controller | 
 **search** | **string** | Optional parameter to search on network name | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse20050**](InlineResponse20050.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewOrganizationInventoryOrders

> InlineResponse200311 PreviewOrganizationInventoryOrders(ctx, organizationId).PreviewOrganizationInventoryOrders(previewOrganizationInventoryOrders).Execute()

Preview the results and status of an order claim by the secure order id



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
    previewOrganizationInventoryOrders := *openapiclient.NewInlineObject279("ClaimId_example") // InlineObject279 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.PreviewOrganizationInventoryOrders(context.Background(), organizationId).PreviewOrganizationInventoryOrders(previewOrganizationInventoryOrders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.PreviewOrganizationInventoryOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewOrganizationInventoryOrders`: InlineResponse200311
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.PreviewOrganizationInventoryOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewOrganizationInventoryOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewOrganizationInventoryOrders** | [**InlineObject279**](InlineObject279.md) |  | 

### Return type

[**InlineResponse200311**](InlineResponse200311.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseFromOrganizationInventory

> InlineResponse200312 ReleaseFromOrganizationInventory(ctx, organizationId).ReleaseFromOrganizationInventory(releaseFromOrganizationInventory).Execute()

Release a list of claimed devices from an organization.



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
    releaseFromOrganizationInventory := *openapiclient.NewInlineObject280() // InlineObject280 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.ReleaseFromOrganizationInventory(context.Background(), organizationId).ReleaseFromOrganizationInventory(releaseFromOrganizationInventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ReleaseFromOrganizationInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseFromOrganizationInventory`: InlineResponse200312
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ReleaseFromOrganizationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseFromOrganizationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseFromOrganizationInventory** | [**InlineObject280**](InlineObject280.md) |  | 

### Return type

[**InlineResponse200312**](InlineResponse200312.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := apiClient.InventoryApi.UpdateOrganizationCellularGatewayEsimsInventory(context.Background(), organizationId, id).UpdateOrganizationCellularGatewayEsimsInventory(updateOrganizationCellularGatewayEsimsInventory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateOrganizationCellularGatewayEsimsInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCellularGatewayEsimsInventory`: InlineResponse200266Items
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.UpdateOrganizationCellularGatewayEsimsInventory`: %v\n", resp)
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

