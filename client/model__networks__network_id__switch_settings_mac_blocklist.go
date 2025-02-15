/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSwitchSettingsMacBlocklist MAC blocklist
type NetworksNetworkIdSwitchSettingsMacBlocklist struct {
	// Enable MAC blocklist
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdSwitchSettingsMacBlocklist instantiates a new NetworksNetworkIdSwitchSettingsMacBlocklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchSettingsMacBlocklist() *NetworksNetworkIdSwitchSettingsMacBlocklist {
	this := NetworksNetworkIdSwitchSettingsMacBlocklist{}
	return &this
}

// NewNetworksNetworkIdSwitchSettingsMacBlocklistWithDefaults instantiates a new NetworksNetworkIdSwitchSettingsMacBlocklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchSettingsMacBlocklistWithDefaults() *NetworksNetworkIdSwitchSettingsMacBlocklist {
	this := NetworksNetworkIdSwitchSettingsMacBlocklist{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchSettingsMacBlocklist) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchSettingsMacBlocklist) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchSettingsMacBlocklist) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdSwitchSettingsMacBlocklist) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdSwitchSettingsMacBlocklist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchSettingsMacBlocklist struct {
	value *NetworksNetworkIdSwitchSettingsMacBlocklist
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchSettingsMacBlocklist) Get() *NetworksNetworkIdSwitchSettingsMacBlocklist {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchSettingsMacBlocklist) Set(val *NetworksNetworkIdSwitchSettingsMacBlocklist) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchSettingsMacBlocklist) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchSettingsMacBlocklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchSettingsMacBlocklist(val *NetworksNetworkIdSwitchSettingsMacBlocklist) *NullableNetworksNetworkIdSwitchSettingsMacBlocklist {
	return &NullableNetworksNetworkIdSwitchSettingsMacBlocklist{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchSettingsMacBlocklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchSettingsMacBlocklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


