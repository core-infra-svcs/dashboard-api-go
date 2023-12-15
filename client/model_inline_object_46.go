/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject46 struct for InlineObject46
type InlineObject46 struct {
	// The name of the new profile. Must be unique.
	Name *string `json:"name,omitempty"`
	TwoFourGhzSettings *NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	PerSsidSettings *NetworksNetworkIdApplianceRfProfilesPerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewInlineObject46 instantiates a new InlineObject46 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject46() *InlineObject46 {
	this := InlineObject46{}
	return &this
}

// NewInlineObject46WithDefaults instantiates a new InlineObject46 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject46WithDefaults() *InlineObject46 {
	this := InlineObject46{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject46) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject46) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject46) SetName(v string) {
	o.Name = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject46) GetTwoFourGhzSettings() NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetTwoFourGhzSettingsOk() (*NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject46) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject46) SetTwoFourGhzSettings(v NetworksNetworkIdApplianceRfProfilesRfProfileIdTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject46) GetFiveGhzSettings() NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetFiveGhzSettingsOk() (*NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject46) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject46) SetFiveGhzSettings(v NetworksNetworkIdApplianceRfProfilesRfProfileIdFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineObject46) GetPerSsidSettings() NetworksNetworkIdApplianceRfProfilesPerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetPerSsidSettingsOk() (*NetworksNetworkIdApplianceRfProfilesPerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineObject46) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineObject46) SetPerSsidSettings(v NetworksNetworkIdApplianceRfProfilesPerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o InlineObject46) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject46 struct {
	value *InlineObject46
	isSet bool
}

func (v NullableInlineObject46) Get() *InlineObject46 {
	return v.value
}

func (v *NullableInlineObject46) Set(val *InlineObject46) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject46) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject46) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject46(val *InlineObject46) *NullableInlineObject46 {
	return &NullableInlineObject46{value: val, isSet: true}
}

func (v NullableInlineObject46) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject46) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


