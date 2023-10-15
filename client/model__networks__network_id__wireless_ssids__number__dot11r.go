/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberDot11r The current setting for 802.11r
type NetworksNetworkIdWirelessSsidsNumberDot11r struct {
	// Whether 802.11r is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// (Optional) Whether 802.11r is adaptive or not.
	Adaptive *bool `json:"adaptive,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberDot11r instantiates a new NetworksNetworkIdWirelessSsidsNumberDot11r object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberDot11r() *NetworksNetworkIdWirelessSsidsNumberDot11r {
	this := NetworksNetworkIdWirelessSsidsNumberDot11r{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberDot11rWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberDot11r object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberDot11rWithDefaults() *NetworksNetworkIdWirelessSsidsNumberDot11r {
	this := NetworksNetworkIdWirelessSsidsNumberDot11r{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAdaptive returns the Adaptive field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) GetAdaptive() bool {
	if o == nil || isNil(o.Adaptive) {
		var ret bool
		return ret
	}
	return *o.Adaptive
}

// GetAdaptiveOk returns a tuple with the Adaptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) GetAdaptiveOk() (*bool, bool) {
	if o == nil || isNil(o.Adaptive) {
    return nil, false
	}
	return o.Adaptive, true
}

// HasAdaptive returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) HasAdaptive() bool {
	if o != nil && !isNil(o.Adaptive) {
		return true
	}

	return false
}

// SetAdaptive gets a reference to the given bool and assigns it to the Adaptive field.
func (o *NetworksNetworkIdWirelessSsidsNumberDot11r) SetAdaptive(v bool) {
	o.Adaptive = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberDot11r) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Adaptive) {
		toSerialize["adaptive"] = o.Adaptive
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberDot11r struct {
	value *NetworksNetworkIdWirelessSsidsNumberDot11r
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDot11r) Get() *NetworksNetworkIdWirelessSsidsNumberDot11r {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDot11r) Set(val *NetworksNetworkIdWirelessSsidsNumberDot11r) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDot11r) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDot11r) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberDot11r(val *NetworksNetworkIdWirelessSsidsNumberDot11r) *NullableNetworksNetworkIdWirelessSsidsNumberDot11r {
	return &NullableNetworksNetworkIdWirelessSsidsNumberDot11r{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberDot11r) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberDot11r) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


