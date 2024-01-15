/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 Settings for SSID 11
type NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 struct {
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *NetworksNetworkIdWirelessRfProfilesApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings11 instantiates a new NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings11() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 {
	this := NetworksNetworkIdWirelessRfProfilesPerSsidSettings11{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings11WithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings11WithDefaults() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 {
	this := NetworksNetworkIdWirelessRfProfilesPerSsidSettings11{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetMinBitrate() float32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetMinBitrateOk() (*float32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetBands() NetworksNetworkIdWirelessRfProfilesApBandSettingsBands {
	if o == nil || isNil(o.Bands) {
		var ret NetworksNetworkIdWirelessRfProfilesApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetBandsOk() (*NetworksNetworkIdWirelessRfProfilesApBandSettingsBands, bool) {
	if o == nil || isNil(o.Bands) {
    return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) HasBands() bool {
	if o != nil && !isNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given NetworksNetworkIdWirelessRfProfilesApBandSettingsBands and assigns it to the Bands field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) SetBands(v NetworksNetworkIdWirelessRfProfilesApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.Bands) {
		toSerialize["bands"] = o.Bands
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11 struct {
	value *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11) Get() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11 {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11) Set(val *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11(val *NetworksNetworkIdWirelessRfProfilesPerSsidSettings11) *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11 {
	return &NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings11) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


