# InlineObject131

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A friendly name or description for the interface or VLAN. | 
**Subnet** | Pointer to **string** | The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24). | [optional] 
**InterfaceIp** | Pointer to **string** | The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch&#39;s management IP. | [optional] 
**MulticastRouting** | Pointer to **string** | Enable multicast support if, multicast routing between VLANs is required. Options are, &#39;disabled&#39;, &#39;enabled&#39; or &#39;IGMP snooping querier&#39;. Default is &#39;disabled&#39;. | [optional] 
**VlanId** | **int32** | The VLAN this routed interface is on. VLAN must be between 1 and 4094. | 
**DefaultGateway** | Pointer to **string** | The next hop for any traffic that isn&#39;t going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface. | [optional] 
**OspfSettings** | Pointer to [**NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings**](NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings.md) |  | [optional] 
**Ipv6** | Pointer to [**NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6**](NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6.md) |  | [optional] 

## Methods

### NewInlineObject131

`func NewInlineObject131(name string, vlanId int32, ) *InlineObject131`

NewInlineObject131 instantiates a new InlineObject131 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject131WithDefaults

`func NewInlineObject131WithDefaults() *InlineObject131`

NewInlineObject131WithDefaults instantiates a new InlineObject131 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject131) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject131) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject131) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *InlineObject131) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject131) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject131) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject131) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *InlineObject131) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *InlineObject131) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *InlineObject131) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *InlineObject131) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetMulticastRouting

`func (o *InlineObject131) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *InlineObject131) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *InlineObject131) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *InlineObject131) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineObject131) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineObject131) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineObject131) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetDefaultGateway

`func (o *InlineObject131) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *InlineObject131) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *InlineObject131) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *InlineObject131) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetOspfSettings

`func (o *InlineObject131) GetOspfSettings() NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *InlineObject131) GetOspfSettingsOk() (*NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *InlineObject131) SetOspfSettings(v NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *InlineObject131) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetIpv6

`func (o *InlineObject131) GetIpv6() NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InlineObject131) GetIpv6Ok() (*NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InlineObject131) SetIpv6(v NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *InlineObject131) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


