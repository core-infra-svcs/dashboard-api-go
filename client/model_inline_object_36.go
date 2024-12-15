/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject36 struct for InlineObject36
type InlineObject36 struct {
	// Desired ESL channel for the device, or 'Auto' (case insensitive) to use the recommended channel
	Channel *string `json:"channel,omitempty"`
	// Turn ESL features on and off for this device
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineObject36 instantiates a new InlineObject36 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject36() *InlineObject36 {
	this := InlineObject36{}
	return &this
}

// NewInlineObject36WithDefaults instantiates a new InlineObject36 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject36WithDefaults() *InlineObject36 {
	this := InlineObject36{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineObject36) GetChannel() string {
	if o == nil || isNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject36) GetChannelOk() (*string, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineObject36) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *InlineObject36) SetChannel(v string) {
	o.Channel = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject36) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject36) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject36) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject36) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineObject36) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject36 struct {
	value *InlineObject36
	isSet bool
}

func (v NullableInlineObject36) Get() *InlineObject36 {
	return v.value
}

func (v *NullableInlineObject36) Set(val *InlineObject36) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject36) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject36) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject36(val *InlineObject36) *NullableInlineObject36 {
	return &NullableInlineObject36{value: val, isSet: true}
}

func (v NullableInlineObject36) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject36) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


