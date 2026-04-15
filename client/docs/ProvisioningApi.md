# \ProvisioningApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationWirelessDevicesProvisioningDeployment**](ProvisioningApi.md#CreateOrganizationWirelessDevicesProvisioningDeployment) | **Post** /organizations/{organizationId}/wireless/devices/provisioning/deployments | Create a zero touch deployment for a wireless access point
[**DeleteOrganizationWirelessDevicesProvisioningDeployment**](ProvisioningApi.md#DeleteOrganizationWirelessDevicesProvisioningDeployment) | **Delete** /organizations/{organizationId}/wireless/devices/provisioning/deployments/{deploymentId} | Delete a zero touch deployment
[**GetOrganizationApiRestProvisioningPipelinesJobs**](ProvisioningApi.md#GetOrganizationApiRestProvisioningPipelinesJobs) | **Get** /organizations/{organizationId}/api/rest/provisioning/pipelines/jobs | List pipeline jobs, with optional status filtering
[**GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline**](ProvisioningApi.md#GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline) | **Get** /organizations/{organizationId}/api/rest/provisioning/pipelines/jobs/overviews/byPipeline | Retrieves pipeline overviews with aggregated job status counts
[**GetOrganizationDevicesProvisioningStatuses**](ProvisioningApi.md#GetOrganizationDevicesProvisioningStatuses) | **Get** /organizations/{organizationId}/devices/provisioning/statuses | List the provisioning statuses information for devices in an organization.
[**GetOrganizationWirelessDevicesProvisioningDeployments**](ProvisioningApi.md#GetOrganizationWirelessDevicesProvisioningDeployments) | **Get** /organizations/{organizationId}/wireless/devices/provisioning/deployments | List the zero touch deployments for wireless access points in an organization
[**UpdateOrganizationWirelessDevicesProvisioningDeployments**](ProvisioningApi.md#UpdateOrganizationWirelessDevicesProvisioningDeployments) | **Put** /organizations/{organizationId}/wireless/devices/provisioning/deployments | Update a zero touch deployment



## CreateOrganizationWirelessDevicesProvisioningDeployment

> InlineResponse200416 CreateOrganizationWirelessDevicesProvisioningDeployment(ctx, organizationId).CreateOrganizationWirelessDevicesProvisioningDeployment(createOrganizationWirelessDevicesProvisioningDeployment).Execute()

Create a zero touch deployment for a wireless access point



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
    createOrganizationWirelessDevicesProvisioningDeployment := *openapiclient.NewInlineObject328([]openapiclient.OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1{*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1New()), "Status_example", "Type_example")}) // InlineObject328 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.CreateOrganizationWirelessDevicesProvisioningDeployment(context.Background(), organizationId).CreateOrganizationWirelessDevicesProvisioningDeployment(createOrganizationWirelessDevicesProvisioningDeployment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.CreateOrganizationWirelessDevicesProvisioningDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessDevicesProvisioningDeployment`: InlineResponse200416
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.CreateOrganizationWirelessDevicesProvisioningDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationWirelessDevicesProvisioningDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationWirelessDevicesProvisioningDeployment** | [**InlineObject328**](InlineObject328.md) |  | 

### Return type

[**InlineResponse200416**](InlineResponse200416.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationWirelessDevicesProvisioningDeployment

> DeleteOrganizationWirelessDevicesProvisioningDeployment(ctx, organizationId, deploymentId).Execute()

Delete a zero touch deployment



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
    deploymentId := "deploymentId_example" // string | Deployment ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.DeleteOrganizationWirelessDevicesProvisioningDeployment(context.Background(), organizationId, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.DeleteOrganizationWirelessDevicesProvisioningDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**deploymentId** | **string** | Deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationWirelessDevicesProvisioningDeploymentRequest struct via the builder pattern


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

> InlineResponse200244 GetOrganizationApiRestProvisioningPipelinesJobs(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PipelineIds(pipelineIds).Status(status).Execute()

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
    resp, r, err := apiClient.ProvisioningApi.GetOrganizationApiRestProvisioningPipelinesJobs(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PipelineIds(pipelineIds).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.GetOrganizationApiRestProvisioningPipelinesJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRestProvisioningPipelinesJobs`: InlineResponse200244
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.GetOrganizationApiRestProvisioningPipelinesJobs`: %v\n", resp)
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

[**InlineResponse200244**](InlineResponse200244.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline

> InlineResponse200245 GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(ctx, organizationId).PipelineIds(pipelineIds).Execute()

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
    resp, r, err := apiClient.ProvisioningApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(context.Background(), organizationId).PipelineIds(pipelineIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: InlineResponse200245
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: %v\n", resp)
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

[**InlineResponse200245**](InlineResponse200245.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesProvisioningStatuses

> []InlineResponse200312 GetOrganizationDevicesProvisioningStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the provisioning statuses information for devices in an organization.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device by network ID. This filter uses multiple exact matches. (optional)
    productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. (optional)
    status := "status_example" // string | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. (optional)
    tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
    tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.GetOrganizationDevicesProvisioningStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.GetOrganizationDevicesProvisioningStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationDevicesProvisioningStatuses`: []InlineResponse200312
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.GetOrganizationDevicesProvisioningStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesProvisioningStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. | 
 **status** | **string** | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]InlineResponse200312**](InlineResponse200312.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesProvisioningDeployments

> []InlineResponse200416 GetOrganizationWirelessDevicesProvisioningDeployments(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).SortBy(sortBy).SortOrder(sortOrder).DeploymentType(deploymentType).Execute()

List the zero touch deployments for wireless access points in an organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 20. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    search := "search_example" // string | Filter by MAC address, serial number, new device name, old device name, or model. (optional)
    sortBy := "sortBy_example" // string | Field used to sort results. Default is 'status'. (optional)
    sortOrder := "sortOrder_example" // string | Sort order. Default is 'asc'. (optional)
    deploymentType := "deploymentType_example" // string | Filter deployments by type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.GetOrganizationWirelessDevicesProvisioningDeployments(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).SortBy(sortBy).SortOrder(sortOrder).DeploymentType(deploymentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.GetOrganizationWirelessDevicesProvisioningDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesProvisioningDeployments`: []InlineResponse200416
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.GetOrganizationWirelessDevicesProvisioningDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesProvisioningDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 20. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **search** | **string** | Filter by MAC address, serial number, new device name, old device name, or model. | 
 **sortBy** | **string** | Field used to sort results. Default is &#39;status&#39;. | 
 **sortOrder** | **string** | Sort order. Default is &#39;asc&#39;. | 
 **deploymentType** | **string** | Filter deployments by type. | 

### Return type

[**[]InlineResponse200416**](InlineResponse200416.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessDevicesProvisioningDeployments

> InlineResponse200416 UpdateOrganizationWirelessDevicesProvisioningDeployments(ctx, organizationId).UpdateOrganizationWirelessDevicesProvisioningDeployments(updateOrganizationWirelessDevicesProvisioningDeployments).Execute()

Update a zero touch deployment



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
    updateOrganizationWirelessDevicesProvisioningDeployments := *openapiclient.NewInlineObject327([]openapiclient.OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1{*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1New()), "Status_example", "Type_example")}) // InlineObject327 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApi.UpdateOrganizationWirelessDevicesProvisioningDeployments(context.Background(), organizationId).UpdateOrganizationWirelessDevicesProvisioningDeployments(updateOrganizationWirelessDevicesProvisioningDeployments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApi.UpdateOrganizationWirelessDevicesProvisioningDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessDevicesProvisioningDeployments`: InlineResponse200416
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApi.UpdateOrganizationWirelessDevicesProvisioningDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationWirelessDevicesProvisioningDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationWirelessDevicesProvisioningDeployments** | [**InlineObject327**](InlineObject327.md) |  | 

### Return type

[**InlineResponse200416**](InlineResponse200416.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

