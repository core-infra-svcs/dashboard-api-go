# \TunnelableApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCampusGatewayClustersTunnelable**](TunnelableApi.md#GetOrganizationCampusGatewayClustersTunnelable) | **Get** /organizations/{organizationId}/campusGateway/clusters/tunnelable | List available campus gateway clusters that can be used for wireless network tunneling



## GetOrganizationCampusGatewayClustersTunnelable

> InlineResponse200337 GetOrganizationCampusGatewayClustersTunnelable(ctx, organizationId).FromNetworkIds(fromNetworkIds).ExcludeSsidNumbers(excludeSsidNumbers).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List available campus gateway clusters that can be used for wireless network tunneling



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
    fromNetworkIds := []string{"Inner_example"} // []string | Filter results by wireless network IDs.
    excludeSsidNumbers := []string{"Inner_example"} // []string | Optional SSID numbers (0-15) to exclude from restriction checks, corresponding to each fromNetworkIds entry. Only one SSID per network can be excluded. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 200. Default is 200. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TunnelableApi.GetOrganizationCampusGatewayClustersTunnelable(context.Background(), organizationId).FromNetworkIds(fromNetworkIds).ExcludeSsidNumbers(excludeSsidNumbers).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TunnelableApi.GetOrganizationCampusGatewayClustersTunnelable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCampusGatewayClustersTunnelable`: InlineResponse200337
    fmt.Fprintf(os.Stdout, "Response from `TunnelableApi.GetOrganizationCampusGatewayClustersTunnelable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCampusGatewayClustersTunnelableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromNetworkIds** | **[]string** | Filter results by wireless network IDs. | 
 **excludeSsidNumbers** | **[]string** | Optional SSID numbers (0-15) to exclude from restriction checks, corresponding to each fromNetworkIds entry. Only one SSID per network can be excluded. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 200. Default is 200. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200337**](InlineResponse200337.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

