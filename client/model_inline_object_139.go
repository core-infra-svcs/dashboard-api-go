/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject139 struct for InlineObject139
type InlineObject139 struct {
	// The mac address of the trusted server being added
	Mac string `json:"mac"`
	// The VLAN of the trusted server being added. It must be between 1 and 4094
	Vlan int32 `json:"vlan"`
	Ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 `json:"ipv4"`
}

// NewInlineObject139 instantiates a new InlineObject139 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject139(mac string, vlan int32, ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) *InlineObject139 {
	this := InlineObject139{}
	this.Mac = mac
	this.Vlan = vlan
	this.Ipv4 = ipv4
	return &this
}

// NewInlineObject139WithDefaults instantiates a new InlineObject139 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject139WithDefaults() *InlineObject139 {
	this := InlineObject139{}
	return &this
}

// GetMac returns the Mac field value
func (o *InlineObject139) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *InlineObject139) SetMac(v string) {
	o.Mac = v
}

// GetVlan returns the Vlan field value
func (o *InlineObject139) GetVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetVlanOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *InlineObject139) SetVlan(v int32) {
	o.Vlan = v
}

// GetIpv4 returns the Ipv4 field value
func (o *InlineObject139) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 {
	if o == nil {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41
		return ret
	}

	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value
// and a boolean to check if the value has been set.
func (o *InlineObject139) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ipv4, true
}

// SetIpv4 sets field value
func (o *InlineObject139) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) {
	o.Ipv4 = v
}

func (o InlineObject139) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mac"] = o.Mac
	}
	if true {
		toSerialize["vlan"] = o.Vlan
	}
	if true {
		toSerialize["ipv4"] = o.Ipv4
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject139 struct {
	value *InlineObject139
	isSet bool
}

func (v NullableInlineObject139) Get() *InlineObject139 {
	return v.value
}

func (v *NullableInlineObject139) Set(val *InlineObject139) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject139) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject139) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject139(val *InlineObject139) *NullableInlineObject139 {
	return &NullableInlineObject139{value: val, isSet: true}
}

func (v NullableInlineObject139) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject139) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


