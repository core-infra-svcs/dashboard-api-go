/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialApplianceRadioSettingsFiveGhzSettings Manual radio settings for 5 GHz.
type DevicesSerialApplianceRadioSettingsFiveGhzSettings struct {
	// Sets a manual channel for 5 GHz. Can be '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161', '165', '169', '173' or '177' or null for using auto channel.
	Channel *int32 `json:"channel,omitempty"`
	// Sets a manual channel width for 5 GHz. Can be '0', '20', '40', '80' or '160' or null for using auto channel width.
	ChannelWidth *int32 `json:"channelWidth,omitempty"`
	// Set a manual target power for 5 GHz. Can be between '8' or '30' or null for using auto power range.
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewDevicesSerialApplianceRadioSettingsFiveGhzSettings instantiates a new DevicesSerialApplianceRadioSettingsFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialApplianceRadioSettingsFiveGhzSettings() *DevicesSerialApplianceRadioSettingsFiveGhzSettings {
	this := DevicesSerialApplianceRadioSettingsFiveGhzSettings{}
	return &this
}

// NewDevicesSerialApplianceRadioSettingsFiveGhzSettingsWithDefaults instantiates a new DevicesSerialApplianceRadioSettingsFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialApplianceRadioSettingsFiveGhzSettingsWithDefaults() *DevicesSerialApplianceRadioSettingsFiveGhzSettings {
	this := DevicesSerialApplianceRadioSettingsFiveGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) GetChannelWidth() int32 {
	if o == nil || isNil(o.ChannelWidth) {
		var ret int32
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) GetChannelWidthOk() (*int32, bool) {
	if o == nil || isNil(o.ChannelWidth) {
    return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) HasChannelWidth() bool {
	if o != nil && !isNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given int32 and assigns it to the ChannelWidth field.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) SetChannelWidth(v int32) {
	o.ChannelWidth = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) GetTargetPower() int32 {
	if o == nil || isNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || isNil(o.TargetPower) {
    return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) HasTargetPower() bool {
	if o != nil && !isNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *DevicesSerialApplianceRadioSettingsFiveGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o DevicesSerialApplianceRadioSettingsFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !isNil(o.TargetPower) {
		toSerialize["targetPower"] = o.TargetPower
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings struct {
	value *DevicesSerialApplianceRadioSettingsFiveGhzSettings
	isSet bool
}

func (v NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings) Get() *DevicesSerialApplianceRadioSettingsFiveGhzSettings {
	return v.value
}

func (v *NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings) Set(val *DevicesSerialApplianceRadioSettingsFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialApplianceRadioSettingsFiveGhzSettings(val *DevicesSerialApplianceRadioSettingsFiveGhzSettings) *NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings {
	return &NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings{value: val, isSet: true}
}

func (v NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialApplianceRadioSettingsFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


