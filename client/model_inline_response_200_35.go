/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20035 struct for InlineResponse20035
type InlineResponse20035 struct {
	// The id
	InterfaceId *string `json:"interfaceId,omitempty"`
	// The name
	Name *string `json:"name,omitempty"`
	// IPv4 subnet
	Subnet *string `json:"subnet,omitempty"`
	// IPv4 address
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// Multicast routing status
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// VLAN id
	VlanId *int32 `json:"vlanId,omitempty"`
	OspfSettings *DevicesSerialSwitchRoutingInterfacesOspfSettings `json:"ospfSettings,omitempty"`
	OspfV3 *DevicesSerialSwitchRoutingInterfacesOspfV3 `json:"ospfV3,omitempty"`
	Ipv6 *DevicesSerialSwitchRoutingInterfacesIpv6 `json:"ipv6,omitempty"`
	// Whether this is the switch's IPv4 uplink
	UplinkV4 *bool `json:"uplinkV4,omitempty"`
	// Whether this is the switch's IPv6 uplink
	UplinkV6 *bool `json:"uplinkV6,omitempty"`
	// IPv4 default gateway
	DefaultGateway *string `json:"defaultGateway,omitempty"`
}

// NewInlineResponse20035 instantiates a new InlineResponse20035 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20035() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// NewInlineResponse20035WithDefaults instantiates a new InlineResponse20035 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20035WithDefaults() *InlineResponse20035 {
	this := InlineResponse20035{}
	return &this
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *InlineResponse20035) GetInterfaceId() string {
	if o == nil || isNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetInterfaceIdOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceId) {
    return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *InlineResponse20035) HasInterfaceId() bool {
	if o != nil && !isNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *InlineResponse20035) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20035) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20035) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20035) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20035) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20035) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20035) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *InlineResponse20035) GetInterfaceIp() string {
	if o == nil || isNil(o.InterfaceIp) {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetInterfaceIpOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceIp) {
    return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *InlineResponse20035) HasInterfaceIp() bool {
	if o != nil && !isNil(o.InterfaceIp) {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *InlineResponse20035) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *InlineResponse20035) GetMulticastRouting() string {
	if o == nil || isNil(o.MulticastRouting) {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || isNil(o.MulticastRouting) {
    return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *InlineResponse20035) HasMulticastRouting() bool {
	if o != nil && !isNil(o.MulticastRouting) {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *InlineResponse20035) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *InlineResponse20035) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *InlineResponse20035) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *InlineResponse20035) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *InlineResponse20035) GetOspfSettings() DevicesSerialSwitchRoutingInterfacesOspfSettings {
	if o == nil || isNil(o.OspfSettings) {
		var ret DevicesSerialSwitchRoutingInterfacesOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetOspfSettingsOk() (*DevicesSerialSwitchRoutingInterfacesOspfSettings, bool) {
	if o == nil || isNil(o.OspfSettings) {
    return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *InlineResponse20035) HasOspfSettings() bool {
	if o != nil && !isNil(o.OspfSettings) {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given DevicesSerialSwitchRoutingInterfacesOspfSettings and assigns it to the OspfSettings field.
func (o *InlineResponse20035) SetOspfSettings(v DevicesSerialSwitchRoutingInterfacesOspfSettings) {
	o.OspfSettings = &v
}

// GetOspfV3 returns the OspfV3 field value if set, zero value otherwise.
func (o *InlineResponse20035) GetOspfV3() DevicesSerialSwitchRoutingInterfacesOspfV3 {
	if o == nil || isNil(o.OspfV3) {
		var ret DevicesSerialSwitchRoutingInterfacesOspfV3
		return ret
	}
	return *o.OspfV3
}

// GetOspfV3Ok returns a tuple with the OspfV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetOspfV3Ok() (*DevicesSerialSwitchRoutingInterfacesOspfV3, bool) {
	if o == nil || isNil(o.OspfV3) {
    return nil, false
	}
	return o.OspfV3, true
}

// HasOspfV3 returns a boolean if a field has been set.
func (o *InlineResponse20035) HasOspfV3() bool {
	if o != nil && !isNil(o.OspfV3) {
		return true
	}

	return false
}

// SetOspfV3 gets a reference to the given DevicesSerialSwitchRoutingInterfacesOspfV3 and assigns it to the OspfV3 field.
func (o *InlineResponse20035) SetOspfV3(v DevicesSerialSwitchRoutingInterfacesOspfV3) {
	o.OspfV3 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineResponse20035) GetIpv6() DevicesSerialSwitchRoutingInterfacesIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret DevicesSerialSwitchRoutingInterfacesIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetIpv6Ok() (*DevicesSerialSwitchRoutingInterfacesIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineResponse20035) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given DevicesSerialSwitchRoutingInterfacesIpv6 and assigns it to the Ipv6 field.
func (o *InlineResponse20035) SetIpv6(v DevicesSerialSwitchRoutingInterfacesIpv6) {
	o.Ipv6 = &v
}

// GetUplinkV4 returns the UplinkV4 field value if set, zero value otherwise.
func (o *InlineResponse20035) GetUplinkV4() bool {
	if o == nil || isNil(o.UplinkV4) {
		var ret bool
		return ret
	}
	return *o.UplinkV4
}

// GetUplinkV4Ok returns a tuple with the UplinkV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetUplinkV4Ok() (*bool, bool) {
	if o == nil || isNil(o.UplinkV4) {
    return nil, false
	}
	return o.UplinkV4, true
}

// HasUplinkV4 returns a boolean if a field has been set.
func (o *InlineResponse20035) HasUplinkV4() bool {
	if o != nil && !isNil(o.UplinkV4) {
		return true
	}

	return false
}

// SetUplinkV4 gets a reference to the given bool and assigns it to the UplinkV4 field.
func (o *InlineResponse20035) SetUplinkV4(v bool) {
	o.UplinkV4 = &v
}

// GetUplinkV6 returns the UplinkV6 field value if set, zero value otherwise.
func (o *InlineResponse20035) GetUplinkV6() bool {
	if o == nil || isNil(o.UplinkV6) {
		var ret bool
		return ret
	}
	return *o.UplinkV6
}

// GetUplinkV6Ok returns a tuple with the UplinkV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetUplinkV6Ok() (*bool, bool) {
	if o == nil || isNil(o.UplinkV6) {
    return nil, false
	}
	return o.UplinkV6, true
}

// HasUplinkV6 returns a boolean if a field has been set.
func (o *InlineResponse20035) HasUplinkV6() bool {
	if o != nil && !isNil(o.UplinkV6) {
		return true
	}

	return false
}

// SetUplinkV6 gets a reference to the given bool and assigns it to the UplinkV6 field.
func (o *InlineResponse20035) SetUplinkV6(v bool) {
	o.UplinkV6 = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *InlineResponse20035) GetDefaultGateway() string {
	if o == nil || isNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20035) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || isNil(o.DefaultGateway) {
    return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *InlineResponse20035) HasDefaultGateway() bool {
	if o != nil && !isNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *InlineResponse20035) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

func (o InlineResponse20035) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !isNil(o.Name) {
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
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !isNil(o.OspfSettings) {
		toSerialize["ospfSettings"] = o.OspfSettings
	}
	if !isNil(o.OspfV3) {
		toSerialize["ospfV3"] = o.OspfV3
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if !isNil(o.UplinkV4) {
		toSerialize["uplinkV4"] = o.UplinkV4
	}
	if !isNil(o.UplinkV6) {
		toSerialize["uplinkV6"] = o.UplinkV6
	}
	if !isNil(o.DefaultGateway) {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20035 struct {
	value *InlineResponse20035
	isSet bool
}

func (v NullableInlineResponse20035) Get() *InlineResponse20035 {
	return v.value
}

func (v *NullableInlineResponse20035) Set(val *InlineResponse20035) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20035) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20035) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20035(val *InlineResponse20035) *NullableInlineResponse20035 {
	return &NullableInlineResponse20035{value: val, isSet: true}
}

func (v NullableInlineResponse20035) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20035) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


