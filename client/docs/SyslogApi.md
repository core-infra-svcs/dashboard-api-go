# \SyslogApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationDevicesSyslogServersByNetwork**](SyslogApi.md#GetOrganizationDevicesSyslogServersByNetwork) | **Get** /organizations/{organizationId}/devices/syslog/servers/byNetwork | Returns syslog servers configured for the networks within an organization.
[**GetOrganizationDevicesSyslogServersRolesByNetwork**](SyslogApi.md#GetOrganizationDevicesSyslogServersRolesByNetwork) | **Get** /organizations/{organizationId}/devices/syslog/servers/roles/byNetwork | Returns roles that can be assigned to a syslog server for a given network.
[**UpdateNetworkDevicesSyslogServers**](SyslogApi.md#UpdateNetworkDevicesSyslogServers) | **Put** /networks/{networkId}/devices/syslog/servers | Updates the syslog servers configuration for a network.



## GetOrganizationDevicesSyslogServersByNetwork

> InlineResponse200365 GetOrganizationDevicesSyslogServersByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Returns syslog servers configured for the networks within an organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | IDs of the networks for which to fetch syslog servers; suggested maximum array size is 100 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyslogApi.GetOrganizationDevicesSyslogServersByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyslogApi.GetOrganizationDevicesSyslogServersByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesSyslogServersByNetwork`: InlineResponse200365
    fmt.Fprintf(os.Stdout, "Response from `SyslogApi.GetOrganizationDevicesSyslogServersByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesSyslogServersByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | IDs of the networks for which to fetch syslog servers; suggested maximum array size is 100 | 

### Return type

[**InlineResponse200365**](InlineResponse200365.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesSyslogServersRolesByNetwork

> InlineResponse200366 GetOrganizationDevicesSyslogServersRolesByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Returns roles that can be assigned to a syslog server for a given network.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | IDs of the networks for which to fetch valid syslog server roles; suggested maximum array size is 100 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyslogApi.GetOrganizationDevicesSyslogServersRolesByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyslogApi.GetOrganizationDevicesSyslogServersRolesByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesSyslogServersRolesByNetwork`: InlineResponse200366
    fmt.Fprintf(os.Stdout, "Response from `SyslogApi.GetOrganizationDevicesSyslogServersRolesByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesSyslogServersRolesByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | IDs of the networks for which to fetch valid syslog server roles; suggested maximum array size is 100 | 

### Return type

[**InlineResponse200366**](InlineResponse200366.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkDevicesSyslogServers

> InlineResponse200131 UpdateNetworkDevicesSyslogServers(ctx, networkId).UpdateNetworkDevicesSyslogServers(updateNetworkDevicesSyslogServers).Execute()

Updates the syslog servers configuration for a network.



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
    updateNetworkDevicesSyslogServers := *openapiclient.NewInlineObject115([]openapiclient.NetworksNetworkIdDevicesSyslogServersServers{*openapiclient.NewNetworksNetworkIdDevicesSyslogServersServers("Host_example", int32(123), []string{"Roles_example"})}) // InlineObject115 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyslogApi.UpdateNetworkDevicesSyslogServers(context.Background(), networkId).UpdateNetworkDevicesSyslogServers(updateNetworkDevicesSyslogServers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyslogApi.UpdateNetworkDevicesSyslogServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkDevicesSyslogServers`: InlineResponse200131
    fmt.Fprintf(os.Stdout, "Response from `SyslogApi.UpdateNetworkDevicesSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkDevicesSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDevicesSyslogServers** | [**InlineObject115**](InlineObject115.md) |  | 

### Return type

[**InlineResponse200131**](InlineResponse200131.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

