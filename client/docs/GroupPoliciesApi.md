# \GroupPoliciesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkGroupPolicy**](GroupPoliciesApi.md#CreateNetworkGroupPolicy) | **Post** /networks/{networkId}/groupPolicies | Create a group policy
[**DeleteNetworkGroupPolicy**](GroupPoliciesApi.md#DeleteNetworkGroupPolicy) | **Delete** /networks/{networkId}/groupPolicies/{groupPolicyId} | Delete a group policy
[**GetNetworkGroupPolicies**](GroupPoliciesApi.md#GetNetworkGroupPolicies) | **Get** /networks/{networkId}/groupPolicies | List the group policies in a network
[**GetNetworkGroupPolicy**](GroupPoliciesApi.md#GetNetworkGroupPolicy) | **Get** /networks/{networkId}/groupPolicies/{groupPolicyId} | Display a group policy
[**UpdateNetworkGroupPolicy**](GroupPoliciesApi.md#UpdateNetworkGroupPolicy) | **Put** /networks/{networkId}/groupPolicies/{groupPolicyId} | Update a group policy



## CreateNetworkGroupPolicy

> InlineResponse200107 CreateNetworkGroupPolicy(ctx, networkId).CreateNetworkGroupPolicy(createNetworkGroupPolicy).Execute()

Create a group policy



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
    createNetworkGroupPolicy := *openapiclient.NewInlineObject111("Name_example") // InlineObject111 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.CreateNetworkGroupPolicy(context.Background(), networkId).CreateNetworkGroupPolicy(createNetworkGroupPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.CreateNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkGroupPolicy`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.CreateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkGroupPolicy** | [**InlineObject111**](InlineObject111.md) |  | 

### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkGroupPolicy

> DeleteNetworkGroupPolicy(ctx, networkId, groupPolicyId).Force(force).Execute()

Delete a group policy



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
    groupPolicyId := "groupPolicyId_example" // string | Group policy ID
    force := true // bool | If true, the system deletes the GP even if there are active clients using the GP. After deletion, active clients that were assigned to that Group Policy will be left without any policy applied. Default is false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.DeleteNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.DeleteNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | If true, the system deletes the GP even if there are active clients using the GP. After deletion, active clients that were assigned to that Group Policy will be left without any policy applied. Default is false. | 

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


## GetNetworkGroupPolicies

> []InlineResponse200107 GetNetworkGroupPolicies(ctx, networkId).Execute()

List the group policies in a network



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
    resp, r, err := apiClient.GroupPoliciesApi.GetNetworkGroupPolicies(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.GetNetworkGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroupPolicies`: []InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.GetNetworkGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroupPolicy

> InlineResponse200107 GetNetworkGroupPolicy(ctx, networkId, groupPolicyId).Execute()

Display a group policy



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
    groupPolicyId := "groupPolicyId_example" // string | Group policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.GetNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.GetNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroupPolicy`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.GetNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkGroupPolicy

> InlineResponse200107 UpdateNetworkGroupPolicy(ctx, networkId, groupPolicyId).UpdateNetworkGroupPolicy(updateNetworkGroupPolicy).Execute()

Update a group policy



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
    groupPolicyId := "groupPolicyId_example" // string | Group policy ID
    updateNetworkGroupPolicy := *openapiclient.NewInlineObject112() // InlineObject112 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.UpdateNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).UpdateNetworkGroupPolicy(updateNetworkGroupPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.UpdateNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkGroupPolicy`: InlineResponse200107
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.UpdateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkGroupPolicy** | [**InlineObject112**](InlineObject112.md) |  | 

### Return type

[**InlineResponse200107**](InlineResponse200107.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

