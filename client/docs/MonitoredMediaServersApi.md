# \MonitoredMediaServersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersApi.md#CreateOrganizationInsightMonitoredMediaServer) | **Post** /organizations/{organizationId}/insight/monitoredMediaServers | Add a media server to be monitored for this organization
[**DeleteOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersApi.md#DeleteOrganizationInsightMonitoredMediaServer) | **Delete** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Delete a monitored media server from this organization
[**GetOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersApi.md#GetOrganizationInsightMonitoredMediaServer) | **Get** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Return a monitored media server for this organization
[**GetOrganizationInsightMonitoredMediaServers**](MonitoredMediaServersApi.md#GetOrganizationInsightMonitoredMediaServers) | **Get** /organizations/{organizationId}/insight/monitoredMediaServers | List the monitored media servers for this organization
[**UpdateOrganizationInsightMonitoredMediaServer**](MonitoredMediaServersApi.md#UpdateOrganizationInsightMonitoredMediaServer) | **Put** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Update a monitored media server for this organization



## CreateOrganizationInsightMonitoredMediaServer

> InlineResponse200291 CreateOrganizationInsightMonitoredMediaServer(ctx, organizationId).CreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer).Execute()

Add a media server to be monitored for this organization



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
    createOrganizationInsightMonitoredMediaServer := *openapiclient.NewInlineObject266("Name_example", "Address_example") // InlineObject266 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoredMediaServersApi.CreateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId).CreateOrganizationInsightMonitoredMediaServer(createOrganizationInsightMonitoredMediaServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersApi.CreateOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInsightMonitoredMediaServer`: InlineResponse200291
    fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersApi.CreateOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInsightMonitoredMediaServer** | [**InlineObject266**](InlineObject266.md) |  | 

### Return type

[**InlineResponse200291**](InlineResponse200291.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationInsightMonitoredMediaServer

> DeleteOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).Execute()

Delete a monitored media server from this organization



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
    monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoredMediaServersApi.DeleteOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersApi.DeleteOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


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


## GetOrganizationInsightMonitoredMediaServer

> InlineResponse200291 GetOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).Execute()

Return a monitored media server for this organization



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
    monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoredMediaServersApi.GetOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersApi.GetOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInsightMonitoredMediaServer`: InlineResponse200291
    fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersApi.GetOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200291**](InlineResponse200291.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInsightMonitoredMediaServers

> []InlineResponse200291 GetOrganizationInsightMonitoredMediaServers(ctx, organizationId).Execute()

List the monitored media servers for this organization



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
    resp, r, err := apiClient.MonitoredMediaServersApi.GetOrganizationInsightMonitoredMediaServers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersApi.GetOrganizationInsightMonitoredMediaServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInsightMonitoredMediaServers`: []InlineResponse200291
    fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersApi.GetOrganizationInsightMonitoredMediaServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInsightMonitoredMediaServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200291**](InlineResponse200291.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationInsightMonitoredMediaServer

> InlineResponse200291 UpdateOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).UpdateOrganizationInsightMonitoredMediaServer(updateOrganizationInsightMonitoredMediaServer).Execute()

Update a monitored media server for this organization



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
    monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID
    updateOrganizationInsightMonitoredMediaServer := *openapiclient.NewInlineObject267() // InlineObject267 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoredMediaServersApi.UpdateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).UpdateOrganizationInsightMonitoredMediaServer(updateOrganizationInsightMonitoredMediaServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredMediaServersApi.UpdateOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationInsightMonitoredMediaServer`: InlineResponse200291
    fmt.Fprintf(os.Stdout, "Response from `MonitoredMediaServersApi.UpdateOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationInsightMonitoredMediaServer** | [**InlineObject267**](InlineObject267.md) |  | 

### Return type

[**InlineResponse200291**](InlineResponse200291.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

