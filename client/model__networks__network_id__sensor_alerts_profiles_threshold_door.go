/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSensorAlertsProfilesThresholdDoor Door open threshold. 'open' must be provided and set to true.
type NetworksNetworkIdSensorAlertsProfilesThresholdDoor struct {
	// Alerting threshold for a door open event. Must be set to true.
	Open bool `json:"open"`
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdDoor instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdDoor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSensorAlertsProfilesThresholdDoor(open bool) *NetworksNetworkIdSensorAlertsProfilesThresholdDoor {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdDoor{}
	this.Open = open
	return &this
}

// NewNetworksNetworkIdSensorAlertsProfilesThresholdDoorWithDefaults instantiates a new NetworksNetworkIdSensorAlertsProfilesThresholdDoor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSensorAlertsProfilesThresholdDoorWithDefaults() *NetworksNetworkIdSensorAlertsProfilesThresholdDoor {
	this := NetworksNetworkIdSensorAlertsProfilesThresholdDoor{}
	return &this
}

// GetOpen returns the Open field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdDoor) GetOpen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Open
}

// GetOpenOk returns a tuple with the Open field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdDoor) GetOpenOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Open, true
}

// SetOpen sets field value
func (o *NetworksNetworkIdSensorAlertsProfilesThresholdDoor) SetOpen(v bool) {
	o.Open = v
}

func (o NetworksNetworkIdSensorAlertsProfilesThresholdDoor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["open"] = o.Open
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor struct {
	value *NetworksNetworkIdSensorAlertsProfilesThresholdDoor
	isSet bool
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor) Get() *NetworksNetworkIdSensorAlertsProfilesThresholdDoor {
	return v.value
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor) Set(val *NetworksNetworkIdSensorAlertsProfilesThresholdDoor) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor(val *NetworksNetworkIdSensorAlertsProfilesThresholdDoor) *NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor {
	return &NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSensorAlertsProfilesThresholdDoor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


