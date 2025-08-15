# \JobsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchNetworkFloorPlansAutoLocateJobs**](JobsApi.md#BatchNetworkFloorPlansAutoLocateJobs) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/batch | Schedule auto locate jobs for one or more floor plans in a network
[**CancelNetworkFloorPlansAutoLocateJob**](JobsApi.md#CancelNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/cancel | Cancel a scheduled or running auto locate job
[**PublishNetworkFloorPlansAutoLocateJob**](JobsApi.md#PublishNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/publish | Update the status of a finished auto locate job to be published, and update device locations
[**RecalculateNetworkFloorPlansAutoLocateJob**](JobsApi.md#RecalculateNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/recalculate | Trigger auto locate recalculation for a job, and optionally set anchors



## BatchNetworkFloorPlansAutoLocateJobs

> InlineResponse200103 BatchNetworkFloorPlansAutoLocateJobs(ctx, networkId).BatchNetworkFloorPlansAutoLocateJobs(batchNetworkFloorPlansAutoLocateJobs).Execute()

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
    batchNetworkFloorPlansAutoLocateJobs := *openapiclient.NewInlineObject106([]openapiclient.NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs{*openapiclient.NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs("FloorPlanId_example")}) // InlineObject106 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.BatchNetworkFloorPlansAutoLocateJobs(context.Background(), networkId).BatchNetworkFloorPlansAutoLocateJobs(batchNetworkFloorPlansAutoLocateJobs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.BatchNetworkFloorPlansAutoLocateJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchNetworkFloorPlansAutoLocateJobs`: InlineResponse200103
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.BatchNetworkFloorPlansAutoLocateJobs`: %v\n", resp)
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

 **batchNetworkFloorPlansAutoLocateJobs** | [**InlineObject106**](InlineObject106.md) |  | 

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
    resp, r, err := apiClient.JobsApi.CancelNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.CancelNetworkFloorPlansAutoLocateJob``: %v\n", err)
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


## PublishNetworkFloorPlansAutoLocateJob

> InlineResponse200104 PublishNetworkFloorPlansAutoLocateJob(ctx, networkId, jobId).PublishNetworkFloorPlansAutoLocateJob(publishNetworkFloorPlansAutoLocateJob).Execute()

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
    publishNetworkFloorPlansAutoLocateJob := *openapiclient.NewInlineObject107() // InlineObject107 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PublishNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).PublishNetworkFloorPlansAutoLocateJob(publishNetworkFloorPlansAutoLocateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PublishNetworkFloorPlansAutoLocateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishNetworkFloorPlansAutoLocateJob`: InlineResponse200104
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PublishNetworkFloorPlansAutoLocateJob`: %v\n", resp)
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


 **publishNetworkFloorPlansAutoLocateJob** | [**InlineObject107**](InlineObject107.md) |  | 

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


## RecalculateNetworkFloorPlansAutoLocateJob

> InlineResponse200105 RecalculateNetworkFloorPlansAutoLocateJob(ctx, networkId, jobId).RecalculateNetworkFloorPlansAutoLocateJob(recalculateNetworkFloorPlansAutoLocateJob).Execute()

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
    recalculateNetworkFloorPlansAutoLocateJob := *openapiclient.NewInlineObject108() // InlineObject108 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.RecalculateNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).RecalculateNetworkFloorPlansAutoLocateJob(recalculateNetworkFloorPlansAutoLocateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.RecalculateNetworkFloorPlansAutoLocateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecalculateNetworkFloorPlansAutoLocateJob`: InlineResponse200105
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.RecalculateNetworkFloorPlansAutoLocateJob`: %v\n", resp)
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


 **recalculateNetworkFloorPlansAutoLocateJob** | [**InlineObject108**](InlineObject108.md) |  | 

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

