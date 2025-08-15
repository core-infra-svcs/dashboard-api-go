# InlineObject162

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A friendly name or description for the interface or VLAN (max length 128 characters). | [optional] 
**Subnet** | Pointer to **NullableString** | The network that this L3 interface is on, in CIDR notation (ex. 10.1.1.0/24). | [optional] 
**SwitchPortId** | Pointer to **NullableString** | Switch Port ID when in Routed mode (CS 17.18 or higher required) | [optional] 
**InterfaceIp** | Pointer to **NullableString** | The IP address that will be used for Layer 3 routing on this VLAN or subnet. This cannot be the same         as the device management IP. | [optional] 
**MulticastRouting** | Pointer to **string** | Enable multicast support if, multicast routing between VLANs is required. Options are:         &#39;disabled&#39;, &#39;enabled&#39; or &#39;IGMP snooping querier&#39;. Default is &#39;disabled&#39;. | [optional] 
**VlanId** | Pointer to **NullableInt32** | The VLAN this L3 interface is on. VLAN must be between 1 and 4094. | [optional] 
**DefaultGateway** | Pointer to **string** | The next hop for any traffic that isn&#39;t going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a L3 interface. Required if this is the first IPv4 interface. | [optional] 
**OspfSettings** | Pointer to [**DevicesSerialSwitchRoutingInterfacesOspfSettings1**](DevicesSerialSwitchRoutingInterfacesOspfSettings1.md) |  | [optional] 
**Ipv6** | Pointer to [**DevicesSerialSwitchRoutingInterfacesIpv61**](DevicesSerialSwitchRoutingInterfacesIpv61.md) |  | [optional] 
**Vrf** | Pointer to [**DevicesSerialSwitchRoutingInterfacesVrf1**](DevicesSerialSwitchRoutingInterfacesVrf1.md) |  | [optional] 
**Loopback** | Pointer to **map[string]interface{}** | The loopback settings of the interface. | [optional] 

## Methods

### NewInlineObject162

`func NewInlineObject162() *InlineObject162`

NewInlineObject162 instantiates a new InlineObject162 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject162WithDefaults

`func NewInlineObject162WithDefaults() *InlineObject162`

NewInlineObject162WithDefaults instantiates a new InlineObject162 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject162) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject162) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject162) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject162) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineObject162) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject162) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject162) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject162) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *InlineObject162) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *InlineObject162) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetSwitchPortId

`func (o *InlineObject162) GetSwitchPortId() string`

GetSwitchPortId returns the SwitchPortId field if non-nil, zero value otherwise.

### GetSwitchPortIdOk

`func (o *InlineObject162) GetSwitchPortIdOk() (*string, bool)`

GetSwitchPortIdOk returns a tuple with the SwitchPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortId

`func (o *InlineObject162) SetSwitchPortId(v string)`

SetSwitchPortId sets SwitchPortId field to given value.

### HasSwitchPortId

`func (o *InlineObject162) HasSwitchPortId() bool`

HasSwitchPortId returns a boolean if a field has been set.

### SetSwitchPortIdNil

`func (o *InlineObject162) SetSwitchPortIdNil(b bool)`

 SetSwitchPortIdNil sets the value for SwitchPortId to be an explicit nil

### UnsetSwitchPortId
`func (o *InlineObject162) UnsetSwitchPortId()`

UnsetSwitchPortId ensures that no value is present for SwitchPortId, not even an explicit nil
### GetInterfaceIp

`func (o *InlineObject162) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineObject162) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineObject162) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *InlineObject162) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### SetInterfaceIpNil

`func (o *InlineObject162) SetInterfaceIpNil(b bool)`

 SetInterfaceIpNil sets the value for InterfaceIp to be an explicit nil

### UnsetInterfaceIp
`func (o *InlineObject162) UnsetInterfaceIp()`

UnsetInterfaceIp ensures that no value is present for InterfaceIp, not even an explicit nil
### GetMulticastRouting

`func (o *InlineObject162) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *InlineObject162) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *InlineObject162) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *InlineObject162) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineObject162) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineObject162) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineObject162) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineObject162) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *InlineObject162) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *InlineObject162) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetDefaultGateway

`func (o *InlineObject162) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *InlineObject162) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *InlineObject162) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *InlineObject162) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetOspfSettings

`func (o *InlineObject162) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesOspfSettings1`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *InlineObject162) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesOspfSettings1, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *InlineObject162) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesOspfSettings1)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *InlineObject162) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineObject162) GetIpv6() DevicesSerialSwitchRoutingInterfacesIpv61`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineObject162) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesIpv61, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineObject162) SetIpv6(v DevicesSerialSwitchRoutingInterfacesIpv61)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineObject162) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetVrf

`func (o *InlineObject162) GetVrf() DevicesSerialSwitchRoutingInterfacesVrf1`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InlineObject162) GetVrfOk() (*DevicesSerialSwitchRoutingInterfacesVrf1, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InlineObject162) SetVrf(v DevicesSerialSwitchRoutingInterfacesVrf1)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InlineObject162) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetLoopback

`func (o *InlineObject162) GetLoopback() map[string]interface{}`

GetLoopback returns the Loopback field if non-nil, zero value otherwise.

### GetLoopbackOk

`func (o *InlineObject162) GetLoopbackOk() (*map[string]interface{}, bool)`

GetLoopbackOk returns a tuple with the Loopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopback

`func (o *InlineObject162) SetLoopback(v map[string]interface{})`

SetLoopback sets Loopback field to given value.

### HasLoopback

`func (o *InlineObject162) HasLoopback() bool`

HasLoopback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


