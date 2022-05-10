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

// InlineObject16 struct for InlineObject16
type InlineObject16 struct {
	// A friendly name or description for the interface or VLAN.
	Name *string `json:"name,omitempty"`
	// The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet *string `json:"subnet,omitempty"`
	// The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'.
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// The VLAN this routed interface is on. VLAN must be between 1 and 4094.
	VlanId *int32 `json:"vlanId,omitempty"`
	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
	DefaultGateway *string `json:"defaultGateway,omitempty"`
	OspfSettings *DevicesSerialSwitchRoutingInterfacesInterfaceIdOspfSettings `json:"ospfSettings,omitempty"`
	Ipv6 *DevicesSerialSwitchRoutingInterfacesInterfaceIdIpv6 `json:"ipv6,omitempty"`
}

// NewInlineObject16 instantiates a new InlineObject16 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// NewInlineObject16WithDefaults instantiates a new InlineObject16 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16WithDefaults() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject16) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject16) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject16) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject16) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject16) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject16) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *InlineObject16) GetInterfaceIp() string {
	if o == nil || o.InterfaceIp == nil {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetInterfaceIpOk() (*string, bool) {
	if o == nil || o.InterfaceIp == nil {
		return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *InlineObject16) HasInterfaceIp() bool {
	if o != nil && o.InterfaceIp != nil {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *InlineObject16) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *InlineObject16) GetMulticastRouting() string {
	if o == nil || o.MulticastRouting == nil {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || o.MulticastRouting == nil {
		return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *InlineObject16) HasMulticastRouting() bool {
	if o != nil && o.MulticastRouting != nil {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *InlineObject16) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineObject16) GetVlanId() int32 {
	if o == nil || o.VlanId == nil {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetVlanIdOk() (*int32, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineObject16) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineObject16) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *InlineObject16) GetDefaultGateway() string {
	if o == nil || o.DefaultGateway == nil {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || o.DefaultGateway == nil {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *InlineObject16) HasDefaultGateway() bool {
	if o != nil && o.DefaultGateway != nil {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *InlineObject16) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *InlineObject16) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesInterfaceIdOspfSettings {
	if o == nil || o.OspfSettings == nil {
		var ret DevicesSerialSwitchRoutingInterfacesInterfaceIdOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesInterfaceIdOspfSettings, bool) {
	if o == nil || o.OspfSettings == nil {
		return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *InlineObject16) HasOspfSettings() bool {
	if o != nil && o.OspfSettings != nil {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given DevicesSerialSwitchRoutingInterfacesInterfaceIdOspfSettings and assigns it to the OspfSettings field.
func (o *InlineObject16) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesInterfaceIdOspfSettings) {
	o.OspfSettings = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineObject16) GetIpv6() DevicesSerialSwitchRoutingInterfacesInterfaceIdIpv6 {
	if o == nil || o.Ipv6 == nil {
		var ret DevicesSerialSwitchRoutingInterfacesInterfaceIdIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesInterfaceIdIpv6, bool) {
	if o == nil || o.Ipv6 == nil {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineObject16) HasIpv6() bool {
	if o != nil && o.Ipv6 != nil {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given DevicesSerialSwitchRoutingInterfacesInterfaceIdIpv6 and assigns it to the Ipv6 field.
func (o *InlineObject16) SetIpv6(v DevicesSerialSwitchRoutingInterfacesInterfaceIdIpv6) {
	o.Ipv6 = &v
}

func (o InlineObject16) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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
	if o.VlanId != nil {
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

type NullableInlineObject16 struct {
	value *InlineObject16
	isSet bool
}

func (v NullableInlineObject16) Get() *InlineObject16 {
	return v.value
}

func (v *NullableInlineObject16) Set(val *InlineObject16) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16(val *InlineObject16) *NullableInlineObject16 {
	return &NullableInlineObject16{value: val, isSet: true}
}

func (v NullableInlineObject16) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


