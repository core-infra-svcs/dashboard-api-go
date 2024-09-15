/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject182 struct for InlineObject182
type InlineObject182 struct {
	// The name of the new profile. Must be unique.
	Name *string `json:"name,omitempty"`
	// Set this profile as the default indoor rf profile. If the profile ID is one of 'indoor' or 'outdoor',   then a new profile will be created from the respective ID and set as the default
	IsIndoorDefault *bool `json:"isIndoorDefault,omitempty"`
	// Set this profile as the default outdoor rf profile. If the profile ID is one of 'indoor' or 'outdoor',   then a new profile will be created from the respective ID and set as the default
	IsOutdoorDefault *bool `json:"isOutdoorDefault,omitempty"`
	// Steers client to best available access point. Can be either true or false.
	ClientBalancingEnabled *bool `json:"clientBalancingEnabled,omitempty"`
	// Minimum bitrate can be set to either 'band' or 'ssid'.
	MinBitrateType *string `json:"minBitrateType,omitempty"`
	// Band selection can be set to either 'ssid' or 'ap'.
	BandSelectionType *string `json:"bandSelectionType,omitempty"`
	ApBandSettings *NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	SixGhzSettings *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings `json:"sixGhzSettings,omitempty"`
	Transmission *InlineResponse200184Transmission `json:"transmission,omitempty"`
	PerSsidSettings *NetworksNetworkIdWirelessRfProfilesPerSsidSettings `json:"perSsidSettings,omitempty"`
	FlexRadios *NetworksNetworkIdWirelessRfProfilesFlexRadios `json:"flexRadios,omitempty"`
}

// NewInlineObject182 instantiates a new InlineObject182 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject182() *InlineObject182 {
	this := InlineObject182{}
	return &this
}

// NewInlineObject182WithDefaults instantiates a new InlineObject182 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject182WithDefaults() *InlineObject182 {
	this := InlineObject182{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject182) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject182) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject182) SetName(v string) {
	o.Name = &v
}

// GetIsIndoorDefault returns the IsIndoorDefault field value if set, zero value otherwise.
func (o *InlineObject182) GetIsIndoorDefault() bool {
	if o == nil || isNil(o.IsIndoorDefault) {
		var ret bool
		return ret
	}
	return *o.IsIndoorDefault
}

// GetIsIndoorDefaultOk returns a tuple with the IsIndoorDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetIsIndoorDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsIndoorDefault) {
    return nil, false
	}
	return o.IsIndoorDefault, true
}

// HasIsIndoorDefault returns a boolean if a field has been set.
func (o *InlineObject182) HasIsIndoorDefault() bool {
	if o != nil && !isNil(o.IsIndoorDefault) {
		return true
	}

	return false
}

// SetIsIndoorDefault gets a reference to the given bool and assigns it to the IsIndoorDefault field.
func (o *InlineObject182) SetIsIndoorDefault(v bool) {
	o.IsIndoorDefault = &v
}

// GetIsOutdoorDefault returns the IsOutdoorDefault field value if set, zero value otherwise.
func (o *InlineObject182) GetIsOutdoorDefault() bool {
	if o == nil || isNil(o.IsOutdoorDefault) {
		var ret bool
		return ret
	}
	return *o.IsOutdoorDefault
}

// GetIsOutdoorDefaultOk returns a tuple with the IsOutdoorDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetIsOutdoorDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsOutdoorDefault) {
    return nil, false
	}
	return o.IsOutdoorDefault, true
}

// HasIsOutdoorDefault returns a boolean if a field has been set.
func (o *InlineObject182) HasIsOutdoorDefault() bool {
	if o != nil && !isNil(o.IsOutdoorDefault) {
		return true
	}

	return false
}

// SetIsOutdoorDefault gets a reference to the given bool and assigns it to the IsOutdoorDefault field.
func (o *InlineObject182) SetIsOutdoorDefault(v bool) {
	o.IsOutdoorDefault = &v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *InlineObject182) GetClientBalancingEnabled() bool {
	if o == nil || isNil(o.ClientBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ClientBalancingEnabled) {
    return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *InlineObject182) HasClientBalancingEnabled() bool {
	if o != nil && !isNil(o.ClientBalancingEnabled) {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *InlineObject182) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *InlineObject182) GetMinBitrateType() string {
	if o == nil || isNil(o.MinBitrateType) {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || isNil(o.MinBitrateType) {
    return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *InlineObject182) HasMinBitrateType() bool {
	if o != nil && !isNil(o.MinBitrateType) {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *InlineObject182) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value if set, zero value otherwise.
func (o *InlineObject182) GetBandSelectionType() string {
	if o == nil || isNil(o.BandSelectionType) {
		var ret string
		return ret
	}
	return *o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.BandSelectionType) {
    return nil, false
	}
	return o.BandSelectionType, true
}

// HasBandSelectionType returns a boolean if a field has been set.
func (o *InlineObject182) HasBandSelectionType() bool {
	if o != nil && !isNil(o.BandSelectionType) {
		return true
	}

	return false
}

// SetBandSelectionType gets a reference to the given string and assigns it to the BandSelectionType field.
func (o *InlineObject182) SetBandSelectionType(v string) {
	o.BandSelectionType = &v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *InlineObject182) GetApBandSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings {
	if o == nil || isNil(o.ApBandSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetApBandSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings, bool) {
	if o == nil || isNil(o.ApBandSettings) {
    return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *InlineObject182) HasApBandSettings() bool {
	if o != nil && !isNil(o.ApBandSettings) {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings and assigns it to the ApBandSettings field.
func (o *InlineObject182) SetApBandSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject182) GetTwoFourGhzSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetTwoFourGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject182) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject182) SetTwoFourGhzSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject182) GetFiveGhzSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetFiveGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject182) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject182) SetFiveGhzSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetSixGhzSettings returns the SixGhzSettings field value if set, zero value otherwise.
func (o *InlineObject182) GetSixGhzSettings() NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings {
	if o == nil || isNil(o.SixGhzSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings
		return ret
	}
	return *o.SixGhzSettings
}

// GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetSixGhzSettingsOk() (*NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings, bool) {
	if o == nil || isNil(o.SixGhzSettings) {
    return nil, false
	}
	return o.SixGhzSettings, true
}

