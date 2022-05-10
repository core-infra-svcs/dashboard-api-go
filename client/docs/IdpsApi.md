# \IdpsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSamlIdp**](IdpsApi.md#CreateOrganizationSamlIdp) | **Post** /organizations/{organizationId}/saml/idps | Create a SAML IdP for your organization.
[**DeleteOrganizationSamlIdp**](IdpsApi.md#DeleteOrganizationSamlIdp) | **Delete** /organizations/{organizationId}/saml/idps/{idpId} | Remove a SAML IdP in your organization.
[**GetOrganizationSamlIdp**](IdpsApi.md#GetOrganizationSamlIdp) | **Get** /organizations/{organizationId}/saml/idps/{idpId} | Get a SAML IdP from your organization.
[**GetOrganizationSamlIdps**](IdpsApi.md#GetOrganizationSamlIdps) | **Get** /organizations/{organizationId}/saml/idps | List the SAML IdPs in your organization.
[**UpdateOrganizationSamlIdp**](IdpsApi.md#UpdateOrganizationSamlIdp) | **Put** /organizations/{organizationId}/saml/idps/{idpId} | Update a SAML IdP in your organization



## CreateOrganizationSamlIdp

> map[string]interface{} CreateOrganizationSamlIdp(ctx, organizationId).CreateOrganizationSamlIdp(createOrganizationSamlIdp).Execute()

Create a SAML IdP for your organization.



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
    organizationId := "organizationId_example" // string | 
    createOrganizationSamlIdp := *openapiclient.NewInlineObject187("X509certSha1Fingerprint_example") // InlineObject187 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpsApi.CreateOrganizationSamlIdp(context.Background(), organizationId).CreateOrganizationSamlIdp(createOrganizationSamlIdp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.CreateOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSamlIdp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.CreateOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSamlIdp** | [**InlineObject187**](InlineObject187.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSamlIdp

> DeleteOrganizationSamlIdp(ctx, organizationId, idpId).Execute()

Remove a SAML IdP in your organization.



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
    organizationId := "organizationId_example" // string | 
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpsApi.DeleteOrganizationSamlIdp(context.Background(), organizationId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.DeleteOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSamlIdpRequest struct via the builder pattern


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


## GetOrganizationSamlIdp

> map[string]interface{} GetOrganizationSamlIdp(ctx, organizationId, idpId).Execute()

Get a SAML IdP from your organization.



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
    organizationId := "organizationId_example" // string | 
    idpId := "idpId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpsApi.GetOrganizationSamlIdp(context.Background(), organizationId, idpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.GetOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlIdp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.GetOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlIdps

> []map[string]interface{} GetOrganizationSamlIdps(ctx, organizationId).Execute()

List the SAML IdPs in your organization.



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
    organizationId := "organizationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpsApi.GetOrganizationSamlIdps(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.GetOrganizationSamlIdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlIdps`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.GetOrganizationSamlIdps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlIdpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSamlIdp

> map[string]interface{} UpdateOrganizationSamlIdp(ctx, organizationId, idpId).UpdateOrganizationSamlIdp(updateOrganizationSamlIdp).Execute()

Update a SAML IdP in your organization



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
    organizationId := "organizationId_example" // string | 
    idpId := "idpId_example" // string | 
    updateOrganizationSamlIdp := *openapiclient.NewInlineObject188() // InlineObject188 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdpsApi.UpdateOrganizationSamlIdp(context.Background(), organizationId, idpId).UpdateOrganizationSamlIdp(updateOrganizationSamlIdp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.UpdateOrganizationSamlIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSamlIdp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.UpdateOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**idpId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSamlIdp** | [**InlineObject188**](InlineObject188.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

