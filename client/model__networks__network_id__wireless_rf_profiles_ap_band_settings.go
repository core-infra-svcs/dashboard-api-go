/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 December, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.28.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesApBandSettings Settings that will be enabled if selectionType is set to 'ap'.
type NetworksNetworkIdWirelessRfProfilesApBandSettings struct {
	// Choice between 'dual', '2.4ghz' or '5ghz'. Defaults to dual.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band. Can be either true or false. Defaults to true.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesApBandSettings instantiates a new NetworksNetworkIdWirelessRfProfilesApBandSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesApBandSettings() *NetworksNetworkIdWirelessRfProfilesApBandSettings {
	this := NetworksNetworkIdWirelessRfProfilesApBandSettings{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesApBandSettingsWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesApBandSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesApBandSettingsWithDefaults() *NetworksNetworkIdWirelessRfProfilesApBandSettings {
	this := NetworksNetworkIdWirelessRfProfilesApBandSettings{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *NetworksNetworkIdWirelessRfProfilesApBandSettings) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o NetworksNetworkIdWirelessRfProfilesApBandSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesApBandSettings struct {
	value *NetworksNetworkIdWirelessRfProfilesApBandSettings
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesApBandSettings) Get() *NetworksNetworkIdWirelessRfProfilesApBandSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesApBandSettings) Set(val *NetworksNetworkIdWirelessRfProfilesApBandSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesApBandSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesApBandSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesApBandSettings(val *NetworksNetworkIdWirelessRfProfilesApBandSettings) *NullableNetworksNetworkIdWirelessRfProfilesApBandSettings {
	return &NullableNetworksNetworkIdWirelessRfProfilesApBandSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesApBandSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesApBandSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


