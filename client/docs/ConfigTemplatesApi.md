# \ConfigTemplatesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationConfigTemplate**](ConfigTemplatesApi.md#CreateOrganizationConfigTemplate) | **Post** /organizations/{organizationId}/configTemplates | Create a new configuration template
[**DeleteOrganizationConfigTemplate**](ConfigTemplatesApi.md#DeleteOrganizationConfigTemplate) | **Delete** /organizations/{organizationId}/configTemplates/{configTemplateId} | Remove a configuration template
[**GetOrganizationConfigTemplate**](ConfigTemplatesApi.md#GetOrganizationConfigTemplate) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId} | Return a single configuration template
[**GetOrganizationConfigTemplateSwitchProfilePort**](ConfigTemplatesApi.md#GetOrganizationConfigTemplateSwitchProfilePort) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Return a switch template port
[**GetOrganizationConfigTemplateSwitchProfilePorts**](ConfigTemplatesApi.md#GetOrganizationConfigTemplateSwitchProfilePorts) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports | Return all the ports of a switch template
[**GetOrganizationConfigTemplateSwitchProfiles**](ConfigTemplatesApi.md#GetOrganizationConfigTemplateSwitchProfiles) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles | List the switch templates for your switch template configuration
[**GetOrganizationConfigTemplates**](ConfigTemplatesApi.md#GetOrganizationConfigTemplates) | **Get** /organizations/{organizationId}/configTemplates | List the configuration templates for this organization
[**UpdateOrganizationConfigTemplate**](ConfigTemplatesApi.md#UpdateOrganizationConfigTemplate) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId} | Update a configuration template
[**UpdateOrganizationConfigTemplateSwitchProfilePort**](ConfigTemplatesApi.md#UpdateOrganizationConfigTemplateSwitchProfilePort) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Update a switch template port



## CreateOrganizationConfigTemplate

> InlineResponse200264 CreateOrganizationConfigTemplate(ctx, organizationId).CreateOrganizationConfigTemplate(createOrganizationConfigTemplate).Execute()

Create a new configuration template



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
    createOrganizationConfigTemplate := *openapiclient.NewInlineObject251("Name_example") // InlineObject251 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.CreateOrganizationConfigTemplate(context.Background(), organizationId).CreateOrganizationConfigTemplate(createOrganizationConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.CreateOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationConfigTemplate`: InlineResponse200264
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.CreateOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationConfigTemplate** | [**InlineObject251**](InlineObject251.md) |  | 

### Return type

[**InlineResponse200264**](InlineResponse200264.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationConfigTemplate

> DeleteOrganizationConfigTemplate(ctx, organizationId, configTemplateId).Execute()

Remove a configuration template



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
    configTemplateId := "configTemplateId_example" // string | Config template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.DeleteOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.DeleteOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationConfigTemplateRequest struct via the builder pattern


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


## GetOrganizationConfigTemplate

> InlineResponse200264 GetOrganizationConfigTemplate(ctx, organizationId, configTemplateId).Execute()

Return a single configuration template



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
    configTemplateId := "configTemplateId_example" // string | Config template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.GetOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.GetOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplate`: InlineResponse200264
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.GetOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200264**](InlineResponse200264.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePort

> InlineResponse200266 GetOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).Execute()

Return a switch template port



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
    configTemplateId := "configTemplateId_example" // string | Config template ID
    profileId := "profileId_example" // string | Profile ID
    portId := "portId_example" // string | Port ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplateSwitchProfilePort`: InlineResponse200266
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**InlineResponse200266**](InlineResponse200266.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePorts

> []InlineResponse200266 GetOrganizationConfigTemplateSwitchProfilePorts(ctx, organizationId, configTemplateId, profileId).Execute()

Return all the ports of a switch template



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
    configTemplateId := "configTemplateId_example" // string | Config template ID
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfilePorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplateSwitchProfilePorts`: []InlineResponse200266
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfilePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]InlineResponse200266**](InlineResponse200266.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfiles

> []InlineResponse200265 GetOrganizationConfigTemplateSwitchProfiles(ctx, organizationId, configTemplateId).Execute()

List the switch templates for your switch template configuration



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
    configTemplateId := "configTemplateId_example" // string | Config template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfiles(context.Background(), organizationId, configTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplateSwitchProfiles`: []InlineResponse200265
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.GetOrganizationConfigTemplateSwitchProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse200265**](InlineResponse200265.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplates

> []InlineResponse200264 GetOrganizationConfigTemplates(ctx, organizationId).Execute()

List the configuration templates for this organization



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
    resp, r, err := apiClient.ConfigTemplatesApi.GetOrganizationConfigTemplates(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.GetOrganizationConfigTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplates`: []InlineResponse200264
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.GetOrganizationConfigTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200264**](InlineResponse200264.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationConfigTemplate

> InlineResponse200264 UpdateOrganizationConfigTemplate(ctx, organizationId, configTemplateId).UpdateOrganizationConfigTemplate(updateOrganizationConfigTemplate).Execute()

Update a configuration template



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
    configTemplateId := "configTemplateId_example" // string | Config template ID
    updateOrganizationConfigTemplate := *openapiclient.NewInlineObject252() // InlineObject252 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.UpdateOrganizationConfigTemplate(context.Background(), organizationId, configTemplateId).UpdateOrganizationConfigTemplate(updateOrganizationConfigTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.UpdateOrganizationConfigTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationConfigTemplate`: InlineResponse200264
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.UpdateOrganizationConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationConfigTemplate** | [**InlineObject252**](InlineObject252.md) |  | 

### Return type

[**InlineResponse200264**](InlineResponse200264.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationConfigTemplateSwitchProfilePort

> InlineResponse200266 UpdateOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePort(updateOrganizationConfigTemplateSwitchProfilePort).Execute()

Update a switch template port



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
    configTemplateId := "configTemplateId_example" // string | Config template ID
    profileId := "profileId_example" // string | Profile ID
    portId := "portId_example" // string | Port ID
    updateOrganizationConfigTemplateSwitchProfilePort := *openapiclient.NewInlineObject253() // InlineObject253 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigTemplatesApi.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePort(updateOrganizationConfigTemplateSwitchProfilePort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigTemplatesApi.UpdateOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationConfigTemplateSwitchProfilePort`: InlineResponse200266
    fmt.Fprintf(os.Stdout, "Response from `ConfigTemplatesApi.UpdateOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateOrganizationConfigTemplateSwitchProfilePort** | [**InlineObject253**](InlineObject253.md) |  | 

### Return type

[**InlineResponse200266**](InlineResponse200266.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

