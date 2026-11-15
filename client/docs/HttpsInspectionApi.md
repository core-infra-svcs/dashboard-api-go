# \HttpsInspectionApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceSecurityHttpsInspectionByNetwork**](HttpsInspectionApi.md#GetOrganizationApplianceSecurityHttpsInspectionByNetwork) | **Get** /organizations/{organizationId}/appliance/security/httpsInspection/byNetwork | List HTTPS inspection status and configuration for appliances in the organization.
[**UpdateNetworkApplianceSecurityHttpsInspection**](HttpsInspectionApi.md#UpdateNetworkApplianceSecurityHttpsInspection) | **Put** /networks/{networkId}/appliance/security/httpsInspection | Update HTTPS inspection configuration for a network.



## GetOrganizationApplianceSecurityHttpsInspectionByNetwork

> InlineResponse200295 GetOrganizationApplianceSecurityHttpsInspectionByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List HTTPS inspection status and configuration for appliances in the organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned list will be filtered to only include these networks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpsInspectionApi.GetOrganizationApplianceSecurityHttpsInspectionByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpsInspectionApi.GetOrganizationApplianceSecurityHttpsInspectionByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceSecurityHttpsInspectionByNetwork`: InlineResponse200295
    fmt.Fprintf(os.Stdout, "Response from `HttpsInspectionApi.GetOrganizationApplianceSecurityHttpsInspectionByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityHttpsInspectionByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned list will be filtered to only include these networks. | 

### Return type

[**InlineResponse200295**](InlineResponse200295.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityHttpsInspection

> InlineResponse20091 UpdateNetworkApplianceSecurityHttpsInspection(ctx, networkId).UpdateNetworkApplianceSecurityHttpsInspection(updateNetworkApplianceSecurityHttpsInspection).Execute()

Update HTTPS inspection configuration for a network.



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
    updateNetworkApplianceSecurityHttpsInspection := *openapiclient.NewInlineObject76(false) // InlineObject76 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpsInspectionApi.UpdateNetworkApplianceSecurityHttpsInspection(context.Background(), networkId).UpdateNetworkApplianceSecurityHttpsInspection(updateNetworkApplianceSecurityHttpsInspection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpsInspectionApi.UpdateNetworkApplianceSecurityHttpsInspection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSecurityHttpsInspection`: InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `HttpsInspectionApi.UpdateNetworkApplianceSecurityHttpsInspection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityHttpsInspectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityHttpsInspection** | [**InlineObject76**](InlineObject76.md) |  | 

### Return type

[**InlineResponse20091**](InlineResponse20091.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

