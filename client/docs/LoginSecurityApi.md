# \LoginSecurityApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationLoginSecurity**](LoginSecurityApi.md#GetOrganizationLoginSecurity) | **Get** /organizations/{organizationId}/loginSecurity | Returns the login security settings for an organization.
[**UpdateOrganizationLoginSecurity**](LoginSecurityApi.md#UpdateOrganizationLoginSecurity) | **Put** /organizations/{organizationId}/loginSecurity | Update the login security settings for an organization



## GetOrganizationLoginSecurity

> InlineResponse200118 GetOrganizationLoginSecurity(ctx, organizationId).Execute()

Returns the login security settings for an organization.



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
    resp, r, err := apiClient.LoginSecurityApi.GetOrganizationLoginSecurity(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginSecurityApi.GetOrganizationLoginSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLoginSecurity`: InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `LoginSecurityApi.GetOrganizationLoginSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLoginSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200118**](InlineResponse200118.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationLoginSecurity

> InlineResponse200118 UpdateOrganizationLoginSecurity(ctx, organizationId).UpdateOrganizationLoginSecurity(updateOrganizationLoginSecurity).Execute()

Update the login security settings for an organization



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
    updateOrganizationLoginSecurity := *openapiclient.NewInlineObject209() // InlineObject209 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginSecurityApi.UpdateOrganizationLoginSecurity(context.Background(), organizationId).UpdateOrganizationLoginSecurity(updateOrganizationLoginSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginSecurityApi.UpdateOrganizationLoginSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationLoginSecurity`: InlineResponse200118
    fmt.Fprintf(os.Stdout, "Response from `LoginSecurityApi.UpdateOrganizationLoginSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationLoginSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationLoginSecurity** | [**InlineObject209**](InlineObject209.md) |  | 

### Return type

[**InlineResponse200118**](InlineResponse200118.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

