# \ThemesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSplashTheme**](ThemesApi.md#CreateOrganizationSplashTheme) | **Post** /organizations/{organizationId}/splash/themes | Create a Splash Theme
[**CreateOrganizationSplashThemeAsset**](ThemesApi.md#CreateOrganizationSplashThemeAsset) | **Post** /organizations/{organizationId}/splash/themes/{themeIdentifier}/assets | Create a Splash Theme Asset
[**DeleteOrganizationSplashTheme**](ThemesApi.md#DeleteOrganizationSplashTheme) | **Delete** /organizations/{organizationId}/splash/themes/{id} | Delete a Splash Theme
[**GetOrganizationSplashThemes**](ThemesApi.md#GetOrganizationSplashThemes) | **Get** /organizations/{organizationId}/splash/themes | List Splash Themes



## CreateOrganizationSplashTheme

> InlineResponse200316 CreateOrganizationSplashTheme(ctx, organizationId).CreateOrganizationSplashTheme(createOrganizationSplashTheme).Execute()

Create a Splash Theme



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
    createOrganizationSplashTheme := *openapiclient.NewInlineObject295() // InlineObject295 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemesApi.CreateOrganizationSplashTheme(context.Background(), organizationId).CreateOrganizationSplashTheme(createOrganizationSplashTheme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesApi.CreateOrganizationSplashTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSplashTheme`: InlineResponse200316
    fmt.Fprintf(os.Stdout, "Response from `ThemesApi.CreateOrganizationSplashTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSplashThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSplashTheme** | [**InlineObject295**](InlineObject295.md) |  | 

### Return type

[**InlineResponse200316**](InlineResponse200316.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSplashThemeAsset

> InlineResponse200315 CreateOrganizationSplashThemeAsset(ctx, organizationId, themeIdentifier).CreateOrganizationSplashThemeAsset(createOrganizationSplashThemeAsset).Execute()

Create a Splash Theme Asset



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
    themeIdentifier := "themeIdentifier_example" // string | Theme identifier
    createOrganizationSplashThemeAsset := *openapiclient.NewInlineObject296() // InlineObject296 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemesApi.CreateOrganizationSplashThemeAsset(context.Background(), organizationId, themeIdentifier).CreateOrganizationSplashThemeAsset(createOrganizationSplashThemeAsset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesApi.CreateOrganizationSplashThemeAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSplashThemeAsset`: InlineResponse200315
    fmt.Fprintf(os.Stdout, "Response from `ThemesApi.CreateOrganizationSplashThemeAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**themeIdentifier** | **string** | Theme identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSplashThemeAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOrganizationSplashThemeAsset** | [**InlineObject296**](InlineObject296.md) |  | 

### Return type

[**InlineResponse200315**](InlineResponse200315.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSplashTheme

> DeleteOrganizationSplashTheme(ctx, organizationId, id).Execute()

Delete a Splash Theme



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemesApi.DeleteOrganizationSplashTheme(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesApi.DeleteOrganizationSplashTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSplashThemeRequest struct via the builder pattern


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


## GetOrganizationSplashThemes

> []InlineResponse200316 GetOrganizationSplashThemes(ctx, organizationId).Execute()

List Splash Themes



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
    resp, r, err := apiClient.ThemesApi.GetOrganizationSplashThemes(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemesApi.GetOrganizationSplashThemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSplashThemes`: []InlineResponse200316
    fmt.Fprintf(os.Stdout, "Response from `ThemesApi.GetOrganizationSplashThemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSplashThemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200316**](InlineResponse200316.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

