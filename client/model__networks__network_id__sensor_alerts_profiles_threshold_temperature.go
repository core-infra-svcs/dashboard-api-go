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

// NetworksNetworkIdSensorAlertsProfilesThresholdTemperature Temperature threshold. One of 'celsius', 'fahrenheit', or 'quality' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdTemperature struct {
	// Alerting threshold in degrees Celsius.
	Celsius *float32 `json:"celsius,omitempty"`
	// Alerting threshold in degrees Fahrenheit.
	Fahrenheit *float32 `json:"fahrenheit,omitempty"`
	// Alerting threshold as a qualitative temperature level.
	Quality *string `json:"quality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdTemperature instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdTemperature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdTemperature() *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdTemperature{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdTemperatureWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdTemperature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdTemperatureWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdTemperature{}
	return &this
}

// GetCelsius returns the Celsius field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) GetCelsius() float32 {
	if o == nil || isNil(o.Celsius) {
		var ret float32
		return ret
	}
	return *o.Celsius
}

// GetCelsiusOk returns a tuple with the Celsius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) GetCelsiusOk() (*float32, bool) {
	if o == nil || isNil(o.Celsius) {
    return nil, false
	}
	return o.Celsius, true
}

// HasCelsius returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) HasCelsius() bool {
	if o != nil && !isNil(o.Celsius) {
		return true
	}

	return false
}

// SetCelsius gets a reference to the given float32 and assigns it to the Celsius field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) SetCelsius(v float32) {
	o.Celsius = &v
}

// GetFahrenheit returns the Fahrenheit field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) GetFahrenheit() float32 {
	if o == nil || isNil(o.Fahrenheit) {
		var ret float32
		return ret
	}
	return *o.Fahrenheit
}

// GetFahrenheitOk returns a tuple with the Fahrenheit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) GetFahrenheitOk() (*float32, bool) {
	if o == nil || isNil(o.Fahrenheit) {
    return nil, false
	}
	return o.Fahrenheit, true
}

// HasFahrenheit returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) HasFahrenheit() bool {
	if o != nil && !isNil(o.Fahrenheit) {
		return true
	}

	return false
}

// SetFahrenheit gets a reference to the given float32 and assigns it to the Fahrenheit field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) SetFahrenheit(v float32) {
	o.Fahrenheit = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) SetQuality(v string) {
	o.Quality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Celsius) {
		toSerialize["celsius"] = o.Celsius
	}
	if !isNil(o.Fahrenheit) {
		toSerialize["fahrenheit"] = o.Fahrenheit
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature(val *NetworksNetworkIdSensorAlertsProfilesThresholdTemperature) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdTemperature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


