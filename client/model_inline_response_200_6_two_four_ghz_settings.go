/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2006TwoFourGhzSettings Manual radio settings for 2.4 GHz
type InlineResponse2006TwoFourGhzSettings struct {
	// Manual channel for 2.4 GHz
	Channel *int32 `json:"channel,omitempty"`
	// Manual target power for 2.4 GHz
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewInlineResponse2006TwoFourGhzSettings instantiates a new InlineResponse2006TwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2006TwoFourGhzSettings() *InlineResponse2006TwoFourGhzSettings {
	this := InlineResponse2006TwoFourGhzSettings{}
	return &this
}

// NewInlineResponse2006TwoFourGhzSettingsWithDefaults instantiates a new InlineResponse2006TwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2006TwoFourGhzSettingsWithDefaults() *InlineResponse2006TwoFourGhzSettings {
	this := InlineResponse2006TwoFourGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineResponse2006TwoFourGhzSettings) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006TwoFourGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineResponse2006TwoFourGhzSettings) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *InlineResponse2006TwoFourGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *InlineResponse2006TwoFourGhzSettings) GetTargetPower() int32 {
	if o == nil || isNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2006TwoFourGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || isNil(o.TargetPower) {
    return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *InlineResponse2006TwoFourGhzSettings) HasTargetPower() bool {
	if o != nil && !isNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *InlineResponse2006TwoFourGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o InlineResponse2006TwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.TargetPower) {
		toSerialize["targetPower"] = o.TargetPower
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2006TwoFourGhzSettings struct {
	value *InlineResponse2006TwoFourGhzSettings
	isSet bool
}

func (v NullableInlineResponse2006TwoFourGhzSettings) Get() *InlineResponse2006TwoFourGhzSettings {
	return v.value
}

func (v *NullableInlineResponse2006TwoFourGhzSettings) Set(val *InlineResponse2006TwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2006TwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2006TwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2006TwoFourGhzSettings(val *InlineResponse2006TwoFourGhzSettings) *NullableInlineResponse2006TwoFourGhzSettings {
	return &NullableInlineResponse2006TwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableInlineResponse2006TwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2006TwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


