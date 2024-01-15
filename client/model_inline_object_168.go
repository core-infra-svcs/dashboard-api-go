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

// InlineObject168 struct for InlineObject168
type InlineObject168 struct {
	// The name of the new profile. Must be unique. This param is required on creation.
	Name string `json:"name"`
	// Steers client to best available access point. Can be either true or false. Defaults to true.
	ClientBalancingEnabled *bool `json:"clientBalancingEnabled,omitempty"`
	// Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	MinBitrateType *string `json:"minBitrateType,omitempty"`
	// Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	BandSelectionType string `json:"bandSelectionType"`
	ApBandSettings *NetworksNetworkIdWirelessRfProfilesApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *InlineResponse200110TwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *InlineResponse200110FiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	SixGhzSettings *NetworksNetworkIdWirelessRfProfilesSixGhzSettings `json:"sixGhzSettings,omitempty"`
	Transmission *InlineResponse200110Transmission `json:"transmission,omitempty"`
	PerSsidSettings *NetworksNetworkIdWirelessRfProfilesPerSsidSettings `json:"perSsidSettings,omitempty"`
	FlexRadios *NetworksNetworkIdWirelessRfProfilesFlexRadios `json:"flexRadios,omitempty"`
}

// NewInlineObject168 instantiates a new InlineObject168 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject168(name string, bandSelectionType string) *InlineObject168 {
	this := InlineObject168{}
	this.Name = name
	this.BandSelectionType = bandSelectionType
	return &this
}

