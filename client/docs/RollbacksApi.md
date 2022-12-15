# \RollbacksApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkFirmwareUpgradesRollback**](RollbacksApi.md#CreateNetworkFirmwareUpgradesRollback) | **Post** /networks/{networkId}/firmwareUpgrades/rollbacks | Rollback a Firmware Upgrade For A Network



## CreateNetworkFirmwareUpgradesRollback

> InlineResponse20020 CreateNetworkFirmwareUpgradesRollback(ctx, networkId).CreateNetworkFirmwareUpgradesRollback(createNetworkFirmwareUpgradesRollback).Execute()

Rollback a Firmware Upgrade For A Network



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
    networkId := "networkId_example" // string | 
    createNetworkFirmwareUpgradesRollback := *openapiclient.NewInlineObject77([]openapiclient.NetworksNetworkIdFirmwareUpgradesRollbacksReasons{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesRollbacksReasons("Category_example", "Comment_example")}) // InlineObject77 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RollbacksApi.CreateNetworkFirmwareUpgradesRollback(context.Background(), networkId).CreateNetworkFirmwareUpgradesRollback(createNetworkFirmwareUpgradesRollback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RollbacksApi.CreateNetworkFirmwareUpgradesRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesRollback`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `RollbacksApi.CreateNetworkFirmwareUpgradesRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesRollback** | [**InlineObject77**](InlineObject77.md) |  | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

