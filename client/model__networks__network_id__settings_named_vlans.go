/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSettingsNamedVlans A hash of Named VLANs options applied to the Network.
type NetworksNetworkIdSettingsNamedVlans struct {
	// Enables / disables Named VLANs on the Network.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdSettingsNamedVlans instantiates a new NetworksNetworkIdSettingsNamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSettingsNamedVlans() *NetworksNetworkIdSettingsNamedVlans {
	this := NetworksNetworkIdSettingsNamedVlans{}
	return &this
}

// NewNetworksNetworkIdSettingsNamedVlansWithDefaults instantiates a new NetworksNetworkIdSettingsNamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSettingsNamedVlansWithDefaults() *NetworksNetworkIdSettingsNamedVlans {
	this := NetworksNetworkIdSettingsNamedVlans{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSettingsNamedVlans) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSettingsNamedVlans) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSettingsNamedVlans) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSettingsNamedVlans) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdSettingsNamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSettingsNamedVlans struct {
	value *NetworksNetworkIdSettingsNamedVlans
	isSet bool
}

func (v NullableNetworksNetworkIdSettingsNamedVlans) Get() *NetworksNetworkIdSettingsNamedVlans {
	return v.value
}

func (v *NullableNetworksNetworkIdSettingsNamedVlans) Set(val *NetworksNetworkIdSettingsNamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSettingsNamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSettingsNamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSettingsNamedVlans(val *NetworksNetworkIdSettingsNamedVlans) *NullableNetworksNetworkIdSettingsNamedVlans {
	return &NullableNetworksNetworkIdSettingsNamedVlans{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSettingsNamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSettingsNamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


