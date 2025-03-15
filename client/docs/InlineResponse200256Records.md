# InlineResponse200256Records

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200256Network**](InlineResponse200256Network.md) |  | [optional] 
**Ip** | Pointer to **string** | The IP address of the client | [optional] 
**Ip6** | Pointer to **string** | The IPv6 address of the client | [optional] 
**Description** | Pointer to **string** | Short description of the client | [optional] 
**FirstSeen** | Pointer to **int32** | Timestamp client was first seen in the network | [optional] 
**LastSeen** | Pointer to **int32** | Timestamp client was last seen in the network | [optional] 
**Os** | Pointer to **string** | The operating system of the client | [optional] 
**User** | Pointer to **string** | The username of the user of the client | [optional] 
**Vlan** | Pointer to **string** | The name of the VLAN that the client is connected to | [optional] 
**Ssid** | Pointer to **string** | The name of the SSID that the client is connected to | [optional] 
**Switchport** | Pointer to **string** | The switch port the client is connected to | [optional] 
**WirelessCapabilities** | Pointer to **string** | Wireless capabilities of the client | [optional] 
**SmInstalled** | Pointer to **bool** | Status of SM for the client | [optional] 
**RecentDeviceMac** | Pointer to **string** | The MAC address of the node that the device was last connected to | [optional] 
**ClientVpnConnections** | Pointer to [**[]InlineResponse200256ClientVpnConnections**](InlineResponse200256ClientVpnConnections.md) | VPN connections associated with the client | [optional] 
**Lldp** | Pointer to **[][]string** | The link layer discover protocol settings for the client | [optional] 
**Cdp** | Pointer to **[][]string** | The Cisco discover protocol settings for the client | [optional] 
**Status** | Pointer to **string** | The connection status of the client | [optional] 

## Methods

### NewInlineResponse200256Records

`func NewInlineResponse200256Records() *InlineResponse200256Records`

NewInlineResponse200256Records instantiates a new InlineResponse200256Records object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200256RecordsWithDefaults

`func NewInlineResponse200256RecordsWithDefaults() *InlineResponse200256Records`

NewInlineResponse200256RecordsWithDefaults instantiates a new InlineResponse200256Records object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200256Records) GetNetwork() InlineResponse200256Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200256Records) GetNetworkOk() (*InlineResponse200256Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200256Records) SetNetwork(v InlineResponse200256Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200256Records) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse200256Records) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse200256Records) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse200256Records) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse200256Records) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *InlineResponse200256Records) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *InlineResponse200256Records) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *InlineResponse200256Records) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *InlineResponse200256Records) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200256Records) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200256Records) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200256Records) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200256Records) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstSeen

`func (o *InlineResponse200256Records) GetFirstSeen() int32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *InlineResponse200256Records) GetFirstSeenOk() (*int32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *InlineResponse200256Records) SetFirstSeen(v int32)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *InlineResponse200256Records) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *InlineResponse200256Records) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *InlineResponse200256Records) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *InlineResponse200256Records) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *InlineResponse200256Records) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetOs

`func (o *InlineResponse200256Records) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InlineResponse200256Records) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InlineResponse200256Records) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *InlineResponse200256Records) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetUser

`func (o *InlineResponse200256Records) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineResponse200256Records) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineResponse200256Records) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineResponse200256Records) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200256Records) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200256Records) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200256Records) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200256Records) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse200256Records) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse200256Records) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse200256Records) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse200256Records) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSwitchport

`func (o *InlineResponse200256Records) GetSwitchport() string`

GetSwitchport returns the Switchport field if non-nil, zero value otherwise.

### GetSwitchportOk

`func (o *InlineResponse200256Records) GetSwitchportOk() (*string, bool)`

GetSwitchportOk returns a tuple with the Switchport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchport

`func (o *InlineResponse200256Records) SetSwitchport(v string)`

SetSwitchport sets Switchport field to given value.

### HasSwitchport

`func (o *InlineResponse200256Records) HasSwitchport() bool`

HasSwitchport returns a boolean if a field has been set.

### GetWirelessCapabilities

`func (o *InlineResponse200256Records) GetWirelessCapabilities() string`

GetWirelessCapabilities returns the WirelessCapabilities field if non-nil, zero value otherwise.

### GetWirelessCapabilitiesOk

`func (o *InlineResponse200256Records) GetWirelessCapabilitiesOk() (*string, bool)`

GetWirelessCapabilitiesOk returns a tuple with the WirelessCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessCapabilities

`func (o *InlineResponse200256Records) SetWirelessCapabilities(v string)`

SetWirelessCapabilities sets WirelessCapabilities field to given value.

### HasWirelessCapabilities

`func (o *InlineResponse200256Records) HasWirelessCapabilities() bool`

HasWirelessCapabilities returns a boolean if a field has been set.

### GetSmInstalled

`func (o *InlineResponse200256Records) GetSmInstalled() bool`

GetSmInstalled returns the SmInstalled field if non-nil, zero value otherwise.

### GetSmInstalledOk

`func (o *InlineResponse200256Records) GetSmInstalledOk() (*bool, bool)`

GetSmInstalledOk returns a tuple with the SmInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmInstalled

`func (o *InlineResponse200256Records) SetSmInstalled(v bool)`

SetSmInstalled sets SmInstalled field to given value.

### HasSmInstalled

`func (o *InlineResponse200256Records) HasSmInstalled() bool`

HasSmInstalled returns a boolean if a field has been set.

### GetRecentDeviceMac

`func (o *InlineResponse200256Records) GetRecentDeviceMac() string`

GetRecentDeviceMac returns the RecentDeviceMac field if non-nil, zero value otherwise.

### GetRecentDeviceMacOk

`func (o *InlineResponse200256Records) GetRecentDeviceMacOk() (*string, bool)`

GetRecentDeviceMacOk returns a tuple with the RecentDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeviceMac

`func (o *InlineResponse200256Records) SetRecentDeviceMac(v string)`

SetRecentDeviceMac sets RecentDeviceMac field to given value.

### HasRecentDeviceMac

`func (o *InlineResponse200256Records) HasRecentDeviceMac() bool`

HasRecentDeviceMac returns a boolean if a field has been set.

### GetClientVpnConnections

`func (o *InlineResponse200256Records) GetClientVpnConnections() []InlineResponse200256ClientVpnConnections`

GetClientVpnConnections returns the ClientVpnConnections field if non-nil, zero value otherwise.

### GetClientVpnConnectionsOk

`func (o *InlineResponse200256Records) GetClientVpnConnectionsOk() (*[]InlineResponse200256ClientVpnConnections, bool)`

GetClientVpnConnectionsOk returns a tuple with the ClientVpnConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnConnections

`func (o *InlineResponse200256Records) SetClientVpnConnections(v []InlineResponse200256ClientVpnConnections)`

SetClientVpnConnections sets ClientVpnConnections field to given value.

### HasClientVpnConnections

`func (o *InlineResponse200256Records) HasClientVpnConnections() bool`

HasClientVpnConnections returns a boolean if a field has been set.

### GetLldp

`func (o *InlineResponse200256Records) GetLldp() [][]string`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse200256Records) GetLldpOk() (*[][]string, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse200256Records) SetLldp(v [][]string)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse200256Records) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse200256Records) GetCdp() [][]string`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse200256Records) GetCdpOk() (*[][]string, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse200256Records) SetCdp(v [][]string)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse200256Records) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200256Records) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200256Records) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200256Records) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200256Records) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


