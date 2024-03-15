# \EntitlementsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdministeredLicensingSubscriptionEntitlements**](EntitlementsApi.md#GetAdministeredLicensingSubscriptionEntitlements) | **Get** /administered/licensing/subscription/entitlements | Retrieve the list of purchasable entitlements



## GetAdministeredLicensingSubscriptionEntitlements

> InlineResponse2001 GetAdministeredLicensingSubscriptionEntitlements(ctx).Skus(skus).Execute()

Retrieve the list of purchasable entitlements



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
    skus := []string{"Inner_example"} // []string | Filter to entitlements with the specified SKUs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.GetAdministeredLicensingSubscriptionEntitlements(context.Background()).Skus(skus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetAdministeredLicensingSubscriptionEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredLicensingSubscriptionEntitlements`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetAdministeredLicensingSubscriptionEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredLicensingSubscriptionEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skus** | **[]string** | Filter to entitlements with the specified SKUs | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

