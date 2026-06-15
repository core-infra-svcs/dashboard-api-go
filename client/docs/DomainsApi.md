# \DomainsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExclusionsNetworkApplianceUmbrellaDomains**](DomainsApi.md#ExclusionsNetworkApplianceUmbrellaDomains) | **Put** /networks/{networkId}/appliance/umbrella/domains/exclusions | Specify one or more domain names to be excluded from being routed to Cisco Umbrella.



## ExclusionsNetworkApplianceUmbrellaDomains

> InlineResponse200102 ExclusionsNetworkApplianceUmbrellaDomains(ctx, networkId).ExclusionsNetworkApplianceUmbrellaDomains(exclusionsNetworkApplianceUmbrellaDomains).Execute()

Specify one or more domain names to be excluded from being routed to Cisco Umbrella.



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
    networkId := "networkId_example" // string | Network ID
    exclusionsNetworkApplianceUmbrellaDomains := *openapiclient.NewInlineObject87([]string{"Domains_example"}) // InlineObject87 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsApi.ExclusionsNetworkApplianceUmbrellaDomains(context.Background(), networkId).ExclusionsNetworkApplianceUmbrellaDomains(exclusionsNetworkApplianceUmbrellaDomains).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.ExclusionsNetworkApplianceUmbrellaDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExclusionsNetworkApplianceUmbrellaDomains`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.ExclusionsNetworkApplianceUmbrellaDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExclusionsNetworkApplianceUmbrellaDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclusionsNetworkApplianceUmbrellaDomains** | [**InlineObject87**](InlineObject87.md) |  | 

### Return type

[**InlineResponse200102**](InlineResponse200102.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

