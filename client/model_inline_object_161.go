/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject161 struct for InlineObject161
type InlineObject161 struct {
	// Optional boolean to retain all the current configs given by the template.
	RetainConfigs *bool `json:"retainConfigs,omitempty"`
}

// NewInlineObject161 instantiates a new InlineObject161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject161() *InlineObject161 {
	this := InlineObject161{}
	return &this
}

// NewInlineObject161WithDefaults instantiates a new InlineObject161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject161WithDefaults() *InlineObject161 {
	this := InlineObject161{}
	return &this
}

// GetRetainConfigs returns the RetainConfigs field value if set, zero value otherwise.
func (o *InlineObject161) GetRetainConfigs() bool {
	if o == nil || isNil(o.RetainConfigs) {
		var ret bool
		return ret
	}
	return *o.RetainConfigs
}

// GetRetainConfigsOk returns a tuple with the RetainConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject161) GetRetainConfigsOk() (*bool, bool) {
	if o == nil || isNil(o.RetainConfigs) {
    return nil, false
	}
	return o.RetainConfigs, true
}

// HasRetainConfigs returns a boolean if a field has been set.
func (o *InlineObject161) HasRetainConfigs() bool {
	if o != nil && !isNil(o.RetainConfigs) {
		return true
	}

	return false
}

// SetRetainConfigs gets a reference to the given bool and assigns it to the RetainConfigs field.
func (o *InlineObject161) SetRetainConfigs(v bool) {
	o.RetainConfigs = &v
}

func (o InlineObject161) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RetainConfigs) {
		toSerialize["retainConfigs"] = o.RetainConfigs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject161 struct {
	value *InlineObject161
	isSet bool
}

func (v NullableInlineObject161) Get() *InlineObject161 {
	return v.value
}

func (v *NullableInlineObject161) Set(val *InlineObject161) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject161) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject161(val *InlineObject161) *NullableInlineObject161 {
	return &NullableInlineObject161{value: val, isSet: true}
}

func (v NullableInlineObject161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


