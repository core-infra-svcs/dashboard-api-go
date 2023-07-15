# InlineResponse20027

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
**Status** | Pointer to **string** | The connection status of the client | [optional] 
**Usage** | Pointer to [**InlineResponse20027Usage**](InlineResponse20027Usage.md) |  | [optional] 
**NamedVlan** | Pointer to **string** | Named VLAN of the client | [optional] 
**AdaptivePolicyGroup** | Pointer to **string** | The adaptive policy group of the client | [optional] 
**DeviceTypePrediction** | Pointer to **string** | Prediction of the client&#39;s device type | [optional] 
**RecentDeviceSerial** | Pointer to **string** | The serial of the node the device was last connected to | [optional] 
**RecentDeviceName** | Pointer to **string** | The name of the node the device was last connected to | [optional] 
**RecentDeviceConnection** | Pointer to **string** | Client&#39;s most recent connection type | [optional] 
**Notes** | Pointer to **string** | Notes on the client | [optional] 
**Ip6Local** | Pointer to **string** | Local IPv6 address of the client | [optional] 
**GroupPolicy8021x** | Pointer to **string** | 802.1x group policy of the client | [optional] 
**PskGroup** | Pointer to **string** | iPSK name of the client | [optional] 

## Methods

### NewInlineResponse20027

`func NewInlineResponse20027() *InlineResponse20027`

NewInlineResponse20027 instantiates a new InlineResponse20027 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20027WithDefaults

`func NewInlineResponse20027WithDefaults() *InlineResponse20027`

NewInlineResponse20027WithDefaults instantiates a new InlineResponse20027 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20027) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20027) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20027) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20027) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse20027) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse20027) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse20027) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse20027) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetIp

`func (o *InlineResponse20027) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineResponse20027) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineResponse20027) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineResponse20027) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *InlineResponse20027) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *InlineResponse20027) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *InlineResponse20027) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *InlineResponse20027) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20027) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20027) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20027) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20027) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstSeen

`func (o *InlineResponse20027) GetFirstSeen() int32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *InlineResponse20027) GetFirstSeenOk() (*int32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *InlineResponse20027) SetFirstSeen(v int32)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *InlineResponse20027) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *InlineResponse20027) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *InlineResponse20027) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *InlineResponse20027) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *InlineResponse20027) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetManufacturer

`func (o *InlineResponse20027) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InlineResponse20027) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InlineResponse20027) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InlineResponse20027) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetOs

`func (o *InlineResponse20027) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InlineResponse20027) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InlineResponse20027) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *InlineResponse20027) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetUser

