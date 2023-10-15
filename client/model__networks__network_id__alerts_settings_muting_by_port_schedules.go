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

// NetworksNetworkIdAlertsSettingsMutingByPortSchedules Mute wireless unreachable alerts based on switch port schedules
type NetworksNetworkIdAlertsSettingsMutingByPortSchedules struct {
	// If true, then wireless unreachable alerts will be muted when caused by a port schedule
	Enabled *bool `json:"enabled,omitempty"`
}

// NewNetworksNetworkIdAlertsSettingsMutingByPortSchedules instantiates a new NetworksNetworkIdAlertsSettingsMutingByPortSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsSettingsMutingByPortSchedules() *NetworksNetworkIdAlertsSettingsMutingByPortSchedules {
	this := NetworksNetworkIdAlertsSettingsMutingByPortSchedules{}
	return &this
}

// NewNetworksNetworkIdAlertsSettingsMutingByPortSchedulesWithDefaults instantiates a new NetworksNetworkIdAlertsSettingsMutingByPortSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsSettingsMutingByPortSchedulesWithDefaults() *NetworksNetworkIdAlertsSettingsMutingByPortSchedules {
	this := NetworksNetworkIdAlertsSettingsMutingByPortSchedules{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsSettingsMutingByPortSchedules) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsSettingsMutingByPortSchedules) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsSettingsMutingByPortSchedules) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworksNetworkIdAlertsSettingsMutingByPortSchedules) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworksNetworkIdAlertsSettingsMutingByPortSchedules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules struct {
	value *NetworksNetworkIdAlertsSettingsMutingByPortSchedules
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules) Get() *NetworksNetworkIdAlertsSettingsMutingByPortSchedules {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules) Set(val *NetworksNetworkIdAlertsSettingsMutingByPortSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules(val *NetworksNetworkIdAlertsSettingsMutingByPortSchedules) *NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules {
	return &NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsSettingsMutingByPortSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


