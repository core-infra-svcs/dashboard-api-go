# InlineResponse20039

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
**DefaultGateway** | Pointer to **string** | IPv4 default gateway | [optional] 

## Methods

### NewInlineResponse20039

`func NewInlineResponse20039() *InlineResponse20039`

NewInlineResponse20039 instantiates a new InlineResponse20039 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20039WithDefaults

`func NewInlineResponse20039WithDefaults() *InlineResponse20039`

NewInlineResponse20039WithDefaults instantiates a new InlineResponse20039 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *InlineResponse20039) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *InlineResponse20039) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *InlineResponse20039) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *InlineResponse20039) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20039) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20039) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20039) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20039) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse20039) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20039) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20039) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20039) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20039) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20039) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20039) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20039) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *InlineResponse20039) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineResponse20039) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineResponse20039) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *InlineResponse20039) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse20039) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20039) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20039) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20039) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSwitchPortId

`func (o *InlineResponse20039) GetSwitchPortId() string`

GetSwitchPortId returns the SwitchPortId field if non-nil, zero value otherwise.

### GetSwitchPortIdOk

`func (o *InlineResponse20039) GetSwitchPortIdOk() (*string, bool)`

GetSwitchPortIdOk returns a tuple with the SwitchPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortId

`func (o *InlineResponse20039) SetSwitchPortId(v string)`

SetSwitchPortId sets SwitchPortId field to given value.

### HasSwitchPortId

`func (o *InlineResponse20039) HasSwitchPortId() bool`

HasSwitchPortId returns a boolean if a field has been set.

### GetMulticastRouting

`func (o *InlineResponse20039) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *InlineResponse20039) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *InlineResponse20039) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *InlineResponse20039) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse20039) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse20039) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse20039) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse20039) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetUplinkV4

`func (o *InlineResponse20039) GetUplinkV4() bool`

GetUplinkV4 returns the UplinkV4 field if non-nil, zero value otherwise.

### GetUplinkV4Ok

`func (o *InlineResponse20039) GetUplinkV4Ok() (*bool, bool)`

GetUplinkV4Ok returns a tuple with the UplinkV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV4

`func (o *InlineResponse20039) SetUplinkV4(v bool)`

SetUplinkV4 sets UplinkV4 field to given value.

### HasUplinkV4

`func (o *InlineResponse20039) HasUplinkV4() bool`

HasUplinkV4 returns a boolean if a field has been set.

### GetUplinkV6

`func (o *InlineResponse20039) GetUplinkV6() bool`

GetUplinkV6 returns the UplinkV6 field if non-nil, zero value otherwise.

### GetUplinkV6Ok

`func (o *InlineResponse20039) GetUplinkV6Ok() (*bool, bool)`

GetUplinkV6Ok returns a tuple with the UplinkV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV6

`func (o *InlineResponse20039) SetUplinkV6(v bool)`

SetUplinkV6 sets UplinkV6 field to given value.

### HasUplinkV6

`func (o *InlineResponse20039) HasUplinkV6() bool`

HasUplinkV6 returns a boolean if a field has been set.

### GetOspfSettings

`func (o *InlineResponse20039) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesOspfSettings`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *InlineResponse20039) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesOspfSettings, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *InlineResponse20039) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesOspfSettings)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *InlineResponse20039) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetOspfV3

`func (o *InlineResponse20039) GetOspfV3() DevicesSerialSwitchRoutingInterfacesOspfV3`

GetOspfV3 returns the OspfV3 field if non-nil, zero value otherwise.

### GetOspfV3Ok

`func (o *InlineResponse20039) GetOspfV3Ok() (*DevicesSerialSwitchRoutingInterfacesOspfV3, bool)`

GetOspfV3Ok returns a tuple with the OspfV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfV3

`func (o *InlineResponse20039) SetOspfV3(v DevicesSerialSwitchRoutingInterfacesOspfV3)`

SetOspfV3 sets OspfV3 field to given value.

### HasOspfV3

`func (o *InlineResponse20039) HasOspfV3() bool`

HasOspfV3 returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20039) GetIpv6() DevicesSerialSwitchRoutingInterfacesIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20039) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20039) SetIpv6(v DevicesSerialSwitchRoutingInterfacesIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20039) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetVrf

`func (o *InlineResponse20039) GetVrf() DevicesSerialSwitchRoutingInterfacesVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineResponse20039) GetVrfOk() (*DevicesSerialSwitchRoutingInterfacesVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineResponse20039) SetVrf(v DevicesSerialSwitchRoutingInterfacesVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineResponse20039) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetLoopback

`func (o *InlineResponse20039) GetLoopback() map[string]interface{}`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *InlineResponse20039) GetLoopbackOk() (*map[string]interface{}, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *InlineResponse20039) SetLoopback(v map[string]interface{})`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *InlineResponse20039) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *InlineResponse20039) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *InlineResponse20039) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *InlineResponse20039) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *InlineResponse20039) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


