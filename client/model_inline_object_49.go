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

// InlineObject49 struct for InlineObject49
type InlineObject49 struct {
	// The name of the new profile. Must be unique.
	Name *string `json:"name,omitempty"`
	TwoFourGhzSettings *NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	PerSsidSettings *NetworksNetworkIdApplianceRfProfilesPerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewInlineObject49 instantiates a new InlineObject49 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject49() *InlineObject49 {
	this := InlineObject49{}
	return &this
}

// NewInlineObject49WithDefaults instantiates a new InlineObject49 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject49WithDefaults() *InlineObject49 {
	this := InlineObject49{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject49) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject49) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject49) SetName(v string) {
	o.Name = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject49) GetTwoFourGhzSettings() NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetTwoFourGhzSettingsOk() (*NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject49) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject49) SetTwoFourGhzSettings(v NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject49) GetFiveGhzSettings() NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetFiveGhzSettingsOk() (*NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject49) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject49) SetFiveGhzSettings(v NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineObject49) GetPerSsidSettings() NetworksNetworkIdApplianceRfProfilesPerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject49) GetPerSsidSettingsOk() (*NetworksNetworkIdApplianceRfProfilesPerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineObject49) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineObject49) SetPerSsidSettings(v NetworksNetworkIdApplianceRfProfilesPerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o InlineObject49) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !isNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if !isNil(o.PerSsidSettings) {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject49 struct {
	value *InlineObject49
	isSet bool
}

func (v NullableInlineObject49) Get() *InlineObject49 {
	return v.value
}

func (v *NullableInlineObject49) Set(val *InlineObject49) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject49) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject49) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject49(val *InlineObject49) *NullableInlineObject49 {
	return &NullableInlineObject49{value: val, isSet: true}
}

func (v NullableInlineObject49) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject49) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


