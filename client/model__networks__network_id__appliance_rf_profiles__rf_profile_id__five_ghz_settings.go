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

// NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings Settings related to 5Ghz band
type NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings struct {
	// Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings instantiates a new NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings() *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings {
	this := NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings{}
	return &this
}

// NewNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettingsWithDefaults instantiates a new NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettingsWithDefaults() *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings {
	this := NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings struct {
	value *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) Get() *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) Set(val *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings(val *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) *NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings {
	return &NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


