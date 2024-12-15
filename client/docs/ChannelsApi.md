# \ChannelsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecalculateOrganizationWirelessRadioAutoRfChannels**](ChannelsApi.md#RecalculateOrganizationWirelessRadioAutoRfChannels) | **Post** /organizations/{organizationId}/wireless/radio/autoRf/channels/recalculate | Recalculates automatically assigned channels for every AP within specified the specified network(s)



## RecalculateOrganizationWirelessRadioAutoRfChannels

> InlineResponse200326 RecalculateOrganizationWirelessRadioAutoRfChannels(ctx, organizationId).RecalculateOrganizationWirelessRadioAutoRfChannels(recalculateOrganizationWirelessRadioAutoRfChannels).Execute()

Recalculates automatically assigned channels for every AP within specified the specified network(s)



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
    recalculateOrganizationWirelessRadioAutoRfChannels := *openapiclient.NewInlineObject274([]string{"NetworkIds_example"}) // InlineObject274 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelsApi.RecalculateOrganizationWirelessRadioAutoRfChannels(context.Background(), organizationId).RecalculateOrganizationWirelessRadioAutoRfChannels(recalculateOrganizationWirelessRadioAutoRfChannels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.RecalculateOrganizationWirelessRadioAutoRfChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecalculateOrganizationWirelessRadioAutoRfChannels`: InlineResponse200326
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.RecalculateOrganizationWirelessRadioAutoRfChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecalculateOrganizationWirelessRadioAutoRfChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recalculateOrganizationWirelessRadioAutoRfChannels** | [**InlineObject274**](InlineObject274.md) |  | 

### Return type

[**InlineResponse200326**](InlineResponse200326.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

