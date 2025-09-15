# \EnrollmentsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceWirelessZigbeeEnrollment**](EnrollmentsApi.md#CreateDeviceWirelessZigbeeEnrollment) | **Post** /devices/{serial}/wireless/zigbee/enrollments | Enqueue a job to start enrolling door locks on zigbee configured wireless devices
[**GetDeviceWirelessZigbeeEnrollment**](EnrollmentsApi.md#GetDeviceWirelessZigbeeEnrollment) | **Get** /devices/{serial}/wireless/zigbee/enrollments/{enrollmentId} | Return an enrollment



## CreateDeviceWirelessZigbeeEnrollment

> InlineResponse20110 CreateDeviceWirelessZigbeeEnrollment(ctx, serial).Execute()

Enqueue a job to start enrolling door locks on zigbee configured wireless devices



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentsApi.CreateDeviceWirelessZigbeeEnrollment(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsApi.CreateDeviceWirelessZigbeeEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceWirelessZigbeeEnrollment`: InlineResponse20110
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentsApi.CreateDeviceWirelessZigbeeEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceWirelessZigbeeEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20110**](InlineResponse20110.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessZigbeeEnrollment

> InlineResponse20049 GetDeviceWirelessZigbeeEnrollment(ctx, serial, enrollmentId).Execute()

Return an enrollment



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
    serial := "serial_example" // string | Serial
    enrollmentId := "enrollmentId_example" // string | Enrollment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnrollmentsApi.GetDeviceWirelessZigbeeEnrollment(context.Background(), serial, enrollmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentsApi.GetDeviceWirelessZigbeeEnrollment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessZigbeeEnrollment`: InlineResponse20049
    fmt.Fprintf(os.Stdout, "Response from `EnrollmentsApi.GetDeviceWirelessZigbeeEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**enrollmentId** | **string** | Enrollment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessZigbeeEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20049**](InlineResponse20049.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

