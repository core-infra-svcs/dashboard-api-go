# \AclsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAdaptivePolicyAcl**](AclsApi.md#CreateOrganizationAdaptivePolicyAcl) | **Post** /organizations/{organizationId}/adaptivePolicy/acls | Creates new adaptive policy ACL
[**DeleteOrganizationAdaptivePolicyAcl**](AclsApi.md#DeleteOrganizationAdaptivePolicyAcl) | **Delete** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Deletes the specified adaptive policy ACL
[**GetOrganizationAdaptivePolicyAcl**](AclsApi.md#GetOrganizationAdaptivePolicyAcl) | **Get** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Returns the adaptive policy ACL information
[**GetOrganizationAdaptivePolicyAcls**](AclsApi.md#GetOrganizationAdaptivePolicyAcls) | **Get** /organizations/{organizationId}/adaptivePolicy/acls | List adaptive policy ACLs in a organization
[**UpdateOrganizationAdaptivePolicyAcl**](AclsApi.md#UpdateOrganizationAdaptivePolicyAcl) | **Put** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Updates an adaptive policy ACL



## CreateOrganizationAdaptivePolicyAcl

> InlineResponse200211 CreateOrganizationAdaptivePolicyAcl(ctx, organizationId).CreateOrganizationAdaptivePolicyAcl(createOrganizationAdaptivePolicyAcl).Execute()

Creates new adaptive policy ACL



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
    createOrganizationAdaptivePolicyAcl := *openapiclient.NewInlineObject211("Name_example", []openapiclient.OrganizationsOrganizationIdAdaptivePolicyAclsRules1{*openapiclient.NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1("Policy_example", "Protocol_example")}, "IpVersion_example") // InlineObject211 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AclsApi.CreateOrganizationAdaptivePolicyAcl(context.Background(), organizationId).CreateOrganizationAdaptivePolicyAcl(createOrganizationAdaptivePolicyAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.CreateOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyAcl`: InlineResponse200211
    fmt.Fprintf(os.Stdout, "Response from `AclsApi.CreateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyAcl** | [**InlineObject211**](InlineObject211.md) |  | 

### Return type

[**InlineResponse200211**](InlineResponse200211.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationAdaptivePolicyAcl

> DeleteOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Deletes the specified adaptive policy ACL



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
    aclId := "aclId_example" // string | Acl ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AclsApi.DeleteOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.DeleteOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyAclRequest struct via the builder pattern


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


## GetOrganizationAdaptivePolicyAcl

> InlineResponse200211 GetOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Returns the adaptive policy ACL information



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
    aclId := "aclId_example" // string | Acl ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AclsApi.GetOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.GetOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyAcl`: InlineResponse200211
    fmt.Fprintf(os.Stdout, "Response from `AclsApi.GetOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200211**](InlineResponse200211.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyAcls

> []InlineResponse200211 GetOrganizationAdaptivePolicyAcls(ctx, organizationId).Execute()

List adaptive policy ACLs in a organization



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
    resp, r, err := apiClient.AclsApi.GetOrganizationAdaptivePolicyAcls(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.GetOrganizationAdaptivePolicyAcls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyAcls`: []InlineResponse200211
    fmt.Fprintf(os.Stdout, "Response from `AclsApi.GetOrganizationAdaptivePolicyAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200211**](InlineResponse200211.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdaptivePolicyAcl

> InlineResponse200211 UpdateOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).UpdateOrganizationAdaptivePolicyAcl(updateOrganizationAdaptivePolicyAcl).Execute()

Updates an adaptive policy ACL



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
    aclId := "aclId_example" // string | Acl ID
    updateOrganizationAdaptivePolicyAcl := *openapiclient.NewInlineObject212() // InlineObject212 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AclsApi.UpdateOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).UpdateOrganizationAdaptivePolicyAcl(updateOrganizationAdaptivePolicyAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclsApi.UpdateOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyAcl`: InlineResponse200211
    fmt.Fprintf(os.Stdout, "Response from `AclsApi.UpdateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyAcl** | [**InlineObject212**](InlineObject212.md) |  | 

### Return type

[**InlineResponse200211**](InlineResponse200211.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

