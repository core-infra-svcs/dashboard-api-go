/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings Settings related to 2.4Ghz band
type NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings struct {
	// Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings instantiates a new NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings() *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings {
	this := NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings{}
	return &this
}

// NewNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettingsWithDefaults instantiates a new NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettingsWithDefaults() *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings {
	this := NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) GetMinBitrate() float32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) GetMinBitrateOk() (*float32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings struct {
	value *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) Get() *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) Set(val *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings(val *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) *NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings {
	return &NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


