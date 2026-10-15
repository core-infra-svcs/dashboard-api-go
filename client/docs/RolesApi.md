# \RolesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCameraRole**](RolesApi.md#CreateOrganizationCameraRole) | **Post** /organizations/{organizationId}/camera/roles | Creates new role for this organization.
[**CreateOrganizationSmAdminsRole**](RolesApi.md#CreateOrganizationSmAdminsRole) | **Post** /organizations/{organizationId}/sm/admins/roles | Create a Limited Access Role
[**DeleteOrganizationCameraRole**](RolesApi.md#DeleteOrganizationCameraRole) | **Delete** /organizations/{organizationId}/camera/roles/{roleId} | Delete an existing role for this organization.
[**DeleteOrganizationSmAdminsRole**](RolesApi.md#DeleteOrganizationSmAdminsRole) | **Delete** /organizations/{organizationId}/sm/admins/roles/{roleId} | Delete a Limited Access Role
[**GetOrganizationCameraRole**](RolesApi.md#GetOrganizationCameraRole) | **Get** /organizations/{organizationId}/camera/roles/{roleId} | Retrieve a single role.
[**GetOrganizationCameraRoles**](RolesApi.md#GetOrganizationCameraRoles) | **Get** /organizations/{organizationId}/camera/roles | List all the roles in this organization
[**GetOrganizationDevicesSyslogServersRolesByNetwork**](RolesApi.md#GetOrganizationDevicesSyslogServersRolesByNetwork) | **Get** /organizations/{organizationId}/devices/syslog/servers/roles/byNetwork | Returns roles that can be assigned to a syslog server for a given network.
[**GetOrganizationSmAdminsRole**](RolesApi.md#GetOrganizationSmAdminsRole) | **Get** /organizations/{organizationId}/sm/admins/roles/{roleId} | Return a Limited Access Role
[**GetOrganizationSmAdminsRoles**](RolesApi.md#GetOrganizationSmAdminsRoles) | **Get** /organizations/{organizationId}/sm/admins/roles | List the Limited Access Roles for an organization
[**UpdateOrganizationCameraRole**](RolesApi.md#UpdateOrganizationCameraRole) | **Put** /organizations/{organizationId}/camera/roles/{roleId} | Update an existing role in this organization.
[**UpdateOrganizationSmAdminsRole**](RolesApi.md#UpdateOrganizationSmAdminsRole) | **Put** /organizations/{organizationId}/sm/admins/roles/{roleId} | Update a Limited Access Role



## CreateOrganizationCameraRole

> InlineResponse200330 CreateOrganizationCameraRole(ctx, organizationId).CreateOrganizationCameraRole(createOrganizationCameraRole).Execute()

Creates new role for this organization.



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
    createOrganizationCameraRole := *openapiclient.NewInlineObject282("Name_example") // InlineObject282 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.CreateOrganizationCameraRole(context.Background(), organizationId).CreateOrganizationCameraRole(createOrganizationCameraRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCameraRole`: InlineResponse200330
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCameraRole** | [**InlineObject282**](InlineObject282.md) |  | 

### Return type

[**InlineResponse200330**](InlineResponse200330.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSmAdminsRole

> InlineResponse200447Items CreateOrganizationSmAdminsRole(ctx, organizationId).CreateOrganizationSmAdminsRole(createOrganizationSmAdminsRole).Execute()

Create a Limited Access Role



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
    createOrganizationSmAdminsRole := *openapiclient.NewInlineObject363("Name_example") // InlineObject363 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.CreateOrganizationSmAdminsRole(context.Background(), organizationId).CreateOrganizationSmAdminsRole(createOrganizationSmAdminsRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateOrganizationSmAdminsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSmAdminsRole`: InlineResponse200447Items
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSmAdminsRole** | [**InlineObject363**](InlineObject363.md) |  | 

### Return type

[**InlineResponse200447Items**](InlineResponse200447Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationCameraRole

> DeleteOrganizationCameraRole(ctx, organizationId, roleId).Execute()

Delete an existing role for this organization.



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
    roleId := "roleId_example" // string | Role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.DeleteOrganizationCameraRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCameraRoleRequest struct via the builder pattern


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


## DeleteOrganizationSmAdminsRole

> DeleteOrganizationSmAdminsRole(ctx, organizationId, roleId).Execute()

Delete a Limited Access Role



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
    roleId := "roleId_example" // string | Role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.DeleteOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteOrganizationSmAdminsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSmAdminsRoleRequest struct via the builder pattern


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


## GetOrganizationCameraRole

> InlineResponse200330 GetOrganizationCameraRole(ctx, organizationId, roleId).Execute()

Retrieve a single role.



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
    roleId := "roleId_example" // string | Role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetOrganizationCameraRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraRole`: InlineResponse200330
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200330**](InlineResponse200330.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraRoles

> []InlineResponse200330 GetOrganizationCameraRoles(ctx, organizationId).Execute()

List all the roles in this organization



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
    resp, r, err := apiClient.RolesApi.GetOrganizationCameraRoles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationCameraRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraRoles`: []InlineResponse200330
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationCameraRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200330**](InlineResponse200330.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesSyslogServersRolesByNetwork

> InlineResponse200384 GetOrganizationDevicesSyslogServersRolesByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Returns roles that can be assigned to a syslog server for a given network.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | IDs of the networks for which to fetch valid syslog server roles; suggested maximum array size is 100 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetOrganizationDevicesSyslogServersRolesByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationDevicesSyslogServersRolesByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesSyslogServersRolesByNetwork`: InlineResponse200384
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationDevicesSyslogServersRolesByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesSyslogServersRolesByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | IDs of the networks for which to fetch valid syslog server roles; suggested maximum array size is 100 | 

### Return type

[**InlineResponse200384**](InlineResponse200384.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmAdminsRole

> InlineResponse200447Items GetOrganizationSmAdminsRole(ctx, organizationId, roleId).Execute()

Return a Limited Access Role



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
    roleId := "roleId_example" // string | Role ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationSmAdminsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmAdminsRole`: InlineResponse200447Items
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200447Items**](InlineResponse200447Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmAdminsRoles

> InlineResponse200447 GetOrganizationSmAdminsRoles(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the Limited Access Roles for an organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetOrganizationSmAdminsRoles(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationSmAdminsRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmAdminsRoles`: InlineResponse200447
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationSmAdminsRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmAdminsRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200447**](InlineResponse200447.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCameraRole

> InlineResponse200330 UpdateOrganizationCameraRole(ctx, organizationId, roleId).UpdateOrganizationCameraRole(updateOrganizationCameraRole).Execute()

Update an existing role in this organization.



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
    roleId := "roleId_example" // string | Role ID
    updateOrganizationCameraRole := *openapiclient.NewInlineObject283() // InlineObject283 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.UpdateOrganizationCameraRole(context.Background(), organizationId, roleId).UpdateOrganizationCameraRole(updateOrganizationCameraRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCameraRole`: InlineResponse200330
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationCameraRole** | [**InlineObject283**](InlineObject283.md) |  | 

### Return type

[**InlineResponse200330**](InlineResponse200330.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSmAdminsRole

> InlineResponse200447Items UpdateOrganizationSmAdminsRole(ctx, organizationId, roleId).UpdateOrganizationSmAdminsRole(updateOrganizationSmAdminsRole).Execute()

Update a Limited Access Role



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
    roleId := "roleId_example" // string | Role ID
    updateOrganizationSmAdminsRole := *openapiclient.NewInlineObject364() // InlineObject364 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.UpdateOrganizationSmAdminsRole(context.Background(), organizationId, roleId).UpdateOrganizationSmAdminsRole(updateOrganizationSmAdminsRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateOrganizationSmAdminsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSmAdminsRole`: InlineResponse200447Items
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSmAdminsRole** | [**InlineObject364**](InlineObject364.md) |  | 

### Return type

[**InlineResponse200447Items**](InlineResponse200447Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

