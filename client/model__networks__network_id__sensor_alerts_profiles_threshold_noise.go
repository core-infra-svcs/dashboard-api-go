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

// NetworksNetworkIdSensorAlertsProfilesThresholdNoise Noise threshold. 'ambient' must be provided.
type NetworksNetworkIdSensorAlertsProfilesThresholdNoise struct {
	Ambient NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient `json:"ambient"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdNoise instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdNoise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdNoise(ambient NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) *NetworksNetworkIdSensorAlertsProfilesThresholdNoise {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdNoise{}
	this.Ambient = ambient
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdNoiseWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdNoise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdNoiseWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdNoise {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdNoise{}
	return &this
}

// GetAmbient returns the Ambient field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoise) GetAmbient() NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient {
	if o == nil {
		var ret NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient
		return ret
	}

	return o.Ambient
}

// GetAmbientOk returns a tuple with the Ambient field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoise) GetAmbientOk() (*NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ambient, true
}

// SetAmbient sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdNoise) SetAmbient(v NetworksNetworkIdSensorAlertsProfilesThresholdNoiseAmbient) {
	o.Ambient = v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdNoise) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ambient"] = o.Ambient
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdNoise
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdNoise {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdNoise) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise(val *NetworksNetworkIdSensorAlertsProfilesThresholdNoise) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdNoise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


