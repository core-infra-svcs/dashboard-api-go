/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings Settings related to 2.4Ghz band
type NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings struct {
	// Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30.
	MaxPower *int32 `json:"maxPower,omitempty"`
	// Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30.
	MinPower *int32 `json:"minPower,omitempty"`
	// Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'.
	ValidAutoChannels []int32 `json:"validAutoChannels,omitempty"`
	// Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering.
	AxEnabled *bool `json:"axEnabled,omitempty"`
	// The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	Rxsop *int32 `json:"rxsop,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings instantiates a new NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings() *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings {
	this := NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettingsWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettingsWithDefaults() *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings {
	this := NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings{}
	return &this
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetMaxPower() int32 {
	if o == nil || isNil(o.MaxPower) {
		var ret int32
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetMaxPowerOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPower) {
    return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) HasMaxPower() bool {
	if o != nil && !isNil(o.MaxPower) {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given int32 and assigns it to the MaxPower field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) SetMaxPower(v int32) {
	o.MaxPower = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetMinPower() int32 {
	if o == nil || isNil(o.MinPower) {
		var ret int32
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetMinPowerOk() (*int32, bool) {
	if o == nil || isNil(o.MinPower) {
    return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) HasMinPower() bool {
	if o != nil && !isNil(o.MinPower) {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given int32 and assigns it to the MinPower field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) SetMinPower(v int32) {
	o.MinPower = &v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetMinBitrate() float32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetMinBitrateOk() (*float32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetValidAutoChannels returns the ValidAutoChannels field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetValidAutoChannels() []int32 {
	if o == nil || isNil(o.ValidAutoChannels) {
		var ret []int32
		return ret
	}
	return o.ValidAutoChannels
}

// GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetValidAutoChannelsOk() ([]int32, bool) {
	if o == nil || isNil(o.ValidAutoChannels) {
    return nil, false
	}
	return o.ValidAutoChannels, true
}

// HasValidAutoChannels returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) HasValidAutoChannels() bool {
	if o != nil && !isNil(o.ValidAutoChannels) {
		return true
	}

	return false
}

// SetValidAutoChannels gets a reference to the given []int32 and assigns it to the ValidAutoChannels field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) SetValidAutoChannels(v []int32) {
	o.ValidAutoChannels = v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetAxEnabled() bool {
	if o == nil || isNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AxEnabled) {
    return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) HasAxEnabled() bool {
	if o != nil && !isNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

// GetRxsop returns the Rxsop field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetRxsop() int32 {
	if o == nil || isNil(o.Rxsop) {
		var ret int32
		return ret
	}
	return *o.Rxsop
}

// GetRxsopOk returns a tuple with the Rxsop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) GetRxsopOk() (*int32, bool) {
	if o == nil || isNil(o.Rxsop) {
    return nil, false
	}
	return o.Rxsop, true
}

// HasRxsop returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) HasRxsop() bool {
	if o != nil && !isNil(o.Rxsop) {
		return true
	}

	return false
}

// SetRxsop gets a reference to the given int32 and assigns it to the Rxsop field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) SetRxsop(v int32) {
	o.Rxsop = &v
}

func (o NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxPower) {
		toSerialize["maxPower"] = o.MaxPower
	}
	if !isNil(o.MinPower) {
		toSerialize["minPower"] = o.MinPower
	}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.ValidAutoChannels) {
		toSerialize["validAutoChannels"] = o.ValidAutoChannels
	}
	if !isNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	if !isNil(o.Rxsop) {
		toSerialize["rxsop"] = o.Rxsop
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings struct {
	value *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) Get() *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) Set(val *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings(val *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings {
	return &NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


