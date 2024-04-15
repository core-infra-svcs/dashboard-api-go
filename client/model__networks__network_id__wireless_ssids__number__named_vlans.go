/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberNamedVlans Named VLAN settings.
type NetworksNetworkIdWirelessSsidsNumberNamedVlans struct {
	Tagging *NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging `json:"tagging,omitempty"`
	Radius *NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius `json:"radius,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlans instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlans() *NetworksNetworkIdWirelessSsidsNumberNamedVlans {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlans{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberNamedVlansWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberNamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberNamedVlansWithDefaults() *NetworksNetworkIdWirelessSsidsNumberNamedVlans {
	this := NetworksNetworkIdWirelessSsidsNumberNamedVlans{}
	return &this
}

// GetTagging returns the Tagging field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) GetTagging() NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging {
	if o == nil || isNil(o.Tagging) {
		var ret NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging
		return ret
	}
	return *o.Tagging
}

// GetTaggingOk returns a tuple with the Tagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) GetTaggingOk() (*NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging, bool) {
	if o == nil || isNil(o.Tagging) {
    return nil, false
	}
	return o.Tagging, true
}

// HasTagging returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) HasTagging() bool {
	if o != nil && !isNil(o.Tagging) {
		return true
	}

	return false
}

// SetTagging gets a reference to the given NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging and assigns it to the Tagging field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) SetTagging(v NetworksNetworkIdWirelessSsidsNumberNamedVlansTagging) {
	o.Tagging = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) GetRadius() NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius {
	if o == nil || isNil(o.Radius) {
		var ret NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) GetRadiusOk() (*NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius, bool) {
	if o == nil || isNil(o.Radius) {
    return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) HasRadius() bool {
	if o != nil && !isNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius and assigns it to the Radius field.
func (o *NetworksNetworkIdWirelessSsidsNumberNamedVlans) SetRadius(v NetworksNetworkIdWirelessSsidsNumberNamedVlansRadius) {
	o.Radius = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberNamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tagging) {
		toSerialize["tagging"] = o.Tagging
	}
	if !isNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans struct {
	value *NetworksNetworkIdWirelessSsidsNumberNamedVlans
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans) Get() *NetworksNetworkIdWirelessSsidsNumberNamedVlans {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans) Set(val *NetworksNetworkIdWirelessSsidsNumberNamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberNamedVlans(val *NetworksNetworkIdWirelessSsidsNumberNamedVlans) *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans {
	return &NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberNamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


