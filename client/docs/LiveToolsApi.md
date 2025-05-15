# \LiveToolsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlinkDeviceLeds**](LiveToolsApi.md#BlinkDeviceLeds) | **Post** /devices/{serial}/blinkLeds | Blink the LEDs on a device
[**CreateDeviceLiveToolsArpTable**](LiveToolsApi.md#CreateDeviceLiveToolsArpTable) | **Post** /devices/{serial}/liveTools/arpTable | Enqueue a job to perform a ARP table request for the device
[**CreateDeviceLiveToolsCableTest**](LiveToolsApi.md#CreateDeviceLiveToolsCableTest) | **Post** /devices/{serial}/liveTools/cableTest | Enqueue a job to perform a cable test for the device on the specified ports
[**CreateDeviceLiveToolsLedsBlink**](LiveToolsApi.md#CreateDeviceLiveToolsLedsBlink) | **Post** /devices/{serial}/liveTools/leds/blink | Enqueue a job to blink LEDs on a device
[**CreateDeviceLiveToolsMacTable**](LiveToolsApi.md#CreateDeviceLiveToolsMacTable) | **Post** /devices/{serial}/liveTools/macTable | Enqueue a job to request the MAC table from the device
[**CreateDeviceLiveToolsPing**](LiveToolsApi.md#CreateDeviceLiveToolsPing) | **Post** /devices/{serial}/liveTools/ping | Enqueue a job to ping a target host from the device
[**CreateDeviceLiveToolsPingDevice**](LiveToolsApi.md#CreateDeviceLiveToolsPingDevice) | **Post** /devices/{serial}/liveTools/pingDevice | Enqueue a job to check connectivity status to the device
[**CreateDeviceLiveToolsThroughputTest**](LiveToolsApi.md#CreateDeviceLiveToolsThroughputTest) | **Post** /devices/{serial}/liveTools/throughputTest | Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput
[**CreateDeviceLiveToolsWakeOnLan**](LiveToolsApi.md#CreateDeviceLiveToolsWakeOnLan) | **Post** /devices/{serial}/liveTools/wakeOnLan | Enqueue a job to send a Wake-on-LAN packet from the device
[**CycleDeviceSwitchPorts**](LiveToolsApi.md#CycleDeviceSwitchPorts) | **Post** /devices/{serial}/switch/ports/cycle | Cycle a set of switch ports
[**GetDeviceLiveToolsArpTable**](LiveToolsApi.md#GetDeviceLiveToolsArpTable) | **Get** /devices/{serial}/liveTools/arpTable/{arpTableId} | Return an ARP table live tool job.
[**GetDeviceLiveToolsCableTest**](LiveToolsApi.md#GetDeviceLiveToolsCableTest) | **Get** /devices/{serial}/liveTools/cableTest/{id} | Return a cable test live tool job.
[**GetDeviceLiveToolsLedsBlink**](LiveToolsApi.md#GetDeviceLiveToolsLedsBlink) | **Get** /devices/{serial}/liveTools/leds/blink/{ledsBlinkId} | Return a blink LEDs job
[**GetDeviceLiveToolsMacTable**](LiveToolsApi.md#GetDeviceLiveToolsMacTable) | **Get** /devices/{serial}/liveTools/macTable/{macTableId} | Return a MAC table live tool job.
[**GetDeviceLiveToolsPing**](LiveToolsApi.md#GetDeviceLiveToolsPing) | **Get** /devices/{serial}/liveTools/ping/{id} | Return a ping job
[**GetDeviceLiveToolsPingDevice**](LiveToolsApi.md#GetDeviceLiveToolsPingDevice) | **Get** /devices/{serial}/liveTools/pingDevice/{id} | Return a ping device job
[**GetDeviceLiveToolsThroughputTest**](LiveToolsApi.md#GetDeviceLiveToolsThroughputTest) | **Get** /devices/{serial}/liveTools/throughputTest/{throughputTestId} | Return a throughput test job
[**GetDeviceLiveToolsWakeOnLan**](LiveToolsApi.md#GetDeviceLiveToolsWakeOnLan) | **Get** /devices/{serial}/liveTools/wakeOnLan/{wakeOnLanId} | Return a Wake-on-LAN job
[**RebootDevice**](LiveToolsApi.md#RebootDevice) | **Post** /devices/{serial}/reboot | Reboot a device



## BlinkDeviceLeds

> InlineResponse2021 BlinkDeviceLeds(ctx, serial).BlinkDeviceLeds(blinkDeviceLeds).Execute()

Blink the LEDs on a device



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
    blinkDeviceLeds := *openapiclient.NewInlineObject6() // InlineObject6 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.BlinkDeviceLeds(context.Background(), serial).BlinkDeviceLeds(blinkDeviceLeds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.BlinkDeviceLeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlinkDeviceLeds`: InlineResponse2021
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.BlinkDeviceLeds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlinkDeviceLedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blinkDeviceLeds** | [**InlineObject6**](InlineObject6.md) |  | 

### Return type

[**InlineResponse2021**](InlineResponse2021.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsArpTable

> InlineResponse2011 CreateDeviceLiveToolsArpTable(ctx, serial).CreateDeviceLiveToolsArpTable(createDeviceLiveToolsArpTable).Execute()

Enqueue a job to perform a ARP table request for the device



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
    createDeviceLiveToolsArpTable := *openapiclient.NewInlineObject16() // InlineObject16 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsArpTable(context.Background(), serial).CreateDeviceLiveToolsArpTable(createDeviceLiveToolsArpTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsArpTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsArpTable`: InlineResponse2011
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTable** | [**InlineObject16**](InlineObject16.md) |  | 

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsCableTest

> InlineResponse2012 CreateDeviceLiveToolsCableTest(ctx, serial).CreateDeviceLiveToolsCableTest(createDeviceLiveToolsCableTest).Execute()

Enqueue a job to perform a cable test for the device on the specified ports



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
    createDeviceLiveToolsCableTest := *openapiclient.NewInlineObject17([]string{"Ports_example"}) // InlineObject17 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsCableTest(context.Background(), serial).CreateDeviceLiveToolsCableTest(createDeviceLiveToolsCableTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsCableTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsCableTest`: InlineResponse2012
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsCableTest** | [**InlineObject17**](InlineObject17.md) |  | 

### Return type

[**InlineResponse2012**](InlineResponse2012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsLedsBlink

> InlineResponse2013 CreateDeviceLiveToolsLedsBlink(ctx, serial).CreateDeviceLiveToolsLedsBlink(createDeviceLiveToolsLedsBlink).Execute()

Enqueue a job to blink LEDs on a device



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
    createDeviceLiveToolsLedsBlink := *openapiclient.NewInlineObject18(int32(123)) // InlineObject18 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsLedsBlink(context.Background(), serial).CreateDeviceLiveToolsLedsBlink(createDeviceLiveToolsLedsBlink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsLedsBlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsLedsBlink`: InlineResponse2013
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsLedsBlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsLedsBlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsLedsBlink** | [**InlineObject18**](InlineObject18.md) |  | 

### Return type

[**InlineResponse2013**](InlineResponse2013.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsMacTable

> InlineResponse2014 CreateDeviceLiveToolsMacTable(ctx, serial).CreateDeviceLiveToolsMacTable(createDeviceLiveToolsMacTable).Execute()

Enqueue a job to request the MAC table from the device



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
    createDeviceLiveToolsMacTable := *openapiclient.NewInlineObject19() // InlineObject19 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsMacTable(context.Background(), serial).CreateDeviceLiveToolsMacTable(createDeviceLiveToolsMacTable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsMacTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsMacTable`: InlineResponse2014
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsMacTable** | [**InlineObject19**](InlineObject19.md) |  | 

### Return type

[**InlineResponse2014**](InlineResponse2014.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPing

> InlineResponse2015 CreateDeviceLiveToolsPing(ctx, serial).CreateDeviceLiveToolsPing(createDeviceLiveToolsPing).Execute()

Enqueue a job to ping a target host from the device



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
    createDeviceLiveToolsPing := *openapiclient.NewInlineObject20("Target_example") // InlineObject20 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsPing(context.Background(), serial).CreateDeviceLiveToolsPing(createDeviceLiveToolsPing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsPing`: InlineResponse2015
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPing** | [**InlineObject20**](InlineObject20.md) |  | 

### Return type

[**InlineResponse2015**](InlineResponse2015.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPingDevice

> InlineResponse2016 CreateDeviceLiveToolsPingDevice(ctx, serial).CreateDeviceLiveToolsPingDevice(createDeviceLiveToolsPingDevice).Execute()

Enqueue a job to check connectivity status to the device



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
    createDeviceLiveToolsPingDevice := *openapiclient.NewInlineObject21() // InlineObject21 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsPingDevice(context.Background(), serial).CreateDeviceLiveToolsPingDevice(createDeviceLiveToolsPingDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsPingDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsPingDevice`: InlineResponse2016
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPingDevice** | [**InlineObject21**](InlineObject21.md) |  | 

### Return type

[**InlineResponse2016**](InlineResponse2016.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsThroughputTest

> InlineResponse2017 CreateDeviceLiveToolsThroughputTest(ctx, serial).CreateDeviceLiveToolsThroughputTest(createDeviceLiveToolsThroughputTest).Execute()

Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput



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
    createDeviceLiveToolsThroughputTest := *openapiclient.NewInlineObject22() // InlineObject22 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsThroughputTest(context.Background(), serial).CreateDeviceLiveToolsThroughputTest(createDeviceLiveToolsThroughputTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsThroughputTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsThroughputTest`: InlineResponse2017
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsThroughputTest** | [**InlineObject22**](InlineObject22.md) |  | 

### Return type

[**InlineResponse2017**](InlineResponse2017.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsWakeOnLan

> InlineResponse2018 CreateDeviceLiveToolsWakeOnLan(ctx, serial).CreateDeviceLiveToolsWakeOnLan(createDeviceLiveToolsWakeOnLan).Execute()

Enqueue a job to send a Wake-on-LAN packet from the device



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
    createDeviceLiveToolsWakeOnLan := *openapiclient.NewInlineObject23(int32(123), "Mac_example") // InlineObject23 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CreateDeviceLiveToolsWakeOnLan(context.Background(), serial).CreateDeviceLiveToolsWakeOnLan(createDeviceLiveToolsWakeOnLan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CreateDeviceLiveToolsWakeOnLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsWakeOnLan`: InlineResponse2018
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CreateDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsWakeOnLan** | [**InlineObject23**](InlineObject23.md) |  | 

### Return type

[**InlineResponse2018**](InlineResponse2018.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CycleDeviceSwitchPorts

> InlineResponse20036 CycleDeviceSwitchPorts(ctx, serial).CycleDeviceSwitchPorts(cycleDeviceSwitchPorts).Execute()

Cycle a set of switch ports



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
    cycleDeviceSwitchPorts := *openapiclient.NewInlineObject27([]string{"Ports_example"}) // InlineObject27 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.CycleDeviceSwitchPorts(context.Background(), serial).CycleDeviceSwitchPorts(cycleDeviceSwitchPorts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.CycleDeviceSwitchPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CycleDeviceSwitchPorts`: InlineResponse20036
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.CycleDeviceSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCycleDeviceSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cycleDeviceSwitchPorts** | [**InlineObject27**](InlineObject27.md) |  | 

### Return type

[**InlineResponse20036**](InlineResponse20036.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsArpTable

> InlineResponse20022 GetDeviceLiveToolsArpTable(ctx, serial, arpTableId).Execute()

Return an ARP table live tool job.



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
    arpTableId := "arpTableId_example" // string | Arp table ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsArpTable(context.Background(), serial, arpTableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsArpTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsArpTable`: InlineResponse20022
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**arpTableId** | **string** | Arp table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20022**](InlineResponse20022.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsCableTest

> InlineResponse20023 GetDeviceLiveToolsCableTest(ctx, serial, id).Execute()

Return a cable test live tool job.



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsCableTest(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsCableTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsCableTest`: InlineResponse20023
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20023**](InlineResponse20023.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsLedsBlink

> InlineResponse20024 GetDeviceLiveToolsLedsBlink(ctx, serial, ledsBlinkId).Execute()

Return a blink LEDs job



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
    ledsBlinkId := "ledsBlinkId_example" // string | Leds blink ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsLedsBlink(context.Background(), serial, ledsBlinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsLedsBlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsLedsBlink`: InlineResponse20024
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsLedsBlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**ledsBlinkId** | **string** | Leds blink ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsLedsBlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20024**](InlineResponse20024.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsMacTable

> InlineResponse20025 GetDeviceLiveToolsMacTable(ctx, serial, macTableId).Execute()

Return a MAC table live tool job.



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
    macTableId := "macTableId_example" // string | Mac table ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsMacTable(context.Background(), serial, macTableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsMacTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsMacTable`: InlineResponse20025
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**macTableId** | **string** | Mac table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20025**](InlineResponse20025.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPing

> InlineResponse20026 GetDeviceLiveToolsPing(ctx, serial, id).Execute()

Return a ping job



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsPing(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsPing`: InlineResponse20026
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20026**](InlineResponse20026.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPingDevice

> InlineResponse20027 GetDeviceLiveToolsPingDevice(ctx, serial, id).Execute()

Return a ping device job



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsPingDevice(context.Background(), serial, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsPingDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsPingDevice`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsThroughputTest

> InlineResponse20028 GetDeviceLiveToolsThroughputTest(ctx, serial, throughputTestId).Execute()

Return a throughput test job



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
    throughputTestId := "throughputTestId_example" // string | Throughput test ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsThroughputTest(context.Background(), serial, throughputTestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsThroughputTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsThroughputTest`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**throughputTestId** | **string** | Throughput test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsWakeOnLan

> InlineResponse20029 GetDeviceLiveToolsWakeOnLan(ctx, serial, wakeOnLanId).Execute()

Return a Wake-on-LAN job



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
    wakeOnLanId := "wakeOnLanId_example" // string | Wake on lan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LiveToolsApi.GetDeviceLiveToolsWakeOnLan(context.Background(), serial, wakeOnLanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.GetDeviceLiveToolsWakeOnLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsWakeOnLan`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.GetDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**wakeOnLanId** | **string** | Wake on lan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootDevice

> InlineResponse2023 RebootDevice(ctx, serial).Execute()

Reboot a device



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
    resp, r, err := apiClient.LiveToolsApi.RebootDevice(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsApi.RebootDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootDevice`: InlineResponse2023
    fmt.Fprintf(os.Stdout, "Response from `LiveToolsApi.RebootDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2023**](InlineResponse2023.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

