# \JobsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchNetworkFloorPlansAutoLocateJobs**](JobsApi.md#BatchNetworkFloorPlansAutoLocateJobs) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/batch | Schedule auto locate jobs for one or more floor plans in a network
[**CancelNetworkFloorPlansAutoLocateJob**](JobsApi.md#CancelNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/cancel | Cancel a scheduled or running auto locate job
[**GetOrganizationApiRestProvisioningPipelinesJobs**](JobsApi.md#GetOrganizationApiRestProvisioningPipelinesJobs) | **Get** /organizations/{organizationId}/api/rest/provisioning/pipelines/jobs | List pipeline jobs, with optional status filtering
[**GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline**](JobsApi.md#GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline) | **Get** /organizations/{organizationId}/api/rest/provisioning/pipelines/jobs/overviews/byPipeline | Retrieves pipeline overviews with aggregated job status counts
[**PublishNetworkFloorPlansAutoLocateJob**](JobsApi.md#PublishNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/publish | Update the status of a finished auto locate job to be published, and update device locations
[**RecalculateNetworkFloorPlansAutoLocateJob**](JobsApi.md#RecalculateNetworkFloorPlansAutoLocateJob) | **Post** /networks/{networkId}/floorPlans/autoLocate/jobs/{jobId}/recalculate | Trigger auto locate recalculation for a job, and optionally set anchors



## BatchNetworkFloorPlansAutoLocateJobs

> InlineResponse200140 BatchNetworkFloorPlansAutoLocateJobs(ctx, networkId).BatchNetworkFloorPlansAutoLocateJobs(batchNetworkFloorPlansAutoLocateJobs).Execute()

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
    batchNetworkFloorPlansAutoLocateJobs := *openapiclient.NewInlineObject125([]openapiclient.NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs{*openapiclient.NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs("FloorPlanId_example")}) // InlineObject125 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.BatchNetworkFloorPlansAutoLocateJobs(context.Background(), networkId).BatchNetworkFloorPlansAutoLocateJobs(batchNetworkFloorPlansAutoLocateJobs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.BatchNetworkFloorPlansAutoLocateJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchNetworkFloorPlansAutoLocateJobs`: InlineResponse200140
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

 **batchNetworkFloorPlansAutoLocateJobs** | [**InlineObject125**](InlineObject125.md) |  | 

### Return type

[**InlineResponse200140**](InlineResponse200140.md)

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


## GetOrganizationApiRestProvisioningPipelinesJobs

> InlineResponse200266 GetOrganizationApiRestProvisioningPipelinesJobs(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PipelineIds(pipelineIds).Status(status).Execute()

List pipeline jobs, with optional status filtering



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    pipelineIds := []string{"Inner_example"} // []string | Pipeline IDs to retrieve jobs for (optional)
    status := "status_example" // string | If provided, filters jobs by status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetOrganizationApiRestProvisioningPipelinesJobs(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PipelineIds(pipelineIds).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetOrganizationApiRestProvisioningPipelinesJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRestProvisioningPipelinesJobs`: InlineResponse200266
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetOrganizationApiRestProvisioningPipelinesJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRestProvisioningPipelinesJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **pipelineIds** | **[]string** | Pipeline IDs to retrieve jobs for | 
 **status** | **string** | If provided, filters jobs by status | 

### Return type

[**InlineResponse200266**](InlineResponse200266.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline

> InlineResponse200267 GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(ctx, organizationId).PipelineIds(pipelineIds).Execute()

Retrieves pipeline overviews with aggregated job status counts



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
    pipelineIds := []string{"Inner_example"} // []string | Pipeline IDs to retrieve overviews for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(context.Background(), organizationId).PipelineIds(pipelineIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: InlineResponse200267
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipelineIds** | **[]string** | Pipeline IDs to retrieve overviews for | 

### Return type

[**InlineResponse200267**](InlineResponse200267.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishNetworkFloorPlansAutoLocateJob

> InlineResponse200141 PublishNetworkFloorPlansAutoLocateJob(ctx, networkId, jobId).PublishNetworkFloorPlansAutoLocateJob(publishNetworkFloorPlansAutoLocateJob).Execute()

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
    publishNetworkFloorPlansAutoLocateJob := *openapiclient.NewInlineObject126() // InlineObject126 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PublishNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).PublishNetworkFloorPlansAutoLocateJob(publishNetworkFloorPlansAutoLocateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PublishNetworkFloorPlansAutoLocateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishNetworkFloorPlansAutoLocateJob`: InlineResponse200141
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


 **publishNetworkFloorPlansAutoLocateJob** | [**InlineObject126**](InlineObject126.md) |  | 

### Return type

[**InlineResponse200141**](InlineResponse200141.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecalculateNetworkFloorPlansAutoLocateJob

> InlineResponse200142 RecalculateNetworkFloorPlansAutoLocateJob(ctx, networkId, jobId).RecalculateNetworkFloorPlansAutoLocateJob(recalculateNetworkFloorPlansAutoLocateJob).Execute()

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
    recalculateNetworkFloorPlansAutoLocateJob := *openapiclient.NewInlineObject127() // InlineObject127 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.RecalculateNetworkFloorPlansAutoLocateJob(context.Background(), networkId, jobId).RecalculateNetworkFloorPlansAutoLocateJob(recalculateNetworkFloorPlansAutoLocateJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.RecalculateNetworkFloorPlansAutoLocateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecalculateNetworkFloorPlansAutoLocateJob`: InlineResponse200142
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


 **recalculateNetworkFloorPlansAutoLocateJob** | [**InlineObject127**](InlineObject127.md) |  | 

### Return type

[**InlineResponse200142**](InlineResponse200142.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

