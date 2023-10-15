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

// NetworksNetworkIdSensorAlertsProfilesThresholdHumidity Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdHumidity struct {
	// Alerting threshold in %RH.
	RelativePercentage *int32 `json:"relativePercentage,omitempty"`
	// Alerting threshold as a qualitative humidity level.
	Quality *string `json:"quality,omitempty"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdHumidity instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdHumidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdHumidity() *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdHumidity{}
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdHumidityWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdHumidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdHumidityWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdHumidity{}
	return &this
}

// GetRelativePercentage returns the RelativePercentage field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) GetRelativePercentage() int32 {
	if o == nil || isNil(o.RelativePercentage) {
		var ret int32
		return ret
	}
	return *o.RelativePercentage
}

// GetRelativePercentageOk returns a tuple with the RelativePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) GetRelativePercentageOk() (*int32, bool) {
	if o == nil || isNil(o.RelativePercentage) {
    return nil, false
	}
	return o.RelativePercentage, true
}

// HasRelativePercentage returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) HasRelativePercentage() bool {
	if o != nil && !isNil(o.RelativePercentage) {
		return true
	}

	return false
}

// SetRelativePercentage gets a reference to the given int32 and assigns it to the RelativePercentage field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) SetRelativePercentage(v int32) {
	o.RelativePercentage = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) SetQuality(v string) {
	o.Quality = &v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelativePercentage) {
		toSerialize["relativePercentage"] = o.RelativePercentage
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity(val *NetworksNetworkIdSensorAlertsProfilesThresholdHumidity) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdHumidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


