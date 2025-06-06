# \SubscriptionApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindAdministeredLicensingSubscriptionSubscription**](SubscriptionApi.md#BindAdministeredLicensingSubscriptionSubscription) | **Post** /administered/licensing/subscription/subscriptions/{subscriptionId}/bind | Bind networks to a subscription
[**ClaimAdministeredLicensingSubscriptionSubscriptions**](SubscriptionApi.md#ClaimAdministeredLicensingSubscriptionSubscriptions) | **Post** /administered/licensing/subscription/subscriptions/claim | Claim a subscription into an organization.
[**GetAdministeredLicensingSubscriptionEntitlements**](SubscriptionApi.md#GetAdministeredLicensingSubscriptionEntitlements) | **Get** /administered/licensing/subscription/entitlements | Retrieve the list of purchasable entitlements
[**GetAdministeredLicensingSubscriptionSubscriptions**](SubscriptionApi.md#GetAdministeredLicensingSubscriptionSubscriptions) | **Get** /administered/licensing/subscription/subscriptions | List available subscriptions
[**GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses**](SubscriptionApi.md#GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses) | **Get** /administered/licensing/subscription/subscriptions/compliance/statuses | Get compliance status for requested subscriptions
[**ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey**](SubscriptionApi.md#ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey) | **Post** /administered/licensing/subscription/subscriptions/claimKey/validate | Find a subscription by claim key



## BindAdministeredLicensingSubscriptionSubscription

> InlineResponse2005 BindAdministeredLicensingSubscriptionSubscription(ctx, subscriptionId).Validate(validate).BindAdministeredLicensingSubscriptionSubscription(bindAdministeredLicensingSubscriptionSubscription).Execute()

Bind networks to a subscription



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
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    validate := true // bool | Check if the provided networks can be bound to the subscription. Returns any licensing problems and does not commit the results. (optional)
    bindAdministeredLicensingSubscriptionSubscription := *openapiclient.NewInlineObject2() // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.BindAdministeredLicensingSubscriptionSubscription(context.Background(), subscriptionId).Validate(validate).BindAdministeredLicensingSubscriptionSubscription(bindAdministeredLicensingSubscriptionSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.BindAdministeredLicensingSubscriptionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindAdministeredLicensingSubscriptionSubscription`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.BindAdministeredLicensingSubscriptionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindAdministeredLicensingSubscriptionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validate** | **bool** | Check if the provided networks can be bound to the subscription. Returns any licensing problems and does not commit the results. | 
 **bindAdministeredLicensingSubscriptionSubscription** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimAdministeredLicensingSubscriptionSubscriptions

> InlineResponse2003 ClaimAdministeredLicensingSubscriptionSubscriptions(ctx).ClaimAdministeredLicensingSubscriptionSubscriptions(claimAdministeredLicensingSubscriptionSubscriptions).Validate(validate).Execute()

Claim a subscription into an organization.



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
    claimAdministeredLicensingSubscriptionSubscriptions := *openapiclient.NewInlineObject("ClaimKey_example", "OrganizationId_example") // InlineObject | 
    validate := true // bool | Check if the provided claim key is valid and can be claimed into the organization. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.ClaimAdministeredLicensingSubscriptionSubscriptions(context.Background()).ClaimAdministeredLicensingSubscriptionSubscriptions(claimAdministeredLicensingSubscriptionSubscriptions).Validate(validate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ClaimAdministeredLicensingSubscriptionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimAdministeredLicensingSubscriptionSubscriptions`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ClaimAdministeredLicensingSubscriptionSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClaimAdministeredLicensingSubscriptionSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **claimAdministeredLicensingSubscriptionSubscriptions** | [**InlineObject**](InlineObject.md) |  | 
 **validate** | **bool** | Check if the provided claim key is valid and can be claimed into the organization. | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdministeredLicensingSubscriptionEntitlements

> InlineResponse2002 GetAdministeredLicensingSubscriptionEntitlements(ctx).Skus(skus).Execute()

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
    resp, r, err := apiClient.SubscriptionApi.GetAdministeredLicensingSubscriptionEntitlements(context.Background()).Skus(skus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetAdministeredLicensingSubscriptionEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredLicensingSubscriptionEntitlements`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetAdministeredLicensingSubscriptionEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredLicensingSubscriptionEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skus** | **[]string** | Filter to entitlements with the specified SKUs | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdministeredLicensingSubscriptionSubscriptions

> []InlineResponse2003 GetAdministeredLicensingSubscriptionSubscriptions(ctx).OrganizationIds(organizationIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SubscriptionIds(subscriptionIds).StartDate(startDate).EndDate(endDate).Statuses(statuses).ProductTypes(productTypes).Skus(skus).Name(name).Execute()

List available subscriptions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    organizationIds := []string{"Inner_example"} // []string | Organizations to get associated subscriptions for
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    subscriptionIds := []string{"Inner_example"} // []string | List of subscription ids to fetch (optional)
    startDate := time.Now() // time.Time | Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use 'startDate[<option>]=?' in the request. Accepted options include lt, gt, lte, gte. (optional)
    endDate := time.Now() // time.Time | Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use 'endDate[<option>]=?' in the request. Accepted options include lt, gt, lte, gte. (optional)
    statuses := []string{"Statuses_example"} // []string | List of statuses that returned subscriptions can have (optional)
    productTypes := []string{"ProductTypes_example"} // []string | List of product types that returned subscriptions need to have entitlements for. (optional)
    skus := []string{"Inner_example"} // []string | List of SKUs that returned subscriptions need to have entitlements for. (optional)
    name := "name_example" // string | Search for subscription name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetAdministeredLicensingSubscriptionSubscriptions(context.Background()).OrganizationIds(organizationIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SubscriptionIds(subscriptionIds).StartDate(startDate).EndDate(endDate).Statuses(statuses).ProductTypes(productTypes).Skus(skus).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetAdministeredLicensingSubscriptionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredLicensingSubscriptionSubscriptions`: []InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetAdministeredLicensingSubscriptionSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredLicensingSubscriptionSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationIds** | **[]string** | Organizations to get associated subscriptions for | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **subscriptionIds** | **[]string** | List of subscription ids to fetch | 
 **startDate** | **time.Time** | Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use &#39;startDate[&lt;option&gt;]&#x3D;?&#39; in the request. Accepted options include lt, gt, lte, gte. | 
 **endDate** | **time.Time** | Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use &#39;endDate[&lt;option&gt;]&#x3D;?&#39; in the request. Accepted options include lt, gt, lte, gte. | 
 **statuses** | **[]string** | List of statuses that returned subscriptions can have | 
 **productTypes** | **[]string** | List of product types that returned subscriptions need to have entitlements for. | 
 **skus** | **[]string** | List of SKUs that returned subscriptions need to have entitlements for. | 
 **name** | **string** | Search for subscription name | 

### Return type

[**[]InlineResponse2003**](InlineResponse2003.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses

> []InlineResponse2004 GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses(ctx).OrganizationIds(organizationIds).SubscriptionIds(subscriptionIds).Execute()

Get compliance status for requested subscriptions



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
    organizationIds := []string{"Inner_example"} // []string | Organizations to get subscription compliance information for
    subscriptionIds := []string{"Inner_example"} // []string | Subscription ids (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses(context.Background()).OrganizationIds(organizationIds).SubscriptionIds(subscriptionIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses`: []InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationIds** | **[]string** | Organizations to get subscription compliance information for | 
 **subscriptionIds** | **[]string** | Subscription ids | 

### Return type

[**[]InlineResponse2004**](InlineResponse2004.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey

> InlineResponse2003 ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(ctx).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(validateAdministeredLicensingSubscriptionSubscriptionsClaimKey).Execute()

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
    resp, r, err := apiClient.SubscriptionApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(context.Background()).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(validateAdministeredLicensingSubscriptionSubscriptionsClaimKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateAdministeredLicensingSubscriptionSubscriptionsClaimKey** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

