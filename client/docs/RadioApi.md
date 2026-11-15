# \RadioApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceApplianceRadioSettings**](RadioApi.md#GetDeviceApplianceRadioSettings) | **Get** /devices/{serial}/appliance/radio/settings | Return the radio settings of an appliance
[**GetDeviceWirelessRadioOverrides**](RadioApi.md#GetDeviceWirelessRadioOverrides) | **Get** /devices/{serial}/wireless/radio/overrides | Return the radio overrides of a device
[**GetDeviceWirelessRadioSettings**](RadioApi.md#GetDeviceWirelessRadioSettings) | **Get** /devices/{serial}/wireless/radio/settings | Return the manually configured radio settings overrides of a device, which take precedence over RF profiles.
[**GetOrganizationWirelessRadioOverridesByDevice**](RadioApi.md#GetOrganizationWirelessRadioOverridesByDevice) | **Get** /organizations/{organizationId}/wireless/radio/overrides/byDevice | Return a list of radio overrides
[**GetOrganizationWirelessRadioRrmByNetwork**](RadioApi.md#GetOrganizationWirelessRadioRrmByNetwork) | **Get** /organizations/{organizationId}/wireless/radio/rrm/byNetwork | List the AutoRF settings of an organization by network
[**RecalculateOrganizationWirelessRadioAutoRfChannels**](RadioApi.md#RecalculateOrganizationWirelessRadioAutoRfChannels) | **Post** /organizations/{organizationId}/wireless/radio/autoRf/channels/recalculate | Recalculates automatically assigned channels for every AP within specified the specified network(s)
[**UpdateDeviceApplianceRadioSettings**](RadioApi.md#UpdateDeviceApplianceRadioSettings) | **Put** /devices/{serial}/appliance/radio/settings | Update the radio settings of an appliance
[**UpdateDeviceWirelessRadioOverrides**](RadioApi.md#UpdateDeviceWirelessRadioOverrides) | **Put** /devices/{serial}/wireless/radio/overrides | Update 2.4 GHz, 5 GHz, and 6 GHz radio settings (channel, channel width, power, and enable/disable) that override RF profiles
[**UpdateDeviceWirelessRadioSettings**](RadioApi.md#UpdateDeviceWirelessRadioSettings) | **Put** /devices/{serial}/wireless/radio/settings | Update 2.4 GHz and 5 GHz radio settings (channel, channel width, power) that override RF profiles
[**UpdateNetworkWirelessRadioRrm**](RadioApi.md#UpdateNetworkWirelessRadioRrm) | **Put** /networks/{networkId}/wireless/radio/rrm | Update the AutoRF settings for a wireless network



## GetDeviceApplianceRadioSettings

> InlineResponse20012 GetDeviceApplianceRadioSettings(ctx, serial).Execute()

Return the radio settings of an appliance



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
    resp, r, err := apiClient.RadioApi.GetDeviceApplianceRadioSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.GetDeviceApplianceRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceApplianceRadioSettings`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.GetDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessRadioOverrides

> InlineResponse20066 GetDeviceWirelessRadioOverrides(ctx, serial).Execute()

Return the radio overrides of a device



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
    resp, r, err := apiClient.RadioApi.GetDeviceWirelessRadioOverrides(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.GetDeviceWirelessRadioOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessRadioOverrides`: InlineResponse20066
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.GetDeviceWirelessRadioOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessRadioOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20066**](InlineResponse20066.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessRadioSettings

> InlineResponse20012 GetDeviceWirelessRadioSettings(ctx, serial).Execute()

Return the manually configured radio settings overrides of a device, which take precedence over RF profiles.



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
    resp, r, err := apiClient.RadioApi.GetDeviceWirelessRadioSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.GetDeviceWirelessRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessRadioSettings`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.GetDeviceWirelessRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessRadioOverridesByDevice

> InlineResponse200523 GetOrganizationWirelessRadioOverridesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()

Return a list of radio overrides



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned radio overrides will be filtered to only include these networks. (optional)
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned radio overrides will be filtered to only include these serials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RadioApi.GetOrganizationWirelessRadioOverridesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.GetOrganizationWirelessRadioOverridesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessRadioOverridesByDevice`: InlineResponse200523
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.GetOrganizationWirelessRadioOverridesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessRadioOverridesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned radio overrides will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned radio overrides will be filtered to only include these serials. | 

### Return type

[**InlineResponse200523**](InlineResponse200523.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessRadioRrmByNetwork

> InlineResponse200524 GetOrganizationWirelessRadioRrmByNetwork(ctx, organizationId).NetworkIds(networkIds).StartingAfter(startingAfter).EndingBefore(endingBefore).PerPage(perPage).SortOrder(sortOrder).Execute()

List the AutoRF settings of an organization by network



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter results by network. (optional)
    startingAfter := "startingAfter_example" // string | Retrieving items after this network ID (optional)
    endingBefore := "endingBefore_example" // string | Retrieving items before this network ID (optional)
    perPage := int32(56) // int32 | Number of items per page (optional)
    sortOrder := "sortOrder_example" // string | The sort order of items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RadioApi.GetOrganizationWirelessRadioRrmByNetwork(context.Background(), organizationId).NetworkIds(networkIds).StartingAfter(startingAfter).EndingBefore(endingBefore).PerPage(perPage).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.GetOrganizationWirelessRadioRrmByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessRadioRrmByNetwork`: InlineResponse200524
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.GetOrganizationWirelessRadioRrmByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessRadioRrmByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter results by network. | 
 **startingAfter** | **string** | Retrieving items after this network ID | 
 **endingBefore** | **string** | Retrieving items before this network ID | 
 **perPage** | **int32** | Number of items per page | 
 **sortOrder** | **string** | The sort order of items | 

### Return type

[**InlineResponse200524**](InlineResponse200524.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecalculateOrganizationWirelessRadioAutoRfChannels

> InlineResponse200522 RecalculateOrganizationWirelessRadioAutoRfChannels(ctx, organizationId).RecalculateOrganizationWirelessRadioAutoRfChannels(recalculateOrganizationWirelessRadioAutoRfChannels).Execute()

Recalculates automatically assigned channels for every AP within specified the specified network(s)



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
    recalculateOrganizationWirelessRadioAutoRfChannels := *openapiclient.NewInlineObject384([]string{"NetworkIds_example"}) // InlineObject384 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RadioApi.RecalculateOrganizationWirelessRadioAutoRfChannels(context.Background(), organizationId).RecalculateOrganizationWirelessRadioAutoRfChannels(recalculateOrganizationWirelessRadioAutoRfChannels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.RecalculateOrganizationWirelessRadioAutoRfChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecalculateOrganizationWirelessRadioAutoRfChannels`: InlineResponse200522
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.RecalculateOrganizationWirelessRadioAutoRfChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recalculateOrganizationWirelessRadioAutoRfChannels** | [**InlineObject384**](InlineObject384.md) |  | 

### Return type

[**InlineResponse200522**](InlineResponse200522.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceApplianceRadioSettings

> InlineResponse20012 UpdateDeviceApplianceRadioSettings(ctx, serial).UpdateDeviceApplianceRadioSettings(updateDeviceApplianceRadioSettings).Execute()

Update the radio settings of an appliance



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
    updateDeviceApplianceRadioSettings := *openapiclient.NewInlineObject6() // InlineObject6 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RadioApi.UpdateDeviceApplianceRadioSettings(context.Background(), serial).UpdateDeviceApplianceRadioSettings(updateDeviceApplianceRadioSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.UpdateDeviceApplianceRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceApplianceRadioSettings`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.UpdateDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceRadioSettings** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessRadioOverrides

> InlineResponse20066 UpdateDeviceWirelessRadioOverrides(ctx, serial).UpdateDeviceWirelessRadioOverrides(updateDeviceWirelessRadioOverrides).Execute()

Update 2.4 GHz, 5 GHz, and 6 GHz radio settings (channel, channel width, power, and enable/disable) that override RF profiles



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
    updateDeviceWirelessRadioOverrides := *openapiclient.NewInlineObject50() // InlineObject50 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RadioApi.UpdateDeviceWirelessRadioOverrides(context.Background(), serial).UpdateDeviceWirelessRadioOverrides(updateDeviceWirelessRadioOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.UpdateDeviceWirelessRadioOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessRadioOverrides`: InlineResponse20066
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.UpdateDeviceWirelessRadioOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessRadioOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessRadioOverrides** | [**InlineObject50**](InlineObject50.md) |  | 

### Return type

[**InlineResponse20066**](InlineResponse20066.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessRadioSettings

> InlineResponse20012 UpdateDeviceWirelessRadioSettings(ctx, serial).UpdateDeviceWirelessRadioSettings(updateDeviceWirelessRadioSettings).Execute()

Update 2.4 GHz and 5 GHz radio settings (channel, channel width, power) that override RF profiles



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
    updateDeviceWirelessRadioSettings := *openapiclient.NewInlineObject51() // InlineObject51 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RadioApi.UpdateDeviceWirelessRadioSettings(context.Background(), serial).UpdateDeviceWirelessRadioSettings(updateDeviceWirelessRadioSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.UpdateDeviceWirelessRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessRadioSettings`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.UpdateDeviceWirelessRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessRadioSettings** | [**InlineObject51**](InlineObject51.md) |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessRadioRrm

> InlineResponse200246 UpdateNetworkWirelessRadioRrm(ctx, networkId).UpdateNetworkWirelessRadioRrm(updateNetworkWirelessRadioRrm).Execute()

Update the AutoRF settings for a wireless network



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
    updateNetworkWirelessRadioRrm := *openapiclient.NewInlineObject217() // InlineObject217 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RadioApi.UpdateNetworkWirelessRadioRrm(context.Background(), networkId).UpdateNetworkWirelessRadioRrm(updateNetworkWirelessRadioRrm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RadioApi.UpdateNetworkWirelessRadioRrm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessRadioRrm`: InlineResponse200246
    fmt.Fprintf(os.Stdout, "Response from `RadioApi.UpdateNetworkWirelessRadioRrm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessRadioRrmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessRadioRrm** | [**InlineObject217**](InlineObject217.md) |  | 

### Return type

[**InlineResponse200246**](InlineResponse200246.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

