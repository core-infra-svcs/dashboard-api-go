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

// NetworksNetworkIdApplianceVlansIpv6 IPv6 configuration on the VLAN
type NetworksNetworkIdApplianceVlansIpv6 struct {
	// Enable IPv6 on VLAN
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix assignments on the VLAN
	PrefixAssignments []NetworksNetworkIdApplianceVlansIpv6PrefixAssignments `json:"prefixAssignments,omitempty"`
}

// NewNetworksNetworkIdApplianceVlansIpv6 instantiates a new NetworksNetworkIdApplianceVlansIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceVlansIpv6() *NetworksNetworkIdApplianceVlansIpv6 {
	this := NetworksNetworkIdApplianceVlansIpv6{}
	return &this
}

// NewNetworksNetworkIdApplianceVlansIpv6WithDefaults instantiates a new NetworksNetworkIdApplianceVlansIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceVlansIpv6WithDefaults() *NetworksNetworkIdApplianceVlansIpv6 {
	this := NetworksNetworkIdApplianceVlansIpv6{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVlansIpv6) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansIpv6) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVlansIpv6) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdApplianceVlansIpv6) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefixAssignments returns the PrefixAssignments field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceVlansIpv6) GetPrefixAssignments() []NetworksNetworkIdApplianceVlansIpv6PrefixAssignments {
	if o == nil || isNil(o.PrefixAssignments) {
		var ret []NetworksNetworkIdApplianceVlansIpv6PrefixAssignments
		return ret
	}
	return o.PrefixAssignments
}

// GetPrefixAssignmentsOk returns a tuple with the PrefixAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceVlansIpv6) GetPrefixAssignmentsOk() ([]NetworksNetworkIdApplianceVlansIpv6PrefixAssignments, bool) {
	if o == nil || isNil(o.PrefixAssignments) {
    return nil, false
	}
	return o.PrefixAssignments, true
}

// HasPrefixAssignments returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceVlansIpv6) HasPrefixAssignments() bool {
	if o != nil && !isNil(o.PrefixAssignments) {
		return true
	}

	return false
}

// SetPrefixAssignments gets a reference to the given []NetworksNetworkIdApplianceVlansIpv6PrefixAssignments and assigns it to the PrefixAssignments field.
func (o *NetworksNetworkIdApplianceVlansIpv6) SetPrefixAssignments(v []NetworksNetworkIdApplianceVlansIpv6PrefixAssignments) {
	o.PrefixAssignments = v
}

func (o NetworksNetworkIdApplianceVlansIpv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.PrefixAssignments) {
		toSerialize["prefixAssignments"] = o.PrefixAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceVlansIpv6 struct {
	value *NetworksNetworkIdApplianceVlansIpv6
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceVlansIpv6) Get() *NetworksNetworkIdApplianceVlansIpv6 {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceVlansIpv6) Set(val *NetworksNetworkIdApplianceVlansIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceVlansIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceVlansIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceVlansIpv6(val *NetworksNetworkIdApplianceVlansIpv6) *NullableNetworksNetworkIdApplianceVlansIpv6 {
	return &NullableNetworksNetworkIdApplianceVlansIpv6{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceVlansIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceVlansIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


