/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdCurrent Electrical current threshold. 'level' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdCurrent struct {
	// Alerting threshold in amps. Must be between 0 and 15.
	Draw float32 `json:"draw"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdCurrent instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdCurrent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdCurrent(draw float32) *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdCurrent{}
	this.Draw = draw
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdCurrentWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdCurrent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdCurrentWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdCurrent{}
	return &this
}

// GetDraw returns the Draw field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent) GetDraw() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Draw
}

// GetDrawOk returns a tuple with the Draw field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent) GetDrawOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Draw, true
}

// SetDraw sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent) SetDraw(v float32) {
	o.Draw = v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdCurrent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["draw"] = o.Draw
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent(val *NetworksNetworkIdSensorAlertsProfilesThresholdCurrent) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCurrent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

