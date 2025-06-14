/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdCo2 CO2 concentration threshold. One of 'concentration' or 'quality' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdCo2 struct {
	// Alerting threshold as CO2 parts per million.
	Concentration *int32 `json:"concentration,omitempty"`
	// Alerting threshold as a qualitative CO2 level.
	Quality *string `json:"quality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdCo2 instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdCo2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdCo2() *NetworksNetworkIdSensorAlertsProfilesThresholdCo2 {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdCo2{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdCo2WithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdCo2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdCo2WithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdCo2 {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdCo2{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) SetConcentration(v int32) {
	o.Concentration = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) SetQuality(v string) {
	o.Quality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdCo2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2 struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdCo2
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdCo2 {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2(val *NetworksNetworkIdSensorAlertsProfilesThresholdCo2) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2 {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdCo2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


