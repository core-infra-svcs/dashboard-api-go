/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 Settings for SSID 6
type NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 struct {
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *NetworksNetworkIdWirelessRfProfilesApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings6 instantiates a new NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings6() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 {
	this := NetworksNetworkIdWirelessRfProfilesPerSsidSettings6{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings6WithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesPerSsidSettings6WithDefaults() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 {
	this := NetworksNetworkIdWirelessRfProfilesPerSsidSettings6{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetMinBitrate() float32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetMinBitrateOk() (*float32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetBands() NetworksNetworkIdWirelessRfProfilesApBandSettingsBands {
	if o == nil || isNil(o.Bands) {
		var ret NetworksNetworkIdWirelessRfProfilesApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetBandsOk() (*NetworksNetworkIdWirelessRfProfilesApBandSettingsBands, bool) {
	if o == nil || isNil(o.Bands) {
    return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) HasBands() bool {
	if o != nil && !isNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given NetworksNetworkIdWirelessRfProfilesApBandSettingsBands and assigns it to the Bands field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) SetBands(v NetworksNetworkIdWirelessRfProfilesApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) MarshalJSON() ([]byte, error) {
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

type NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6 struct {
	value *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6) Get() *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6 {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6) Set(val *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6(val *NetworksNetworkIdWirelessRfProfilesPerSsidSettings6) *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6 {
	return &NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesPerSsidSettings6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


