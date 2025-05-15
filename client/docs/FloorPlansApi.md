# \FloorPlansApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchNetworkFloorPlansAutoLocateJobs**](FloorPlansApi.md#BatchNetworkFloorPlansAutoLocateJobs) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/batch | Schedule auto locate jobs for one or more floor plans in a network
[**BatchNetworkFloorPlansDevicesUpdate**](FloorPlansApi.md#BatchNetworkFloorPlansDevicesUpdate) | **Post** /networks/{networkId}/floorPlans/devices/batchUpdate | Update floorplan assignments for a batch of devices
[**CancelNetworkFloorPlansAutoLocateJob**](FloorPlansApi.md#CancelNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/cancel | Cancel a scheduled or running auto locate job
[**CreateNetworkFloorPlan**](FloorPlansApi.md#CreateNetworkFloorPlan) | **Post** /networks/{networkId}/floorPlans | Upload a floor plan
[**DeleteNetworkFloorPlan**](FloorPlansApi.md#DeleteNetworkFloorPlan) | **Delete** /networks/{networkId}/floorPlans/{floorPlanId} | Destroy a floor plan
[**GetNetworkFloorPlan**](FloorPlansApi.md#GetNetworkFloorPlan) | **Get** /networks/{networkId}/floorPlans/{floorPlanId} | Find a floor plan by ID
[**GetNetworkFloorPlans**](FloorPlansApi.md#GetNetworkFloorPlans) | **Get** /networks/{networkId}/floorPlans | List the floor plans that belong to your network
[**GetOrganizationFloorPlansAutoLocateDevices**](FloorPlansApi.md#GetOrganizationFloorPlansAutoLocateDevices) | **Get** /organizations/{organizationId}/floorPlans/autoLocate/devices | List auto locate details for each device in your organization
[**GetOrganizationFloorPlansAutoLocateStatuses**](FloorPlansApi.md#GetOrganizationFloorPlansAutoLocateStatuses) | **Get** /organizations/{organizationId}/floorPlans/autoLocate/statuses | List the status of auto locate for each floorplan in your organization
[**PublishNetworkFloorPlansAutoLocateJob**](FloorPlansApi.md#PublishNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/publish | Update the status of a finished auto locate job to be published, and update device locations
[**RecalculateNetworkFloorPlansAutoLocateJob**](FloorPlansApi.md#RecalculateNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/recalculate | Trigger auto locate recalculation for a job, and optionally set anchors
[**UpdateNetworkFloorPlan**](FloorPlansApi.md#UpdateNetworkFloorPlan) | **Put** /networks/{networkId}/floorPlans/{floorPlanId} | Update a floor plan&#39;s geolocation and other meta data



## BatchNetworkFloorPlansAutoLocateJobs

> InlineResponse200102 BatchNetworkFloorPlansAutoLocateJobs(ctx, networkId).BatchNetworkFloorPlansAutoLocateJobs(batchNetworkFloorPlansAutoLocateJobs).Execute()

Schedule auto locate jobs for one or more floor plans in a network



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
    batchNetworkFloorPlansAutoLocateJobs := *openapiclient.NewInlineObject104([]openapiclient.NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs{*openapiclient.NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs("FloorPlanId_example")}) // InlineObject104 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.BatchNetworkFloorPlansAutoLocateJobs(context.Background(), networkId).BatchNetworkFloorPlansAutoLocateJobs(batchNetworkFloorPlansAutoLocateJobs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.BatchNetworkFloorPlansAutoLocateJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchNetworkFloorPlansAutoLocateJobs`: InlineResponse200102
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.BatchNetworkFloorPlansAutoLocateJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchNetworkFloorPlansAutoLocateJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchNetworkFloorPlansAutoLocateJobs** | [**InlineObject104**](InlineObject104.md) |  | 

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


## BatchNetworkFloorPlansDevicesUpdate

> InlineResponse200105 BatchNetworkFloorPlansDevicesUpdate(ctx, networkId).BatchNetworkFloorPlansDevicesUpdate(batchNetworkFloorPlansDevicesUpdate).Execute()

Update floorplan assignments for a batch of devices



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
    batchNetworkFloorPlansDevicesUpdate := *openapiclient.NewInlineObject107([]openapiclient.NetworksNetworkIdFloorPlansDevicesBatchUpdateAssignments{*openapiclient.NewNetworksNetworkIdFloorPlansDevicesBatchUpdateAssignments("Serial_example", *openapiclient.NewNetworksNetworkIdFloorPlansDevicesBatchUpdateFloorPlan("Id_example"))}) // InlineObject107 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.BatchNetworkFloorPlansDevicesUpdate(context.Background(), networkId).BatchNetworkFloorPlansDevicesUpdate(batchNetworkFloorPlansDevicesUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.BatchNetworkFloorPlansDevicesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchNetworkFloorPlansDevicesUpdate`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.BatchNetworkFloorPlansDevicesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchNetworkFloorPlansDevicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchNetworkFloorPlansDevicesUpdate** | [**InlineObject107**](InlineObject107.md) |  | 

### Return type

[**InlineResponse200105**](InlineResponse200105.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelNetworkFloorPlansAutoLocateJob

> CancelNetworkFloorPlansAutoLocateJob(ctx, networkId, jobId).Execute()

Cancel a scheduled or running auto locate job



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
    jobId := "jobId_example" // string | Job ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.CancelNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.CancelNetworkFloorPlansAutoLocateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelNetworkFloorPlansAutoLocateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFloorPlan

> InlineResponse200101 CreateNetworkFloorPlan(ctx, networkId).CreateNetworkFloorPlan(createNetworkFloorPlan).Execute()

Upload a floor plan



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
    createNetworkFloorPlan := *openapiclient.NewInlineObject103("Name_example", "ImageContents_example") // InlineObject103 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.CreateNetworkFloorPlan(context.Background(), networkId).CreateNetworkFloorPlan(createNetworkFloorPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.CreateNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFloorPlan`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.CreateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFloorPlan** | [**InlineObject103**](InlineObject103.md) |  | 

### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkFloorPlan

> InlineResponse200101 DeleteNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Destroy a floor plan



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
    floorPlanId := "floorPlanId_example" // string | Floor plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.DeleteNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.DeleteNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkFloorPlan`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.DeleteNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFloorPlan

> InlineResponse200101 GetNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Find a floor plan by ID



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
    floorPlanId := "floorPlanId_example" // string | Floor plan ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.GetNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.GetNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFloorPlan`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.GetNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFloorPlans

> []InlineResponse200101 GetNetworkFloorPlans(ctx, networkId).Execute()

List the floor plans that belong to your network



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
    resp, r, err := apiClient.FloorPlansApi.GetNetworkFloorPlans(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.GetNetworkFloorPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFloorPlans`: []InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.GetNetworkFloorPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationFloorPlansAutoLocateDevices

> []InlineResponse200283 GetOrganizationFloorPlansAutoLocateDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).FloorPlanIds(floorPlanIds).Execute()

List auto locate details for each device in your organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more network IDs (optional)
    floorPlanIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more floorplan IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.GetOrganizationFloorPlansAutoLocateDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).FloorPlanIds(floorPlanIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.GetOrganizationFloorPlansAutoLocateDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationFloorPlansAutoLocateDevices`: []InlineResponse200283
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.GetOrganizationFloorPlansAutoLocateDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationFloorPlansAutoLocateDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by one or more network IDs | 
 **floorPlanIds** | **[]string** | Optional parameter to filter devices by one or more floorplan IDs | 

### Return type

[**[]InlineResponse200283**](InlineResponse200283.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationFloorPlansAutoLocateStatuses

> []InlineResponse200284 GetOrganizationFloorPlansAutoLocateStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).FloorPlanIds(floorPlanIds).Execute()

List the status of auto locate for each floorplan in your organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter floorplans by one or more network IDs (optional)
    floorPlanIds := []string{"Inner_example"} // []string | Optional parameter to filter floorplans by one or more floorplan IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.GetOrganizationFloorPlansAutoLocateStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).FloorPlanIds(floorPlanIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.GetOrganizationFloorPlansAutoLocateStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationFloorPlansAutoLocateStatuses`: []InlineResponse200284
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.GetOrganizationFloorPlansAutoLocateStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationFloorPlansAutoLocateStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 10000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter floorplans by one or more network IDs | 
 **floorPlanIds** | **[]string** | Optional parameter to filter floorplans by one or more floorplan IDs | 

### Return type

[**[]InlineResponse200284**](InlineResponse200284.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishNetworkFloorPlansAutoLocateJob

> InlineResponse200103 PublishNetworkFloorPlansAutoLocateJob(ctx, networkId, jobId).PublishNetworkFloorPlansAutoLocateJob(publishNetworkFloorPlansAutoLocateJob).Execute()

Update the status of a finished auto locate job to be published, and update device locations



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
    jobId := "jobId_example" // string | Job ID
    publishNetworkFloorPlansAutoLocateJob := *openapiclient.NewInlineObject105() // InlineObject105 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.PublishNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).PublishNetworkFloorPlansAutoLocateJob(publishNetworkFloorPlansAutoLocateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.PublishNetworkFloorPlansAutoLocateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishNetworkFloorPlansAutoLocateJob`: InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.PublishNetworkFloorPlansAutoLocateJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishNetworkFloorPlansAutoLocateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publishNetworkFloorPlansAutoLocateJob** | [**InlineObject105**](InlineObject105.md) |  | 

### Return type

[**InlineResponse200103**](InlineResponse200103.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecalculateNetworkFloorPlansAutoLocateJob

> InlineResponse200104 RecalculateNetworkFloorPlansAutoLocateJob(ctx, networkId, jobId).RecalculateNetworkFloorPlansAutoLocateJob(recalculateNetworkFloorPlansAutoLocateJob).Execute()

Trigger auto locate recalculation for a job, and optionally set anchors



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
    jobId := "jobId_example" // string | Job ID
    recalculateNetworkFloorPlansAutoLocateJob := *openapiclient.NewInlineObject106() // InlineObject106 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.RecalculateNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).RecalculateNetworkFloorPlansAutoLocateJob(recalculateNetworkFloorPlansAutoLocateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.RecalculateNetworkFloorPlansAutoLocateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecalculateNetworkFloorPlansAutoLocateJob`: InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.RecalculateNetworkFloorPlansAutoLocateJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecalculateNetworkFloorPlansAutoLocateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recalculateNetworkFloorPlansAutoLocateJob** | [**InlineObject106**](InlineObject106.md) |  | 

### Return type

[**InlineResponse200104**](InlineResponse200104.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFloorPlan

> InlineResponse200101 UpdateNetworkFloorPlan(ctx, networkId, floorPlanId).UpdateNetworkFloorPlan(updateNetworkFloorPlan).Execute()

Update a floor plan's geolocation and other meta data



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
    floorPlanId := "floorPlanId_example" // string | Floor plan ID
    updateNetworkFloorPlan := *openapiclient.NewInlineObject108() // InlineObject108 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.UpdateNetworkFloorPlan(context.Background(), networkId, floorPlanId).UpdateNetworkFloorPlan(updateNetworkFloorPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.UpdateNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFloorPlan`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.UpdateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkFloorPlan** | [**InlineObject108**](InlineObject108.md) |  | 

### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

