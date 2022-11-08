/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSingleLanIpv6 IPv6 configuration on the VLAN
type NetworksNetworkIdApplianceSingleLanIpv6 struct {
	// Enable IPv6 on VLAN.
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix assignments on the VLAN
	PrefixAssignments []NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments `json:"prefixAssignments,omitempty"`
}

// NewNetworksNetworkIdApplianceSingleLanIpv6 instantiates a new NetworksNetworkIdApplianceSingleLanIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSingleLanIpv6() *NetworksNetworkIdApplianceSingleLanIpv6 {
	this := NetworksNetworkIdApplianceSingleLanIpv6{}
	return &this
}

// NewNetworksNetworkIdApplianceSingleLanIpv6WithDefaults instantiates a new NetworksNetworkIdApplianceSingleLanIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSingleLanIpv6WithDefaults() *NetworksNetworkIdApplianceSingleLanIpv6 {
	this := NetworksNetworkIdApplianceSingleLanIpv6{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefixAssignments returns the PrefixAssignments field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) GetPrefixAssignments() []NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments {
	if o == nil || isNil(o.PrefixAssignments) {
		var ret []NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments
		return ret
	}
	return o.PrefixAssignments
}

// GetPrefixAssignmentsOk returns a tuple with the PrefixAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) GetPrefixAssignmentsOk() ([]NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments, bool) {
	if o == nil || isNil(o.PrefixAssignments) {
    return nil, false
	}
	return o.PrefixAssignments, true
}

// HasPrefixAssignments returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) HasPrefixAssignments() bool {
	if o != nil && !isNil(o.PrefixAssignments) {
		return true
	}

	return false
}

// SetPrefixAssignments gets a reference to the given []NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments and assigns it to the PrefixAssignments field.
func (o *NetworksNetworkIdApplianceSingleLanIpv6) SetPrefixAssignments(v []NetworksNetworkIdApplianceSingleLanIpv6PrefixAssignments) {
	o.PrefixAssignments = v
}

func (o NetworksNetworkIdApplianceSingleLanIpv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.PrefixAssignments) {
		toSerialize["prefixAssignments"] = o.PrefixAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSingleLanIpv6 struct {
	value *NetworksNetworkIdApplianceSingleLanIpv6
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6) Get() *NetworksNetworkIdApplianceSingleLanIpv6 {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6) Set(val *NetworksNetworkIdApplianceSingleLanIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSingleLanIpv6(val *NetworksNetworkIdApplianceSingleLanIpv6) *NullableNetworksNetworkIdApplianceSingleLanIpv6 {
	return &NullableNetworksNetworkIdApplianceSingleLanIpv6{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSingleLanIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSingleLanIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