`func (o *InlineResponse20027) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineResponse20027) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineResponse20027) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineResponse20027) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse20027) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse20027) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse20027) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse20027) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse20027) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse20027) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse20027) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse20027) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSwitchport

`func (o *InlineResponse20027) GetSwitchport() string`

GetSwitchport returns the Switchport field if non-nil, zero value otherwise.

### GetSwitchportOk

`func (o *InlineResponse20027) GetSwitchportOk() (*string, bool)`

GetSwitchportOk returns a tuple with the Switchport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchport

`func (o *InlineResponse20027) SetSwitchport(v string)`

SetSwitchport sets Switchport field to given value.

### HasSwitchport

`func (o *InlineResponse20027) HasSwitchport() bool`

HasSwitchport returns a boolean if a field has been set.

### GetWirelessCapabilities

`func (o *InlineResponse20027) GetWirelessCapabilities() string`

GetWirelessCapabilities returns the WirelessCapabilities field if non-nil, zero value otherwise.

### GetWirelessCapabilitiesOk

`func (o *InlineResponse20027) GetWirelessCapabilitiesOk() (*string, bool)`

GetWirelessCapabilitiesOk returns a tuple with the WirelessCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessCapabilities

`func (o *InlineResponse20027) SetWirelessCapabilities(v string)`

SetWirelessCapabilities sets WirelessCapabilities field to given value.

### HasWirelessCapabilities

`func (o *InlineResponse20027) HasWirelessCapabilities() bool`

HasWirelessCapabilities returns a boolean if a field has been set.

### GetSmInstalled

`func (o *InlineResponse20027) GetSmInstalled() bool`

GetSmInstalled returns the SmInstalled field if non-nil, zero value otherwise.

### GetSmInstalledOk

`func (o *InlineResponse20027) GetSmInstalledOk() (*bool, bool)`

GetSmInstalledOk returns a tuple with the SmInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmInstalled

`func (o *InlineResponse20027) SetSmInstalled(v bool)`

SetSmInstalled sets SmInstalled field to given value.

### HasSmInstalled

`func (o *InlineResponse20027) HasSmInstalled() bool`

HasSmInstalled returns a boolean if a field has been set.

### GetRecentDeviceMac

`func (o *InlineResponse20027) GetRecentDeviceMac() string`

GetRecentDeviceMac returns the RecentDeviceMac field if non-nil, zero value otherwise.

### GetRecentDeviceMacOk

`func (o *InlineResponse20027) GetRecentDeviceMacOk() (*string, bool)`

GetRecentDeviceMacOk returns a tuple with the RecentDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeviceMac

`func (o *InlineResponse20027) SetRecentDeviceMac(v string)`

SetRecentDeviceMac sets RecentDeviceMac field to given value.

### HasRecentDeviceMac

`func (o *InlineResponse20027) HasRecentDeviceMac() bool`

HasRecentDeviceMac returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20027) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20027) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20027) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20027) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsage

`func (o *InlineResponse20027) GetUsage() InlineResponse20027Usage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *InlineResponse20027) GetUsageOk() (*InlineResponse20027Usage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *InlineResponse20027) SetUsage(v InlineResponse20027Usage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *InlineResponse20027) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetNamedVlan

`func (o *InlineResponse20027) GetNamedVlan() string`

GetNamedVlan returns the NamedVlan field if non-nil, zero value otherwise.

### GetNamedVlanOk

`func (o *InlineResponse20027) GetNamedVlanOk() (*string, bool)`

GetNamedVlanOk returns a tuple with the NamedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlan

`func (o *InlineResponse20027) SetNamedVlan(v string)`

SetNamedVlan sets NamedVlan field to given value.

### HasNamedVlan

`func (o *InlineResponse20027) HasNamedVlan() bool`

HasNamedVlan returns a boolean if a field has been set.

### GetAdaptivePolicyGroup

`func (o *InlineResponse20027) GetAdaptivePolicyGroup() string`

GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupOk

`func (o *InlineResponse20027) GetAdaptivePolicyGroupOk() (*string, bool)`

GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroup

`func (o *InlineResponse20027) SetAdaptivePolicyGroup(v string)`

SetAdaptivePolicyGroup sets AdaptivePolicyGroup field to given value.

### HasAdaptivePolicyGroup

`func (o *InlineResponse20027) HasAdaptivePolicyGroup() bool`

HasAdaptivePolicyGroup returns a boolean if a field has been set.

### GetDeviceTypePrediction

`func (o *InlineResponse20027) GetDeviceTypePrediction() string`

GetDeviceTypePrediction returns the DeviceTypePrediction field if non-nil, zero value otherwise.

### GetDeviceTypePredictionOk

`func (o *InlineResponse20027) GetDeviceTypePredictionOk() (*string, bool)`

GetDeviceTypePredictionOk returns a tuple with the DeviceTypePrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypePrediction

`func (o *InlineResponse20027) SetDeviceTypePrediction(v string)`

SetDeviceTypePrediction sets DeviceTypePrediction field to given value.

### HasDeviceTypePrediction

`func (o *InlineResponse20027) HasDeviceTypePrediction() bool`

HasDeviceTypePrediction returns a boolean if a field has been set.

### GetRecentDeviceSerial

`func (o *InlineResponse20027) GetRecentDeviceSerial() string`

GetRecentDeviceSerial returns the RecentDeviceSerial field if non-nil, zero value otherwise.

### GetRecentDeviceSerialOk

`func (o *InlineResponse20027) GetRecentDeviceSerialOk() (*string, bool)`

