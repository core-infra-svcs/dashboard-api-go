/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdPm25 PM2.5 concentration threshold. One of 'concentration' or 'quality' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdPm25 struct {
	// Alerting threshold as PM2.5 parts per million.
	Concentration *int32 `json:"concentration,omitempty"`
	// Alerting threshold as a qualitative PM2.5 level.
	Quality *string `json:"quality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdPm25 instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdPm25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdPm25() *NetworksNetworkIdSensorAlertsProfilesThresholdPm25 {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdPm25{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdPm25WithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdPm25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdPm25WithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdPm25 {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdPm25{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) SetConcentration(v int32) {
	o.Concentration = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) SetQuality(v string) {
	o.Quality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdPm25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25 struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdPm25
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdPm25 {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25(val *NetworksNetworkIdSensorAlertsProfilesThresholdPm25) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25 {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdPm25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


