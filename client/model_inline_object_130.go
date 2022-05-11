/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject130 struct for InlineObject130
type InlineObject130 struct {
	// The name of the new profile. Must be unique. This param is required on creation.
	Name string `json:"name"`
	// Steers client to best available access point. Can be either true or false. Defaults to true.
	ClientBalancingEnabled *bool `json:"clientBalancingEnabled,omitempty"`
	// Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	MinBitrateType *string `json:"minBitrateType,omitempty"`
	// Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	BandSelectionType string `json:"bandSelectionType"`
	ApBandSettings *NetworksNetworkIdWirelessRfProfilesApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *NetworksNetworkIdWirelessRfProfilesFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	PerSsidSettings *NetworksNetworkIdWirelessRfProfilesPerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewInlineObject130 instantiates a new InlineObject130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject130(name string, bandSelectionType string) *InlineObject130 {
	this := InlineObject130{}
	this.Name = name
	this.BandSelectionType = bandSelectionType
	return &this
}

// NewInlineObject130WithDefaults instantiates a new InlineObject130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject130WithDefaults() *InlineObject130 {
	this := InlineObject130{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject130) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject130) SetName(v string) {
	o.Name = v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *InlineObject130) GetClientBalancingEnabled() bool {
	if o == nil || o.ClientBalancingEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || o.ClientBalancingEnabled == nil {
		return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *InlineObject130) HasClientBalancingEnabled() bool {
	if o != nil && o.ClientBalancingEnabled != nil {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *InlineObject130) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *InlineObject130) GetMinBitrateType() string {
	if o == nil || o.MinBitrateType == nil {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || o.MinBitrateType == nil {
		return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *InlineObject130) HasMinBitrateType() bool {
	if o != nil && o.MinBitrateType != nil {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *InlineObject130) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value
func (o *InlineObject130) GetBandSelectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BandSelectionType, true
}

// SetBandSelectionType sets field value
func (o *InlineObject130) SetBandSelectionType(v string) {
	o.BandSelectionType = v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *InlineObject130) GetApBandSettings() NetworksNetworkIdWirelessRfProfilesApBandSettings {
	if o == nil || o.ApBandSettings == nil {
		var ret NetworksNetworkIdWirelessRfProfilesApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetApBandSettingsOk() (*NetworksNetworkIdWirelessRfProfilesApBandSettings, bool) {
	if o == nil || o.ApBandSettings == nil {
		return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *InlineObject130) HasApBandSettings() bool {
	if o != nil && o.ApBandSettings != nil {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesApBandSettings and assigns it to the ApBandSettings field.
func (o *InlineObject130) SetApBandSettings(v NetworksNetworkIdWirelessRfProfilesApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject130) GetTwoFourGhzSettings() NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings {
	if o == nil || o.TwoFourGhzSettings == nil {
		var ret NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetTwoFourGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings, bool) {
	if o == nil || o.TwoFourGhzSettings == nil {
		return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject130) HasTwoFourGhzSettings() bool {
	if o != nil && o.TwoFourGhzSettings != nil {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject130) SetTwoFourGhzSettings(v NetworksNetworkIdWirelessRfProfilesTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject130) GetFiveGhzSettings() NetworksNetworkIdWirelessRfProfilesFiveGhzSettings {
	if o == nil || o.FiveGhzSettings == nil {
		var ret NetworksNetworkIdWirelessRfProfilesFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetFiveGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesFiveGhzSettings, bool) {
	if o == nil || o.FiveGhzSettings == nil {
		return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject130) HasFiveGhzSettings() bool {
	if o != nil && o.FiveGhzSettings != nil {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject130) SetFiveGhzSettings(v NetworksNetworkIdWirelessRfProfilesFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineObject130) GetPerSsidSettings() NetworksNetworkIdWirelessRfProfilesPerSsidSettings {
	if o == nil || o.PerSsidSettings == nil {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject130) GetPerSsidSettingsOk() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings, bool) {
	if o == nil || o.PerSsidSettings == nil {
		return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineObject130) HasPerSsidSettings() bool {
	if o != nil && o.PerSsidSettings != nil {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineObject130) SetPerSsidSettings(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o InlineObject130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ClientBalancingEnabled != nil {
		toSerialize["clientBalancingEnabled"] = o.ClientBalancingEnabled
	}
	if o.MinBitrateType != nil {
		toSerialize["minBitrateType"] = o.MinBitrateType
	}
	if true {
		toSerialize["bandSelectionType"] = o.BandSelectionType
	}
	if o.ApBandSettings != nil {
		toSerialize["apBandSettings"] = o.ApBandSettings
	}
	if o.TwoFourGhzSettings != nil {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if o.FiveGhzSettings != nil {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if o.PerSsidSettings != nil {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject130 struct {
	value *InlineObject130
	isSet bool
}

func (v NullableInlineObject130) Get() *InlineObject130 {
	return v.value
}

func (v *NullableInlineObject130) Set(val *InlineObject130) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject130) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject130(val *InlineObject130) *NullableInlineObject130 {
	return &NullableInlineObject130{value: val, isSet: true}
}

func (v NullableInlineObject130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


