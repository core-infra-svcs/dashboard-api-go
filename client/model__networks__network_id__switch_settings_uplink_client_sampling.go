/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchSettingsUplinkClientSampling Uplink client sampling
type NetworksNetworkIdSwitchSettingsUplinkClientSampling struct {
	// Enable uplink client sampling
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdSwitchSettingsUplinkClientSampling instantiates a new NetworksNetworkIdSwitchSettingsUplinkClientSampling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchSettingsUplinkClientSampling() *NetworksNetworkIdSwitchSettingsUplinkClientSampling {
	this := NetworksNetworkIdSwitchSettingsUplinkClientSampling{}
	return &this
}

// NewNetworksNetworkIdSwitchSettingsUplinkClientSamplingWithDefaults instantiates a new NetworksNetworkIdSwitchSettingsUplinkClientSampling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchSettingsUplinkClientSamplingWithDefaults() *NetworksNetworkIdSwitchSettingsUplinkClientSampling {
	this := NetworksNetworkIdSwitchSettingsUplinkClientSampling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchSettingsUplinkClientSampling) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchSettingsUplinkClientSampling) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchSettingsUplinkClientSampling) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSwitchSettingsUplinkClientSampling) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdSwitchSettingsUplinkClientSampling) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling struct {
	value *NetworksNetworkIdSwitchSettingsUplinkClientSampling
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling) Get() *NetworksNetworkIdSwitchSettingsUplinkClientSampling {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling) Set(val *NetworksNetworkIdSwitchSettingsUplinkClientSampling) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchSettingsUplinkClientSampling(val *NetworksNetworkIdSwitchSettingsUplinkClientSampling) *NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling {
	return &NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchSettingsUplinkClientSampling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


