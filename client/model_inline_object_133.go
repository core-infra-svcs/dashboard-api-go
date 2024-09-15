/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject133 struct for InlineObject133
type InlineObject133 struct {
	// The mac address of the trusted server being added
	Mac string `json:"mac"`
	// The VLAN of the trusted server being added. It must be between 1 and 4094
	Vlan int32 `json:"vlan"`
	Ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 `json:"ipv4"`
}

// NewInlineObject133 instantiates a new InlineObject133 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject133(mac string, vlan int32, ipv4 NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) *InlineObject133 {
	this := InlineObject133{}
	this.Mac = mac
	this.Vlan = vlan
	this.Ipv4 = ipv4
	return &this
}

// NewInlineObject133WithDefaults instantiates a new InlineObject133 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject133WithDefaults() *InlineObject133 {
	this := InlineObject133{}
	return &this
}

// GetMac returns the Mac field value
func (o *InlineObject133) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *InlineObject133) GetMacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *InlineObject133) SetMac(v string) {
	o.Mac = v
}

// GetVlan returns the Vlan field value
func (o *InlineObject133) GetVlan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *InlineObject133) GetVlanOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vlan, true
}

// SetVlan sets field value
func (o *InlineObject133) SetVlan(v int32) {
	o.Vlan = v
}

// GetIpv4 returns the Ipv4 field value
func (o *InlineObject133) GetIpv4() NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41 {
	if o == nil {
		var ret NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41
		return ret
	}

	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value
// and a boolean to check if the value has been set.
func (o *InlineObject133) GetIpv4Ok() (*NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ipv4, true
}

// SetIpv4 sets field value
func (o *InlineObject133) SetIpv4(v NetworksNetworkIdSwitchDhcpServerPolicyArpInspectionTrustedServersIpv41) {
	o.Ipv4 = v
}

func (o InlineObject133) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject133 struct {
	value *InlineObject133
	isSet bool
}

func (v NullableInlineObject133) Get() *InlineObject133 {
	return v.value
}

func (v *NullableInlineObject133) Set(val *InlineObject133) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject133) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject133) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject133(val *InlineObject133) *NullableInlineObject133 {
	return &NullableInlineObject133{value: val, isSet: true}
}

func (v NullableInlineObject133) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject133) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


