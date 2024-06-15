/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdFrequency Electrical frequency threshold. 'level' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdFrequency struct {
	// Alerting threshold in hertz. Must be between 0 and 60.
	Level float32 `json:"level"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdFrequency instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdFrequency(level float32) *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdFrequency{}
	this.Level = level
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdFrequencyWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdFrequencyWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdFrequency{}
	return &this
}

// GetLevel returns the Level field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency) GetLevel() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency) GetLevelOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency) SetLevel(v float32) {
	o.Level = v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdFrequency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency(val *NetworksNetworkIdSensorAlertsProfilesThresholdFrequency) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


