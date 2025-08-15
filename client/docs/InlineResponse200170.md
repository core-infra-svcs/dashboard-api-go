# InlineResponse200170

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **string** | The ID | [optional] 
**Name** | Pointer to **string** | The name | [optional] 
**Mode** | Pointer to **string** | The mode | [optional] 
**Subnet** | Pointer to **string** | IPv4 subnet | [optional] 
**InterfaceIp** | Pointer to **string** | IPv4 address | [optional] 
**Serial** | Pointer to **string** | Device serial | [optional] 
**SwitchPortId** | Pointer to **string** | Switch Port ID when in Routed mode | [optional] 
**MulticastRouting** | Pointer to **string** | Multicast routing status | [optional] 
**VlanId** | Pointer to **int32** | VLAN ID | [optional] 
**UplinkV4** | Pointer to **bool** | When true, this interface is used as static IPv4 uplink | [optional] 
**UplinkV6** | Pointer to **bool** | When true, this interface is used as static IPv6 uplink | [optional] 
**OspfSettings** | Pointer to [**DevicesSerialSwitchRoutingInterfacesOspfSettings**](DevicesSerialSwitchRoutingInterfacesOspfSettings.md) |  | [optional] 
**OspfV3** | Pointer to [**DevicesSerialSwitchRoutingInterfacesOspfV3**](DevicesSerialSwitchRoutingInterfacesOspfV3.md) |  | [optional] 
**Ipv6** | Pointer to [**DevicesSerialSwitchRoutingInterfacesIpv6**](DevicesSerialSwitchRoutingInterfacesIpv6.md) |  | [optional] 
**Vrf** | Pointer to [**DevicesSerialSwitchRoutingInterfacesVrf**](DevicesSerialSwitchRoutingInterfacesVrf.md) |  | [optional] 
**Loopback** | Pointer to **map[string]interface{}** | Loopback Interface settings | [optional] 

## Methods

### NewInlineResponse200170

`func NewInlineResponse200170() *InlineResponse200170`

NewInlineResponse200170 instantiates a new InlineResponse200170 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200170WithDefaults

`func NewInlineResponse200170WithDefaults() *InlineResponse200170`

NewInlineResponse200170WithDefaults instantiates a new InlineResponse200170 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *InlineResponse200170) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *InlineResponse200170) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *InlineResponse200170) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *InlineResponse200170) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200170) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200170) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200170) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200170) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse200170) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200170) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200170) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200170) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse200170) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse200170) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse200170) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse200170) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *InlineResponse200170) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineResponse200170) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineResponse200170) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *InlineResponse200170) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200170) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200170) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200170) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200170) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSwitchPortId

`func (o *InlineResponse200170) GetSwitchPortId() string`

GetSwitchPortId returns the SwitchPortId field if non-nil, zero value otherwise.

### GetSwitchPortIdOk

`func (o *InlineResponse200170) GetSwitchPortIdOk() (*string, bool)`

GetSwitchPortIdOk returns a tuple with the SwitchPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortId

`func (o *InlineResponse200170) SetSwitchPortId(v string)`

SetSwitchPortId sets SwitchPortId field to given value.

### HasSwitchPortId

`func (o *InlineResponse200170) HasSwitchPortId() bool`

HasSwitchPortId returns a boolean if a field has been set.

### GetMulticastRouting

`func (o *InlineResponse200170) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *InlineResponse200170) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *InlineResponse200170) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *InlineResponse200170) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse200170) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse200170) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse200170) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse200170) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetUplinkV4

`func (o *InlineResponse200170) GetUplinkV4() bool`

GetUplinkV4 returns the UplinkV4 field if non-nil, zero value otherwise.

### GetUplinkV4Ok

`func (o *InlineResponse200170) GetUplinkV4Ok() (*bool, bool)`

GetUplinkV4Ok returns a tuple with the UplinkV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV4

`func (o *InlineResponse200170) SetUplinkV4(v bool)`

SetUplinkV4 sets UplinkV4 field to given value.

### HasUplinkV4

`func (o *InlineResponse200170) HasUplinkV4() bool`

HasUplinkV4 returns a boolean if a field has been set.

### GetUplinkV6

`func (o *InlineResponse200170) GetUplinkV6() bool`

GetUplinkV6 returns the UplinkV6 field if non-nil, zero value otherwise.

### GetUplinkV6Ok

`func (o *InlineResponse200170) GetUplinkV6Ok() (*bool, bool)`

GetUplinkV6Ok returns a tuple with the UplinkV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV6

`func (o *InlineResponse200170) SetUplinkV6(v bool)`

SetUplinkV6 sets UplinkV6 field to given value.

### HasUplinkV6

`func (o *InlineResponse200170) HasUplinkV6() bool`

HasUplinkV6 returns a boolean if a field has been set.

### GetOspfSettings

`func (o *InlineResponse200170) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesOspfSettings`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *InlineResponse200170) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesOspfSettings, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *InlineResponse200170) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesOspfSettings)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *InlineResponse200170) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetOspfV3

`func (o *InlineResponse200170) GetOspfV3() DevicesSerialSwitchRoutingInterfacesOspfV3`

GetOspfV3 returns the OspfV3 field if non-nil, zero value otherwise.

### GetOspfV3Ok

`func (o *InlineResponse200170) GetOspfV3Ok() (*DevicesSerialSwitchRoutingInterfacesOspfV3, bool)`

GetOspfV3Ok returns a tuple with the OspfV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfV3

`func (o *InlineResponse200170) SetOspfV3(v DevicesSerialSwitchRoutingInterfacesOspfV3)`

SetOspfV3 sets OspfV3 field to given value.

### HasOspfV3

`func (o *InlineResponse200170) HasOspfV3() bool`

HasOspfV3 returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse200170) GetIpv6() DevicesSerialSwitchRoutingInterfacesIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse200170) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse200170) SetIpv6(v DevicesSerialSwitchRoutingInterfacesIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse200170) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse200170) GetVrf() DevicesSerialSwitchRoutingInterfacesVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse200170) GetVrfOk() (*DevicesSerialSwitchRoutingInterfacesVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse200170) SetVrf(v DevicesSerialSwitchRoutingInterfacesVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse200170) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetLoopback

`func (o *InlineResponse200170) GetLoopback() map[string]interface{}`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *InlineResponse200170) GetLoopbackOk() (*map[string]interface{}, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *InlineResponse200170) SetLoopback(v map[string]interface{})`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *InlineResponse200170) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


