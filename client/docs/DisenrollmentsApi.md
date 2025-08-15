# \DisenrollmentsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationWirelessZigbeeDisenrollment**](DisenrollmentsApi.md#CreateOrganizationWirelessZigbeeDisenrollment) | **Post** /organizations/{organizationId}/wireless/zigbee/disenrollments | Enqueue a job to start disenrolling door locks on zigbee configured wireless devices
[**GetOrganizationWirelessZigbeeDisenrollment**](DisenrollmentsApi.md#GetOrganizationWirelessZigbeeDisenrollment) | **Get** /organizations/{organizationId}/wireless/zigbee/disenrollments/{disenrollmentId} | Return a disenrollment



## CreateOrganizationWirelessZigbeeDisenrollment

> InlineResponse20123 CreateOrganizationWirelessZigbeeDisenrollment(ctx, organizationId).CreateOrganizationWirelessZigbeeDisenrollment(createOrganizationWirelessZigbeeDisenrollment).Execute()

Enqueue a job to start disenrolling door locks on zigbee configured wireless devices



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
    createOrganizationWirelessZigbeeDisenrollment := *openapiclient.NewInlineObject309() // InlineObject309 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisenrollmentsApi.CreateOrganizationWirelessZigbeeDisenrollment(context.Background(), organizationId).CreateOrganizationWirelessZigbeeDisenrollment(createOrganizationWirelessZigbeeDisenrollment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisenrollmentsApi.CreateOrganizationWirelessZigbeeDisenrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessZigbeeDisenrollment`: InlineResponse20123
    fmt.Fprintf(os.Stdout, "Response from `DisenrollmentsApi.CreateOrganizationWirelessZigbeeDisenrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessZigbeeDisenrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationWirelessZigbeeDisenrollment** | [**InlineObject309**](InlineObject309.md) |  | 

### Return type

[**InlineResponse20123**](InlineResponse20123.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessZigbeeDisenrollment

> InlineResponse200374 GetOrganizationWirelessZigbeeDisenrollment(ctx, organizationId, disenrollmentId).Execute()

Return a disenrollment



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
    disenrollmentId := "disenrollmentId_example" // string | Disenrollment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DisenrollmentsApi.GetOrganizationWirelessZigbeeDisenrollment(context.Background(), organizationId, disenrollmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisenrollmentsApi.GetOrganizationWirelessZigbeeDisenrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessZigbeeDisenrollment`: InlineResponse200374
    fmt.Fprintf(os.Stdout, "Response from `DisenrollmentsApi.GetOrganizationWirelessZigbeeDisenrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**disenrollmentId** | **string** | Disenrollment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessZigbeeDisenrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200374**](InlineResponse200374.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

