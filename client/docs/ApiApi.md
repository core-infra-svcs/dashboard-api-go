# \ApiApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAdministeredIdentitiesMeApiKeys**](ApiApi.md#GenerateAdministeredIdentitiesMeApiKeys) | **Post** /administered/identities/me/api/keys/generate | Generates an API key for an identity
[**GetAdministeredIdentitiesMeApiKeys**](ApiApi.md#GetAdministeredIdentitiesMeApiKeys) | **Get** /administered/identities/me/api/keys | List the non-sensitive metadata associated with the API keys that belong to the user
[**RevokeAdministeredIdentitiesMeApiKeys**](ApiApi.md#RevokeAdministeredIdentitiesMeApiKeys) | **Post** /administered/identities/me/api/keys/{suffix}/revoke | Revokes an identity&#39;s API key, using the last four characters of the key



## GenerateAdministeredIdentitiesMeApiKeys

> InlineResponse202 GenerateAdministeredIdentitiesMeApiKeys(ctx).Execute()

Generates an API key for an identity



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiApi.GenerateAdministeredIdentitiesMeApiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.GenerateAdministeredIdentitiesMeApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAdministeredIdentitiesMeApiKeys`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `ApiApi.GenerateAdministeredIdentitiesMeApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAdministeredIdentitiesMeApiKeysRequest struct via the builder pattern


### Return type

[**InlineResponse202**](InlineResponse202.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdministeredIdentitiesMeApiKeys

> []InlineResponse2001 GetAdministeredIdentitiesMeApiKeys(ctx).Execute()

List the non-sensitive metadata associated with the API keys that belong to the user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiApi.GetAdministeredIdentitiesMeApiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.GetAdministeredIdentitiesMeApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredIdentitiesMeApiKeys`: []InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ApiApi.GetAdministeredIdentitiesMeApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredIdentitiesMeApiKeysRequest struct via the builder pattern


### Return type

[**[]InlineResponse2001**](InlineResponse2001.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAdministeredIdentitiesMeApiKeys

> RevokeAdministeredIdentitiesMeApiKeys(ctx, suffix).Execute()

Revokes an identity's API key, using the last four characters of the key



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
    suffix := "suffix_example" // string | Suffix

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiApi.RevokeAdministeredIdentitiesMeApiKeys(context.Background(), suffix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.RevokeAdministeredIdentitiesMeApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**suffix** | **string** | Suffix | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAdministeredIdentitiesMeApiKeysRequest struct via the builder pattern


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

