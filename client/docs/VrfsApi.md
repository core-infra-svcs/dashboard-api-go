# \VrfsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceRoutingVrfsSettings**](VrfsApi.md#GetOrganizationApplianceRoutingVrfsSettings) | **Get** /organizations/{organizationId}/appliance/routing/vrfs/settings | Return the VRF setting for an organization.
[**UpdateOrganizationApplianceRoutingVrfsSettings**](VrfsApi.md#UpdateOrganizationApplianceRoutingVrfsSettings) | **Put** /organizations/{organizationId}/appliance/routing/vrfs/settings | Update the VRF setting for an organization.



## GetOrganizationApplianceRoutingVrfsSettings

> InlineResponse200287 GetOrganizationApplianceRoutingVrfsSettings(ctx, organizationId).Execute()

Return the VRF setting for an organization.



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
    resp, r, err := apiClient.VrfsApi.GetOrganizationApplianceRoutingVrfsSettings(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VrfsApi.GetOrganizationApplianceRoutingVrfsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceRoutingVrfsSettings`: InlineResponse200287
    fmt.Fprintf(os.Stdout, "Response from `VrfsApi.GetOrganizationApplianceRoutingVrfsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceRoutingVrfsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200287**](InlineResponse200287.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceRoutingVrfsSettings

> InlineResponse200287 UpdateOrganizationApplianceRoutingVrfsSettings(ctx, organizationId).UpdateOrganizationApplianceRoutingVrfsSettings(updateOrganizationApplianceRoutingVrfsSettings).Execute()

Update the VRF setting for an organization.



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
    updateOrganizationApplianceRoutingVrfsSettings := *openapiclient.NewInlineObject257(false) // InlineObject257 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VrfsApi.UpdateOrganizationApplianceRoutingVrfsSettings(context.Background(), organizationId).UpdateOrganizationApplianceRoutingVrfsSettings(updateOrganizationApplianceRoutingVrfsSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VrfsApi.UpdateOrganizationApplianceRoutingVrfsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceRoutingVrfsSettings`: InlineResponse200287
    fmt.Fprintf(os.Stdout, "Response from `VrfsApi.UpdateOrganizationApplianceRoutingVrfsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceRoutingVrfsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceRoutingVrfsSettings** | [**InlineObject257**](InlineObject257.md) |  | 

### Return type

[**InlineResponse200287**](InlineResponse200287.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

