# \MerakiAuthUsersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#CreateNetworkMerakiAuthUser) | **Post** /networks/{networkId}/merakiAuthUsers | Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)
[**DeleteNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#DeleteNetworkMerakiAuthUser) | **Delete** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Deauthorize a user
[**GetNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#GetNetworkMerakiAuthUser) | **Get** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Return the Meraki Auth splash guest, RADIUS, or client VPN user
[**GetNetworkMerakiAuthUsers**](MerakiAuthUsersApi.md#GetNetworkMerakiAuthUsers) | **Get** /networks/{networkId}/merakiAuthUsers | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)
[**UpdateNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#UpdateNetworkMerakiAuthUser) | **Put** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)



## CreateNetworkMerakiAuthUser

> InlineResponse20034 CreateNetworkMerakiAuthUser(ctx, networkId).CreateNetworkMerakiAuthUser(createNetworkMerakiAuthUser).Execute()

Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)



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
    createNetworkMerakiAuthUser := *openapiclient.NewInlineObject89("Email_example", []openapiclient.NetworksNetworkIdMerakiAuthUsersAuthorizations1{*openapiclient.NewNetworksNetworkIdMerakiAuthUsersAuthorizations1()}) // InlineObject89 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.CreateNetworkMerakiAuthUser(context.Background(), networkId).CreateNetworkMerakiAuthUser(createNetworkMerakiAuthUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.CreateNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkMerakiAuthUser`: InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.CreateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMerakiAuthUser** | [**InlineObject89**](InlineObject89.md) |  | 

### Return type

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkMerakiAuthUser

> DeleteNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Execute()

Deauthorize a user



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
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.DeleteNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.DeleteNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMerakiAuthUser

> InlineResponse20034 GetNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Execute()

Return the Meraki Auth splash guest, RADIUS, or client VPN user



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
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.GetNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.GetNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMerakiAuthUser`: InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.GetNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMerakiAuthUsers

> []InlineResponse20034 GetNetworkMerakiAuthUsers(ctx, networkId).Execute()

List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)



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
    resp, r, err := apiClient.MerakiAuthUsersApi.GetNetworkMerakiAuthUsers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.GetNetworkMerakiAuthUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMerakiAuthUsers`: []InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.GetNetworkMerakiAuthUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20034**](InlineResponse20034.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkMerakiAuthUser

> InlineResponse20034 UpdateNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUser(updateNetworkMerakiAuthUser).Execute()

Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)



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
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID
    updateNetworkMerakiAuthUser := *openapiclient.NewInlineObject90() // InlineObject90 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.UpdateNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUser(updateNetworkMerakiAuthUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.UpdateNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkMerakiAuthUser`: InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.UpdateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMerakiAuthUser** | [**InlineObject90**](InlineObject90.md) |  | 

### Return type

[**InlineResponse20034**](InlineResponse20034.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

