# \TargetGroupsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSmTargetGroup**](TargetGroupsApi.md#CreateNetworkSmTargetGroup) | **Post** /networks/{networkId}/sm/targetGroups | Add a target group
[**DeleteNetworkSmTargetGroup**](TargetGroupsApi.md#DeleteNetworkSmTargetGroup) | **Delete** /networks/{networkId}/sm/targetGroups/{targetGroupId} | Delete a target group from a network
[**GetNetworkSmTargetGroup**](TargetGroupsApi.md#GetNetworkSmTargetGroup) | **Get** /networks/{networkId}/sm/targetGroups/{targetGroupId} | Return a target group
[**GetNetworkSmTargetGroups**](TargetGroupsApi.md#GetNetworkSmTargetGroups) | **Get** /networks/{networkId}/sm/targetGroups | List the target groups in this network
[**UpdateNetworkSmTargetGroup**](TargetGroupsApi.md#UpdateNetworkSmTargetGroup) | **Put** /networks/{networkId}/sm/targetGroups/{targetGroupId} | Update a target group



## CreateNetworkSmTargetGroup

> InlineResponse200143 CreateNetworkSmTargetGroup(ctx, networkId).CreateNetworkSmTargetGroup(createNetworkSmTargetGroup).Execute()

Add a target group



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
    createNetworkSmTargetGroup := *openapiclient.NewInlineObject134() // InlineObject134 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetGroupsApi.CreateNetworkSmTargetGroup(context.Background(), networkId).CreateNetworkSmTargetGroup(createNetworkSmTargetGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.CreateNetworkSmTargetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSmTargetGroup`: InlineResponse200143
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.CreateNetworkSmTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSmTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSmTargetGroup** | [**InlineObject134**](InlineObject134.md) |  | 

### Return type

[**InlineResponse200143**](InlineResponse200143.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSmTargetGroup

> DeleteNetworkSmTargetGroup(ctx, networkId, targetGroupId).Execute()

Delete a target group from a network



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
    targetGroupId := "targetGroupId_example" // string | Target group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetGroupsApi.DeleteNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.DeleteNetworkSmTargetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**targetGroupId** | **string** | Target group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSmTargetGroupRequest struct via the builder pattern


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


## GetNetworkSmTargetGroup

> InlineResponse200143 GetNetworkSmTargetGroup(ctx, networkId, targetGroupId).WithDetails(withDetails).Execute()

Return a target group



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
    targetGroupId := "targetGroupId_example" // string | Target group ID
    withDetails := true // bool | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetGroupsApi.GetNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).WithDetails(withDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.GetNetworkSmTargetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmTargetGroup`: InlineResponse200143
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.GetNetworkSmTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**targetGroupId** | **string** | Target group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDetails** | **bool** | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response | 

### Return type

[**InlineResponse200143**](InlineResponse200143.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmTargetGroups

> []InlineResponse200143 GetNetworkSmTargetGroups(ctx, networkId).WithDetails(withDetails).Execute()

List the target groups in this network



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
    withDetails := true // bool | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetGroupsApi.GetNetworkSmTargetGroups(context.Background(), networkId).WithDetails(withDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.GetNetworkSmTargetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmTargetGroups`: []InlineResponse200143
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.GetNetworkSmTargetGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmTargetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDetails** | **bool** | Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response | 

### Return type

[**[]InlineResponse200143**](InlineResponse200143.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSmTargetGroup

> InlineResponse200143 UpdateNetworkSmTargetGroup(ctx, networkId, targetGroupId).UpdateNetworkSmTargetGroup(updateNetworkSmTargetGroup).Execute()

Update a target group



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
    targetGroupId := "targetGroupId_example" // string | Target group ID
    updateNetworkSmTargetGroup := *openapiclient.NewInlineObject135() // InlineObject135 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TargetGroupsApi.UpdateNetworkSmTargetGroup(context.Background(), networkId, targetGroupId).UpdateNetworkSmTargetGroup(updateNetworkSmTargetGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.UpdateNetworkSmTargetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSmTargetGroup`: InlineResponse200143
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.UpdateNetworkSmTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**targetGroupId** | **string** | Target group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSmTargetGroup** | [**InlineObject135**](InlineObject135.md) |  | 

### Return type

[**InlineResponse200143**](InlineResponse200143.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

