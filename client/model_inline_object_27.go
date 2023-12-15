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

// InlineObject27 struct for InlineObject27
type InlineObject27 struct {
	// The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).
	RfProfileId *string `json:"rfProfileId,omitempty"`
	TwoFourGhzSettings *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *DevicesSerialWirelessRadioSettingsFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
}

// NewInlineObject27 instantiates a new InlineObject27 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject27() *InlineObject27 {
	this := InlineObject27{}
	return &this
}

// NewInlineObject27WithDefaults instantiates a new InlineObject27 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject27WithDefaults() *InlineObject27 {
	this := InlineObject27{}
	return &this
}

// GetRfProfileId returns the RfProfileId field value if set, zero value otherwise.
func (o *InlineObject27) GetRfProfileId() string {
	if o == nil || isNil(o.RfProfileId) {
		var ret string
		return ret
	}
	return *o.RfProfileId
}

// GetRfProfileIdOk returns a tuple with the RfProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject27) GetRfProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.RfProfileId) {
    return nil, false
	}
	return o.RfProfileId, true
}

// HasRfProfileId returns a boolean if a field has been set.
func (o *InlineObject27) HasRfProfileId() bool {
	if o != nil && !isNil(o.RfProfileId) {
		return true
	}

	return false
}

// SetRfProfileId gets a reference to the given string and assigns it to the RfProfileId field.
func (o *InlineObject27) SetRfProfileId(v string) {
	o.RfProfileId = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject27) GetTwoFourGhzSettings() DevicesSerialApplianceRadioSettingsTwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret DevicesSerialApplianceRadioSettingsTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject27) GetTwoFourGhzSettingsOk() (*DevicesSerialApplianceRadioSettingsTwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject27) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given DevicesSerialApplianceRadioSettingsTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject27) SetTwoFourGhzSettings(v DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject27) GetFiveGhzSettings() DevicesSerialWirelessRadioSettingsFiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret DevicesSerialWirelessRadioSettingsFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject27) GetFiveGhzSettingsOk() (*DevicesSerialWirelessRadioSettingsFiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject27) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given DevicesSerialWirelessRadioSettingsFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject27) SetFiveGhzSettings(v DevicesSerialWirelessRadioSettingsFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

func (o InlineObject27) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RfProfileId) {
		toSerialize["rfProfileId"] = o.RfProfileId
	}
	if !isNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !isNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject27 struct {
	value *InlineObject27
	isSet bool
}

func (v NullableInlineObject27) Get() *InlineObject27 {
	return v.value
}

func (v *NullableInlineObject27) Set(val *InlineObject27) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject27) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject27) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject27(val *InlineObject27) *NullableInlineObject27 {
	return &NullableInlineObject27{value: val, isSet: true}
}

func (v NullableInlineObject27) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject27) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


