/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdTvoc TVOC concentration threshold. One of 'concentration' or 'quality' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdTvoc struct {
	// Alerting threshold as TVOC micrograms per cubic meter.
	Concentration *int32 `json:"concentration,omitempty"`
	// Alerting threshold as a qualitative TVOC level.
	Quality *string `json:"quality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdTvoc instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdTvoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdTvoc() *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdTvoc{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdTvocWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdTvoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdTvocWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdTvoc{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) SetConcentration(v int32) {
	o.Concentration = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) SetQuality(v string) {
	o.Quality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc(val *NetworksNetworkIdSensorAlertsProfilesThresholdTvoc) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTvoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


