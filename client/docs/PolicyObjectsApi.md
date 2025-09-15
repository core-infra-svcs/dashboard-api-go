# \PolicyObjectsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationPolicyObject**](PolicyObjectsApi.md#CreateOrganizationPolicyObject) | **Post** /organizations/{organizationId}/policyObjects | Creates a new Policy Object.
[**CreateOrganizationPolicyObjectsGroup**](PolicyObjectsApi.md#CreateOrganizationPolicyObjectsGroup) | **Post** /organizations/{organizationId}/policyObjects/groups | Creates a new Policy Object Group.
[**DeleteOrganizationPolicyObject**](PolicyObjectsApi.md#DeleteOrganizationPolicyObject) | **Delete** /organizations/{organizationId}/policyObjects/{policyObjectId} | Deletes a Policy Object.
[**DeleteOrganizationPolicyObjectsGroup**](PolicyObjectsApi.md#DeleteOrganizationPolicyObjectsGroup) | **Delete** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Deletes a Policy Object Group.
[**GetOrganizationPolicyObject**](PolicyObjectsApi.md#GetOrganizationPolicyObject) | **Get** /organizations/{organizationId}/policyObjects/{policyObjectId} | Shows details of a Policy Object.
[**GetOrganizationPolicyObjects**](PolicyObjectsApi.md#GetOrganizationPolicyObjects) | **Get** /organizations/{organizationId}/policyObjects | Lists Policy Objects belonging to the organization.
[**GetOrganizationPolicyObjectsGroup**](PolicyObjectsApi.md#GetOrganizationPolicyObjectsGroup) | **Get** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Shows details of a Policy Object Group.
[**GetOrganizationPolicyObjectsGroups**](PolicyObjectsApi.md#GetOrganizationPolicyObjectsGroups) | **Get** /organizations/{organizationId}/policyObjects/groups | Lists Policy Object Groups belonging to the organization.
[**UpdateOrganizationPolicyObject**](PolicyObjectsApi.md#UpdateOrganizationPolicyObject) | **Put** /organizations/{organizationId}/policyObjects/{policyObjectId} | Updates a Policy Object.
[**UpdateOrganizationPolicyObjectsGroup**](PolicyObjectsApi.md#UpdateOrganizationPolicyObjectsGroup) | **Put** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Updates a Policy Object Group.



## CreateOrganizationPolicyObject

> InlineResponse200313 CreateOrganizationPolicyObject(ctx, organizationId).CreateOrganizationPolicyObject(createOrganizationPolicyObject).Execute()

Creates a new Policy Object.



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
    createOrganizationPolicyObject := *openapiclient.NewInlineObject287("Name_example", "Category_example", "Type_example") // InlineObject287 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.CreateOrganizationPolicyObject(context.Background(), organizationId).CreateOrganizationPolicyObject(createOrganizationPolicyObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.CreateOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPolicyObject`: InlineResponse200313
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.CreateOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPolicyObject** | [**InlineObject287**](InlineObject287.md) |  | 

### Return type

[**InlineResponse200313**](InlineResponse200313.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationPolicyObjectsGroup

> InlineResponse200314 CreateOrganizationPolicyObjectsGroup(ctx, organizationId).CreateOrganizationPolicyObjectsGroup(createOrganizationPolicyObjectsGroup).Execute()

Creates a new Policy Object Group.



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
    createOrganizationPolicyObjectsGroup := *openapiclient.NewInlineObject288("Name_example") // InlineObject288 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.CreateOrganizationPolicyObjectsGroup(context.Background(), organizationId).CreateOrganizationPolicyObjectsGroup(createOrganizationPolicyObjectsGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.CreateOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPolicyObjectsGroup`: InlineResponse200314
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.CreateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPolicyObjectsGroup** | [**InlineObject288**](InlineObject288.md) |  | 

### Return type

[**InlineResponse200314**](InlineResponse200314.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPolicyObject

> DeleteOrganizationPolicyObject(ctx, organizationId, policyObjectId).Execute()

Deletes a Policy Object.



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
    policyObjectId := "policyObjectId_example" // string | Policy object ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.DeleteOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.DeleteOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyObjectRequest struct via the builder pattern


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


## DeleteOrganizationPolicyObjectsGroup

> DeleteOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Deletes a Policy Object Group.



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
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.DeleteOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.DeleteOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyObjectsGroupRequest struct via the builder pattern


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


## GetOrganizationPolicyObject

> InlineResponse200313 GetOrganizationPolicyObject(ctx, organizationId, policyObjectId).Execute()

Shows details of a Policy Object.



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
    policyObjectId := "policyObjectId_example" // string | Policy object ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.GetOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.GetOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObject`: InlineResponse200313
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.GetOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200313**](InlineResponse200313.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjects

> InlineResponse200313 GetOrganizationPolicyObjects(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Lists Policy Objects belonging to the organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.GetOrganizationPolicyObjects(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.GetOrganizationPolicyObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjects`: InlineResponse200313
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.GetOrganizationPolicyObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 10 - 5000. Default is 5000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200313**](InlineResponse200313.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjectsGroup

> InlineResponse200314 GetOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Shows details of a Policy Object Group.



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
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.GetOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.GetOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjectsGroup`: InlineResponse200314
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.GetOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200314**](InlineResponse200314.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationPolicyObjectsGroups

> InlineResponse200314 GetOrganizationPolicyObjectsGroups(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Lists Policy Object Groups belonging to the organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.GetOrganizationPolicyObjectsGroups(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.GetOrganizationPolicyObjectsGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjectsGroups`: InlineResponse200314
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.GetOrganizationPolicyObjectsGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200314**](InlineResponse200314.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPolicyObject

> InlineResponse200313 UpdateOrganizationPolicyObject(ctx, organizationId, policyObjectId).UpdateOrganizationPolicyObject(updateOrganizationPolicyObject).Execute()

Updates a Policy Object.



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
    policyObjectId := "policyObjectId_example" // string | Policy object ID
    updateOrganizationPolicyObject := *openapiclient.NewInlineObject290() // InlineObject290 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.UpdateOrganizationPolicyObject(context.Background(), organizationId, policyObjectId).UpdateOrganizationPolicyObject(updateOrganizationPolicyObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.UpdateOrganizationPolicyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPolicyObject`: InlineResponse200313
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.UpdateOrganizationPolicyObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectId** | **string** | Policy object ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPolicyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPolicyObject** | [**InlineObject290**](InlineObject290.md) |  | 

### Return type

[**InlineResponse200313**](InlineResponse200313.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationPolicyObjectsGroup

> InlineResponse200314 UpdateOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroup(updateOrganizationPolicyObjectsGroup).Execute()

Updates a Policy Object Group.



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
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID
    updateOrganizationPolicyObjectsGroup := *openapiclient.NewInlineObject289() // InlineObject289 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyObjectsApi.UpdateOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroup(updateOrganizationPolicyObjectsGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyObjectsApi.UpdateOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPolicyObjectsGroup`: InlineResponse200314
    fmt.Fprintf(os.Stdout, "Response from `PolicyObjectsApi.UpdateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPolicyObjectsGroup** | [**InlineObject289**](InlineObject289.md) |  | 

### Return type

[**InlineResponse200314**](InlineResponse200314.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

