# \ApiApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAdministeredIdentitiesMeApiKeys**](ApiApi.md#GenerateAdministeredIdentitiesMeApiKeys) | **Post** /administered/identities/me/api/keys/generate | Generates an API key for an identity
[**GetAdministeredIdentitiesMeApiKeys**](ApiApi.md#GetAdministeredIdentitiesMeApiKeys) | **Get** /administered/identities/me/api/keys | List the non-sensitive metadata associated with the API keys that belong to the user
[**GetOrganizationApiRestProvisioningPipelinesJobs**](ApiApi.md#GetOrganizationApiRestProvisioningPipelinesJobs) | **Get** /organizations/{organizationId}/api/rest/provisioning/pipelines/jobs | List pipeline jobs, with optional status filtering
[**GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline**](ApiApi.md#GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline) | **Get** /organizations/{organizationId}/api/rest/provisioning/pipelines/jobs/overviews/byPipeline | Retrieves pipeline overviews with aggregated job status counts
[**RevokeAdministeredIdentitiesMeApiKeys**](ApiApi.md#RevokeAdministeredIdentitiesMeApiKeys) | **Post** /administered/identities/me/api/keys/{suffix}/revoke | Revokes an identity&#39;s API key, using the last four characters of the key



## GenerateAdministeredIdentitiesMeApiKeys

> InlineResponse202 GenerateAdministeredIdentitiesMeApiKeys(ctx).Execute()

Generates an API key for an identity



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiApi.GenerateAdministeredIdentitiesMeApiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.GenerateAdministeredIdentitiesMeApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAdministeredIdentitiesMeApiKeys`: InlineResponse202
    fmt.Fprintf(os.Stdout, "Response from `ApiApi.GenerateAdministeredIdentitiesMeApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAdministeredIdentitiesMeApiKeysRequest struct via the builder pattern


### Return type

[**InlineResponse202**](InlineResponse202.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdministeredIdentitiesMeApiKeys

> []InlineResponse2001 GetAdministeredIdentitiesMeApiKeys(ctx).Execute()

List the non-sensitive metadata associated with the API keys that belong to the user



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiApi.GetAdministeredIdentitiesMeApiKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.GetAdministeredIdentitiesMeApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredIdentitiesMeApiKeys`: []InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `ApiApi.GetAdministeredIdentitiesMeApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredIdentitiesMeApiKeysRequest struct via the builder pattern


### Return type

[**[]InlineResponse2001**](InlineResponse2001.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRestProvisioningPipelinesJobs

> InlineResponse200270 GetOrganizationApiRestProvisioningPipelinesJobs(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PipelineIds(pipelineIds).Status(status).Execute()

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
    resp, r, err := apiClient.ApiApi.GetOrganizationApiRestProvisioningPipelinesJobs(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).PipelineIds(pipelineIds).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.GetOrganizationApiRestProvisioningPipelinesJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRestProvisioningPipelinesJobs`: InlineResponse200270
    fmt.Fprintf(os.Stdout, "Response from `ApiApi.GetOrganizationApiRestProvisioningPipelinesJobs`: %v\n", resp)
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

[**InlineResponse200270**](InlineResponse200270.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline

> InlineResponse200271 GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(ctx, organizationId).PipelineIds(pipelineIds).Execute()

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
    resp, r, err := apiClient.ApiApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline(context.Background(), organizationId).PipelineIds(pipelineIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: InlineResponse200271
    fmt.Fprintf(os.Stdout, "Response from `ApiApi.GetOrganizationApiRestProvisioningPipelinesJobsOverviewsByPipeline`: %v\n", resp)
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

[**InlineResponse200271**](InlineResponse200271.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAdministeredIdentitiesMeApiKeys

> RevokeAdministeredIdentitiesMeApiKeys(ctx, suffix).Execute()

Revokes an identity's API key, using the last four characters of the key



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
    suffix := "suffix_example" // string | Suffix

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiApi.RevokeAdministeredIdentitiesMeApiKeys(context.Background(), suffix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.RevokeAdministeredIdentitiesMeApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**suffix** | **string** | Suffix | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAdministeredIdentitiesMeApiKeysRequest struct via the builder pattern


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