GetRecentDeviceSerialOk returns a tuple with the RecentDeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeviceSerial

`func (o *InlineResponse20027) SetRecentDeviceSerial(v string)`

SetRecentDeviceSerial sets RecentDeviceSerial field to given value.

### HasRecentDeviceSerial

`func (o *InlineResponse20027) HasRecentDeviceSerial() bool`

HasRecentDeviceSerial returns a boolean if a field has been set.

### GetRecentDeviceName

`func (o *InlineResponse20027) GetRecentDeviceName() string`

GetRecentDeviceName returns the RecentDeviceName field if non-nil, zero value otherwise.

### GetRecentDeviceNameOk

`func (o *InlineResponse20027) GetRecentDeviceNameOk() (*string, bool)`

GetRecentDeviceNameOk returns a tuple with the RecentDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeviceName

`func (o *InlineResponse20027) SetRecentDeviceName(v string)`

SetRecentDeviceName sets RecentDeviceName field to given value.

### HasRecentDeviceName

`func (o *InlineResponse20027) HasRecentDeviceName() bool`

HasRecentDeviceName returns a boolean if a field has been set.

### GetRecentDeviceConnection

`func (o *InlineResponse20027) GetRecentDeviceConnection() string`

GetRecentDeviceConnection returns the RecentDeviceConnection field if non-nil, zero value otherwise.

### GetRecentDeviceConnectionOk

`func (o *InlineResponse20027) GetRecentDeviceConnectionOk() (*string, bool)`

GetRecentDeviceConnectionOk returns a tuple with the RecentDeviceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeviceConnection

`func (o *InlineResponse20027) SetRecentDeviceConnection(v string)`

SetRecentDeviceConnection sets RecentDeviceConnection field to given value.

### HasRecentDeviceConnection

`func (o *InlineResponse20027) HasRecentDeviceConnection() bool`

HasRecentDeviceConnection returns a boolean if a field has been set.

### GetNotes

`func (o *InlineResponse20027) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InlineResponse20027) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InlineResponse20027) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InlineResponse20027) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIp6Local

`func (o *InlineResponse20027) GetIp6Local() string`

GetIp6Local returns the Ip6Local field if non-nil, zero value otherwise.

### GetIp6LocalOk

`func (o *InlineResponse20027) GetIp6LocalOk() (*string, bool)`

GetIp6LocalOk returns a tuple with the Ip6Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6Local

`func (o *InlineResponse20027) SetIp6Local(v string)`

SetIp6Local sets Ip6Local field to given value.

### HasIp6Local

`func (o *InlineResponse20027) HasIp6Local() bool`

HasIp6Local returns a boolean if a field has been set.

### GetGroupPolicy8021x

`func (o *InlineResponse20027) GetGroupPolicy8021x() string`

GetGroupPolicy8021x returns the GroupPolicy8021x field if non-nil, zero value otherwise.

### GetGroupPolicy8021xOk

`func (o *InlineResponse20027) GetGroupPolicy8021xOk() (*string, bool)`

GetGroupPolicy8021xOk returns a tuple with the GroupPolicy8021x field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicy8021x

`func (o *InlineResponse20027) SetGroupPolicy8021x(v string)`

SetGroupPolicy8021x sets GroupPolicy8021x field to given value.

### HasGroupPolicy8021x

`func (o *InlineResponse20027) HasGroupPolicy8021x() bool`

HasGroupPolicy8021x returns a boolean if a field has been set.

### GetPskGroup

`func (o *InlineResponse20027) GetPskGroup() string`

GetPskGroup returns the PskGroup field if non-nil, zero value otherwise.

### GetPskGroupOk

`func (o *InlineResponse20027) GetPskGroupOk() (*string, bool)`

GetPskGroupOk returns a tuple with the PskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskGroup

`func (o *InlineResponse20027) SetPskGroup(v string)`

SetPskGroup sets PskGroup field to given value.

### HasPskGroup

`func (o *InlineResponse20027) HasPskGroup() bool`

HasPskGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


