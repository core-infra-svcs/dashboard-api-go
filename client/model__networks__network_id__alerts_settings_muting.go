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

// NetworksNetworkIdAlertsSettingsMuting Mute alerts under certain conditions
type NetworksNetworkIdAlertsSettingsMuting struct {
	ByPortSchedules *NetworksNetworkIdAlertsSettingsMutingByPortSchedules `json:"byPortSchedules,omitempty"`
}

// NewNetworksNetworkIdAlertsSettingsMuting instantiates a new NetworksNetworkIdAlertsSettingsMuting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdAlertsSettingsMuting() *NetworksNetworkIdAlertsSettingsMuting {
	this := NetworksNetworkIdAlertsSettingsMuting{}
	return &this
}

// NewNetworksNetworkIdAlertsSettingsMutingWithDefaults instantiates a new NetworksNetworkIdAlertsSettingsMuting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdAlertsSettingsMutingWithDefaults() *NetworksNetworkIdAlertsSettingsMuting {
	this := NetworksNetworkIdAlertsSettingsMuting{}
	return &this
}

// GetByPortSchedules returns the ByPortSchedules field value if set, zero value otherwise.
func (o *NetworksNetworkIdAlertsSettingsMuting) GetByPortSchedules() NetworksNetworkIdAlertsSettingsMutingByPortSchedules {
	if o == nil || isNil(o.ByPortSchedules) {
		var ret NetworksNetworkIdAlertsSettingsMutingByPortSchedules
		return ret
	}
	return *o.ByPortSchedules
}

// GetByPortSchedulesOk returns a tuple with the ByPortSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdAlertsSettingsMuting) GetByPortSchedulesOk() (*NetworksNetworkIdAlertsSettingsMutingByPortSchedules, bool) {
	if o == nil || isNil(o.ByPortSchedules) {
    return nil, false
	}
	return o.ByPortSchedules, true
}

// HasByPortSchedules returns a boolean if a field has been set.
func (o *NetworksNetworkIdAlertsSettingsMuting) HasByPortSchedules() bool {
	if o != nil && !isNil(o.ByPortSchedules) {
		return true
	}

	return false
}

// SetByPortSchedules gets a reference to the given NetworksNetworkIdAlertsSettingsMutingByPortSchedules and assigns it to the ByPortSchedules field.
func (o *NetworksNetworkIdAlertsSettingsMuting) SetByPortSchedules(v NetworksNetworkIdAlertsSettingsMutingByPortSchedules) {
	o.ByPortSchedules = &v
}

func (o NetworksNetworkIdAlertsSettingsMuting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByPortSchedules) {
		toSerialize["byPortSchedules"] = o.ByPortSchedules
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdAlertsSettingsMuting struct {
	value *NetworksNetworkIdAlertsSettingsMuting
	isSet bool
}

func (v NullableNetworksNetworkIdAlertsSettingsMuting) Get() *NetworksNetworkIdAlertsSettingsMuting {
	return v.value
}

func (v *NullableNetworksNetworkIdAlertsSettingsMuting) Set(val *NetworksNetworkIdAlertsSettingsMuting) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdAlertsSettingsMuting) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdAlertsSettingsMuting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdAlertsSettingsMuting(val *NetworksNetworkIdAlertsSettingsMuting) *NullableNetworksNetworkIdAlertsSettingsMuting {
	return &NullableNetworksNetworkIdAlertsSettingsMuting{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdAlertsSettingsMuting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdAlertsSettingsMuting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


