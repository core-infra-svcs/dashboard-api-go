# \WirelessApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWirelessRfProfile**](WirelessApi.md#CreateNetworkWirelessRfProfile) | **Post** /networks/{networkId}/wireless/rfProfiles | Creates new RF profile for this network
[**CreateNetworkWirelessSsidIdentityPsk**](WirelessApi.md#CreateNetworkWirelessSsidIdentityPsk) | **Post** /networks/{networkId}/wireless/ssids/{number}/identityPsks | Create an Identity PSK
[**DeleteNetworkWirelessRfProfile**](WirelessApi.md#DeleteNetworkWirelessRfProfile) | **Delete** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Delete a RF Profile
[**DeleteNetworkWirelessSsidIdentityPsk**](WirelessApi.md#DeleteNetworkWirelessSsidIdentityPsk) | **Delete** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Delete an Identity PSK
[**GetDeviceWirelessBluetoothSettings**](WirelessApi.md#GetDeviceWirelessBluetoothSettings) | **Get** /devices/{serial}/wireless/bluetooth/settings | Return the bluetooth settings for a wireless device
[**GetDeviceWirelessConnectionStats**](WirelessApi.md#GetDeviceWirelessConnectionStats) | **Get** /devices/{serial}/wireless/connectionStats | Aggregated connectivity info for a given AP on this network
[**GetDeviceWirelessLatencyStats**](WirelessApi.md#GetDeviceWirelessLatencyStats) | **Get** /devices/{serial}/wireless/latencyStats | Aggregated latency info for a given AP on this network
[**GetDeviceWirelessRadioSettings**](WirelessApi.md#GetDeviceWirelessRadioSettings) | **Get** /devices/{serial}/wireless/radio/settings | Return the radio settings of a device
[**GetDeviceWirelessStatus**](WirelessApi.md#GetDeviceWirelessStatus) | **Get** /devices/{serial}/wireless/status | Return the SSID statuses of an access point
[**GetNetworkWirelessAirMarshal**](WirelessApi.md#GetNetworkWirelessAirMarshal) | **Get** /networks/{networkId}/wireless/airMarshal | List Air Marshal scan results from a network
[**GetNetworkWirelessAlternateManagementInterface**](WirelessApi.md#GetNetworkWirelessAlternateManagementInterface) | **Get** /networks/{networkId}/wireless/alternateManagementInterface | Return alternate management interface and devices with IP assigned
[**GetNetworkWirelessBilling**](WirelessApi.md#GetNetworkWirelessBilling) | **Get** /networks/{networkId}/wireless/billing | Return the billing settings of this network
[**GetNetworkWirelessBluetoothSettings**](WirelessApi.md#GetNetworkWirelessBluetoothSettings) | **Get** /networks/{networkId}/wireless/bluetooth/settings | Return the Bluetooth settings for a network. &lt;a href&#x3D;\&quot;https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\&quot;&gt;Bluetooth settings&lt;/a&gt; must be enabled on the network.
[**GetNetworkWirelessChannelUtilizationHistory**](WirelessApi.md#GetNetworkWirelessChannelUtilizationHistory) | **Get** /networks/{networkId}/wireless/channelUtilizationHistory | Return AP channel utilization over time for a device or network client
[**GetNetworkWirelessClientConnectionStats**](WirelessApi.md#GetNetworkWirelessClientConnectionStats) | **Get** /networks/{networkId}/wireless/clients/{clientId}/connectionStats | Aggregated connectivity info for a given client on this network
[**GetNetworkWirelessClientConnectivityEvents**](WirelessApi.md#GetNetworkWirelessClientConnectivityEvents) | **Get** /networks/{networkId}/wireless/clients/{clientId}/connectivityEvents | List the wireless connectivity events for a client within a network in the timespan.
[**GetNetworkWirelessClientCountHistory**](WirelessApi.md#GetNetworkWirelessClientCountHistory) | **Get** /networks/{networkId}/wireless/clientCountHistory | Return wireless client counts over time for a network, device, or network client
[**GetNetworkWirelessClientLatencyHistory**](WirelessApi.md#GetNetworkWirelessClientLatencyHistory) | **Get** /networks/{networkId}/wireless/clients/{clientId}/latencyHistory | Return the latency history for a client
[**GetNetworkWirelessClientLatencyStats**](WirelessApi.md#GetNetworkWirelessClientLatencyStats) | **Get** /networks/{networkId}/wireless/clients/{clientId}/latencyStats | Aggregated latency info for a given client on this network
[**GetNetworkWirelessClientsConnectionStats**](WirelessApi.md#GetNetworkWirelessClientsConnectionStats) | **Get** /networks/{networkId}/wireless/clients/connectionStats | Aggregated connectivity info for this network, grouped by clients
[**GetNetworkWirelessClientsLatencyStats**](WirelessApi.md#GetNetworkWirelessClientsLatencyStats) | **Get** /networks/{networkId}/wireless/clients/latencyStats | Aggregated latency info for this network, grouped by clients
[**GetNetworkWirelessConnectionStats**](WirelessApi.md#GetNetworkWirelessConnectionStats) | **Get** /networks/{networkId}/wireless/connectionStats | Aggregated connectivity info for this network
[**GetNetworkWirelessDataRateHistory**](WirelessApi.md#GetNetworkWirelessDataRateHistory) | **Get** /networks/{networkId}/wireless/dataRateHistory | Return PHY data rates over time for a network, device, or network client
[**GetNetworkWirelessDevicesConnectionStats**](WirelessApi.md#GetNetworkWirelessDevicesConnectionStats) | **Get** /networks/{networkId}/wireless/devices/connectionStats | Aggregated connectivity info for this network, grouped by node
[**GetNetworkWirelessDevicesLatencyStats**](WirelessApi.md#GetNetworkWirelessDevicesLatencyStats) | **Get** /networks/{networkId}/wireless/devices/latencyStats | Aggregated latency info for this network, grouped by node
[**GetNetworkWirelessFailedConnections**](WirelessApi.md#GetNetworkWirelessFailedConnections) | **Get** /networks/{networkId}/wireless/failedConnections | List of all failed client connection events on this network in a given time range
[**GetNetworkWirelessLatencyHistory**](WirelessApi.md#GetNetworkWirelessLatencyHistory) | **Get** /networks/{networkId}/wireless/latencyHistory | Return average wireless latency over time for a network, device, or network client
[**GetNetworkWirelessLatencyStats**](WirelessApi.md#GetNetworkWirelessLatencyStats) | **Get** /networks/{networkId}/wireless/latencyStats | Aggregated latency info for this network
[**GetNetworkWirelessMeshStatuses**](WirelessApi.md#GetNetworkWirelessMeshStatuses) | **Get** /networks/{networkId}/wireless/meshStatuses | List wireless mesh statuses for repeaters
[**GetNetworkWirelessRfProfile**](WirelessApi.md#GetNetworkWirelessRfProfile) | **Get** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkWirelessRfProfiles**](WirelessApi.md#GetNetworkWirelessRfProfiles) | **Get** /networks/{networkId}/wireless/rfProfiles | List RF profiles for this network
[**GetNetworkWirelessSettings**](WirelessApi.md#GetNetworkWirelessSettings) | **Get** /networks/{networkId}/wireless/settings | Return the wireless settings for a network
[**GetNetworkWirelessSignalQualityHistory**](WirelessApi.md#GetNetworkWirelessSignalQualityHistory) | **Get** /networks/{networkId}/wireless/signalQualityHistory | Return signal quality (SNR/RSSI) over time for a device or network client
[**GetNetworkWirelessSsid**](WirelessApi.md#GetNetworkWirelessSsid) | **Get** /networks/{networkId}/wireless/ssids/{number} | Return a single MR SSID
[**GetNetworkWirelessSsidBonjourForwarding**](WirelessApi.md#GetNetworkWirelessSsidBonjourForwarding) | **Get** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | List the Bonjour forwarding setting and rules for the SSID
[**GetNetworkWirelessSsidDeviceTypeGroupPolicies**](WirelessApi.md#GetNetworkWirelessSsidDeviceTypeGroupPolicies) | **Get** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | List the device type group policies for the SSID
[**GetNetworkWirelessSsidEapOverride**](WirelessApi.md#GetNetworkWirelessSsidEapOverride) | **Get** /networks/{networkId}/wireless/ssids/{number}/eapOverride | Return the EAP overridden parameters for an SSID
[**GetNetworkWirelessSsidFirewallL3FirewallRules**](WirelessApi.md#GetNetworkWirelessSsidFirewallL3FirewallRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules | Return the L3 firewall rules for an SSID on an MR network
[**GetNetworkWirelessSsidFirewallL7FirewallRules**](WirelessApi.md#GetNetworkWirelessSsidFirewallL7FirewallRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules | Return the L7 firewall rules for an SSID on an MR network
[**GetNetworkWirelessSsidHotspot20**](WirelessApi.md#GetNetworkWirelessSsidHotspot20) | **Get** /networks/{networkId}/wireless/ssids/{number}/hotspot20 | Return the Hotspot 2.0 settings for an SSID
[**GetNetworkWirelessSsidIdentityPsk**](WirelessApi.md#GetNetworkWirelessSsidIdentityPsk) | **Get** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Return an Identity PSK
[**GetNetworkWirelessSsidIdentityPsks**](WirelessApi.md#GetNetworkWirelessSsidIdentityPsks) | **Get** /networks/{networkId}/wireless/ssids/{number}/identityPsks | List all Identity PSKs in a wireless network
[**GetNetworkWirelessSsidSchedules**](WirelessApi.md#GetNetworkWirelessSsidSchedules) | **Get** /networks/{networkId}/wireless/ssids/{number}/schedules | List the outage schedule for the SSID
[**GetNetworkWirelessSsidSplashSettings**](WirelessApi.md#GetNetworkWirelessSsidSplashSettings) | **Get** /networks/{networkId}/wireless/ssids/{number}/splash/settings | Display the splash page settings for the given SSID
[**GetNetworkWirelessSsidTrafficShapingRules**](WirelessApi.md#GetNetworkWirelessSsidTrafficShapingRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Display the traffic shaping settings for a SSID on an MR network
[**GetNetworkWirelessSsidVpn**](WirelessApi.md#GetNetworkWirelessSsidVpn) | **Get** /networks/{networkId}/wireless/ssids/{number}/vpn | List the VPN settings for the SSID.
[**GetNetworkWirelessSsids**](WirelessApi.md#GetNetworkWirelessSsids) | **Get** /networks/{networkId}/wireless/ssids | List the MR SSIDs in a network
[**GetNetworkWirelessUsageHistory**](WirelessApi.md#GetNetworkWirelessUsageHistory) | **Get** /networks/{networkId}/wireless/usageHistory | Return AP usage over time for a device or network client
[**GetOrganizationWirelessDevicesChannelUtilizationByDevice**](WirelessApi.md#GetOrganizationWirelessDevicesChannelUtilizationByDevice) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/byDevice | Get average channel utilization for all bands in a network, split by AP
[**GetOrganizationWirelessDevicesChannelUtilizationByNetwork**](WirelessApi.md#GetOrganizationWirelessDevicesChannelUtilizationByNetwork) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork | Get average channel utilization across all bands for all networks in the organization
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval**](WirelessApi.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval | Get a time-series of average channel utilization for all bands, segmented by device.
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval**](WirelessApi.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval | Get a time-series of average channel utilization for all bands
[**GetOrganizationWirelessDevicesEthernetStatuses**](WirelessApi.md#GetOrganizationWirelessDevicesEthernetStatuses) | **Get** /organizations/{organizationId}/wireless/devices/ethernet/statuses | Endpoint to see power status for wireless devices
[**UpdateDeviceWirelessBluetoothSettings**](WirelessApi.md#UpdateDeviceWirelessBluetoothSettings) | **Put** /devices/{serial}/wireless/bluetooth/settings | Update the bluetooth settings for a wireless device
[**UpdateDeviceWirelessRadioSettings**](WirelessApi.md#UpdateDeviceWirelessRadioSettings) | **Put** /devices/{serial}/wireless/radio/settings | Update the radio settings of a device
[**UpdateNetworkWirelessAlternateManagementInterface**](WirelessApi.md#UpdateNetworkWirelessAlternateManagementInterface) | **Put** /networks/{networkId}/wireless/alternateManagementInterface | Update alternate management interface and device static IP
[**UpdateNetworkWirelessBilling**](WirelessApi.md#UpdateNetworkWirelessBilling) | **Put** /networks/{networkId}/wireless/billing | Update the billing settings
[**UpdateNetworkWirelessBluetoothSettings**](WirelessApi.md#UpdateNetworkWirelessBluetoothSettings) | **Put** /networks/{networkId}/wireless/bluetooth/settings | Update the Bluetooth settings for a network
[**UpdateNetworkWirelessRfProfile**](WirelessApi.md#UpdateNetworkWirelessRfProfile) | **Put** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Updates specified RF profile for this network
[**UpdateNetworkWirelessSettings**](WirelessApi.md#UpdateNetworkWirelessSettings) | **Put** /networks/{networkId}/wireless/settings | Update the wireless settings for a network
[**UpdateNetworkWirelessSsid**](WirelessApi.md#UpdateNetworkWirelessSsid) | **Put** /networks/{networkId}/wireless/ssids/{number} | Update the attributes of an MR SSID
[**UpdateNetworkWirelessSsidBonjourForwarding**](WirelessApi.md#UpdateNetworkWirelessSsidBonjourForwarding) | **Put** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | Update the bonjour forwarding setting and rules for the SSID
[**UpdateNetworkWirelessSsidDeviceTypeGroupPolicies**](WirelessApi.md#UpdateNetworkWirelessSsidDeviceTypeGroupPolicies) | **Put** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | Update the device type group policies for the SSID
[**UpdateNetworkWirelessSsidEapOverride**](WirelessApi.md#UpdateNetworkWirelessSsidEapOverride) | **Put** /networks/{networkId}/wireless/ssids/{number}/eapOverride | Update the EAP overridden parameters for an SSID.
[**UpdateNetworkWirelessSsidFirewallL3FirewallRules**](WirelessApi.md#UpdateNetworkWirelessSsidFirewallL3FirewallRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules | Update the L3 firewall rules of an SSID on an MR network
[**UpdateNetworkWirelessSsidFirewallL7FirewallRules**](WirelessApi.md#UpdateNetworkWirelessSsidFirewallL7FirewallRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules | Update the L7 firewall rules of an SSID on an MR network
[**UpdateNetworkWirelessSsidHotspot20**](WirelessApi.md#UpdateNetworkWirelessSsidHotspot20) | **Put** /networks/{networkId}/wireless/ssids/{number}/hotspot20 | Update the Hotspot 2.0 settings of an SSID
[**UpdateNetworkWirelessSsidIdentityPsk**](WirelessApi.md#UpdateNetworkWirelessSsidIdentityPsk) | **Put** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Update an Identity PSK
[**UpdateNetworkWirelessSsidSchedules**](WirelessApi.md#UpdateNetworkWirelessSsidSchedules) | **Put** /networks/{networkId}/wireless/ssids/{number}/schedules | Update the outage schedule for the SSID
[**UpdateNetworkWirelessSsidSplashSettings**](WirelessApi.md#UpdateNetworkWirelessSsidSplashSettings) | **Put** /networks/{networkId}/wireless/ssids/{number}/splash/settings | Modify the splash page settings for the given SSID
[**UpdateNetworkWirelessSsidTrafficShapingRules**](WirelessApi.md#UpdateNetworkWirelessSsidTrafficShapingRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Update the traffic shaping settings for an SSID on an MR network
[**UpdateNetworkWirelessSsidVpn**](WirelessApi.md#UpdateNetworkWirelessSsidVpn) | **Put** /networks/{networkId}/wireless/ssids/{number}/vpn | Update the VPN settings for the SSID



## CreateNetworkWirelessRfProfile

> InlineResponse20092 CreateNetworkWirelessRfProfile(ctx, networkId).CreateNetworkWirelessRfProfile(createNetworkWirelessRfProfile).Execute()

Creates new RF profile for this network



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
    createNetworkWirelessRfProfile := *openapiclient.NewInlineObject154("Name_example", "BandSelectionType_example") // InlineObject154 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.CreateNetworkWirelessRfProfile(context.Background(), networkId).CreateNetworkWirelessRfProfile(createNetworkWirelessRfProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.CreateNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessRfProfile`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.CreateNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWirelessRfProfile** | [**InlineObject154**](InlineObject154.md) |  | 

### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWirelessSsidIdentityPsk

> map[string]interface{} CreateNetworkWirelessSsidIdentityPsk(ctx, networkId, number).CreateNetworkWirelessSsidIdentityPsk(createNetworkWirelessSsidIdentityPsk).Execute()

Create an Identity PSK



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
    number := "number_example" // string | Number
    createNetworkWirelessSsidIdentityPsk := *openapiclient.NewInlineObject164("Name_example", "GroupPolicyId_example") // InlineObject164 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.CreateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number).CreateNetworkWirelessSsidIdentityPsk(createNetworkWirelessSsidIdentityPsk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.CreateNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessSsidIdentityPsk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.CreateNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkWirelessSsidIdentityPsk** | [**InlineObject164**](InlineObject164.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWirelessRfProfile

> DeleteNetworkWirelessRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.DeleteNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.DeleteNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWirelessSsidIdentityPsk

> DeleteNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).Execute()

Delete an Identity PSK



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
    number := "number_example" // string | Number
    identityPskId := "identityPskId_example" // string | Identity psk ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.DeleteNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.DeleteNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 
**identityPskId** | **string** | Identity psk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessBluetoothSettings

> InlineResponse20010 GetDeviceWirelessBluetoothSettings(ctx, serial).Execute()

Return the bluetooth settings for a wireless device



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
    resp, r, err := apiClient.WirelessApi.GetDeviceWirelessBluetoothSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetDeviceWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessBluetoothSettings`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessConnectionStats

> InlineResponse20011 GetDeviceWirelessConnectionStats(ctx, serial).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for a given AP on this network



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetDeviceWirelessConnectionStats(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetDeviceWirelessConnectionStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessConnectionStats`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetDeviceWirelessConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessConnectionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessLatencyStats

> map[string]interface{} GetDeviceWirelessLatencyStats(ctx, serial).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for a given AP on this network



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    fields := "fields_example" // string | Partial selection: If present, this call will return only the selected fields of [\"rawDistribution\", \"avg\"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetDeviceWirelessLatencyStats(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetDeviceWirelessLatencyStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessLatencyStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetDeviceWirelessLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessLatencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **fields** | **string** | Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string. | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessRadioSettings

> map[string]interface{} GetDeviceWirelessRadioSettings(ctx, serial).Execute()

Return the radio settings of a device



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
    resp, r, err := apiClient.WirelessApi.GetDeviceWirelessRadioSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetDeviceWirelessRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessRadioSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetDeviceWirelessRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessStatus

> map[string]interface{} GetDeviceWirelessStatus(ctx, serial).Execute()

Return the SSID statuses of an access point



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
    resp, r, err := apiClient.WirelessApi.GetDeviceWirelessStatus(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetDeviceWirelessStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetDeviceWirelessStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessAirMarshal

> []map[string]interface{} GetNetworkWirelessAirMarshal(ctx, networkId).T0(t0).Timespan(timespan).Execute()

List Air Marshal scan results from a network



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
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessAirMarshal(context.Background(), networkId).T0(t0).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessAirMarshal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessAirMarshal`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessAirMarshal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessAirMarshalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessAlternateManagementInterface

> map[string]interface{} GetNetworkWirelessAlternateManagementInterface(ctx, networkId).Execute()

Return alternate management interface and devices with IP assigned



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
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessAlternateManagementInterface(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessAlternateManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessAlternateManagementInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessBilling

> map[string]interface{} GetNetworkWirelessBilling(ctx, networkId).Execute()

Return the billing settings of this network



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
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessBilling(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessBilling`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessBluetoothSettings

> InlineResponse20085 GetNetworkWirelessBluetoothSettings(ctx, networkId).Execute()

Return the Bluetooth settings for a network. <a href=\"https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\">Bluetooth settings</a> must be enabled on the network.



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
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessBluetoothSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessBluetoothSettings`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessChannelUtilizationHistory

> []InlineResponse20086 GetNetworkWirelessChannelUtilizationHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Execute()

Return AP channel utilization over time for a device or network client



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
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)
    autoResolution := true // bool | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. (optional)
    clientId := "clientId_example" // string | Filter results by network client to return per-device, per-band AP channel utilization metrics inner joined by the queried client's connection history. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by device to return AP channel utilization metrics for the queried device; either :band or :clientId must be jointly specified. (optional)
    apTag := "apTag_example" // string | Filter results by AP tag to return AP channel utilization metrics for devices labeled with the given tag; either :clientId or :deviceSerial must be jointly specified. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessChannelUtilizationHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessChannelUtilizationHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessChannelUtilizationHistory`: []InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessChannelUtilizationHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessChannelUtilizationHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 600, 1200, 3600, 14400, 86400. The default is 86400. | 
 **autoResolution** | **bool** | Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false. | 
 **clientId** | **string** | Filter results by network client to return per-device, per-band AP channel utilization metrics inner joined by the queried client&#39;s connection history. | 
 **deviceSerial** | **string** | Filter results by device to return AP channel utilization metrics for the queried device; either :band or :clientId must be jointly specified. | 
 **apTag** | **string** | Filter results by AP tag to return AP channel utilization metrics for devices labeled with the given tag; either :clientId or :deviceSerial must be jointly specified. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). | 

### Return type

[**[]InlineResponse20086**](InlineResponse20086.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientConnectionStats

> map[string]interface{} GetNetworkWirelessClientConnectionStats(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for a given client on this network



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
    clientId := "clientId_example" // string | Client ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessClientConnectionStats(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessClientConnectionStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessClientConnectionStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessClientConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientConnectionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientConnectivityEvents

> []map[string]interface{} GetNetworkWirelessClientConnectivityEvents(ctx, networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Types(types).IncludedSeverities(includedSeverities).Band(band).SsidNumber(ssidNumber).DeviceSerial(deviceSerial).Execute()

List the wireless connectivity events for a client within a network in the timespan.



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
    clientId := "clientId_example" // string | Client ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    types := []string{"Types_example"} // []string | A list of event types to include. If not specified, events of all types will be returned. Valid types are 'assoc', 'disassoc', 'auth', 'deauth', 'dns', 'dhcp', 'roam', 'connection' and/or 'sticky'. (optional)
    includedSeverities := []string{"IncludedSeverities_example"} // []string | A list of severities to include. If not specified, events of all severities will be returned. Valid severities are 'good', 'info', 'warn' and/or 'bad'. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5', '6'). (optional)
    ssidNumber := int32(56) // int32 | An SSID number to include. If not specified, events for all SSIDs will be returned. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by an AP's serial number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessClientConnectivityEvents(context.Background(), networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Types(types).IncludedSeverities(includedSeverities).Band(band).SsidNumber(ssidNumber).DeviceSerial(deviceSerial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessClientConnectivityEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessClientConnectivityEvents`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessClientConnectivityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientConnectivityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **types** | **[]string** | A list of event types to include. If not specified, events of all types will be returned. Valid types are &#39;assoc&#39;, &#39;disassoc&#39;, &#39;auth&#39;, &#39;deauth&#39;, &#39;dns&#39;, &#39;dhcp&#39;, &#39;roam&#39;, &#39;connection&#39; and/or &#39;sticky&#39;. | 
 **includedSeverities** | **[]string** | A list of severities to include. If not specified, events of all severities will be returned. Valid severities are &#39;good&#39;, &#39;info&#39;, &#39;warn&#39; and/or &#39;bad&#39;. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39;, &#39;6&#39;). | 
 **ssidNumber** | **int32** | An SSID number to include. If not specified, events for all SSIDs will be returned. | 
 **deviceSerial** | **string** | Filter results by an AP&#39;s serial number. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientCountHistory

> []InlineResponse20087 GetNetworkWirelessClientCountHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()

Return wireless client counts over time for a network, device, or network client



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
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)
    autoResolution := true // bool | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. (optional)
    clientId := "clientId_example" // string | Filter results by network client to return per-device client counts over time inner joined by the queried client's connection history. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by device. (optional)
    apTag := "apTag_example" // string | Filter results by AP tag. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). (optional)
    ssid := int32(56) // int32 | Filter results by SSID number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessClientCountHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessClientCountHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessClientCountHistory`: []InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessClientCountHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientCountHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. | 
 **autoResolution** | **bool** | Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false. | 
 **clientId** | **string** | Filter results by network client to return per-device client counts over time inner joined by the queried client&#39;s connection history. | 
 **deviceSerial** | **string** | Filter results by device. | 
 **apTag** | **string** | Filter results by AP tag. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). | 
 **ssid** | **int32** | Filter results by SSID number. | 

### Return type

[**[]InlineResponse20087**](InlineResponse20087.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientLatencyHistory

> []map[string]interface{} GetNetworkWirelessClientLatencyHistory(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Return the latency history for a client



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
    clientId := "clientId_example" // string | Client ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 791 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 1 day. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 86400. The default is 86400. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessClientLatencyHistory(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessClientLatencyHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessClientLatencyHistory`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessClientLatencyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientLatencyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 791 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 86400. The default is 86400. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientLatencyStats

> map[string]interface{} GetNetworkWirelessClientLatencyStats(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for a given client on this network



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
    clientId := "clientId_example" // string | Client ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    fields := "fields_example" // string | Partial selection: If present, this call will return only the selected fields of [\"rawDistribution\", \"avg\"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessClientLatencyStats(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessClientLatencyStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessClientLatencyStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessClientLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientLatencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **fields** | **string** | Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string. | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientsConnectionStats

> []map[string]interface{} GetNetworkWirelessClientsConnectionStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for this network, grouped by clients



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessClientsConnectionStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessClientsConnectionStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessClientsConnectionStats`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessClientsConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientsConnectionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientsLatencyStats

> []map[string]interface{} GetNetworkWirelessClientsLatencyStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for this network, grouped by clients



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    fields := "fields_example" // string | Partial selection: If present, this call will return only the selected fields of [\"rawDistribution\", \"avg\"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessClientsLatencyStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessClientsLatencyStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessClientsLatencyStats`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessClientsLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientsLatencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **fields** | **string** | Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessConnectionStats

> InlineResponse20088 GetNetworkWirelessConnectionStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for this network



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessConnectionStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessConnectionStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessConnectionStats`: InlineResponse20088
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessConnectionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 

### Return type

[**InlineResponse20088**](InlineResponse20088.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessDataRateHistory

> []InlineResponse20089 GetNetworkWirelessDataRateHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()

Return PHY data rates over time for a network, device, or network client



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
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)
    autoResolution := true // bool | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. (optional)
    clientId := "clientId_example" // string | Filter results by network client. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by device. (optional)
    apTag := "apTag_example" // string | Filter results by AP tag. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). (optional)
    ssid := int32(56) // int32 | Filter results by SSID number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessDataRateHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessDataRateHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessDataRateHistory`: []InlineResponse20089
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessDataRateHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessDataRateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. | 
 **autoResolution** | **bool** | Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false. | 
 **clientId** | **string** | Filter results by network client. | 
 **deviceSerial** | **string** | Filter results by device. | 
 **apTag** | **string** | Filter results by AP tag. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). | 
 **ssid** | **int32** | Filter results by SSID number. | 

### Return type

[**[]InlineResponse20089**](InlineResponse20089.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessDevicesConnectionStats

> []InlineResponse20011 GetNetworkWirelessDevicesConnectionStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for this network, grouped by node



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessDevicesConnectionStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessDevicesConnectionStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessDevicesConnectionStats`: []InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessDevicesConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessDevicesConnectionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 

### Return type

[**[]InlineResponse20011**](InlineResponse20011.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessDevicesLatencyStats

> []map[string]interface{} GetNetworkWirelessDevicesLatencyStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for this network, grouped by node



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    fields := "fields_example" // string | Partial selection: If present, this call will return only the selected fields of [\"rawDistribution\", \"avg\"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessDevicesLatencyStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessDevicesLatencyStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessDevicesLatencyStats`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessDevicesLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessDevicesLatencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **fields** | **string** | Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessFailedConnections

> []InlineResponse20090 GetNetworkWirelessFailedConnections(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Serial(serial).ClientId(clientId).Execute()

List of all failed client connection events on this network in a given time range



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    serial := "serial_example" // string | Filter by AP (optional)
    clientId := "clientId_example" // string | Filter by client MAC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessFailedConnections(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Serial(serial).ClientId(clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessFailedConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessFailedConnections`: []InlineResponse20090
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessFailedConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessFailedConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **serial** | **string** | Filter by AP | 
 **clientId** | **string** | Filter by client MAC | 

### Return type

[**[]InlineResponse20090**](InlineResponse20090.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessLatencyHistory

> []InlineResponse20091 GetNetworkWirelessLatencyHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).AccessCategory(accessCategory).Execute()

Return average wireless latency over time for a network, device, or network client



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
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)
    autoResolution := true // bool | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. (optional)
    clientId := "clientId_example" // string | Filter results by network client. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by device. (optional)
    apTag := "apTag_example" // string | Filter results by AP tag. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). (optional)
    ssid := int32(56) // int32 | Filter results by SSID number. (optional)
    accessCategory := "accessCategory_example" // string | Filter by access category. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessLatencyHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).AccessCategory(accessCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessLatencyHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessLatencyHistory`: []InlineResponse20091
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessLatencyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessLatencyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. | 
 **autoResolution** | **bool** | Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false. | 
 **clientId** | **string** | Filter results by network client. | 
 **deviceSerial** | **string** | Filter results by device. | 
 **apTag** | **string** | Filter results by AP tag. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). | 
 **ssid** | **int32** | Filter results by SSID number. | 
 **accessCategory** | **string** | Filter by access category. | 

### Return type

[**[]InlineResponse20091**](InlineResponse20091.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessLatencyStats

> map[string]interface{} GetNetworkWirelessLatencyStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for this network



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    fields := "fields_example" // string | Partial selection: If present, this call will return only the selected fields of [\"rawDistribution\", \"avg\"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessLatencyStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessLatencyStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessLatencyStats`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessLatencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **fields** | **string** | Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string. | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessMeshStatuses

> map[string]interface{} GetNetworkWirelessMeshStatuses(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List wireless mesh statuses for repeaters



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessMeshStatuses(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessMeshStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessMeshStatuses`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessMeshStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessMeshStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 500. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessRfProfile

> InlineResponse20092 GetNetworkWirelessRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessRfProfile`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessRfProfiles

> InlineResponse20092 GetNetworkWirelessRfProfiles(ctx, networkId).IncludeTemplateProfiles(includeTemplateProfiles).Execute()

List RF profiles for this network



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
    includeTemplateProfiles := true // bool | If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessRfProfiles(context.Background(), networkId).IncludeTemplateProfiles(includeTemplateProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessRfProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessRfProfiles`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTemplateProfiles** | **bool** | If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false. | 

### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSettings

> InlineResponse20093 GetNetworkWirelessSettings(ctx, networkId).Execute()

Return the wireless settings for a network



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
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSettings`: InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20093**](InlineResponse20093.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSignalQualityHistory

> []InlineResponse20094 GetNetworkWirelessSignalQualityHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()

Return signal quality (SNR/RSSI) over time for a device or network client



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
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)
    autoResolution := true // bool | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. (optional)
    clientId := "clientId_example" // string | Filter results by network client. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by device. (optional)
    apTag := "apTag_example" // string | Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). (optional)
    ssid := int32(56) // int32 | Filter results by SSID number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSignalQualityHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSignalQualityHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSignalQualityHistory`: []InlineResponse20094
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSignalQualityHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSignalQualityHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. | 
 **autoResolution** | **bool** | Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false. | 
 **clientId** | **string** | Filter results by network client. | 
 **deviceSerial** | **string** | Filter results by device. | 
 **apTag** | **string** | Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). | 
 **ssid** | **int32** | Filter results by SSID number. | 

### Return type

[**[]InlineResponse20094**](InlineResponse20094.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsid

> map[string]interface{} GetNetworkWirelessSsid(ctx, networkId, number).Execute()

Return a single MR SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsid(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsid`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidBonjourForwarding

> map[string]interface{} GetNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).Execute()

List the Bonjour forwarding setting and rules for the SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidBonjourForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidBonjourForwarding`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidDeviceTypeGroupPolicies

> map[string]interface{} GetNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).Execute()

List the device type group policies for the SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidDeviceTypeGroupPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidEapOverride

> InlineResponse20095 GetNetworkWirelessSsidEapOverride(ctx, networkId, number).Execute()

Return the EAP overridden parameters for an SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidEapOverride(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidEapOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidEapOverride`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidEapOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidEapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidFirewallL3FirewallRules

> map[string]interface{} GetNetworkWirelessSsidFirewallL3FirewallRules(ctx, networkId, number).Execute()

Return the L3 firewall rules for an SSID on an MR network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidFirewallL3FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidFirewallL7FirewallRules

> map[string]interface{} GetNetworkWirelessSsidFirewallL7FirewallRules(ctx, networkId, number).Execute()

Return the L7 firewall rules for an SSID on an MR network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidFirewallL7FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidHotspot20

> map[string]interface{} GetNetworkWirelessSsidHotspot20(ctx, networkId, number).Execute()

Return the Hotspot 2.0 settings for an SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidHotspot20(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidHotspot20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidHotspot20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidHotspot20`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidHotspot20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidIdentityPsk

> InlineResponse20096 GetNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).Execute()

Return an Identity PSK



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
    number := "number_example" // string | Number
    identityPskId := "identityPskId_example" // string | Identity psk ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidIdentityPsk`: InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 
**identityPskId** | **string** | Identity psk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20096**](InlineResponse20096.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidIdentityPsks

> []InlineResponse20096 GetNetworkWirelessSsidIdentityPsks(ctx, networkId, number).Execute()

List all Identity PSKs in a wireless network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidIdentityPsks(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidIdentityPsks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidIdentityPsks`: []InlineResponse20096
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidIdentityPsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidIdentityPsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20096**](InlineResponse20096.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidSchedules

> map[string]interface{} GetNetworkWirelessSsidSchedules(ctx, networkId, number).Execute()

List the outage schedule for the SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidSchedules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidSplashSettings

> InlineResponse20097 GetNetworkWirelessSsidSplashSettings(ctx, networkId, number).Execute()

Display the splash page settings for the given SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidSplashSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidSplashSettings`: InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidSplashSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidSplashSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20097**](InlineResponse20097.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidTrafficShapingRules

> map[string]interface{} GetNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).Execute()

Display the traffic shaping settings for a SSID on an MR network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidTrafficShapingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidVpn

> map[string]interface{} GetNetworkWirelessSsidVpn(ctx, networkId, number).Execute()

List the VPN settings for the SSID.



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsidVpn(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsidVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidVpn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsidVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsids

> []map[string]interface{} GetNetworkWirelessSsids(ctx, networkId).Execute()

List the MR SSIDs in a network



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
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessSsids(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessSsids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsids`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessUsageHistory

> []InlineResponse20098 GetNetworkWirelessUsageHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()

Return AP usage over time for a device or network client



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
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)
    autoResolution := true // bool | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. (optional)
    clientId := "clientId_example" // string | Filter results by network client to return per-device AP usage over time inner joined by the queried client's connection history. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by device. Requires :band. (optional)
    apTag := "apTag_example" // string | Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). (optional)
    ssid := int32(56) // int32 | Filter results by SSID number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetNetworkWirelessUsageHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetNetworkWirelessUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessUsageHistory`: []InlineResponse20098
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetNetworkWirelessUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. | 
 **autoResolution** | **bool** | Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false. | 
 **clientId** | **string** | Filter results by network client to return per-device AP usage over time inner joined by the queried client&#39;s connection history. | 
 **deviceSerial** | **string** | Filter results by device. Requires :band. | 
 **apTag** | **string** | Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). | 
 **ssid** | **int32** | Filter results by SSID number. | 

### Return type

[**[]InlineResponse20098**](InlineResponse20098.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationByDevice

> []InlineResponse200155 GetOrganizationWirelessDevicesChannelUtilizationByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get average channel utilization for all bands in a network, split by AP



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationByDevice`: []InlineResponse200155
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200155**](InlineResponse200155.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationByNetwork

> []InlineResponse200156 GetOrganizationWirelessDevicesChannelUtilizationByNetwork(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get average channel utilization across all bands for all networks in the organization



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationByNetwork(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationByNetwork`: []InlineResponse200156
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200156**](InlineResponse200156.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval

> []InlineResponse200157 GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands, segmented by device.



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: []InlineResponse200157
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200157**](InlineResponse200157.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval

> []InlineResponse200158 GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands



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
    networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
    serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: []InlineResponse200158
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]InlineResponse200158**](InlineResponse200158.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesEthernetStatuses

> []InlineResponse200159 GetOrganizationWirelessDevicesEthernetStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Endpoint to see power status for wireless devices



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.GetOrganizationWirelessDevicesEthernetStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.GetOrganizationWirelessDevicesEthernetStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesEthernetStatuses`: []InlineResponse200159
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.GetOrganizationWirelessDevicesEthernetStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesEthernetStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 

### Return type

[**[]InlineResponse200159**](InlineResponse200159.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessBluetoothSettings

> InlineResponse20010 UpdateDeviceWirelessBluetoothSettings(ctx, serial).UpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings).Execute()

Update the bluetooth settings for a wireless device



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
    updateDeviceWirelessBluetoothSettings := *openapiclient.NewInlineObject25() // InlineObject25 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateDeviceWirelessBluetoothSettings(context.Background(), serial).UpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateDeviceWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessBluetoothSettings`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessBluetoothSettings** | [**InlineObject25**](InlineObject25.md) |  | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessRadioSettings

> map[string]interface{} UpdateDeviceWirelessRadioSettings(ctx, serial).UpdateDeviceWirelessRadioSettings(updateDeviceWirelessRadioSettings).Execute()

Update the radio settings of a device



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
    updateDeviceWirelessRadioSettings := *openapiclient.NewInlineObject26() // InlineObject26 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateDeviceWirelessRadioSettings(context.Background(), serial).UpdateDeviceWirelessRadioSettings(updateDeviceWirelessRadioSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateDeviceWirelessRadioSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessRadioSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateDeviceWirelessRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessRadioSettings** | [**InlineObject26**](InlineObject26.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessAlternateManagementInterface

> map[string]interface{} UpdateNetworkWirelessAlternateManagementInterface(ctx, networkId).UpdateNetworkWirelessAlternateManagementInterface(updateNetworkWirelessAlternateManagementInterface).Execute()

Update alternate management interface and device static IP



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
    updateNetworkWirelessAlternateManagementInterface := *openapiclient.NewInlineObject151() // InlineObject151 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessAlternateManagementInterface(context.Background(), networkId).UpdateNetworkWirelessAlternateManagementInterface(updateNetworkWirelessAlternateManagementInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessAlternateManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessAlternateManagementInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessAlternateManagementInterface** | [**InlineObject151**](InlineObject151.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessBilling

> map[string]interface{} UpdateNetworkWirelessBilling(ctx, networkId).UpdateNetworkWirelessBilling(updateNetworkWirelessBilling).Execute()

Update the billing settings



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
    updateNetworkWirelessBilling := *openapiclient.NewInlineObject152() // InlineObject152 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessBilling(context.Background(), networkId).UpdateNetworkWirelessBilling(updateNetworkWirelessBilling).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessBilling`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessBilling** | [**InlineObject152**](InlineObject152.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessBluetoothSettings

> InlineResponse20085 UpdateNetworkWirelessBluetoothSettings(ctx, networkId).UpdateNetworkWirelessBluetoothSettings(updateNetworkWirelessBluetoothSettings).Execute()

Update the Bluetooth settings for a network



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
    updateNetworkWirelessBluetoothSettings := *openapiclient.NewInlineObject153() // InlineObject153 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessBluetoothSettings(context.Background(), networkId).UpdateNetworkWirelessBluetoothSettings(updateNetworkWirelessBluetoothSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessBluetoothSettings`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessBluetoothSettings** | [**InlineObject153**](InlineObject153.md) |  | 

### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessRfProfile

> InlineResponse20092 UpdateNetworkWirelessRfProfile(ctx, networkId, rfProfileId).UpdateNetworkWirelessRfProfile(updateNetworkWirelessRfProfile).Execute()

Updates specified RF profile for this network



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID
    updateNetworkWirelessRfProfile := *openapiclient.NewInlineObject155() // InlineObject155 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkWirelessRfProfile(updateNetworkWirelessRfProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessRfProfile`: InlineResponse20092
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessRfProfile** | [**InlineObject155**](InlineObject155.md) |  | 

### Return type

[**InlineResponse20092**](InlineResponse20092.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSettings

> InlineResponse20093 UpdateNetworkWirelessSettings(ctx, networkId).UpdateNetworkWirelessSettings(updateNetworkWirelessSettings).Execute()

Update the wireless settings for a network



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
    updateNetworkWirelessSettings := *openapiclient.NewInlineObject156() // InlineObject156 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSettings(context.Background(), networkId).UpdateNetworkWirelessSettings(updateNetworkWirelessSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSettings`: InlineResponse20093
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessSettings** | [**InlineObject156**](InlineObject156.md) |  | 

### Return type

[**InlineResponse20093**](InlineResponse20093.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsid

> map[string]interface{} UpdateNetworkWirelessSsid(ctx, networkId, number).UpdateNetworkWirelessSsid(updateNetworkWirelessSsid).Execute()

Update the attributes of an MR SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsid := *openapiclient.NewInlineObject157() // InlineObject157 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsid(context.Background(), networkId, number).UpdateNetworkWirelessSsid(updateNetworkWirelessSsid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsid`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsid** | [**InlineObject157**](InlineObject157.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidBonjourForwarding

> map[string]interface{} UpdateNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).UpdateNetworkWirelessSsidBonjourForwarding(updateNetworkWirelessSsidBonjourForwarding).Execute()

Update the bonjour forwarding setting and rules for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidBonjourForwarding := *openapiclient.NewInlineObject158() // InlineObject158 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).UpdateNetworkWirelessSsidBonjourForwarding(updateNetworkWirelessSsidBonjourForwarding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidBonjourForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidBonjourForwarding`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidBonjourForwarding** | [**InlineObject158**](InlineObject158.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidDeviceTypeGroupPolicies

> map[string]interface{} UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()

Update the device type group policies for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidDeviceTypeGroupPolicies := *openapiclient.NewInlineObject159() // InlineObject159 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidDeviceTypeGroupPolicies** | [**InlineObject159**](InlineObject159.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidEapOverride

> InlineResponse20095 UpdateNetworkWirelessSsidEapOverride(ctx, networkId, number).UpdateNetworkWirelessSsidEapOverride(updateNetworkWirelessSsidEapOverride).Execute()

Update the EAP overridden parameters for an SSID.



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidEapOverride := *openapiclient.NewInlineObject160() // InlineObject160 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidEapOverride(context.Background(), networkId, number).UpdateNetworkWirelessSsidEapOverride(updateNetworkWirelessSsidEapOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidEapOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidEapOverride`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidEapOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidEapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidEapOverride** | [**InlineObject160**](InlineObject160.md) |  | 

### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidFirewallL3FirewallRules

> map[string]interface{} UpdateNetworkWirelessSsidFirewallL3FirewallRules(ctx, networkId, number).UpdateNetworkWirelessSsidFirewallL3FirewallRules(updateNetworkWirelessSsidFirewallL3FirewallRules).Execute()

Update the L3 firewall rules of an SSID on an MR network



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidFirewallL3FirewallRules := *openapiclient.NewInlineObject161() // InlineObject161 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidFirewallL3FirewallRules(updateNetworkWirelessSsidFirewallL3FirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidFirewallL3FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidFirewallL3FirewallRules** | [**InlineObject161**](InlineObject161.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidFirewallL7FirewallRules

> map[string]interface{} UpdateNetworkWirelessSsidFirewallL7FirewallRules(ctx, networkId, number).UpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules).Execute()

Update the L7 firewall rules of an SSID on an MR network



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidFirewallL7FirewallRules := *openapiclient.NewInlineObject162() // InlineObject162 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidFirewallL7FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidFirewallL7FirewallRules** | [**InlineObject162**](InlineObject162.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidHotspot20

> map[string]interface{} UpdateNetworkWirelessSsidHotspot20(ctx, networkId, number).UpdateNetworkWirelessSsidHotspot20(updateNetworkWirelessSsidHotspot20).Execute()

Update the Hotspot 2.0 settings of an SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidHotspot20 := *openapiclient.NewInlineObject163() // InlineObject163 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidHotspot20(context.Background(), networkId, number).UpdateNetworkWirelessSsidHotspot20(updateNetworkWirelessSsidHotspot20).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidHotspot20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidHotspot20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidHotspot20`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidHotspot20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidHotspot20** | [**InlineObject163**](InlineObject163.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidIdentityPsk

> map[string]interface{} UpdateNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).UpdateNetworkWirelessSsidIdentityPsk(updateNetworkWirelessSsidIdentityPsk).Execute()

Update an Identity PSK



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
    number := "number_example" // string | Number
    identityPskId := "identityPskId_example" // string | Identity psk ID
    updateNetworkWirelessSsidIdentityPsk := *openapiclient.NewInlineObject165() // InlineObject165 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).UpdateNetworkWirelessSsidIdentityPsk(updateNetworkWirelessSsidIdentityPsk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidIdentityPsk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 
**identityPskId** | **string** | Identity psk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkWirelessSsidIdentityPsk** | [**InlineObject165**](InlineObject165.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidSchedules

> map[string]interface{} UpdateNetworkWirelessSsidSchedules(ctx, networkId, number).UpdateNetworkWirelessSsidSchedules(updateNetworkWirelessSsidSchedules).Execute()

Update the outage schedule for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidSchedules := *openapiclient.NewInlineObject166() // InlineObject166 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidSchedules(context.Background(), networkId, number).UpdateNetworkWirelessSsidSchedules(updateNetworkWirelessSsidSchedules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidSchedules** | [**InlineObject166**](InlineObject166.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidSplashSettings

> InlineResponse20097 UpdateNetworkWirelessSsidSplashSettings(ctx, networkId, number).UpdateNetworkWirelessSsidSplashSettings(updateNetworkWirelessSsidSplashSettings).Execute()

Modify the splash page settings for the given SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidSplashSettings := *openapiclient.NewInlineObject167() // InlineObject167 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).UpdateNetworkWirelessSsidSplashSettings(updateNetworkWirelessSsidSplashSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidSplashSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidSplashSettings`: InlineResponse20097
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidSplashSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidSplashSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidSplashSettings** | [**InlineObject167**](InlineObject167.md) |  | 

### Return type

[**InlineResponse20097**](InlineResponse20097.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidTrafficShapingRules

> map[string]interface{} UpdateNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()

Update the traffic shaping settings for an SSID on an MR network



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidTrafficShapingRules := *openapiclient.NewInlineObject168() // InlineObject168 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidTrafficShapingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidTrafficShapingRules** | [**InlineObject168**](InlineObject168.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidVpn

> map[string]interface{} UpdateNetworkWirelessSsidVpn(ctx, networkId, number).UpdateNetworkWirelessSsidVpn(updateNetworkWirelessSsidVpn).Execute()

Update the VPN settings for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidVpn := *openapiclient.NewInlineObject169() // InlineObject169 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessApi.UpdateNetworkWirelessSsidVpn(context.Background(), networkId, number).UpdateNetworkWirelessSsidVpn(updateNetworkWirelessSsidVpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessApi.UpdateNetworkWirelessSsidVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidVpn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessApi.UpdateNetworkWirelessSsidVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidVpn** | [**InlineObject169**](InlineObject169.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

