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

// InlineObject54 struct for InlineObject54
type InlineObject54 struct {
	// The name of the new profile. Must be unique. This param is required on creation.
	Name string `json:"name"`
	TwoFourGhzSettings *NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *NetworksNetworkIdApplianceRfProfilesFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	PerSsidSettings *NetworksNetworkIdApplianceRfProfilesPerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewInlineObject54 instantiates a new InlineObject54 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject54(name string) *InlineObject54 {
	this := InlineObject54{}
	this.Name = name
	return &this
}

// NewInlineObject54WithDefaults instantiates a new InlineObject54 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject54WithDefaults() *InlineObject54 {
	this := InlineObject54{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject54) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject54) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject54) SetName(v string) {
	o.Name = v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject54) GetTwoFourGhzSettings() NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject54) GetTwoFourGhzSettingsOk() (*NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject54) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject54) SetTwoFourGhzSettings(v NetworksNetworkIdApplianceRfProfilesTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject54) GetFiveGhzSettings() NetworksNetworkIdApplianceRfProfilesFiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject54) GetFiveGhzSettingsOk() (*NetworksNetworkIdApplianceRfProfilesFiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject54) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject54) SetFiveGhzSettings(v NetworksNetworkIdApplianceRfProfilesFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *InlineObject54) GetPerSsidSettings() NetworksNetworkIdApplianceRfProfilesPerSsidSettings {
	if o == nil || isNil(o.PerSsidSettings) {
		var ret NetworksNetworkIdApplianceRfProfilesPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject54) GetPerSsidSettingsOk() (*NetworksNetworkIdApplianceRfProfilesPerSsidSettings, bool) {
	if o == nil || isNil(o.PerSsidSettings) {
    return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *InlineObject54) HasPerSsidSettings() bool {
	if o != nil && !isNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given NetworksNetworkIdApplianceRfProfilesPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *InlineObject54) SetPerSsidSettings(v NetworksNetworkIdApplianceRfProfilesPerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o InlineObject54) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableInlineObject54 struct {
	value *InlineObject54
	isSet bool
}

func (v NullableInlineObject54) Get() *InlineObject54 {
	return v.value
}

func (v *NullableInlineObject54) Set(val *InlineObject54) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject54) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject54) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject54(val *InlineObject54) *NullableInlineObject54 {
	return &NullableInlineObject54{value: val, isSet: true}
}

func (v NullableInlineObject54) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject54) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