// HasSixGhzSettings returns a boolean if a field has been set.
func (o *InlineObject182) HasSixGhzSettings() bool {
	if o != nil && !isNil(o.SixGhzSettings) {
		return true
	}

	return false
}

// SetSixGhzSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings and assigns it to the SixGhzSettings field.
func (o *InlineObject182) SetSixGhzSettings(v NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) {
	o.SixGhzSettings = &v
}

// GetTransmission returns the Transmission field value if set, zero value otherwise.
func (o *InlineObject182) GetTransmission() InlineResponse200184Transmission {
	if o == nil || isNil(o.Transmission) {
		var ret InlineResponse200184Transmission
		return ret
	}
	return *o.Transmission
}

// GetTransmissionOk returns a tuple with the Transmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetTransmissionOk() (*InlineResponse200184Transmission, bool) {
	if o == nil || isNil(o.Transmission) {
    return nil, false
	}
	return o.Transmission, true
}

// HasTransmission returns a boolean if a field has been set.
func (o *InlineObject182) HasTransmission() bool {
	if o != nil && !isNil(o.Transmission) {
		return true
	}

	return false
}

// SetTransmission gets a reference to the given InlineResponse200184Transmission and assigns it to the Transmission field.
func (o *InlineObject182) SetTransmission(v InlineResponse200184Transmission) {
	o.Transmission = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineObject182) GetPerSsidSettings() NetworksNetworkIdWirelessRfProfilesPerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret NetworksNetworkIdWirelessRfProfilesPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetPerSsidSettingsOk() (*NetworksNetworkIdWirelessRfProfilesPerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineObject182) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given NetworksNetworkIdWirelessRfProfilesPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineObject182) SetPerSsidSettings(v NetworksNetworkIdWirelessRfProfilesPerSsidSettings) {
	o.PerSsidSettings = &v
}

// GetFlexRadios returns the FlexRadios field value if set, zero value otherwise.
func (o *InlineObject182) GetFlexRadios() NetworksNetworkIdWirelessRfProfilesFlexRadios {
	if o == nil || isNil(o.FlexRadios) {
		var ret NetworksNetworkIdWirelessRfProfilesFlexRadios
		return ret
	}
	return *o.FlexRadios
}

// GetFlexRadiosOk returns a tuple with the FlexRadios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject182) GetFlexRadiosOk() (*NetworksNetworkIdWirelessRfProfilesFlexRadios, bool) {
	if o == nil || isNil(o.FlexRadios) {
    return nil, false
	}
	return o.FlexRadios, true
}

// HasFlexRadios returns a boolean if a field has been set.
func (o *InlineObject182) HasFlexRadios() bool {
	if o != nil && !isNil(o.FlexRadios) {
		return true
	}

	return false
}

// SetFlexRadios gets a reference to the given NetworksNetworkIdWirelessRfProfilesFlexRadios and assigns it to the FlexRadios field.
func (o *InlineObject182) SetFlexRadios(v NetworksNetworkIdWirelessRfProfilesFlexRadios) {
	o.FlexRadios = &v
}

func (o InlineObject182) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsIndoorDefault) {
		toSerialize["isIndoorDefault"] = o.IsIndoorDefault
	}
	if !isNil(o.IsOutdoorDefault) {
		toSerialize["isOutdoorDefault"] = o.IsOutdoorDefault
	}
	if !isNil(o.ClientBalancingEnabled) {
		toSerialize["clientBalancingEnabled"] = o.ClientBalancingEnabled
	}
	if !isNil(o.MinBitrateType) {
		toSerialize["minBitrateType"] = o.MinBitrateType
	}
	if !isNil(o.BandSelectionType) {
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

type NullableInlineObject182 struct {
	value *InlineObject182
	isSet bool
}

func (v NullableInlineObject182) Get() *InlineObject182 {
	return v.value
}

func (v *NullableInlineObject182) Set(val *InlineObject182) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject182) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject182) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject182(val *InlineObject182) *NullableInlineObject182 {
	return &NullableInlineObject182{value: val, isSet: true}
}

func (v NullableInlineObject182) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject182) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