// NewInlineObject168WithDefaults instantiates a new InlineObject168 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject168WithDefaults() *InlineObject168 {
	this := InlineObject168{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject168) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject168) SetName(v string) {
	o.Name = v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *InlineObject168) GetClientBalancingEnabled() bool {
	if o == nil || isNil(o.ClientBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ClientBalancingEnabled) {
    return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *InlineObject168) HasClientBalancingEnabled() bool {
	if o != nil && !isNil(o.ClientBalancingEnabled) {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *InlineObject168) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *InlineObject168) GetMinBitrateType() string {
	if o == nil || isNil(o.MinBitrateType) {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || isNil(o.MinBitrateType) {
    return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *InlineObject168) HasMinBitrateType() bool {
	if o != nil && !isNil(o.MinBitrateType) {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *InlineObject168) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value
func (o *InlineObject168) GetBandSelectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BandSelectionType, true
}

// SetBandSelectionType sets field value
func (o *InlineObject168) SetBandSelectionType(v string) {
	o.BandSelectionType = v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *InlineObject168) GetApBandSettings() NetworksNetworkIdWirelessRfProfilesApBandSettings {
	if o == nil || isNil(o.ApBandSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetApBandSettingsOk() (*NetworksNetworkIdWirelessRfProfilesApBandSettings, bool) {
	if o == nil || isNil(o.ApBandSettings) {
    return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *InlineObject168) HasApBandSettings() bool {
	if o != nil && !isNil(o.ApBandSettings) {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesApBandSettings and assigns it to the ApBandSettings field.
func (o *InlineObject168) SetApBandSettings(v NetworksNetworkIdWirelessRfProfilesApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject168) GetTwoFourGhzSettings() InlineResponse200110TwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret InlineResponse200110TwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetTwoFourGhzSettingsOk() (*InlineResponse200110TwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject168) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given InlineResponse200110TwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject168) SetTwoFourGhzSettings(v InlineResponse200110TwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject168) GetFiveGhzSettings() InlineResponse200110FiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret InlineResponse200110FiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetFiveGhzSettingsOk() (*InlineResponse200110FiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject168) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given InlineResponse200110FiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject168) SetFiveGhzSettings(v InlineResponse200110FiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetSixGhzSettings returns the SixGhzSettings field value if set, zero value otherwise.
func (o *InlineObject168) GetSixGhzSettings() NetworksNetworkIdWirelessRfProfilesSixGhzSettings {
	if o == nil || isNil(o.SixGhzSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesSixGhzSettings
		return ret
	}
	return *o.SixGhzSettings
}

// GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetSixGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesSixGhzSettings, bool) {
	if o == nil || isNil(o.SixGhzSettings) {
    return nil, false
	}
	return o.SixGhzSettings, true
}

// HasSixGhzSettings returns a boolean if a field has been set.
func (o *InlineObject168) HasSixGhzSettings() bool {
	if o != nil && !isNil(o.SixGhzSettings) {
		return true
	}

	return false
}

// SetSixGhzSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesSixGhzSettings and assigns it to the SixGhzSettings field.
func (o *InlineObject168) SetSixGhzSettings(v NetworksNetworkIdWirelessRfProfilesSixGhzSettings) {
	o.SixGhzSettings = &v
}

// GetTransmission returns the Transmission field value if set, zero value otherwise.
func (o *InlineObject168) GetTransmission() InlineResponse200110Transmission {
	if o == nil || isNil(o.Transmission) {
		var ret InlineResponse200110Transmission
		return ret
	}
	return *o.Transmission
}

// GetTransmissionOk returns a tuple with the Transmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetTransmissionOk() (*InlineResponse200110Transmission, bool) {
	if o == nil || isNil(o.Transmission) {
    return nil, false
	}
	return o.Transmission, true
}

// HasTransmission returns a boolean if a field has been set.
func (o *InlineObject168) HasTransmission() bool {
	if o != nil && !isNil(o.Transmission) {
		return true
	}

	return false
}

// SetTransmission gets a reference to the given InlineResponse200110Transmission and assigns it to the Transmission field.
func (o *InlineObject168) SetTransmission(v InlineResponse200110Transmission) {
	o.Transmission = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineObject168) GetPerSsidSettings() NetworksNetworkIdWirelessRfProfilesPerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetPerSsidSettingsOk() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineObject168) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineObject168) SetPerSsidSettings(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings) {
	o.PerSsidSettings = &v
}

// GetFlexRadios returns the FlexRadios field value if set, zero value otherwise.
func (o *InlineObject168) GetFlexRadios() NetworksNetworkIdWirelessRfProfilesFlexRadios {
	if o == nil || isNil(o.FlexRadios) {
		var ret NetworksNetworkIdWirelessRfProfilesFlexRadios
		return ret
	}
	return *o.FlexRadios
}

// GetFlexRadiosOk returns a tuple with the FlexRadios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject168) GetFlexRadiosOk() (*NetworksNetworkIdWirelessRfProfilesFlexRadios, bool) {
	if o == nil || isNil(o.FlexRadios) {
    return nil, false
	}
	return o.FlexRadios, true
}

// HasFlexRadios returns a boolean if a field has been set.
func (o *InlineObject168) HasFlexRadios() bool {
	if o != nil && !isNil(o.FlexRadios) {
		return true
	}

	return false
}

// SetFlexRadios gets a reference to the given NetworksNetworkIdWirelessRfProfilesFlexRadios and assigns it to the FlexRadios field.
func (o *InlineObject168) SetFlexRadios(v NetworksNetworkIdWirelessRfProfilesFlexRadios) {
	o.FlexRadios = &v
}

func (o InlineObject168) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ClientBalancingEnabled) {
		toSerialize["clientBalancingEnabled"] = o.ClientBalancingEnabled
	}
	if !isNil(o.MinBitrateType) {
		toSerialize["minBitrateType"] = o.MinBitrateType
	}
	if true {
		toSerialize["bandSelectionType"] = o.BandSelectionType
	}
	if !isNil(o.ApBandSettings) {
		toSerialize["apBandSettings"] = o.ApBandSettings
	}
	if !isNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !isNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if !isNil(o.SixGhzSettings) {
		toSerialize["sixGhzSettings"] = o.SixGhzSettings
	}
	if !isNil(o.Transmission) {
		toSerialize["transmission"] = o.Transmission
	}
	if !isNil(o.PerSsidSettings) {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	if !isNil(o.FlexRadios) {
		toSerialize["flexRadios"] = o.FlexRadios
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject168 struct {
	value *InlineObject168
	isSet bool
}

func (v NullableInlineObject168) Get() *InlineObject168 {
	return v.value
}

func (v *NullableInlineObject168) Set(val *InlineObject168) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject168) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject168) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject168(val *InlineObject168) *NullableInlineObject168 {
	return &NullableInlineObject168{value: val, isSet: true}
}

func (v NullableInlineObject168) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject168) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


