# \DeploymentsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationWirelessDevicesProvisioningDeployment**](DeploymentsApi.md#CreateOrganizationWirelessDevicesProvisioningDeployment) | **Post** /organizations/{organizationId}/wireless/devices/provisioning/deployments | Create a zero touch deployment for a wireless access point
[**DeleteOrganizationWirelessDevicesProvisioningDeployment**](DeploymentsApi.md#DeleteOrganizationWirelessDevicesProvisioningDeployment) | **Delete** /organizations/{organizationId}/wireless/devices/provisioning/deployments/{deploymentId} | Delete a zero touch deployment
[**GetOrganizationWirelessDevicesProvisioningDeployments**](DeploymentsApi.md#GetOrganizationWirelessDevicesProvisioningDeployments) | **Get** /organizations/{organizationId}/wireless/devices/provisioning/deployments | List the zero touch deployments for wireless access points in an organization
[**UpdateOrganizationWirelessDevicesProvisioningDeployments**](DeploymentsApi.md#UpdateOrganizationWirelessDevicesProvisioningDeployments) | **Put** /organizations/{organizationId}/wireless/devices/provisioning/deployments | Update a zero touch deployment



## CreateOrganizationWirelessDevicesProvisioningDeployment

> InlineResponse200389 CreateOrganizationWirelessDevicesProvisioningDeployment(ctx, organizationId).CreateOrganizationWirelessDevicesProvisioningDeployment(createOrganizationWirelessDevicesProvisioningDeployment).Execute()

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
    createOrganizationWirelessDevicesProvisioningDeployment := *openapiclient.NewInlineObject310([]openapiclient.OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1{*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1New()), "Status_example", "Type_example")}) // InlineObject310 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.CreateOrganizationWirelessDevicesProvisioningDeployment(context.Background(), organizationId).CreateOrganizationWirelessDevicesProvisioningDeployment(createOrganizationWirelessDevicesProvisioningDeployment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CreateOrganizationWirelessDevicesProvisioningDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationWirelessDevicesProvisioningDeployment`: InlineResponse200389
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.CreateOrganizationWirelessDevicesProvisioningDeployment`: %v\n", resp)
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

 **createOrganizationWirelessDevicesProvisioningDeployment** | [**InlineObject310**](InlineObject310.md) |  | 

### Return type

[**InlineResponse200389**](InlineResponse200389.md)

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
    resp, r, err := apiClient.DeploymentsApi.DeleteOrganizationWirelessDevicesProvisioningDeployment(context.Background(), organizationId, deploymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DeleteOrganizationWirelessDevicesProvisioningDeployment``: %v\n", err)
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


## GetOrganizationWirelessDevicesProvisioningDeployments

> []InlineResponse200389 GetOrganizationWirelessDevicesProvisioningDeployments(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).SortBy(sortBy).SortOrder(sortOrder).DeploymentType(deploymentType).Execute()

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
    resp, r, err := apiClient.DeploymentsApi.GetOrganizationWirelessDevicesProvisioningDeployments(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Search(search).SortBy(sortBy).SortOrder(sortOrder).DeploymentType(deploymentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetOrganizationWirelessDevicesProvisioningDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesProvisioningDeployments`: []InlineResponse200389
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetOrganizationWirelessDevicesProvisioningDeployments`: %v\n", resp)
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

[**[]InlineResponse200389**](InlineResponse200389.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationWirelessDevicesProvisioningDeployments

> InlineResponse200389 UpdateOrganizationWirelessDevicesProvisioningDeployments(ctx, organizationId).UpdateOrganizationWirelessDevicesProvisioningDeployments(updateOrganizationWirelessDevicesProvisioningDeployments).Execute()

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
    updateOrganizationWirelessDevicesProvisioningDeployments := *openapiclient.NewInlineObject309([]openapiclient.OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1{*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1(*openapiclient.NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1New()), "Status_example", "Type_example")}) // InlineObject309 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentsApi.UpdateOrganizationWirelessDevicesProvisioningDeployments(context.Background(), organizationId).UpdateOrganizationWirelessDevicesProvisioningDeployments(updateOrganizationWirelessDevicesProvisioningDeployments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.UpdateOrganizationWirelessDevicesProvisioningDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationWirelessDevicesProvisioningDeployments`: InlineResponse200389
    fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.UpdateOrganizationWirelessDevicesProvisioningDeployments`: %v\n", resp)
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

 **updateOrganizationWirelessDevicesProvisioningDeployments** | [**InlineObject309**](InlineObject309.md) |  | 

### Return type

[**InlineResponse200389**](InlineResponse200389.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

