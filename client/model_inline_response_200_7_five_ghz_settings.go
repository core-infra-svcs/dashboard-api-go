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

// InlineResponse2007FiveGhzSettings Manual radio settings for 5 GHz
type InlineResponse2007FiveGhzSettings struct {
	// Manual channel for 5 GHz
	Channel *int32 `json:"channel,omitempty"`
	// Manual channel width for 5 GHz
	ChannelWidth *int32 `json:"channelWidth,omitempty"`
	// Manual target power for 5 GHz
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewInlineResponse2007FiveGhzSettings instantiates a new InlineResponse2007FiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007FiveGhzSettings() *InlineResponse2007FiveGhzSettings {
	this := InlineResponse2007FiveGhzSettings{}
	return &this
}

// NewInlineResponse2007FiveGhzSettingsWithDefaults instantiates a new InlineResponse2007FiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007FiveGhzSettingsWithDefaults() *InlineResponse2007FiveGhzSettings {
	this := InlineResponse2007FiveGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineResponse2007FiveGhzSettings) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007FiveGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineResponse2007FiveGhzSettings) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *InlineResponse2007FiveGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *InlineResponse2007FiveGhzSettings) GetChannelWidth() int32 {
	if o == nil || isNil(o.ChannelWidth) {
		var ret int32
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007FiveGhzSettings) GetChannelWidthOk() (*int32, bool) {
	if o == nil || isNil(o.ChannelWidth) {
    return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *InlineResponse2007FiveGhzSettings) HasChannelWidth() bool {
	if o != nil && !isNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given int32 and assigns it to the ChannelWidth field.
func (o *InlineResponse2007FiveGhzSettings) SetChannelWidth(v int32) {
	o.ChannelWidth = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *InlineResponse2007FiveGhzSettings) GetTargetPower() int32 {
	if o == nil || isNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007FiveGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || isNil(o.TargetPower) {
    return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *InlineResponse2007FiveGhzSettings) HasTargetPower() bool {
	if o != nil && !isNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *InlineResponse2007FiveGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o InlineResponse2007FiveGhzSettings) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2007FiveGhzSettings struct {
	value *InlineResponse2007FiveGhzSettings
	isSet bool
}

func (v NullableInlineResponse2007FiveGhzSettings) Get() *InlineResponse2007FiveGhzSettings {
	return v.value
}

func (v *NullableInlineResponse2007FiveGhzSettings) Set(val *InlineResponse2007FiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007FiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007FiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007FiveGhzSettings(val *InlineResponse2007FiveGhzSettings) *NullableInlineResponse2007FiveGhzSettings {
	return &NullableInlineResponse2007FiveGhzSettings{value: val, isSet: true}
}

func (v NullableInlineResponse2007FiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007FiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


