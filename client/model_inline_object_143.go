/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject143 struct for InlineObject143
type InlineObject143 struct {
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

// NewInlineObject143 instantiates a new InlineObject143 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject143(name string, vlanId int32) *InlineObject143 {
	this := InlineObject143{}
	this.Name = name
	this.VlanId = vlanId
	return &this
}

// NewInlineObject143WithDefaults instantiates a new InlineObject143 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject143WithDefaults() *InlineObject143 {
	this := InlineObject143{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject143) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject143) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject143) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject143) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject143) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *InlineObject143) GetInterfaceIp() string {
	if o == nil || isNil(o.InterfaceIp) {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetInterfaceIpOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceIp) {
    return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *InlineObject143) HasInterfaceIp() bool {
	if o != nil && !isNil(o.InterfaceIp) {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *InlineObject143) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *InlineObject143) GetMulticastRouting() string {
	if o == nil || isNil(o.MulticastRouting) {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || isNil(o.MulticastRouting) {
    return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *InlineObject143) HasMulticastRouting() bool {
	if o != nil && !isNil(o.MulticastRouting) {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *InlineObject143) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value
func (o *InlineObject143) GetVlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetVlanIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *InlineObject143) SetVlanId(v int32) {
	o.VlanId = v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *InlineObject143) GetDefaultGateway() string {
	if o == nil || isNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || isNil(o.DefaultGateway) {
    return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *InlineObject143) HasDefaultGateway() bool {
	if o != nil && !isNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *InlineObject143) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *InlineObject143) GetOspfSettings() NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings {
	if o == nil || isNil(o.OspfSettings) {
		var ret NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetOspfSettingsOk() (*NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings, bool) {
	if o == nil || isNil(o.OspfSettings) {
    return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *InlineObject143) HasOspfSettings() bool {
	if o != nil && !isNil(o.OspfSettings) {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings and assigns it to the OspfSettings field.
func (o *InlineObject143) SetOspfSettings(v NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesOspfSettings) {
	o.OspfSettings = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineObject143) GetIpv6() NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject143) GetIpv6Ok() (*NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineObject143) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6 and assigns it to the Ipv6 field.
func (o *InlineObject143) SetIpv6(v NetworksNetworkIdSwitchStacksSwitchStackIdRoutingInterfacesIpv6) {
	o.Ipv6 = &v
}

func (o InlineObject143) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject143 struct {
	value *InlineObject143
	isSet bool
}

func (v NullableInlineObject143) Get() *InlineObject143 {
	return v.value
}

func (v *NullableInlineObject143) Set(val *InlineObject143) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject143) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject143) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject143(val *InlineObject143) *NullableInlineObject143 {
	return &NullableInlineObject143{value: val, isSet: true}
}

func (v NullableInlineObject143) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject143) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


