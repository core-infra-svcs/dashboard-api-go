# \ClaimKeyApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey**](ClaimKeyApi.md#ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey) | **Post** /administered/licensing/subscription/subscriptions/claimKey/validate | Find a subscription by claim key



## ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey

> InlineResponse2002 ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(ctx).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(validateAdministeredLicensingSubscriptionSubscriptionsClaimKey).Execute()

Find a subscription by claim key



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
    validateAdministeredLicensingSubscriptionSubscriptionsClaimKey := *openapiclient.NewInlineObject1("ClaimKey_example") // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClaimKeyApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(context.Background()).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(validateAdministeredLicensingSubscriptionSubscriptionsClaimKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClaimKeyApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `ClaimKeyApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateAdministeredLicensingSubscriptionSubscriptionsClaimKey** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

