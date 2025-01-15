/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200158 struct for InlineResponse200158
type InlineResponse200158 struct {
	DefaultSettings *InlineResponse200158DefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings.
	Overrides []InlineResponse200158Overrides `json:"overrides,omitempty"`
}

// NewInlineResponse200158 instantiates a new InlineResponse200158 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200158() *InlineResponse200158 {
	this := InlineResponse200158{}
	return &this
}

// NewInlineResponse200158WithDefaults instantiates a new InlineResponse200158 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200158WithDefaults() *InlineResponse200158 {
	this := InlineResponse200158{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *InlineResponse200158) GetDefaultSettings() InlineResponse200158DefaultSettings {
	if o == nil || isNil(o.DefaultSettings) {
		var ret InlineResponse200158DefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200158) GetDefaultSettingsOk() (*InlineResponse200158DefaultSettings, bool) {
	if o == nil || isNil(o.DefaultSettings) {
    return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *InlineResponse200158) HasDefaultSettings() bool {
	if o != nil && !isNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given InlineResponse200158DefaultSettings and assigns it to the DefaultSettings field.
func (o *InlineResponse200158) SetDefaultSettings(v InlineResponse200158DefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineResponse200158) GetOverrides() []InlineResponse200158Overrides {
	if o == nil || isNil(o.Overrides) {
		var ret []InlineResponse200158Overrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200158) GetOverridesOk() ([]InlineResponse200158Overrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineResponse200158) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []InlineResponse200158Overrides and assigns it to the Overrides field.
func (o *InlineResponse200158) SetOverrides(v []InlineResponse200158Overrides) {
	o.Overrides = v
}

func (o InlineResponse200158) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200158 struct {
	value *InlineResponse200158
	isSet bool
}

func (v NullableInlineResponse200158) Get() *InlineResponse200158 {
	return v.value
}

func (v *NullableInlineResponse200158) Set(val *InlineResponse200158) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200158) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200158) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200158(val *InlineResponse200158) *NullableInlineResponse200158 {
	return &NullableInlineResponse200158{value: val, isSet: true}
}

func (v NullableInlineResponse200158) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200158) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


