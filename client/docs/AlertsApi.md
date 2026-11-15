# \AlertsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSensorAlertsProfile**](AlertsApi.md#CreateNetworkSensorAlertsProfile) | **Post** /networks/{networkId}/sensor/alerts/profiles | Creates a sensor alert profile for a network.
[**CreateOrganizationAlertsProfile**](AlertsApi.md#CreateOrganizationAlertsProfile) | **Post** /organizations/{organizationId}/alerts/profiles | Create an organization-wide alert configuration
[**CreateOrganizationAssuranceAlertsProfile**](AlertsApi.md#CreateOrganizationAssuranceAlertsProfile) | **Post** /organizations/{organizationId}/assurance/alerts/profiles | Create an alert profile
[**DeleteNetworkSensorAlertsProfile**](AlertsApi.md#DeleteNetworkSensorAlertsProfile) | **Delete** /networks/{networkId}/sensor/alerts/profiles/{id} | Deletes a sensor alert profile from a network.
[**DeleteOrganizationAlertsProfile**](AlertsApi.md#DeleteOrganizationAlertsProfile) | **Delete** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Removes an organization-wide alert config
[**DeleteOrganizationAssuranceAlertsProfile**](AlertsApi.md#DeleteOrganizationAssuranceAlertsProfile) | **Delete** /organizations/{organizationId}/assurance/alerts/profiles/{profileId} | Delete an alert profile for this organization
[**DismissOrganizationAssuranceAlerts**](AlertsApi.md#DismissOrganizationAssuranceAlerts) | **Post** /organizations/{organizationId}/assurance/alerts/dismiss | Dismiss health alerts
[**GetNetworkAlertsHistory**](AlertsApi.md#GetNetworkAlertsHistory) | **Get** /networks/{networkId}/alerts/history | Return the alert history for this network
[**GetNetworkAlertsSettings**](AlertsApi.md#GetNetworkAlertsSettings) | **Get** /networks/{networkId}/alerts/settings | Return the alert configuration for this network
[**GetNetworkHealthAlerts**](AlertsApi.md#GetNetworkHealthAlerts) | **Get** /networks/{networkId}/health/alerts | Return all global alerts on this network
[**GetNetworkSensorAlertsCurrentOverviewByMetric**](AlertsApi.md#GetNetworkSensorAlertsCurrentOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/current/overview/byMetric | Return an overview of currently alerting sensors by metric
[**GetNetworkSensorAlertsOverviewByMetric**](AlertsApi.md#GetNetworkSensorAlertsOverviewByMetric) | **Get** /networks/{networkId}/sensor/alerts/overview/byMetric | Return an overview of alert occurrences over a timespan, by metric
[**GetNetworkSensorAlertsProfile**](AlertsApi.md#GetNetworkSensorAlertsProfile) | **Get** /networks/{networkId}/sensor/alerts/profiles/{id} | Show details of a sensor alert profile for a network.
[**GetNetworkSensorAlertsProfiles**](AlertsApi.md#GetNetworkSensorAlertsProfiles) | **Get** /networks/{networkId}/sensor/alerts/profiles | Lists all sensor alert profiles for a network.
[**GetOrganizationAlertsProfiles**](AlertsApi.md#GetOrganizationAlertsProfiles) | **Get** /organizations/{organizationId}/alerts/profiles | List all organization-wide alert configurations
[**GetOrganizationAssuranceAlert**](AlertsApi.md#GetOrganizationAssuranceAlert) | **Get** /organizations/{organizationId}/assurance/alerts/{id} | Return a singular Health Alert by its id
[**GetOrganizationAssuranceAlerts**](AlertsApi.md#GetOrganizationAssuranceAlerts) | **Get** /organizations/{organizationId}/assurance/alerts | Return all health alerts for an organization
[**GetOrganizationAssuranceAlertsOverview**](AlertsApi.md#GetOrganizationAssuranceAlertsOverview) | **Get** /organizations/{organizationId}/assurance/alerts/overview | Return overview of active health alerts for an organization
[**GetOrganizationAssuranceAlertsOverviewByNetwork**](AlertsApi.md#GetOrganizationAssuranceAlertsOverviewByNetwork) | **Get** /organizations/{organizationId}/assurance/alerts/overview/byNetwork | Return a Summary of Alerts grouped by network and severity
[**GetOrganizationAssuranceAlertsOverviewByType**](AlertsApi.md#GetOrganizationAssuranceAlertsOverviewByType) | **Get** /organizations/{organizationId}/assurance/alerts/overview/byType | Return a Summary of Alerts grouped by type and severity
[**GetOrganizationAssuranceAlertsOverviewHistorical**](AlertsApi.md#GetOrganizationAssuranceAlertsOverviewHistorical) | **Get** /organizations/{organizationId}/assurance/alerts/overview/historical | Returns historical health alert overviews
[**GetOrganizationAssuranceAlertsProfiles**](AlertsApi.md#GetOrganizationAssuranceAlertsProfiles) | **Get** /organizations/{organizationId}/assurance/alerts/profiles | List the alert profiles for this organization
[**GetOrganizationAssuranceAlertsTaxonomyCategories**](AlertsApi.md#GetOrganizationAssuranceAlertsTaxonomyCategories) | **Get** /organizations/{organizationId}/assurance/alerts/taxonomy/categories | Return a list of Category Types
[**GetOrganizationAssuranceAlertsTaxonomyTypes**](AlertsApi.md#GetOrganizationAssuranceAlertsTaxonomyTypes) | **Get** /organizations/{organizationId}/assurance/alerts/taxonomy/types | Return a list of alert types
[**RestoreOrganizationAssuranceAlerts**](AlertsApi.md#RestoreOrganizationAssuranceAlerts) | **Post** /organizations/{organizationId}/assurance/alerts/restore | Restore health alerts from dismissed
[**UpdateNetworkAlertsSettings**](AlertsApi.md#UpdateNetworkAlertsSettings) | **Put** /networks/{networkId}/alerts/settings | Update the alert configuration for this network
[**UpdateNetworkSensorAlertsProfile**](AlertsApi.md#UpdateNetworkSensorAlertsProfile) | **Put** /networks/{networkId}/sensor/alerts/profiles/{id} | Updates a sensor alert profile for a network.
[**UpdateOrganizationAlertsProfile**](AlertsApi.md#UpdateOrganizationAlertsProfile) | **Put** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Update an organization-wide alert config
[**UpdateOrganizationAssuranceAlertsProfile**](AlertsApi.md#UpdateOrganizationAssuranceAlertsProfile) | **Put** /organizations/{organizationId}/assurance/alerts/profiles/{profileId} | Update an alert profile



## CreateNetworkSensorAlertsProfile

> InlineResponse200161 CreateNetworkSensorAlertsProfile(ctx, networkId).CreateNetworkSensorAlertsProfile(createNetworkSensorAlertsProfile).Execute()

Creates a sensor alert profile for a network.



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
    createNetworkSensorAlertsProfile := *openapiclient.NewInlineObject144("Name_example", []openapiclient.NetworksNetworkIdSensorAlertsProfilesConditions{*openapiclient.NewNetworksNetworkIdSensorAlertsProfilesConditions("Metric_example", *openapiclient.NewNetworksNetworkIdSensorAlertsProfilesThreshold())}) // InlineObject144 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.CreateNetworkSensorAlertsProfile(context.Background(), networkId).CreateNetworkSensorAlertsProfile(createNetworkSensorAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSensorAlertsProfile`: InlineResponse200161
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSensorAlertsProfile** | [**InlineObject144**](InlineObject144.md) |  | 

### Return type

[**InlineResponse200161**](InlineResponse200161.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAlertsProfile

> InlineResponse200274 CreateOrganizationAlertsProfile(ctx, organizationId).CreateOrganizationAlertsProfile(createOrganizationAlertsProfile).Execute()

Create an organization-wide alert configuration



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
    createOrganizationAlertsProfile := *openapiclient.NewInlineObject251("Type_example", *openapiclient.NewOrganizationsOrganizationIdAlertsProfilesAlertCondition1(), *openapiclient.NewOrganizationsOrganizationIdAlertsProfilesRecipients(), []string{"NetworkTags_example"}) // InlineObject251 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.CreateOrganizationAlertsProfile(context.Background(), organizationId).CreateOrganizationAlertsProfile(createOrganizationAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAlertsProfile`: InlineResponse200274
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAlertsProfile** | [**InlineObject251**](InlineObject251.md) |  | 

### Return type

[**InlineResponse200274**](InlineResponse200274.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationAssuranceAlertsProfile

> InlineResponse200323Items CreateOrganizationAssuranceAlertsProfile(ctx, organizationId).CreateOrganizationAssuranceAlertsProfile(createOrganizationAssuranceAlertsProfile).Execute()

Create an alert profile



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
    createOrganizationAssuranceAlertsProfile := *openapiclient.NewInlineObject278("Name_example", []string{"NetworkIds_example"}, []string{"AlertTypes_example"}, *openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration(false, *openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations(*openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsEmail(false), *openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsWebhook(false)))) // InlineObject278 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.CreateOrganizationAssuranceAlertsProfile(context.Background(), organizationId).CreateOrganizationAssuranceAlertsProfile(createOrganizationAssuranceAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateOrganizationAssuranceAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAssuranceAlertsProfile`: InlineResponse200323Items
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateOrganizationAssuranceAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAssuranceAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAssuranceAlertsProfile** | [**InlineObject278**](InlineObject278.md) |  | 

### Return type

[**InlineResponse200323Items**](InlineResponse200323Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSensorAlertsProfile

> DeleteNetworkSensorAlertsProfile(ctx, networkId, id).Execute()

Deletes a sensor alert profile from a network.



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.DeleteNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSensorAlertsProfileRequest struct via the builder pattern


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


## DeleteOrganizationAlertsProfile

> DeleteOrganizationAlertsProfile(ctx, organizationId, alertConfigId).Execute()

Removes an organization-wide alert config



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
    alertConfigId := "alertConfigId_example" // string | Alert config ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.DeleteOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAlertsProfileRequest struct via the builder pattern


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


## DeleteOrganizationAssuranceAlertsProfile

> DeleteOrganizationAssuranceAlertsProfile(ctx, organizationId, profileId).Execute()

Delete an alert profile for this organization



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
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.DeleteOrganizationAssuranceAlertsProfile(context.Background(), organizationId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteOrganizationAssuranceAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAssuranceAlertsProfileRequest struct via the builder pattern


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


## DismissOrganizationAssuranceAlerts

> DismissOrganizationAssuranceAlerts(ctx, organizationId).DismissOrganizationAssuranceAlerts(dismissOrganizationAssuranceAlerts).Execute()

Dismiss health alerts



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
    dismissOrganizationAssuranceAlerts := *openapiclient.NewInlineObject277([]string{"AlertIds_example"}) // InlineObject277 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.DismissOrganizationAssuranceAlerts(context.Background(), organizationId).DismissOrganizationAssuranceAlerts(dismissOrganizationAssuranceAlerts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DismissOrganizationAssuranceAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDismissOrganizationAssuranceAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dismissOrganizationAssuranceAlerts** | [**InlineObject277**](InlineObject277.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsHistory

> []InlineResponse20070 GetNetworkAlertsHistory(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the alert history for this network



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.GetNetworkAlertsHistory(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNetworkAlertsHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAlertsHistory`: []InlineResponse20070
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNetworkAlertsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]InlineResponse20070**](InlineResponse20070.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsSettings

> InlineResponse20071 GetNetworkAlertsSettings(ctx, networkId).Execute()

Return the alert configuration for this network



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
    resp, r, err := apiClient.AlertsApi.GetNetworkAlertsSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNetworkAlertsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAlertsSettings`: InlineResponse20071
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20071**](InlineResponse20071.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkHealthAlerts

> []InlineResponse200150 GetNetworkHealthAlerts(ctx, networkId).Execute()

Return all global alerts on this network



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
    resp, r, err := apiClient.AlertsApi.GetNetworkHealthAlerts(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNetworkHealthAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkHealthAlerts`: []InlineResponse200150
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNetworkHealthAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkHealthAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200150**](InlineResponse200150.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsCurrentOverviewByMetric

> InlineResponse200159 GetNetworkSensorAlertsCurrentOverviewByMetric(ctx, networkId).Execute()

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
    resp, r, err := apiClient.AlertsApi.GetNetworkSensorAlertsCurrentOverviewByMetric(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNetworkSensorAlertsCurrentOverviewByMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsCurrentOverviewByMetric`: InlineResponse200159
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNetworkSensorAlertsCurrentOverviewByMetric`: %v\n", resp)
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

[**InlineResponse200159**](InlineResponse200159.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsOverviewByMetric

> []InlineResponse200160 GetNetworkSensorAlertsOverviewByMetric(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

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
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 900, 3600, 86400, 604800, 2629746. The default is 604800. Interval is calculated if time params are provided. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.GetNetworkSensorAlertsOverviewByMetric(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNetworkSensorAlertsOverviewByMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsOverviewByMetric`: []InlineResponse200160
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNetworkSensorAlertsOverviewByMetric`: %v\n", resp)
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
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 900, 3600, 86400, 604800, 2629746. The default is 604800. Interval is calculated if time params are provided. | 

### Return type

[**[]InlineResponse200160**](InlineResponse200160.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsProfile

> InlineResponse200161 GetNetworkSensorAlertsProfile(ctx, networkId, id).Execute()

Show details of a sensor alert profile for a network.



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.GetNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsProfile`: InlineResponse200161
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200161**](InlineResponse200161.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsProfiles

> []InlineResponse200161 GetNetworkSensorAlertsProfiles(ctx, networkId).Execute()

Lists all sensor alert profiles for a network.



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
    resp, r, err := apiClient.AlertsApi.GetNetworkSensorAlertsProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNetworkSensorAlertsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsProfiles`: []InlineResponse200161
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNetworkSensorAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200161**](InlineResponse200161.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAlertsProfiles

> []InlineResponse200274 GetOrganizationAlertsProfiles(ctx, organizationId).Execute()

List all organization-wide alert configurations



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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAlertsProfiles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAlertsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAlertsProfiles`: []InlineResponse200274
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200274**](InlineResponse200274.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlert

> InlineResponse200318 GetOrganizationAssuranceAlert(ctx, organizationId, id).Execute()

Return a singular Health Alert by its id



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlert(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlert`: InlineResponse200318
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200318**](InlineResponse200318.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlerts

> []InlineResponse200318 GetOrganizationAssuranceAlerts(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

Return all health alerts for an organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 4 - 300. Default is 30. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'. (optional)
    networkId := "networkId_example" // string | Optional parameter to filter alerts by network ids. (optional)
    severity := "severity_example" // string | Optional parameter to filter by severity type. (optional)
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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlerts(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlerts`: []InlineResponse200318
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 4 - 300. Default is 30. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;ascending&#39;. | 
 **networkId** | **string** | Optional parameter to filter alerts by network ids. | 
 **severity** | **string** | Optional parameter to filter by severity type. | 
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

[**[]InlineResponse200318**](InlineResponse200318.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverview

> InlineResponse200319 GetOrganizationAssuranceAlertsOverview(ctx, organizationId).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlertsOverview(context.Background(), organizationId).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlertsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverview`: InlineResponse200319
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlertsOverview`: %v\n", resp)
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

[**InlineResponse200319**](InlineResponse200319.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewByNetwork

> InlineResponse200320 GetOrganizationAssuranceAlertsOverviewByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlertsOverviewByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlertsOverviewByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverviewByNetwork`: InlineResponse200320
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlertsOverviewByNetwork`: %v\n", resp)
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

[**InlineResponse200320**](InlineResponse200320.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewByType

> InlineResponse200321 GetOrganizationAssuranceAlertsOverviewByType(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).IncludeDeviceTags(includeDeviceTags).IncludeNetworks(includeNetworks).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()

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
    includeDeviceTags := true // bool | Include grouped device tags for each alert type in the response. (optional)
    includeNetworks := true // bool | Include affected networks for each alert type in the response. (optional)
    suppressAlertsForOfflineNodes := true // bool | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlertsOverviewByType(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).NetworkId(networkId).Severity(severity).Types(types).TsStart(tsStart).TsEnd(tsEnd).Category(category).SortBy(sortBy).Serials(serials).DeviceTypes(deviceTypes).DeviceTags(deviceTags).Active(active).Dismissed(dismissed).Resolved(resolved).IncludeDeviceTags(includeDeviceTags).IncludeNetworks(includeNetworks).SuppressAlertsForOfflineNodes(suppressAlertsForOfflineNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlertsOverviewByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverviewByType`: InlineResponse200321
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlertsOverviewByType`: %v\n", resp)
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
 **includeDeviceTags** | **bool** | Include grouped device tags for each alert type in the response. | 
 **includeNetworks** | **bool** | Include affected networks for each alert type in the response. | 
 **suppressAlertsForOfflineNodes** | **bool** | When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false. | 

### Return type

[**InlineResponse200321**](InlineResponse200321.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsOverviewHistorical

> InlineResponse200322 GetOrganizationAssuranceAlertsOverviewHistorical(ctx, organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).Execute()

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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlertsOverviewHistorical(context.Background(), organizationId).SegmentDuration(segmentDuration).TsStart(tsStart).NetworkId(networkId).Severity(severity).Types(types).TsEnd(tsEnd).Category(category).Serials(serials).DeviceTypes(deviceTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlertsOverviewHistorical``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsOverviewHistorical`: InlineResponse200322
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlertsOverviewHistorical`: %v\n", resp)
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

[**InlineResponse200322**](InlineResponse200322.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsProfiles

> InlineResponse200323 GetOrganizationAssuranceAlertsProfiles(ctx, organizationId).Execute()

List the alert profiles for this organization



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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlertsProfiles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlertsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsProfiles`: InlineResponse200323
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200323**](InlineResponse200323.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsTaxonomyCategories

> []InlineResponse200324 GetOrganizationAssuranceAlertsTaxonomyCategories(ctx, organizationId).Execute()

Return a list of Category Types



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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlertsTaxonomyCategories(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlertsTaxonomyCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsTaxonomyCategories`: []InlineResponse200324
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlertsTaxonomyCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsTaxonomyCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200324**](InlineResponse200324.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAssuranceAlertsTaxonomyTypes

> []InlineResponse200325 GetOrganizationAssuranceAlertsTaxonomyTypes(ctx, organizationId).Execute()

Return a list of alert types



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
    resp, r, err := apiClient.AlertsApi.GetOrganizationAssuranceAlertsTaxonomyTypes(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetOrganizationAssuranceAlertsTaxonomyTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAssuranceAlertsTaxonomyTypes`: []InlineResponse200325
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetOrganizationAssuranceAlertsTaxonomyTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAssuranceAlertsTaxonomyTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200325**](InlineResponse200325.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreOrganizationAssuranceAlerts

> RestoreOrganizationAssuranceAlerts(ctx, organizationId).RestoreOrganizationAssuranceAlerts(restoreOrganizationAssuranceAlerts).Execute()

Restore health alerts from dismissed



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
    restoreOrganizationAssuranceAlerts := *openapiclient.NewInlineObject280([]string{"AlertIds_example"}) // InlineObject280 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.RestoreOrganizationAssuranceAlerts(context.Background(), organizationId).RestoreOrganizationAssuranceAlerts(restoreOrganizationAssuranceAlerts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.RestoreOrganizationAssuranceAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreOrganizationAssuranceAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreOrganizationAssuranceAlerts** | [**InlineObject280**](InlineObject280.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAlertsSettings

> InlineResponse20071 UpdateNetworkAlertsSettings(ctx, networkId).UpdateNetworkAlertsSettings(updateNetworkAlertsSettings).Execute()

Update the alert configuration for this network



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
    updateNetworkAlertsSettings := *openapiclient.NewInlineObject53() // InlineObject53 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.UpdateNetworkAlertsSettings(context.Background(), networkId).UpdateNetworkAlertsSettings(updateNetworkAlertsSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateNetworkAlertsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAlertsSettings`: InlineResponse20071
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAlertsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkAlertsSettings** | [**InlineObject53**](InlineObject53.md) |  | 

### Return type

[**InlineResponse20071**](InlineResponse20071.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSensorAlertsProfile

> InlineResponse200161 UpdateNetworkSensorAlertsProfile(ctx, networkId, id).UpdateNetworkSensorAlertsProfile(updateNetworkSensorAlertsProfile).Execute()

Updates a sensor alert profile for a network.



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
    id := "id_example" // string | ID
    updateNetworkSensorAlertsProfile := *openapiclient.NewInlineObject145() // InlineObject145 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.UpdateNetworkSensorAlertsProfile(context.Background(), networkId, id).UpdateNetworkSensorAlertsProfile(updateNetworkSensorAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSensorAlertsProfile`: InlineResponse200161
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSensorAlertsProfile** | [**InlineObject145**](InlineObject145.md) |  | 

### Return type

[**InlineResponse200161**](InlineResponse200161.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAlertsProfile

> InlineResponse200274 UpdateOrganizationAlertsProfile(ctx, organizationId, alertConfigId).UpdateOrganizationAlertsProfile(updateOrganizationAlertsProfile).Execute()

Update an organization-wide alert config



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
    alertConfigId := "alertConfigId_example" // string | Alert config ID
    updateOrganizationAlertsProfile := *openapiclient.NewInlineObject252() // InlineObject252 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.UpdateOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).UpdateOrganizationAlertsProfile(updateOrganizationAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAlertsProfile`: InlineResponse200274
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAlertsProfile** | [**InlineObject252**](InlineObject252.md) |  | 

### Return type

[**InlineResponse200274**](InlineResponse200274.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAssuranceAlertsProfile

> InlineResponse200323Items UpdateOrganizationAssuranceAlertsProfile(ctx, organizationId, profileId).UpdateOrganizationAssuranceAlertsProfile(updateOrganizationAssuranceAlertsProfile).Execute()

Update an alert profile



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
    profileId := "profileId_example" // string | Profile ID
    updateOrganizationAssuranceAlertsProfile := *openapiclient.NewInlineObject279("Name_example", []string{"NetworkIds_example"}, []string{"AlertTypes_example"}, *openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfiguration(false, *openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinations(*openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsEmail(false), *openapiclient.NewOrganizationsOrganizationIdAssuranceAlertsProfilesConfigurationAlertDestinationsWebhook(false)))) // InlineObject279 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertsApi.UpdateOrganizationAssuranceAlertsProfile(context.Background(), organizationId, profileId).UpdateOrganizationAssuranceAlertsProfile(updateOrganizationAssuranceAlertsProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateOrganizationAssuranceAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAssuranceAlertsProfile`: InlineResponse200323Items
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateOrganizationAssuranceAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAssuranceAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAssuranceAlertsProfile** | [**InlineObject279**](InlineObject279.md) |  | 

### Return type

[**InlineResponse200323Items**](InlineResponse200323Items.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

