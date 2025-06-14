/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject161 struct for InlineObject161
type InlineObject161 struct {
	// A friendly name or description for the interface or VLAN (max length 128 characters).
	Name string `json:"name"`
	// The network that this L3 interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet NullableString `json:"subnet,omitempty"`
	// The IP address that will be used for Layer 3 routing on this VLAN or subnet. This cannot be the same         as the device management IP.
	InterfaceIp NullableString `json:"interfaceIp,omitempty"`
	// Enable multicast support if, multicast routing between VLANs is required. Options are:         'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// The VLAN this L3 interface is on. VLAN must be between 1 and 4094.
	VlanId NullableInt32 `json:"vlanId,omitempty"`
	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a L3 interface. Required if this is the first IPv4 interface.
	DefaultGateway *string `json:"defaultGateway,omitempty"`
	OspfSettings *DevicesSerialSwitchRoutingInterfacesOspfSettings1 `json:"ospfSettings,omitempty"`
	Ipv6 *DevicesSerialSwitchRoutingInterfacesIpv61 `json:"ipv6,omitempty"`
}

// NewInlineObject161 instantiates a new InlineObject161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject161(name string) *InlineObject161 {
	this := InlineObject161{}
	this.Name = name
	return &this
}

// NewInlineObject161WithDefaults instantiates a new InlineObject161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject161WithDefaults() *InlineObject161 {
	this := InlineObject161{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject161) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject161) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject161) GetSubnet() string {
	if o == nil || isNil(o.Subnet.Get()) {
		var ret string
		return ret
	}
	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject161) GetSubnetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject161) HasSubnet() bool {
	if o != nil && o.Subnet.IsSet() {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given NullableString and assigns it to the Subnet field.
func (o *InlineObject161) SetSubnet(v string) {
	o.Subnet.Set(&v)
}
// SetSubnetNil sets the value for Subnet to be an explicit nil
func (o *InlineObject161) SetSubnetNil() {
	o.Subnet.Set(nil)
}

// UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
func (o *InlineObject161) UnsetSubnet() {
	o.Subnet.Unset()
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject161) GetInterfaceIp() string {
	if o == nil || isNil(o.InterfaceIp.Get()) {
		var ret string
		return ret
	}
	return *o.InterfaceIp.Get()
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject161) GetInterfaceIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.InterfaceIp.Get(), o.InterfaceIp.IsSet()
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *InlineObject161) HasInterfaceIp() bool {
	if o != nil && o.InterfaceIp.IsSet() {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given NullableString and assigns it to the InterfaceIp field.
func (o *InlineObject161) SetInterfaceIp(v string) {
	o.InterfaceIp.Set(&v)
}
// SetInterfaceIpNil sets the value for InterfaceIp to be an explicit nil
func (o *InlineObject161) SetInterfaceIpNil() {
	o.InterfaceIp.Set(nil)
}

// UnsetInterfaceIp ensures that no value is present for InterfaceIp, not even an explicit nil
func (o *InlineObject161) UnsetInterfaceIp() {
	o.InterfaceIp.Unset()
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *InlineObject161) GetMulticastRouting() string {
	if o == nil || isNil(o.MulticastRouting) {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || isNil(o.MulticastRouting) {
    return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *InlineObject161) HasMulticastRouting() bool {
	if o != nil && !isNil(o.MulticastRouting) {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *InlineObject161) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject161) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId.Get()) {
		var ret int32
		return ret
	}
	return *o.VlanId.Get()
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject161) GetVlanIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.VlanId.Get(), o.VlanId.IsSet()
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineObject161) HasVlanId() bool {
	if o != nil && o.VlanId.IsSet() {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given NullableInt32 and assigns it to the VlanId field.
func (o *InlineObject161) SetVlanId(v int32) {
	o.VlanId.Set(&v)
}
// SetVlanIdNil sets the value for VlanId to be an explicit nil
func (o *InlineObject161) SetVlanIdNil() {
	o.VlanId.Set(nil)
}

// UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
func (o *InlineObject161) UnsetVlanId() {
	o.VlanId.Unset()
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *InlineObject161) GetDefaultGateway() string {
	if o == nil || isNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || isNil(o.DefaultGateway) {
    return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *InlineObject161) HasDefaultGateway() bool {
	if o != nil && !isNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *InlineObject161) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *InlineObject161) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesOspfSettings1 {
	if o == nil || isNil(o.OspfSettings) {
		var ret DevicesSerialSwitchRoutingInterfacesOspfSettings1
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesOspfSettings1, bool) {
	if o == nil || isNil(o.OspfSettings) {
    return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *InlineObject161) HasOspfSettings() bool {
	if o != nil && !isNil(o.OspfSettings) {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given DevicesSerialSwitchRoutingInterfacesOspfSettings1 and assigns it to the OspfSettings field.
func (o *InlineObject161) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesOspfSettings1) {
	o.OspfSettings = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineObject161) GetIpv6() DevicesSerialSwitchRoutingInterfacesIpv61 {
	if o == nil || isNil(o.Ipv6) {
		var ret DevicesSerialSwitchRoutingInterfacesIpv61
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesIpv61, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineObject161) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given DevicesSerialSwitchRoutingInterfacesIpv61 and assigns it to the Ipv6 field.
func (o *InlineObject161) SetIpv6(v DevicesSerialSwitchRoutingInterfacesIpv61) {
	o.Ipv6 = &v
}

func (o InlineObject161) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Subnet.IsSet() {
		toSerialize["subnet"] = o.Subnet.Get()
	}
	if o.InterfaceIp.IsSet() {
		toSerialize["interfaceIp"] = o.InterfaceIp.Get()
	}
	if !isNil(o.MulticastRouting) {
		toSerialize["multicastRouting"] = o.MulticastRouting
	}
	if o.VlanId.IsSet() {
		toSerialize["vlanId"] = o.VlanId.Get()
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

type NullableInlineObject161 struct {
	value *InlineObject161
	isSet bool
}

func (v NullableInlineObject161) Get() *InlineObject161 {
	return v.value
}

func (v *NullableInlineObject161) Set(val *InlineObject161) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject161) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject161(val *InlineObject161) *NullableInlineObject161 {
	return &NullableInlineObject161{value: val, isSet: true}
}

func (v NullableInlineObject161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


