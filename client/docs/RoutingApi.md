# \RoutingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceSwitchRoutingInterface**](RoutingApi.md#CreateDeviceSwitchRoutingInterface) | **Post** /devices/{serial}/switch/routing/interfaces | Create a layer 3 interface for a switch
[**CreateDeviceSwitchRoutingStaticRoute**](RoutingApi.md#CreateDeviceSwitchRoutingStaticRoute) | **Post** /devices/{serial}/switch/routing/staticRoutes | Create a layer 3 static route for a switch
[**CreateNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingApi.md#CreateNetworkSwitchRoutingMulticastRendezvousPoint) | **Post** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | Create a multicast rendezvous point
[**CreateNetworkSwitchStackRoutingInterface**](RoutingApi.md#CreateNetworkSwitchStackRoutingInterface) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | Create a layer 3 interface for a switch stack
[**CreateNetworkSwitchStackRoutingStaticRoute**](RoutingApi.md#CreateNetworkSwitchStackRoutingStaticRoute) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | Create a layer 3 static route for a switch stack
[**DeleteDeviceSwitchRoutingInterface**](RoutingApi.md#DeleteDeviceSwitchRoutingInterface) | **Delete** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Delete a layer 3 interface from the switch
[**DeleteDeviceSwitchRoutingStaticRoute**](RoutingApi.md#DeleteDeviceSwitchRoutingStaticRoute) | **Delete** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch
[**DeleteNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingApi.md#DeleteNetworkSwitchRoutingMulticastRendezvousPoint) | **Delete** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Delete a multicast rendezvous point
[**DeleteNetworkSwitchStackRoutingInterface**](RoutingApi.md#DeleteNetworkSwitchStackRoutingInterface) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Delete a layer 3 interface from a switch stack
[**DeleteNetworkSwitchStackRoutingStaticRoute**](RoutingApi.md#DeleteNetworkSwitchStackRoutingStaticRoute) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch stack
[**GetDeviceSwitchRoutingInterface**](RoutingApi.md#GetDeviceSwitchRoutingInterface) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Return a layer 3 interface for a switch
[**GetDeviceSwitchRoutingInterfaceDhcp**](RoutingApi.md#GetDeviceSwitchRoutingInterfaceDhcp) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch
[**GetDeviceSwitchRoutingInterfaces**](RoutingApi.md#GetDeviceSwitchRoutingInterfaces) | **Get** /devices/{serial}/switch/routing/interfaces | List layer 3 interfaces for a switch
[**GetDeviceSwitchRoutingStaticRoute**](RoutingApi.md#GetDeviceSwitchRoutingStaticRoute) | **Get** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch
[**GetDeviceSwitchRoutingStaticRoutes**](RoutingApi.md#GetDeviceSwitchRoutingStaticRoutes) | **Get** /devices/{serial}/switch/routing/staticRoutes | List layer 3 static routes for a switch
[**GetNetworkSwitchRoutingMulticast**](RoutingApi.md#GetNetworkSwitchRoutingMulticast) | **Get** /networks/{networkId}/switch/routing/multicast | Return multicast settings for a network
[**GetNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingApi.md#GetNetworkSwitchRoutingMulticastRendezvousPoint) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Return a multicast rendezvous point
[**GetNetworkSwitchRoutingMulticastRendezvousPoints**](RoutingApi.md#GetNetworkSwitchRoutingMulticastRendezvousPoints) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | List multicast rendezvous points
[**GetNetworkSwitchRoutingOspf**](RoutingApi.md#GetNetworkSwitchRoutingOspf) | **Get** /networks/{networkId}/switch/routing/ospf | Return layer 3 OSPF routing configuration
[**GetNetworkSwitchStackRoutingInterface**](RoutingApi.md#GetNetworkSwitchStackRoutingInterface) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Return a layer 3 interface from a switch stack
[**GetNetworkSwitchStackRoutingInterfaceDhcp**](RoutingApi.md#GetNetworkSwitchStackRoutingInterfaceDhcp) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch stack
[**GetNetworkSwitchStackRoutingInterfaces**](RoutingApi.md#GetNetworkSwitchStackRoutingInterfaces) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | List layer 3 interfaces for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoute**](RoutingApi.md#GetNetworkSwitchStackRoutingStaticRoute) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoutes**](RoutingApi.md#GetNetworkSwitchStackRoutingStaticRoutes) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | List layer 3 static routes for a switch stack
[**UpdateDeviceSwitchRoutingInterface**](RoutingApi.md#UpdateDeviceSwitchRoutingInterface) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch
[**UpdateDeviceSwitchRoutingInterfaceDhcp**](RoutingApi.md#UpdateDeviceSwitchRoutingInterfaceDhcp) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch
[**UpdateDeviceSwitchRoutingStaticRoute**](RoutingApi.md#UpdateDeviceSwitchRoutingStaticRoute) | **Put** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch
[**UpdateNetworkSwitchRoutingMulticast**](RoutingApi.md#UpdateNetworkSwitchRoutingMulticast) | **Put** /networks/{networkId}/switch/routing/multicast | Update multicast settings for a network
[**UpdateNetworkSwitchRoutingMulticastRendezvousPoint**](RoutingApi.md#UpdateNetworkSwitchRoutingMulticastRendezvousPoint) | **Put** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Update a multicast rendezvous point
[**UpdateNetworkSwitchRoutingOspf**](RoutingApi.md#UpdateNetworkSwitchRoutingOspf) | **Put** /networks/{networkId}/switch/routing/ospf | Update layer 3 OSPF routing configuration
[**UpdateNetworkSwitchStackRoutingInterface**](RoutingApi.md#UpdateNetworkSwitchStackRoutingInterface) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch stack
[**UpdateNetworkSwitchStackRoutingInterfaceDhcp**](RoutingApi.md#UpdateNetworkSwitchStackRoutingInterfaceDhcp) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch stack
[**UpdateNetworkSwitchStackRoutingStaticRoute**](RoutingApi.md#UpdateNetworkSwitchStackRoutingStaticRoute) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch stack



## CreateDeviceSwitchRoutingInterface

> InlineResponse20039 CreateDeviceSwitchRoutingInterface(ctx, serial).CreateDeviceSwitchRoutingInterface(createDeviceSwitchRoutingInterface).Execute()

Create a layer 3 interface for a switch



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
    createDeviceSwitchRoutingInterface := *openapiclient.NewInlineObject29("Name_example") // InlineObject29 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.CreateDeviceSwitchRoutingInterface(context.Background(), serial).CreateDeviceSwitchRoutingInterface(createDeviceSwitchRoutingInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.CreateDeviceSwitchRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceSwitchRoutingInterface`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.CreateDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSwitchRoutingInterface** | [**InlineObject29**](InlineObject29.md) |  | 

### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceSwitchRoutingStaticRoute

> InlineResponse20041 CreateDeviceSwitchRoutingStaticRoute(ctx, serial).CreateDeviceSwitchRoutingStaticRoute(createDeviceSwitchRoutingStaticRoute).Execute()

Create a layer 3 static route for a switch



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
    createDeviceSwitchRoutingStaticRoute := *openapiclient.NewInlineObject32("Subnet_example", "NextHopIp_example") // InlineObject32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.CreateDeviceSwitchRoutingStaticRoute(context.Background(), serial).CreateDeviceSwitchRoutingStaticRoute(createDeviceSwitchRoutingStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.CreateDeviceSwitchRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceSwitchRoutingStaticRoute`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.CreateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSwitchRoutingStaticRoute** | [**InlineObject32**](InlineObject32.md) |  | 

### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchRoutingMulticastRendezvousPoint

> InlineResponse200164 CreateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId).CreateNetworkSwitchRoutingMulticastRendezvousPoint(createNetworkSwitchRoutingMulticastRendezvousPoint).Execute()

Create a multicast rendezvous point



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
    createNetworkSwitchRoutingMulticastRendezvousPoint := *openapiclient.NewInlineObject152("InterfaceIp_example", "MulticastGroup_example") // InlineObject152 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).CreateNetworkSwitchRoutingMulticastRendezvousPoint(createNetworkSwitchRoutingMulticastRendezvousPoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.CreateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchRoutingMulticastRendezvousPoint`: InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.CreateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchRoutingMulticastRendezvousPoint** | [**InlineObject152**](InlineObject152.md) |  | 

### Return type

[**InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchStackRoutingInterface

> InlineResponse20039 CreateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId).CreateNetworkSwitchStackRoutingInterface(createNetworkSwitchStackRoutingInterface).Execute()

Create a layer 3 interface for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    createNetworkSwitchStackRoutingInterface := *openapiclient.NewInlineObject159("Name_example") // InlineObject159 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.CreateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId).CreateNetworkSwitchStackRoutingInterface(createNetworkSwitchStackRoutingInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.CreateNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchStackRoutingInterface`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.CreateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchStackRoutingInterface** | [**InlineObject159**](InlineObject159.md) |  | 

### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchStackRoutingStaticRoute

> InlineResponse20041 CreateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId).CreateNetworkSwitchStackRoutingStaticRoute(createNetworkSwitchStackRoutingStaticRoute).Execute()

Create a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    createNetworkSwitchStackRoutingStaticRoute := *openapiclient.NewInlineObject162("Subnet_example", "NextHopIp_example") // InlineObject162 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).CreateNetworkSwitchStackRoutingStaticRoute(createNetworkSwitchStackRoutingStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.CreateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchStackRoutingStaticRoute`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.CreateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchStackRoutingStaticRoute** | [**InlineObject162**](InlineObject162.md) |  | 

### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceSwitchRoutingInterface

> DeleteDeviceSwitchRoutingInterface(ctx, serial, interfaceId).Execute()

Delete a layer 3 interface from the switch



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
    interfaceId := "interfaceId_example" // string | Interface ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteDeviceSwitchRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


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


## DeleteDeviceSwitchRoutingStaticRoute

> DeleteDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).Execute()

Delete a layer 3 static route for a switch



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
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteDeviceSwitchRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


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


## DeleteNetworkSwitchRoutingMulticastRendezvousPoint

> DeleteNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Delete a multicast rendezvous point



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
    rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


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


## DeleteNetworkSwitchStackRoutingInterface

> DeleteNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Delete a layer 3 interface from a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    interfaceId := "interfaceId_example" // string | Interface ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


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


## DeleteNetworkSwitchStackRoutingStaticRoute

> DeleteNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Delete a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


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


## GetDeviceSwitchRoutingInterface

> InlineResponse20039 GetDeviceSwitchRoutingInterface(ctx, serial, interfaceId).Execute()

Return a layer 3 interface for a switch



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
    interfaceId := "interfaceId_example" // string | Interface ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetDeviceSwitchRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSwitchRoutingInterface`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingInterfaceDhcp

> InlineResponse20040 GetDeviceSwitchRoutingInterfaceDhcp(ctx, serial, interfaceId).Execute()

Return a layer 3 interface DHCP configuration for a switch



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
    interfaceId := "interfaceId_example" // string | Interface ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSwitchRoutingInterfaceDhcp`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingInterfaces

> []InlineResponse20039 GetDeviceSwitchRoutingInterfaces(ctx, serial).Protocol(protocol).Execute()

List layer 3 interfaces for a switch



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
    protocol := "protocol_example" // string | Optional parameter to filter L3 interfaces by protocol. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetDeviceSwitchRoutingInterfaces(context.Background(), serial).Protocol(protocol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetDeviceSwitchRoutingInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSwitchRoutingInterfaces`: []InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetDeviceSwitchRoutingInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **protocol** | **string** | Optional parameter to filter L3 interfaces by protocol. | 

### Return type

[**[]InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingStaticRoute

> InlineResponse20041 GetDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).Execute()

Return a layer 3 static route for a switch



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
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetDeviceSwitchRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSwitchRoutingStaticRoute`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingStaticRoutes

> []InlineResponse20041 GetDeviceSwitchRoutingStaticRoutes(ctx, serial).Execute()

List layer 3 static routes for a switch



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
    resp, r, err := apiClient.RoutingApi.GetDeviceSwitchRoutingStaticRoutes(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetDeviceSwitchRoutingStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSwitchRoutingStaticRoutes`: []InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetDeviceSwitchRoutingStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticast

> InlineResponse200163 GetNetworkSwitchRoutingMulticast(ctx, networkId).Execute()

Return multicast settings for a network



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
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchRoutingMulticast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchRoutingMulticast`: InlineResponse200163
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchRoutingMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200163**](InlineResponse200163.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticastRendezvousPoint

> InlineResponse200164 GetNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Return a multicast rendezvous point



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
    rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchRoutingMulticastRendezvousPoint`: InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticastRendezvousPoints

> []InlineResponse200164 GetNetworkSwitchRoutingMulticastRendezvousPoints(ctx, networkId).Execute()

List multicast rendezvous points



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
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchRoutingMulticastRendezvousPoints`: []InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchRoutingMulticastRendezvousPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingOspf

> InlineResponse200165 GetNetworkSwitchRoutingOspf(ctx, networkId).Execute()

Return layer 3 OSPF routing configuration



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
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchRoutingOspf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchRoutingOspf`: InlineResponse200165
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200165**](InlineResponse200165.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterface

> InlineResponse20039 GetNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface from a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    interfaceId := "interfaceId_example" // string | Interface ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingInterface`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterfaceDhcp

> InlineResponse20040 GetNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface DHCP configuration for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    interfaceId := "interfaceId_example" // string | Interface ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingInterfaceDhcp`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterfaces

> []InlineResponse20039 GetNetworkSwitchStackRoutingInterfaces(ctx, networkId, switchStackId).Protocol(protocol).Execute()

List layer 3 interfaces for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    protocol := "protocol_example" // string | Optional parameter to filter L3 interfaces by protocol. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingInterfaces(context.Background(), networkId, switchStackId).Protocol(protocol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchStackRoutingInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingInterfaces`: []InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchStackRoutingInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **protocol** | **string** | Optional parameter to filter L3 interfaces by protocol. | 

### Return type

[**[]InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingStaticRoute

> InlineResponse20041 GetNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Return a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    staticRouteId := "staticRouteId_example" // string | Static route ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingStaticRoute`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingStaticRoutes

> []InlineResponse20041 GetNetworkSwitchStackRoutingStaticRoutes(ctx, networkId, switchStackId).Execute()

List layer 3 static routes for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetNetworkSwitchStackRoutingStaticRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStackRoutingStaticRoutes`: []InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetNetworkSwitchStackRoutingStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingInterface

> InlineResponse20039 UpdateDeviceSwitchRoutingInterface(ctx, serial, interfaceId).UpdateDeviceSwitchRoutingInterface(updateDeviceSwitchRoutingInterface).Execute()

Update a layer 3 interface for a switch



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
    interfaceId := "interfaceId_example" // string | Interface ID
    updateDeviceSwitchRoutingInterface := *openapiclient.NewInlineObject30() // InlineObject30 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).UpdateDeviceSwitchRoutingInterface(updateDeviceSwitchRoutingInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateDeviceSwitchRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceSwitchRoutingInterface`: InlineResponse20039
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchRoutingInterface** | [**InlineObject30**](InlineObject30.md) |  | 

### Return type

[**InlineResponse20039**](InlineResponse20039.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingInterfaceDhcp

> InlineResponse20040 UpdateDeviceSwitchRoutingInterfaceDhcp(ctx, serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcp(updateDeviceSwitchRoutingInterfaceDhcp).Execute()

Update a layer 3 interface DHCP configuration for a switch



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
    interfaceId := "interfaceId_example" // string | Interface ID
    updateDeviceSwitchRoutingInterfaceDhcp := *openapiclient.NewInlineObject31() // InlineObject31 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcp(updateDeviceSwitchRoutingInterfaceDhcp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceSwitchRoutingInterfaceDhcp`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchRoutingInterfaceDhcp** | [**InlineObject31**](InlineObject31.md) |  | 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingStaticRoute

> InlineResponse20041 UpdateDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).UpdateDeviceSwitchRoutingStaticRoute(updateDeviceSwitchRoutingStaticRoute).Execute()

Update a layer 3 static route for a switch



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
    staticRouteId := "staticRouteId_example" // string | Static route ID
    updateDeviceSwitchRoutingStaticRoute := *openapiclient.NewInlineObject33() // InlineObject33 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).UpdateDeviceSwitchRoutingStaticRoute(updateDeviceSwitchRoutingStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateDeviceSwitchRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceSwitchRoutingStaticRoute`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchRoutingStaticRoute** | [**InlineObject33**](InlineObject33.md) |  | 

### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingMulticast

> InlineResponse200163 UpdateNetworkSwitchRoutingMulticast(ctx, networkId).UpdateNetworkSwitchRoutingMulticast(updateNetworkSwitchRoutingMulticast).Execute()

Update multicast settings for a network



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
    updateNetworkSwitchRoutingMulticast := *openapiclient.NewInlineObject151() // InlineObject151 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateNetworkSwitchRoutingMulticast(context.Background(), networkId).UpdateNetworkSwitchRoutingMulticast(updateNetworkSwitchRoutingMulticast).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateNetworkSwitchRoutingMulticast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchRoutingMulticast`: InlineResponse200163
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateNetworkSwitchRoutingMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchRoutingMulticast** | [**InlineObject151**](InlineObject151.md) |  | 

### Return type

[**InlineResponse200163**](InlineResponse200163.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingMulticastRendezvousPoint

> InlineResponse200164 UpdateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).UpdateNetworkSwitchRoutingMulticastRendezvousPoint(updateNetworkSwitchRoutingMulticastRendezvousPoint).Execute()

Update a multicast rendezvous point



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
    rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID
    updateNetworkSwitchRoutingMulticastRendezvousPoint := *openapiclient.NewInlineObject153("InterfaceIp_example", "MulticastGroup_example") // InlineObject153 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).UpdateNetworkSwitchRoutingMulticastRendezvousPoint(updateNetworkSwitchRoutingMulticastRendezvousPoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: InlineResponse200164
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchRoutingMulticastRendezvousPoint** | [**InlineObject153**](InlineObject153.md) |  | 

### Return type

[**InlineResponse200164**](InlineResponse200164.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingOspf

> InlineResponse200165 UpdateNetworkSwitchRoutingOspf(ctx, networkId).UpdateNetworkSwitchRoutingOspf(updateNetworkSwitchRoutingOspf).Execute()

Update layer 3 OSPF routing configuration



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
    updateNetworkSwitchRoutingOspf := *openapiclient.NewInlineObject154() // InlineObject154 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateNetworkSwitchRoutingOspf(context.Background(), networkId).UpdateNetworkSwitchRoutingOspf(updateNetworkSwitchRoutingOspf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateNetworkSwitchRoutingOspf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchRoutingOspf`: InlineResponse200165
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchRoutingOspf** | [**InlineObject154**](InlineObject154.md) |  | 

### Return type

[**InlineResponse200165**](InlineResponse200165.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingInterface

> InlineResponse200168 UpdateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterface(updateNetworkSwitchStackRoutingInterface).Execute()

Update a layer 3 interface for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    interfaceId := "interfaceId_example" // string | Interface ID
    updateNetworkSwitchStackRoutingInterface := *openapiclient.NewInlineObject160() // InlineObject160 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterface(updateNetworkSwitchStackRoutingInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateNetworkSwitchStackRoutingInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStackRoutingInterface`: InlineResponse200168
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterface** | [**InlineObject160**](InlineObject160.md) |  | 

### Return type

[**InlineResponse200168**](InlineResponse200168.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingInterfaceDhcp

> InlineResponse20040 UpdateNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcp(updateNetworkSwitchStackRoutingInterfaceDhcp).Execute()

Update a layer 3 interface DHCP configuration for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    interfaceId := "interfaceId_example" // string | Interface ID
    updateNetworkSwitchStackRoutingInterfaceDhcp := *openapiclient.NewInlineObject161() // InlineObject161 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcp(updateNetworkSwitchStackRoutingInterfaceDhcp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStackRoutingInterfaceDhcp`: InlineResponse20040
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterfaceDhcp** | [**InlineObject161**](InlineObject161.md) |  | 

### Return type

[**InlineResponse20040**](InlineResponse20040.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingStaticRoute

> InlineResponse20041 UpdateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).UpdateNetworkSwitchStackRoutingStaticRoute(updateNetworkSwitchStackRoutingStaticRoute).Execute()

Update a layer 3 static route for a switch stack



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
    switchStackId := "switchStackId_example" // string | Switch stack ID
    staticRouteId := "staticRouteId_example" // string | Static route ID
    updateNetworkSwitchStackRoutingStaticRoute := *openapiclient.NewInlineObject163() // InlineObject163 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).UpdateNetworkSwitchStackRoutingStaticRoute(updateNetworkSwitchStackRoutingStaticRoute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.UpdateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStackRoutingStaticRoute`: InlineResponse20041
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.UpdateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingStaticRoute** | [**InlineObject163**](InlineObject163.md) |  | 

### Return type

[**InlineResponse20041**](InlineResponse20041.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

