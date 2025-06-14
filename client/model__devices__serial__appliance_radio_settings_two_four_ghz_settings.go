/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialApplianceRadioSettingsTwoFourGhzSettings Manual radio settings for 2.4 GHz.
type DevicesSerialApplianceRadioSettingsTwoFourGhzSettings struct {
	// Sets a manual channel for 2.4 GHz. Can be '1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12', '13' or '14' or null for using auto channel.
	Channel *int32 `json:"channel,omitempty"`
	// Set a manual target power for 2.4 GHz (dBm). Enter null for using auto power range.
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewDevicesSerialApplianceRadioSettingsTwoFourGhzSettings instantiates a new DevicesSerialApplianceRadioSettingsTwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialApplianceRadioSettingsTwoFourGhzSettings() *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings {
	this := DevicesSerialApplianceRadioSettingsTwoFourGhzSettings{}
	return &this
}

// NewDevicesSerialApplianceRadioSettingsTwoFourGhzSettingsWithDefaults instantiates a new DevicesSerialApplianceRadioSettingsTwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialApplianceRadioSettingsTwoFourGhzSettingsWithDefaults() *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings {
	this := DevicesSerialApplianceRadioSettingsTwoFourGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) GetTargetPower() int32 {
	if o == nil || isNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || isNil(o.TargetPower) {
    return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) HasTargetPower() bool {
	if o != nil && !isNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.TargetPower) {
		toSerialize["targetPower"] = o.TargetPower
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings struct {
	value *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings
	isSet bool
}

func (v NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings) Get() *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings {
	return v.value
}

func (v *NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings) Set(val *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings(val *DevicesSerialApplianceRadioSettingsTwoFourGhzSettings) *NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings {
	return &NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialApplianceRadioSettingsTwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


