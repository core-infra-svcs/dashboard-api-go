/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdVoltage Voltage threshold. 'level' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdVoltage struct {
	// Alerting threshold in volts. Must be between 0 and 250.
	Level float32 `json:"level"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdVoltage instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdVoltage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdVoltage(level float32) *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdVoltage{}
	this.Level = level
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdVoltageWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdVoltage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdVoltageWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdVoltage{}
	return &this
}

// GetLevel returns the Level field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage) GetLevel() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage) GetLevelOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage) SetLevel(v float32) {
	o.Level = v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdVoltage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage(val *NetworksNetworkIdSensorAlertsProfilesThresholdVoltage) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdVoltage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


