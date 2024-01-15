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

// InlineResponse200110 struct for InlineResponse200110
type InlineResponse200110 struct {
	// The name of the new profile. Must be unique.
	Id *string `json:"id,omitempty"`
	// The network ID of the RF Profile
	NetworkId *string `json:"networkId,omitempty"`
	// The name of the new profile. Must be unique. This param is required on creation.
	Name *string `json:"name,omitempty"`
	// Steers client to best available access point. Can be either true or false. Defaults to true.
	ClientBalancingEnabled *bool `json:"clientBalancingEnabled,omitempty"`
	// Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	MinBitrateType *string `json:"minBitrateType,omitempty"`
	// Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	BandSelectionType *string `json:"bandSelectionType,omitempty"`
	ApBandSettings *InlineResponse200110ApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *InlineResponse200110TwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *InlineResponse200110FiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	SixGhzSettings *InlineResponse200110SixGhzSettings `json:"sixGhzSettings,omitempty"`
	Transmission *InlineResponse200110Transmission `json:"transmission,omitempty"`
	PerSsidSettings *InlineResponse200110PerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewInlineResponse200110 instantiates a new InlineResponse200110 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200110() *InlineResponse200110 {
	this := InlineResponse200110{}
	return &this
}

// NewInlineResponse200110WithDefaults instantiates a new InlineResponse200110 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200110WithDefaults() *InlineResponse200110 {
	this := InlineResponse200110{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200110) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200110) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200110) SetId(v string) {
	o.Id = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *InlineResponse200110) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *InlineResponse200110) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *InlineResponse200110) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200110) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200110) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200110) SetName(v string) {
	o.Name = &v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *InlineResponse200110) GetClientBalancingEnabled() bool {
	if o == nil || isNil(o.ClientBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ClientBalancingEnabled) {
    return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *InlineResponse200110) HasClientBalancingEnabled() bool {
	if o != nil && !isNil(o.ClientBalancingEnabled) {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *InlineResponse200110) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *InlineResponse200110) GetMinBitrateType() string {
	if o == nil || isNil(o.MinBitrateType) {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || isNil(o.MinBitrateType) {
    return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *InlineResponse200110) HasMinBitrateType() bool {
	if o != nil && !isNil(o.MinBitrateType) {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *InlineResponse200110) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value if set, zero value otherwise.
func (o *InlineResponse200110) GetBandSelectionType() string {
	if o == nil || isNil(o.BandSelectionType) {
		var ret string
		return ret
	}
	return *o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.BandSelectionType) {
    return nil, false
	}
	return o.BandSelectionType, true
}

// HasBandSelectionType returns a boolean if a field has been set.
func (o *InlineResponse200110) HasBandSelectionType() bool {
	if o != nil && !isNil(o.BandSelectionType) {
		return true
	}

	return false
}

// SetBandSelectionType gets a reference to the given string and assigns it to the BandSelectionType field.
func (o *InlineResponse200110) SetBandSelectionType(v string) {
	o.BandSelectionType = &v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *InlineResponse200110) GetApBandSettings() InlineResponse200110ApBandSettings {
	if o == nil || isNil(o.ApBandSettings) {
		var ret InlineResponse200110ApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetApBandSettingsOk() (*InlineResponse200110ApBandSettings, bool) {
	if o == nil || isNil(o.ApBandSettings) {
    return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *InlineResponse200110) HasApBandSettings() bool {
	if o != nil && !isNil(o.ApBandSettings) {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given InlineResponse200110ApBandSettings and assigns it to the ApBandSettings field.
func (o *InlineResponse200110) SetApBandSettings(v InlineResponse200110ApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse200110) GetTwoFourGhzSettings() InlineResponse200110TwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret InlineResponse200110TwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetTwoFourGhzSettingsOk() (*InlineResponse200110TwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse200110) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given InlineResponse200110TwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineResponse200110) SetTwoFourGhzSettings(v InlineResponse200110TwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse200110) GetFiveGhzSettings() InlineResponse200110FiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret InlineResponse200110FiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetFiveGhzSettingsOk() (*InlineResponse200110FiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse200110) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given InlineResponse200110FiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineResponse200110) SetFiveGhzSettings(v InlineResponse200110FiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetSixGhzSettings returns the SixGhzSettings field value if set, zero value otherwise.
func (o *InlineResponse200110) GetSixGhzSettings() InlineResponse200110SixGhzSettings {
	if o == nil || isNil(o.SixGhzSettings) {
		var ret InlineResponse200110SixGhzSettings
		return ret
	}
	return *o.SixGhzSettings
}

// GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetSixGhzSettingsOk() (*InlineResponse200110SixGhzSettings, bool) {
	if o == nil || isNil(o.SixGhzSettings) {
    return nil, false
	}
	return o.SixGhzSettings, true
}

// HasSixGhzSettings returns a boolean if a field has been set.
func (o *InlineResponse200110) HasSixGhzSettings() bool {
	if o != nil && !isNil(o.SixGhzSettings) {
		return true
	}

	return false
}

// SetSixGhzSettings gets a reference to the given InlineResponse200110SixGhzSettings and assigns it to the SixGhzSettings field.
func (o *InlineResponse200110) SetSixGhzSettings(v InlineResponse200110SixGhzSettings) {
	o.SixGhzSettings = &v
}

// GetTransmission returns the Transmission field value if set, zero value otherwise.
func (o *InlineResponse200110) GetTransmission() InlineResponse200110Transmission {
	if o == nil || isNil(o.Transmission) {
		var ret InlineResponse200110Transmission
		return ret
	}
	return *o.Transmission
}

// GetTransmissionOk returns a tuple with the Transmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetTransmissionOk() (*InlineResponse200110Transmission, bool) {
	if o == nil || isNil(o.Transmission) {
    return nil, false
	}
	return o.Transmission, true
}

// HasTransmission returns a boolean if a field has been set.
func (o *InlineResponse200110) HasTransmission() bool {
	if o != nil && !isNil(o.Transmission) {
		return true
	}

	return false
}

// SetTransmission gets a reference to the given InlineResponse200110Transmission and assigns it to the Transmission field.
func (o *InlineResponse200110) SetTransmission(v InlineResponse200110Transmission) {
	o.Transmission = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineResponse200110) GetPerSsidSettings() InlineResponse200110PerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret InlineResponse200110PerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200110) GetPerSsidSettingsOk() (*InlineResponse200110PerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineResponse200110) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given InlineResponse200110PerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineResponse200110) SetPerSsidSettings(v InlineResponse200110PerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o InlineResponse200110) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
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
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200110 struct {
	value *InlineResponse200110
	isSet bool
}

func (v NullableInlineResponse200110) Get() *InlineResponse200110 {
	return v.value
}

func (v *NullableInlineResponse200110) Set(val *InlineResponse200110) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200110) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200110) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200110(val *InlineResponse200110) *NullableInlineResponse200110 {
	return &NullableInlineResponse200110{value: val, isSet: true}
}

func (v NullableInlineResponse200110) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200110) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


