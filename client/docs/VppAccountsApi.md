# \VppAccountsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSmVppAccount**](VppAccountsApi.md#GetOrganizationSmVppAccount) | **Get** /organizations/{organizationId}/sm/vppAccounts/{vppAccountId} | Get a hash containing the unparsed token of the VPP account with the given ID
[**GetOrganizationSmVppAccounts**](VppAccountsApi.md#GetOrganizationSmVppAccounts) | **Get** /organizations/{organizationId}/sm/vppAccounts | List the VPP accounts in the organization



## GetOrganizationSmVppAccount

> InlineResponse200144 GetOrganizationSmVppAccount(ctx, organizationId, vppAccountId).Execute()

Get a hash containing the unparsed token of the VPP account with the given ID



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
    vppAccountId := "vppAccountId_example" // string | Vpp account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VppAccountsApi.GetOrganizationSmVppAccount(context.Background(), organizationId, vppAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VppAccountsApi.GetOrganizationSmVppAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmVppAccount`: InlineResponse200144
    fmt.Fprintf(os.Stdout, "Response from `VppAccountsApi.GetOrganizationSmVppAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**vppAccountId** | **string** | Vpp account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmVppAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200144**](InlineResponse200144.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmVppAccounts

> []InlineResponse200144 GetOrganizationSmVppAccounts(ctx, organizationId).Execute()

List the VPP accounts in the organization



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
    resp, r, err := apiClient.VppAccountsApi.GetOrganizationSmVppAccounts(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VppAccountsApi.GetOrganizationSmVppAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmVppAccounts`: []InlineResponse200144
    fmt.Fprintf(os.Stdout, "Response from `VppAccountsApi.GetOrganizationSmVppAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmVppAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200144**](InlineResponse200144.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

