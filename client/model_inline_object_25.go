/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 April, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject25 struct for InlineObject25
type InlineObject25 struct {
	// The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).
	RfProfileId *string `json:"rfProfileId,omitempty"`
	TwoFourGhzSettings *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *DevicesSerialWirelessRadioSettingsFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
}

// NewInlineObject25 instantiates a new InlineObject25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject25() *InlineObject25 {
	this := InlineObject25{}
	return &this
}

// NewInlineObject25WithDefaults instantiates a new InlineObject25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject25WithDefaults() *InlineObject25 {
	this := InlineObject25{}
	return &this
}

// GetRfProfileId returns the RfProfileId field value if set, zero value otherwise.
func (o *InlineObject25) GetRfProfileId() string {
	if o == nil || isNil(o.RfProfileId) {
		var ret string
		return ret
	}
	return *o.RfProfileId
}

// GetRfProfileIdOk returns a tuple with the RfProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject25) GetRfProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.RfProfileId) {
    return nil, false
	}
	return o.RfProfileId, true
}

// HasRfProfileId returns a boolean if a field has been set.
func (o *InlineObject25) HasRfProfileId() bool {
	if o != nil && !isNil(o.RfProfileId) {
		return true
	}

	return false
}

// SetRfProfileId gets a reference to the given string and assigns it to the RfProfileId field.
func (o *InlineObject25) SetRfProfileId(v string) {
	o.RfProfileId = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *InlineObject25) GetTwoFourGhzSettings() DevicesSerialWirelessRadioSettingsTwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret DevicesSerialWirelessRadioSettingsTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject25) GetTwoFourGhzSettingsOk() (*DevicesSerialWirelessRadioSettingsTwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *InlineObject25) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given DevicesSerialWirelessRadioSettingsTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *InlineObject25) SetTwoFourGhzSettings(v DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *InlineObject25) GetFiveGhzSettings() DevicesSerialWirelessRadioSettingsFiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret DevicesSerialWirelessRadioSettingsFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject25) GetFiveGhzSettingsOk() (*DevicesSerialWirelessRadioSettingsFiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *InlineObject25) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given DevicesSerialWirelessRadioSettingsFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *InlineObject25) SetFiveGhzSettings(v DevicesSerialWirelessRadioSettingsFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

func (o InlineObject25) MarshalJSON() ([]byte, error) {
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

type NullableInlineObject25 struct {
	value *InlineObject25
	isSet bool
}

func (v NullableInlineObject25) Get() *InlineObject25 {
	return v.value
}

func (v *NullableInlineObject25) Set(val *InlineObject25) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject25) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject25(val *InlineObject25) *NullableInlineObject25 {
	return &NullableInlineObject25{value: val, isSet: true}
}

func (v NullableInlineObject25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


