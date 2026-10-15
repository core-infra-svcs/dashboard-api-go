# \CellularApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate**](CellularApi.md#BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate) | **Post** /organizations/{organizationId}/devices/cellular/data/profiles/assignments/batchCreate | Assign devices to a Cellular Data Management Profile in batch
[**BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete**](CellularApi.md#BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete) | **Post** /organizations/{organizationId}/devices/cellular/data/profiles/assignments/bulkDelete | Unassign devices from a Cellular Data Management Profile in batch
[**CreateDeviceCellularUplinksBandsMasksUpdate**](CellularApi.md#CreateDeviceCellularUplinksBandsMasksUpdate) | **Post** /devices/{serial}/cellular/uplinks/bands/masks/update | Update the cellular band masks for a device
[**CreateOrganizationDevicesCellularDataProfile**](CellularApi.md#CreateOrganizationDevicesCellularDataProfile) | **Post** /organizations/{organizationId}/devices/cellular/data/profiles | Add a cellular data management profile to this organization
[**DeleteOrganizationDevicesCellularDataProfile**](CellularApi.md#DeleteOrganizationDevicesCellularDataProfile) | **Delete** /organizations/{organizationId}/devices/cellular/data/profiles/{profileId} | Delete a cellular data management profile from this organization
[**GetDeviceCellularSims**](CellularApi.md#GetDeviceCellularSims) | **Get** /devices/{serial}/cellular/sims | Return the SIM and APN configurations for a cellular device.
[**GetOrganizationDevicesCellularDataDevices**](CellularApi.md#GetOrganizationDevicesCellularDataDevices) | **Get** /organizations/{organizationId}/devices/cellular/data/devices | List devices eligible for Cellular Data Management profile assignment in this organization
[**GetOrganizationDevicesCellularDataProfiles**](CellularApi.md#GetOrganizationDevicesCellularDataProfiles) | **Get** /organizations/{organizationId}/devices/cellular/data/profiles | List cellular data management profiles in this organization
[**GetOrganizationDevicesCellularDataProfilesAssignments**](CellularApi.md#GetOrganizationDevicesCellularDataProfilesAssignments) | **Get** /organizations/{organizationId}/devices/cellular/data/profiles/assignments | List Cellular Data Management Profile assignments in this organization
[**GetOrganizationDevicesCellularDataUsageByDevice**](CellularApi.md#GetOrganizationDevicesCellularDataUsageByDevice) | **Get** /organizations/{organizationId}/devices/cellular/data/usage/byDevice | List current cellular data usage for devices in this organization
[**GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval**](CellularApi.md#GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/devices/cellular/data/usage/history/byDevice/byInterval | List historical cellular data usage grouped by device and interval in this organization
[**GetOrganizationDevicesCellularGeolocations**](CellularApi.md#GetOrganizationDevicesCellularGeolocations) | **Get** /organizations/{organizationId}/devices/cellular/geolocations | List the latest cellular geolocation telemetry for devices in an organization
[**GetOrganizationDevicesCellularUplinksBandsByDevice**](CellularApi.md#GetOrganizationDevicesCellularUplinksBandsByDevice) | **Get** /organizations/{organizationId}/devices/cellular/uplinks/bands/byDevice | List the latest cellular uplink signal information for devices in an organization
[**GetOrganizationDevicesCellularUplinksTowersByDevice**](CellularApi.md#GetOrganizationDevicesCellularUplinksTowersByDevice) | **Get** /organizations/{organizationId}/devices/cellular/uplinks/towers/byDevice | List the latest cellular tower information for devices in an organization
[**UpdateDeviceCellularGeolocations**](CellularApi.md#UpdateDeviceCellularGeolocations) | **Put** /devices/{serial}/cellular/geolocations | Update the enablement of the geolocation feature for a device
[**UpdateDeviceCellularSims**](CellularApi.md#UpdateDeviceCellularSims) | **Put** /devices/{serial}/cellular/sims | Updates the SIM and APN configurations for a cellular device.
[**UpdateOrganizationDevicesCellularDataProfile**](CellularApi.md#UpdateOrganizationDevicesCellularDataProfile) | **Put** /organizations/{organizationId}/devices/cellular/data/profiles/{profileId} | Update a Cellular Data Management Profile



## BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate

> InlineResponse200365 BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(ctx, organizationId).BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(batchOrganizationDevicesCellularDataProfilesAssignmentsCreate).Execute()

Assign devices to a Cellular Data Management Profile in batch



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
    batchOrganizationDevicesCellularDataProfilesAssignmentsCreate := *openapiclient.NewInlineObject296([]openapiclient.OrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateItems{*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateItems(*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateProfile("Id_example"), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBatchCreateDevice("Serial_example"))}) // InlineObject296 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(context.Background(), organizationId).BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate(batchOrganizationDevicesCellularDataProfilesAssignmentsCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate`: InlineResponse200365
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.BatchOrganizationDevicesCellularDataProfilesAssignmentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchOrganizationDevicesCellularDataProfilesAssignmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchOrganizationDevicesCellularDataProfilesAssignmentsCreate** | [**InlineObject296**](InlineObject296.md) |  | 

### Return type

[**InlineResponse200365**](InlineResponse200365.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete

> BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(ctx, organizationId).BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete).Execute()

Unassign devices from a Cellular Data Management Profile in batch



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
    bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete := *openapiclient.NewInlineObject297([]openapiclient.OrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteItems{*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteItems(*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteProfile("Id_example"), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesAssignmentsBulkDeleteDevice("Serial_example"))}) // InlineObject297 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(context.Background(), organizationId).BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete(bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.BulkOrganizationDevicesCellularDataProfilesAssignmentsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkOrganizationDevicesCellularDataProfilesAssignmentsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkOrganizationDevicesCellularDataProfilesAssignmentsDelete** | [**InlineObject297**](InlineObject297.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceCellularUplinksBandsMasksUpdate

> InlineResponse20028 CreateDeviceCellularUplinksBandsMasksUpdate(ctx, serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()

Update the cellular band masks for a device



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
    createDeviceCellularUplinksBandsMasksUpdate := *openapiclient.NewInlineObject17("Slot_example", "Type_example", []string{"Masked_example"}) // InlineObject17 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.CreateDeviceCellularUplinksBandsMasksUpdate(context.Background(), serial).CreateDeviceCellularUplinksBandsMasksUpdate(createDeviceCellularUplinksBandsMasksUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.CreateDeviceCellularUplinksBandsMasksUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceCellularUplinksBandsMasksUpdate`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.CreateDeviceCellularUplinksBandsMasksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceCellularUplinksBandsMasksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceCellularUplinksBandsMasksUpdate** | [**InlineObject17**](InlineObject17.md) |  | 

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationDevicesCellularDataProfile

> InlineResponse200363 CreateOrganizationDevicesCellularDataProfile(ctx, organizationId).CreateOrganizationDevicesCellularDataProfile(createOrganizationDevicesCellularDataProfile).Execute()

Add a cellular data management profile to this organization



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
    createOrganizationDevicesCellularDataProfile := *openapiclient.NewInlineObject295("Name_example", "Description_example", []openapiclient.OrganizationsOrganizationIdDevicesCellularDataProfilesRules{*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesRules("Slot_example", *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesUplink(int32(123), false), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesCap(int32(123), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm("Resets_example", *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesCapTermStarts())))}) // InlineObject295 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.CreateOrganizationDevicesCellularDataProfile(context.Background(), organizationId).CreateOrganizationDevicesCellularDataProfile(createOrganizationDevicesCellularDataProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.CreateOrganizationDevicesCellularDataProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationDevicesCellularDataProfile`: InlineResponse200363
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.CreateOrganizationDevicesCellularDataProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationDevicesCellularDataProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationDevicesCellularDataProfile** | [**InlineObject295**](InlineObject295.md) |  | 

### Return type

[**InlineResponse200363**](InlineResponse200363.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationDevicesCellularDataProfile

> DeleteOrganizationDevicesCellularDataProfile(ctx, organizationId, profileId).Execute()

Delete a cellular data management profile from this organization



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
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.DeleteOrganizationDevicesCellularDataProfile(context.Background(), organizationId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.DeleteOrganizationDevicesCellularDataProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationDevicesCellularDataProfileRequest struct via the builder pattern


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


## GetDeviceCellularSims

> InlineResponse20027 GetDeviceCellularSims(ctx, serial).Execute()

Return the SIM and APN configurations for a cellular device.



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
    resp, r, err := apiClient.CellularApi.GetDeviceCellularSims(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetDeviceCellularSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularSims`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataDevices

> InlineResponse200361 GetOrganizationDevicesCellularDataDevices(ctx, organizationId).IncludeAssigned(includeAssigned).IncludedSerials(includedSerials).ExcludedSerials(excludedSerials).IncludedProfileIds(includedProfileIds).ExcludedProfileIds(excludedProfileIds).DeviceTypes(deviceTypes).Slots(slots).Name(name).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List devices eligible for Cellular Data Management profile assignment in this organization



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
    includeAssigned := true // bool | Whether to include devices that have already been assigned to a Cellular Data Management Profile (optional)
    includedSerials := []string{"Inner_example"} // []string | List of device serials to force-include in the response when the devices would otherwise be filtered out. This override is primarily useful for keeping selected devices visible while paging through results. Maximum 1000 serials. (optional)
    excludedSerials := []string{"Inner_example"} // []string | List of device serials to force-exclude from the response when the devices would otherwise be returned. This override is primarily useful for hiding selected devices while paging through results. Maximum 1000 serials. (optional)
    includedProfileIds := []string{"Inner_example"} // []string | List of Cellular Data Management Profile IDs to include in the results. Maximum 1000 profile IDs. (optional)
    excludedProfileIds := []string{"Inner_example"} // []string | List of Cellular Data Management Profile IDs to exclude from the results. Maximum 1000 profile IDs. (optional)
    deviceTypes := []string{"Inner_example"} // []string | List of device types to filter by. Maximum 1000 device types. (optional)
    slots := []string{"Slots_example"} // []string | List of SIM slot types that devices must support. Accepted values are sim1, sim2, and esim. Maximum 3 slots. (optional)
    name := "name_example" // string | Name of the device to filter by (partial matches allowed) (optional)
    serials := []string{"Inner_example"} // []string | List of device serials to filter by. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularDataDevices(context.Background(), organizationId).IncludeAssigned(includeAssigned).IncludedSerials(includedSerials).ExcludedSerials(excludedSerials).IncludedProfileIds(includedProfileIds).ExcludedProfileIds(excludedProfileIds).DeviceTypes(deviceTypes).Slots(slots).Name(name).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularDataDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataDevices`: InlineResponse200361
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularDataDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAssigned** | **bool** | Whether to include devices that have already been assigned to a Cellular Data Management Profile | 
 **includedSerials** | **[]string** | List of device serials to force-include in the response when the devices would otherwise be filtered out. This override is primarily useful for keeping selected devices visible while paging through results. Maximum 1000 serials. | 
 **excludedSerials** | **[]string** | List of device serials to force-exclude from the response when the devices would otherwise be returned. This override is primarily useful for hiding selected devices while paging through results. Maximum 1000 serials. | 
 **includedProfileIds** | **[]string** | List of Cellular Data Management Profile IDs to include in the results. Maximum 1000 profile IDs. | 
 **excludedProfileIds** | **[]string** | List of Cellular Data Management Profile IDs to exclude from the results. Maximum 1000 profile IDs. | 
 **deviceTypes** | **[]string** | List of device types to filter by. Maximum 1000 device types. | 
 **slots** | **[]string** | List of SIM slot types that devices must support. Accepted values are sim1, sim2, and esim. Maximum 3 slots. | 
 **name** | **string** | Name of the device to filter by (partial matches allowed) | 
 **serials** | **[]string** | List of device serials to filter by. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200361**](InlineResponse200361.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataProfiles

> InlineResponse200362 GetOrganizationDevicesCellularDataProfiles(ctx, organizationId).ProfileIds(profileIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List cellular data management profiles in this organization



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
    profileIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by Data Management Profile ID. (optional)
    serials := []string{"Inner_example"} // []string | Devices to find Cellular Data Management Profiles for. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularDataProfiles(context.Background(), organizationId).ProfileIds(profileIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularDataProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataProfiles`: InlineResponse200362
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularDataProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to filter the results by Data Management Profile ID. | 
 **serials** | **[]string** | Devices to find Cellular Data Management Profiles for. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200362**](InlineResponse200362.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataProfilesAssignments

> InlineResponse200364 GetOrganizationDevicesCellularDataProfilesAssignments(ctx, organizationId).ProfileIds(profileIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List Cellular Data Management Profile assignments in this organization



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
    profileIds := []string{"Inner_example"} // []string | Optional parameter to find assignments by Profile IDs. Maximum 1000 profile IDs. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to find assignments by Device Serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularDataProfilesAssignments(context.Background(), organizationId).ProfileIds(profileIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularDataProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataProfilesAssignments`: InlineResponse200364
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularDataProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profileIds** | **[]string** | Optional parameter to find assignments by Profile IDs. Maximum 1000 profile IDs. | 
 **serials** | **[]string** | Optional parameter to find assignments by Device Serials. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200364**](InlineResponse200364.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataUsageByDevice

> InlineResponse200367 GetOrganizationDevicesCellularDataUsageByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List current cellular data usage for devices in this organization



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
    serials := []string{"Inner_example"} // []string | Filter the results by device serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularDataUsageByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularDataUsageByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataUsageByDevice`: InlineResponse200367
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularDataUsageByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataUsageByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Filter the results by device serials. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200367**](InlineResponse200367.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval

> InlineResponse200368 GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

List historical cellular data usage grouped by device and interval in this organization



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
    serials := []string{"Inner_example"} // []string | Required parameter to filter the results by device serials. Maximum 10 serials.
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 10. Default is 5. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 366 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 86400. Interval is calculated if time params are provided. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval`: InlineResponse200368
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularDataUsageHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularDataUsageHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Required parameter to filter the results by device serials. Maximum 10 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 10. Default is 5. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 366 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 1200, 14400, 86400. The default is 86400. Interval is calculated if time params are provided. | 

### Return type

[**InlineResponse200368**](InlineResponse200368.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularGeolocations

> InlineResponse200369 GetOrganizationDevicesCellularGeolocations(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the latest cellular geolocation telemetry for devices in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter the results by device serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularGeolocations(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularGeolocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularGeolocations`: InlineResponse200369
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularGeolocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularGeolocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter the results by device serials. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200369**](InlineResponse200369.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularUplinksBandsByDevice

> InlineResponse200370 GetOrganizationDevicesCellularUplinksBandsByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the latest cellular uplink signal information for devices in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter the results by device serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularUplinksBandsByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularUplinksBandsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularUplinksBandsByDevice`: InlineResponse200370
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularUplinksBandsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularUplinksBandsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter the results by device serials. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200370**](InlineResponse200370.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesCellularUplinksTowersByDevice

> InlineResponse200371 GetOrganizationDevicesCellularUplinksTowersByDevice(ctx, organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the latest cellular tower information for devices in an organization



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter the results by device serials. Maximum 1000 serials. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.GetOrganizationDevicesCellularUplinksTowersByDevice(context.Background(), organizationId).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.GetOrganizationDevicesCellularUplinksTowersByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesCellularUplinksTowersByDevice`: InlineResponse200371
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.GetOrganizationDevicesCellularUplinksTowersByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesCellularUplinksTowersByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter the results by device serials. Maximum 1000 serials. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200371**](InlineResponse200371.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGeolocations

> InlineResponse20026 UpdateDeviceCellularGeolocations(ctx, serial).UpdateDeviceCellularGeolocations(updateDeviceCellularGeolocations).Execute()

Update the enablement of the geolocation feature for a device



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
    updateDeviceCellularGeolocations := *openapiclient.NewInlineObject15(false) // InlineObject15 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.UpdateDeviceCellularGeolocations(context.Background(), serial).UpdateDeviceCellularGeolocations(updateDeviceCellularGeolocations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.UpdateDeviceCellularGeolocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularGeolocations`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.UpdateDeviceCellularGeolocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGeolocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGeolocations** | [**InlineObject15**](InlineObject15.md) |  | 

### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularSims

> InlineResponse20027 UpdateDeviceCellularSims(ctx, serial).UpdateDeviceCellularSims(updateDeviceCellularSims).Execute()

Updates the SIM and APN configurations for a cellular device.



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
    updateDeviceCellularSims := *openapiclient.NewInlineObject16() // InlineObject16 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.UpdateDeviceCellularSims(context.Background(), serial).UpdateDeviceCellularSims(updateDeviceCellularSims).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.UpdateDeviceCellularSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularSims`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.UpdateDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularSims** | [**InlineObject16**](InlineObject16.md) |  | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationDevicesCellularDataProfile

> InlineResponse200366 UpdateOrganizationDevicesCellularDataProfile(ctx, organizationId, profileId).UpdateOrganizationDevicesCellularDataProfile(updateOrganizationDevicesCellularDataProfile).Execute()

Update a Cellular Data Management Profile



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
    profileId := "profileId_example" // string | Profile ID
    updateOrganizationDevicesCellularDataProfile := *openapiclient.NewInlineObject298([]openapiclient.OrganizationsOrganizationIdDevicesCellularDataProfilesRules{*openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesRules("Slot_example", *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesUplink(int32(123), false), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesCap(int32(123), *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesCapTerm("Resets_example", *openapiclient.NewOrganizationsOrganizationIdDevicesCellularDataProfilesCapTermStarts())))}) // InlineObject298 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularApi.UpdateOrganizationDevicesCellularDataProfile(context.Background(), organizationId, profileId).UpdateOrganizationDevicesCellularDataProfile(updateOrganizationDevicesCellularDataProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularApi.UpdateOrganizationDevicesCellularDataProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationDevicesCellularDataProfile`: InlineResponse200366
    fmt.Fprintf(os.Stdout, "Response from `CellularApi.UpdateOrganizationDevicesCellularDataProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationDevicesCellularDataProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationDevicesCellularDataProfile** | [**InlineObject298**](InlineObject298.md) |  | 

### Return type

[**InlineResponse200366**](InlineResponse200366.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

