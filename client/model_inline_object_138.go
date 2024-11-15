/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject138 struct for InlineObject138
type InlineObject138 struct {
	// The updated mac address of the trusted server
	Mac *string `json:"mac,omitempty"`
	// The updated VLAN of the trusted server. It must be between 1 and 4094
	Vlan *int32 `json:"vlan,omitempty"`
	Ipv4 *NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 `json:"ipv4,omitempty"`
}

// NewInlineObject138 instantiates a new InlineObject138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject138() *InlineObject138 {
	this := InlineObject138{}
	return &this
}

// NewInlineObject138WithDefaults instantiates a new InlineObject138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject138WithDefaults() *InlineObject138 {
	this := InlineObject138{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineObject138) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineObject138) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineObject138) SetMac(v string) {
	o.Mac = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *InlineObject138) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *InlineObject138) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *InlineObject138) SetVlan(v int32) {
	o.Vlan = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *InlineObject138) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject138) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *InlineObject138) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4 and assigns it to the Ipv4 field.
func (o *InlineObject138) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersTrustedServerIdIpv4) {
	o.Ipv4 = &v
}

func (o InlineObject138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
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


