# InlineResponse20017

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the client | [optional] 
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**Ip** | Pointer to **string** | The IP address of the client | [optional] 
**Ip6** | Pointer to **string** | The IPv6 address of the client | [optional] 
**Description** | Pointer to **string** | Short description of the client | [optional] 
**FirstSeen** | Pointer to **int32** | Timestamp client was first seen in the network | [optional] 
**LastSeen** | Pointer to **int32** | Timestamp client was last seen in the network | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the client | [optional] 
**Os** | Pointer to **string** | The operating system of the client | [optional] 
**User** | Pointer to **string** | The username of the user of the client | [optional] 
**Vlan** | Pointer to **string** | The name of the VLAN that the client is connected to | [optional] 
**Ssid** | Pointer to **string** | The name of the SSID that the client is connected to | [optional] 
**Switchport** | Pointer to **string** | The switch port that the client is connected to | [optional] 
**WirelessCapabilities** | Pointer to **string** | Wireless capabilities of the client | [optional] 
**SmInstalled** | Pointer to **bool** | Status of SM for the client | [optional] 
**RecentDeviceMac** | Pointer to **string** | The MAC address of the node that the device was last connected to | [optional] 
**ClientVpnConnections** | Pointer to [**[]InlineResponse20017ClientVpnConnections**](InlineResponse20017ClientVpnConnections.md) | VPN connections associated with the client | [optional] 
**Lldp** | Pointer to **[][]string** | The link layer discover protocol settings for the client | [optional] 
**Cdp** | Pointer to **[][]string** | The Cisco discover protocol settings for the client | [optional] 
**Status** | Pointer to **string** | The connection status of the client | [optional] 

## Methods

### NewInlineResponse20017

`func NewInlineResponse20017() *InlineResponse20017`

NewInlineResponse20017 instantiates a new InlineResponse20017 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20017WithDefaults

`func NewInlineResponse20017WithDefaults() *InlineResponse20017`

NewInlineResponse20017WithDefaults instantiates a new InlineResponse20017 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20017) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20017) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20017) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20017) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse20017) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse20017) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse20017) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse20017) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse20017) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20017) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20017) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20017) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *InlineResponse20017) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *InlineResponse20017) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *InlineResponse20017) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *InlineResponse20017) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20017) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20017) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20017) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20017) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstSeen

`func (o *InlineResponse20017) GetFirstSeen() int32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *InlineResponse20017) GetFirstSeenOk() (*int32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *InlineResponse20017) SetFirstSeen(v int32)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *InlineResponse20017) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *InlineResponse20017) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *InlineResponse20017) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *InlineResponse20017) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *InlineResponse20017) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetManufacturer

`func (o *InlineResponse20017) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InlineResponse20017) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InlineResponse20017) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InlineResponse20017) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetOs

`func (o *InlineResponse20017) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InlineResponse20017) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InlineResponse20017) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *InlineResponse20017) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetUser

`func (o *InlineResponse20017) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineResponse20017) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineResponse20017) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineResponse20017) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20017) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20017) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20017) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20017) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse20017) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse20017) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse20017) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse20017) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSwitchport

`func (o *InlineResponse20017) GetSwitchport() string`

GetSwitchport returns the Switchport field if non-nil, zero value otherwise.

### GetSwitchportOk

`func (o *InlineResponse20017) GetSwitchportOk() (*string, bool)`

GetSwitchportOk returns a tuple with the Switchport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchport

`func (o *InlineResponse20017) SetSwitchport(v string)`

SetSwitchport sets Switchport field to given value.

### HasSwitchport

`func (o *InlineResponse20017) HasSwitchport() bool`

HasSwitchport returns a boolean if a field has been set.

### GetWirelessCapabilities

`func (o *InlineResponse20017) GetWirelessCapabilities() string`

GetWirelessCapabilities returns the WirelessCapabilities field if non-nil, zero value otherwise.

### GetWirelessCapabilitiesOk

`func (o *InlineResponse20017) GetWirelessCapabilitiesOk() (*string, bool)`

GetWirelessCapabilitiesOk returns a tuple with the WirelessCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessCapabilities

`func (o *InlineResponse20017) SetWirelessCapabilities(v string)`

SetWirelessCapabilities sets WirelessCapabilities field to given value.

### HasWirelessCapabilities

`func (o *InlineResponse20017) HasWirelessCapabilities() bool`

HasWirelessCapabilities returns a boolean if a field has been set.

### GetSmInstalled

`func (o *InlineResponse20017) GetSmInstalled() bool`

GetSmInstalled returns the SmInstalled field if non-nil, zero value otherwise.

### GetSmInstalledOk

`func (o *InlineResponse20017) GetSmInstalledOk() (*bool, bool)`

GetSmInstalledOk returns a tuple with the SmInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmInstalled

`func (o *InlineResponse20017) SetSmInstalled(v bool)`

SetSmInstalled sets SmInstalled field to given value.

### HasSmInstalled

`func (o *InlineResponse20017) HasSmInstalled() bool`

HasSmInstalled returns a boolean if a field has been set.

### GetRecentDeviceMac

`func (o *InlineResponse20017) GetRecentDeviceMac() string`

GetRecentDeviceMac returns the RecentDeviceMac field if non-nil, zero value otherwise.

### GetRecentDeviceMacOk

`func (o *InlineResponse20017) GetRecentDeviceMacOk() (*string, bool)`

GetRecentDeviceMacOk returns a tuple with the RecentDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeviceMac

`func (o *InlineResponse20017) SetRecentDeviceMac(v string)`

SetRecentDeviceMac sets RecentDeviceMac field to given value.

### HasRecentDeviceMac

`func (o *InlineResponse20017) HasRecentDeviceMac() bool`

HasRecentDeviceMac returns a boolean if a field has been set.

### GetClientVpnConnections

`func (o *InlineResponse20017) GetClientVpnConnections() []InlineResponse20017ClientVpnConnections`

GetClientVpnConnections returns the ClientVpnConnections field if non-nil, zero value otherwise.

### GetClientVpnConnectionsOk

`func (o *InlineResponse20017) GetClientVpnConnectionsOk() (*[]InlineResponse20017ClientVpnConnections, bool)`

GetClientVpnConnectionsOk returns a tuple with the ClientVpnConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnConnections

`func (o *InlineResponse20017) SetClientVpnConnections(v []InlineResponse20017ClientVpnConnections)`

SetClientVpnConnections sets ClientVpnConnections field to given value.

### HasClientVpnConnections

`func (o *InlineResponse20017) HasClientVpnConnections() bool`

HasClientVpnConnections returns a boolean if a field has been set.

### GetLldp

`func (o *InlineResponse20017) GetLldp() [][]string`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *InlineResponse20017) GetLldpOk() (*[][]string, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *InlineResponse20017) SetLldp(v [][]string)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *InlineResponse20017) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *InlineResponse20017) GetCdp() [][]string`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *InlineResponse20017) GetCdpOk() (*[][]string, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *InlineResponse20017) SetCdp(v [][]string)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *InlineResponse20017) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20017) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20017) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20017) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20017) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


