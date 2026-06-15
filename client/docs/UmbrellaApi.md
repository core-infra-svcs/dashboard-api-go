# \UmbrellaApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworkApplianceUmbrellaPolicies**](UmbrellaApi.md#AddNetworkApplianceUmbrellaPolicies) | **Post** /networks/{networkId}/appliance/umbrella/policies/add | Add one Cisco Umbrella DNS security policy to an MX network by policy ID
[**ConnectNetworkApplianceUmbrellaAccount**](UmbrellaApi.md#ConnectNetworkApplianceUmbrellaAccount) | **Post** /networks/{networkId}/appliance/umbrella/account/connect | Connect a Cisco Umbrella account to this network
[**DisconnectNetworkApplianceUmbrellaAccount**](UmbrellaApi.md#DisconnectNetworkApplianceUmbrellaAccount) | **Post** /networks/{networkId}/appliance/umbrella/account/disconnect | Disconnect Umbrella account from this network
[**ExclusionsNetworkApplianceUmbrellaDomains**](UmbrellaApi.md#ExclusionsNetworkApplianceUmbrellaDomains) | **Put** /networks/{networkId}/appliance/umbrella/domains/exclusions | Specify one or more domain names to be excluded from being routed to Cisco Umbrella.
[**ProtectionNetworkApplianceUmbrella**](UmbrellaApi.md#ProtectionNetworkApplianceUmbrella) | **Put** /networks/{networkId}/appliance/umbrella/protection | Enable or disable umbrella protection for an appliance network
[**RemoveNetworkApplianceUmbrellaPolicies**](UmbrellaApi.md#RemoveNetworkApplianceUmbrellaPolicies) | **Post** /networks/{networkId}/appliance/umbrella/policies/remove | Remove one Cisco Umbrella DNS security policy from an MX network by policy ID



## AddNetworkApplianceUmbrellaPolicies

> InlineResponse200103 AddNetworkApplianceUmbrellaPolicies(ctx, networkId).AddNetworkApplianceUmbrellaPolicies(addNetworkApplianceUmbrellaPolicies).Execute()

Add one Cisco Umbrella DNS security policy to an MX network by policy ID



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
    addNetworkApplianceUmbrellaPolicies := *openapiclient.NewInlineObject88(*openapiclient.NewNetworksNetworkIdApplianceUmbrellaPoliciesAddPolicy("Id_example")) // InlineObject88 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UmbrellaApi.AddNetworkApplianceUmbrellaPolicies(context.Background(), networkId).AddNetworkApplianceUmbrellaPolicies(addNetworkApplianceUmbrellaPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UmbrellaApi.AddNetworkApplianceUmbrellaPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNetworkApplianceUmbrellaPolicies`: InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `UmbrellaApi.AddNetworkApplianceUmbrellaPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkApplianceUmbrellaPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addNetworkApplianceUmbrellaPolicies** | [**InlineObject88**](InlineObject88.md) |  | 

### Return type

[**InlineResponse200103**](InlineResponse200103.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectNetworkApplianceUmbrellaAccount

> InlineResponse200101 ConnectNetworkApplianceUmbrellaAccount(ctx, networkId).ConnectNetworkApplianceUmbrellaAccount(connectNetworkApplianceUmbrellaAccount).Execute()

Connect a Cisco Umbrella account to this network



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
    connectNetworkApplianceUmbrellaAccount := *openapiclient.NewInlineObject86(*openapiclient.NewNetworksNetworkIdApplianceUmbrellaAccountConnectApi("Key_example", "Secret_example")) // InlineObject86 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UmbrellaApi.ConnectNetworkApplianceUmbrellaAccount(context.Background(), networkId).ConnectNetworkApplianceUmbrellaAccount(connectNetworkApplianceUmbrellaAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UmbrellaApi.ConnectNetworkApplianceUmbrellaAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectNetworkApplianceUmbrellaAccount`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `UmbrellaApi.ConnectNetworkApplianceUmbrellaAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectNetworkApplianceUmbrellaAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectNetworkApplianceUmbrellaAccount** | [**InlineObject86**](InlineObject86.md) |  | 

### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectNetworkApplianceUmbrellaAccount

> DisconnectNetworkApplianceUmbrellaAccount(ctx, networkId).Execute()

Disconnect Umbrella account from this network



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
    resp, r, err := apiClient.UmbrellaApi.DisconnectNetworkApplianceUmbrellaAccount(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UmbrellaApi.DisconnectNetworkApplianceUmbrellaAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectNetworkApplianceUmbrellaAccountRequest struct via the builder pattern


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


## ExclusionsNetworkApplianceUmbrellaDomains

> InlineResponse200102 ExclusionsNetworkApplianceUmbrellaDomains(ctx, networkId).ExclusionsNetworkApplianceUmbrellaDomains(exclusionsNetworkApplianceUmbrellaDomains).Execute()

Specify one or more domain names to be excluded from being routed to Cisco Umbrella.



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
    exclusionsNetworkApplianceUmbrellaDomains := *openapiclient.NewInlineObject87([]string{"Domains_example"}) // InlineObject87 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UmbrellaApi.ExclusionsNetworkApplianceUmbrellaDomains(context.Background(), networkId).ExclusionsNetworkApplianceUmbrellaDomains(exclusionsNetworkApplianceUmbrellaDomains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UmbrellaApi.ExclusionsNetworkApplianceUmbrellaDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExclusionsNetworkApplianceUmbrellaDomains`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `UmbrellaApi.ExclusionsNetworkApplianceUmbrellaDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExclusionsNetworkApplianceUmbrellaDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclusionsNetworkApplianceUmbrellaDomains** | [**InlineObject87**](InlineObject87.md) |  | 

### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtectionNetworkApplianceUmbrella

> InlineResponse200104 ProtectionNetworkApplianceUmbrella(ctx, networkId).ProtectionNetworkApplianceUmbrella(protectionNetworkApplianceUmbrella).Execute()

Enable or disable umbrella protection for an appliance network



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
    protectionNetworkApplianceUmbrella := *openapiclient.NewInlineObject90(false) // InlineObject90 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UmbrellaApi.ProtectionNetworkApplianceUmbrella(context.Background(), networkId).ProtectionNetworkApplianceUmbrella(protectionNetworkApplianceUmbrella).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UmbrellaApi.ProtectionNetworkApplianceUmbrella``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtectionNetworkApplianceUmbrella`: InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `UmbrellaApi.ProtectionNetworkApplianceUmbrella`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProtectionNetworkApplianceUmbrellaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **protectionNetworkApplianceUmbrella** | [**InlineObject90**](InlineObject90.md) |  | 

### Return type

[**InlineResponse200104**](InlineResponse200104.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNetworkApplianceUmbrellaPolicies

> RemoveNetworkApplianceUmbrellaPolicies(ctx, networkId).RemoveNetworkApplianceUmbrellaPolicies(removeNetworkApplianceUmbrellaPolicies).Execute()

Remove one Cisco Umbrella DNS security policy from an MX network by policy ID



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
    removeNetworkApplianceUmbrellaPolicies := *openapiclient.NewInlineObject89(*openapiclient.NewNetworksNetworkIdApplianceUmbrellaPoliciesRemovePolicy("Id_example")) // InlineObject89 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UmbrellaApi.RemoveNetworkApplianceUmbrellaPolicies(context.Background(), networkId).RemoveNetworkApplianceUmbrellaPolicies(removeNetworkApplianceUmbrellaPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UmbrellaApi.RemoveNetworkApplianceUmbrellaPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkApplianceUmbrellaPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeNetworkApplianceUmbrellaPolicies** | [**InlineObject89**](InlineObject89.md) |  | 

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

