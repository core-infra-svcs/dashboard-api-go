/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdWater Water detection threshold. 'present' must be provided and set to true.
type NetworksNetworkIdSensorAlertsProfilesThresholdWater struct {
	// Alerting threshold for a water detection event. Must be set to true.
	Present bool `json:"present"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdWater instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdWater object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdWater(present bool) *NetworksNetworkIdSensorAlertsProfilesThresholdWater {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdWater{}
	this.Present = present
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdWaterWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdWater object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdWaterWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdWater {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdWater{}
	return &this
}

// GetPresent returns the Present field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdWater) GetPresent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Present
}

// GetPresentOk returns a tuple with the Present field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdWater) GetPresentOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Present, true
}

// SetPresent sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdWater) SetPresent(v bool) {
	o.Present = v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdWater) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["present"] = o.Present
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdWater
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdWater {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdWater) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdWater(val *NetworksNetworkIdSensorAlertsProfilesThresholdWater) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdWater) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


