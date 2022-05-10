/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject15 struct for InlineObject15
type InlineObject15 struct {
	// A friendly name or description for the interface or VLAN.
	Name string `json:"name"`
	// The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet *string `json:"subnet,omitempty"`
	// The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// The VLAN this routed interface is on. VLAN must be between 1 and 4094.
	VlanId int32 `json:"vlanId"`
	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
	DefaultGateway *string `json:"defaultGateway,omitempty"`
	OspfSettings *DevicesSerialSwitchRoutingInterfacesOspfSettings `json:"ospfSettings,omitempty"`
	Ipv6 *DevicesSerialSwitchRoutingInterfacesIpv6 `json:"ipv6,omitempty"`
}

// NewInlineObject15 instantiates a new InlineObject15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject15(name string, vlanId int32) *InlineObject15 {
	this := InlineObject15{}
	this.Name = name
	this.VlanId = vlanId
	return &this
}

// NewInlineObject15WithDefaults instantiates a new InlineObject15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject15WithDefaults() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject15) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject15) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject15) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject15) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject15) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *InlineObject15) GetInterfaceIp() string {
	if o == nil || o.InterfaceIp == nil {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetInterfaceIpOk() (*string, bool) {
	if o == nil || o.InterfaceIp == nil {
		return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *InlineObject15) HasInterfaceIp() bool {
	if o != nil && o.InterfaceIp != nil {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *InlineObject15) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *InlineObject15) GetMulticastRouting() string {
	if o == nil || o.MulticastRouting == nil {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || o.MulticastRouting == nil {
		return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *InlineObject15) HasMulticastRouting() bool {
	if o != nil && o.MulticastRouting != nil {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *InlineObject15) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value
func (o *InlineObject15) GetVlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetVlanIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *InlineObject15) SetVlanId(v int32) {
	o.VlanId = v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *InlineObject15) GetDefaultGateway() string {
	if o == nil || o.DefaultGateway == nil {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || o.DefaultGateway == nil {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *InlineObject15) HasDefaultGateway() bool {
	if o != nil && o.DefaultGateway != nil {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *InlineObject15) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *InlineObject15) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesOspfSettings {
	if o == nil || o.OspfSettings == nil {
		var ret DevicesSerialSwitchRoutingInterfacesOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesOspfSettings, bool) {
	if o == nil || o.OspfSettings == nil {
		return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *InlineObject15) HasOspfSettings() bool {
	if o != nil && o.OspfSettings != nil {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given DevicesSerialSwitchRoutingInterfacesOspfSettings and assigns it to the OspfSettings field.
func (o *InlineObject15) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesOspfSettings) {
	o.OspfSettings = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineObject15) GetIpv6() DevicesSerialSwitchRoutingInterfacesIpv6 {
	if o == nil || o.Ipv6 == nil {
		var ret DevicesSerialSwitchRoutingInterfacesIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesIpv6, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineObject15) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given DevicesSerialSwitchRoutingInterfacesIpv6 and assigns it to the Ipv6 field.
func (o *InlineObject15) SetIpv6(v DevicesSerialSwitchRoutingInterfacesIpv6) {
	o.Ipv6 = &v
}

func (o InlineObject15) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.InterfaceIp != nil {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if o.MulticastRouting != nil {
		toSerialize["multicastRouting"] = o.MulticastRouting
	}
	if true {
		toSerialize["vlanId"] = o.VlanId
	}
	if o.DefaultGateway != nil {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if o.OspfSettings != nil {
		toSerialize["ospfSettings"] = o.OspfSettings
	}
	if o.Ipv6 != nil {
		toSerialize["ipv6"] = o.Ipv6
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject15 struct {
	value *InlineObject15
	isSet bool
}

func (v NullableInlineObject15) Get() *InlineObject15 {
	return v.value
}

func (v *NullableInlineObject15) Set(val *InlineObject15) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject15) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject15(val *InlineObject15) *NullableInlineObject15 {
	return &NullableInlineObject15{value: val, isSet: true}
}

func (v NullableInlineObject15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


