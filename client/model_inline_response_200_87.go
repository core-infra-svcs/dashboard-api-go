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

// InlineResponse20087 struct for InlineResponse20087
type InlineResponse20087 struct {
	DefaultSettings *InlineResponse20087DefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings.       An empty array will clear the multicast settings.
	Overrides []InlineResponse20087Overrides `json:"overrides,omitempty"`
}

// NewInlineResponse20087 instantiates a new InlineResponse20087 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20087() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// NewInlineResponse20087WithDefaults instantiates a new InlineResponse20087 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20087WithDefaults() *InlineResponse20087 {
	this := InlineResponse20087{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *InlineResponse20087) GetDefaultSettings() InlineResponse20087DefaultSettings {
	if o == nil || isNil(o.DefaultSettings) {
		var ret InlineResponse20087DefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetDefaultSettingsOk() (*InlineResponse20087DefaultSettings, bool) {
	if o == nil || isNil(o.DefaultSettings) {
    return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *InlineResponse20087) HasDefaultSettings() bool {
	if o != nil && !isNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given InlineResponse20087DefaultSettings and assigns it to the DefaultSettings field.
func (o *InlineResponse20087) SetDefaultSettings(v InlineResponse20087DefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *InlineResponse20087) GetOverrides() []InlineResponse20087Overrides {
	if o == nil || isNil(o.Overrides) {
		var ret []InlineResponse20087Overrides
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20087) GetOverridesOk() ([]InlineResponse20087Overrides, bool) {
	if o == nil || isNil(o.Overrides) {
    return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *InlineResponse20087) HasOverrides() bool {
	if o != nil && !isNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []InlineResponse20087Overrides and assigns it to the Overrides field.
func (o *InlineResponse20087) SetOverrides(v []InlineResponse20087Overrides) {
	o.Overrides = v
}

func (o InlineResponse20087) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !isNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20087 struct {
	value *InlineResponse20087
	isSet bool
}

func (v NullableInlineResponse20087) Get() *InlineResponse20087 {
	return v.value
}

func (v *NullableInlineResponse20087) Set(val *InlineResponse20087) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20087) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20087) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20087(val *InlineResponse20087) *NullableInlineResponse20087 {
	return &NullableInlineResponse20087{value: val, isSet: true}
}

func (v NullableInlineResponse20087) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20087) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


