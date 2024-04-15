/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceRfProfilesFiveGhzSettings Settings related to 5Ghz band
type NetworksNetworkIdApplianceRfProfilesFiveGhzSettings struct {
	// Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettings instantiates a new NetworksNetworkIdApplianceRfProfilesFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettings() *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings {
	this := NetworksNetworkIdApplianceRfProfilesFiveGhzSettings{}
	return &this
}

// NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettingsWithDefaults instantiates a new NetworksNetworkIdApplianceRfProfilesFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceRfProfilesFiveGhzSettingsWithDefaults() *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings {
	this := NetworksNetworkIdApplianceRfProfilesFiveGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings struct {
	value *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings) Get() *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings) Set(val *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings(val *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) *NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings {
	return &NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceRfProfilesFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


