# InlineResponse20035

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **string** | The id | [optional] 
**Name** | Pointer to **string** | The name | [optional] 
**Subnet** | Pointer to **string** | IPv4 subnet | [optional] 
**InterfaceIp** | Pointer to **string** | IPv4 address | [optional] 
**MulticastRouting** | Pointer to **string** | Multicast routing status | [optional] 
**VlanId** | Pointer to **int32** | VLAN id | [optional] 
**OspfSettings** | Pointer to [**DevicesSerialSwitchRoutingInterfacesOspfSettings**](DevicesSerialSwitchRoutingInterfacesOspfSettings.md) |  | [optional] 
**OspfV3** | Pointer to [**DevicesSerialSwitchRoutingInterfacesOspfV3**](DevicesSerialSwitchRoutingInterfacesOspfV3.md) |  | [optional] 
**Ipv6** | Pointer to [**DevicesSerialSwitchRoutingInterfacesIpv6**](DevicesSerialSwitchRoutingInterfacesIpv6.md) |  | [optional] 
**UplinkV4** | Pointer to **bool** | Whether this is the switch&#39;s IPv4 uplink | [optional] 
**UplinkV6** | Pointer to **bool** | Whether this is the switch&#39;s IPv6 uplink | [optional] 
**DefaultGateway** | Pointer to **string** | IPv4 default gateway | [optional] 

## Methods

### NewInlineResponse20035

`func NewInlineResponse20035() *InlineResponse20035`

NewInlineResponse20035 instantiates a new InlineResponse20035 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20035WithDefaults

`func NewInlineResponse20035WithDefaults() *InlineResponse20035`

NewInlineResponse20035WithDefaults instantiates a new InlineResponse20035 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *InlineResponse20035) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *InlineResponse20035) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *InlineResponse20035) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *InlineResponse20035) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20035) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20035) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20035) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20035) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse20035) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse20035) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse20035) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse20035) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *InlineResponse20035) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineResponse20035) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineResponse20035) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *InlineResponse20035) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetMulticastRouting

`func (o *InlineResponse20035) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *InlineResponse20035) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *InlineResponse20035) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *InlineResponse20035) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse20035) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse20035) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse20035) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse20035) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetOspfSettings

`func (o *InlineResponse20035) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesOspfSettings`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *InlineResponse20035) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesOspfSettings, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *InlineResponse20035) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesOspfSettings)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *InlineResponse20035) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetOspfV3

`func (o *InlineResponse20035) GetOspfV3() DevicesSerialSwitchRoutingInterfacesOspfV3`

GetOspfV3 returns the OspfV3 field if non-nil, zero value otherwise.

### GetOspfV3Ok

`func (o *InlineResponse20035) GetOspfV3Ok() (*DevicesSerialSwitchRoutingInterfacesOspfV3, bool)`

GetOspfV3Ok returns a tuple with the OspfV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfV3

`func (o *InlineResponse20035) SetOspfV3(v DevicesSerialSwitchRoutingInterfacesOspfV3)`

SetOspfV3 sets OspfV3 field to given value.

### HasOspfV3

`func (o *InlineResponse20035) HasOspfV3() bool`

HasOspfV3 returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineResponse20035) GetIpv6() DevicesSerialSwitchRoutingInterfacesIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineResponse20035) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineResponse20035) SetIpv6(v DevicesSerialSwitchRoutingInterfacesIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineResponse20035) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetUplinkV4

`func (o *InlineResponse20035) GetUplinkV4() bool`

GetUplinkV4 returns the UplinkV4 field if non-nil, zero value otherwise.

### GetUplinkV4Ok

`func (o *InlineResponse20035) GetUplinkV4Ok() (*bool, bool)`

GetUplinkV4Ok returns a tuple with the UplinkV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV4

`func (o *InlineResponse20035) SetUplinkV4(v bool)`

SetUplinkV4 sets UplinkV4 field to given value.

### HasUplinkV4

`func (o *InlineResponse20035) HasUplinkV4() bool`

HasUplinkV4 returns a boolean if a field has been set.

### GetUplinkV6

`func (o *InlineResponse20035) GetUplinkV6() bool`

GetUplinkV6 returns the UplinkV6 field if non-nil, zero value otherwise.

### GetUplinkV6Ok

`func (o *InlineResponse20035) GetUplinkV6Ok() (*bool, bool)`

GetUplinkV6Ok returns a tuple with the UplinkV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV6

`func (o *InlineResponse20035) SetUplinkV6(v bool)`

SetUplinkV6 sets UplinkV6 field to given value.

### HasUplinkV6

`func (o *InlineResponse20035) HasUplinkV6() bool`

HasUplinkV6 returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *InlineResponse20035) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *InlineResponse20035) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *InlineResponse20035) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *InlineResponse20035) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


