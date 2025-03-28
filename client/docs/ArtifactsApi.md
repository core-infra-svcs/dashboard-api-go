# \ArtifactsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCameraCustomAnalyticsArtifact**](ArtifactsApi.md#CreateOrganizationCameraCustomAnalyticsArtifact) | **Post** /organizations/{organizationId}/camera/customAnalytics/artifacts | Create custom analytics artifact
[**DeleteOrganizationCameraCustomAnalyticsArtifact**](ArtifactsApi.md#DeleteOrganizationCameraCustomAnalyticsArtifact) | **Delete** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Delete Custom Analytics Artifact
[**GetOrganizationCameraCustomAnalyticsArtifact**](ArtifactsApi.md#GetOrganizationCameraCustomAnalyticsArtifact) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Get Custom Analytics Artifact
[**GetOrganizationCameraCustomAnalyticsArtifacts**](ArtifactsApi.md#GetOrganizationCameraCustomAnalyticsArtifacts) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts | List Custom Analytics Artifacts



## CreateOrganizationCameraCustomAnalyticsArtifact

> InlineResponse20115 CreateOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId).CreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact).Execute()

Create custom analytics artifact



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
    createOrganizationCameraCustomAnalyticsArtifact := *openapiclient.NewInlineObject236() // InlineObject236 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsApi.CreateOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId).CreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.CreateOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCameraCustomAnalyticsArtifact`: InlineResponse20115
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.CreateOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCameraCustomAnalyticsArtifact** | [**InlineObject236**](InlineObject236.md) |  | 

### Return type

[**InlineResponse20115**](InlineResponse20115.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationCameraCustomAnalyticsArtifact

> DeleteOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Delete Custom Analytics Artifact



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
    artifactId := "artifactId_example" // string | Artifact ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsApi.DeleteOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.DeleteOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**artifactId** | **string** | Artifact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


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


## GetOrganizationCameraCustomAnalyticsArtifact

> InlineResponse200243 GetOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Get Custom Analytics Artifact



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
    artifactId := "artifactId_example" // string | Artifact ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsApi.GetOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.GetOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraCustomAnalyticsArtifact`: InlineResponse200243
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.GetOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**artifactId** | **string** | Artifact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200243**](InlineResponse200243.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraCustomAnalyticsArtifacts

> []InlineResponse200243 GetOrganizationCameraCustomAnalyticsArtifacts(ctx, organizationId).Execute()

List Custom Analytics Artifacts



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
    resp, r, err := apiClient.ArtifactsApi.GetOrganizationCameraCustomAnalyticsArtifacts(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.GetOrganizationCameraCustomAnalyticsArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraCustomAnalyticsArtifacts`: []InlineResponse200243
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.GetOrganizationCameraCustomAnalyticsArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200243**](InlineResponse200243.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

