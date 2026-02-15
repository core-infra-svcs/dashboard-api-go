# \MqttApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWirelessMqttSettings**](MqttApi.md#GetOrganizationWirelessMqttSettings) | **Get** /organizations/{organizationId}/wireless/mqtt/settings | Return MQTT Settings for networks
[**UpdateOrganizationWirelessMqttSettings**](MqttApi.md#UpdateOrganizationWirelessMqttSettings) | **Put** /organizations/{organizationId}/wireless/mqtt/settings | Add new broker config for wireless MQTT



## GetOrganizationWirelessMqttSettings

> InlineResponse200385 GetOrganizationWirelessMqttSettings(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Return MQTT Settings for networks



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 250. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter mqtt settings by network ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttApi.GetOrganizationWirelessMqttSettings(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttApi.GetOrganizationWirelessMqttSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessMqttSettings`: InlineResponse200385
    fmt.Fprintf(os.Stdout, "Response from `MqttApi.GetOrganizationWirelessMqttSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessMqttSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 250. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter mqtt settings by network ID. | 

### Return type

[**InlineResponse200385**](InlineResponse200385.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessMqttSettings

> InlineResponse200385Items UpdateOrganizationWirelessMqttSettings(ctx, organizationId).UpdateOrganizationWirelessMqttSettings(updateOrganizationWirelessMqttSettings).Execute()

Add new broker config for wireless MQTT



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
    updateOrganizationWirelessMqttSettings := *openapiclient.NewInlineObject311(*openapiclient.NewOrganizationsOrganizationIdWirelessMqttSettingsNetwork("Id_example"), *openapiclient.NewOrganizationsOrganizationIdWirelessMqttSettingsMqtt(false)) // InlineObject311 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttApi.UpdateOrganizationWirelessMqttSettings(context.Background(), organizationId).UpdateOrganizationWirelessMqttSettings(updateOrganizationWirelessMqttSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttApi.UpdateOrganizationWirelessMqttSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessMqttSettings`: InlineResponse200385Items
    fmt.Fprintf(os.Stdout, "Response from `MqttApi.UpdateOrganizationWirelessMqttSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessMqttSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationWirelessMqttSettings** | [**InlineObject311**](InlineObject311.md) |  | 

### Return type

[**InlineResponse200385Items**](InlineResponse200385Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

