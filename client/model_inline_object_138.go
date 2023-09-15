/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject138 struct for InlineObject138
type InlineObject138 struct {
	// A friendly name or description for the interface or VLAN.
	Name string `json:"name"`
	// The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet *string `json:"subnet,omitempty"`
	// The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// The VLAN this routed interface is on. VLAN must be between 1 and 4094.
	VlanId int32 `json:"vlanId"`
	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
	DefaultGateway *string `json:"defaultGateway,omitempty"`
	OspfSettings *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings `json:"ospfSettings,omitempty"`
	Ipv6 *NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6 `json:"ipv6,omitempty"`
}

// NewInlineObject138 instantiates a new InlineObject138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject138(name string, vlanId int32) *InlineObject138 {
	this := InlineObject138{}
	this.Name = name
	this.VlanId = vlanId
	return &this
}

// NewInlineObject138WithDefaults instantiates a new InlineObject138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject138WithDefaults() *InlineObject138 {
	this := InlineObject138{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject138) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject138) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject138) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject138) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject138) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *InlineObject138) GetInterfaceIp() string {
	if o == nil || isNil(o.InterfaceIp) {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetInterfaceIpOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceIp) {
    return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *InlineObject138) HasInterfaceIp() bool {
	if o != nil && !isNil(o.InterfaceIp) {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *InlineObject138) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *InlineObject138) GetMulticastRouting() string {
	if o == nil || isNil(o.MulticastRouting) {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || isNil(o.MulticastRouting) {
    return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *InlineObject138) HasMulticastRouting() bool {
	if o != nil && !isNil(o.MulticastRouting) {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *InlineObject138) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value
func (o *InlineObject138) GetVlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetVlanIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *InlineObject138) SetVlanId(v int32) {
	o.VlanId = v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *InlineObject138) GetDefaultGateway() string {
	if o == nil || isNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || isNil(o.DefaultGateway) {
    return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *InlineObject138) HasDefaultGateway() bool {
	if o != nil && !isNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *InlineObject138) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *InlineObject138) GetOspfSettings() NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings {
	if o == nil || isNil(o.OspfSettings) {
		var ret NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetOspfSettingsOk() (*NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings, bool) {
	if o == nil || isNil(o.OspfSettings) {
    return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *InlineObject138) HasOspfSettings() bool {
	if o != nil && !isNil(o.OspfSettings) {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings and assigns it to the OspfSettings field.
func (o *InlineObject138) SetOspfSettings(v NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) {
	o.OspfSettings = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineObject138) GetIpv6() NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetIpv6Ok() (*NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineObject138) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6 and assigns it to the Ipv6 field.
func (o *InlineObject138) SetIpv6(v NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6) {
	o.Ipv6 = &v
}

func (o InlineObject138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.InterfaceIp) {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if !isNil(o.MulticastRouting) {
		toSerialize["multicastRouting"] = o.MulticastRouting
	}
	if true {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.DefaultGateway) {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if !isNil(o.OspfSettings) {
		toSerialize["ospfSettings"] = o.OspfSettings
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject138 struct {
	value *InlineObject138
	isSet bool
}

func (v NullableInlineObject138) Get() *InlineObject138 {
	return v.value
}

func (v *NullableInlineObject138) Set(val *InlineObject138) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject138) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject138) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject138(val *InlineObject138) *NullableInlineObject138 {
	return &NullableInlineObject138{value: val, isSet: true}
}

func (v NullableInlineObject138) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject138) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


