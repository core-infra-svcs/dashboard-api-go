/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdRealPower Real power threshold. 'draw' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdRealPower struct {
	// Alerting threshold in watts. Must be between 0 and 3750.
	Draw float32 `json:"draw"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdRealPower instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdRealPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdRealPower(draw float32) *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdRealPower{}
	this.Draw = draw
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdRealPowerWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdRealPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdRealPowerWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdRealPower{}
	return &this
}

// GetDraw returns the Draw field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower) GetDraw() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Draw
}

// GetDrawOk returns a tuple with the Draw field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower) GetDrawOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Draw, true
}

// SetDraw sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower) SetDraw(v float32) {
	o.Draw = v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdRealPower) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["draw"] = o.Draw
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower(val *NetworksNetworkIdSensorAlertsProfilesThresholdRealPower) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdRealPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


