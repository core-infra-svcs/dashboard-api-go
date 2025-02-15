# \CommunicationPlansApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans**](CommunicationPlansApi.md#GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans) | **Get** /organizations/{organizationId}/cellularGateway/esims/serviceProviders/accounts/communicationPlans | The communication plans available for a given provider.



## GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans

> InlineResponse200247 GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(ctx, organizationId).AccountIds(accountIds).Execute()

The communication plans available for a given provider.



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
    accountIds := []string{"Inner_example"} // []string | Account IDs that communication plans will be fetched for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommunicationPlansApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans(context.Background(), organizationId).AccountIds(accountIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationPlansApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: InlineResponse200247
    fmt.Fprintf(os.Stdout, "Response from `CommunicationPlansApi.GetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIds** | **[]string** | Account IDs that communication plans will be fetched for | 

### Return type

[**InlineResponse200247**](InlineResponse200247.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

