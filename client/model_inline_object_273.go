/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject273 struct for InlineObject273
type InlineObject273 struct {
	// Boolean for updating SAML SSO enabled settings.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineObject273 instantiates a new InlineObject273 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject273() *InlineObject273 {
	this := InlineObject273{}
	return &this
}

// NewInlineObject273WithDefaults instantiates a new InlineObject273 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject273WithDefaults() *InlineObject273 {
	this := InlineObject273{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject273) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject273) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject273) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject273) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineObject273) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject273 struct {
	value *InlineObject273
	isSet bool
}

func (v NullableInlineObject273) Get() *InlineObject273 {
	return v.value
}

func (v *NullableInlineObject273) Set(val *InlineObject273) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject273) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject273) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject273(val *InlineObject273) *NullableInlineObject273 {
	return &NullableInlineObject273{value: val, isSet: true}
}

func (v NullableInlineObject273) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject273) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


