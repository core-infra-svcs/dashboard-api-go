# \VlansApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkApplianceVlan**](VlansApi.md#CreateNetworkApplianceVlan) | **Post** /networks/{networkId}/appliance/vlans | Add a VLAN
[**DeleteNetworkApplianceVlan**](VlansApi.md#DeleteNetworkApplianceVlan) | **Delete** /networks/{networkId}/appliance/vlans/{vlanId} | Delete a VLAN from a network
[**GetNetworkApplianceVlan**](VlansApi.md#GetNetworkApplianceVlan) | **Get** /networks/{networkId}/appliance/vlans/{vlanId} | Return a VLAN
[**GetNetworkApplianceVlans**](VlansApi.md#GetNetworkApplianceVlans) | **Get** /networks/{networkId}/appliance/vlans | List the VLANs for an MX network
[**GetNetworkApplianceVlansSettings**](VlansApi.md#GetNetworkApplianceVlansSettings) | **Get** /networks/{networkId}/appliance/vlans/settings | Returns the enabled status of VLANs for the network
[**UpdateNetworkApplianceVlan**](VlansApi.md#UpdateNetworkApplianceVlan) | **Put** /networks/{networkId}/appliance/vlans/{vlanId} | Update a VLAN
[**UpdateNetworkApplianceVlansSettings**](VlansApi.md#UpdateNetworkApplianceVlansSettings) | **Put** /networks/{networkId}/appliance/vlans/settings | Enable/Disable VLANs for the given network



## CreateNetworkApplianceVlan

> InlineResponse2019 CreateNetworkApplianceVlan(ctx, networkId).CreateNetworkApplianceVlan(createNetworkApplianceVlan).Execute()

Add a VLAN



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
    createNetworkApplianceVlan := *openapiclient.NewInlineObject74("Id_example", "Name_example") // InlineObject74 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlansApi.CreateNetworkApplianceVlan(context.Background(), networkId).CreateNetworkApplianceVlan(createNetworkApplianceVlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlansApi.CreateNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceVlan`: InlineResponse2019
    fmt.Fprintf(os.Stdout, "Response from `VlansApi.CreateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceVlan** | [**InlineObject74**](InlineObject74.md) |  | 

### Return type

[**InlineResponse2019**](InlineResponse2019.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceVlan

> DeleteNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Delete a VLAN from a network



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
    vlanId := "vlanId_example" // string | Vlan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlansApi.DeleteNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlansApi.DeleteNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceVlanRequest struct via the builder pattern


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


## GetNetworkApplianceVlan

> InlineResponse20073 GetNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Return a VLAN



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
    vlanId := "vlanId_example" // string | Vlan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlansApi.GetNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlansApi.GetNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlan`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `VlansApi.GetNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlans

> []InlineResponse20073 GetNetworkApplianceVlans(ctx, networkId).Execute()

List the VLANs for an MX network



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
    resp, r, err := apiClient.VlansApi.GetNetworkApplianceVlans(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlansApi.GetNetworkApplianceVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlans`: []InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `VlansApi.GetNetworkApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlansSettings

> InlineResponse20074 GetNetworkApplianceVlansSettings(ctx, networkId).Execute()

Returns the enabled status of VLANs for the network



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
    resp, r, err := apiClient.VlansApi.GetNetworkApplianceVlansSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlansApi.GetNetworkApplianceVlansSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceVlansSettings`: InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `VlansApi.GetNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20074**](InlineResponse20074.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlan

> InlineResponse20073 UpdateNetworkApplianceVlan(ctx, networkId, vlanId).UpdateNetworkApplianceVlan(updateNetworkApplianceVlan).Execute()

Update a VLAN



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
    vlanId := "vlanId_example" // string | Vlan ID
    updateNetworkApplianceVlan := *openapiclient.NewInlineObject76() // InlineObject76 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlansApi.UpdateNetworkApplianceVlan(context.Background(), networkId, vlanId).UpdateNetworkApplianceVlan(updateNetworkApplianceVlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlansApi.UpdateNetworkApplianceVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVlan`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `VlansApi.UpdateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceVlan** | [**InlineObject76**](InlineObject76.md) |  | 

### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlansSettings

> InlineResponse20074 UpdateNetworkApplianceVlansSettings(ctx, networkId).UpdateNetworkApplianceVlansSettings(updateNetworkApplianceVlansSettings).Execute()

Enable/Disable VLANs for the given network



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
    updateNetworkApplianceVlansSettings := *openapiclient.NewInlineObject75() // InlineObject75 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlansApi.UpdateNetworkApplianceVlansSettings(context.Background(), networkId).UpdateNetworkApplianceVlansSettings(updateNetworkApplianceVlansSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlansApi.UpdateNetworkApplianceVlansSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceVlansSettings`: InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `VlansApi.UpdateNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVlansSettings** | [**InlineObject75**](InlineObject75.md) |  | 

### Return type

[**InlineResponse20074**](InlineResponse20074.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

