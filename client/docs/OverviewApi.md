# \OverviewApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraAnalyticsOverview**](OverviewApi.md#GetDeviceCameraAnalyticsOverview) | **Get** /devices/{serial}/camera/analytics/overview | Returns an overview of aggregate analytics data for a timespan
[**GetNetworkClientsOverview**](OverviewApi.md#GetNetworkClientsOverview) | **Get** /networks/{networkId}/clients/overview | Return overview statistics for network clients
[**GetNetworkSensorAlertsCurrentOverviewByMetric**](OverviewApi.md#GetNetworkSensorAlertsCurrentOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/current/overview/byMetric | Return an overview of currently alerting sensors by metric
[**GetNetworkSensorAlertsOverviewByMetric**](OverviewApi.md#GetNetworkSensorAlertsOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/overview/byMetric | Return an overview of alert occurrences over a timespan, by metric
[**GetOrganizationAdaptivePolicyOverview**](OverviewApi.md#GetOrganizationAdaptivePolicyOverview) | **Get** /organizations/{organizationId}/adaptivePolicy/overview | Returns adaptive policy aggregate statistics for an organization
[**GetOrganizationApiRequestsOverview**](OverviewApi.md#GetOrganizationApiRequestsOverview) | **Get** /organizations/{organizationId}/apiRequests/overview | Return an aggregated overview of API requests data
[**GetOrganizationApiRequestsOverviewResponseCodesByInterval**](OverviewApi.md#GetOrganizationApiRequestsOverviewResponseCodesByInterval) | **Get** /organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval | Tracks organizations&#39; API requests by response code across a given time period
[**GetOrganizationApplianceUplinksStatusesOverview**](OverviewApi.md#GetOrganizationApplianceUplinksStatusesOverview) | **Get** /organizations/{organizationId}/appliance/uplinks/statuses/overview | Returns an overview of uplink statuses
[**GetOrganizationAssuranceAlertsOverview**](OverviewApi.md#GetOrganizationAssuranceAlertsOverview) | **Get** /organizations/{organizationId}/assurance/alerts/overview | Return overview of active health alerts for an organization
[**GetOrganizationAssuranceAlertsOverviewByNetwork**](OverviewApi.md#GetOrganizationAssuranceAlertsOverviewByNetwork) | **Get** /organizations/{organizationId}/assurance/alerts/overview/byNetwork | Return a Summary of Alerts grouped by network and severity
[**GetOrganizationAssuranceAlertsOverviewByType**](OverviewApi.md#GetOrganizationAssuranceAlertsOverviewByType) | **Get** /organizations/{organizationId}/assurance/alerts/overview/byType | Return a Summary of Alerts grouped by type and severity
[**GetOrganizationAssuranceAlertsOverviewHistorical**](OverviewApi.md#GetOrganizationAssuranceAlertsOverviewHistorical) | **Get** /organizations/{organizationId}/assurance/alerts/overview/historical | Returns historical health alert overviews
[**GetOrganizationClientsOverview**](OverviewApi.md#GetOrganizationClientsOverview) | **Get** /organizations/{organizationId}/clients/overview | Return summary information around client data usage (in kb) across the given organization.
[**GetOrganizationDevicesOverviewByModel**](OverviewApi.md#GetOrganizationDevicesOverviewByModel) | **Get** /organizations/{organizationId}/devices/overview/byModel | Lists the count for each device model
[**GetOrganizationDevicesStatusesOverview**](OverviewApi.md#GetOrganizationDevicesStatusesOverview) | **Get** /organizations/{organizationId}/devices/statuses/overview | Return an overview of current device statuses
[**GetOrganizationLicensesOverview**](OverviewApi.md#GetOrganizationLicensesOverview) | **Get** /organizations/{organizationId}/licenses/overview | Return an overview of the license state for an organization
[**GetOrganizationSwitchPortsClientsOverviewByDevice**](OverviewApi.md#GetOrganizationSwitchPortsClientsOverviewByDevice) | **Get** /organizations/{organizationId}/switch/ports/clients/overview/byDevice | List the number of clients for all switchports with at least one online client in an organization.
[**GetOrganizationSwitchPortsOverview**](OverviewApi.md#GetOrganizationSwitchPortsOverview) | **Get** /organizations/{organizationId}/switch/ports/overview | Returns the counts of all active ports for the requested timespan, grouped by speed
[**GetOrganizationWirelessClientsOverviewByDevice**](OverviewApi.md#GetOrganizationWirelessClientsOverviewByDevice) | **Get** /organizations/{organizationId}/wireless/clients/overview/byDevice | List access point client count at the moment in an organization
[**GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval**](OverviewApi.md#GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/wirelessController/clients/overview/history/byDevice/byInterval | List wireless client counts of wireless LAN controllers over time in an organization
[**GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice**](OverviewApi.md#GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice) | **Get** /organizations/{organizationId}/wirelessController/devices/interfaces/packets/overview/byDevice | Retrieve the packet counters for the interfaces of a Wireless LAN controller
[**GetOrganizationWirelessControllerOverviewByDevice**](OverviewApi.md#GetOrganizationWirelessControllerOverviewByDevice) | **Get** /organizations/{organizationId}/wirelessController/overview/byDevice | List the overview information of wireless LAN controllers in an organization and it is updated every minute.



## GetDeviceCameraAnalyticsOverview

> []InlineResponse20012 GetDeviceCameraAnalyticsOverview(ctx, serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()

Returns an overview of aggregate analytics data for a timespan



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. (optional)
    objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetDeviceCameraAnalyticsOverview(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetDeviceCameraAnalyticsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraAnalyticsOverview`: []InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetDeviceCameraAnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. | 
 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]InlineResponse20012**](InlineResponse20012.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientsOverview

> InlineResponse20090 GetNetworkClientsOverview(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Return overview statistics for network clients



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetNetworkClientsOverview(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetNetworkClientsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientsOverview`: InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetNetworkClientsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800. | 

### Return type

[**InlineResponse20090**](InlineResponse20090.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsCurrentOverviewByMetric

> InlineResponse200119 GetNetworkSensorAlertsCurrentOverviewByMetric(ctx, networkId).Execute()

Return an overview of currently alerting sensors by metric



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetNetworkSensorAlertsCurrentOverviewByMetric(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetNetworkSensorAlertsCurrentOverviewByMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsCurrentOverviewByMetric`: InlineResponse200119
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetNetworkSensorAlertsCurrentOverviewByMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsCurrentOverviewByMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200119**](InlineResponse200119.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsOverviewByMetric

> []InlineResponse200120 GetNetworkSensorAlertsOverviewByMetric(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Return an overview of alert occurrences over a timespan, by metric



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 731 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 366 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 366 days. The default is 7 days. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 900, 3600, 86400, 604800, 2592000. The default is 604800. Interval is calculated if time params are provided. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetNetworkSensorAlertsOverviewByMetric(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetNetworkSensorAlertsOverviewByMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsOverviewByMetric`: []InlineResponse200120
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetNetworkSensorAlertsOverviewByMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsOverviewByMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 731 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 366 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 366 days. The default is 7 days. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 900, 3600, 86400, 604800, 2592000. The default is 604800. Interval is calculated if time params are provided. | 

### Return type

[**[]InlineResponse200120**](InlineResponse200120.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyOverview

> InlineResponse200218 GetOrganizationAdaptivePolicyOverview(ctx, organizationId).Execute()

Returns adaptive policy aggregate statistics for an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationAdaptivePolicyOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationAdaptivePolicyOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyOverview`: InlineResponse200218
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationAdaptivePolicyOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200218**](InlineResponse200218.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRequestsOverview

> InlineResponse200224 GetOrganizationApiRequestsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return an aggregated overview of API requests data



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationApiRequestsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationApiRequestsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRequestsOverview`: InlineResponse200224
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationApiRequestsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRequestsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. | 

### Return type

[**InlineResponse200224**](InlineResponse200224.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRequestsOverviewResponseCodesByInterval

> []InlineResponse200225 GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()

Tracks organizations' API requests by response code across a given time period



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. (optional)
    version := int32(56) // int32 | Filter by API version of the endpoint. Allowable values are: [0, 1] (optional)
    operationIds := []string{"Inner_example"} // []string | Filter by operation ID of the endpoint (optional)
    sourceIps := []string{"Inner_example"} // []string | Filter by source IP that made the API request (optional)
    adminIds := []string{"Inner_example"} // []string | Filter by admin ID of user that made the API request (optional)
    userAgent := "userAgent_example" // string | Filter by user agent string for API request. This will filter by a complete or partial match. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRequestsOverviewResponseCodesByInterval`: []InlineResponse200225
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationApiRequestsOverviewResponseCodesByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. | 
 **version** | **int32** | Filter by API version of the endpoint. Allowable values are: [0, 1] | 
 **operationIds** | **[]string** | Filter by operation ID of the endpoint | 
 **sourceIps** | **[]string** | Filter by source IP that made the API request | 
 **adminIds** | **[]string** | Filter by admin ID of user that made the API request | 
 **userAgent** | **string** | Filter by user agent string for API request. This will filter by a complete or partial match. | 

### Return type

[**[]InlineResponse200225**](InlineResponse200225.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinksStatusesOverview

> InlineResponse200236 GetOrganizationApplianceUplinksStatusesOverview(ctx, organizationId).Execute()

Returns an overview of uplink statuses



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationApplianceUplinksStatusesOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationApplianceUplinksStatusesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceUplinksStatusesOverview`: InlineResponse200236
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationApplianceUplinksStatusesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksStatusesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200236**](InlineResponse200236.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverview

> InlineResponse200244 GetOrganizationAssuranceAlertsOverview(ctx, organizationId).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return overview of active health alerts for an organization



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
    organizationId := "organizationId_example" // string | Organization ID
    networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network ids. (optional)
    severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
    types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
    tsStart := time.Now() // time.Time | Optional parameter to filter by starting timestamp (optional)
    tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp (optional)
    category := "category_example" // string | Optional parameter to filter by category. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
    deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)
    deviceTags := []string{"Inner_example"} // []string | Optional parameter to filter by device tags (optional)
    active := true // bool | Optional parameter to filter by active alerts defaults to true (optional)
    dismissed := true // bool | Optional parameter to filter by dismissed alerts defaults to false (optional)
    resolved := true // bool | Optional parameter to filter by resolved alerts defaults to false (optional)
    suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationAssuranceAlertsOverview(context.Background(), organizationId).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationAssuranceAlertsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverview`: InlineResponse200244
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationAssuranceAlertsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkId** | **string** | Optional parameter to filter alerts overview by network ids. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsStart** | **time.Time** | Optional parameter to filter by starting timestamp | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp | 
 **category** | **string** | Optional parameter to filter by category. | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 
 **deviceTags** | **[]string** | Optional parameter to filter by device tags | 
 **active** | **bool** | Optional parameter to filter by active alerts defaults to true | 
 **dismissed** | **bool** | Optional parameter to filter by dismissed alerts defaults to false | 
 **resolved** | **bool** | Optional parameter to filter by resolved alerts defaults to false | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**InlineResponse200244**](InlineResponse200244.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewByNetwork

> InlineResponse200245 GetOrganizationAssuranceAlertsOverviewByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return a Summary of Alerts grouped by network and severity



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
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'. (optional)
    networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network id. (optional)
    severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
    types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
    tsStart := time.Now() // time.Time | Optional parameter to filter by starting timestamp (optional)
    tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp (optional)
    category := "category_example" // string | Optional parameter to filter by category. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
    deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)
    deviceTags := []string{"Inner_example"} // []string | Optional parameter to filter by device tags (optional)
    active := true // bool | Optional parameter to filter by active alerts defaults to true (optional)
    dismissed := true // bool | Optional parameter to filter by dismissed alerts defaults to false (optional)
    resolved := true // bool | Optional parameter to filter by resolved alerts defaults to false (optional)
    suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationAssuranceAlertsOverviewByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationAssuranceAlertsOverviewByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverviewByNetwork`: InlineResponse200245
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationAssuranceAlertsOverviewByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;ascending&#39;. | 
 **networkId** | **string** | Optional parameter to filter alerts overview by network id. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsStart** | **time.Time** | Optional parameter to filter by starting timestamp | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp | 
 **category** | **string** | Optional parameter to filter by category. | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 
 **deviceTags** | **[]string** | Optional parameter to filter by device tags | 
 **active** | **bool** | Optional parameter to filter by active alerts defaults to true | 
 **dismissed** | **bool** | Optional parameter to filter by dismissed alerts defaults to false | 
 **resolved** | **bool** | Optional parameter to filter by resolved alerts defaults to false | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewByType

> InlineResponse200246 GetOrganizationAssuranceAlertsOverviewByType(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return a Summary of Alerts grouped by type and severity



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
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'. (optional)
    networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network ids. (optional)
    severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
    types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
    tsStart := time.Now() // time.Time | Optional parameter to filter by starting timestamp (optional)
    tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp (optional)
    category := "category_example" // string | Optional parameter to filter by category. (optional)
    sortBy := "sortBy_example" // string | Optional parameter to set column to sort by. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
    deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)
    deviceTags := []string{"Inner_example"} // []string | Optional parameter to filter by device tags (optional)
    active := true // bool | Optional parameter to filter by active alerts defaults to true (optional)
    dismissed := true // bool | Optional parameter to filter by dismissed alerts defaults to false (optional)
    resolved := true // bool | Optional parameter to filter by resolved alerts defaults to false (optional)
    suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationAssuranceAlertsOverviewByType(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationAssuranceAlertsOverviewByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverviewByType`: InlineResponse200246
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationAssuranceAlertsOverviewByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;ascending&#39;. | 
 **networkId** | **string** | Optional parameter to filter alerts overview by network ids. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsStart** | **time.Time** | Optional parameter to filter by starting timestamp | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp | 
 **category** | **string** | Optional parameter to filter by category. | 
 **sortBy** | **string** | Optional parameter to set column to sort by. | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 
 **deviceTags** | **[]string** | Optional parameter to filter by device tags | 
 **active** | **bool** | Optional parameter to filter by active alerts defaults to true | 
 **dismissed** | **bool** | Optional parameter to filter by dismissed alerts defaults to false | 
 **resolved** | **bool** | Optional parameter to filter by resolved alerts defaults to false | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**InlineResponse200246**](InlineResponse200246.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewHistorical

> InlineResponse200247 GetOrganizationAssuranceAlertsOverviewHistorical(ctx, organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).Execute()

Returns historical health alert overviews



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
    organizationId := "organizationId_example" // string | Organization ID
    segmentDuration := int32(56) // int32 | Amount of time in seconds for each segment in the returned dataset
    tsStart := time.Now() // time.Time | Parameter to define starting timestamp of historical totals
    networkId := "networkId_example" // string | Optional parameter to filter alerts overview by network ids. (optional)
    severity := "severity_example" // string | Optional parameter to filter alerts overview by severity type. (optional)
    types := []string{"Types_example"} // []string | Optional parameter to filter by alert type. (optional)
    tsEnd := time.Now() // time.Time | Optional parameter to filter by end timestamp defaults to the current time (optional)
    category := "category_example" // string | Optional parameter to filter by category. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter by primary device serial (optional)
    deviceTypes := []string{"DeviceTypes_example"} // []string | Optional parameter to filter by device types (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationAssuranceAlertsOverviewHistorical(context.Background(), organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationAssuranceAlertsOverviewHistorical``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverviewHistorical`: InlineResponse200247
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationAssuranceAlertsOverviewHistorical`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsOverviewHistoricalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segmentDuration** | **int32** | Amount of time in seconds for each segment in the returned dataset | 
 **tsStart** | **time.Time** | Parameter to define starting timestamp of historical totals | 
 **networkId** | **string** | Optional parameter to filter alerts overview by network ids. | 
 **severity** | **string** | Optional parameter to filter alerts overview by severity type. | 
 **types** | **[]string** | Optional parameter to filter by alert type. | 
 **tsEnd** | **time.Time** | Optional parameter to filter by end timestamp defaults to the current time | 
 **category** | **string** | Optional parameter to filter by category. | 
 **serials** | **[]string** | Optional parameter to filter by primary device serial | 
 **deviceTypes** | **[]string** | Optional parameter to filter by device types | 

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


## GetOrganizationClientsOverview

> InlineResponse200268 GetOrganizationClientsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return summary information around client data usage (in kb) across the given organization.



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationClientsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationClientsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationClientsOverview`: InlineResponse200268
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationClientsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClientsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**InlineResponse200268**](InlineResponse200268.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesOverviewByModel

> InlineResponse200278 GetOrganizationDevicesOverviewByModel(ctx, organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()

Lists the count for each device model



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
    models := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by networkId. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationDevicesOverviewByModel(context.Background(), organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationDevicesOverviewByModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesOverviewByModel`: InlineResponse200278
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationDevicesOverviewByModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesOverviewByModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **models** | **[]string** | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by networkId. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 

### Return type

[**InlineResponse200278**](InlineResponse200278.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesStatusesOverview

> InlineResponse200286 GetOrganizationDevicesStatusesOverview(ctx, organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()

Return an overview of current device statuses



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
    productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. (optional)
    networkIds := []string{"Inner_example"} // []string | An optional parameter to filter device statuses by network. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationDevicesStatusesOverview(context.Background(), organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationDevicesStatusesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesStatusesOverview`: InlineResponse200286
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationDevicesStatusesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesStatusesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productTypes** | **[]string** | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, campusGateway, and secureConnect. | 
 **networkIds** | **[]string** | An optional parameter to filter device statuses by network. | 

### Return type

[**InlineResponse200286**](InlineResponse200286.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLicensesOverview

> InlineResponse200308 GetOrganizationLicensesOverview(ctx, organizationId).Execute()

Return an overview of the license state for an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationLicensesOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationLicensesOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicensesOverview`: InlineResponse200308
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationLicensesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200308**](InlineResponse200308.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSwitchPortsClientsOverviewByDevice

> InlineResponse200345 GetOrganizationSwitchPortsClientsOverviewByDevice(ctx, organizationId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()

List the number of clients for all switchports with at least one online client in an organization.



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
    organizationId := "organizationId_example" // string | Organization ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 20. Default is 20. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    configurationUpdatedAfter := time.Now() // time.Time | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. (optional)
    mac := "mac_example" // string | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided MAC addresses. (optional)
    name := "name_example" // string | Optional parameter to filter items to switches with names that contain the search term or are an exact match. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches in one of the provided networks. (optional)
    portProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. (optional)
    serial := "serial_example" // string | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided serials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationSwitchPortsClientsOverviewByDevice(context.Background(), organizationId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationSwitchPortsClientsOverviewByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSwitchPortsClientsOverviewByDevice`: InlineResponse200345
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationSwitchPortsClientsOverviewByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsClientsOverviewByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 20. Default is 20. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **time.Time** | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. | 
 **mac** | **string** | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. | 
 **macs** | **[]string** | Optional parameter to filter items to switches that have one of the provided MAC addresses. | 
 **name** | **string** | Optional parameter to filter items to switches with names that contain the search term or are an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter items to switches in one of the provided networks. | 
 **portProfileIds** | **[]string** | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. | 
 **serial** | **string** | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. | 
 **serials** | **[]string** | Optional parameter to filter items to switches that have one of the provided serials. | 

### Return type

[**InlineResponse200345**](InlineResponse200345.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSwitchPortsOverview

> InlineResponse200346 GetOrganizationSwitchPortsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns the counts of all active ports for the requested timespan, grouped by speed



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationSwitchPortsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationSwitchPortsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSwitchPortsOverview`: InlineResponse200346
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationSwitchPortsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**InlineResponse200346**](InlineResponse200346.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessClientsOverviewByDevice

> InlineResponse200356 GetOrganizationWirelessClientsOverviewByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewayClusterIds(campusGatewayClusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List access point client count at the moment in an organization



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter access points client counts by network ID. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter access points client counts by its serial numbers. This filter uses multiple exact matches. (optional)
    campusGatewayClusterIds := []string{"Inner_example"} // []string | Optional parameter to filter access points client counts by MCG cluster IDs. This filter uses multiple exact matches. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationWirelessClientsOverviewByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).CampusGatewayClusterIds(campusGatewayClusterIds).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationWirelessClientsOverviewByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessClientsOverviewByDevice`: InlineResponse200356
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationWirelessClientsOverviewByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessClientsOverviewByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter access points client counts by network ID. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter access points client counts by its serial numbers. This filter uses multiple exact matches. | 
 **campusGatewayClusterIds** | **[]string** | Optional parameter to filter access points client counts by MCG cluster IDs. This filter uses multiple exact matches. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200356**](InlineResponse200356.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval

> InlineResponse200382 GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Resolution(resolution).Execute()

List wireless client counts of wireless LAN controllers over time in an organization



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval`: InlineResponse200382
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerClientsOverviewHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. | 

### Return type

[**InlineResponse200382**](InlineResponse200382.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice

> InlineResponse200390 GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(ctx, organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Retrieve the packet counters for the interfaces of a Wireless LAN controller



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
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    names := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 hour. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice(context.Background(), organizationId).Serials(serials).Names(names).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice`: InlineResponse200390
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerDevicesInterfacesPacketsOverviewByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **names** | **[]string** | Optional parameter to filter wireless LAN controller by its interface name. This filter uses multiple exact matches. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 1 day after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 1 day. The default is 1 hour. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200390**](InlineResponse200390.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessControllerOverviewByDevice

> InlineResponse200395 GetOrganizationWirelessControllerOverviewByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the overview information of wireless LAN controllers in an organization and it is updated every minute.



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
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverviewApi.GetOrganizationWirelessControllerOverviewByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverviewApi.GetOrganizationWirelessControllerOverviewByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessControllerOverviewByDevice`: InlineResponse200395
    fmt.Fprintf(os.Stdout, "Response from `OverviewApi.GetOrganizationWirelessControllerOverviewByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessControllerOverviewByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Optional parameter to filter wireless LAN controllers by network ID. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter wireless LAN controller by its cloud ID. This filter uses multiple exact matches. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**InlineResponse200395**](InlineResponse200395.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

